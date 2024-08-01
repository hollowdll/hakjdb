package server

import (
	"context"

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

	if s.credentialStore.IsServerPasswordEnabled() {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return status.Error(codes.InvalidArgument, kvdberrors.ErrMissingMetadata.Error())
		}

		passwordValues := md.Get(common.GrpcMetadataKeyPassword)
		if len(passwordValues) < 1 {
			return status.Errorf(codes.Unauthenticated, kvdberrors.ErrInvalidCredentials.Error())
		}
		password := passwordValues[0]

		// clear password
		defer func() {
			for i := range passwordValues {
				passwordValues[i] = ""
			}
			password = ""
			md.Set(common.GrpcMetadataKeyPassword, "")
		}()

		err := s.credentialStore.IsCorrectServerPassword([]byte(password))
		if err != nil {
			return status.Error(codes.Unauthenticated, kvdberrors.ErrInvalidCredentials.Error())
		}
	}
	return nil
}
