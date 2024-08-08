package server_test

import (
	"context"
	"testing"

	"github.com/hollowdll/kvdb/cmd/kvdbserver/config"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestAuthorizeIncomingRpcCall(t *testing.T) {
	cfg := config.DefaultConfig()

	t.Run("PasswordEnabled", func(t *testing.T) {
		s := server.NewKvdbServer(cfg, testutil.DisabledLogger())
		password := "pass321"
		s.EnablePasswordProtection(password)

		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyPassword, password))
		err := s.AuthorizeIncomingRpcCall(ctx)
		assert.NoErrorf(t, err, "expected no error; error = %s", err)
	})

	t.Run("PasswordDisabled", func(t *testing.T) {
		s := server.NewKvdbServer(cfg, testutil.DisabledLogger())
		err := s.AuthorizeIncomingRpcCall(context.Background())
		assert.NoErrorf(t, err, "expected no error; error = %s", err)
	})

	t.Run("InvalidCredentials", func(t *testing.T) {
		s := server.NewKvdbServer(cfg, testutil.DisabledLogger())
		password := "pass321!"
		s.EnablePasswordProtection(password)

		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyPassword, "incorrect"))
		err := s.AuthorizeIncomingRpcCall(ctx)
		require.Error(t, err, "expected error")

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.Unauthenticated, st.Code(), "expected status = %s; got = %s", codes.Unauthenticated, st.Code())
	})
}
