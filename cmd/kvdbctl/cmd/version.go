package cmd

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/api/v0/serverpb"
	"github.com/hollowdll/kvdb/cmd/kvdbctl/client"
	"github.com/hollowdll/kvdb/version"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  "Show the client, server, and API versions.",
	Run: func(cmd *cobra.Command, args []string) {
		showVersionInfo()
	},
}

func showVersionInfo() {
	fmt.Printf("kvdbctl version: %s\n", version.Version)

	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()
	res, err := client.GrpcServerClient.GetServerInfo(ctx, &serverpb.GetServerInfoRequest{})
	client.CheckGrpcError(err)

	fmt.Printf("server version: %s\n", res.GeneralInfo.KvdbVersion)
	fmt.Printf("API version: %s\n", res.GeneralInfo.ApiVersion)
}
