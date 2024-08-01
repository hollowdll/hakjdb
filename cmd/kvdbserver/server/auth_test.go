package server_test

import (
	"context"
	"testing"

	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestAuthorizeIncomingRpcCall(t *testing.T) {
	t.Run("PasswordEnabled", func(t *testing.T) {
		server := server.NewServer()
		password := "pass321"
		server.DisableLogger()
		server.EnablePasswordProtection(password)

		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyPassword, password))
		err := server.AuthorizeIncomingRpcCall(ctx)
		assert.NoErrorf(t, err, "expected no error; error = %s", err)
	})

	t.Run("PasswordDisabled", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()

		err := server.AuthorizeIncomingRpcCall(context.Background())
		assert.NoErrorf(t, err, "expected no error; error = %s", err)
	})

	t.Run("InvalidCredentials", func(t *testing.T) {
		server := server.NewServer()
		password := "pass321!"
		server.DisableLogger()
		server.EnablePasswordProtection(password)

		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyPassword, "incorrect"))
		err := server.AuthorizeIncomingRpcCall(ctx)
		require.Error(t, err, "expected error")

		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.Unauthenticated, st.Code(), "expected status = %s; got = %s", codes.Unauthenticated, st.Code())
	})
}
