package cmd

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserverpb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdDeleteAllKeys = &cobra.Command{
	Use:   "deletekeys",
	Short: "Delete all the keys of a database",
	Long:  "Deletes all the keys of a database.",
	Run: func(cmd *cobra.Command, args []string) {
		deleteAllKeys()
	},
}

func init() {
	cmdDeleteAllKeys.Flags().StringVarP(&dbName, "database", "d", "", "database to use")
}

func deleteAllKeys() {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	_, err := client.GrpcStorageClient.DeleteAllKeys(ctx, &kvdbserverpb.DeleteAllKeysRequest{})
	client.CheckGrpcError(err)

	fmt.Println("OK")
}
