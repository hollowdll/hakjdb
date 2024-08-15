package hashmap

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/api/v0/kvpb"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdSetHashMap = &cobra.Command{
	Use:   "set KEY FIELD VALUE [FIELD VALUE ...]",
	Short: "Set HashMap fields and values",
	Long: `Set the specified fields and their values of a HashMap key.
If the specified fields exist, they will be overwritten with the new values.
Creates the key if it doesn't exist.
Overwrites the key if it is holding a value of another data type.
This command can set multiple fields.
`,
	Args: cobra.MatchAll(cobra.MinimumNArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)%2 == 0 {
			cobra.CheckErr("invalid number of arguments")
		}

		fieldValueMap := make(map[string][]byte)
		for i := 1; i < len(args); i += 2 {
			fieldValueMap[args[i]] = []byte(args[i+1])
		}

		setHashMap(args[0], fieldValueMap)
	},
}

func init() {
	cmdSetHashMap.Flags().StringVarP(&dbName, "database", "d", "", client.DBFlagMsg)
}

func setHashMap(key string, fieldValueMap map[string][]byte) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	res, err := client.GrpcHashMapKVClient.SetHashMap(ctx, &kvpb.SetHashMapRequest{Key: key, FieldValueMap: fieldValueMap})
	client.CheckGrpcError(err)

	fmt.Printf("%d\n", res.FieldsAddedCount)
}
