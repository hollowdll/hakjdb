package validation

import (
	"strings"
	"testing"
)

func TestValidateDBName(t *testing.T) {
	maxDBNameSize := 64
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
			dbName: strings.Repeat("a", maxDBNameSize),
			valid:  true,
		},
		{
			dbName: strings.Repeat("a", maxDBNameSize+1),
			valid:  false,
		},
	}
	for _, test := range cases {
		err := ValidateDBName(test.dbName)
		if test.valid && err != nil {
			t.Errorf("database name %s should be valid but is invalid", test.dbName)
		} else if !test.valid && err == nil {
			t.Errorf("database name %s should be invalid but is valid", test.dbName)
		}
	}
}

func TestValidateDBKey(t *testing.T) {
	maxKeySize := 1024
	type TestCase struct {
		key   string
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
			key:   strings.Repeat("a", maxKeySize),
			valid: true,
		},
		{
			key:   strings.Repeat("a", maxKeySize+1),
			valid: false,
		},
	}
	for _, test := range cases {
		err := ValidateDBKey(test.key)
		if test.valid && err != nil {
			t.Errorf("key %s should be valid but is invalid", test.key)
		} else if !test.valid && err == nil {
			t.Errorf("key %s should be invalid but is valid", test.key)
		}
	}
}

func TestIsBlank(t *testing.T) {
	type TestCase struct {
		input    string
		expected bool
	}
	cases := []TestCase{
		{
			input:    "",
			expected: true,
		},
		{
			input:    " ",
			expected: true,
		},
		{
			input:    "  ",
			expected: true,
		},
		{
			input:    "   ",
			expected: true,
		},
		{
			input:    "           ",
			expected: true,
		},
		{
			input:    "a",
			expected: false,
		},
		{
			input:    "  a  ",
			expected: false,
		},
		{
			input:    " a b ",
			expected: false,
		},
	}
	for _, test := range cases {
		result := isBlank(test.input)
		if result != test.expected {
			t.Errorf("input = '%s'; got = %v; expected = %v", test.input, result, test.expected)
		}
	}
}

func TestIsTooLong(t *testing.T) {
	type TestCase struct {
		input       string
		targetBytes int
		expected    bool
	}
	cases := []TestCase{
		{
			input:       "abc",
			targetBytes: 2,
			expected:    true,
		},
		{
			input:       "abc",
			targetBytes: 4,
			expected:    false,
		},
		{
			input:       "abcdef",
			targetBytes: 6,
			expected:    false,
		},
	}
	for _, tc := range cases {
		result := isTooLong(tc.input, tc.targetBytes)
		if result != tc.expected {
			t.Errorf("input = '%s'; got = %v; expected = %v", tc.input, result, tc.expected)
		}
	}
}

func TestDBNameContainsValidCharacters(t *testing.T) {
	cases := []string{
		"ABC09",
		"x_y-zfhL123",
		"XYZ-abc",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_",
	}
	for _, tc := range cases {
		if !dbNameContainsValidCharacters(tc) {
			t.Errorf("database name '%s' contains invalid characters; got = %v; expected = %v", tc, false, true)
		}
	}
}
