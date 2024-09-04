package cmd

import (
	"github.com/hollowdll/hakjdb/cmd/hakjctl/client"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/cmd/connect"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/cmd/db"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/cmd/hashmap"
	"github.com/hollowdll/hakjdb/internal/common"
	"github.com/hollowdll/hakjdb/version"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	dbName  string
	rootCmd = &cobra.Command{
		Use:     "hakjctl",
		Short:   "CLI tool for HakjDB key-value data store",
		Long:    "CLI tool for HakjDB key-value data store. Use it to control and interact with HakjDB servers.",
		Version: version.Version,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(client.InitClient)

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
