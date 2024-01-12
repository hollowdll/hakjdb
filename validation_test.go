package kvdb

import (
	"strings"
	"testing"
)

func TestValidateDatabaseName(t *testing.T) {
	type TestCase struct {
		dbName string
		valid  bool
	}
	cases := []TestCase{
		{
			dbName: "",
			valid:  false,
		},
		{
			dbName: "   ",
			valid:  false,
		},
		{
			dbName: "!test?",
			valid:  false,
		},
		{
			dbName: "test",
			valid:  true,
		},
		{
			dbName: "{[()]};:,.",
			valid:  false,
		},
		{
			dbName: "this-is_valid",
			valid:  true,
		},
		{
			dbName: "tHis-iS_valiD0123",
			valid:  true,
		},
		{
			dbName: strings.Repeat("a", DbNameMaxSize+1),
			valid:  false,
		},
	}
	for _, test := range cases {
		err := validateDatabaseName(test.dbName)
		if test.valid && err != nil {
			t.Errorf("database name %s should be valid but is invalid", test.dbName)
		} else if !test.valid && err == nil {
			t.Errorf("database name %s should be invalid but is valid", test.dbName)
		}
	}
}

func TestValidateDatabaseKey(t *testing.T) {
	type TestCase struct {
		key   DatabaseKey
		valid bool
	}
	cases := []TestCase{
		{
			key:   "",
			valid: false,
		},
		{
			key:   "   ",
			valid: false,
		},
		{
			key:   "valid",
			valid: true,
		},
		{
			key:   "012VaLId!?-a_b_c/",
			valid: true,
		},
		{
			key:   "{[()]};:,.",
			valid: true,
		},
		{
			key:   DatabaseKey(strings.Repeat("a", DbKeyMaxSize+1)),
			valid: false,
		},
	}
	for _, test := range cases {
		err := validateDatabaseKey(test.key)
		if test.valid && err != nil {
			t.Errorf("key %s should be valid but is invalid", test.key)
		} else if !test.valid && err == nil {
			t.Errorf("key %s should be invalid but is valid", test.key)
		}
	}
}