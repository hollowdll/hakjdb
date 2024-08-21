package kv

import (
	"context"

	"github.com/hollowdll/hakjdb/api/v1/kvpb"
	grpcerrors "github.com/hollowdll/hakjdb/cmd/hakjserver/grpc/errors"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/server"
)

type StringKVServiceServer struct {
	srv server.StringKVService
	kvpb.UnimplementedStringKVServiceServer
}

func NewStringKVServiceServer(s *server.HakjServer) kvpb.StringKVServiceServer {
	return &StringKVServiceServer{srv: s}
}

// SetString is the implementation of RPC SetString.
func (s *StringKVServiceServer) SetString(ctx context.Context, req *kvpb.SetStringRequest) (res *kvpb.SetStringResponse, err error) {
	res, err = s.srv.SetString(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetString is the implementation of RPC GetString.
func (s *StringKVServiceServer) GetString(ctx context.Context, req *kvpb.GetStringRequest) (res *kvpb.GetStringResponse, err error) {
	res, err = s.srv.GetString(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}
