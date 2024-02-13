package cmd

import (
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/cmd/connect"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/cmd/db"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/config"
	"github.com/hollowdll/kvdb/version"
	"github.com/spf13/cobra"
)

var (
	dbName  string
	rootCmd = &cobra.Command{
		Use:     "kvdb-cli",
		Short:   "CLI tool for kvdb key-value store",
		Long:    "CLI tool for kvdb key-value store",
		Version: version.Version,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(config.InitConfig, client.InitClient)

	rootCmd.AddCommand(db.CmdDb)
	rootCmd.AddCommand(connect.CmdConnect)
	rootCmd.AddCommand(cmdGetString)
	rootCmd.AddCommand(cmdSetString)
	rootCmd.AddCommand(cmdDeleteKey)
	rootCmd.AddCommand(cmdDeleteAllKeys)
	rootCmd.AddCommand(cmdInfo)

	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
