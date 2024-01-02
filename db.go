package kvdb

import (
	"sync"
)

// DatabaseKey represents key-value pair key. Key is stored as string.
type DatabaseKey string

// DatabaseStringValue represents key-value pair string value. Value is stored as string.
type DatabaseStringValue string

// Database containing key-value pairs of data.
type Database struct {
	Name  string
	data  map[DatabaseKey]DatabaseStringValue
	mutex sync.RWMutex
}

// Creates a new instance of Database.
func newDatabase(name string) *Database {
	return &Database{
		Name: name,
		data: make(map[DatabaseKey]DatabaseStringValue),
	}
}

// CreateDatabase creates a new database with a name. Validates input.
func CreateDatabase(name string) (*Database, error) {
	err := validateDatabaseName(name)
	if err != nil {
		return nil, err
	}

	return newDatabase(name), nil
}

// GetString retrieves a string value using a key.
func (db *Database) GetString(key DatabaseKey) DatabaseStringValue {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	return db.data[key]
}

// SetString sets a string value using a key. Validates key before storing.
func (db *Database) SetString(key DatabaseKey, value DatabaseStringValue) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	err := validateDatabaseKey(key)
	if err != nil {
		return err
	}

	db.data[key] = value
	return nil
}
