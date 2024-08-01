package storage

import (
	"context"

	"github.com/hollowdll/kvdb/api/v0/kvpb"
	grpcerrors "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc/errors"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
)

const (
	getAllKeysRPCName    string = "GetAllKeys"
	getKeyTypeRPCName    string = "GetKeyType"
	deleteKeysRPCName    string = "DeleteKeys"
	deleteAllKeysRPCName string = "DeleteAllKeys"
)

type GeneralKVServiceServer struct {
	srv server.GeneralKVService
	kvpb.UnimplementedGeneralKVServiceServer
}

func NewGeneralKVServiceServer(s *server.KvdbServer) kvpb.GeneralKVServiceServer {
	return &GeneralKVServiceServer{srv: s}
}

// GetAllKeys is the implementation of RPC GetAllKeys.
func (s *GeneralKVServiceServer) GetAllKeys(ctx context.Context, req *kvpb.GetAllKeysRequest) (res *kvpb.GetAllKeysResponse, err error) {
	logger := s.srv.Logger()
	dbName := s.srv.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", getAllKeysRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", getAllKeysRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", getAllKeysRPCName, dbName, req)
		}
	}()

	res, err = s.srv.GetAllKeys(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetKeyType is the implementation of RPC GetKeyType.
func (s *GeneralKVServiceServer) GetKeyType(ctx context.Context, req *kvpb.GetKeyTypeRequest) (res *kvpb.GetKeyTypeResponse, err error) {
	logger := s.srv.Logger()
	dbName := s.srv.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", getKeyTypeRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", getKeyTypeRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", getKeyTypeRPCName, dbName, req)
		}
	}()

	res, err = s.srv.GetKeyType(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// DeleteKeys is the implementation of RPC DeleteKeys.
func (s *GeneralKVServiceServer) DeleteKeys(ctx context.Context, req *kvpb.DeleteKeysRequest) (res *kvpb.DeleteKeysResponse, err error) {
	logger := s.srv.Logger()
	dbName := s.srv.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", deleteKeysRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", deleteKeysRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", deleteKeysRPCName, dbName, req)
		}
	}()

	res, err = s.srv.DeleteKeys(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// DeleteAllKeys is the implementation of RPC DeleteAllKeys.
func (s *GeneralKVServiceServer) DeleteAllKeys(ctx context.Context, req *kvpb.DeleteAllKeysRequest) (res *kvpb.DeleteAllKeysResponse, err error) {
	logger := s.srv.Logger()
	dbName := s.srv.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", deleteAllKeysRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", deleteAllKeysRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", deleteAllKeysRPCName, dbName, req)
		}
	}()

	res, err = s.srv.DeleteAllKeys(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}
