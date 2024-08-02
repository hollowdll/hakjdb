package kvdb

import (
	"reflect"
	"sync"
	"time"
)

// DBKey represents a database key.
type DBKey string

// HashMapFieldKey represents the key of a HashMap field.
type HashMapFieldKey string

// StringValue represents a String data type value.
type StringValue []byte

// HashMapValue represents a HashMap data type value.
type HashMapValue map[HashMapFieldKey]HashMapField

// StringKey is a database key that holds a String value.
type StringKey struct {
	value StringValue
}

// HashMapKey is a database key that holds a HashMap value.
type HashMapKey struct {
	value HashMapValue
}

// HashMapField is a HashMap field that holds a String value.
type HashMapField struct {
	value StringValue
}

// dbStoredDataExperimental holds the data stored in a database.
// THIS IS EXPERIMENTAL!
type dbStoredDataExperimental struct {
	// stringData holds String keys.
	stringData map[DBKey]StringKey
	// hashMapData holds HashMap keys.
	hashMapData map[DBKey]HashMapKey
}

func newDBStoredDataExperimental() *dbStoredDataExperimental {
	return &dbStoredDataExperimental{
		stringData:  make(map[DBKey]StringKey),
		hashMapData: make(map[DBKey]HashMapKey),
	}
}

// DB is a database used as a namespace for storing key-value pairs.
type DB struct {
	// The name of the database.
	Name string
	// The description of the database.
	Description string
	// Timestamp describing when the database was created.
	CreatedAt time.Time
	// Timestamp describing when the database was updated.
	UpdatedAt time.Time
	// The data stored in this database.
	storedData dbStoredDataExperimental
	// The current number of keys in this database.
	keyCount uint32
	mu       sync.RWMutex
}

func newDB(name string) *DB {
	return &DB{
		Name:        name,
		Description: "",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		storedData:  *newDBStoredDataExperimental(),
		keyCount:    0,
	}
}

func (db *DB) update() {
	db.UpdatedAt = time.Now().UTC()
}

// keyExists returns true if the key exists in the database.
func (db *DB) keyExists(key DBKey) bool {
	_, exists := db.storedData.stringData[key]
	if exists {
		return true
	}
	_, exists = db.storedData.hashMapData[key]
	return exists
}

// GetKeyCount returns the number of keys in the database.
func (db *Database) GetKeyCount() uint32 {
	db.mu.RLock()
	defer db.mu.RUnlock()

	return db.keyCount
}

// GetStoredSizeBytes returns the size of stored data in bytes.
func (db *Database) GetStoredSizeBytes() uint64 {
	db.mu.RLock()
	defer db.mu.RUnlock()
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

// CreateDatabase creates a new database with the given name.
func CreateDatabase(name string) *Database {
	return newDatabase(name)
}

// GetKeyType returns the data type of the key if it exists.
// The returned bool is true if the key exists and false if it doesn't.
func (db *Database) GetKeyType(key string) (DBKeyType, bool) {
	_, exists := db.storedData.stringData[key]
	if exists {
		return StringKeyTypeName, true
	}
	_, exists = db.storedData.hashMapData[key]
	if exists {
		return HashMapKeyTypeName, true
	}

	return "", false
}

// GetString retrieves a string value using a key.
// The returned bool is true if the key exists.
func (db *Database) GetString(key string) (string, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	value, exists := db.storedData.stringData[key]
	return value, exists
}

// SetString sets a string value using a key, overwriting previous value.
// Creates the key if it doesn't exist.
func (db *Database) SetString(key string, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if !db.keyExists(key) {
		db.keyCount++
	}

	// Overwrite other data types
	delete(db.storedData.hashMapData, key)

	db.storedData.stringData[key] = value
	db.update()
}

// DeleteKeys deletes the specified keys and the values they are holding.
// Returns the number of keys that were deleted.
func (db *Database) DeleteKeys(keys []string) uint32 {
	db.mu.Lock()
	defer db.mu.Unlock()

	var keysDeleted uint32 = 0
	for _, key := range keys {
		_, ok := db.storedData.stringData[string(key)]
		if ok {
			delete(db.storedData.stringData, string(key))
			keysDeleted++
			db.keyCount--
			continue
		}
		_, ok = db.storedData.hashMapData[string(key)]
		if ok {
			delete(db.storedData.hashMapData, string(key))
			keysDeleted++
			db.keyCount--
		}
	}

	if keysDeleted > 0 {
		db.update()
	}

	return keysDeleted
}

// DeleteAllKeys deletes all the keys.
func (db *Database) DeleteAllKeys() {
	db.mu.Lock()
	defer db.mu.Unlock()
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
	db.mu.RLock()
	defer db.mu.RUnlock()

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
func (db *Database) SetHashMap(key string, fields map[string]string, maxFieldLimit uint32) uint32 {
	db.mu.Lock()
	defer db.mu.Unlock()

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
		_, exists := db.storedData.hashMapData[key][field]
		if !exists {
			// ignore new fields if max limit is reached
			if uint32(len(db.storedData.hashMapData[key])) >= maxFieldLimit {
				continue
			}
			fieldsAdded++
		}
		db.storedData.hashMapData[key][field] = fieldValue
	}
	db.update()

	return fieldsAdded
}

// GetHashMapFieldValues returns HashMap field values using a key.
// The returned bool is true if the key exists,
// or false if the key doesn't exist.
func (db *Database) GetHashMapFieldValues(key string, fields []string) (map[string]*HashMapFieldValue, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	fieldValueMap := make(map[string]*HashMapFieldValue)
	keyValue, keyExists := db.storedData.hashMapData[key]
	if !keyExists {
		return fieldValueMap, false
	}

	for _, field := range fields {
		value, ok := keyValue[field]
		fieldValue := &HashMapFieldValue{
			Value: value,
			Ok:    ok,
		}
		fieldValueMap[field] = fieldValue
	}

	return fieldValueMap, keyExists
}

// DeleteHashMapFields removes fields from a HashMap using a key.
// Returns the number of removed fields. The returned bool is true if the key exists and holds a HashMap.
func (db *Database) DeleteHashMapFields(key string, fields []string) (uint32, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()

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
func (db *Database) GetAllHashMapFieldsAndValues(key string) (map[string]string, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	value, exists := db.storedData.hashMapData[key]
	if !exists {
		return make(map[string]string), false
	}

	return value, true
}

// GetHashMapFieldCount returns the number of fields in a HashMap.
func (db *Database) GetHashMapFieldCount(key string) uint32 {
	db.mu.RLock()
	defer db.mu.RUnlock()

	return uint32(len(db.storedData.hashMapData[key]))
}
