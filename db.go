package kvdb

import (
	"fmt"
	"math"
	"reflect"
	"sync"
	"time"
)

// DbMaxKeyCount is the maximum number of keys a database can hold.
const DbMaxKeyCount uint32 = math.MaxUint32 / 2

// DatabaseKey represents key-value pair key. Key is stored as string.
type DatabaseKey string

// DatabaseStringValue represents key-value pair string value. Value is stored as string.
type DatabaseStringValue string

// Database containing key-value pairs of data.
type Database struct {
	// Name of the database.
	Name string
	// UTC timestamp describing when the database was created.
	CreatedAt time.Time
	// UTC timestamp describing when the database was updated.
	UpdatedAt time.Time
	data      map[DatabaseKey]DatabaseStringValue
	mutex     sync.RWMutex
}

// Creates a new instance of Database.
func newDatabase(name string) *Database {
	return &Database{
		Name:      name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		data:      make(map[DatabaseKey]DatabaseStringValue),
	}
}

// update updates the database changing some of its fields.
func (db *Database) update() {
	db.UpdatedAt = time.Now()
}

// keyExists returns true if key exists in the database.
func (db *Database) keyExists(key DatabaseKey) bool {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	_, exists := db.data[key]
	return exists
}

// GetKeyCount returns the number of keys in the database.
func (db *Database) GetKeyCount() uint32 {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	return uint32(len(db.data))
}

// GetStoredSizeBytes returns the size of stored data in the database in bytes.
func (db *Database) GetStoredSizeBytes() uint64 {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	var size uint64
	for key, value := range db.data {
		size += uint64(reflect.TypeOf(key).Size())
		size += uint64(len(key))
		size += uint64(reflect.TypeOf(value).Size())
		size += uint64(len(value))
	}

	return size
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
	err := validateDatabaseKey(key)
	if err != nil {
		return err
	}

	// Max key count exceeded
	if db.GetKeyCount() >= DbMaxKeyCount {
		return fmt.Errorf("max key count exceeded (%d keys)", DbMaxKeyCount)
	}

	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.data[key] = value
	db.update()

	return nil
}

// DeleteKey deletes a key and its value. Returns true if key exists and it was deleted.
// Returns false if key doesn't exist.
func (db *Database) DeleteKey(key DatabaseKey) bool {
	if !db.keyExists(key) {
		return false
	}

	db.mutex.Lock()
	defer db.mutex.Unlock()
	delete(db.data, key)

	return true
}
