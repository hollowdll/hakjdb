package config

import (
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configFileName string = ".kvdb-cli"
	configFileType string = "yaml"

	// ConfigKeyHost is the configuration key for host.
	ConfigKeyHost string = "host"
	// ConfigKeyPort is the configuration key for port.
	ConfigKeyPort string = "port"
	// ConfigKeyDatabase is the configuration key for default database.
	ConfigKeyDatabase string = "default_db"

	// EnvPrefix is the prefix for environment variables.
	EnvPrefix string = "KVDBCLI"
	// EnvVarPassword is the environment variable for password.
	EnvVarPassword string = EnvPrefix + "_PASSWORD"

	// DefaultDatabase is the name of the default database to use.
	DefaultDatabase string = "default"
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
	viper.SetDefault(ConfigKeyDatabase, DefaultDatabase)

	viper.SafeWriteConfig()
	cobra.CheckErr(viper.ReadInConfig())
}
