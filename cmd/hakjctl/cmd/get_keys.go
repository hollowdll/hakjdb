package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/hakjdb/api/v1/kvpb"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/client"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/config"
	"github.com/hollowdll/hakjdb/internal/common"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdGetKeys = &cobra.Command{
	Use:   "getkeys",
	Short: "List keys",
	Long:  "List all the keys of a database.",
	Example: `# Use the default database
hakjctl getkeys

# Specify the database to use
hakjctl getkeys -d default`,
	Run: func(cmd *cobra.Command, args []string) {
		getKeys()
	},
}

func init() {
	cmdGetKeys.Flags().StringVarP(&dbName, "database", "d", "", client.DBFlagMsg)
}

func getKeys() {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, config.GetCmdTimeout())
	defer cancel()

	res, err := client.GrpcGeneralKVClient.GetAllKeys(ctx, &kvpb.GetAllKeysRequest{})
	client.CheckGrpcError(err)

	for i, key := range res.Keys {
		res.Keys[i] = fmt.Sprintf("%d) %s", i+1, key)
	}
	if len(res.Keys) > 0 {
		fmt.Println(strings.Join(res.Keys, "\n"))
	}
}
