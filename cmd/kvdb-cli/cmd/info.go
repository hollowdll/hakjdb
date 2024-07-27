package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/kvdb/api/v0/serverpb"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdInfo = &cobra.Command{
	Use:   "info",
	Short: "Show information about the server",
	Long:  "Shows information about the server.",
	Run: func(cmd *cobra.Command, args []string) {
		showServerInfo()
	},
}

func showServerInfo() {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()
	res, err := client.GrpcServerClient.GetServerInfo(ctx, &serverpb.GetServerInfoRequest{})
	client.CheckGrpcError(err)

	var info strings.Builder
	info.WriteString("** General **\n")
	info.WriteString(fmt.Sprintf("kvdb_version: %s\n", res.GeneralInfo.KvdbVersion))
	info.WriteString(fmt.Sprintf("go_version: %s\n", res.GeneralInfo.GoVersion))
	info.WriteString(fmt.Sprintf("os: %s\n", res.GeneralInfo.Os))
	info.WriteString(fmt.Sprintf("arch: %s\n", res.GeneralInfo.Arch))
	info.WriteString(fmt.Sprintf("process_id: %d\n", res.GeneralInfo.ProcessId))
	info.WriteString(fmt.Sprintf("uptime_seconds: %d\n", res.GeneralInfo.UptimeSeconds))
	info.WriteString(fmt.Sprintf("tcp_port: %d\n", res.GeneralInfo.TcpPort))

	if res.GeneralInfo.TlsEnabled {
		info.WriteString("tls_enabled: yes\n")
	} else {
		info.WriteString("tls_enabled: no\n")
	}

	if res.GeneralInfo.PasswordEnabled {
		info.WriteString("password_enabled: yes\n")
	} else {
		info.WriteString("password_enabled: no\n")
	}

	if res.GeneralInfo.LogfileEnabled {
		info.WriteString("logfile_enabled: yes\n")
	} else {
		info.WriteString("logfile_enabled: no\n")
	}

	if res.GeneralInfo.DebugEnabled {
		info.WriteString("debug_enabled: yes\n")
	} else {
		info.WriteString("debug_enabled: no\n")
	}

	info.WriteString("\n** Databases **\n")
	info.WriteString(fmt.Sprintf("db_count: %d\n", res.DbInfo.DbCount))
	info.WriteString(fmt.Sprintf("default_db: %s\n", res.DbInfo.DefaultDb))

	info.WriteString("\n** Data storage **\n")
	info.WriteString(fmt.Sprintf("total_data_size: %d B\n", res.StorageInfo.TotalDataSize))
	info.WriteString(fmt.Sprintf("total_keys: %d\n", res.StorageInfo.TotalKeys))

	info.WriteString("\n** Clients **\n")
	info.WriteString(fmt.Sprintf("client_connections: %d\n", res.ClientInfo.ClientConnections))
	info.WriteString(fmt.Sprintf("max_client_connections: %d\n", res.ClientInfo.MaxClientConnections))

	info.WriteString("\n** Memory **\n")
	info.WriteString(fmt.Sprintf("memory_alloc: %.1f MB\n", common.BytesToMegabytes(res.MemoryInfo.MemoryAlloc)))
	info.WriteString(fmt.Sprintf("memory_total_alloc: %.1f MB\n", common.BytesToMegabytes(res.MemoryInfo.MemoryTotalAlloc)))
	info.WriteString(fmt.Sprintf("memory_sys: %.1f MB", common.BytesToMegabytes(res.MemoryInfo.MemorySys)))

	fmt.Println(info.String())
}
