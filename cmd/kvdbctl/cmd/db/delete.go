package db

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/api/v0/dbpb"
	"github.com/hollowdll/kvdb/cmd/kvdbctl/client"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdDbDelete = &cobra.Command{
	Use:   "delete NAME",
	Short: "Delete a database",
	Long:  "Delete a database with the specified name.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		deleteDatabase(args[0])
	},
}

func deleteDatabase(name string) {
	if !client.PromptConfirm(fmt.Sprintf("Delete database '%s' and all its data? Yes/No: ", name)) {
		return
	}
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	res, err := client.GrpcDBClient.DeleteDB(ctx, &dbpb.DeleteDBRequest{DbName: name})
	client.CheckGrpcError(err)
	fmt.Println(res.DbName)
}
