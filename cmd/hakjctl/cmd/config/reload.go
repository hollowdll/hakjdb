package config

import (
	"context"
	"fmt"

	"github.com/hollowdll/hakjdb/api/v1/serverpb"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/client"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/config"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdConfigReload = &cobra.Command{
	Use:   "reload",
	Short: "Reload configurations",
	Long:  "Reload configurations on the HakjDB server from its configuration sources.",
	Example: `# Reload configurations
hakjctl config reload`,
	Run: func(cmd *cobra.Command, args []string) {
		reloadServerConfig()
	},
}

func reloadServerConfig() {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, config.GetCmdTimeout())
	defer cancel()

	_, err := client.GrpcServerClient.ReloadConfig(ctx, &serverpb.ReloadConfigRequest{})
	client.CheckGrpcError(err)
	fmt.Println("OK")
}
