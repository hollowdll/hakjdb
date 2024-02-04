package cmd

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdInfo = &cobra.Command{
	Use:   "info",
	Short: "Show information about the server",
	Long:  "Show information about the server",
	Run: func(cmd *cobra.Command, args []string) {
		showServerInfo()
	},
}

func showServerInfo() {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeoutSeconds)
	defer cancel()
	response, err := client.GrpcServerClient.GetServerInfo(ctx, &kvdbserver.GetServerInfoRequest{})
	client.CheckGrpcError(err)

	var info string
	info += fmt.Sprintf("kvdb_version: %s\n", response.Data.GetKvdbVersion())
	info += fmt.Sprintf("go_version: %s\n", response.Data.GetGoVersion())
	info += fmt.Sprintf("db_count: %d\n", response.Data.GetDbCount())
	info += fmt.Sprintf("total_data_size: %dB\n", response.Data.GetTotalDataSize())
	info += fmt.Sprintf("os: %s\n", response.Data.GetOs())
	info += fmt.Sprintf("arch: %s\n", response.Data.GetArch())
	info += fmt.Sprintf("process_id: %d\n", response.Data.GetProcessId())
	info += fmt.Sprintf("uptime_seconds: %d\n", response.Data.GetUptimeSeconds())
	info += fmt.Sprintf("tcp_port: %d", response.Data.GetTcpPort())

	fmt.Println(info)
}
