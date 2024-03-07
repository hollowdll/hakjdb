package hashmap

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdGetHashMapFieldValue = &cobra.Command{
	Use:   "get [key] [field]",
	Short: "Get a HashMap field value",
	Long:  "Get a HashMap field value using a key",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		getHashMapFieldValue(args[0], args[1])
	},
}

func init() {
	cmdGetHashMapFieldValue.Flags().StringVarP(&dbName, "db", "d", "", "database to use")
}

func getHashMapFieldValue(key string, field string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	res, err := client.GrpcStorageClient.GetHashMapFieldValue(ctx, &kvdbserver.GetHashMapFieldValueRequest{Key: key, Field: field})
	client.CheckGrpcError(err)

	if res.Ok {
		fmt.Printf("\"%s\"\n", res.Value)
	} else {
		fmt.Println(client.ValueNone)
	}
}
