package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdDbLs = &cobra.Command{
	Use:   "ls",
	Short: "List all databases",
	Long:  "List all databases",
	Run: func(cmd *cobra.Command, args []string) {
		showDatabaseNames()
	},
}

func showDatabaseNames() {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeoutSeconds)
	defer cancel()
	response, err := client.GrpcDatabaseClient.GetAllDatabases(ctx, &kvdbserver.GetAllDatabasesRequest{})
	client.CheckGrpcError(err)

	if len(response.DbNames) > 0 {
		fmt.Println(strings.Join(response.DbNames, "\n"))
	}
}
