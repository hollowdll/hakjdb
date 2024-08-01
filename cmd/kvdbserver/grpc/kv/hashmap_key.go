package storage

import (
	"context"

	"github.com/hollowdll/kvdb/api/v0/kvpb"
	grpcerrors "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc/errors"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/validation"
)

const (
	setHashMapRPCName                   string = "SetHashMap"
	getHashMapFieldValuesRPCName        string = "GetHashMapFieldValues"
	getAllHashMapFieldsAndValuesRPCName string = "GetAllHashMapFieldsAndValues"
	deleteHashMapFieldsRPCName          string = "DeleteHashMapFields"
)

type HashMapKVServiceServer struct {
	srv server.HashMapKVService
	kvpb.UnimplementedHashMapKVServiceServer
}

func NewHashMapKVServiceServer(s *server.KvdbServer) kvpb.HashMapKVServiceServer {
	return &HashMapKVServiceServer{srv: s}
}

// SetHashMap is the implementation of RPC SetHashMap.
func (s *HashMapKVServiceServer) SetHashMap(ctx context.Context, req *kvpb.SetHashMapRequest) (res *kvpb.SetHashMapResponse, err error) {
	logger := s.srv.Logger()
	dbName := s.srv.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", setHashMapRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", setHashMapRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", setHashMapRPCName, dbName, req)
		}
	}()

	if err = validation.ValidateDBKey(req.Key); err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	res, err = s.srv.SetHashMap(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetHashMapFieldValues is the implementation of RPC GetHashMapFieldValues.
func (s *HashMapKVServiceServer) GetHashMapFieldValues(ctx context.Context, req *kvpb.GetHashMapFieldValuesRequest) (res *kvpb.GetHashMapFieldValuesResponse, err error) {
	logger := s.srv.Logger()
	dbName := s.srv.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", getHashMapFieldValuesRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", getHashMapFieldValuesRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", getHashMapFieldValuesRPCName, dbName, req)
		}
	}()

	res, err = s.srv.GetHashMapFieldValues(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetAllHashMapFieldsAndValues is the implementation of RPC GetAllHashMapFieldsAndValues.
func (s *HashMapKVServiceServer) GetAllHashMapFieldsAndValues(ctx context.Context, req *kvpb.GetAllHashMapFieldsAndValuesRequest) (res *kvpb.GetAllHashMapFieldsAndValuesResponse, err error) {
	logger := s.srv.Logger()
	dbName := s.srv.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", getAllHashMapFieldsAndValuesRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", getAllHashMapFieldsAndValuesRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", getAllHashMapFieldsAndValuesRPCName, dbName, req)
		}
	}()

	res, err = s.srv.GetAllHashMapFieldsAndValues(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// DeleteHashMapFields is the implementation of RPC DeleteHashMapFields.
func (s *HashMapKVServiceServer) DeleteHashMapFields(ctx context.Context, req *kvpb.DeleteHashMapFieldsRequest) (res *kvpb.DeleteHashMapFieldsResponse, err error) {
	logger := s.srv.Logger()
	dbName := s.srv.GetDBNameFromContext(ctx)
	logger.Debugf("%s: (call) db = %s %v", deleteHashMapFieldsRPCName, dbName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", deleteHashMapFieldsRPCName, err)
		} else {
			logger.Debugf("%s: (success) db = %s %v", deleteHashMapFieldsRPCName, dbName, req)
		}
	}()

	res, err = s.srv.DeleteHashMapFields(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}
