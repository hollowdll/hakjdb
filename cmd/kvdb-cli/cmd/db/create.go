package db

import (
	"context"
	"log"
	"time"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/cobra"
)

var dbName string
var cmdCreateDb = &cobra.Command{
	Use:   "create [flags]",
	Short: "Create a new database",
	Long:  "Create a new database",
	Run: func(cmd *cobra.Command, args []string) {
		createDatabase()
	},
}
var cmdTimeout = time.Second

func init() {
	cmdCreateDb.Flags().StringVarP(&dbName, "name", "n", "", "name of the database (required)")
	cmdCreateDb.MarkFlagRequired("name")
}

func createDatabase() {
	response, err := client.GrpcClient.CreateDatabase(context.Background(), &kvdbserver.CreateDatabaseRequest{Name: dbName})
	if err != nil {
		log.Fatalf("cannot create database: %v", err)
	}

	log.Printf("Created database: %s", response.GetName())
}
