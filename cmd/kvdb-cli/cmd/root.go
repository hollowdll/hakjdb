package cmd

import (
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/cmd/db"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "kvdb-cli",
	Short:   "CLI tool for kvdb key-value database",
	Long:    "CLI tool for kvdb key-value database",
	Version: "0.0.0",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(db.CmdDb)

	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
