package cmd

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/api/v0/storagepb"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdGetKeyType = &cobra.Command{
	Use:   "keytype [key]",
	Short: "Get the data type of a key",
	Long:  "Gets the data type of the value a key is holding.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		getKeyType(args[0])
	},
}

func init() {
	cmdGetKeyType.Flags().StringVarP(&dbName, "database", "d", "", "database to use")
}

func getKeyType(key string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	response, err := client.GrpcGeneralKeyClient.GetKeyType(ctx, &storagepb.GetKeyTypeRequest{Key: key})
	client.CheckGrpcError(err)

	if response.Ok {
		fmt.Printf("\"%s\"\n", response.KeyType)
	} else {
		fmt.Println(client.ValueNone)
	}
}
