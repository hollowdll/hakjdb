package server_test

import (
	"context"
	"testing"

	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateDatabase(t *testing.T) {
	t.Run("DatabaseNonExistent", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "test"

		request := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		response, err := server.CreateDatabase(context.Background(), request)

		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, response, "expected response to be non-nil")
		assert.Equalf(t, dbName, response.DbName, "expected DbName = %s; got = %s", dbName, response.DbName)
	})

	t.Run("DatabaseAlreadyExists", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "test"

		request := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), request)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		response, err := server.CreateDatabase(context.Background(), request)
		require.Error(t, err, "expected error")
		require.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.AlreadyExists, st.Code(), "expected status = %s; got = %s", codes.AlreadyExists, st.Code())
	})

	t.Run("InvalidArguments", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "   "

		request := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		response, err := server.CreateDatabase(context.Background(), request)
		require.Error(t, err, "expected error")
		require.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})
}

func TestGetAllDatabases(t *testing.T) {
	t.Run("NoDatabases", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		expected := 0
		request := &kvdbserver.GetAllDatabasesRequest{}
		response, err := server.GetAllDatabases(context.Background(), request)

		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, response, "expected response to be non-nil")
		assert.Equalf(t, expected, len(response.DbNames), "expected databases = %d; got = %d", expected, len(response.DbNames))
	})

	t.Run("MultipleDatabases", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()

		dbs := []string{"db0", "db1", "db2"}
		for _, db := range dbs {
			request := &kvdbserver.CreateDatabaseRequest{DbName: db}
			_, err := server.CreateDatabase(context.Background(), request)
			require.NoErrorf(t, err, "expected no error; error = %s", err)
		}

		request := &kvdbserver.GetAllDatabasesRequest{}
		response, err := server.GetAllDatabases(context.Background(), request)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, response, "expected response to be non-nil")
		assert.Equalf(t, len(dbs), len(response.DbNames), "expected databases = %d; got = %d", len(dbs), len(response.DbNames))

		for _, db := range response.DbNames {
			assert.Equalf(t, true, common.StringInSlice(db, dbs), "expected database name %s to be in %v", db, dbs)
		}
	})
}

func TestGetDatabaseInfo(t *testing.T) {
	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"

		request := &kvdbserver.GetDatabaseInfoRequest{DbName: dbName}
		response, err := server.GetDatabaseInfo(context.Background(), request)
		require.Error(t, err, "expected error")
		require.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.NotFound, st.Code(), "expected status = %s; got = %s", codes.NotFound, st.Code())
	})

	t.Run("DatabaseExists", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"

		requestCreate := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), requestCreate)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		requestGet := &kvdbserver.GetDatabaseInfoRequest{DbName: dbName}
		response, err := server.GetDatabaseInfo(context.Background(), requestGet)
		expectedKeyCount := uint32(0)
		expectedDataSize := uint64(0)

		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, response, "expected response to be non-nil")
		assert.Equalf(t, dbName, response.Data.Name, "expected database name = %s; got = %s", dbName, response.Data.Name)
		assert.Equalf(t, expectedKeyCount, response.Data.KeyCount, "expected keys = %d; got = %d", expectedKeyCount, response.Data.KeyCount)
		assert.Equalf(t, expectedDataSize, response.Data.DataSize, "expected data size = %d; got = %d", expectedDataSize, response.Data.DataSize)
	})
}

func TestDefaultDatabase(t *testing.T) {
	t.Run("GetDatabaseInfo", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)

		request := &kvdbserver.GetDatabaseInfoRequest{DbName: dbName}
		response, err := server.GetDatabaseInfo(context.Background(), request)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, response, "expected response to be non-nil")
		assert.Equalf(t, dbName, response.Data.Name, "expected database name = %s; got = %s", dbName, response.Data.Name)
	})

	t.Run("GetAllDatabases", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)

		request := &kvdbserver.GetAllDatabasesRequest{}
		response, err := server.GetAllDatabases(context.Background(), request)
		expectedDbCount := 1

		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, response, "expected response to be non-nil")
		assert.Equalf(t, expectedDbCount, len(response.DbNames), "expected databases = %d; got = %d", expectedDbCount, len(response.DbNames))
	})
}

func TestDeleteDatabase(t *testing.T) {
	t.Run("DatabaseExists", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)

		req := &kvdbserver.DeleteDatabaseRequest{DbName: dbName}
		res, err := server.DeleteDatabase(context.Background(), req)

		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res, "expected response to be non-nil")
		assert.Equalf(t, dbName, res.DbName, "expected db name = &s; got = %s", dbName, res.DbName)
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"

		req := &kvdbserver.DeleteDatabaseRequest{DbName: dbName}
		res, err := server.DeleteDatabase(context.Background(), req)
		require.Error(t, err, "expected error")
		require.Nil(t, res, "expected response to be nil")

		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.NotFound, st.Code(), "expected status = %s; got = %s", codes.NotFound, st.Code())
	})
}
