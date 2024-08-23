package hashmap

import (
	"context"
	"fmt"

	"github.com/hollowdll/hakjdb/api/v1/kvpb"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/client"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/config"
	"github.com/hollowdll/hakjdb/internal/common"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdDeleteHashMapFields = &cobra.Command{
	Use:   "delete KEY FIELD [FIELD ...]",
	Short: "Remove fields from a HashMap key value",
	Long: `Remove the specified fields from a HashMap key value.
Ignores fields that do not exist. This command can remove multiple fields.
Returns the number of fields that were removed.
`,
	Example: `# Use the default database
hakjctl hashmap delete key1 field1

# Specify the database to use
hakjctl hashmap delete key1 field2 -d default

# Remove multiple fields
hakjctl hashmap delete key1 field3 field4 field5`,
	Args: cobra.MatchAll(cobra.MinimumNArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		deleteHashMapFields(args[0], args[1:])
	},
}

func init() {
	cmdDeleteHashMapFields.Flags().StringVarP(&dbName, "database", "d", "", client.DBFlagMsg)
}

func deleteHashMapFields(key string, fields []string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, config.GetCmdTimeout())
	defer cancel()

	res, err := client.GrpcHashMapKVClient.DeleteHashMapFields(ctx, &kvpb.DeleteHashMapFieldsRequest{Key: key, Fields: fields})
	client.CheckGrpcError(err)

	if res.Ok {
		fmt.Printf("%d\n", res.FieldsRemovedCount)
	} else {
		fmt.Println(client.ValueNone)
	}
}
