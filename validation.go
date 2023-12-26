package kvdb

import (
	"errors"
	"strings"
)

func isEmpty(input string) bool {
	return len(input) == 0
}

func containsWhitespace(input string) bool {
	return strings.ContainsAny(input, " .\t\n\r\v\f")
}

func containsInvalidSymbol(input string) bool {
	return strings.ContainsAny(input, ".")
}

func validateDatabaseName(name string) error {
	if isEmpty(name) {
		return errors.New("Database name is empty")
	}
	if containsWhitespace(name) {
		return errors.New("Database name contains whitespace")
	}
	if containsInvalidSymbol(name) {
		return errors.New("Database name contains invalid symbol")
	}

	return nil
}
