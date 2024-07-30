package storage

import (
	"context"

	"github.com/hollowdll/kvdb/api/v0/storagepb"
	rpcerrors "github.com/hollowdll/kvdb/cmd/kvdbserver/rpc/errors"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
)

const (
	getAllKeysRPCName    string = "GetAllKeys"
	getKeyTypeRPCName    string = "GetKeyType"
	deleteKeysRPCName    string = "DeleteKeys"
	deleteAllKeysRPCName string = "DeleteAllKeys"
)

type GeneralKeyServiceServer struct {
	gks server.GeneralKeyService
	storagepb.UnimplementedGeneralKeyServiceServer
}

func NewGeneralKeyServiceServer(s *server.KvdbServer) storagepb.GeneralKeyServiceServer {
	return &GeneralKeyServiceServer{gks: s}
}

// GetAllKeys is the implementation of RPC GetAllKeys.
func (s *GeneralKeyServiceServer) GetAllKeys(ctx context.Context, req *storagepb.GetAllKeysRequest) (res *storagepb.GetAllKeysResponse, err error) {
	logger := s.gks.Logger()
	dbName := s.gks.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", getAllKeysRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", getAllKeysRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", getAllKeysRPCName, dbName, req)
		}
	}()

	res, err = s.gks.GetAllKeys(ctx, req)
	if err != nil {
		return nil, rpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetKeyType is the implementation of RPC GetKeyType.
func (s *GeneralKeyServiceServer) GetKeyType(ctx context.Context, req *storagepb.GetKeyTypeRequest) (res *storagepb.GetKeyTypeResponse, err error) {
	logger := s.gks.Logger()
	dbName := s.gks.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", getKeyTypeRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", getKeyTypeRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", getKeyTypeRPCName, dbName, req)
		}
	}()

	res, err = s.gks.GetKeyType(ctx, req)
	if err != nil {
		return nil, rpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// DeleteKeys is the implementation of RPC DeleteKeys.
func (s *GeneralKeyServiceServer) DeleteKeys(ctx context.Context, req *storagepb.DeleteKeysRequest) (res *storagepb.DeleteKeysResponse, err error) {
	logger := s.gks.Logger()
	dbName := s.gks.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", deleteKeysRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", deleteKeysRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", deleteKeysRPCName, dbName, req)
		}
	}()

	res, err = s.gks.DeleteKeys(ctx, req)
	if err != nil {
		return nil, rpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// DeleteAllKeys is the implementation of RPC DeleteAllKeys.
func (s *GeneralKeyServiceServer) DeleteAllKeys(ctx context.Context, req *storagepb.DeleteAllKeysRequest) (res *storagepb.DeleteAllKeysResponse, err error) {
	logger := s.gks.Logger()
	dbName := s.gks.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", deleteAllKeysRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", deleteAllKeysRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", deleteAllKeysRPCName, dbName, req)
		}
	}()

	res, err = s.gks.DeleteAllKeys(ctx, req)
	if err != nil {
		return nil, rpcerrors.ToGrpcError(err)
	}

	return res, nil
}
