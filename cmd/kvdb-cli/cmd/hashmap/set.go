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

var cmdSetHashMap = &cobra.Command{
	Use:   "set [key] [field value ...]",
	Short: "Set HashMap fields and values",
	Long:  "Set HashMap fields and their corresponding values using a key",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)%2 == 0 {
			cobra.CheckErr("invalid number of arguments")
		}

		fields := make(map[string]string)
		for i := 1; i < len(args); i += 2 {
			fields[args[i]] = args[i+1]
		}

		setHashMap(args[0], fields)
	},
}

func init() {
	cmdSetHashMap.Flags().StringVarP(&dbName, "database", "d", "", "database to use")
}

func setHashMap(key string, fields map[string]string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	_, err := client.GrpcStorageClient.SetHashMap(ctx, &kvdbserver.SetHashMapRequest{Key: key, Fields: fields})
	client.CheckGrpcError(err)

	fmt.Println("OK")
}
