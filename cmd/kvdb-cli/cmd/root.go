package cmd

import (
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/cmd/db"
	"github.com/hollowdll/kvdb/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "kvdb-cli",
	Short:   "CLI tool for kvdb key-value database",
	Long:    "CLI tool for kvdb key-value database",
	Version: version.Version,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(db.CmdDb)
	rootCmd.AddCommand(cmdConnect)

	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
