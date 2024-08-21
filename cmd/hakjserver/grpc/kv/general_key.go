package kv

import (
	"context"

	"github.com/hollowdll/hakjdb/api/v1/kvpb"
	grpcerrors "github.com/hollowdll/hakjdb/cmd/hakjserver/grpc/errors"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/server"
)

type GeneralKVServiceServer struct {
	srv server.GeneralKVService
	kvpb.UnimplementedGeneralKVServiceServer
}

func NewGeneralKVServiceServer(s *server.HakjServer) kvpb.GeneralKVServiceServer {
	return &GeneralKVServiceServer{srv: s}
}

// GetAllKeys is the implementation of RPC GetAllKeys.
func (s *GeneralKVServiceServer) GetAllKeys(ctx context.Context, req *kvpb.GetAllKeysRequest) (res *kvpb.GetAllKeysResponse, err error) {
	res, err = s.srv.GetAllKeys(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetKeyType is the implementation of RPC GetKeyType.
func (s *GeneralKVServiceServer) GetKeyType(ctx context.Context, req *kvpb.GetKeyTypeRequest) (res *kvpb.GetKeyTypeResponse, err error) {
	res, err = s.srv.GetKeyType(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// DeleteKeys is the implementation of RPC DeleteKeys.
func (s *GeneralKVServiceServer) DeleteKeys(ctx context.Context, req *kvpb.DeleteKeysRequest) (res *kvpb.DeleteKeysResponse, err error) {
	res, err = s.srv.DeleteKeys(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// DeleteAllKeys is the implementation of RPC DeleteAllKeys.
func (s *GeneralKVServiceServer) DeleteAllKeys(ctx context.Context, req *kvpb.DeleteAllKeysRequest) (res *kvpb.DeleteAllKeysResponse, err error) {
	res, err = s.srv.DeleteAllKeys(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}
