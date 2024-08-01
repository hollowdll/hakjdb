package storage

import (
	"context"

	"github.com/hollowdll/kvdb/api/v0/kvpb"
	grpcerrors "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc/errors"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/validation"
)

const (
	setStringRPCName string = "SetString"
	getStringRPCName string = "GetString"
)

type StringKVServiceServer struct {
	srv server.StringKVService
	kvpb.UnimplementedStringKVServiceServer
}

func NewStringKVServiceServer(s *server.KvdbServer) kvpb.StringKVServiceServer {
	return &StringKVServiceServer{srv: s}
}

// SetString is the implementation of RPC SetString.
func (s *StringKVServiceServer) SetString(ctx context.Context, req *kvpb.SetStringRequest) (res *kvpb.SetStringResponse, err error) {
	logger := s.srv.Logger()
	dbName := s.srv.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", setStringRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", setStringRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", setStringRPCName, dbName, req)
		}
	}()

	if err = validation.ValidateDBKey(req.Key); err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	res, err = s.srv.SetString(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetString is the implementation of RPC GetString.
func (s *StringKVServiceServer) GetString(ctx context.Context, req *kvpb.GetStringRequest) (res *kvpb.GetStringResponse, err error) {
	logger := s.srv.Logger()
	dbName := s.srv.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", getStringRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", getStringRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", getStringRPCName, dbName, req)
		}
	}()

	res, err = s.srv.GetString(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}
