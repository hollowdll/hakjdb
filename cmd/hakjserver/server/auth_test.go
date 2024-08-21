package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/hollowdll/hakjdb"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/auth"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/config"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/server"
	"github.com/hollowdll/hakjdb/internal/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestAuthorizeIncomingRpcCall(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.AuthEnabled = true
	cfg.AuthTokenSecretKey = "test-key"
	cfg.AuthTokenTTL = 60

	t.Run("AuthDisabled", func(t *testing.T) {
		cfg := config.DefaultConfig()
		s := server.NewKvdbServer(cfg, hakjdb.DisabledLogger())
		ctx := context.Background()
		err := s.AuthorizeIncomingRpcCall(ctx)
		assert.NoErrorf(t, err, "expected no error; error = %v", err)
	})

	t.Run("ValidToken", func(t *testing.T) {
		s := server.NewKvdbServer(cfg, hakjdb.DisabledLogger())
		password := "pass1234"
		user := "root"
		s.EnableAuth(password)
		jwtOpts := &auth.JWTOptions{
			SignKey: cfg.AuthTokenSecretKey,
			TTL:     time.Duration(cfg.AuthTokenTTL) * time.Second,
		}
		token, err := auth.GenerateJWT(jwtOpts, user)
		require.NoErrorf(t, err, "expected no error; error = %v", err)

		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyAuthToken, token))
		err = s.AuthorizeIncomingRpcCall(ctx)
		assert.NoErrorf(t, err, "expected no error; error = %v", err)
	})

	t.Run("MetadataNotSent", func(t *testing.T) {
		s := server.NewKvdbServer(cfg, hakjdb.DisabledLogger())
		password := "pass1234"
		s.EnableAuth(password)

		ctx := context.Background()
		err := s.AuthorizeIncomingRpcCall(ctx)
		require.Error(t, err, "expected error")

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.Unauthenticated, st.Code(), "expected status = %s; got = %s", codes.Unauthenticated, st.Code())
	})

	t.Run("TokenNotSent", func(t *testing.T) {
		s := server.NewKvdbServer(cfg, hakjdb.DisabledLogger())
		password := "pass1234"
		s.EnableAuth(password)

		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("incorrect-123", "incorrect"))
		err := s.AuthorizeIncomingRpcCall(ctx)
		require.Error(t, err, "expected error")

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.Unauthenticated, st.Code(), "expected status = %s; got = %s", codes.Unauthenticated, st.Code())
	})

	t.Run("InvalidToken", func(t *testing.T) {
		s := server.NewKvdbServer(cfg, hakjdb.DisabledLogger())
		password := "pass321!"
		s.EnableAuth(password)

		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyAuthToken, "invalid-token"))
		err := s.AuthorizeIncomingRpcCall(ctx)
		require.Error(t, err, "expected error")

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.Unauthenticated, st.Code(), "expected status = %s; got = %s", codes.Unauthenticated, st.Code())
	})
}
