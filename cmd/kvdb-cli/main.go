package main

import (
	"fmt"
	"os"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/cmd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address = "localhost:12345"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: failed to connect to the server:", err)
		os.Exit(1)
	}
	defer conn.Close()
	client.InitClient(conn)

	cmd.Execute()
}
