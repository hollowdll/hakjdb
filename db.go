package kvdb

import (
	"fmt"
	"math"
	"sync"
	"time"
)

const dbMaxKeyCount uint32 = math.MaxUint32 / 2

// DatabaseKey represents key-value pair key. Key is stored as string.
type DatabaseKey string

// DatabaseStringValue represents key-value pair string value. Value is stored as string.
type DatabaseStringValue string

// Database containing key-value pairs of data.
type Database struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	data      map[DatabaseKey]DatabaseStringValue
	mutex     sync.RWMutex
}

// Creates a new instance of Database.
func newDatabase(name string) *Database {
	return &Database{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		data:      make(map[DatabaseKey]DatabaseStringValue),
	}
}

// update updates the database changing some of its fields.
func (db *Database) update() {
	db.UpdatedAt = time.Now()
}

// GetKeyCount returns the number of keys in the database.
func (db *Database) GetKeyCount() uint32 {
	return uint32(len(db.data))
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

	// Max key count exceeded
	if db.GetKeyCount() >= dbMaxKeyCount {
		return fmt.Errorf("max key count exceeded (%d keys)", dbMaxKeyCount)
	}

	db.data[key] = value
	db.update()

	return nil
}
