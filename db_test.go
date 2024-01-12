package kvdb

import "testing"

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

// TODO: test case for overwriting key
func TestDatabaseSetString(t *testing.T) {
	db := newDatabase("test")
	err := db.SetString("key1", "value1")
	if err != nil {
		t.Fatalf("error setting string value: %v", err)
	}
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
}
