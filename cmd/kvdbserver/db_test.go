package main_test

import (
	"context"
	"testing"

	main "github.com/hollowdll/kvdb/cmd/kvdbserver"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
)

func TestCreateDatabase(t *testing.T) {
	server := main.NewServer()

	request := &kvdbserver.CreateDatabaseRequest{DbName: "test"}
	_, err := server.CreateDatabase(context.Background(), request)

	if err != nil {
		t.Error(err)
	}
}
