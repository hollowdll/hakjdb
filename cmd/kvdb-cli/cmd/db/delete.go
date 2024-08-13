package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/kvdb/api/v0/dbpb"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdDbDelete = &cobra.Command{
	Use:   "delete NAME",
	Short: "Delete a database",
	Long:  "Deletes a database.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		deleteDatabase(args[0])
	},
}

func deleteDatabase(name string) {
	if !promptConfirmDelete(name) {
		return
	}
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	res, err := client.GrpcDBClient.DeleteDB(ctx, &dbpb.DeleteDBRequest{DbName: name})
	client.CheckGrpcError(err)
	fmt.Println(res.DbName)
}

func promptConfirmDelete(dbName string) bool {
	var input string
	fmt.Printf("Delete database '%s' and all its data? Yes/No: ", dbName)
	_, err := fmt.Scanln(&input)
	input = strings.TrimSpace(input)
	if input == "" {
		return false
	}
	cobra.CheckErr(err)

	return strings.ToLower(input) == "yes"
}
