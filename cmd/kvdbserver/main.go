package main

import (
	"github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/config"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/grpc"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/version"
)

func start() {
	logger := kvdb.NewDefaultLogger()
	defer logger.CloseLogFile()
	logger.Infof("Starting kvdb v%s server ...", version.Version)
	cfg := config.LoadConfig(logger)
	s := server.NewKvdbServer(cfg, logger)
	s.Init()
	grpcServer := grpc.SetupGrpcServer(s)
	s.SetupListener()
	grpc.ServeGrpcServer(s, grpcServer)
}

func main() {
	start()
}
