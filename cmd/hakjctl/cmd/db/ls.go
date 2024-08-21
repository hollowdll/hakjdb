package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/hakjdb/api/v1/dbpb"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/client"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/config"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdDbLs = &cobra.Command{
	Use:   "ls",
	Short: "List all the databases",
	Long:  "List the names of all the databases that exist on the server.",
	Example: `# List the names of all the databases
hakjctl db ls`,
	Run: func(cmd *cobra.Command, args []string) {
		showDatabaseNames()
	},
}

func showDatabaseNames() {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, config.GetCmdTimeout())
	defer cancel()
	response, err := client.GrpcDBClient.GetAllDBs(ctx, &dbpb.GetAllDBsRequest{})
	client.CheckGrpcError(err)

	if len(response.DbNames) > 0 {
		fmt.Println(strings.Join(response.DbNames, "\n"))
	}
}
