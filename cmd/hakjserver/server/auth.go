package server

import (
	"context"
	"time"

	"github.com/hollowdll/hakjdb/cmd/hakjserver/auth"
	hakjerrors "github.com/hollowdll/hakjdb/errors"
	"github.com/hollowdll/hakjdb/internal/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthorizeIncomingRpcCall checks that incoming RPC call provides valid credentials.
func (s *HakjServer) AuthorizeIncomingRpcCall(ctx context.Context) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.Cfg.AuthEnabled {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return status.Error(codes.Unauthenticated, hakjerrors.ErrMissingMetadata.Error())
		}

		values := md.Get(common.GrpcMetadataKeyAuthToken)
		if len(values) < 1 {
			return status.Errorf(codes.Unauthenticated, hakjerrors.ErrInvalidAuthToken.Error())
		}
		tokenStr := values[0]

		// clear token
		defer func() {
			for i := range values {
				values[i] = ""
			}
			tokenStr = ""
			md.Set(common.GrpcMetadataKeyAuthToken, "")
		}()

		opts := &auth.JWTOptions{
			SignKey: s.Cfg.AuthTokenSecretKey,
			TTL:     time.Duration(s.Cfg.AuthTokenTTL) * time.Second,
		}
		_, err := auth.ValidateJWT(tokenStr, opts)
		if err != nil {
			lg := s.Logger()
			lg.Debugf("failed to validate JWT token: %v", err)
			return status.Error(codes.Unauthenticated, hakjerrors.ErrInvalidAuthToken.Error())
		}
	}
	return nil
}
