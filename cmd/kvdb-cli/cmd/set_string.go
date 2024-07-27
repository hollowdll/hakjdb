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

var cmdSetString = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set a string value",
	Long: `
Sets a key to hold a String value.
Creates the key if it doesn't exist.
Overwrites the key if it is holding a value of another data type.
`,
	Args: cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		setString(args[0], args[1])
	},
}

func init() {
	cmdSetString.Flags().StringVarP(&dbName, "database", "d", "", "database to use")
}

func setString(key string, value string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	_, err := client.GrpcStringKeyClient.SetString(ctx, &storagepb.SetStringRequest{Key: key, Value: value})
	client.CheckGrpcError(err)

	fmt.Println("OK")
}
