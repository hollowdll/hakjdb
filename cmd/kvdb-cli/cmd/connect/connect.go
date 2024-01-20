package connect

import (
	"github.com/spf13/cobra"
)

var CmdConnect = &cobra.Command{
	Use:   "connect",
	Short: "Manage connection settings",
	Long:  "Manage connection settings to a kvdb server",
}

func init() {
	CmdConnect.AddCommand(cmdConnectLs)
	CmdConnect.AddCommand(cmdConnectSet)
}
