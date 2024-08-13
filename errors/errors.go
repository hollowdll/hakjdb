package errors

import (
	"errors"
)

var (
	// ErrDatabaseNotFound is returned when a database couldn't be found.
	ErrDatabaseNotFound = errors.New("database not found")
	// ErrDatabaseExists is returned when setting a database name that already exists.
	ErrDatabaseExists = errors.New("database already exists")
	// ErrDatabaseNameRequired is returned when setting a blank database name.
	ErrDatabaseNameRequired = errors.New("database name required")
	// ErrDatabaseNameTooLong is returned when setting a database name that is too long.
	ErrDatabaseNameTooLong = errors.New("database name too long")
	// ErrDatabaseNameInvalid is returned when setting a database name that contains invalid characters.
	ErrDatabaseNameInvalid = errors.New("database name contains invalid characters")
	// ErrDatabaseDescriptionTooLong is returned when setting a database description that is too long.
	ErrDatabaseDescriptionTooLong = errors.New("database description too long")

	// ErrDatabaseKeyRequired is returned when inserting a key with a blank name.
	ErrDatabaseKeyRequired = errors.New("key required")
	// ErrDatabaseKeyTooLong is returned when inserting a key that is too long.
	ErrDatabaseKeyTooLong = errors.New("key too long")
	// ErrDatabaseKeyInvalid is returned when inserting a key with a name that contains
	// invalid characters.
	//
	// DEPRECATED.
	ErrDatabaseKeyInvalid = errors.New("key contains invalid characters")
	// ErrMaxKeysReached is returned when trying to insert keys into a database
	// that has reached the maximum key limit.
	ErrMaxKeysReached = errors.New("max key limit reached")

	// ErrMissingMetadata is returned when gRPC requires metadata
	// but it is missing.
	ErrMissingMetadata = errors.New("missing metadata")
	// ErrMissingKeyInMetadata is returned when a required key is missing
	// in gRPC metadata.
	ErrMissingKeyInMetadata = errors.New("missing key in metadata")

	// ErrInvalidCredentials is returned in authorization process
	// if provided credentials are incorrect.
	ErrInvalidCredentials = errors.New("invalid credentials")

	// ErrLogFileNotEnabled is returned when trying to access logs but log file is not enabled.
	ErrLogFileNotEnabled = errors.New("log file is not enabled")
	// ErrReadLogFile is returned when reading from the log file fails.
	ErrReadLogFile = errors.New("cannot read log file")

	// ErrGetOSInfo is returned when getting information about the OS fails.
	ErrGetOSInfo = errors.New("cannot get information about OS")

	// ErrMaxHashMapFieldsReached is returned when trying to insert fields into a HashMap
	// that has reached the maximum field limit.
	//
	// DEPRECATED.
	ErrMaxHashMapFieldsReached = errors.New("max HashMap field limit reached")

	// ErrMaxClientConnectionsReached is returned when a new client tries to connect to the server
	// but the maximum number of client connections is reached.
	ErrMaxClientConnectionsReached = errors.New("max client connections reached")
)
