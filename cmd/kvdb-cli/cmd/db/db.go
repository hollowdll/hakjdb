package db

import (
	"github.com/spf13/cobra"
)

var CmdDb = &cobra.Command{
	Use:   "db [command]",
	Short: "Manage databases",
	Long:  "Manage databases",
}

func init() {
	CmdDb.AddCommand(cmdCreateDb)
	CmdDb.AddCommand(cmdDbLs)
}
