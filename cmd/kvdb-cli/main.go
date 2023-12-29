package main

import (
	"fmt"
	"os"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/cmd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address = fmt.Sprintf("%s:%d", cmd.Hostname, cmd.Port)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: failed to connect to the server:", err)
		os.Exit(1)
	}
	defer conn.Close()
	client.InitClient(conn)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
