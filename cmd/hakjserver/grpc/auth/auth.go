package auth

import (
	"context"

	"github.com/hollowdll/kvdb/api/v0/authpb"
	grpcerrors "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc/errors"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
)

type AuthServiceServer struct {
	srv server.AuthService
	authpb.UnimplementedAuthServiceServer
}

func NewAuthServiceServer(s *server.KvdbServer) authpb.AuthServiceServer {
	return &AuthServiceServer{srv: s}
}

// Authenticate is the implementation of RPC Authenticate.
func (s *AuthServiceServer) Authenticate(ctx context.Context, req *authpb.AuthenticateRequest) (*authpb.AuthenticateResponse, error) {
	res, err := s.srv.Authenticate(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}
	return res, nil
}
