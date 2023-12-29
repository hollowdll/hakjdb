package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/cobra"
)

// CommandTimeout specifies how long to wait until a command terminates.
var CommandTimeout = time.Second * 10

var dbName string
var cmdCreateDb = &cobra.Command{
	Use:   "create [flags]",
	Short: "Create a new database",
	Long:  "Create a new database",
	Run: func(cmd *cobra.Command, args []string) {
		createDatabase()
	},
}

func init() {
	cmdCreateDb.Flags().StringVarP(&dbName, "name", "n", "", "name of the database (required)")
	cmdCreateDb.MarkFlagRequired("name")
}

func createDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), CommandTimeout)
	defer cancel()
	response, err := client.GrpcClient.CreateDatabase(ctx, &kvdbserver.CreateDatabaseRequest{Name: dbName})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: cannot create database:", err)
	}

	fmt.Println("Created database:", response.GetName())
}
