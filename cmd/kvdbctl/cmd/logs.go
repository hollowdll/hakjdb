package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/kvdb/api/v0/serverpb"
	"github.com/hollowdll/kvdb/cmd/kvdbctl/client"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdLogs = &cobra.Command{
	Use:   "logs",
	Short: "Get logs from the server",
	Long:  "Get logs from the server if the server's log file is enabled. Currently gets all the logs.",
	Example: `# Get all logs
kvdbctl logs`,
	Run: func(cmd *cobra.Command, args []string) {
		getLogs()
	},
}

func getLogs() {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	res, err := client.GrpcServerClient.GetLogs(ctx, &serverpb.GetLogsRequest{})
	client.CheckGrpcError(err)

	if len(res.Logs) > 0 {
		fmt.Println(strings.Join(res.Logs, "\n"))
	}
}
