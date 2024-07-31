package grpc

import (
	"github.com/hollowdll/kvdb/api/v0/dbpb"
	"github.com/hollowdll/kvdb/api/v0/serverpb"
	"github.com/hollowdll/kvdb/api/v0/storagepb"
	dbrpc "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc/db"
	serverrpc "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc/server"
	storagerpc "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc/storage"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func SetupGrpcServer(s *server.KvdbServer) *grpc.Server {
	logger := s.Logger()
	logger.Infof("Setting up gRPC server ...")
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(s.AuthInterceptor),
	}

	if !s.Cfg.TLSEnabled {
		logger.Warning("TLS is disabled. Connections will not be encrypted")
	} else {
		logger.Info("Attempting to enable TLS ...")
		cert := s.GetTLSCert()
		opts = append(opts, grpc.Creds(credentials.NewServerTLSFromCert(&cert)))
		logger.Info("TLS is enabled. Connections will be encrypted")
	}

	grpcServer := grpc.NewServer(opts...)
	serverpb.RegisterServerServiceServer(grpcServer, serverrpc.NewServerServiceServer(s))
	dbpb.RegisterDatabaseServiceServer(grpcServer, dbrpc.NewDBServiceServer(s))
	storagepb.RegisterGeneralKeyServiceServer(grpcServer, storagerpc.NewGeneralKeyServiceServer(s))
	storagepb.RegisterStringKeyServiceServer(grpcServer, storagerpc.NewStringKeyServiceServer(s))
	storagepb.RegisterHashMapKeyServiceServer(grpcServer, storagerpc.NewHashMapKeyServiceServer(s))

	return grpcServer
}

func ServeGrpcServer(s *server.KvdbServer, grpcServer *grpc.Server) {
	logger := s.Logger()
	if err := grpcServer.Serve(s.ClientConnListener); err != nil {
		logger.Errorf("Failed to accept incoming connection: %v", err)
	}
}
