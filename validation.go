package kvdb

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	// dbNameMaxSize is the maximum length of database name in bytes.
	dbNameMaxSize int = 32
	// dbKeyMaxSize is the maximum length of database key in bytes.
	dbKeyMaxSize int = 32
)

// isEmpty returns true if input is too long.
func isEmpty(input string) bool {
	return len(input) == 0
}

// isTooLong returns true if input is longer than target.
// Target should be in bytes.
func isTooLong(input string, targetBytes int) bool {
	return len(input) > targetBytes
}

// dbNamecontainsValidCharacters checks if database name
// contains valid characters by matching it against a regexp.
func dbNamecontainsValidCharacters(name string) bool {
	pattern := regexp.MustCompile("^[A-Za-z0-9-_]+$")
	return pattern.MatchString(name)
}

// dbKeycontainsValidCharacters checks if database key
// contains valid characters by matching it against a regexp.
func dbKeycontainsValidCharacters(key DatabaseKey) bool {
	pattern := regexp.MustCompile("^[A-Za-z0-9-_:]+$")
	return pattern.MatchString(string(key))
}

// validateDatabaseName validates database name.
// Returns error if validation error is matched.
func validateDatabaseName(name string) error {
	if isEmpty(name) {
		return errors.New("database name is empty")
	}
	if isTooLong(name, dbNameMaxSize) {
		return fmt.Errorf("database name is too long (max %d bytes)", dbNameMaxSize)
	}
	if !dbNamecontainsValidCharacters(name) {
		return errors.New("database name contains invalid characters")
	}

	return nil
}

// validateDatabaseKey validates database key.
// Returns error if validation error is matched.
func validateDatabaseKey(key DatabaseKey) error {
	if isEmpty(string(key)) {
		return errors.New("key is empty")
	}
	if isTooLong(string(key), dbKeyMaxSize) {
		return fmt.Errorf("key is too long (max %d bytes)", dbKeyMaxSize)
	}
	if !dbKeycontainsValidCharacters(key) {
		return errors.New("key contains invalid characters")
	}

	return nil
}
