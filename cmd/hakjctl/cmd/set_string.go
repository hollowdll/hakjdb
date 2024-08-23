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

var cmdSetString = &cobra.Command{
	Use:   "set KEY VALUE",
	Short: "Set the value of a String key",
	Long: `Set the value of a String key.
Creates the key if it doesn't exist.
Overwrites the key if it is holding a value of another data type.
`,
	Example: `# Use the default database
hakjctl set key1 "Hello world!"

# Specify the database to use
hakjctl set key2 "value123" -d default`,
	Args: cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		setString(args[0], []byte(args[1]))
	},
}

func init() {
	cmdSetString.Flags().StringVarP(&dbName, "database", "d", "", client.DBFlagMsg)
}

func setString(key string, value []byte) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, config.GetCmdTimeout())
	defer cancel()

	_, err := client.GrpcStringKVClient.SetString(ctx, &kvpb.SetStringRequest{Key: key, Value: value})
	client.CheckGrpcError(err)

	fmt.Println("OK")
}
