package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"

	kvdb "github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/hollowdll/kvdb/version"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	kvdbserver.UnimplementedDatabaseServiceServer
	kvdbserver.UnimplementedServerServiceServer
	kvdbserver.UnimplementedStorageServiceServer
	startTime       time.Time
	databases       map[string]*kvdb.Database
	credentialStore InMemoryCredentialStore
	logger          kvdb.Logger
	mutex           sync.RWMutex
}

// portInUse is the TCP/IP port the server uses.
var portInUse uint16

func NewServer() *Server {
	return &Server{
		startTime:       time.Now(),
		databases:       make(map[string]*kvdb.Database),
		credentialStore: *newInMemoryCredentialStore(),
		logger:          kvdb.NewDefaultLogger(),
	}
}

// DisableLogger disables all log outputs from this server.
func (s *Server) DisableLogger() {
	s.logger.Disable()
}

// getTotalDataSize returns the total amount of stored data on this server in bytes.
func (s *Server) getTotalDataSize() uint64 {
	var sum uint64
	for _, db := range s.databases {
		sum += db.GetStoredSizeBytes()
	}

	return sum
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

// GetServerInfo returns information about the server.
func (s *Server) GetServerInfo(ctx context.Context, req *kvdbserver.GetServerInfoRequest) (res *kvdbserver.GetServerInfoResponse, err error) {
	s.logger.Debug("Attempt to get server info")
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to get server info: %s", err)
		} else {
			s.logger.Debug("Get server info success")
		}
	}()

	osInfo, err := getOsInfo()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	info := &kvdbserver.ServerInfo{
		KvdbVersion:   version.Version,
		GoVersion:     runtime.Version(),
		DbCount:       uint32(len(s.databases)),
		TotalDataSize: s.getTotalDataSize(),
		Os:            osInfo,
		Arch:          runtime.GOARCH,
		ProcessId:     uint32(os.Getpid()),
		UptimeSeconds: uint64(time.Since(s.startTime).Seconds()),
		TcpPort:       uint32(portInUse),
	}

	return &kvdbserver.GetServerInfoResponse{Data: info}, nil
}

// initServer initializes the server.
// Returns the initialized Server and grpc.Server.
func initServer() (*Server, *grpc.Server) {
	server := NewServer()
	initConfig(server)
	server.logger.ClearFlags()

	// Enable debug logs.
	if viper.GetBool(ConfigKeyDebugEnabled) {
		server.logger.EnableDebug()
		server.logger.Info("Debug mode is enabled. Debug messages will be logged.")
	}

	grpcServer := grpc.NewServer()
	kvdbserver.RegisterDatabaseServiceServer(grpcServer, server)
	kvdbserver.RegisterServerServiceServer(grpcServer, server)
	kvdbserver.RegisterStorageServiceServer(grpcServer, server)

	return server, grpcServer
}

// StartServer initializes and starts the server.
func StartServer() {
	server, grpcServer := initServer()
	portInUse = viper.GetUint16(ConfigKeyPort)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", portInUse))
	if err != nil {
		server.logger.Fatalf("Failed to listen: %v", err)
	}
	server.logger.Infof("Server listening at %v", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		server.logger.Fatalf("Failed to serve gRPC: %v", err)
	}
}
