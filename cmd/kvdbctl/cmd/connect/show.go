package connect

import (
	"fmt"

	"github.com/hollowdll/kvdb/cmd/kvdbctl/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdConnectShow = &cobra.Command{
	Use:   "show",
	Short: "Show connection settings",
	Long:  "Show the currently configured connection settings used to connect to a server.",
	Run: func(cmd *cobra.Command, args []string) {
		showConnectionSettings()
	},
}

func showConnectionSettings() {
	var output string
	output += fmt.Sprintf("host: %s\n", viper.GetString(config.ConfigKeyHost))
	output += fmt.Sprintf("port: %d\n", viper.GetUint16(config.ConfigKeyPort))
	output += fmt.Sprintf("database: %s", viper.GetString(config.ConfigKeyDatabase))

	fmt.Println(output)
}
