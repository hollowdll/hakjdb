package main

import (
	"log"
	"net"

	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"google.golang.org/grpc"
)

func main() {
	server := newDatabaseServer()
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	kvdbserver.RegisterDatabaseServer(grpcServer, server)
	log.Printf("Server listening at %v", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to server gRPC: %v", err)
	}

	/* test stuff
	db, err := kvdb.CreateDatabase("valid_db1")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	fmt.Println("DB:", db.Name)

	logger := kvdb.NewLogger()
	if err := logger.LogMessage(kvdb.LogTypeInfo, "Test log"); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
	*/
}
