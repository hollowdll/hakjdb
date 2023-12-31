package connect

import (
	"github.com/spf13/cobra"
)

var CmdConnect = &cobra.Command{
	Use:   "connect [command]",
	Short: "Manage connection settings to a server",
	Long:  "Manage connection settings to a server",
}

func init() {
	CmdConnect.AddCommand(cmdConnectLs)
	CmdConnect.AddCommand(cmdConnectSet)
}
