package connect

import (
	"fmt"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdConnectShow = &cobra.Command{
	Use:   "show",
	Short: "Show connection settings",
	Long:  "Show connection settings to a kvdb server",
	Run: func(cmd *cobra.Command, args []string) {
		showConnectionSettings()
	},
}

func showConnectionSettings() {
	var output string
	output += fmt.Sprintf("Host: %s\n", viper.GetString(config.ConfigKeyHost))
	output += fmt.Sprintf("Port: %d", viper.GetUint16(config.ConfigKeyPort))
	fmt.Println(output)
}
