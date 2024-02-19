package server

import (
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/viper"
)

const (
	dataDirName    string = "data"
	configFileName string = ".kvdbserver"
	configFileType string = "json"
	logFileName    string = "kvdb.log"

	// EnvPrefix is the prefix that environment variables use.
	EnvPrefix string = "KVDB"
	// ConfigKeyPort is the configuration key for port.
	ConfigKeyPort string = "port"
	// ConfigKeyDebugEnabled is the configuration key for debug mode.
	ConfigKeyDebugEnabled string = "debug_enabled"
	// ConfigKeyDefaultDatabase is the configuration key for default database.
	ConfigKeyDefaultDatabase string = "default_db"
	// ConfigKeyLogFileEnabled is the configuration key for enabling log file.
	ConfigKeyLogFileEnabled string = "logfile_enabled"

	// DefaultDatabase is the name of the default database.
	DefaultDatabase string = "default"

	// EnvVarPassword is the environment variable for server password.
	EnvVarPassword string = EnvPrefix + "_PASSWORD"
)

// initConfig initializes server configurations.
func initConfig(s *Server) {
	parentDir, err := common.GetExecParentDirPath()
	if err != nil {
		s.logger.Fatalf("Failed to get executable's parent directory: %s", err)
	}
	configDirPath, err := common.GetDirPath(parentDir, dataDirName)
	if err != nil {
		s.logger.Fatalf("Failed to get data directory: %s", err)
	}

	viper.AddConfigPath(configDirPath)
	viper.SetConfigType(configFileType)
	viper.SetConfigName(configFileName)

	viper.SetDefault(ConfigKeyPort, common.ServerDefaultPort)
	viper.SetDefault(ConfigKeyDebugEnabled, false)
	viper.SetDefault(ConfigKeyDefaultDatabase, DefaultDatabase)
	viper.SetDefault(ConfigKeyLogFileEnabled, false)

	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv()

	viper.SafeWriteConfig()
	if err = viper.ReadInConfig(); err != nil {
		s.logger.Fatalf("Failed to load configuration: %s", err)
	}
}
