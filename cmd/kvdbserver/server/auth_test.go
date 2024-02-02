package server_test

import (
	"strings"
	"testing"

	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
}
