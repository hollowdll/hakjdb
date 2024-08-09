package kvdb

import (
	"bytes"
	"testing"
	"time"

	"github.com/hollowdll/kvdb/internal/common"
)

func TestGetKeyType(t *testing.T) {
	dbCfg := DBConfig{
		MaxHashMapFields: common.HashMapMaxFields,
	}

	t.Run("KeyNotFound", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		keyType, ok := db.GetKeyType("key1")

		expectedKeyType := DBKeyType("")
		if keyType != expectedKeyType {
			t.Errorf("expected key type = %s; got = %s", expectedKeyType, keyType)
		}

		expectedOk := false
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})

	t.Run("StringKey", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		db.SetString("key1", []byte("value"))
		keyType, ok := db.GetKeyType("key1")

		expectedKeyType := DBKeyType("String")
		if keyType != expectedKeyType {
			t.Errorf("expected key type = %s; got = %s", expectedKeyType, keyType)
		}

		expectedOk := true
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})

	t.Run("HashMapKey", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		db.SetHashMap("key1", make(map[string][]byte))
		keyType, ok := db.GetKeyType("key1")

		expectedKeyType := DBKeyType("HashMap")
		if keyType != expectedKeyType {
			t.Errorf("expected key type = %s; got = %s", expectedKeyType, keyType)
		}

		expectedOk := true
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})
}

func TestSetString(t *testing.T) {
	dbCfg := DBConfig{
		MaxHashMapFields: common.HashMapMaxFields,
	}

	t.Run("SetNonExistentKey", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		db.SetString("key1", []byte("value1"))

		var expectedKeys = 1
		keys := db.GetKeyCount()
		if keys != expectedKeys {
			t.Errorf("expected keys = %d; got = %d", expectedKeys, keys)
		}
	})

	t.Run("OverwriteExistingKey", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		db.SetString("key1", []byte("value1"))
		db.SetString("key1", []byte("value2"))

		var expectedKeys = 1
		keys := db.GetKeyCount()
		if keys != expectedKeys {
			t.Errorf("expected keys = %d; got = %d", expectedKeys, keys)
		}
	})

	t.Run("DatabaseIsUpdated", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		originalTime := db.UpdatedAt()

		time.Sleep(10 * time.Millisecond)
		db.SetString("key1", []byte("value1"))

		updatedTime := db.UpdatedAt()
		if !updatedTime.After(originalTime) {
			t.Errorf("expected time %s to be after %s", updatedTime, originalTime)
		}
	})
}

func TestGetStringKey(t *testing.T) {
	dbCfg := DBConfig{
		MaxHashMapFields: common.HashMapMaxFields,
	}

	t.Run("GetNonExistentKey", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		kv, ok := db.GetStringKey("key1")

		expectedValue := StringValue("")
		if !bytes.Equal(kv.Value, expectedValue) {
			t.Errorf("expected value = %s; got = %s", expectedValue, kv.Value)
		}

		expectedOk := false
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})

	t.Run("GetExistingKey", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		db.SetString("key1", []byte("value1"))
		kv, ok := db.GetStringKey("key1")

		expectedValue := StringValue("value1")
		if !bytes.Equal(kv.Value, expectedValue) {
			t.Errorf("expected value = %s; got = %s", expectedValue, kv.Value)
		}

		expectedOk := true
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})
}

func TestDeleteKeys(t *testing.T) {
	dbCfg := DBConfig{
		MaxHashMapFields: common.HashMapMaxFields,
	}

	t.Run("DeleteNonExistentKey", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		result := db.DeleteKeys([]string{"key1"})
		var expectedResult uint32 = 0

		if result != expectedResult {
			t.Errorf("expected result = %v; got = %v", expectedResult, result)
		}

		result = db.DeleteKeys([]string{"key2", "key3"})
		expectedResult = 0

		if result != expectedResult {
			t.Errorf("expected result = %v; got = %v", expectedResult, result)
		}
	})

	t.Run("DeleteExistingKey", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		db.SetString("key1", []byte("value1"))

		result := db.DeleteKeys([]string{"key1"})
		var expectedResult uint32 = 1
		if result != expectedResult {
			t.Errorf("expected result = %v; got = %v", expectedResult, result)
		}

		result = db.DeleteKeys([]string{"key1"})
		expectedResult = 0
		if result != expectedResult {
			t.Errorf("expected result = %v; got = %v", expectedResult, result)
		}
	})

	t.Run("DeleteMultipleExistingKeys", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		db.SetString("key1", []byte("val1"))
		db.SetString("key2", []byte("val2"))
		db.SetString("key3", []byte("val3"))

		result := db.DeleteKeys([]string{"key1", "key2", "key3"})
		var expectedResult uint32 = 3
		if result != expectedResult {
			t.Errorf("expected result = %v; got = %v", expectedResult, result)
		}

		result = db.DeleteKeys([]string{"key3", "key1", "key2"})
		expectedResult = 0
		if result != expectedResult {
			t.Errorf("expected result = %v; got = %v", expectedResult, result)
		}
	})

	t.Run("DeleteMultipleKeysButNotAll", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		db.SetString("key1", []byte("val1"))
		db.SetString("key2", []byte("val2"))
		db.SetString("key3", []byte("val3"))
		db.SetString("key4", []byte("val4"))

		result := db.DeleteKeys([]string{"key3", "key4"})
		var expectedResult uint32 = 2
		if result != expectedResult {
			t.Errorf("expected result = %v; got = %v", expectedResult, result)
		}

		result = db.DeleteKeys([]string{"key2"})
		expectedResult = 1
		if result != expectedResult {
			t.Errorf("expected result = %v; got = %v", expectedResult, result)
		}

		result = db.DeleteKeys([]string{"key2", "key3", "key4"})
		expectedResult = 0
		if result != expectedResult {
			t.Errorf("expected result = %v; got = %v", expectedResult, result)
		}

		result = db.DeleteKeys([]string{"key1", "key2", "key3", "key4"})
		expectedResult = 1
		if result != expectedResult {
			t.Errorf("expected result = %v; got = %v", expectedResult, result)
		}
	})

	t.Run("DatabaseIsNotUpdatedIfKeyNotExists", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		originalTime := db.UpdatedAt()

		time.Sleep(10 * time.Millisecond)
		db.DeleteKeys([]string{"key1"})

		updatedTime := db.UpdatedAt()
		if !updatedTime.Equal(originalTime) {
			t.Errorf("expected times to be equal; updated time = %s; original time = %s", updatedTime, originalTime)
		}
	})

	t.Run("DatabaseIsUpdatedIfKeyDeleted", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		db.SetString("key1", []byte("val1"))

		originalTime := db.UpdatedAt()
		time.Sleep(10 * time.Millisecond)
		db.DeleteKeys([]string{"key1"})

		updatedTime := db.UpdatedAt()
		if !updatedTime.After(originalTime) {
			t.Errorf("expected time %s to be after %s", updatedTime, originalTime)
		}
	})
}

func TestGetKeyCount(t *testing.T) {
	dbCfg := DBConfig{
		MaxHashMapFields: common.HashMapMaxFields,
	}
	db := NewDB("test", "", dbCfg)
	count := db.GetKeyCount()
	if count != 0 {
		t.Fatalf("key count should be 0 but got %d", count)
	}

	db.SetString("key1", []byte("value1"))
	count = db.GetKeyCount()
	if count != 1 {
		t.Fatalf("key count should be 1 but got %d", count)
	}

	db.SetHashMap("key2", map[string][]byte{
		"field1": []byte("value1"),
	})
	count = db.GetKeyCount()
	if count != 2 {
		t.Fatalf("key count should be 2 but got %d", count)
	}

	db.DeleteKeys([]string{"key1"})
	count = db.GetKeyCount()
	if count != 1 {
		t.Fatalf("key count should be 1 but got %d", count)
	}

	db.DeleteKeys([]string{"key2"})
	count = db.GetKeyCount()
	if count != 0 {
		t.Fatalf("key count should be 0 but got %d", count)
	}

	db.SetString("key1", []byte("value1"))
	db.SetHashMap("key2", map[string][]byte{
		"field1": []byte("value1"),
	})
	count = db.GetKeyCount()
	if count != 2 {
		t.Fatalf("key count should be 2 but got %d", count)
	}
}

func TestDeleteAllKeys(t *testing.T) {
	dbCfg := DBConfig{
		MaxHashMapFields: common.HashMapMaxFields,
	}

	t.Run("NoKeys", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		db.DeleteAllKeys()
		count := db.GetKeyCount()
		expectedCount := 0
		if count != 0 {
			t.Errorf("expected keys = %d; got = %d", expectedCount, count)
		}
	})

	t.Run("MultipleKeys", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		keys := []string{"key1", "key2", "key3"}
		for _, key := range keys {
			db.SetString(key, []byte("value"))
		}

		count := db.GetKeyCount()
		var expectedCount = 3
		if count != expectedCount {
			t.Fatalf("expected keys = %d; got = %d", expectedCount, count)
		}

		db.DeleteAllKeys()
		count = db.GetKeyCount()
		expectedCount = 0
		if count != expectedCount {
			t.Errorf("expected keys = %d; got = %d", expectedCount, count)
		}
	})
}

func TestGetAllKeys(t *testing.T) {
	dbCfg := DBConfig{
		MaxHashMapFields: common.HashMapMaxFields,
	}

	t.Run("NoKeys", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		keys := db.GetAllKeys()
		expectedKeys := 0
		if len(keys) != expectedKeys {
			t.Errorf("expected keys = %d; got = %d", expectedKeys, len(keys))
		}
	})

	t.Run("MultipleKeys", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		keys := []string{"key1", "key2", "key3"}
		for _, key := range keys {
			db.SetString(key, []byte("val"))
		}

		actualKeys := db.GetAllKeys()
		if len(actualKeys) != len(keys) {
			t.Fatalf("expected keys = %d; got = %d", len(keys), len(actualKeys))
		}

		for _, key := range actualKeys {
			if !common.StringInSlice(key, keys) {
				t.Fatalf("expected key %s to be in %v", key, keys)
			}
		}
	})
}

func TestSetHashMap(t *testing.T) {
	dbCfg := DBConfig{
		MaxHashMapFields: common.HashMapMaxFields,
	}
	fields := make(map[string][]byte)
	fields["field1"] = []byte("value1")
	fields["field2"] = []byte("value2")
	fields["field3"] = []byte("value3")

	t.Run("SetNonExistentKey", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		fieldsAdded := db.SetHashMap("key1", fields)

		var expectedFieldsAdded uint32 = 3
		if fieldsAdded != expectedFieldsAdded {
			t.Errorf("expected fields added = %d; got = %d", expectedFieldsAdded, fieldsAdded)
		}

		var expectedKeys = 1
		keys := db.GetKeyCount()
		if keys != expectedKeys {
			t.Errorf("expected keys = %d; got = %d", expectedKeys, keys)
		}
	})

	t.Run("OverwriteExistingHashMapKey", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		db.SetHashMap("key1", fields)
		fieldsAdded := db.SetHashMap("key1", make(map[string][]byte))

		var expectedFieldsAdded uint32 = 0
		if fieldsAdded != expectedFieldsAdded {
			t.Errorf("expected fields added = %d; got = %d", expectedFieldsAdded, fieldsAdded)
		}

		var expectedKeys = 1
		keys := db.GetKeyCount()
		if keys != expectedKeys {
			t.Errorf("expected keys = %d; got = %d", expectedKeys, keys)
		}
	})

	t.Run("DatabaseIsUpdated", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		originalTime := db.UpdatedAt()

		time.Sleep(10 * time.Millisecond)
		db.SetHashMap("key1", fields)

		updatedTime := db.UpdatedAt()
		if !updatedTime.After(originalTime) {
			t.Errorf("expected time %s to be after %s", updatedTime, originalTime)
		}
	})

	t.Run("MaxFieldLimitReached", func(t *testing.T) {
		dbCfg := dbCfg
		dbCfg.MaxHashMapFields = 2
		db := NewDB("test", "", dbCfg)
		fields2 := make(map[string][]byte)
		fields2["field4"] = []byte("val")
		fields2["field5"] = []byte("val")
		fields2["field6"] = []byte("val")

		fieldsAdded1 := db.SetHashMap("key1", fields)
		fieldsAdded2 := db.SetHashMap("key1", fields2)

		var expectedFieldsAdded1 uint32 = 2
		if fieldsAdded1 != expectedFieldsAdded1 {
			t.Errorf("expected fields added = %d; got = %d", expectedFieldsAdded1, fieldsAdded1)
		}

		var expectedFieldsAdded2 uint32 = 0
		if fieldsAdded2 != expectedFieldsAdded2 {
			t.Errorf("expected fields added = %d; got = %d", expectedFieldsAdded2, fieldsAdded2)
		}
	})
}

func TestGetHashMapFieldValues(t *testing.T) {
	dbCfg := DBConfig{
		MaxHashMapFields: common.HashMapMaxFields,
	}
	fields := make(map[string][]byte)
	fields["field1"] = []byte("value1")
	fields["field2"] = []byte("value2")
	fields["field3"] = []byte("value3")

	t.Run("GetNonExistentKey", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		values, ok := db.GetHashMapFieldValues("key1", []string{"field2"})

		expectedFields := 0
		if len(values) != expectedFields {
			t.Errorf("expected fields = %d; got = %d", expectedFields, len(values))
		}

		expectedOk := false
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})

	t.Run("GetNonExistentField", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		key := "key1"
		db.SetHashMap(key, fields)
		values, _ := db.GetHashMapFieldValues(key, []string{"field1234"})

		expectedValue := []byte("")
		if !bytes.Equal(values["field1234"].FieldValue.Value, expectedValue) {
			t.Errorf("expected value = %s; got = %s", expectedValue, values["field1234"].FieldValue.Value)
		}

		expectedOk := false
		if values["field1234"].Ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, values["field1234"].Ok)
		}
	})

	t.Run("GetExistingKeyAndFields", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		key := "key1"
		db.SetHashMap(key, fields)
		values, ok := db.GetHashMapFieldValues(key, []string{"field1", "field2", "field3"})

		for field, expectedValue := range fields {
			if !bytes.Equal(values[field].FieldValue.Value, expectedValue) {
				t.Errorf("expected value = %s; got = %s", expectedValue, values[field].FieldValue.Value)
			}
			if !values[field].Ok {
				t.Errorf("expected ok = %v; got = %v", true, values[field].Ok)
			}
		}

		expectedOk := true
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})
}

func TestDeleteHashMapFields(t *testing.T) {
	dbCfg := DBConfig{
		MaxHashMapFields: common.HashMapMaxFields,
	}
	fields := make(map[string][]byte)
	fields["field1"] = []byte("val")
	fields["field2"] = []byte("val")
	fields["field3"] = []byte("val")

	t.Run("KeyNotFound", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		fieldsRemoved, ok := db.DeleteHashMapFields("key1", []string{"field2", "field3"})

		var expectedFieldsRemoved uint32 = 0
		if fieldsRemoved != expectedFieldsRemoved {
			t.Errorf("expected fields removed = %d; got = %d", expectedFieldsRemoved, fieldsRemoved)
		}

		expectedOk := false
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})

	t.Run("FieldsExist", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		db.SetHashMap("key1", fields)
		fieldsRemoved, ok := db.DeleteHashMapFields("key1", []string{"field2", "field3"})

		var expectedFieldsRemoved uint32 = 2
		if fieldsRemoved != expectedFieldsRemoved {
			t.Errorf("expected fields removed = %d; got = %d", expectedFieldsRemoved, fieldsRemoved)
		}

		expectedOk := true
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})

	t.Run("FieldsNotFound", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		db.SetHashMap("key1", fields)
		fieldsRemoved, ok := db.DeleteHashMapFields("key1", []string{"field123", "field1234"})

		var expectedFieldsRemoved uint32 = 0
		if fieldsRemoved != expectedFieldsRemoved {
			t.Errorf("expected fields removed = %d; got = %d", expectedFieldsRemoved, fieldsRemoved)
		}

		expectedOk := true
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})

	t.Run("DuplicateFields", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		db.SetHashMap("key1", fields)
		fieldsRemoved, ok := db.DeleteHashMapFields("key1", []string{"field1", "field1", "field1"})

		var expectedFieldsRemoved uint32 = 1
		if fieldsRemoved != expectedFieldsRemoved {
			t.Errorf("expected fields removed = %d; got = %d", expectedFieldsRemoved, fieldsRemoved)
		}

		expectedOk := true
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})
}

func TestGetHashMapKey(t *testing.T) {
	dbCfg := DBConfig{
		MaxHashMapFields: common.HashMapMaxFields,
	}
	fields := make(map[string][]byte)
	fields["field1"] = []byte("val")
	fields["field2"] = []byte("val")
	fields["field3"] = []byte("val")

	t.Run("GetNonExistentKey", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		kv, ok := db.GetHashMapKey("key1")

		expectedElements := 0
		if len(kv.Value) != expectedElements {
			t.Errorf("expected elements = %d; got = %d", expectedElements, len(kv.Value))
		}

		expectedOk := false
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})

	t.Run("GetExistingKey", func(t *testing.T) {
		db := NewDB("test", "", dbCfg)
		key := "key1"
		db.SetHashMap(key, fields)
		kv, ok := db.GetHashMapKey(key)

		expectedElements := 3
		if len(kv.Value) != expectedElements {
			t.Errorf("expected elements = %d; got = %d", expectedElements, len(kv.Value))
		}

		expectedOk := true
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}

		for expectedField, expectedValue := range fields {
			actualValue, exists := kv.Value[expectedField]
			if !exists {
				t.Errorf("expected field '%s'", expectedField)
			}
			if !bytes.Equal(expectedValue, actualValue.Value) {
				t.Errorf("expected field value = %s; got = %s", expectedValue, actualValue)
			}
		}
	})
}
