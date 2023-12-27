package kvdb

import (
	"errors"
	"regexp"
)

func isEmpty(input string) bool {
	return len(input) == 0
}

func containsValidCharacters(input string) bool {
	pattern := regexp.MustCompile("^[A-Za-z0-9-_]+$")
	return pattern.MatchString(input)
}

func validateDatabaseName(name string) error {
	if isEmpty(name) {
		return errors.New("database name is empty")
	}
	if !containsValidCharacters(name) {
		return errors.New("database name contains invalid characters")
	}

	return nil
}
