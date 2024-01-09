package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdGetString = &cobra.Command{
	Use:   "get [key]",
	Short: "Get a string value",
	Long:  "Get a string value",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		getString(args[0])
	},
}

func init() {
	cmdGetString.Flags().StringVarP(&dbName, "db", "d", "", "name of the database")
}

func getString(key string) {
	// Send database name in metadata
	md := metadata.Pairs(common.GrpcMetadataKeyDbName, dbName)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.ClientCtxTimeout)
	defer cancel()

	response, err := client.GrpcStorageClient.GetString(ctx, &kvdbserver.GetStringRequest{Key: key})
	client.CheckGrpcError(err)

	fmt.Fprintf(os.Stdout, "\"%s\"\n", response.GetValue())
}
