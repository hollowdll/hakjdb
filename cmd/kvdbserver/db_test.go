package main_test

import (
	"context"
	"testing"

	main "github.com/hollowdll/kvdb/cmd/kvdbserver"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateDatabase(t *testing.T) {
	t.Run("DatabaseNonExistent", func(t *testing.T) {
		server := main.NewServer()
		dbName := "test"

		request := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		response, err := server.CreateDatabase(context.Background(), request)

		assert.NoErrorf(t, err, "expected no error; error = %s", err)
		assert.NotNil(t, response, "expected response to be non-nil")
		assert.Equalf(t, dbName, response.DbName, "expected DbName = %s; got = %s", dbName, response.DbName)
	})

	t.Run("DatabaseAlreadyExists", func(t *testing.T) {
		server := main.NewServer()
		dbName := "test"

		request := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), request)
		assert.NoErrorf(t, err, "expected no error; error = %s", err)
		response, err := server.CreateDatabase(context.Background(), request)

		assert.Error(t, err, "expected error")
		assert.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		assert.NotNil(t, st, "expected status to be non-nil")
		assert.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.AlreadyExists, st.Code(), "expected status = %s; got = %s", codes.AlreadyExists, st.Code())
	})

	t.Run("InvalidArguments", func(t *testing.T) {
		server := main.NewServer()
		dbName := "   "

		request := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		response, err := server.CreateDatabase(context.Background(), request)

		assert.Error(t, err, "expected no error")
		assert.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		assert.NotNil(t, st, "expected status to be non-nil")
		assert.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})
}
