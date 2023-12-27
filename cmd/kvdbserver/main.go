package main

import (
	"fmt"
	"os"

	kvdb "github.com/hollowdll/kvdb"
)

func main() {
	fmt.Println("kvdb server")

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
}
