package main

import (
	"fmt"
	"os"

	kvdb "github.com/hollowdll/kvdb"
)

func main() {
	fmt.Println("kvdb server")

	logger := kvdb.NewLogger()
	err := logger.LogMessage(kvdb.LogTypeInfo, "Test log")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}
