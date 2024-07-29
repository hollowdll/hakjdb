package server

import (
	"context"

	"github.com/hollowdll/kvdb/api/v0/serverpb"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	kvdberrors "github.com/hollowdll/kvdb/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	getServerInfoRPCName string = "GetServerInfo"
	getLogsRPCName       string = "GetLogs"
)

type ServerServiceServer struct {
	ss server.ServerService
	serverpb.UnimplementedServerServiceServer
}

func NewServerServiceServer(s *server.KvdbServer) serverpb.ServerServiceServer {
	return &ServerServiceServer{ss: s}
}

// GetServerInfo is the implementation of RPC GetServerInfo.
func (s *ServerServiceServer) GetServerInfo(ctx context.Context, req *serverpb.GetServerInfoRequest) (res *serverpb.GetServerInfoResponse, err error) {
	logger := s.ss.Logger()
	logger.Debugf("%s: (call) %v", getServerInfoRPCName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", getServerInfoRPCName, err)
		} else {
			logger.Debugf("%s: (success) %v", getServerInfoRPCName, req)
		}
	}()

	res, err = s.ss.GetServerInfo(ctx, req)
	if err != nil {
		switch err {
		case kvdberrors.ErrGetOSInfo:
			return nil, status.Error(codes.Internal, err.Error())
		default:
			return nil, status.Error(codes.Unknown, err.Error())
		}
	}

	return res, nil
}

// GetLogs is the implementation of RPC GetLogs.
func (s *ServerServiceServer) GetLogs(ctx context.Context, req *serverpb.GetLogsRequest) (res *serverpb.GetLogsResponse, err error) {
	logger := s.ss.Logger()
	logger.Debugf("%s: (call) %v", getLogsRPCName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", getLogsRPCName, err)
		} else {
			logger.Debugf("%s: (success) %v", getLogsRPCName, req)
		}
	}()

	res, err = s.ss.GetLogs(ctx, req)
	if err != nil {
		switch err {
		case kvdberrors.ErrLogFileNotEnabled:
			return nil, status.Error(codes.FailedPrecondition, err.Error())
		case kvdberrors.ErrReadLogFile:
			return nil, status.Error(codes.Internal, err.Error())
		default:
			return nil, status.Error(codes.Unknown, err.Error())
		}
	}

	return res, nil
}
