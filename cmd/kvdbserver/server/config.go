package server

import (
	"path/filepath"

	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/viper"
)

const (
	dataDirName    string = "data"
	configFileName string = ".kvdbserver"
	configFileType string = "yaml"
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
	// ConfigKeyTlsEnabled is the configuration key for enabling TLS.
	ConfigKeyTlsEnabled string = "tls_enabled"
	// ConfigKeyTlsCertPath is the configuration key for TLS certificate file path.
	ConfigKeyTlsCertPath string = "tls_cert_path"
	// ConfigKeyTlsPrivKeyPath is the configuration key for TLS private key file path.
	ConfigKeyTlsPrivKeyPath string = "tls_private_key_path"
	// ConfigKeyMaxClientConnections is the configuration key for maximum client connections.
	ConfigKeyMaxClientConnections string = "max_client_connections"

	// DefaultDatabase is the name of the default database.
	DefaultDatabase string = "default"

	// EnvVarPassword is the environment variable for server password.
	EnvVarPassword string = EnvPrefix + "_PASSWORD"
)

// initConfig initializes server configurations.
func initConfig(s *Server) {
	parentDir, err := common.GetExecParentDirPath()
	if err != nil {
		s.logger.Fatalf("Failed to get executable's parent directory: %v", err)
	}
	dataDirPath, err := common.GetDirPath(parentDir, dataDirName)
	if err != nil {
		s.logger.Fatalf("Failed to get data directory: %v", err)
	}

	s.SetLogFilePath(filepath.Join(dataDirPath, logFileName))

	viper.AddConfigPath(dataDirPath)
	viper.SetConfigType(configFileType)
	viper.SetConfigName(configFileName)

	viper.SetDefault(ConfigKeyPort, common.ServerDefaultPort)
	viper.SetDefault(ConfigKeyDebugEnabled, false)
	viper.SetDefault(ConfigKeyDefaultDatabase, DefaultDatabase)
	viper.SetDefault(ConfigKeyLogFileEnabled, false)
	viper.SetDefault(ConfigKeyTlsEnabled, false)
	viper.SetDefault(ConfigKeyTlsCertPath, "")
	viper.SetDefault(ConfigKeyTlsPrivKeyPath, "")
	viper.SetDefault(ConfigKeyMaxClientConnections, common.DefaultMaxClientConnections)

	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv()

	viper.SafeWriteConfig()
	if err = viper.ReadInConfig(); err != nil {
		s.logger.Fatalf("Failed to load configuration: %v", err)
	}
}
