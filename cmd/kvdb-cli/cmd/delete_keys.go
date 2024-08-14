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

var (
	cmdDeleteKeys = &cobra.Command{
		Use:   "delete [key ...]",
		Short: "Delete keys",
		Long: `Deletes the specified keys and the values they are holding.
Ignores keys that do not exist.
This command can delete multiple keys.
All the keys of a database can be deleted with --all option.
`,
		Run: func(cmd *cobra.Command, args []string) {
			deleteKeys(args)
		},
	}
	deleteAll bool = false
)

func init() {
	cmdDeleteKeys.Flags().StringVarP(&dbName, "database", "d", "", "The database to use. If not present, the default database is used")
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
