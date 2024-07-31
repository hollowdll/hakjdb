package storage

import (
	"context"

	"github.com/hollowdll/kvdb/api/v0/storagepb"
	rpcerrors "github.com/hollowdll/kvdb/cmd/kvdbserver/rpc/errors"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
)

const (
	setHashMapRPCName                   string = "SetHashMap"
	getHashMapFieldValuesRPCName        string = "GetHashMapFieldValues"
	getAllHashMapFieldsAndValuesRPCName string = "GetAllHashMapFieldsAndValues"
	deleteHashMapFieldsRPCName          string = "DeleteHashMapFields"
)

type HashMapKeyServiceServer struct {
	hks server.HashMapKeyService
	storagepb.UnimplementedHashMapKeyServiceServer
}

func NewHashMapKeyServiceServer(s *server.KvdbServer) storagepb.HashMapKeyServiceServer {
	return &HashMapKeyServiceServer{hks: s}
}

// SetHashMap is the implementation of RPC SetHashMap.
func (s *HashMapKeyServiceServer) SetHashMap(ctx context.Context, req *storagepb.SetHashMapRequest) (res *storagepb.SetHashMapResponse, err error) {
	logger := s.hks.Logger()
	dbName := s.hks.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", setHashMapRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", setHashMapRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", setHashMapRPCName, dbName, req)
		}
	}()

	res, err = s.hks.SetHashMap(ctx, req)
	if err != nil {
		return nil, rpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetHashMapFieldValues is the implementation of RPC GetHashMapFieldValues.
func (s *HashMapKeyServiceServer) GetHashMapFieldValues(ctx context.Context, req *storagepb.GetHashMapFieldValueRequest) (res *storagepb.GetHashMapFieldValueResponse, err error) {
	logger := s.hks.Logger()
	dbName := s.hks.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", getHashMapFieldValuesRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", getHashMapFieldValuesRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", getHashMapFieldValuesRPCName, dbName, req)
		}
	}()

	res, err = s.hks.GetHashMapFieldValues(ctx, req)
	if err != nil {
		return nil, rpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetAllHashMapFieldsAndValues is the implementation of RPC GetAllHashMapFieldsAndValues.
func (s *HashMapKeyServiceServer) GetAllHashMapFieldsAndValues(ctx context.Context, req *storagepb.GetAllHashMapFieldsAndValuesRequest) (res *storagepb.GetAllHashMapFieldsAndValuesResponse, err error) {
	logger := s.hks.Logger()
	dbName := s.hks.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", getAllHashMapFieldsAndValuesRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", getAllHashMapFieldsAndValuesRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", getAllHashMapFieldsAndValuesRPCName, dbName, req)
		}
	}()

	res, err = s.hks.GetAllHashMapFieldsAndValues(ctx, req)
	if err != nil {
		return nil, rpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// DeleteHashMapFields is the implementation of RPC DeleteHashMapFields.
func (s *HashMapKeyServiceServer) DeleteHashMapFields(ctx context.Context, req *storagepb.DeleteHashMapFieldsRequest) (res *storagepb.DeleteHashMapFieldsResponse, err error) {
	logger := s.hks.Logger()
	dbName := s.hks.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", deleteHashMapFieldsRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", deleteHashMapFieldsRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", deleteHashMapFieldsRPCName, dbName, req)
		}
	}()

	res, err = s.hks.DeleteHashMapFields(ctx, req)
	if err != nil {
		return nil, rpcerrors.ToGrpcError(err)
	}

	return res, nil
}
