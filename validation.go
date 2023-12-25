package kvdb

import (
	"errors"
	"strings"
)

func isEmpty(input string) bool {
	return len(input) == 0
}

func containsWhitespace(input string) bool {
	return strings.ContainsAny(input, " \t\n\r\v\f")
}

func validateDatabaseName(name string) error {
	if isEmpty(name) {
		return errors.New("Database name cannot be empty")
	}
	if containsWhitespace(name) {
		return errors.New("Database name cannot contain whitespace")
	}

	return nil
}
