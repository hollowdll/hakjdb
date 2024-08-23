package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/hakjdb/api/v1/serverpb"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/client"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/config"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdLogs = &cobra.Command{
	Use:   "logs",
	Short: "Get logs from the server",
	Long:  "Get logs from the server if the server's log file is enabled. Currently gets all the logs.",
	Example: `# Get all logs
hakjctl logs`,
	Run: func(cmd *cobra.Command, args []string) {
		getLogs()
	},
}

func getLogs() {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, config.GetCmdTimeout())
	defer cancel()

	res, err := client.GrpcServerClient.GetLogs(ctx, &serverpb.GetLogsRequest{})
	client.CheckGrpcError(err)

	if len(res.Logs) > 0 {
		fmt.Println(strings.Join(res.Logs, "\n"))
	}
}
