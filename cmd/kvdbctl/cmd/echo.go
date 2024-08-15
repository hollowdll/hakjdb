package cmd

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/api/v0/echopb"
	"github.com/hollowdll/kvdb/cmd/kvdbctl/client"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdEcho = &cobra.Command{
	Use:   "echo [MESSAGE]",
	Short: "Test connection",
	Long: `Test connection to the server. Sends a message to the server and returns the same message back.
Can be useful for verifying that the server is still alive and can process requests.
`,
	Example: `# Send an empty message
kvdbctl echo
""

# Send message "Hello"
kvdbctl echo "Hello"
"Hello"`,
	Run: func(cmd *cobra.Command, args []string) {
		msg := ""
		if len(args) > 0 {
			msg = args[0]
		}
		echo(msg)
	},
}

func echo(msg string) {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	res, err := client.GrpcEchoClient.UnaryEcho(ctx, &echopb.UnaryEchoRequest{Msg: msg})
	client.CheckGrpcError(err)
	fmt.Printf("\"%s\"\n", res.Msg)
}
