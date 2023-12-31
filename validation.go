package kvdb

import (
	"errors"
	"fmt"
	"regexp"
)

// Database name maximum length in bytes.
const dbNameMaxSize int = 32

func isEmpty(input string) bool {
	return len(input) == 0
}

func dbNameTooLong(name string) bool {
	return len(name) > dbNameMaxSize
}

func dbNamecontainsValidCharacters(name string) bool {
	pattern := regexp.MustCompile("^[A-Za-z0-9-_]+$")
	return pattern.MatchString(name)
}

func validateDatabaseName(name string) error {
	if isEmpty(name) {
		return errors.New("database name is empty")
	}
	if dbNameTooLong(name) {
		return fmt.Errorf("database name is too long (max %d bytes)", dbNameMaxSize)
	}
	if !dbNamecontainsValidCharacters(name) {
		return errors.New("database name contains invalid characters")
	}

	return nil
}
