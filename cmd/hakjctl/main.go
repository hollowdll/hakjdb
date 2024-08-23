package main

import (
	"os"

	"github.com/hollowdll/hakjdb/cmd/hakjctl/client"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/cmd"
)

func main() {
	defer client.CloseConnections()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
