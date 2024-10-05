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
	"google.golang.org/grpc/credentials"
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
	cfgMu           sync.RWMutex

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

func (s *HakjServer) Config() config.ServerConfig {
	s.cfgMu.RLock()
	cfg := s.Cfg
	s.cfgMu.RUnlock()
	return cfg
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
	cfg := s.Config()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return cfg.DefaultDB
	}

	dbName := md.Get(common.GrpcMetadataKeyDbName)
	if len(dbName) < 1 {
		return cfg.DefaultDB
	}

	return dbName[0]
}

// EnableLogFile enables logger to write logs to the log file.
func (s *HakjServer) EnableLogFile() {
	s.loggerMu.Lock()
	defer s.loggerMu.Unlock()
	err := s.logger.EnableLogFile(s.Cfg.LogFilePath)
	if err != nil {
		s.logger.Fatalf("Failed to enable log file: %v", err)
	}
}

// CloseLogger closes logger and releases its possible resources.
func (s *HakjServer) CloseLogger() {
	s.loggerMu.Lock()
	defer s.loggerMu.Unlock()
	err := s.logger.CloseLogFile()
	if err != nil {
		s.logger.Fatalf("Failed to close log file: %v", err)
	}
}

// EnableAuth enables authentication.
func (s *HakjServer) EnableAuth(rootPassword string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.credentialStore.SetPassword(auth.RootUserName, []byte(rootPassword)); err != nil {
		s.logger.Fatalf("Failed to set root password: %v", err)
	}
	s.logger.Info("Authentication is enabled. Clients need to authenticate")
	if rootPassword == "" {
		s.logger.Warning("Using empty password. Consider changing it to a strong password")
	}
}

func (s *HakjServer) DisableAuth() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.credentialStore.SetPassword(auth.RootUserName, []byte("")); err != nil {
		s.logger.Fatalf("Failed to clear root password: %v", err)
	}
	s.logger.Warning("Authentication is disabled")
}

// CreateDefaultDatabase creates an empty default database.
func (s *HakjServer) CreateDefaultDatabase(name string) {
	cfg := s.Config()
	if err := validation.ValidateDBName(name); err != nil {
		s.logger.Fatalf("Failed to create default database: %v", err)
	}
	dbConfig := hakjdb.DBConfig{MaxHashMapFields: cfg.MaxHashMapFields}
	db := hakjdb.NewDB(name, "", dbConfig)
	s.dbs[db.Name()] = db
	s.logger.Infof("Created default database '%s'", db.Name())
}

// DBMaxKeysReached returns true if a database has reached or exceeded the maximum key limit.
func (s *HakjServer) DBMaxKeysReached(db *hakjdb.DB) bool {
	cfg := s.Config()
	return uint32(db.GetKeyCount()) >= cfg.MaxKeysPerDB
}

// Init initializes the server.
func (s *HakjServer) Init() {
	s.logger.Info("Initializing server ...")
	cfg := s.Config()

	if cfg.LogFileEnabled {
		s.EnableLogFile()
		s.logger.Infof("Log file is enabled. Logs will be written to the log file. The file is located at %s", cfg.LogFilePath)
	}

	if cfg.DebugEnabled {
		s.logger.Info("Debug mode is enabled")
	}

	if cfg.AuthEnabled {
		s.logger.Info("Enabling authentication")
		password, _ := config.GetPassword()
		s.EnableAuth(password)
	} else {
		s.logger.Warning("Authentication is disabled")
	}

	s.CreateDefaultDatabase(cfg.DefaultDB)
}

func (s *HakjServer) GetTLSCredentials() credentials.TransportCredentials {
	logger := s.Logger()
	cfg := s.Config()
	serverCert, err := tls.LoadX509KeyPair(cfg.TLSCertPath, cfg.TLSPrivKeyPath)
	if err != nil {
		logger.Fatalf("Failed to load TLS private/public key pair: %v", err)
	}

	clientAuth := tls.NoClientCert
	certPool := x509.NewCertPool()
	if cfg.TLSClientCertAuthEnabled {
		caCert, err := os.ReadFile(cfg.TLSCACertPath)
		if err != nil {
			logger.Fatalf("Failed to read TLS CA certificate: %v", err)
		}
		if !certPool.AppendCertsFromPEM(caCert) {
			logger.Fatal("Failed to parse TLS CA certificate")
		}
		clientAuth = tls.RequireAndVerifyClientCert
		logger.Infof("Using client certificate authentication in TLS. Clients are required to send a client certificate signed by the server's root CA certificate")
	}

	return credentials.NewTLS(
		&tls.Config{
			ClientAuth:   clientAuth,
			Certificates: []tls.Certificate{serverCert},
			ClientCAs:    certPool,
		})
}

func (s *HakjServer) SetupListener() {
	logger := s.Logger()
	cfg := s.Config()
	logger.Info("Setting up listener ...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.PortInUse))
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}
	logger.Infof("Server listening at %v", lis.Addr())

	connListener := NewClientConnListener(lis, s, cfg.MaxClientConnections)
	s.ClientConnListener = connListener
}

func (s *HakjServer) ProcessConfigReload(cfg *config.ServerConfig) {
	s.loggerMu.Lock()
	defer s.loggerMu.Unlock()

	logLevel, logLevelStr, ok := hakjdb.GetLogLevelFromStr(config.GetLogLevelStr())
	if !ok {
		s.logger.Warning("Invalid log level configured. Default log level will be used")
	}
	s.logger.Infof("Using log level %s", logLevelStr)
	// modify the original
	s.logger.SetLogLevel(logLevel)

	if cfg.VerboseLogsEnabled {
		s.logger.Info("Verbose logs are enabled")
	}

	if cfg.LogFileEnabled {
		err := s.logger.CloseLogFile()
		if err != nil {
			s.logger.Fatalf("Failed to close log file: %v", err)
		}
		err = s.logger.EnableLogFile(cfg.LogFilePath)
		if err != nil {
			s.logger.Fatalf("Failed to enable log file: %v", err)
		}
		s.logger.Infof("Log file is enabled. Logs will be written to the log file. The file is located at %s", cfg.LogFilePath)
	}

	if cfg.AuthEnabled {
		s.logger.Info("Enabling authentication")
		password, _ := config.GetPassword()
		s.EnableAuth(password)
	} else {
		s.logger.Info("Disabling authentication")
		s.DisableAuth()
	}

	s.ClientConnListener.mu.Lock()
	s.ClientConnListener.maxClientConnections = cfg.MaxClientConnections
	s.ClientConnListener.mu.Unlock()
}
