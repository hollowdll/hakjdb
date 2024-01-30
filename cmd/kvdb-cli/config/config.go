package config

import (
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configFileName string = ".kvdb-cli"
	configFileType string = "json"
	// ConfigKeyHost is the configuration key for host.
	ConfigKeyHost string = "host"
	// ConfigKeyPort is the configuration key for port.
	ConfigKeyPort string = "port"
)

// InitConfig initializes and loads configurations.
func InitConfig() {
	configDirPath, err := common.GetExecParentDirPath()
	cobra.CheckErr(err)

	viper.AddConfigPath(configDirPath)
	viper.SetConfigType(configFileType)
	viper.SetConfigName(configFileName)

	viper.SetDefault(ConfigKeyHost, common.ServerDefaultHost)
	viper.SetDefault(ConfigKeyPort, common.ServerDefaultPort)

	viper.SafeWriteConfig()
	cobra.CheckErr(viper.ReadInConfig())
}
