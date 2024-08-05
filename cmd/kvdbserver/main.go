package main

import (
	"github.com/hollowdll/kvdb/cmd/kvdbserver/config"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/grpc"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/logging"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/version"
)

func start() {
	logger := logging.NewDefaultLogger()
	defer func() {
		if err := logger.CloseLogFile(); err != nil {
			logger.Errorf("Failed to close log file: %v", err)
		}
	}()
	logger.Infof("Starting kvdb v%s server ...", version.Version)
	cfg := config.LoadConfig(logger)
	logger.SetLogLevelReady(true)
	s := server.NewKvdbServer(cfg, logger)
	s.Init()
	grpcServer := grpc.SetupGrpcServer(s)
	s.SetupListener()
	grpc.ServeGrpcServer(s, grpcServer)
}

func main() {
	start()
}
