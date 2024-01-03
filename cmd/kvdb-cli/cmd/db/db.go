package db

import (
	"github.com/spf13/cobra"
)

var (
	dbName string
	CmdDb  = &cobra.Command{
		Use:   "db [command]",
		Short: "Manage databases",
		Long:  "Manage databases",
	}
)

func init() {
	CmdDb.AddCommand(cmdDbCreate)
	CmdDb.AddCommand(cmdDbLs)
	CmdDb.AddCommand(cmdDbInfo)
}
