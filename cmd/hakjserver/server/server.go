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

	"github.com/hollowdll/hakjdb"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/auth"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/config"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/validation"
	hakjerrors "github.com/hollowdll/hakjdb/errors"
	"github.com/hollowdll/hakjdb/internal/common"
	"google.golang.org/grpc/metadata"
)

// ClientConnListener is a client connection listener
// that accepts new connections and tracks active connections.
type ClientConnListener struct {
	net.Listener
	server *HakjServer

	// clientConnectionsCount is the current number of client connections.
	clientConnectionsCount uint32
	// maxClientConnections is the maximum number of concurrent client connections allowed.
	maxClientConnections uint32
	// totalClientConnections is the number of total client connections created.
	totalClientConnections uint64

	mu sync.RWMutex
}

func NewClientConnListener(lis net.Listener, s *HakjServer, maxConnections uint32) *ClientConnListener {
	return &ClientConnListener{
		Listener:               lis,
		server:                 s,
		clientConnectionsCount: 0,
		maxClientConnections:   maxConnections,
		totalClientConnections: 0,
	}
}

func (l *ClientConnListener) Accept() (net.Conn, error) {
	conn, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}

	l.mu.Lock()
	logger := l.server.Logger()
	l.clientConnectionsCount++
	l.totalClientConnections++
	id := l.totalClientConnections
	logger.Debugf("Client ID %d connected, currently connected clients: %d\n", id, l.clientConnectionsCount)

	clientConn := &clientConn{Conn: conn, id: id, release: func() {
		l.mu.Lock()
		if l.clientConnectionsCount > 0 {
			l.clientConnectionsCount--
		}
		logger.Debugf("Client ID %d disconnected, currently connected clients: %d\n", id, l.clientConnectionsCount)
		l.mu.Unlock()
	}}

	if l.clientConnectionsCount > l.maxClientConnections {
		logger.Errorf("Incoming connection denied: %s", hakjerrors.ErrMaxClientConnectionsReached.Error())
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
	id      uint64
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

type HakjServer struct {
	startTime time.Time

	// dbs holds the databases and their names that exist on the server.
	dbs map[string]*hakjdb.DB

	credentialStore auth.CredentialStore
	logger          hakjdb.Logger
	loggerMu        sync.RWMutex

	// Cfg is the configuration that the server is configured with.
	// It is not intended to be changed after the server has been set up.
	Cfg config.ServerConfig

	*ClientConnListener
	mu sync.RWMutex
}

func NewHakjServer(cfg config.ServerConfig, lg hakjdb.Logger) *HakjServer {
	return &HakjServer{
		startTime:          time.Now(),
		dbs:                make(map[string]*hakjdb.DB),
		credentialStore:    auth.NewInMemoryCredentialStore(),
		logger:             lg,
		Cfg:                cfg,
		ClientConnListener: nil,
	}
}

func (s *HakjServer) Logger() hakjdb.Logger {
	s.loggerMu.RLock()
	l := s.logger
	s.loggerMu.RUnlock()
	return l
}

// totalStoredDataSize returns the total amount of stored data on this server in bytes.
func (s *HakjServer) totalStoredDataSize() uint64 {
	var sum uint64
	for _, db := range s.dbs {
		sum += db.GetEstimatedStorageSizeBytes()
	}

	return sum
}

// dbExists returns true if a database with the given name exists on the server.
func (s *HakjServer) dbExists(name string) bool {
	_, exists := s.dbs[name]
	return exists
}

// GetDBNameFromContext gets the database name from the incoming context gRPC metadata.
func (s *HakjServer) GetDBNameFromContext(ctx context.Context) string {
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

// EnableLogFile enables logger to write logs to the log file.
func (s *HakjServer) EnableLogFile() {
	s.loggerMu.RLock()
	defer s.loggerMu.RUnlock()
	err := s.logger.EnableLogFile(s.Cfg.LogFilePath)
	if err != nil {
		s.logger.Fatalf("Failed to enable log file: %v", err)
	}
}

// CloseLogger closes logger and releases its possible resources.
func (s *HakjServer) CloseLogger() {
	s.loggerMu.RLock()
	defer s.loggerMu.RUnlock()
	err := s.logger.CloseLogFile()
	if err != nil {
		s.logger.Fatalf("Failed to close log file: %v", err)
	}
}

// EnableAuth enables authentication.
func (s *HakjServer) EnableAuth(rootPassword string) {
	if err := s.credentialStore.SetPassword(auth.RootUserName, []byte(rootPassword)); err != nil {
		s.logger.Fatalf("Failed to set root password: %v", err)
	}
	s.logger.Infof("Authentication is enabled. Clients need to authenticate.")
	if rootPassword == "" {
		s.logger.Warning("Using empty password. Consider changing it to a strong password.")
	}
}

// CreateDefaultDatabase creates an empty default database.
func (s *HakjServer) CreateDefaultDatabase(name string) {
	if err := validation.ValidateDBName(name); err != nil {
		s.logger.Fatalf("Failed to create default database: %v", err)
	}
	dbConfig := hakjdb.DBConfig{MaxHashMapFields: s.Cfg.MaxHashMapFields}
	db := hakjdb.NewDB(name, "", dbConfig)
	s.dbs[db.Name()] = db
	s.logger.Infof("Created default database '%s'", db.Name())
}

// DBMaxKeysReached returns true if a database has reached or exceeded the maximum key limit.
func (s *HakjServer) DBMaxKeysReached(db *hakjdb.DB) bool {
	return uint32(db.GetKeyCount()) >= s.Cfg.MaxKeysPerDB
}

// Init initializes the server.
func (s *HakjServer) Init() {
	s.logger.Infof("Initializing server ...")

	if s.Cfg.LogFileEnabled {
		s.EnableLogFile()
		s.logger.Infof("Log file is enabled. Logs will be written to the log file. The file is located at %s", s.Cfg.LogFilePath)
	}

	if s.Cfg.DebugEnabled {
		s.logger.Info("Debug mode is enabled")
	}

	if s.Cfg.AuthEnabled {
		s.logger.Info("Enabling authentication")
		password, _ := config.ShouldUsePassword()
		s.EnableAuth(password)
	} else {
		s.logger.Warning("Authentication is disabled")
	}

	s.CreateDefaultDatabase(s.Cfg.DefaultDB)
}

func (s *HakjServer) GetTLSCert() tls.Certificate {
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

func (s *HakjServer) SetupListener() {
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
