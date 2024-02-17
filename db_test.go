package kvdb

import (
	"testing"
	"time"
)

func TestCreateDatabase(t *testing.T) {
	db, err := CreateDatabase("test")
	expectedDb := newDatabase("test")

	if err != nil {
		t.Fatal(err)
	}
	if db == nil {
		t.Fatal("expected db but got nil")
	}

	if db.GetName() != expectedDb.GetName() {
		t.Fatalf("expected db with name %s but got %s", expectedDb.GetName(), db.GetName())
	}
}

func TestDatabaseSetString(t *testing.T) {
	t.Run("SetNonExistentKey", func(t *testing.T) {
		db := newDatabase("test")
		err := db.SetString("key1", "value1")
		if err != nil {
			t.Fatal(err)
		}

		var expectedKeys uint32 = 1
		keys := db.GetKeyCount()
		if keys != expectedKeys {
			t.Errorf("expected keys = %d; got = %d", expectedKeys, keys)
		}
	})

	t.Run("OverwriteExistingKey", func(t *testing.T) {
		db := newDatabase("test")
		err := db.SetString("key1", "value1")
		if err != nil {
			t.Fatal(err)
		}

		err = db.SetString("key1", "value2")
		if err != nil {
			t.Fatalf("error overwriting key: %v", err)
		}

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
		err := db.SetString("key1", "value1")
		if err != nil {
			t.Fatal(err)
		}

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
		err := db.SetString("key1", "value1")
		if err != nil {
			t.Fatal(err)
		}

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
		err := db.SetString("key1", "value1")
		if err != nil {
			t.Fatal(err)
		}

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
		err := db.SetString(key, expectedValue)
		if err != nil {
			t.Fatal(err)
		}
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

	err := db.SetString("key1", "value1")
	if err != nil {
		t.Fatalf("error setting string value")
	}
	count = db.GetKeyCount()
	if count != 1 {
		t.Fatalf("key count should be 1 but got %d", count)
	}

	err = db.SetString("key2", "value2")
	if err != nil {
		t.Fatalf("error setting string value")
	}
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
			err := db.SetString(key, "value")
			if err != nil {
				t.Fatalf("error setting string value")
			}
		}

		count := db.GetKeyCount()
		if count != uint32(len(keys)) {
			t.Errorf("expected keys = %d; got = %d", len(keys), count)
		}

		db.DeleteAllKeys()
		count = db.GetKeyCount()
		expectedCount := 0
		if count != 0 {
			t.Errorf("expected keys = %d; got = %d", expectedCount, count)
		}
	})
}
