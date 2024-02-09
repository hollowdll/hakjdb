package db

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/config"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc/metadata"
)

var cmdDbDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete a database",
	Long:  "Delete a database",
	Run: func(cmd *cobra.Command, args []string) {
		deleteDatabase()
	},
}

func init() {
	cmdDbDelete.Flags().StringVarP(&dbName, "name", "n", "", "name of the database")
}

func deleteDatabase() {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()
	if len(dbName) < 1 {
		dbName = viper.GetString(config.ConfigKeyDatabase)
	}
	res, err := client.GrpcDatabaseClient.DeleteDatabase(ctx, &kvdbserver.DeleteDatabaseRequest{DbName: dbName})
	client.CheckGrpcError(err)

	if res.Ok {
		fmt.Println("OK")
	} else {
		fmt.Println(client.ValueNone)
	}
}
