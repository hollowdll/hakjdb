package server

import (
	"context"

	"github.com/hollowdll/kvdb/api/v0/serverpb"
	grpcerrors "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc/errors"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
)

type ServerServiceServer struct {
	srv server.ServerService
	serverpb.UnimplementedServerServiceServer
}

func NewServerServiceServer(s *server.KvdbServer) serverpb.ServerServiceServer {
	return &ServerServiceServer{srv: s}
}

// GetServerInfo is the implementation of RPC GetServerInfo.
func (s *ServerServiceServer) GetServerInfo(ctx context.Context, req *serverpb.GetServerInfoRequest) (res *serverpb.GetServerInfoResponse, err error) {
	res, err = s.srv.GetServerInfo(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetLogs is the implementation of RPC GetLogs.
func (s *ServerServiceServer) GetLogs(ctx context.Context, req *serverpb.GetLogsRequest) (res *serverpb.GetLogsResponse, err error) {
	res, err = s.srv.GetLogs(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}
