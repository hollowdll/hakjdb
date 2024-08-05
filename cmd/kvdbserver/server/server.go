package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	kvdb "github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/auth"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/config"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/validation"
	kvdberrors "github.com/hollowdll/kvdb/errors"
	"github.com/hollowdll/kvdb/internal/common"
	"google.golang.org/grpc/metadata"
)

// ClientConnListener is a client connection listener
// that accepts new connections and tracks active connections.
type ClientConnListener struct {
	net.Listener
	server                 *KvdbServer
	clientConnectionsCount uint32
	maxClientConnections   uint32
	mu                     sync.RWMutex
}

func NewClientConnListener(lis net.Listener, s *KvdbServer, maxConnections uint32) *ClientConnListener {
	return &ClientConnListener{
		Listener:               lis,
		server:                 s,
		clientConnectionsCount: 0,
		maxClientConnections:   maxConnections,
	}
}

func (l *ClientConnListener) Accept() (net.Conn, error) {
	conn, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}

	logger := l.server.Logger()
	l.mu.Lock()
	l.clientConnectionsCount++
	logger.Debugf("Client connected, total clients: %d\n", l.clientConnectionsCount)

	clientConn := &clientConn{Conn: conn, release: func() {
		l.mu.Lock()
		if l.clientConnectionsCount > 0 {
			l.clientConnectionsCount--
		}
		logger.Debugf("Client disconnected, total clients: %d\n", l.clientConnectionsCount)
		l.mu.Unlock()
	}}

	if l.clientConnectionsCount > l.maxClientConnections {
		logger.Errorf("Incoming connection denied: %s", kvdberrors.ErrMaxClientConnectionsReached.Error())
		l.mu.Unlock()
		clientConn.Close()
	} else {
		l.mu.Unlock()
	}

	return clientConn, nil
}

type clientConn struct {
	net.Conn
	release func()
	closed  bool
	mu      sync.Mutex
}

func (c *clientConn) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil
	}

	c.closed = true
	err := c.Conn.Close()
	c.release()
	return err
}

type KvdbServer struct {
	startTime time.Time

	// dbs holds the databases and their names that exist on the server.
	dbs map[string]*kvdb.DB

	credentialStore auth.CredentialStore
	logger          kvdb.Logger
	loggerMu        sync.RWMutex

	// Cfg is the configuration that the server is configured with.
	// It is not intended to be changed after the server has been set up.
	Cfg config.ServerConfig

	*ClientConnListener
	mu sync.RWMutex
}

func NewKvdbServer(cfg config.ServerConfig, lg kvdb.Logger) *KvdbServer {
	return &KvdbServer{
		startTime:          time.Now(),
		dbs:                make(map[string]*kvdb.DB),
		credentialStore:    auth.NewInMemoryCredentialStore(),
		logger:             lg,
		Cfg:                cfg,
		ClientConnListener: nil,
	}
}

func (s *KvdbServer) Logger() kvdb.Logger {
	s.loggerMu.RLock()
	l := s.logger
	s.loggerMu.RUnlock()
	return l
}

// totalStoredDataSize returns the total amount of stored data on this server in bytes.
func (s *KvdbServer) totalStoredDataSize() uint64 {
	var sum uint64
	for _, db := range s.dbs {
		sum += db.GetEstimatedStorageSizeBytes()
	}

	return sum
}

// dbExists returns true if a database with the given name exists on the server.
func (s *KvdbServer) dbExists(name string) bool {
	_, exists := s.dbs[name]
	return exists
}

// GetDBNameFromContext gets the database name from the incoming context gRPC metadata.
func (s *KvdbServer) GetDBNameFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return s.Cfg.DefaultDB
	}

	dbName := md.Get(common.GrpcMetadataKeyDbName)
	if len(dbName) < 1 {
		return s.Cfg.DefaultDB
	}

	return dbName[0]
}

// DisableLogger disables all log outputs from this server.
func (s *KvdbServer) DisableLogger() {
	s.logger.Disable()
}

// ActivateDebugMode activates debug mode.
func (s *KvdbServer) ActivateDebugMode() {
	s.logger.EnableDebug()
}

// EnableLogFile enables logger to write logs to the log file.
func (s *KvdbServer) EnableLogFile() {
	err := s.logger.EnableLogFile(s.Cfg.LogFilePath)
	if err != nil {
		s.logger.Fatalf("Failed to enable log file: %v", err)
	}
}

// CloseLogger closes logger and releases its possible resources.
func (s *KvdbServer) CloseLogger() {
	s.loggerMu.RLock()
	defer s.loggerMu.RUnlock()
	err := s.logger.CloseLogFile()
	if err != nil {
		s.logger.Fatalf("Failed to close log file: %v", err)
	}
}

// EnablePasswordProtection enables server password protection and sets the password.
func (s *KvdbServer) EnablePasswordProtection(password string) {
	if err := s.credentialStore.SetServerPassword([]byte(password)); err != nil {
		s.logger.Fatalf("Failed to set server password: %v", err)
	}
	s.logger.Infof("Password protection is enabled. Clients need to authenticate using password.")
}

// CreateDefaultDatabase creates an empty default database.
func (s *KvdbServer) CreateDefaultDatabase(name string) {
	if err := validation.ValidateDBName(name); err != nil {
		s.logger.Fatalf("Failed to create default database: %v", err)
	}
	dbConfig := kvdb.DBConfig{MaxHashMapFields: s.Cfg.MaxHashMapFields}
	db := kvdb.NewDB(name, "", dbConfig)
	s.dbs[db.Name()] = db
	s.logger.Infof("Created default database '%s'", db.Name())
}

// DBMaxKeysReached returns true if a database has reached or exceeded the maximum key limit.
func (s *KvdbServer) DBMaxKeysReached(db *kvdb.DB) bool {
	return db.GetKeyCount() >= s.Cfg.MaxKeysPerDB
}

// Init initializes the server.
func (s *KvdbServer) Init() {
	s.logger.Infof("Initializing server ...")

	if s.Cfg.LogFileEnabled {
		s.EnableLogFile()
		s.logger.Infof("Log file is enabled. Logs will be written to the log file. The file is located at %s", s.Cfg.LogFilePath)
	}

	if s.Cfg.DebugEnabled {
		s.ActivateDebugMode()
		s.logger.Info("Debug mode is enabled")
	}

	password, ok := config.ShouldUsePassword()
	if ok {
		s.EnablePasswordProtection(password)
	} else {
		s.logger.Warning("Password protection is disabled")
	}

	s.CreateDefaultDatabase(s.Cfg.DefaultDB)
}

func (s *KvdbServer) GetTLSCert() tls.Certificate {
	logger := s.Logger()
	certBytes, err := os.ReadFile(s.Cfg.TLSCertPath)
	if err != nil {
		logger.Fatalf("Failed to read TLS certificate: %v", err)
	}
	privKeyBytes, err := os.ReadFile(s.Cfg.TLSPrivKeyPath)
	if err != nil {
		logger.Fatalf("Failed to read TLS private key: %v", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(certBytes) {
		logger.Fatal("Failed to parse TLS certificate")
	}
	cert, err := tls.X509KeyPair(certBytes, privKeyBytes)
	if err != nil {
		logger.Fatalf("Failed to parse TLS public/private key pair: %v", err)
	}

	return cert
}

func (s *KvdbServer) SetupListener() {
	logger := s.Logger()
	logger.Infof("Setting up listener ...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Cfg.PortInUse))
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}
	logger.Infof("Server listening at %v", lis.Addr())

	connListener := NewClientConnListener(lis, s, s.Cfg.MaxClientConnections)
	s.ClientConnListener = connListener
}
