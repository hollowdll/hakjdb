package cmd

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/api/v0/kvpb"
	"github.com/hollowdll/kvdb/cmd/kvdbctl/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var (
	cmdDeleteKeys = &cobra.Command{
		Use:   "delete [KEY ...]",
		Short: "Delete keys",
		Long: `Delete the specified keys and the values they are holding.
Ignores keys that do not exist.
This command can delete multiple keys.
All the keys of a database can be deleted with --all option.
Returns the number of keys that were deleted or OK if all the keys were deleted.
`,
		Example: `# Use the default database
kvdbctl delete key1

# Specify the database to use
kvdbctl delete key2 --database default

# Delete multiple keys
kvdbctl delete key3 key4 key5

# Delete all the keys
kvdbctl delete --all`,
		Run: func(cmd *cobra.Command, args []string) {
			deleteKeys(args)
		},
	}
	deleteAll bool = false
)

func init() {
	cmdDeleteKeys.Flags().StringVarP(&dbName, "database", "d", "", client.DBFlagMsg)
	cmdDeleteKeys.Flags().BoolVar(&deleteAll, "all", false, "Delete all the keys of the database that is being used")
}

func deleteKeys(keys []string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	if deleteAll {
		dbName := md.Get(common.GrpcMetadataKeyDbName)[0]
		if !client.PromptConfirm(fmt.Sprintf("Delete all the keys of database '%s'? Yes/No: ", dbName)) {
			return
		}
		_, err := client.GrpcGeneralKVClient.DeleteAllKeys(ctx, &kvpb.DeleteAllKeysRequest{})
		client.CheckGrpcError(err)
		fmt.Println("OK")
		return
	}

	res, err := client.GrpcGeneralKVClient.DeleteKeys(ctx, &kvpb.DeleteKeysRequest{Keys: keys})
	client.CheckGrpcError(err)
	fmt.Printf("%d\n", res.KeysDeletedCount)
}
