package kv

import (
	"context"

	"github.com/hollowdll/kvdb/api/v0/kvpb"
	grpcerrors "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc/errors"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
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
