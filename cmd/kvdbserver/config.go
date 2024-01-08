package main

import (
	"github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/viper"
)

const (
	configFileName string = ".kvdbserver"
	configFileType string = "json"

	configKeyPort      string = "port"
	configKeyDebugMode string = "debug_mode"
)

// initConfig initializes server configurations.
func initConfig(s *server) {
	configDirPath, err := kvdb.GetDataDirPath()
	if err != nil {
		s.logger.Fatalf("Failed to get data directory: %s", err)
	}

	viper.AddConfigPath(configDirPath)
	viper.SetConfigType(configFileType)
	viper.SetConfigName(configFileName)

	viper.SetDefault(configKeyPort, common.ServerDefaultPort)
	viper.SetDefault(configKeyDebugMode, false)

	viper.SetEnvPrefix("kvdb")
	viper.AutomaticEnv()

	viper.SafeWriteConfig()
	if err = viper.ReadInConfig(); err != nil {
		s.logger.Fatalf("Failed to load configuration: %s", err)
	}
}
