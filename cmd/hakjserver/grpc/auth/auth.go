package auth

import (
	"context"

	"github.com/hollowdll/hakjdb/api/v1/authpb"
	grpcerrors "github.com/hollowdll/hakjdb/cmd/hakjserver/grpc/errors"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/server"
)

type AuthServiceServer struct {
	srv server.AuthService
	authpb.UnimplementedAuthServiceServer
}

func NewAuthServiceServer(s *server.HakjServer) authpb.AuthServiceServer {
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
