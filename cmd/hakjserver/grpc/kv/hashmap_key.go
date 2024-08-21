package kv

import (
	"context"

	"github.com/hollowdll/kvdb/api/v0/kvpb"
	grpcerrors "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc/errors"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
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
	res, err = s.srv.SetHashMap(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetHashMapFieldValues is the implementation of RPC GetHashMapFieldValues.
func (s *HashMapKVServiceServer) GetHashMapFieldValues(ctx context.Context, req *kvpb.GetHashMapFieldValuesRequest) (res *kvpb.GetHashMapFieldValuesResponse, err error) {
	res, err = s.srv.GetHashMapFieldValues(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetAllHashMapFieldsAndValues is the implementation of RPC GetAllHashMapFieldsAndValues.
func (s *HashMapKVServiceServer) GetAllHashMapFieldsAndValues(ctx context.Context, req *kvpb.GetAllHashMapFieldsAndValuesRequest) (res *kvpb.GetAllHashMapFieldsAndValuesResponse, err error) {
	res, err = s.srv.GetAllHashMapFieldsAndValues(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// DeleteHashMapFields is the implementation of RPC DeleteHashMapFields.
func (s *HashMapKVServiceServer) DeleteHashMapFields(ctx context.Context, req *kvpb.DeleteHashMapFieldsRequest) (res *kvpb.DeleteHashMapFieldsResponse, err error) {
	res, err = s.srv.DeleteHashMapFields(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}
