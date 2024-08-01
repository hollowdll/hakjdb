package cmd

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/api/v0/kvpb"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdDeleteKeys = &cobra.Command{
	Use:   "delete [key ...]",
	Short: "Delete keys",
	Long: `
Deletes the specified keys and the values they are holding.
Ignores keys that do not exist.
This command can delete multiple keys.
`,
	Args: cobra.MatchAll(cobra.MinimumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		deleteKeys(args[0:])
	},
}

func init() {
	cmdDeleteKeys.Flags().StringVarP(&dbName, "database", "d", "", "database to use")
}

func deleteKeys(keys []string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	res, err := client.GrpcGeneralKVClient.DeleteKeys(ctx, &kvpb.DeleteKeysRequest{Keys: keys})
	client.CheckGrpcError(err)

	fmt.Printf("%d\n", res.KeysDeletedCount)
}
