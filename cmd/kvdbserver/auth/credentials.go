package auth

import (
	"github.com/hollowdll/kvdb/errors"
	"golang.org/x/crypto/bcrypt"
)

const RootUserName string = "root"

// CredentialStore is an interface for managing server credentials.
type CredentialStore interface {
	SetServerPassword(password []byte) error
	IsCorrectServerPassword(password []byte) error
	IsServerPasswordEnabled() bool

	// SetServerPassword sets a new password for user.
	// If the user doesn't exist, it is added along with the password.
	// The password is hashed using bcrypt before storing it in memory.
	// Max password size is 72 bytes.
	SetPassword(user string, password []byte) error
	// IsCorrectServerPassword checks if the provided password matches the user's stored password.
	// Returns nil if matches, otherwise an error is returned.
	IsCorrectPassword(user string, password []byte) error
}

// InMemoryCredentialStore is an implementation of interface CredentialStore.
// It stores server credentials in memory.
type InMemoryCredentialStore struct {
	// DEPRECATED
	serverPasswordEnabled bool
	// DEPRECATED
	serverPasswordHash []byte
	userPasswords      map[string][]byte
}

func NewInMemoryCredentialStore() *InMemoryCredentialStore {
	return &InMemoryCredentialStore{
		serverPasswordEnabled: false,
		serverPasswordHash:    nil,
		userPasswords: map[string][]byte{
			RootUserName: []byte(""),
		},
	}
}

// DEPRECATED
func (cs *InMemoryCredentialStore) IsServerPasswordEnabled() bool {
	return cs.serverPasswordEnabled
}

// SetServerPassword sets a new password for the server.
// The password is hashed using bcrypt before storing it in memory.
// If password is set, clients must authenticate using it.
// Max password size is 72 bytes.
//
// DEPRECATED
func (cs *InMemoryCredentialStore) SetServerPassword(password []byte) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	cs.serverPasswordHash = hashedPassword
	cs.serverPasswordEnabled = true

	return nil
}

// IsCorrectServerPassword checks if provided password matches the server password.
// Returns nil if matches, otherwise an error is returned.
//
// DEPRECATED.
func (cs *InMemoryCredentialStore) IsCorrectServerPassword(password []byte) error {
	return bcrypt.CompareHashAndPassword(cs.serverPasswordHash, password)
}

func (cs *InMemoryCredentialStore) SetPassword(user string, password []byte) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	cs.userPasswords[user] = hashedPassword
	return nil
}

func (cs *InMemoryCredentialStore) IsCorrectPassword(user string, password []byte) error {
	userPassword, ok := cs.userPasswords[user]
	if !ok {
		return errors.ErrUserNotFound
	}
	return bcrypt.CompareHashAndPassword(userPassword, password)
}
