package storage

import (
	"context"

	"github.com/hollowdll/kvdb/api/v0/storagepb"
	rpcerrors "github.com/hollowdll/kvdb/cmd/kvdbserver/rpc/errors"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/validation"
)

const (
	setStringRPCName string = "SetString"
	getStringRPCName string = "GetString"
)

type StringKeyServiceServer struct {
	sks server.StringKeyService
	storagepb.UnimplementedStringKeyServiceServer
}

func NewStringKeyServiceServer(s *server.KvdbServer) storagepb.StringKeyServiceServer {
	return &StringKeyServiceServer{sks: s}
}

// SetString is the implementation of RPC SetString.
func (s *StringKeyServiceServer) SetString(ctx context.Context, req *storagepb.SetStringRequest) (res *storagepb.SetStringResponse, err error) {
	logger := s.sks.Logger()
	dbName := s.sks.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", setStringRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", setStringRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", setStringRPCName, dbName, req)
		}
	}()

	if err = validation.ValidateDBKey(req.Key); err != nil {
		return nil, rpcerrors.ToGrpcError(err)
	}

	res, err = s.sks.SetString(ctx, req)
	if err != nil {
		return nil, rpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetString is the implementation of RPC GetString.
func (s *StringKeyServiceServer) GetString(ctx context.Context, req *storagepb.GetStringRequest) (res *storagepb.GetStringResponse, err error) {
	logger := s.sks.Logger()
	dbName := s.sks.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", getStringRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", getStringRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", getStringRPCName, dbName, req)
		}
	}()

	res, err = s.sks.GetString(ctx, req)
	if err != nil {
		return nil, rpcerrors.ToGrpcError(err)
	}

	return res, nil
}
