package validation

import (
	"regexp"
	"strings"

	"github.com/hollowdll/hakjdb/errors"
)

const (
	// DbNameMaxSize is the maximum length of a database name in bytes.
	DbNameMaxSize int = 64
	// DbDescriptionMaxSize is the maximum length of a database description in bytes.
	DbDescMaxSize int = 255
	// DbKeyMaxSize is the maximum length of a database key in bytes.
	DbKeyMaxSize int = 1024
)

// isBlank returns true if input is blank.
func isBlank(input string) bool {
	return len(strings.TrimSpace(input)) == 0
}

// isTooLong returns true if input is longer than target.
// Target should be in bytes.
func isTooLong(input string, targetBytes int) bool {
	return len(input) > targetBytes
}

// databaseNameContainsValidCharacters checks if database name
// contains valid characters by matching it against a regexp.
func dbNameContainsValidCharacters(name string) bool {
	pattern := regexp.MustCompile("^[A-Za-z0-9-_]+$")
	return pattern.MatchString(name)
}

// ValidateDBName validates database name.
// Returns error if validation error is matched.
func ValidateDBName(name string) error {
	if isBlank(name) {
		return errors.ErrDatabaseNameRequired
	}
	if isTooLong(name, DbNameMaxSize) {
		return errors.ErrDatabaseNameTooLong
	}
	if !dbNameContainsValidCharacters(name) {
		return errors.ErrDatabaseNameInvalid
	}
	return nil
}

// ValidateDBDesc validates database description.
// Returns error if validation error is matched.
func ValidateDBDesc(desc string) error {
	if isTooLong(desc, DbDescMaxSize) {
		return errors.ErrDatabaseDescriptionTooLong
	}
	return nil
}

// ValidateDBKey validates database key.
// Returns error if validation error is matched.
func ValidateDBKey(key string) error {
	if isBlank(string(key)) {
		return errors.ErrDatabaseKeyRequired
	}
	if isTooLong(string(key), DbKeyMaxSize) {
		return errors.ErrDatabaseKeyTooLong
	}
	return nil
}
