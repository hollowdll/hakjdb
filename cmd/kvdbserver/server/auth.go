package server

import (
	"context"

	kvdberrors "github.com/hollowdll/kvdb/errors"
	"github.com/hollowdll/kvdb/internal/common"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// InMemoryCredentialStore stores server credentials like passwords in memory.
type InMemoryCredentialStore struct {
	serverPasswordHash []byte
}

func NewInMemoryCredentialStore() *InMemoryCredentialStore {
	return &InMemoryCredentialStore{
		serverPasswordHash: nil,
	}
}

// SetServerPassword sets a new password for the server.
// The password is hashed using bcrypt before storing it in memory.
// If password is set, clients must authenticate using it.
// Max password size is 72 bytes.
func (cs *InMemoryCredentialStore) SetServerPassword(password []byte) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	cs.serverPasswordHash = hashedPassword

	return nil
}

// IsCorrectServerPassword checks if provided password matches the server password.
// Returns nil if matches, otherwise an error is returned.
func (cs *InMemoryCredentialStore) IsCorrectServerPassword(password []byte) error {
	return bcrypt.CompareHashAndPassword(cs.serverPasswordHash, password)
}

// authInterceptor is unary interceptor to handle authorization for RPC calls.
func (s *Server) authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if err := s.AuthorizeIncomingRpcCall(ctx); err != nil {
		s.mutex.RLock()
		s.logger.Errorf("Failed to authorize request: %v", err)
		s.mutex.RUnlock()

		return nil, err
	}

	return handler(ctx, req)
}

// AuthorizeIncomingRpcCall checks that incoming RPC call provides valid credentials.
func (s *Server) AuthorizeIncomingRpcCall(ctx context.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.passwordEnabled {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return status.Error(codes.InvalidArgument, kvdberrors.ErrMissingMetadata.Error())
		}

		passwordValues := md.Get(common.GrpcMetadataKeyPassword)
		if len(passwordValues) < 1 {
			return status.Errorf(codes.Unauthenticated, kvdberrors.ErrInvalidCredentials.Error())
		}
		password := passwordValues[0]

		err := s.CredentialStore.IsCorrectServerPassword([]byte(password))
		if err != nil {
			return status.Error(codes.Unauthenticated, kvdberrors.ErrInvalidCredentials.Error())
		}
	}
	return nil
}
