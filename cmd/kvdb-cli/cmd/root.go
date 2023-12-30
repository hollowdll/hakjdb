package cmd

import (
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/cmd/db"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configFileName string = ".kvdb-cli"
	configFileType string = "json"
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
	cobra.OnInitialize(initConfig)

	viper.SetDefault("host", common.ServerDefaultHost)
	viper.SetDefault("port", common.ServerDefaultPort)

	rootCmd.AddCommand(db.CmdDb)
	rootCmd.AddCommand(cmdConnect)

	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func initConfig() {
	configDirPath, err := common.GetExecParentDirPath()
	cobra.CheckErr(err)
	// configFilePath := filepath.Join(configDirPath, fmt.Sprintf("%s.%s", configFileName, configFileType))

	viper.AddConfigPath(configDirPath)
	viper.SetConfigType(configFileType)
	viper.SetConfigName(configFileName)

	viper.SafeWriteConfig()
	err = viper.ReadInConfig()
	cobra.CheckErr(err)
}
