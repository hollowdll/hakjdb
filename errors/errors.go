package errors

import "errors"

var (
	// ErrDatabaseNotFound is returned when a database couldn't be found.
	ErrDatabaseNotFound = errors.New("database not found")

	// ErrDatabaseExists is returned when creating a database that already exists.
	ErrDatabaseExists = errors.New("database already exists")

	// ErrDatabaseNameEmpty is returned when creating a database with a blank name.
	ErrDatabaseNameRequired = errors.New("database name required")

	// ErrDatabaseNameTooLong is returned when creating a database with a name
	// that is larger than DbNameMaxSize.
	ErrDatabaseNameTooLong = errors.New("database name too long")

	// ErrDatabaseNameInvalid is returned when creating a database with a name
	// that contains invalid characters.
	ErrDatabaseNameInvalid = errors.New("database name contains invalid characters")

	// ErrDatabaseKeyRequired is returned when inserting a key with a blank name.
	ErrDatabaseKeyRequired = errors.New("key required")

	// ErrDatabaseKeyTooLong is returned when inserting a key that is larger than DbKeyMaxSize.
	ErrDatabaseKeyTooLong = errors.New("key required")

	// ErrDatabaseKeyInvalid is returned when inserting a key with a name that contains
	// invalid characters.
	ErrDatabaseKeyInvalid = errors.New("key contains invalid characters")

	// ErrMaxKeysExceeded is returned when trying to insert keys into a database
	// that has reached DbMaxKeyCount.
	ErrMaxKeysExceeded = errors.New("max keys exceeded")

	// ErrMissingMetadata is returned when gRPC requires metadata
	// but it is missing.
	ErrMissingMetadata = errors.New("missing metadata")

	// ErrMissingKeyInMetadata is returned when a required key is missing
	// in gRPC metadata.
	ErrMissingKeyInMetadata = errors.New("missing key in metadata")

	// ErrWriteLogFile is returned when a write operation to a log file fails.
	// Combine this with another error if needed.
	ErrWriteLogFile = errors.New("failed to write to log file")
)
