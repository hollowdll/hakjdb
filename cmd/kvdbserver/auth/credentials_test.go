package auth

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetServerPassword(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		credentialStore := NewInMemoryCredentialStore()
		password := "pass123"
		user := "root"

		err := credentialStore.SetPassword(user, []byte(password))
		assert.NoErrorf(t, err, "expected no error; error = %v", err)
	})

	t.Run("PasswordMaxLength", func(t *testing.T) {
		credentialStore := NewInMemoryCredentialStore()
		password := strings.Repeat("a", 72)
		user := "root"

		err := credentialStore.SetPassword(user, []byte(password))
		assert.NoErrorf(t, err, "expected no error; error = %v", err)
	})

	t.Run("PasswordTooLong", func(t *testing.T) {
		credentialStore := NewInMemoryCredentialStore()
		password := strings.Repeat("a", 73)
		user := "root"

		err := credentialStore.SetPassword(user, []byte(password))
		assert.Error(t, err, "expected error")
	})

	t.Run("PasswordEmpty", func(t *testing.T) {
		credentialStore := NewInMemoryCredentialStore()
		password := ""
		user := "root"

		err := credentialStore.SetPassword(user, []byte(password))
		assert.NoErrorf(t, err, "expected no error; error = %v", err)
	})
}

func TestIsCorrectServerPassword(t *testing.T) {
	t.Run("CorrectPassword", func(t *testing.T) {
		credentialStore := NewInMemoryCredentialStore()
		password := "pass123?"
		user := "root"

		err := credentialStore.SetPassword(user, []byte(password))
		require.NoErrorf(t, err, "expected no error; error = %v", err)

		err = credentialStore.IsCorrectPassword(user, []byte(password))
		assert.NoErrorf(t, err, "expected no error; error = %v", err)
	})

	t.Run("IncorrectPassword", func(t *testing.T) {
		credentialStore := NewInMemoryCredentialStore()
		password := "pass123!!"
		user := "root"

		err := credentialStore.SetPassword(user, []byte(password))
		require.NoErrorf(t, err, "expected no error; error = %v", err)

		err = credentialStore.IsCorrectPassword(user, []byte("pass123?"))
		assert.Error(t, err, "expected error")
	})

	t.Run("PasswordIsNotSet", func(t *testing.T) {
		credentialStore := NewInMemoryCredentialStore()
		password := "pass123"
		user := "root"

		err := credentialStore.IsCorrectPassword(user, []byte(password))
		assert.Error(t, err, "expected error")
	})
}
