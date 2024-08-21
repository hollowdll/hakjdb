package main

import (
	"github.com/hollowdll/hakjdb"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/config"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/grpc"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/server"
	"github.com/hollowdll/hakjdb/version"
)

func start() {
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

func main() {
	start()
}
