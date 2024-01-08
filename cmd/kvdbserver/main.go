package main

import (
	"fmt"
	"net"

	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	initServer()
}

func initServer() {
	server := newServer()
	initConfig(server)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetUint16(configKeyPort)))
	if err != nil {
		server.logger.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	kvdbserver.RegisterDatabaseServiceServer(grpcServer, server)
	kvdbserver.RegisterServerServiceServer(grpcServer, server)
	kvdbserver.RegisterStorageServiceServer(grpcServer, server)

	server.logger.Infof("Server listening at %v", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		server.logger.Fatalf("Failed to serve gRPC: %v", err)
	}
}
