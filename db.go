package kvdb

import (
	"sync"
)

// Database containing key-value pairs of data.
type Database struct {
	Name  string
	data  map[DatabaseKey]DatabaseValue
	mutex sync.RWMutex
}

// Creates a new instance of Database.
func newDatabase(name string) *Database {
	return &Database{
		Name: name,
		data: make(map[DatabaseKey]DatabaseValue),
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

// Retrieves a value using a key.
func (db *Database) Get(key DatabaseKey) DatabaseValue {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	return db.data[key]
}

// Sets a value using a key.
func (db *Database) Set(key DatabaseKey, value DatabaseValue) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.data[key] = value
}
