package main

import (
	"log"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/cmd"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address = "localhost:12345"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	client.GrpcClient = kvdbserver.NewDatabaseClient(conn)

	cmd.Execute()
}
