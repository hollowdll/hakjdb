package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/kvdb/api/v0/kvpb"
	"github.com/hollowdll/kvdb/cmd/kvdbctl/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdGetKeys = &cobra.Command{
	Use:   "getkeys",
	Short: "List keys",
	Long:  "List all the keys of a database.",
	Example: `# Use the default database
kvdbctl getkeys

# Specify the database to use
kvdbctl getkeys --database default`,
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
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
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
