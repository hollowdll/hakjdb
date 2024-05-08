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

var cmdDeleteKey = &cobra.Command{
	Use:   "delete [key ...]",
	Short: "Delete specified keys",
	Long:  "Deletes the specified keys and the values they are holding.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		deleteKey(args[0:])
	},
}

func init() {
	cmdDeleteKey.Flags().StringVarP(&dbName, "database", "d", "", "database to use")
}

func deleteKey(keys []string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	res, err := client.GrpcStorageClient.DeleteKey(ctx, &kvdbserverpb.DeleteKeyRequest{Keys: keys})
	client.CheckGrpcError(err)

	fmt.Printf("%d\n", res.KeysDeleted)
}
