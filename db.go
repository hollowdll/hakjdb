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

// Creates a new database with a name. Validates input.
func CreateDatabase(name string) (*Database, error) {
	err := validateDatabaseName(name)
	if err != nil {
		return nil, err
	}

	return newDatabase(name), nil
}

// Retrieves a string value using a key.
func (db *Database) GetString(key DatabaseKey) DatabaseStringValue {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	return db.data[key]
}

// Sets a string value using a key.
func (db *Database) SetString(key DatabaseKey, value DatabaseStringValue) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.data[key] = value
}
