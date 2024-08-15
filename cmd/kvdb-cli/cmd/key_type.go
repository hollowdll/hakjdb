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

var cmdGetKeyType = &cobra.Command{
	Use:   "keytype KEY",
	Short: "Get the data type of a key",
	Long:  "Get the data type of a key.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		getKeyType(args[0])
	},
}

func init() {
	cmdGetKeyType.Flags().StringVarP(&dbName, "database", "d", "", client.DBFlagMsg)
}

func getKeyType(key string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	response, err := client.GrpcGeneralKVClient.GetKeyType(ctx, &kvpb.GetKeyTypeRequest{Key: key})
	client.CheckGrpcError(err)

	if response.Ok {
		fmt.Printf("\"%s\"\n", response.KeyType)
	} else {
		fmt.Println(client.ValueNone)
	}
}
