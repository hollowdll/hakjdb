package kvdb

import (
	"sync"
	"time"
	"unsafe"
)

// DBKey represents a database key.
type DBKey string

// StringValue represents a String data type value.
type StringValue []byte

// HashMapValue represents a HashMap data type value.
type HashMapValue map[string]HashMapField

// StringKey is a database key that holds a String value.
type StringKey struct {
	Value StringValue
}

// HashMapKey is a database key that holds a HashMap value.
type HashMapKey struct {
	Value HashMapValue
}

// HashMapField is a HashMap field that holds a String value.
type HashMapField struct {
	Value StringValue
}

type HashMapFieldValueResult struct {
	// Value is the value the field is holding.
	FieldValue HashMapField
	// Ok is true if the field exists. Otherwise false.
	Ok bool
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

// DBConfig contains fields to configure a database.
type DBConfig struct {
	// The maximum number of fields a HashMap key value can hold.
	MaxHashMapFields uint32
}

// DB is a database used as a namespace for storing key-value pairs.
type DB struct {
	// The name of the database.
	name string
	// The description of the database.
	description string
	// Timestamp describing when the database was created.
	createdAt time.Time
	// Timestamp describing when the database was updated.
	updatedAt time.Time
	// The data stored in this database.
	storedData dbStoredDataExperimental
	// The current number of keys in this database.
	keyCount uint32
	cfg      DBConfig
	mu       sync.RWMutex
}

func NewDB(name string, desc string, cfg DBConfig) *DB {
	return &DB{
		name:        name,
		description: desc,
		createdAt:   time.Now().UTC(),
		updatedAt:   time.Now().UTC(),
		storedData:  *newDBStoredDataExperimental(),
		keyCount:    0,
		cfg:         cfg,
	}
}

func (db *DB) Name() string {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.name
}

func (db *DB) Description() string {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.description
}

func (db *DB) CreatedAt() time.Time {
	return db.createdAt
}

func (db *DB) UpdatedAt() time.Time {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.updatedAt
}

func (db *DB) UpdateName(newName string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.name = newName
	db.update()
}

func (db *DB) UpdateDescription(newDescription string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.description = newDescription
	db.update()
}

func (db *DB) update() {
	db.updatedAt = time.Now().UTC()
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
func (db *DB) GetKeyCount() uint32 {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.keyCount
}

// GetEstimatedStorageSizeBytes returns the estimated size of stored data in bytes.
func (db *DB) GetEstimatedStorageSizeBytes() uint64 {
	db.mu.RLock()
	defer db.mu.RUnlock()
	var size uint64

	for k, v := range db.storedData.stringData {
		size += stringKeyEstimatedMemoryUsageBytes(k, v)
	}

	for k, v := range db.storedData.hashMapData {
		size += hashMapKeyEstimatedMemoryUsageBytes(k, v)
	}

	return size
}

func storedTypeBytes(t any) uint64 {
	return uint64(unsafe.Sizeof(t))
}

func stringKeyEstimatedMemoryUsageBytes(k DBKey, v StringKey) uint64 {
	var size uint64
	size += storedTypeBytes(k) + uint64(len(k))
	size += storedTypeBytes(v) + uint64(len(v.Value))
	return size
}

func hashMapKeyEstimatedMemoryUsageBytes(k DBKey, v HashMapKey) uint64 {
	var size uint64
	size += storedTypeBytes(k) + uint64(len(k))
	size += storedTypeBytes(v)
	for field, fieldValue := range v.Value {
		size += storedTypeBytes(field) + uint64(len(field))
		size += storedTypeBytes(fieldValue) + uint64(len(fieldValue.Value))
	}
	return size
}

// GetKeyType returns the data type of the key if it exists.
// The returned bool is true if the key exists and false if it doesn't.
func (db *DB) GetKeyType(key string) (DBKeyType, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	_, exists := db.storedData.stringData[DBKey(key)]
	if exists {
		return StringKeyTypeName, true
	}
	_, exists = db.storedData.hashMapData[DBKey(key)]
	if exists {
		return HashMapKeyTypeName, true
	}

	return "", false
}

// GetStringKey retrieves a String key value.
// The returned bool is true if the key exists and holds a String.
func (db *DB) GetStringKey(key string) (StringKey, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	kv, ok := db.storedData.stringData[DBKey(key)]
	return kv, ok
}

// SetString sets a String key value, overwriting previous key.
// Creates the key if it doesn't exist.
func (db *DB) SetString(key string, value []byte) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if !db.keyExists(DBKey(key)) {
		db.keyCount++
	}

	// Overwrite other data types
	delete(db.storedData.hashMapData, DBKey(key))

	db.storedData.stringData[DBKey(key)] = StringKey{Value: value}
	db.update()
}

// DeleteKeys deletes the specified keys.
// Returns the number of keys that were deleted.
func (db *DB) DeleteKeys(keys []string) uint32 {
	db.mu.Lock()
	defer db.mu.Unlock()

	var keysDeleted uint32 = 0
	for _, key := range keys {
		_, ok := db.storedData.stringData[DBKey(key)]
		if ok {
			delete(db.storedData.stringData, DBKey(key))
			keysDeleted++
			db.keyCount--
			continue
		}
		_, ok = db.storedData.hashMapData[DBKey(key)]
		if ok {
			delete(db.storedData.hashMapData, DBKey(key))
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
func (db *DB) DeleteAllKeys() {
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

// GetAllKeys returns all the keys.
func (db *DB) GetAllKeys() []string {
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

// SetHashMap sets the specified fields in a HashMap key value, overwriting previous fields.
// Creates the key if it doesn't exist. Returns the number of added fields.
func (db *DB) SetHashMap(key string, fields map[string][]byte) uint32 {
	db.mu.Lock()
	defer db.mu.Unlock()

	if !db.keyExists(DBKey(key)) {
		db.keyCount++
	}

	// Overwrite other data types
	delete(db.storedData.stringData, DBKey(key))

	_, exists := db.storedData.hashMapData[DBKey(key)]
	if !exists {
		db.storedData.hashMapData[DBKey(key)] = HashMapKey{
			Value: HashMapValue(make(map[string]HashMapField)),
		}
	}

	var fieldsAdded uint32 = 0
	for field, fieldValue := range fields {
		_, exists := db.storedData.hashMapData[DBKey(key)].Value[field]
		if !exists {
			// ignore new fields if max limit is reached
			if uint32(len(db.storedData.hashMapData[DBKey(key)].Value)) >= db.cfg.MaxHashMapFields {
				continue
			}
			fieldsAdded++
		}
		db.storedData.hashMapData[DBKey(key)].Value[field] = HashMapField{Value: fieldValue}
	}
	db.update()

	return fieldsAdded
}

// GetHashMapFieldValues returns a HashMap key value's field values.
// The returned bool is true if the key exists,
// or false if the key doesn't exist.
func (db *DB) GetHashMapFieldValues(key string, fields []string) (map[string]*HashMapFieldValueResult, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	fieldValueMap := make(map[string]*HashMapFieldValueResult)
	kv, keyExists := db.storedData.hashMapData[DBKey(key)]
	if !keyExists {
		return fieldValueMap, false
	}

	for _, field := range fields {
		value, ok := kv.Value[field]
		fieldValue := &HashMapFieldValueResult{
			FieldValue: value,
			Ok:         ok,
		}
		fieldValueMap[field] = fieldValue
	}

	return fieldValueMap, keyExists
}

// GetHashMapKey retrieves a HashMap key value.
// The returned map is empty if the key doesn't exist.
// The returned bool is true if the key exists and holds a HashMap.
func (db *DB) GetHashMapKey(key string) (HashMapKey, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	kv, ok := db.storedData.hashMapData[DBKey(key)]
	return kv, ok
}

// DeleteHashMapFields removes the specified fields from a HashMap key value.
// Returns the number of removed fields. The returned bool is true if the key exists and holds a HashMap.
func (db *DB) DeleteHashMapFields(key string, fields []string) (uint32, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()

	kv, keyExists := db.storedData.hashMapData[DBKey(key)]
	if !keyExists {
		return 0, false
	}
	if len(kv.Value) == 0 {
		return 0, true
	}

	var fieldsRemoved uint32 = 0
	for _, field := range fields {
		_, fieldExists := db.storedData.hashMapData[DBKey(key)].Value[field]
		if fieldExists {
			delete(db.storedData.hashMapData[DBKey(key)].Value, field)
			fieldsRemoved++
		}
	}

	if fieldsRemoved > 0 {
		db.update()
	}

	return fieldsRemoved, true
}

/* DISABLED
// GetHashMapFieldCount returns the number of fields in a HashMap key value.
func (db *DB) GetHashMapFieldCount(key string) uint32 {
	db.mu.RLock()
	defer db.mu.RUnlock()

	return uint32(len(db.storedData.hashMapData[DBKey(key)].value))
}
*/
