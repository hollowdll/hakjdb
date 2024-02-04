package db

import (
	"context"
	"fmt"
	"time"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdDbInfo = &cobra.Command{
	Use:   "info",
	Short: "Show information about a database",
	Long:  "Show information about a database",
	Run: func(cmd *cobra.Command, args []string) {
		showDbInfo()
	},
}

func init() {
	cmdDbInfo.Flags().StringVarP(&dbName, "name", "n", "", "name of the database (required)")
	cmdDbInfo.MarkFlagRequired("name")
}

func showDbInfo() {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeoutSeconds)
	defer cancel()
	response, err := client.GrpcDatabaseClient.GetDatabaseInfo(ctx, &kvdbserver.GetDatabaseInfoRequest{DbName: dbName})
	client.CheckGrpcError(err)

	var info string
	info += fmt.Sprintf("name: %s\n", response.Data.GetName())
	info += fmt.Sprintf("created_at: %s00:00\n", response.Data.GetCreatedAt().AsTime().Format(time.RFC3339))
	info += fmt.Sprintf("updated_at: %s00:00\n", response.Data.GetUpdatedAt().AsTime().Format(time.RFC3339))
	info += fmt.Sprintf("key_count: %d\n", response.Data.GetKeyCount())
	info += fmt.Sprintf("data_size: %dB", response.Data.GetDataSize())

	fmt.Println(info)
}
