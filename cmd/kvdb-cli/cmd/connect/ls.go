package connect

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdConnectLs = &cobra.Command{
	Use:   "ls",
	Short: "Show connection settings to a server",
	Long:  "Show connection settings to a server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(formatConnectionSettings())
	},
}

func formatConnectionSettings() string {
	return fmt.Sprintf("Host: %s\nPort: %d", viper.GetString("host"), viper.GetUint16("port"))
}
