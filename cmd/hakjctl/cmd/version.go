package cmd

import (
	"context"
	"fmt"

	"github.com/hollowdll/hakjdb/api/v1/serverpb"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/client"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/config"
	"github.com/hollowdll/hakjdb/version"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  "Show the client, server, and API versions.",
	Example: `# Show all version information
hakjctl version`,
	Run: func(cmd *cobra.Command, args []string) {
		showVersionInfo()
	},
}

func showVersionInfo() {
	fmt.Printf("hakjctl version: %s\n", version.Version)

	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, config.GetCmdTimeout())
	defer cancel()
	res, err := client.GrpcServerClient.GetServerInfo(ctx, &serverpb.GetServerInfoRequest{})
	client.CheckGrpcError(err)

	fmt.Printf("server version: %s\n", res.GeneralInfo.ServerVersion)
	fmt.Printf("API version: %s\n", res.GeneralInfo.ApiVersion)
}
