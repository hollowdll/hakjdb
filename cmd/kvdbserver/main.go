package main

import (
	"fmt"
	"net"

	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"google.golang.org/grpc"
)

func main() {
	server := newServer()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", common.ServerDefaultPort)) // env var later
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
