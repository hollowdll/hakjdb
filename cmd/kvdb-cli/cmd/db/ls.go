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

var cmdDbLs = &cobra.Command{
	Use:   "ls",
	Short: "List all databases",
	Long:  "Lists all the databases that exist on the server.",
	Run: func(cmd *cobra.Command, args []string) {
		showDatabaseNames()
	},
}

func showDatabaseNames() {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()
	response, err := client.GrpcDatabaseClient.GetAllDatabases(ctx, &dbpb.GetAllDatabasesRequest{})
	client.CheckGrpcError(err)

	if len(response.DbNames) > 0 {
		fmt.Println(strings.Join(response.DbNames, "\n"))
	}
}
