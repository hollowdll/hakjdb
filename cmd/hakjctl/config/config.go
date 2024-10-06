package config

import (
	"time"

	"github.com/hollowdll/hakjdb/internal/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configFileName     string = "hakjctl-config"
	configFileType     string = "yaml"
	CacheDirSubDirName string = ".hakjctl"
	TokenCacheFileName string = "._hakjctl_token_cache___"

	// ConfigKeyHost is the configuration key for host.
	ConfigKeyHost string = "host"
	// ConfigKeyPort is the configuration key for port.
	ConfigKeyPort string = "port"
	// ConfigKeyDatabase is the configuration key for default database.
	ConfigKeyDatabase string = "default_db"
	// ConfigKeyTlsCertPath is the configuration key for TLS client certificate path.
	ConfigKeyTLSClientCertPath string = "tls_client_cert_path"
	// ConfigKeyTLSClientKeyPath is the configuration key for TLS client key path.
	ConfigKeyTLSClientKeyPath string = "tls_client_key_path"
	// ConfigKeyTLSCACertPath is the configuration key for TLS CA certificate path.
	ConfigKeyTLSCACertPath string = "tls_ca_cert_path"
	// ConfigKeyCommandTimeout is the configuration key for setting command timeout.
	ConfigKeyCommandTimeout string = "command_timeout"

	// EnvPrefix is the prefix for environment variables.
	EnvPrefix string = "HAKJCTL"
	// EnvVarPassword is the environment variable for password.
	EnvVarPassword string = EnvPrefix + "_PASSWORD"

	// DefaultDatabase is the name of the default database to use.
	DefaultDatabase string = "default"
	// DefaultCommandTimeout is the default command timeout in seconds.
	DefaultCommandTimeout uint32 = 10
)

var cfgRegistry *viper.Viper

func InitCfgRegistry() {
	if cfgRegistry == nil {
		cfgRegistry = viper.New()
	}
}

// InitConfig initializes and loads configurations.
func InitConfig() {
	configDirPath, err := common.GetExecParentDirPath()
	cobra.CheckErr(err)
	initConfigFile(configDirPath)
	setConfigDefaults()
	cfgRegistry.SafeWriteConfig()
	cobra.CheckErr(cfgRegistry.ReadInConfig())
}

func initConfigFile(configDirPath string) {
	cfgRegistry.AddConfigPath(configDirPath)
	cfgRegistry.SetConfigType(configFileType)
	cfgRegistry.SetConfigName(configFileName)
}

func setConfigDefaults() {
	cfgRegistry.SetDefault(ConfigKeyHost, common.ServerDefaultHost)
	cfgRegistry.SetDefault(ConfigKeyPort, common.ServerDefaultPort)
	cfgRegistry.SetDefault(ConfigKeyDatabase, DefaultDatabase)
	cfgRegistry.SetDefault(ConfigKeyTLSClientCertPath, "")
	cfgRegistry.SetDefault(ConfigKeyTLSClientKeyPath, "")
	cfgRegistry.SetDefault(ConfigKeyTLSCACertPath, "")
	cfgRegistry.SetDefault(ConfigKeyCommandTimeout, DefaultCommandTimeout)
}

// GetCmdTimeout gets the configured command timeout.
// Command timeout is the maximum number of seconds to wait before a request is cancelled.
func GetCmdTimeout() time.Duration {
	return time.Duration(cfgRegistry.GetUint32(ConfigKeyCommandTimeout)) * time.Second
}

// The returned string is the file path. The returned bool is true if the path is set.
func LookupTLSCACert() (string, bool) {
	path := cfgRegistry.GetString(ConfigKeyTLSCACertPath)
	return path, path != ""
}

// The returned string is the file path. The returned bool is true if the path is set.
func LookupTLSClientCert() (string, bool) {
	path := cfgRegistry.GetString(ConfigKeyTLSClientCertPath)
	return path, path != ""
}

// The returned string is the file path. The returned bool is true if the path is set.
func LookupTLSClientKey() (string, bool) {
	path := cfgRegistry.GetString(ConfigKeyTLSClientKeyPath)
	return path, path != ""
}

func WriteFile() error {
	return cfgRegistry.WriteConfig()
}

func GetDefaultDB() string {
	return cfgRegistry.GetString(ConfigKeyDatabase)
}

func GetPort() uint16 {
	return cfgRegistry.GetUint16(ConfigKeyPort)
}

func GetHost() string {
	return cfgRegistry.GetString(ConfigKeyHost)
}

func SetDefaultDB(name string) {
	cfgRegistry.Set(ConfigKeyDatabase, name)
}

func SetPort(port uint16) {
	cfgRegistry.Set(ConfigKeyPort, port)
}

func SetHost(host string) {
	cfgRegistry.Set(ConfigKeyHost, host)
}
