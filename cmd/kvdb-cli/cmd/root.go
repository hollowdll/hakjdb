package cmd

import (
	"github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/cmd/db"
	"github.com/hollowdll/kvdb/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const configFileName string = ".kvdb-cli"

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
	cobra.OnInitialize(initConfig)

	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 12345)

	rootCmd.AddCommand(db.CmdDb)
	rootCmd.AddCommand(cmdConnect)

	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func initConfig() {
	configDirPath, err := kvdb.GetExecParentDirPath()
	cobra.CheckErr(err)

	viper.AddConfigPath(configDirPath)
	viper.SetConfigType("json")
	viper.SetConfigName(configFileName)

	err = viper.ReadInConfig()
	cobra.CheckErr(err)
}
