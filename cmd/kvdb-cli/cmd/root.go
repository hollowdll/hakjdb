package cmd

import (
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/cmd/connect"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/cmd/db"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/cmd/hashmap"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/config"
	"github.com/hollowdll/kvdb/version"
	"github.com/spf13/cobra"
)

var (
	dbName  string
	rootCmd = &cobra.Command{
		Use:     "kvdb-cli",
		Short:   "CLI tool for kvdb key-value store",
		Long:    "CLI tool for kvdb key-value store. Use it to manage kvdb servers.",
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
	rootCmd.AddCommand(hashmap.CmdHashMap)
	rootCmd.AddCommand(cmdGetString)
	rootCmd.AddCommand(cmdSetString)
	rootCmd.AddCommand(cmdDeleteKeys)
	rootCmd.AddCommand(cmdInfo)
	rootCmd.AddCommand(cmdGetKeys)
	rootCmd.AddCommand(cmdLogs)
	rootCmd.AddCommand(cmdGetKeyType)
	rootCmd.AddCommand(cmdEcho)

	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
