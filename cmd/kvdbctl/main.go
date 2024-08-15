package main

import (
	"os"

	"github.com/hollowdll/kvdb/cmd/kvdbctl/client"
	"github.com/hollowdll/kvdb/cmd/kvdbctl/cmd"
)

func main() {
	defer client.CloseConnections()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
