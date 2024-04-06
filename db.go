package kvdb

import (
	"reflect"
	"sync"
	"time"
)

// DatabaseKey represents key-value pair key. Key is stored as string.
type DatabaseKey string

// DatabaseStringValue represents key-value pair string value. Value is stored as string.
type DatabaseStringValue string

// DatabaseData holds the data stored in a database.
type databaseStoredData struct {
	// stringData holds String keys.
	stringData map[DatabaseKey]DatabaseStringValue
	// hashMapData holds HashMap keys.
	hashMapData map[DatabaseKey]map[string]string
}

func newDatabaseStoredData() *databaseStoredData {
	return &databaseStoredData{
		stringData:  make(map[DatabaseKey]DatabaseStringValue),
		hashMapData: make(map[DatabaseKey]map[string]string),
	}
}

// Database is a namespace for storing key-value pairs.
type Database struct {
	// The name of the database.
	Name string
	// UTC timestamp describing when the database was created.
	CreatedAt time.Time
	// UTC timestamp describing when the database was updated.
	UpdatedAt time.Time
	// The data stored in this database.
	storedData databaseStoredData
	// The current number of keys in this database.
	keyCount uint32
	mutex    sync.RWMutex
}

// Creates a new instance of Database.
func newDatabase(name string) *Database {
	return &Database{
		Name:       name,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		storedData: *newDatabaseStoredData(),
		keyCount:   0,
	}
}

// GetName returns the name of the database.
func (db *Database) GetName() string {
	return db.Name
}

// update updates the database changing some of its fields.
func (db *Database) update() {
	db.UpdatedAt = time.Now().UTC()
}

// keyExists returns true if the key exists in the database.
func (db *Database) keyExists(key DatabaseKey) bool {
	_, exists := db.storedData.stringData[key]
	if exists {
		return true
	}
	_, exists = db.storedData.hashMapData[key]
	return exists
}

// GetKeyCount returns the number of keys in the database.
func (db *Database) GetKeyCount() uint32 {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	return db.keyCount
}

// GetStoredSizeBytes returns the size of stored data in bytes.
func (db *Database) GetStoredSizeBytes() uint64 {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	var size uint64

	for key, value := range db.storedData.stringData {
		size += uint64(reflect.TypeOf(key).Size())
		size += uint64(len(key))
		size += uint64(reflect.TypeOf(value).Size())
		size += uint64(len(value))
	}

	for key, value := range db.storedData.hashMapData {
		size += uint64(reflect.TypeOf(key).Size())
		size += uint64(len(key))
		size += uint64(reflect.TypeOf(value).Size())
		size += uint64(len(value))
		for field, fieldValue := range value {
			size += uint64(reflect.TypeOf(field).Size())
			size += uint64(len(field))
			size += uint64(reflect.TypeOf(fieldValue).Size())
			size += uint64(len(fieldValue))
		}
	}

	return size
}

// CreateDatabase creates a new database with a name.
func CreateDatabase(name string) *Database {
	return newDatabase(name)
}

// GetTypeOfKey returns the data type of the key if it exists.
// The returned bool is true if the key exists and false if it doesn't.
func (db *Database) GetTypeOfKey(key DatabaseKey) (string, bool) {
	_, exists := db.storedData.stringData[key]
	if exists {
		return "String", true
	}
	_, exists = db.storedData.hashMapData[key]
	if exists {
		return "HashMap", true
	}

	return "", false
}

// GetString retrieves a string value using a key.
// The returned bool is true if the key exists.
func (db *Database) GetString(key DatabaseKey) (DatabaseStringValue, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	value, exists := db.storedData.stringData[key]
	return value, exists
}

// SetString sets a string value using a key, overwriting previous value.
// Creates the key if it doesn't exist.
func (db *Database) SetString(key DatabaseKey, value DatabaseStringValue) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if !db.keyExists(key) {
		db.keyCount++
	}

	// Overwrite other data types
	delete(db.storedData.hashMapData, key)

	db.storedData.stringData[key] = value
	db.update()
}

// DeleteKey deletes a key and the value it is holding.
// Returns true if the key exists and it was deleted.
// Returns false if the key doesn't exist.
func (db *Database) DeleteKey(key DatabaseKey) bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if !db.keyExists(key) {
		return false
	}

	_, exists := db.storedData.stringData[key]
	if exists {
		delete(db.storedData.stringData, key)
		db.keyCount--
	}
	_, exists = db.storedData.hashMapData[key]
	if exists {
		delete(db.storedData.hashMapData, key)
		db.keyCount--
	}
	db.update()

	return true
}

// DeleteAllKeys deletes all the keys.
func (db *Database) DeleteAllKeys() {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	for key := range db.storedData.stringData {
		delete(db.storedData.stringData, key)
	}
	for key := range db.storedData.hashMapData {
		delete(db.storedData.hashMapData, key)
	}

	db.keyCount = 0
	db.update()
}

// GetKeys returns all the keys.
func (db *Database) GetKeys() []string {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	var keys []string
	for key := range db.storedData.stringData {
		keys = append(keys, string(key))
	}
	for key := range db.storedData.hashMapData {
		keys = append(keys, string(key))
	}

	return keys
}

// SetHashMap sets fields in a HashMap value using a key, overwriting previous fields.
// Creates the key if it doesn't exist. Returns the number of added fields.
func (db *Database) SetHashMap(key DatabaseKey, fields map[string]string, maxFieldLimit uint32) uint32 {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if !db.keyExists(key) {
		db.keyCount++
	}

	// Overwrite other data types
	delete(db.storedData.stringData, key)

	_, exists := db.storedData.hashMapData[key]
	if !exists {
		db.storedData.hashMapData[key] = make(map[string]string)
	}

	var fieldsAdded uint32 = 0
	for field, fieldValue := range fields {
		if uint32(len(db.storedData.hashMapData[key])) >= maxFieldLimit {
			return fieldsAdded
		}

		_, exists := db.storedData.hashMapData[key][field]
		if !exists {
			fieldsAdded++
		}
		db.storedData.hashMapData[key][field] = fieldValue
	}
	db.update()

	return fieldsAdded
}

// GetHashMapFieldValue returns a single HashMap field value using a key.
// The returned bool is true if the field exists in the HashMap,
// or false if the key or field doesn't exist.
func (db *Database) GetHashMapFieldValue(key DatabaseKey, field string) (string, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	keyValue, exists := db.storedData.hashMapData[key]
	if !exists {
		return "", false
	}

	fieldValue, exists := keyValue[field]
	return fieldValue, exists
}

// DeleteHashMapFields removes fields from a HashMap using a key.
// Returns the number of removed fields. The returned bool is true if the key exists and holds a HashMap.
func (db *Database) DeleteHashMapFields(key DatabaseKey, fields []string) (uint32, bool) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	keyValue, keyExists := db.storedData.hashMapData[key]
	if !keyExists {
		return 0, false
	}
	if len(keyValue) == 0 {
		return 0, true
	}

	var fieldsRemoved uint32 = 0
	for _, field := range fields {
		_, fieldExists := db.storedData.hashMapData[key][field]
		if fieldExists {
			delete(db.storedData.hashMapData[key], field)
			fieldsRemoved++
		}
	}

	if fieldsRemoved > 0 {
		db.update()
	}

	return fieldsRemoved, true
}

// GetAllHashMapFieldsAndValues returns all the fields and values of a HashMap.
// The returned map is empty if the key doesn't exist.
// The returned bool is true if the key exists and holds a HashMap.
func (db *Database) GetAllHashMapFieldsAndValues(key DatabaseKey) (map[string]string, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	value, exists := db.storedData.hashMapData[key]
	if !exists {
		return make(map[string]string), false
	}

	return value, true
}

// GetHashMapFieldCount returns the number of fields in a HashMap.
func (db *Database) GetHashMapFieldCount(key DatabaseKey) uint32 {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	return uint32(len(db.storedData.hashMapData[key]))
}
