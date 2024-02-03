package db

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdDbCreate = &cobra.Command{
	Use:   "create",
	Short: "Create a new database",
	Long:  "Create a new database",
	Run: func(cmd *cobra.Command, args []string) {
		createDatabase()
	},
}

func init() {
	cmdDbCreate.Flags().StringVarP(&dbName, "name", "n", "", "name of the database (required)")
	cmdDbCreate.MarkFlagRequired("name")
}

func createDatabase() {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeoutSeconds)
	defer cancel()
	_, err := client.GrpcDatabaseClient.CreateDatabase(ctx, &kvdbserver.CreateDatabaseRequest{DbName: dbName})
	client.CheckGrpcError(err)

	fmt.Println("OK")
}
