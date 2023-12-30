package main

import (
	"os"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/cmd"
)

func main() {
	defer client.CloseConnections()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
