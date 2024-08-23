package cmd

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

var cmdGetString = &cobra.Command{
	Use:   "get KEY",
	Short: "Get the value of a String key",
	Long:  "Get the value of a String key. Returns (None) if the key doesn't exist.",
	Example: `# Use the default database
hakjctl get key1

# Specify the database to use
hakjctl get key1 -d default`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		getString(args[0])
	},
}

func init() {
	cmdGetString.Flags().StringVarP(&dbName, "database", "d", "", client.DBFlagMsg)
}

func getString(key string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, config.GetCmdTimeout())
	defer cancel()

	response, err := client.GrpcStringKVClient.GetString(ctx, &kvpb.GetStringRequest{Key: key})
	client.CheckGrpcError(err)

	if response.Ok {
		fmt.Printf("\"%s\"\n", response.Value)
	} else {
		fmt.Println(client.ValueNone)
	}
}
