package cmd

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdSetString = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set a string value",
	Long:  "Set a string value",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		setString(args[0], args[1])
	},
}

func init() {
	cmdSetString.Flags().StringVarP(&dbName, "db", "d", "", "name of the database")
}

func setString(key string, value string) {
	// Send database name in metadata
	md := metadata.Pairs(common.GrpcMetadataKeyDbName, dbName)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.ClientCtxTimeout)
	defer cancel()

	_, err := client.GrpcStorageClient.SetString(ctx, &kvdbserver.SetStringRequest{Key: key, Value: value})
	cobra.CheckErr(err)

	fmt.Println("OK")
}
