package grpc

import (
	"github.com/hollowdll/hakjdb/api/v1/authpb"
	"github.com/hollowdll/hakjdb/api/v1/dbpb"
	"github.com/hollowdll/hakjdb/api/v1/echopb"
	"github.com/hollowdll/hakjdb/api/v1/kvpb"
	"github.com/hollowdll/hakjdb/api/v1/serverpb"
	authrpc "github.com/hollowdll/hakjdb/cmd/hakjserver/grpc/auth"
	dbrpc "github.com/hollowdll/hakjdb/cmd/hakjserver/grpc/db"
	echorpc "github.com/hollowdll/hakjdb/cmd/hakjserver/grpc/echo"
	kvrpc "github.com/hollowdll/hakjdb/cmd/hakjserver/grpc/kv"
	serverrpc "github.com/hollowdll/hakjdb/cmd/hakjserver/grpc/server"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetupGrpcServer(s *server.HakjServer) *grpc.Server {
	logger := s.Logger()
	logger.Infof("Setting up gRPC server ...")
	var opts []grpc.ServerOption
	chainUnaryInterceptors := []grpc.UnaryServerInterceptor{
		newLogUnaryInterceptor(s),
		newAuthUnaryInterceptor(s),
		newHeaderUnaryInterceptor(s),
	}
	opts = append(opts, grpc.ChainUnaryInterceptor(chainUnaryInterceptors...))

	if !s.Cfg.TLSEnabled {
		logger.Warning("TLS is disabled. Connections will not be encrypted")
	} else {
		logger.Info("Attempting to enable TLS ...")
		opts = append(opts, grpc.Creds(s.GetTLSCredentials()))
		logger.Info("TLS is enabled. Connections will be encrypted")
	}

	grpcServer := grpc.NewServer(opts...)
	echopb.RegisterEchoServiceServer(grpcServer, echorpc.NewEchoServiceServer())
	serverpb.RegisterServerServiceServer(grpcServer, serverrpc.NewServerServiceServer(s))
	dbpb.RegisterDBServiceServer(grpcServer, dbrpc.NewDBServiceServer(s))
	kvpb.RegisterGeneralKVServiceServer(grpcServer, kvrpc.NewGeneralKVServiceServer(s))
	kvpb.RegisterStringKVServiceServer(grpcServer, kvrpc.NewStringKVServiceServer(s))
	kvpb.RegisterHashMapKVServiceServer(grpcServer, kvrpc.NewHashMapKVServiceServer(s))
	authpb.RegisterAuthServiceServer(grpcServer, authrpc.NewAuthServiceServer(s))

	// enable gRPC server reflection in debug mode
	if s.Cfg.DebugEnabled {
		logger.Info("Debug mode detected: enabling gRPC server reflection ...")
		reflection.Register(grpcServer)
		logger.Info("gRPC server reflection enabled")
	}

	return grpcServer
}

func ServeGrpcServer(s *server.HakjServer, grpcServer *grpc.Server) {
	logger := s.Logger()
	if err := grpcServer.Serve(s.ClientConnListener); err != nil {
		logger.Errorf("Failed to accept incoming connection: %v", err)
	}
}
