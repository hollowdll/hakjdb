package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"

	kvdb "github.com/hollowdll/kvdb"
	kvdberrors "github.com/hollowdll/kvdb/errors"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserverpb"
	"github.com/hollowdll/kvdb/version"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	kvdbserverpb.UnimplementedDatabaseServiceServer
	kvdbserverpb.UnimplementedServerServiceServer
	kvdbserverpb.UnimplementedStorageServiceServer
	startTime       time.Time
	databases       map[string]*kvdb.Database
	CredentialStore InMemoryCredentialStore
	// True if the server is password protected.
	passwordEnabled bool
	logger          kvdb.Logger
	logFilePath     string
	logFileEnabled  bool
	tlsEnabled      bool
	// The maximum number of keys a database can hold.
	maxKeysPerDb uint32
	// The maximum number of fields a HashMap can hold.
	maxHashMapFields uint32
	// The name of the default database that is created at server startup.
	defaultDb string
	// The TCP/IP port the server is using.
	portInUse uint16
	mutex     sync.RWMutex
}

// ServerOptions contains options that can be passed to the server when creating it.
type ServerOptions struct {
	MaxKeysPerDb     uint32
	MaxHashMapFields uint32
}

func NewServer() *Server {
	return &Server{
		startTime:        time.Now(),
		databases:        make(map[string]*kvdb.Database),
		CredentialStore:  *NewInMemoryCredentialStore(),
		passwordEnabled:  false,
		logger:           kvdb.NewDefaultLogger(),
		logFilePath:      "",
		logFileEnabled:   false,
		tlsEnabled:       false,
		maxKeysPerDb:     common.DbMaxKeyCount,
		maxHashMapFields: common.HashMapMaxFields,
		defaultDb:        DefaultDatabase,
		portInUse:        common.ServerDefaultPort,
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
		maxKeysPerDb:     options.MaxKeysPerDb,
		maxHashMapFields: options.MaxHashMapFields,
		defaultDb:        DefaultDatabase,
		portInUse:        common.ServerDefaultPort,
	}
	if newServer.maxKeysPerDb == 0 {
		newServer.maxKeysPerDb = common.DbMaxKeyCount
	}
	if newServer.maxHashMapFields == 0 {
		newServer.maxHashMapFields = common.HashMapMaxFields
	}

	return newServer
}

// DisableLogger disables all log outputs from this server.
func (s *Server) DisableLogger() {
	s.logger.Disable()
}

// EnableDebugLogs enables server debug logs.
func (s *Server) EnableDebugLogs() {
	s.logger.EnableDebug()
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

// getOsInfo returns information about the server's operating system.
func getOsInfo() (string, error) {
	osInfo := runtime.GOOS

	switch osInfo {
	case "linux":
		cmd := exec.Command("uname", "-r", "-m")
		output, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return "Linux " + strings.TrimSpace(string(output)), nil
	case "windows":
		cmd := exec.Command("cmd", "/c", "ver")
		output, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(output)), nil
	default:
		return osInfo, nil
	}
}

// GetServerInfo is the implementation of RPC GetServerInfo.
func (s *Server) GetServerInfo(ctx context.Context, req *kvdbserverpb.GetServerInfoRequest) (res *kvdbserverpb.GetServerInfoResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	logPrefix := "GetServerInfo"
	s.logger.Debugf("%s: (attempt) %v", logPrefix, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: (success) %v", logPrefix, req)
		}
	}()

	osInfo, err := getOsInfo()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	info := &kvdbserverpb.ServerInfo{
		KvdbVersion:   version.Version,
		GoVersion:     runtime.Version(),
		DbCount:       uint32(len(s.databases)),
		TotalDataSize: s.getTotalDataSize(),
		Os:            osInfo,
		Arch:          runtime.GOARCH,
		ProcessId:     uint32(os.Getpid()),
		UptimeSeconds: uint64(time.Since(s.startTime).Seconds()),
		TcpPort:       uint32(s.portInUse),
	}

	return &kvdbserverpb.GetServerInfoResponse{Data: info}, nil
}

// GetLogs is the implementation of RPC GetLogs.
func (s *Server) GetLogs(ctx context.Context, req *kvdbserverpb.GetLogsRequest) (res *kvdbserverpb.GetLogsResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	logPrefix := "GetLogs"
	s.logger.Debugf("%s: (attempt) %v", logPrefix, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: (success) %v", logPrefix, req)
		}
	}()

	if !s.logFileEnabled {
		return nil, status.Errorf(codes.FailedPrecondition, "%s: enable server log file to get logs", kvdberrors.ErrLogFileNotEnabled.Error())
	}
	s.logger.Debugf("%s: log file is enabled", logPrefix)

	lines, err := common.ReadFileLines(s.logFilePath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &kvdbserverpb.GetLogsResponse{Logs: lines, LogfileEnabled: true}, nil
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
		server.EnableDebugLogs()
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
		grpcServer = grpc.NewServer(grpc.UnaryInterceptor(server.authInterceptor))
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
		grpcServer = grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(server.authInterceptor))
		server.logger.Info("TLS is enabled. Connections will be encrypted")
	}

	kvdbserverpb.RegisterDatabaseServiceServer(grpcServer, server)
	kvdbserverpb.RegisterServerServiceServer(grpcServer, server)
	kvdbserverpb.RegisterStorageServiceServer(grpcServer, server)

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

	if err := grpcServer.Serve(listener); err != nil {
		server.logger.Fatalf("Failed to serve gRPC: %v", err)
	}
}
