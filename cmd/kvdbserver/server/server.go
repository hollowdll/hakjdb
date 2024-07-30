package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"runtime"
	"sync"
	"time"

	kvdb "github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/api/v0/dbpb"
	"github.com/hollowdll/kvdb/api/v0/serverpb"
	"github.com/hollowdll/kvdb/api/v0/storagepb"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/auth"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/config"
	kvdberrors "github.com/hollowdll/kvdb/errors"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/version"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	getServerInfoRPCName string = "GetServerInfo"
	getLogsRPCName       string = "GetLogs"
)

// ClientConnListener is a client connection listener
// that accepts new connections and tracks active connections.
type ClientConnListener struct {
	net.Listener
	logger               kvdb.Logger
	clientConnections    uint32
	maxClientConnections uint32
	mu                   sync.RWMutex
}

func NewClientConnListener(listener net.Listener, server *Server, maxConnections uint32) *ClientConnListener {
	return &ClientConnListener{
		Listener:             listener,
		logger:               server.logger,
		clientConnections:    0,
		maxClientConnections: maxConnections,
	}
}

func (l *ClientConnListener) Accept() (net.Conn, error) {
	conn, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}

	l.mu.Lock()
	l.clientConnections++
	l.logger.Debugf("Client connected, total clients: %d\n", l.clientConnections)

	clientConn := &clientConn{Conn: conn, release: func() {
		l.mu.Lock()
		if l.clientConnections > 0 {
			l.clientConnections--
		}
		l.logger.Debugf("Client disconnected, total clients: %d\n", l.clientConnections)
		l.mu.Unlock()
	}}

	if l.clientConnections > l.maxClientConnections {
		l.logger.Errorf("Incoming connection denied: %s", kvdberrors.ErrMaxClientConnectionsReached.Error())
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
	startTime       time.Time
	databases       map[string]*kvdb.Database
	credentialStore auth.CredentialStore
	logger          kvdb.Logger
	loggerMu        sync.RWMutex
	Cfg             config.ServerConfig
	*ClientConnListener
	mu sync.RWMutex
}

func NewKvdbServer(cfg config.ServerConfig) *KvdbServer {
	return &KvdbServer{
		startTime:          time.Now(),
		databases:          make(map[string]*kvdb.Database),
		credentialStore:    auth.NewInMemoryCredentialStore(),
		logger:             kvdb.NewDefaultLogger(),
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

// getTotalDataSize returns the total amount of stored data on this server in bytes.
func (s *KvdbServer) getTotalDataSize() uint64 {
	var sum uint64
	for _, db := range s.databases {
		sum += db.GetStoredSizeBytes()
	}

	return sum
}

// databaseExists returns true if a database with the given name exists on the server.
func (s *KvdbServer) dbExists(name string) bool {
	_, exists := s.databases[name]
	return exists
}

// getDatabaseNameFromContext gets the database name from the incoming context gRPC metadata.
func (s *KvdbServer) getDBNameFromContext(ctx context.Context) string {
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

type Server struct {
	serverpb.UnimplementedServerServiceServer
	dbpb.UnimplementedDatabaseServiceServer
	storagepb.UnimplementedGeneralKeyServiceServer
	storagepb.UnimplementedStringKeyServiceServer
	storagepb.UnimplementedHashMapKeyServiceServer
	startTime       time.Time
	databases       map[string]*kvdb.Database
	CredentialStore InMemoryCredentialStore
	// True if the server is password protected.
	passwordEnabled bool
	logger          kvdb.Logger
	logFilePath     string
	logFileEnabled  bool
	tlsEnabled      bool
	debugEnabled    bool
	// The maximum number of keys a database can hold.
	maxKeysPerDb uint32
	// The maximum number of fields a HashMap can hold.
	maxHashMapFields uint32
	// The name of the default database that is created at server startup.
	defaultDb string
	// The TCP/IP port the server is using.
	portInUse uint16
	*ClientConnListener
	mutex sync.RWMutex
}

// ServerOptions contains options that can be passed to the server when creating it.
type ServerOptions struct {
	MaxKeysPerDb         uint32
	MaxHashMapFields     uint32
	MaxClientConnections uint32
}

func NewServer() *Server {
	return &Server{
		startTime:          time.Now(),
		databases:          make(map[string]*kvdb.Database),
		CredentialStore:    *NewInMemoryCredentialStore(),
		passwordEnabled:    false,
		logger:             kvdb.NewDefaultLogger(),
		logFilePath:        "",
		logFileEnabled:     false,
		tlsEnabled:         false,
		debugEnabled:       false,
		maxKeysPerDb:       common.DbMaxKeyCount,
		maxHashMapFields:   common.HashMapMaxFields,
		defaultDb:          DefaultDatabase,
		portInUse:          common.ServerDefaultPort,
		ClientConnListener: nil,
	}
}

func NewServerWithOptions(options *ServerOptions) *Server {
	newServer := &Server{
		startTime:        time.Now(),
		databases:        make(map[string]*kvdb.Database),
		CredentialStore:  *NewInMemoryCredentialStore(),
		passwordEnabled:  false,
		logger:           kvdb.NewDefaultLogger(),
		logFilePath:      "",
		logFileEnabled:   false,
		tlsEnabled:       false,
		debugEnabled:     false,
		maxKeysPerDb:     options.MaxKeysPerDb,
		maxHashMapFields: options.MaxHashMapFields,
		defaultDb:        DefaultDatabase,
		portInUse:        common.ServerDefaultPort,
		ClientConnListener: &ClientConnListener{
			Listener:             nil,
			maxClientConnections: options.MaxClientConnections,
		},
	}
	if options.MaxKeysPerDb == 0 {
		newServer.maxKeysPerDb = common.DbMaxKeyCount
	}
	if options.MaxHashMapFields == 0 {
		newServer.maxHashMapFields = common.HashMapMaxFields
	}
	if options.MaxClientConnections == 0 {
		newServer.ClientConnListener.maxClientConnections = common.DefaultMaxClientConnections
	}

	return newServer
}

// DisableLogger disables all log outputs from this server.
func (s *Server) DisableLogger() {
	s.logger.Disable()
}

// EnableDebugMode enables debug mode.
func (s *Server) EnableDebugMode() {
	s.logger.EnableDebug()
	s.debugEnabled = true
}

// SetLogFilePath sets the file path to the log file.
func (s *Server) SetLogFilePath(filePath string) {
	s.logFilePath = filePath
}

// EnableLogFile enables logger to write logs to the log file.
func (s *Server) EnableLogFile() {
	err := s.logger.EnableLogFile(s.logFilePath)
	if err != nil {
		s.logger.Fatalf("Failed to enable log file: %v", err)
	}
	s.logFileEnabled = true
}

// EnableTls enables TLS.
func (s *Server) EnableTls() {
	s.tlsEnabled = true
}

// CloseLogger closes logger and releases its possible resources.
func (s *Server) CloseLogger() {
	err := s.logger.CloseLogFile()
	if err != nil {
		s.logger.Fatalf("Failed to close log file: %v", err)
	}
}

// EnablePasswordProtection enables server password protection and sets the password.
func (s *Server) EnablePasswordProtection(password string) {
	if err := s.CredentialStore.SetServerPassword([]byte(password)); err != nil {
		s.logger.Fatalf("Failed to set server password: %v", err)
	}
	s.passwordEnabled = true
	s.logger.Infof("Password protection is enabled. Clients need to authenticate using password.")
}

// getTotalDataSize returns the total amount of stored data on this server in bytes.
func (s *Server) getTotalDataSize() uint64 {
	var sum uint64
	for _, db := range s.databases {
		sum += db.GetStoredSizeBytes()
	}

	return sum
}

// CreateDefaultDatabase creates an empty default database.
func (s *Server) CreateDefaultDatabase(name string) {
	if err := kvdb.ValidateDatabaseName(name); err != nil {
		s.logger.Fatalf("Failed to create default database: %v", err)
	}
	db := kvdb.CreateDatabase(name)
	s.databases[db.Name] = db
	s.defaultDb = db.Name
	s.logger.Infof("Created default database '%s'", db.Name)
}

// SetPort sets the port that the server should use.
func (s *Server) SetPort(port uint16) {
	s.portInUse = port
	s.logger.Infof("Configured port to use: %d", port)
}

// DbMaxKeysReached returns true if a database has reached or exceeded the maximum key limit.
func (s *Server) DbMaxKeysReached(db *kvdb.Database) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return db.GetKeyCount() >= s.maxKeysPerDb
}

// getDatabaseNameFromContext gets the database name from the incoming gRPC metadata.
func (s *Server) getDatabaseNameFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return s.defaultDb
	}

	dbName := md.Get(common.GrpcMetadataKeyDbName)
	if len(dbName) < 1 {
		return s.defaultDb
	}

	return dbName[0]
}

// GetServerInfo is the implementation of RPC GetServerInfo.
func (s *Server) GetServerInfo(ctx context.Context, req *serverpb.GetServerInfoRequest) (res *serverpb.GetServerInfoResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.ClientConnListener != nil {
		s.ClientConnListener.mu.RLock()
		defer s.ClientConnListener.mu.RUnlock()
	}

	s.logger.Debugf("%s: (call) %v", getServerInfoRPCName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", getServerInfoRPCName, err)
		} else {
			s.logger.Debugf("%s: (success) %v", getServerInfoRPCName, req)
		}
	}()

	osInfo, err := getOsInfo()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	var totalKeys uint64
	for _, db := range s.databases {
		totalKeys += uint64(db.GetKeyCount())
	}

	generalInfo := &serverpb.GeneralInfo{
		KvdbVersion:     version.Version,
		GoVersion:       runtime.Version(),
		Os:              osInfo,
		Arch:            runtime.GOARCH,
		ProcessId:       uint32(os.Getpid()),
		UptimeSeconds:   uint64(time.Since(s.startTime).Seconds()),
		TcpPort:         uint32(s.portInUse),
		TlsEnabled:      s.tlsEnabled,
		PasswordEnabled: s.passwordEnabled,
		LogfileEnabled:  s.logFileEnabled,
		DebugEnabled:    s.debugEnabled,
	}
	memoryInfo := &serverpb.MemoryInfo{
		MemoryAlloc:      m.Alloc,
		MemoryTotalAlloc: m.TotalAlloc,
		MemorySys:        m.Sys,
	}
	storageInfo := &serverpb.StorageInfo{
		TotalDataSize: s.getTotalDataSize(),
		TotalKeys:     totalKeys,
	}
	clientInfo := &serverpb.ClientInfo{
		ClientConnections:    s.ClientConnListener.clientConnections,
		MaxClientConnections: s.ClientConnListener.maxClientConnections,
	}
	dbInfo := &serverpb.DatabaseInfo{
		DbCount:   uint32(len(s.databases)),
		DefaultDb: s.defaultDb,
	}

	return &serverpb.GetServerInfoResponse{
		GeneralInfo: generalInfo,
		MemoryInfo:  memoryInfo,
		StorageInfo: storageInfo,
		ClientInfo:  clientInfo,
		DbInfo:      dbInfo,
	}, nil
}

// GetLogs is the implementation of RPC GetLogs.
func (s *Server) GetLogs(ctx context.Context, req *serverpb.GetLogsRequest) (res *serverpb.GetLogsResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	s.logger.Debugf("%s: (call) %v", getLogsRPCName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", getLogsRPCName, err)
		} else {
			s.logger.Debugf("%s: (success) %v", getLogsRPCName, req)
		}
	}()

	if !s.logFileEnabled {
		return nil, status.Errorf(codes.FailedPrecondition, "%s: enable server log file to get logs", kvdberrors.ErrLogFileNotEnabled.Error())
	}
	s.logger.Debugf("%s: log file is enabled", getLogsRPCName)

	lines, err := common.ReadFileLines(s.logFilePath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &serverpb.GetLogsResponse{Logs: lines}, nil
}

// initServer initializes the server.
// Returns the initialized Server and grpc.Server.
func initServer() (*Server, *grpc.Server) {
	server := NewServer()
	initConfig(server)
	server.logger.ClearFlags()
	server.logger.Infof("Starting kvdb v%s server ...", version.Version)

	if viper.GetBool(ConfigKeyLogFileEnabled) {
		server.EnableLogFile()
		server.logger.Infof("Log file is enabled. Logs will be written to the log file. The file is located at %s", server.logFilePath)
	}

	if viper.GetBool(ConfigKeyDebugEnabled) {
		server.EnableDebugMode()
		server.logger.Info("Debug mode is enabled. Debug messages will be logged")
	}

	if viper.GetBool(ConfigKeyTlsEnabled) {
		server.EnableTls()
	}

	password, present := os.LookupEnv(EnvVarPassword)
	if present {
		server.EnablePasswordProtection(password)
	} else {
		server.logger.Warning("Password protection is disabled")
	}

	server.CreateDefaultDatabase(viper.GetString(ConfigKeyDefaultDatabase))
	server.SetPort(viper.GetUint16(ConfigKeyPort))

	var grpcServer *grpc.Server = nil
	if !server.tlsEnabled {
		grpcServer = grpc.NewServer(grpc.UnaryInterceptor(server.AuthInterceptor))
		server.logger.Warning("TLS is disabled. Connections will not be encrypted")
	} else {
		server.logger.Info("Attempting to enable TLS ...")
		certBytes, err := os.ReadFile(viper.GetString(ConfigKeyTlsCertPath))
		if err != nil {
			server.logger.Fatalf("Failed to read TLS certificate: %v", err)
		}
		keyBytes, err := os.ReadFile(viper.GetString(ConfigKeyTlsPrivKeyPath))
		if err != nil {
			server.logger.Fatalf("Failed to read TLS private key: %v", err)
		}

		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(certBytes) {
			server.logger.Fatal("Failed to parse TLS certificate")
		}
		cert, err := tls.X509KeyPair(certBytes, keyBytes)
		if err != nil {
			server.logger.Fatalf("Failed to parse TLS public/private key pair: %v", err)
		}

		creds := credentials.NewServerTLSFromCert(&cert)
		grpcServer = grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(server.AuthInterceptor))
		server.logger.Info("TLS is enabled. Connections will be encrypted")
	}

	serverpb.RegisterServerServiceServer(grpcServer, server)
	dbpb.RegisterDatabaseServiceServer(grpcServer, server)
	storagepb.RegisterGeneralKeyServiceServer(grpcServer, server)
	storagepb.RegisterStringKeyServiceServer(grpcServer, server)
	storagepb.RegisterHashMapKeyServiceServer(grpcServer, server)

	return server, grpcServer
}

// StartServer initializes and starts the server.
func StartServer() {
	server, grpcServer := initServer()
	defer server.CloseLogger()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", server.portInUse))
	if err != nil {
		server.logger.Fatalf("Failed to listen: %v", err)
	}
	server.logger.Infof("Server listening at %v", listener.Addr())

	connListener := NewClientConnListener(listener, server, viper.GetUint32(ConfigKeyMaxClientConnections))
	server.ClientConnListener = connListener

	if err := grpcServer.Serve(connListener); err != nil {
		server.logger.Errorf("Failed to accept incoming connection: %v", err)
	}
}
