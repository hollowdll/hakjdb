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
