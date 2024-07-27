package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/kvdb/api/v0/dbpb"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc/metadata"
)

var cmdDbDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete a database",
	Long:  "Deletes a database.",
	Run: func(cmd *cobra.Command, args []string) {
		deleteDatabase()
	},
}

func init() {
	cmdDbDelete.Flags().StringVarP(&dbName, "name", "n", "", "name of the database")
}

func deleteDatabase() {
	if !promptConfirmDelete() {
		return
	}
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()
	if len(dbName) < 1 {
		dbName = viper.GetString(config.ConfigKeyDatabase)
	}

	res, err := client.GrpcDatabaseClient.DeleteDatabase(ctx, &dbpb.DeleteDatabaseRequest{DbName: dbName})
	client.CheckGrpcError(err)

	fmt.Println(res.DbName)
}

func promptConfirmDelete() bool {
	var input string
	fmt.Print("Delete database and all its data? Yes/No: ")
	_, err := fmt.Scanln(&input)
	input = strings.TrimSpace(input)
	if input == "" {
		return false
	}
	cobra.CheckErr(err)

	return strings.ToLower(input) == "yes"
}
