package server

import (
	"context"

	"github.com/hollowdll/hakjdb/api/v1/serverpb"
	grpcerrors "github.com/hollowdll/hakjdb/cmd/hakjserver/grpc/errors"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/server"
)

type ServerServiceServer struct {
	srv server.ServerService
	serverpb.UnimplementedServerServiceServer
}

func NewServerServiceServer(s *server.HakjServer) serverpb.ServerServiceServer {
	return &ServerServiceServer{srv: s}
}

// GetServerInfo is the implementation of RPC GetServerInfo.
func (s *ServerServiceServer) GetServerInfo(ctx context.Context, req *serverpb.GetServerInfoRequest) (*serverpb.GetServerInfoResponse, error) {
	res, err := s.srv.GetServerInfo(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}
	return res, nil
}

// GetLogs is the implementation of RPC GetLogs.
func (s *ServerServiceServer) GetLogs(ctx context.Context, req *serverpb.GetLogsRequest) (*serverpb.GetLogsResponse, error) {
	res, err := s.srv.GetLogs(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}
	return res, nil
}

// ReloadConfig is the implementation of RPC ReloadConfig.
func (s *ServerServiceServer) ReloadConfig(ctx context.Context, req *serverpb.ReloadConfigRequest) (*serverpb.ReloadConfigResponse, error) {
	res, err := s.srv.ReloadConfig(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}
	return res, nil
}
