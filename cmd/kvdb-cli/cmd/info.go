package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/proto/kvdbserverpb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdInfo = &cobra.Command{
	Use:   "info",
	Short: "Show information about the server",
	Long:  "Shows information about the kvdb server.",
	Run: func(cmd *cobra.Command, args []string) {
		showServerInfo()
	},
}

func showServerInfo() {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()
	res, err := client.GrpcServerClient.GetServerInfo(ctx, &kvdbserverpb.GetServerInfoRequest{})
	client.CheckGrpcError(err)

	var info strings.Builder
	info.WriteString(fmt.Sprintf("kvdb_version: %s\n", res.Data.KvdbVersion))
	info.WriteString(fmt.Sprintf("go_version: %s\n", res.Data.GoVersion))
	info.WriteString(fmt.Sprintf("db_count: %d\n", res.Data.DbCount))
	info.WriteString(fmt.Sprintf("total_data_size: %dB\n", res.Data.TotalDataSize))
	info.WriteString(fmt.Sprintf("os: %s\n", res.Data.Os))
	info.WriteString(fmt.Sprintf("arch: %s\n", res.Data.Arch))
	info.WriteString(fmt.Sprintf("process_id: %d\n", res.Data.ProcessId))
	info.WriteString(fmt.Sprintf("uptime_seconds: %d\n", res.Data.UptimeSeconds))
	info.WriteString(fmt.Sprintf("tcp_port: %d\n", res.Data.TcpPort))
	info.WriteString(fmt.Sprintf("default_db: %s\n", res.Data.DefaultDb))

	if res.Data.TlsEnabled {
		info.WriteString("tls_enabled: yes\n")
	} else {
		info.WriteString("tls_enabled: no\n")
	}

	if res.Data.PasswordEnabled {
		info.WriteString("password_enabled: yes\n")
	} else {
		info.WriteString("password_enabled: no\n")
	}

	if res.Data.LogfileEnabled {
		info.WriteString("logfile_enabled: yes\n")
	} else {
		info.WriteString("logfile_enabled: no\n")
	}

	if res.Data.DebugEnabled {
		info.WriteString("debug_enabled: yes")
	} else {
		info.WriteString("debug_enabled: no")
	}

	fmt.Println(info.String())
}
