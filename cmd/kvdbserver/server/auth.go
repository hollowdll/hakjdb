package server

import (
	"context"
	"time"

	"github.com/hollowdll/kvdb/cmd/kvdbserver/auth"
	kvdberrors "github.com/hollowdll/kvdb/errors"
	"github.com/hollowdll/kvdb/internal/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthorizeIncomingRpcCall checks that incoming RPC call provides valid credentials.
func (s *KvdbServer) AuthorizeIncomingRpcCall(ctx context.Context) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.Cfg.AuthEnabled {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return status.Error(codes.Unauthenticated, kvdberrors.ErrMissingMetadata.Error())
		}

		values := md.Get(common.GrpcMetadataKeyAuthToken)
		if len(values) < 1 {
			return status.Errorf(codes.Unauthenticated, kvdberrors.ErrInvalidAuthToken.Error())
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
			TTL:     time.Duration(s.Cfg.AuthTokenTTL),
		}
		_, err := auth.ValidateJWT(tokenStr, opts)
		if err != nil {
			return status.Error(codes.Unauthenticated, kvdberrors.ErrInvalidAuthToken.Error())
		}
	}
	return nil
}
