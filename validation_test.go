package kvdb

import "testing"

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
			dbName: "This_should_be_too_long_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
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
