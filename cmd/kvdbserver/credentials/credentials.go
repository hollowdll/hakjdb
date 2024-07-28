package credentials

import "golang.org/x/crypto/bcrypt"

// CredentialStore is an interface for managing server credentials.
type CredentialStore interface {
	SetServerPassword(password []byte) error
	IsCorrectServerPassword(password []byte) error
	IsServerPasswordEnabled() bool
}

// InMemoryCredentialStore is an implementation of interface CredentialStore.
// It stores server credentials in memory.
type InMemoryCredentialStore struct {
	serverPasswordEnabled bool
	serverPasswordHash    []byte
}

func NewInMemoryCredentialStore() *InMemoryCredentialStore {
	return &InMemoryCredentialStore{
		serverPasswordEnabled: false,
		serverPasswordHash:    nil,
	}
}

func (cs *InMemoryCredentialStore) IsServerPasswordEnabled() bool {
	return cs.serverPasswordEnabled
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
	cs.serverPasswordEnabled = true

	return nil
}

// IsCorrectServerPassword checks if provided password matches the server password.
// Returns nil if matches, otherwise an error is returned.
func (cs *InMemoryCredentialStore) IsCorrectServerPassword(password []byte) error {
	return bcrypt.CompareHashAndPassword(cs.serverPasswordHash, password)
}
