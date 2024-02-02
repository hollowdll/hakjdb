package server_test

import (
	"strings"
	"testing"

	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/stretchr/testify/assert"
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
