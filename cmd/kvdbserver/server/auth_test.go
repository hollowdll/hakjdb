package server_test

import (
	"context"
	"strings"
	"testing"

	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func TestSetServerPassword(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		credentialStore := server.NewInMemoryCredentialStore()
		password := "pass123"

		err := credentialStore.SetServerPassword([]byte(password))
		assert.NoErrorf(t, err, "expected no error; error = %v", err)
	})

	t.Run("PasswordTooLong", func(t *testing.T) {
		credentialStore := server.NewInMemoryCredentialStore()
		password := strings.Repeat("a", 73)

		err := credentialStore.SetServerPassword([]byte(password))
		assert.Error(t, err, "expected error")
	})
}

func TestIsCorrectServerPassword(t *testing.T) {
	t.Run("CorrectPassword", func(t *testing.T) {
		credentialStore := server.NewInMemoryCredentialStore()
		password := "pass123?"

		err := credentialStore.SetServerPassword([]byte(password))
		require.NoErrorf(t, err, "expected no error; error = %v", err)

		err = credentialStore.IsCorrectServerPassword([]byte(password))
		assert.NoErrorf(t, err, "expected no error; error = %v", err)
	})

	t.Run("IncorrectPassword", func(t *testing.T) {
		credentialStore := server.NewInMemoryCredentialStore()
		password := "pass123!!"

		err := credentialStore.SetServerPassword([]byte(password))
		require.NoErrorf(t, err, "expected no error; error = %v", err)

		err = credentialStore.IsCorrectServerPassword([]byte("pass123?"))
		assert.Error(t, err, "expected error")
	})

	t.Run("PasswordIsNil", func(t *testing.T) {
		credentialStore := server.NewInMemoryCredentialStore()
		password := "pass123"

		err := credentialStore.IsCorrectServerPassword([]byte(password))
		assert.Error(t, err, "expected error")
	})
}

func TestAuthorizeIncomingRpcCall(t *testing.T) {
	t.Run("PasswordEnabled", func(t *testing.T) {
		server := server.NewServer()
		password := "pass321"
		server.DisableLogger()
		server.EnablePassword()
		server.CredentialStore.SetServerPassword([]byte(password))

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
}
