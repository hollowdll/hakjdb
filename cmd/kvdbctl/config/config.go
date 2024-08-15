package config

import (
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configFileName string = "kvdbctl-config"
	configFileType string = "yaml"

	// ConfigKeyHost is the configuration key for host.
	ConfigKeyHost string = "host"
	// ConfigKeyPort is the configuration key for port.
	ConfigKeyPort string = "port"
	// ConfigKeyDatabase is the configuration key for default database.
	ConfigKeyDatabase string = "default_db"
	// ConfigKeyTlsEnabled is the configuration key for enabling TLS.
	ConfigKeyTlsEnabled string = "tls_enabled"
	// ConfigKeyTlsCertPath is the configuration key for TLS certificate path.
	ConfigKeyTlsCertPath string = "tls_cert_path"

	// EnvPrefix is the prefix for environment variables.
	EnvPrefix string = "KVDBCTL"
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
	viper.SetDefault(ConfigKeyTlsEnabled, false)
	viper.SetDefault(ConfigKeyTlsCertPath, "")

	viper.SafeWriteConfig()
	cobra.CheckErr(viper.ReadInConfig())
}
