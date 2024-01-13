package main_test

import (
	"context"
	"testing"

	main "github.com/hollowdll/kvdb/cmd/kvdbserver"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/stretchr/testify/assert"
)

// TODO: test cases
func TestCreateDatabase(t *testing.T) {
	server := main.NewServer()

	request := &kvdbserver.CreateDatabaseRequest{DbName: "test"}
	_, err := server.CreateDatabase(context.Background(), request)

	assert.NoErrorf(t, err, "error creating database: %s", err)
}
