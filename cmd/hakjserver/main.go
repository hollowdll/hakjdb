package main

import (
	"os"

	"github.com/hollowdll/hakjdb/cmd/hakjserver/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
