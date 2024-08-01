package config

import (
	"os"
	"path/filepath"

	"github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/viper"
)

const (
	dataDirName    string = "data"
	configFileName string = "kvdbserver-config"
	configFileType string = "yaml"
	logFileName    string = "kvdbserver.log"

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
	ConfigKeyTLSEnabled string = "tls_enabled"
	// ConfigKeyTlsCertPath is the configuration key for TLS certificate file path.
	ConfigKeyTLSCertPath string = "tls_cert_path"
	// ConfigKeyTlsPrivKeyPath is the configuration key for TLS private key file path.
	ConfigKeyTLSPrivKeyPath string = "tls_private_key_path"
	// ConfigKeyMaxClientConnections is the configuration key for maximum client connections.
	ConfigKeyMaxClientConnections string = "max_client_connections"

	// DefaultDatabase is the name of the default database.
	DefaultDatabase string = "default"

	// EnvVarPassword is the environment variable for server password.
	EnvVarPassword string = EnvPrefix + "_PASSWORD"
)

// ServerConfig holds the server's configuration.
type ServerConfig struct {
	LogFileEnabled bool
	TLSEnabled     bool
	DebugEnabled   bool
	// The name of the default database that is created at server startup.
	DefaultDB   string
	LogFilePath string
	// The maximum number of keys a database can hold.
	MaxKeysPerDB uint32
	// The maximum number of fields a HashMap can hold.
	MaxHashMapFields uint32
	// The TCP/IP port the server listens at.
	PortInUse            uint16
	MaxClientConnections uint32
	TLSCertPath          string
	TLSPrivKeyPath       string
}

// LoadConfig loads server configurations.
func LoadConfig(lg kvdb.Logger) ServerConfig {
	lg.Infof("Loading configurations ...")
	parentDir, err := common.GetExecParentDirPath()
	if err != nil {
		lg.Fatalf("Failed to get the executable's parent directory: %v", err)
	}
	dataDirPath, err := common.GetDirPath(parentDir, dataDirName)
	if err != nil {
		lg.Fatalf("Failed to get the data directory: %v", err)
	}

	viper.AddConfigPath(dataDirPath)
	viper.SetConfigType(configFileType)
	viper.SetConfigName(configFileName)

	viper.SetDefault(ConfigKeyPort, common.ServerDefaultPort)
	viper.SetDefault(ConfigKeyDebugEnabled, false)
	viper.SetDefault(ConfigKeyDefaultDatabase, DefaultDatabase)
	viper.SetDefault(ConfigKeyLogFileEnabled, false)
	viper.SetDefault(ConfigKeyTLSEnabled, false)
	viper.SetDefault(ConfigKeyTLSCertPath, "")
	viper.SetDefault(ConfigKeyTLSPrivKeyPath, "")
	viper.SetDefault(ConfigKeyMaxClientConnections, common.DefaultMaxClientConnections)

	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv()
	viper.SafeWriteConfig()
	if err = viper.ReadInConfig(); err != nil {
		lg.Fatalf("Failed to read configuration file: %v", err)
	}

	return ServerConfig{
		LogFileEnabled:       viper.GetBool(ConfigKeyLogFileEnabled),
		TLSEnabled:           viper.GetBool(ConfigKeyTLSEnabled),
		DebugEnabled:         viper.GetBool(ConfigKeyDebugEnabled),
		DefaultDB:            viper.GetString(ConfigKeyDefaultDatabase),
		LogFilePath:          filepath.Join(dataDirPath, logFileName),
		MaxKeysPerDB:         common.DbMaxKeyCount,
		MaxHashMapFields:     common.HashMapMaxFields,
		PortInUse:            viper.GetUint16(ConfigKeyPort),
		MaxClientConnections: viper.GetUint32(ConfigKeyMaxClientConnections),
		TLSCertPath:          viper.GetString(ConfigKeyTLSCertPath),
		TLSPrivKeyPath:       viper.GetString(ConfigKeyTLSPrivKeyPath),
	}
}

// ShouldUsePassword returns the server password if it is set with an environment variable.
// The returned bool is true if it is set and false if not.
func ShouldUsePassword() (string, bool) {
	return getEnvVar(EnvVarPassword)
}

func getEnvVar(envVar string) (string, bool) {
	return os.LookupEnv(envVar)
}
