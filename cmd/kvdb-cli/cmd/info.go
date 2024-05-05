package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/internal/common"
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
	info.WriteString("** General **\n")
	info.WriteString(fmt.Sprintf("kvdb_version: %s\n", res.Data.GeneralInfo.KvdbVersion))
	info.WriteString(fmt.Sprintf("go_version: %s\n", res.Data.GeneralInfo.GoVersion))
	info.WriteString(fmt.Sprintf("db_count: %d\n", res.Data.GeneralInfo.DbCount))
	info.WriteString(fmt.Sprintf("os: %s\n", res.Data.GeneralInfo.Os))
	info.WriteString(fmt.Sprintf("arch: %s\n", res.Data.GeneralInfo.Arch))
	info.WriteString(fmt.Sprintf("process_id: %d\n", res.Data.GeneralInfo.ProcessId))
	info.WriteString(fmt.Sprintf("uptime_seconds: %d\n", res.Data.GeneralInfo.UptimeSeconds))
	info.WriteString(fmt.Sprintf("tcp_port: %d\n", res.Data.GeneralInfo.TcpPort))
	info.WriteString(fmt.Sprintf("default_db: %s\n", res.Data.GeneralInfo.DefaultDb))

	if res.Data.GeneralInfo.TlsEnabled {
		info.WriteString("tls_enabled: yes\n")
	} else {
		info.WriteString("tls_enabled: no\n")
	}

	if res.Data.GeneralInfo.PasswordEnabled {
		info.WriteString("password_enabled: yes\n")
	} else {
		info.WriteString("password_enabled: no\n")
	}

	if res.Data.GeneralInfo.LogfileEnabled {
		info.WriteString("logfile_enabled: yes\n")
	} else {
		info.WriteString("logfile_enabled: no\n")
	}

	if res.Data.GeneralInfo.DebugEnabled {
		info.WriteString("debug_enabled: yes\n")
	} else {
		info.WriteString("debug_enabled: no\n")
	}

	info.WriteString("\n** Data storage **\n")
	info.WriteString(fmt.Sprintf("total_data_size: %d B\n", res.Data.StorageInfo.TotalDataSize))
	info.WriteString(fmt.Sprintf("total_keys: %d\n", res.Data.StorageInfo.TotalKeys))

	info.WriteString("\n** Clients **\n")
	info.WriteString(fmt.Sprintf("client_connections: %d\n", res.Data.ClientInfo.ClientConnections))
	info.WriteString(fmt.Sprintf("max_client_connections: %d\n", res.Data.ClientInfo.MaxClientConnections))

	info.WriteString("\n** Memory **\n")
	info.WriteString(fmt.Sprintf("memory_alloc: %.1f MB\n", common.BytesToMegabytes(res.Data.MemoryInfo.MemoryAlloc)))
	info.WriteString(fmt.Sprintf("memory_total_alloc: %.1f MB\n", common.BytesToMegabytes(res.Data.MemoryInfo.MemoryTotalAlloc)))
	info.WriteString(fmt.Sprintf("memory_sys: %.1f MB", common.BytesToMegabytes(res.Data.MemoryInfo.MemorySys)))

	fmt.Println(info.String())
}
