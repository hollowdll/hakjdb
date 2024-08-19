package cmd

import (
	"github.com/hollowdll/kvdb/cmd/kvdbctl/client"
	"github.com/hollowdll/kvdb/cmd/kvdbctl/cmd/connect"
	"github.com/hollowdll/kvdb/cmd/kvdbctl/cmd/db"
	"github.com/hollowdll/kvdb/cmd/kvdbctl/cmd/hashmap"
	"github.com/hollowdll/kvdb/cmd/kvdbctl/config"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/version"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	dbName  string
	rootCmd = &cobra.Command{
		Use:     "kvdbctl",
		Short:   "CLI tool for kvdb key-value data store",
		Long:    "CLI tool for kvdb key-value data store. Use it to manage kvdb servers.",
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
	rootCmd.AddCommand(cmdAuthenticate)
	rootCmd.AddCommand(cmdVersion)

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.DisableAutoGenTag = true

	dir := createGeneratedDocsDir()
	err := doc.GenMarkdownTree(rootCmd, dir)
	cobra.CheckErr(err)
}

func createGeneratedDocsDir() string {
	parentDir, err := common.GetExecParentDirPath()
	cobra.CheckErr(err)
	dir, err := common.GetDirPath(parentDir, "./generated-docs")
	cobra.CheckErr(err)
	return dir
}
