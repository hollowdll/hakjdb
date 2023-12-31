package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/cobra"
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
	ctx, cancel := context.WithTimeout(context.Background(), client.ClientCtxTimeout)
	defer cancel()
	response, err := client.GrpcDatabaseClient.GetAllDatabases(ctx, &kvdbserver.Empty{})
	cobra.CheckErr(err)

	if len(response.Names) > 0 {
		fmt.Println(strings.Join(response.Names, "\n"))
	}
}
