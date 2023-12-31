package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"google.golang.org/grpc"
)

func main() {
	server := newServer()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", common.ServerDefaultPort)) // env var later
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	kvdbserver.RegisterDatabaseServer(grpcServer, server)
	log.Printf("Server listening at %v", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
