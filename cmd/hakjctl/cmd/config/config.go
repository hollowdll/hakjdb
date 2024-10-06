package config

import (
	"github.com/spf13/cobra"
)

var CmdConfig = &cobra.Command{
	Use:   "config",
	Short: "Manage configurations",
	Long:  "Manage HakjDB server configurations.",
}

func init() {
	CmdConfig.AddCommand(cmdConfigReload)
}
