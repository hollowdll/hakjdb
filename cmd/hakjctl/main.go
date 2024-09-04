package main

import (
	"os"

	"github.com/hollowdll/hakjdb/cmd/hakjctl/client"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/cmd"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/config"
)

func main() {
	defer client.CloseConnections()
	config.InitConfig()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
