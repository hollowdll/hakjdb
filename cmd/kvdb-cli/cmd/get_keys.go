package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/kvdb/api/v0/kvpb"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdGetKeys = &cobra.Command{
	Use:   "getkeys",
	Short: "Get all the keys of a database",
	Long:  "Gets all the keys of a database.",
	Run: func(cmd *cobra.Command, args []string) {
		getKeys()
	},
}

func init() {
	cmdGetKeys.Flags().StringVarP(&dbName, "database", "d", "", "database to use")
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
