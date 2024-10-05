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
	viper.BindPFlag(config.ConfigKeyPort, rootCmd.Flags().Lookup("port"))
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
