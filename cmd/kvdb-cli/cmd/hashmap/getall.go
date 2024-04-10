package hashmap

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserverpb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdGetAllHashMapFieldsAndValues = &cobra.Command{
	Use:   "getall [key]",
	Short: "Get all the fields and values of a HashMap",
	Long:  "Get all the fields and values of a HashMap using a key",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		getAllHashMapFieldsAndValues(args[0])
	},
}

func init() {
	cmdGetAllHashMapFieldsAndValues.Flags().StringVarP(&dbName, "database", "d", "", "database to use")
}

func getAllHashMapFieldsAndValues(key string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	res, err := client.GrpcStorageClient.GetAllHashMapFieldsAndValues(ctx, &kvdbserverpb.GetAllHashMapFieldsAndValuesRequest{Key: key})
	client.CheckGrpcError(err)

	if res.Ok {
		if len(res.FieldValueMap) > 0 {
			var builder strings.Builder
			element := 0
			for field, value := range res.FieldValueMap {
				element++
				builder.WriteString(fmt.Sprintf("%d) \"%s\": \"%s\"\n", element, field, value))
			}
			fmt.Print(builder.String())
		}
	} else {
		fmt.Println(client.ValueNone)
	}
}
