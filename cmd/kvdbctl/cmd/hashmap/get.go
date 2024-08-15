package hashmap

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/kvdb/api/v0/kvpb"
	"github.com/hollowdll/kvdb/cmd/kvdbctl/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdGetHashMapFieldValues = &cobra.Command{
	Use:   "get KEY FIELD [FIELD ...]",
	Short: "Get field values of a HashMap key",
	Long: `Get the values of the specified fields of a HashMap key.
This command can return multiple values.
`,
	Args: cobra.MatchAll(cobra.MinimumNArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		getHashMapFieldValues(args[0], args[1:])
	},
}

func init() {
	cmdGetHashMapFieldValues.Flags().StringVarP(&dbName, "database", "d", "", client.DBFlagMsg)
}

func getHashMapFieldValues(key string, fields []string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	res, err := client.GrpcHashMapKVClient.GetHashMapFieldValues(ctx, &kvpb.GetHashMapFieldValuesRequest{
		Key:    key,
		Fields: fields,
	})
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
