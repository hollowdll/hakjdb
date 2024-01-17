package server

import (
	"fmt"
	"net"

	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// initServer initializes the server.
func InitServer() {
	server := NewServer()
	initConfig(server)
	server.logger.ClearFlags()

	// Enable debug logs.
	if viper.GetBool(ConfigKeyDebugMode) {
		server.logger.EnableDebug()
		server.logger.Info("Debug mode is enabled. Debug messages will be logged.")
	}

	grpcServer := grpc.NewServer()
	kvdbserver.RegisterDatabaseServiceServer(grpcServer, server)
	kvdbserver.RegisterServerServiceServer(grpcServer, server)
	kvdbserver.RegisterStorageServiceServer(grpcServer, server)

	startServer(server, grpcServer)
}

// startServer starts the server.
func startServer(server *Server, grpcServer *grpc.Server) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetUint16(ConfigKeyPort)))
	if err != nil {
		server.logger.Fatalf("Failed to listen: %v", err)
	}
	server.logger.Infof("Server listening at %v", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		server.logger.Fatalf("Failed to serve gRPC: %v", err)
	}
}
