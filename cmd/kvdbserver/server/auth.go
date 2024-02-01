package server

import "golang.org/x/crypto/bcrypt"

// InMemoryCredentialStore stores server credentials like passwords in memory.
type InMemoryCredentialStore struct {
	serverPasswordHash []byte
}

func newInMemoryCredentialStore() *InMemoryCredentialStore {
	return &InMemoryCredentialStore{
		serverPasswordHash: nil,
	}
}

// SetServerPassword sets a new password for the server.
// If password is set, clients must authenticate using it.
func (cs *InMemoryCredentialStore) SetServerPassword(password []byte) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	cs.serverPasswordHash = hashedPassword

	return nil
}

// IsCorrectServerPassword checks if provided password matches the server password.
// Returns true if matches, otherwise false.
func (cs *InMemoryCredentialStore) IsCorrectServerPassword(password []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(cs.serverPasswordHash, password)
	if err != nil {
		return false, err
	}

	return true, nil
}
