package cmd

import (
	"github.com/hollowdll/hakjdb"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/config"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/grpc"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/server"
	"github.com/hollowdll/hakjdb/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "hakjserver",
	Short:   "HakjDB server process",
	Long:    `HakjDB server process that listens for requests from HakjDB clients.`,
	Version: version.Version,
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	parseCmdFlags()
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.DisableAutoGenTag = true
}

func parseCmdFlags() {
	rootCmd.Flags().Uint16("port", config.DefaultPort, "server's TCP/IP port")
	rootCmd.Flags().String("password", config.DefaultPassword, "server password used for authentication")
	rootCmd.Flags().Bool("debug-enabled", config.DefaultDebugEnabled, "enable debug mode")
	rootCmd.Flags().String("default-db", config.DefaultDatabase, "the name of the default database that is created at server startup")
	rootCmd.Flags().Bool("logfile-enabled", config.DefaultLogFileEnabled, "enable the log file")
	rootCmd.Flags().Bool("tls-enabled", config.DefaultTLSEnabled, "enable TLS")
	rootCmd.Flags().Bool("tls-client-cert-auth-enabled", config.DefaultTLSClientCertAuthEnabled, "enable TLS client certification authentication (mTLS)")
	rootCmd.Flags().String("tls-ca-cert-path", config.DefaultTLSCACertPath, "file path to the TLS CA certificate")
	rootCmd.Flags().String("tls-cert-path", config.DefaultTLSCertPath, "file path to the TLS certificate")
	rootCmd.Flags().String("tls-key-path", config.DefaultTLSPrivKeyPath, "file path to the TLS private key")
	rootCmd.Flags().Uint32("max-client-connections", config.DefaultMaxClientConnections, "the maximum number of active client connections allowed")
	rootCmd.Flags().String("log-level", config.DefaultLogLevel, "can be debug, info, warning, error or fatal")
	rootCmd.Flags().Bool("verbose-logs-enabled", config.DefaultVerboseLogsEnabled, "enable verbose logs")
	rootCmd.Flags().Bool("auth-enabled", config.DefaultAuthEnabled, "enable authentication")
	rootCmd.Flags().String("auth-token-secret-key", config.DefaultAuthTokenSecretKey, "secret key used to sign JWT tokens")
	rootCmd.Flags().Uint32("auth-token-ttl", config.DefaultAuthTokenTTL, "JWT token time to live in seconds")

	viper.BindPFlag(config.ConfigKeyPort, rootCmd.Flags().Lookup("port"))
	viper.BindPFlag(config.ConfigKeyPassword, rootCmd.Flags().Lookup("password"))
	viper.BindPFlag(config.ConfigKeyDebugEnabled, rootCmd.Flags().Lookup("debug-enabled"))
	viper.BindPFlag(config.ConfigKeyDefaultDatabase, rootCmd.Flags().Lookup("default-db"))
	viper.BindPFlag(config.ConfigKeyLogFileEnabled, rootCmd.Flags().Lookup("logfile-enabled"))
	viper.BindPFlag(config.ConfigKeyTLSEnabled, rootCmd.Flags().Lookup("tls-enabled"))
	viper.BindPFlag(config.ConfigKeyTLSClientCertAuthEnabled, rootCmd.Flags().Lookup("tls-client-cert-auth-enabled"))
	viper.BindPFlag(config.ConfigKeyTLSCACertPath, rootCmd.Flags().Lookup("tls-ca-cert-path"))
	viper.BindPFlag(config.ConfigKeyTLSCertPath, rootCmd.Flags().Lookup("tls-cert-path"))
	viper.BindPFlag(config.ConfigKeyTLSPrivKeyPath, rootCmd.Flags().Lookup("tls-key-path"))
	viper.BindPFlag(config.ConfigKeyMaxClientConnections, rootCmd.Flags().Lookup("max-client-connections"))
	viper.BindPFlag(config.ConfigKeyLogLevel, rootCmd.Flags().Lookup("log-level"))
	viper.BindPFlag(config.ConfigKeyVerboseLogsEnabled, rootCmd.Flags().Lookup("verbose-logs-enabled"))
	viper.BindPFlag(config.ConfigKeyAuthEnabled, rootCmd.Flags().Lookup("auth-enabled"))
	viper.BindPFlag(config.ConfigKeyAuthTokenSecretKey, rootCmd.Flags().Lookup("auth-token-secret-key"))
	viper.BindPFlag(config.ConfigKeyAuthTokenTTL, rootCmd.Flags().Lookup("auth-token-ttl"))
}

func startServer() {
	logger := hakjdb.NewDefaultLogger()
	defer func() {
		if err := logger.CloseLogFile(); err != nil {
			logger.Errorf("Failed to close log file: %v", err)
		}
	}()
	logger.Infof("Starting HakjDB v%s server ...", version.Version)
	logger.Infof("API version %s", version.APIVersion)
	cfg := config.LoadConfig(logger)
	s := server.NewHakjServer(cfg, logger)
	s.Init()
	grpcServer := grpc.SetupGrpcServer(s)
	s.SetupListener()
	grpc.ServeGrpcServer(s, grpcServer)
}
