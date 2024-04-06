package kvdb

import (
	"testing"
	"time"

	"github.com/hollowdll/kvdb/internal/common"
)

func TestCreateDatabase(t *testing.T) {
	db := CreateDatabase("test")
	expectedDb := newDatabase("test")

	if db == nil {
		t.Fatal("expected db but got nil")
	}

	if db.GetName() != expectedDb.GetName() {
		t.Fatalf("expected db with name %s but got %s", expectedDb.GetName(), db.GetName())
	}
}

func TestGetTypeOfKey(t *testing.T) {
	t.Run("KeyNotFound", func(t *testing.T) {
		db := newDatabase("test")
		keyType, ok := db.GetTypeOfKey("key1")

		expectedKeyType := ""
		if keyType != expectedKeyType {
			t.Errorf("expected key type = %s; got = %s", expectedKeyType, keyType)
		}

		expectedOk := false
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})

	t.Run("String", func(t *testing.T) {
		db := newDatabase("test")
		db.SetString("key1", "value")
		keyType, ok := db.GetTypeOfKey("key1")

		expectedKeyType := "String"
		if keyType != expectedKeyType {
			t.Errorf("expected key type = %s; got = %s", expectedKeyType, keyType)
		}

		expectedOk := true
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})

	t.Run("HashMap", func(t *testing.T) {
		db := newDatabase("test")
		db.SetHashMap("key1", make(map[string]string), common.HashMapMaxFields)
		keyType, ok := db.GetTypeOfKey("key1")

		expectedKeyType := "HashMap"
		if keyType != expectedKeyType {
			t.Errorf("expected key type = %s; got = %s", expectedKeyType, keyType)
		}

		expectedOk := true
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})
}

func TestDatabaseSetString(t *testing.T) {
	t.Run("SetNonExistentKey", func(t *testing.T) {
		db := newDatabase("test")
		db.SetString("key1", "value1")

		var expectedKeys uint32 = 1
		keys := db.GetKeyCount()
		if keys != expectedKeys {
			t.Errorf("expected keys = %d; got = %d", expectedKeys, keys)
		}
	})

	t.Run("OverwriteExistingKey", func(t *testing.T) {
		db := newDatabase("test")
		db.SetString("key1", "value1")
		db.SetString("key1", "value2")

		var expectedKeys uint32 = 1
		keys := db.GetKeyCount()
		if keys != expectedKeys {
			t.Errorf("expected keys = %d; got = %d", expectedKeys, keys)
		}
	})

	t.Run("DatabaseIsUpdated", func(t *testing.T) {
		db := newDatabase("test")
		originalTime := db.UpdatedAt

		time.Sleep(10 * time.Millisecond)
		db.SetString("key1", "value1")

		updatedTime := db.UpdatedAt
		if !updatedTime.After(originalTime) {
			t.Errorf("expected time %s to be after %s", updatedTime, originalTime)
		}
	})
}

func TestDatabaseDeleteKey(t *testing.T) {
	t.Run("DeleteNonExistentKey", func(t *testing.T) {
		db := newDatabase("test")
		result := db.DeleteKey("key1")
		expectedResult := false

		if result != expectedResult {
			t.Errorf("expected result = %v; got = %v", expectedResult, result)
		}
	})

	t.Run("DeleteExistingKey", func(t *testing.T) {
		db := newDatabase("test")
		db.SetString("key1", "value1")

		result := db.DeleteKey("key1")
		expectedResult := true
		if result != expectedResult {
			t.Errorf("expected result = %v; got = %v", expectedResult, result)
		}

		result = db.DeleteKey("key1")
		expectedResult = false
		if result != expectedResult {
			t.Errorf("expected result = %v; got = %v", expectedResult, result)
		}
	})

	t.Run("DatabaseIsNotUpdatedIfKeyNonExistent", func(t *testing.T) {
		db := newDatabase("test")
		originalTime := db.UpdatedAt

		time.Sleep(10 * time.Millisecond)
		db.DeleteKey("key1")

		updatedTime := db.UpdatedAt
		if !updatedTime.Equal(originalTime) {
			t.Errorf("expected times to be equal; updated time = %s; original time = %s", updatedTime, originalTime)
		}
	})

	t.Run("DatabaseIsUpdatedIfKeyDeleted", func(t *testing.T) {
		db := newDatabase("test")
		db.SetString("key1", "value1")

		originalTime := db.UpdatedAt
		time.Sleep(10 * time.Millisecond)
		db.DeleteKey("key1")

		updatedTime := db.UpdatedAt
		if !updatedTime.After(originalTime) {
			t.Errorf("expected time %s to be after %s", updatedTime, originalTime)
		}
	})
}

func TestDatabaseGetString(t *testing.T) {
	t.Run("GetNonExistentKey", func(t *testing.T) {
		db := newDatabase("test")
		value, found := db.GetString("key1")

		expectedValue := DatabaseStringValue("")
		if value != expectedValue {
			t.Errorf("expected value = %s; got = %s", expectedValue, value)
		}

		expectedFound := false
		if found != expectedFound {
			t.Errorf("expected found = %v; got = %v", expectedFound, found)
		}
	})

	t.Run("GetExistingKey", func(t *testing.T) {
		db := newDatabase("test")
		expectedValue := DatabaseStringValue("value1")
		key := DatabaseKey("key1")
		db.SetString(key, expectedValue)
		value, found := db.GetString(key)

		if value != expectedValue {
			t.Errorf("expected value = %s; got = %s", expectedValue, value)
		}

		expectedFound := true
		if found != expectedFound {
			t.Errorf("expected found = %v; got = %v", expectedFound, found)
		}
	})
}

func TestGetDatabaseKeyCount(t *testing.T) {
	db := newDatabase("test")
	count := db.GetKeyCount()
	if count != 0 {
		t.Fatalf("key count should be 0 but got %d", count)
	}

	db.SetString("key1", "value1")
	count = db.GetKeyCount()
	if count != 1 {
		t.Fatalf("key count should be 1 but got %d", count)
	}

	db.SetString("key2", "value2")
	count = db.GetKeyCount()
	if count != 2 {
		t.Fatalf("key count should be 2 but got %d", count)
	}

	db.DeleteKey("key1")
	count = db.GetKeyCount()
	if count != 1 {
		t.Fatalf("key count should be 1 but got %d", count)
	}
}

func TestDeleteAllKeys(t *testing.T) {
	t.Run("NoKeys", func(t *testing.T) {
		db := newDatabase("test")
		db.DeleteAllKeys()
		count := db.GetKeyCount()
		expectedCount := 0
		if count != 0 {
			t.Errorf("expected keys = %d; got = %d", expectedCount, count)
		}
	})

	t.Run("MultipleKeys", func(t *testing.T) {
		db := newDatabase("test")
		keys := []DatabaseKey{"key1", "key2", "key3"}
		for _, key := range keys {
			db.SetString(key, "value")
		}

		count := db.GetKeyCount()
		if count != uint32(len(keys)) {
			t.Fatalf("expected keys = %d; got = %d", len(keys), count)
		}

		db.DeleteAllKeys()
		count = db.GetKeyCount()
		var expectedCount uint32 = 0
		if count != expectedCount {
			t.Errorf("expected keys = %d; got = %d", expectedCount, count)
		}
	})
}

func TestGetKeys(t *testing.T) {
	t.Run("NoKeys", func(t *testing.T) {
		db := newDatabase("test")
		keys := db.GetKeys()
		expectedKeys := 0
		if len(keys) != expectedKeys {
			t.Errorf("expected keys = %d; got = %d", expectedKeys, len(keys))
		}
	})

	t.Run("MultipleKeys", func(t *testing.T) {
		db := newDatabase("test")
		keys := []string{"key1", "key2", "key3"}
		for _, key := range keys {
			db.SetString(DatabaseKey(key), "value")
		}

		actualKeys := db.GetKeys()
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
	fields := make(map[string]string)
	fields["field1"] = "value1"
	fields["field2"] = "value2"
	fields["field3"] = "value3"

	t.Run("SetNonExistentKey", func(t *testing.T) {
		db := newDatabase("test")
		db.SetHashMap("key1", fields, common.HashMapMaxFields)

		var expectedKeys uint32 = 1
		keys := db.GetKeyCount()
		if keys != expectedKeys {
			t.Errorf("expected keys = %d; got = %d", expectedKeys, keys)
		}
	})

	t.Run("OverwriteExistingHashMapKey", func(t *testing.T) {
		db := newDatabase("test")
		db.SetHashMap("key1", fields, common.HashMapMaxFields)
		db.SetHashMap("key1", make(map[string]string), common.HashMapMaxFields)

		var expectedKeys uint32 = 1
		keys := db.GetKeyCount()
		if keys != expectedKeys {
			t.Errorf("expected keys = %d; got = %d", expectedKeys, keys)
		}
	})

	t.Run("DatabaseIsUpdated", func(t *testing.T) {
		db := newDatabase("test")
		originalTime := db.UpdatedAt

		time.Sleep(10 * time.Millisecond)
		db.SetHashMap("key1", fields, common.HashMapMaxFields)

		updatedTime := db.UpdatedAt
		if !updatedTime.After(originalTime) {
			t.Errorf("expected time %s to be after %s", updatedTime, originalTime)
		}
	})
}

func TestGetHashMapFieldValue(t *testing.T) {
	fields := make(map[string]string)
	fields["field1"] = "value1"
	fields["field2"] = "value2"
	fields["field3"] = "value3"

	t.Run("GetNonExistentKey", func(t *testing.T) {
		db := newDatabase("test")
		value, ok := db.GetHashMapFieldValue("key1", "field1")

		expectedValue := ""
		if value != expectedValue {
			t.Errorf("expected value = %s; got = %s", expectedValue, value)
		}

		expectedOk := false
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})

	t.Run("GetNonExistentField", func(t *testing.T) {
		db := newDatabase("test")
		key := DatabaseKey("key1")
		db.SetHashMap(key, fields, common.HashMapMaxFields)
		value, ok := db.GetHashMapFieldValue(key, "field12345")

		expectedValue := ""
		if value != expectedValue {
			t.Errorf("expected value = %s; got = %s", expectedValue, value)
		}

		expectedOk := false
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})

	t.Run("GetExistingKeyAndField", func(t *testing.T) {
		db := newDatabase("test")
		key := DatabaseKey("key1")
		db.SetHashMap(key, fields, common.HashMapMaxFields)
		value, ok := db.GetHashMapFieldValue(key, "field2")

		expectedValue := "value2"
		if value != expectedValue {
			t.Errorf("expected value = %s; got = %s", expectedValue, value)
		}

		expectedOk := true
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})
}

func TestDeleteHashMapFields(t *testing.T) {
	fields := make(map[string]string)
	fields["field1"] = "value1"
	fields["field2"] = "value2"
	fields["field3"] = "value3"

	t.Run("KeyNotFound", func(t *testing.T) {
		db := newDatabase("test")
		fieldsRemoved, ok := db.DeleteHashMapFields("key1", []string{"field2", "field3"})

		var expectedFieldsRemoved uint32 = 0
		if fieldsRemoved != expectedFieldsRemoved {
			t.Errorf("expected value = %d; got = %d", expectedFieldsRemoved, fieldsRemoved)
		}

		expectedOk := false
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})

	t.Run("FieldsExist", func(t *testing.T) {
		db := newDatabase("test")
		db.SetHashMap("key1", fields, common.HashMapMaxFields)
		fieldsRemoved, ok := db.DeleteHashMapFields("key1", []string{"field2", "field3"})

		var expectedFieldsRemoved uint32 = 2
		if fieldsRemoved != expectedFieldsRemoved {
			t.Errorf("expected value = %d; got = %d", expectedFieldsRemoved, fieldsRemoved)
		}

		expectedOk := true
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})

	t.Run("FieldsNotFound", func(t *testing.T) {
		db := newDatabase("test")
		db.SetHashMap("key1", fields, common.HashMapMaxFields)
		fieldsRemoved, ok := db.DeleteHashMapFields("key1", []string{"field123", "field1234"})

		var expectedFieldsRemoved uint32 = 0
		if fieldsRemoved != expectedFieldsRemoved {
			t.Errorf("expected value = %d; got = %d", expectedFieldsRemoved, fieldsRemoved)
		}

		expectedOk := true
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})

	t.Run("DuplicateFields", func(t *testing.T) {
		db := newDatabase("test")
		db.SetHashMap("key1", fields, common.HashMapMaxFields)
		fieldsRemoved, ok := db.DeleteHashMapFields("key1", []string{"field1", "field1", "field1"})

		var expectedFieldsRemoved uint32 = 1
		if fieldsRemoved != expectedFieldsRemoved {
			t.Errorf("expected value = %d; got = %d", expectedFieldsRemoved, fieldsRemoved)
		}

		expectedOk := true
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})
}

func TestGetAllHashMapFieldsAndValues(t *testing.T) {
	fields := make(map[string]string)
	fields["field1"] = "value123"
	fields["field2"] = "value777"
	fields["field3"] = "value915"

	t.Run("GetNonExistentKey", func(t *testing.T) {
		db := newDatabase("test")
		result, ok := db.GetAllHashMapFieldsAndValues("key1")

		if result == nil {
			t.Fatalf("expected result but got nil")
		}

		expectedElements := 0
		if len(result) != expectedElements {
			t.Errorf("expected elements = %d; got = %d", expectedElements, len(result))
		}

		expectedOk := false
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}
	})

	t.Run("GetExistingKey", func(t *testing.T) {
		db := newDatabase("test")
		key := DatabaseKey("key1")
		db.SetHashMap(key, fields, common.HashMapMaxFields)
		result, ok := db.GetAllHashMapFieldsAndValues(key)

		expectedElements := 3
		if len(result) != expectedElements {
			t.Errorf("expected elements = %d; got = %d", expectedElements, len(result))
		}

		expectedOk := true
		if ok != expectedOk {
			t.Errorf("expected ok = %v; got = %v", expectedOk, ok)
		}

		for expectedField, expectedValue := range fields {
			actualValue, exists := result[expectedField]
			if !exists {
				t.Errorf("expected field '%s' in result", expectedField)
			}
			if expectedValue != actualValue {
				t.Errorf("expected field value = %s; got = %s", expectedValue, actualValue)
			}
		}
	})
}
