package config

import (
	"os"
	"path/filepath"

	"github.com/hollowdll/hakjdb"
	"github.com/hollowdll/hakjdb/internal/common"
	"github.com/spf13/viper"
)

const (
	dataDirName    string = "data"
	configFileName string = "hakjserver-config"
	configFileType string = "yaml"
	logFileName    string = "hakjserver.log"

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
	// ConfigKeyTLSCACertPath is the configuration key for TLS CA certificate file path.
	ConfigKeyTLSCACertPath string = "tls_ca_cert_path"
	// ConfigKeyMaxClientConnections is the configuration key for maximum client connections.
	ConfigKeyMaxClientConnections string = "max_client_connections"
	// ConfigKeyLogLevel is the configuration key for log level.
	ConfigKeyLogLevel string = "log_level"
	// VerboseLogsEnabled is the configuration key for enabling verbose logs.
	ConfigKeyVerboseLogsEnabled string = "verbose_logs_enabled"
	// ConfigKeyAuthEnabled is the configuration key for enabling authentication.
	ConfigKeyAuthEnabled string = "auth_enabled"
	// ConfigKeyAuthTokenSecretKey is the configuration key for setting the secret key used to sign JWT tokens.
	ConfigKeyAuthTokenSecretKey string = "auth_token_secret_key"
	// ConfigKeyAuthTokenTTL is the configuration key for setting the JWT token time to live in seconds.
	ConfigKeyAuthTokenTTL string = "auth_token_ttl"

	// EnvPrefix is the prefix that environment variables use.
	EnvPrefix string = "HAKJ"
	// EnvVarPassword is the environment variable for server password.
	EnvVarPassword string = EnvPrefix + "_PASSWORD"

	DefaultLogFileEnabled       bool   = false
	DefaultTLSEnabled           bool   = false
	DefaultDebugEnabled         bool   = false
	DefaultVerboseLogsEnabled   bool   = false
	DefaultAuthEnabled          bool   = false
	DefaultDatabase             string = "default"
	DefaultPort                 uint16 = common.ServerDefaultPort
	DefaultLogFilePath          string = ""
	DefaultMaxKeysPerDB         uint32 = common.DbMaxKeyCount
	DefaultMaxHashMapFields     uint32 = common.HashMapMaxFields
	DefaultMaxClientConnections uint32 = common.DefaultMaxClientConnections
	DefaultTLSCertPath          string = ""
	DefaultTLSPrivKeyPath       string = ""
	DefaultTLSCACertPath        string = ""
	DefaultLogLevel             string = hakjdb.DefaultLogLevelStr
	DefaultAuthTokenSecretKey   string = ""
	DefaultAuthTokenTTL         uint32 = 900
)

// ServerConfig holds the server's configuration.
type ServerConfig struct {
	LogFileEnabled     bool
	TLSEnabled         bool
	DebugEnabled       bool
	VerboseLogsEnabled bool
	AuthEnabled        bool

	// The name of the default database that is created at server startup.
	DefaultDB string
	// File path to the log file if it is enabled.
	// ONLY SERVER CAN CONFIGURE.
	LogFilePath string
	// The maximum number of keys a database can hold.
	// ONLY SERVER CAN CONFIGURE.
	MaxKeysPerDB uint32
	// The maximum number of fields a HashMap can hold.
	// ONLY SERVER CAN CONFIGURE.
	MaxHashMapFields uint32
	// The TCP/IP port the server listens at.
	PortInUse uint16
	// The maximum number of active client connections allowed.
	MaxClientConnections uint32

	TLSCertPath    string
	TLSPrivKeyPath string
	TLSCACertPath  string

	// Secret key used to sign JWT tokens.
	AuthTokenSecretKey string
	// JWT token time to live in seconds.
	AuthTokenTTL uint32
}

// LoadConfig loads server configurations.
func LoadConfig(lg hakjdb.Logger) ServerConfig {
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

	viper.SetDefault(ConfigKeyPort, DefaultPort)
	viper.SetDefault(ConfigKeyDebugEnabled, DefaultDebugEnabled)
	viper.SetDefault(ConfigKeyDefaultDatabase, DefaultDatabase)
	viper.SetDefault(ConfigKeyLogFileEnabled, DefaultLogFileEnabled)
	viper.SetDefault(ConfigKeyTLSEnabled, DefaultTLSEnabled)
	viper.SetDefault(ConfigKeyVerboseLogsEnabled, DefaultVerboseLogsEnabled)
	viper.SetDefault(ConfigKeyAuthEnabled, DefaultAuthEnabled)
	viper.SetDefault(ConfigKeyTLSCertPath, DefaultTLSCertPath)
	viper.SetDefault(ConfigKeyTLSPrivKeyPath, DefaultTLSPrivKeyPath)
	viper.SetDefault(ConfigKeyTLSCACertPath, DefaultTLSCACertPath)
	viper.SetDefault(ConfigKeyMaxClientConnections, DefaultMaxClientConnections)
	viper.SetDefault(ConfigKeyLogLevel, DefaultLogLevel)
	viper.SetDefault(ConfigKeyAuthTokenSecretKey, DefaultAuthTokenSecretKey)
	viper.SetDefault(ConfigKeyAuthTokenTTL, DefaultAuthTokenTTL)

	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv()
	viper.SafeWriteConfig()
	if err = viper.ReadInConfig(); err != nil {
		lg.Errorf("Failed to read configuration file: %v", err)
	}

	logLevel, logLevelStr, ok := hakjdb.GetLogLevelFromStr(viper.GetString(ConfigKeyLogLevel))
	if !ok {
		lg.Warning("Invalid log level configured. Default log level will be used")
	}
	lg.Infof("Using log level %s", logLevelStr)
	lg.SetLogLevel(logLevel)

	if viper.GetBool(ConfigKeyVerboseLogsEnabled) {
		lg.Info("verbose logs are enabled")
	}

	return ServerConfig{
		LogFileEnabled:       viper.GetBool(ConfigKeyLogFileEnabled),
		TLSEnabled:           viper.GetBool(ConfigKeyTLSEnabled),
		DebugEnabled:         viper.GetBool(ConfigKeyDebugEnabled),
		VerboseLogsEnabled:   viper.GetBool(ConfigKeyVerboseLogsEnabled),
		AuthEnabled:          viper.GetBool(ConfigKeyAuthEnabled),
		DefaultDB:            viper.GetString(ConfigKeyDefaultDatabase),
		LogFilePath:          filepath.Join(dataDirPath, logFileName),
		MaxKeysPerDB:         DefaultMaxKeysPerDB,
		MaxHashMapFields:     DefaultMaxHashMapFields,
		PortInUse:            viper.GetUint16(ConfigKeyPort),
		MaxClientConnections: viper.GetUint32(ConfigKeyMaxClientConnections),
		TLSCertPath:          viper.GetString(ConfigKeyTLSCertPath),
		TLSPrivKeyPath:       viper.GetString(ConfigKeyTLSPrivKeyPath),
		TLSCACertPath:        viper.GetString(ConfigKeyTLSCACertPath),
		AuthTokenSecretKey:   viper.GetString(ConfigKeyAuthTokenSecretKey),
		AuthTokenTTL:         viper.GetUint32(ConfigKeyAuthTokenTTL),
	}
}

// DefaultConfig returns the default configurations.
func DefaultConfig() ServerConfig {
	return ServerConfig{
		LogFileEnabled:       DefaultLogFileEnabled,
		TLSEnabled:           DefaultTLSEnabled,
		DebugEnabled:         DefaultDebugEnabled,
		VerboseLogsEnabled:   DefaultVerboseLogsEnabled,
		AuthEnabled:          DefaultAuthEnabled,
		DefaultDB:            DefaultDatabase,
		LogFilePath:          DefaultLogFilePath,
		MaxKeysPerDB:         DefaultMaxKeysPerDB,
		MaxHashMapFields:     DefaultMaxHashMapFields,
		PortInUse:            DefaultPort,
		MaxClientConnections: DefaultMaxClientConnections,
		TLSCertPath:          DefaultTLSCertPath,
		TLSPrivKeyPath:       DefaultTLSPrivKeyPath,
		TLSCACertPath:        DefaultTLSCACertPath,
		AuthTokenSecretKey:   DefaultAuthTokenSecretKey,
		AuthTokenTTL:         DefaultAuthTokenTTL,
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
