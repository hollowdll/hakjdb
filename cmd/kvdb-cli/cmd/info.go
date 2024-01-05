package cmd

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/cobra"
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
	ctx, cancel := context.WithTimeout(context.Background(), client.ClientCtxTimeout)
	defer cancel()
	response, err := client.GrpcServerClient.GetServerInfo(ctx, &kvdbserver.GetServerInfoRequest{})
	cobra.CheckErr(err)

	var info string
	info += fmt.Sprintf("server_version: %s\n", response.Info.GetVersion())
	info += fmt.Sprintf("go_version: %s\n", response.Info.GetGoVersion())
	info += fmt.Sprintf("db_count: %d\n", response.Info.GetDbCount())
	info += fmt.Sprintf("total_data_size: %dB\n", response.Info.GetTotalDataSize())
	info += fmt.Sprintf("os: %s\n", response.Info.GetOs())
	info += fmt.Sprintf("arch: %s\n", response.Info.GetArch())
	info += fmt.Sprintf("process_id: %d\n", response.Info.GetProcessId())
	info += fmt.Sprintf("uptime_seconds: %d\n", response.Info.GetUptimeSeconds())
	info += fmt.Sprintf("tcp_port: %d", response.Info.GetTcpPort())

	fmt.Println(info)
}
