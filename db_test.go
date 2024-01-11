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
