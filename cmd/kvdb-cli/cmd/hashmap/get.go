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

var cmdGetHashMapFieldValue = &cobra.Command{
	Use:   "get [key] [field ...]",
	Short: "Get HashMap field values",
	Long: "Gets the values of the specified fields in the HashMap stored at a key. " +
		"This command can return multiple values.",
	Args: cobra.MatchAll(cobra.MinimumNArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		getHashMapFieldValue(args[0], args[1:])
	},
}

func init() {
	cmdGetHashMapFieldValue.Flags().StringVarP(&dbName, "database", "d", "", "database to use")
}

func getHashMapFieldValue(key string, fields []string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	res, err := client.GrpcStorageClient.GetHashMapFieldValue(ctx, &kvdbserverpb.GetHashMapFieldValueRequest{Key: key, Fields: fields})
	client.CheckGrpcError(err)

	if res.Ok {
		if len(res.FieldValueMap) > 0 {
			var builder strings.Builder
			element := 0
			for field, value := range res.FieldValueMap {
				element++
				if value.Ok {
					builder.WriteString(fmt.Sprintf("%d) \"%s\": \"%s\"\n", element, field, value.Value))
				} else {
					builder.WriteString(fmt.Sprintf("%d) \"%s\": %s\n", element, field, client.ValueNone))
				}
			}
			fmt.Print(builder.String())
		}
	} else {
		fmt.Println(client.ValueNone)
	}
}
