package connect

import (
	"fmt"

	"github.com/hollowdll/hakjdb/cmd/hakjctl/config"
	"github.com/spf13/cobra"
)

var cmdConnectShow = &cobra.Command{
	Use:   "show",
	Short: "Show connection settings",
	Long:  "Show the currently configured connection settings used to connect to a HakjDB server.",
	Example: `# Show all connection settings
hakjctl connect show`,
	Run: func(cmd *cobra.Command, args []string) {
		showConnectionSettings()
	},
}

func showConnectionSettings() {
	var output string
	output += fmt.Sprintf("host: %s\n", config.GetHost())
	output += fmt.Sprintf("port: %d\n", config.GetPort())
	output += fmt.Sprintf("database: %s", config.GetDefaultDB())

	fmt.Println(output)
}
