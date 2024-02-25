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
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestSetString(t *testing.T) {
	t.Run("MissingMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()

		request := &kvdbserver.SetStringRequest{Key: "key1", Value: "value1"}
		response, err := server.SetString(context.Background(), request)
		require.Error(t, err, "expected error")
		require.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})

	t.Run("MissingDatabaseInMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs("wrong-key", dbName))

		request := &kvdbserver.SetStringRequest{Key: "key1", Value: "value1"}
		response, err := server.SetString(ctxMd, request)
		require.Error(t, err, "expected error")
		require.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())

	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		request := &kvdbserver.SetStringRequest{Key: "key1", Value: "value1"}
		response, err := server.SetString(ctxMd, request)
		require.Error(t, err, "expected error")
		require.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.NotFound, st.Code(), "expected status = %s; got = %s", codes.NotFound, st.Code())
	})

	t.Run("Success", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		requestCreate := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), requestCreate)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		request := &kvdbserver.SetStringRequest{Key: "key1", Value: "value1"}
		response, err := server.SetString(ctxMd, request)
		assert.NoErrorf(t, err, "expected no error; error = %s", err)
		assert.NotNil(t, response, "expected response to be non-nil")
	})

	t.Run("InvalidInput", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		requestCreate := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), requestCreate)
		assert.NoErrorf(t, err, "expected no error; error = %s", err)

		request := &kvdbserver.SetStringRequest{Key: "      ", Value: "value1"}
		response, err := server.SetString(ctxMd, request)
		require.Error(t, err, "expected error")
		require.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})
}

func TestGetString(t *testing.T) {
	t.Run("MissingMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()

		request := &kvdbserver.GetStringRequest{Key: "key1"}
		response, err := server.GetString(context.Background(), request)
		require.Error(t, err, "expected error")
		require.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})

	t.Run("MissingDatabaseInMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs("wrong-key", dbName))

		request := &kvdbserver.GetStringRequest{Key: "key1"}
		response, err := server.GetString(ctxMd, request)
		require.Error(t, err, "expected error")
		require.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		request := &kvdbserver.GetStringRequest{Key: "key1"}
		response, err := server.GetString(ctxMd, request)
		require.Error(t, err, "expected error")
		require.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.NotFound, st.Code(), "expected status = %s; got = %s", codes.NotFound, st.Code())
	})

	t.Run("SuccessKeyNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		requestCreate := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), requestCreate)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		request := &kvdbserver.GetStringRequest{Key: "key1"}
		response, err := server.GetString(ctxMd, request)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, response, "expected response to be non-nil")
		assert.Equalf(t, false, response.Found, "expected found = %v; got = %v", false, response.Found)
		assert.Equalf(t, "", response.Value, "expected empty string; got = %s", response.Value)
	})

	t.Run("SuccessKeyFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		expectedValue := "value1"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		requestCreate := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), requestCreate)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		requestSet := &kvdbserver.SetStringRequest{Key: "key1", Value: expectedValue}
		_, err = server.SetString(ctxMd, requestSet)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		requestGet := &kvdbserver.GetStringRequest{Key: "key1"}
		response, err := server.GetString(ctxMd, requestGet)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, response, "expected response to be non-nil")
		assert.Equalf(t, true, response.Found, "expected found = %v; got = %v", true, response.Found)
		assert.Equalf(t, expectedValue, response.Value, "expected value = %s; got = %s", expectedValue, response.Value)
	})
}

func TestDeleteKey(t *testing.T) {
	t.Run("MissingMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()

		request := &kvdbserver.DeleteKeyRequest{Key: "key1"}
		response, err := server.DeleteKey(context.Background(), request)
		require.Error(t, err, "expected error")
		require.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})

	t.Run("MissingDatabaseInMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs("wrong-key", dbName))

		request := &kvdbserver.DeleteKeyRequest{Key: "key1"}
		response, err := server.DeleteKey(ctxMd, request)
		require.Error(t, err, "expected error")
		require.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		request := &kvdbserver.DeleteKeyRequest{Key: "key1"}
		response, err := server.DeleteKey(ctxMd, request)
		require.Error(t, err, "expected error")
		require.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.NotFound, st.Code(), "expected status = %s; got = %s", codes.NotFound, st.Code())
	})

	t.Run("SuccessKeyNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		requestCreate := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), requestCreate)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		request := &kvdbserver.DeleteKeyRequest{Key: "key1"}
		response, err := server.DeleteKey(ctxMd, request)
		expected := false
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, response, "expected response to be non-nil")
		assert.Equalf(t, expected, response.Ok, "expected ok = %v; got = %v", expected, response.Ok)
	})

	t.Run("SuccessKeyFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		requestCreate := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), requestCreate)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		requestSet := &kvdbserver.SetStringRequest{Key: "key1", Value: "v"}
		_, err = server.SetString(ctxMd, requestSet)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		requestGet := &kvdbserver.DeleteKeyRequest{Key: "key1"}
		response, err := server.DeleteKey(ctxMd, requestGet)
		expected := true
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, response, "expected response to be non-nil")
		assert.Equalf(t, expected, response.Ok, "expected ok = %v; got = %v", expected, response.Ok)
	})
}

func TestDeleteAllKeys(t *testing.T) {
	t.Run("MissingMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()

		req := &kvdbserver.DeleteAllKeysRequest{}
		res, err := server.DeleteAllKeys(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, res)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})

	t.Run("MissingDatabaseInMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("wrong-key", dbName))

		req := &kvdbserver.DeleteAllKeysRequest{}
		res, err := server.DeleteAllKeys(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserver.DeleteAllKeysRequest{}
		res, err := server.DeleteAllKeys(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.NotFound, st.Code(), "expected status = %s; got = %s", codes.NotFound, st.Code())
	})

	t.Run("NoKeysPresent", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqCreate := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		server.CreateDatabase(context.Background(), reqCreate)

		req := &kvdbserver.DeleteAllKeysRequest{}
		response, err := server.DeleteAllKeys(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, response)
	})

	t.Run("KeysPresent", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqCreate := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		server.CreateDatabase(context.Background(), reqCreate)

		reqSet := &kvdbserver.SetStringRequest{Key: "key1", Value: "v"}
		server.SetString(ctx, reqSet)

		req := &kvdbserver.DeleteAllKeysRequest{}
		response, err := server.DeleteAllKeys(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, response)
	})
}

func TestGetKeys(t *testing.T) {
	t.Run("MissingMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()

		req := &kvdbserver.GetKeysRequest{}
		res, err := server.GetKeys(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, res)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})

	t.Run("MissingDatabaseInMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("wrong-key", dbName))

		req := &kvdbserver.GetKeysRequest{}
		res, err := server.GetKeys(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserver.GetKeysRequest{}
		res, err := server.GetKeys(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.NotFound, st.Code(), "expected status = %s; got = %s", codes.NotFound, st.Code())
	})

	t.Run("NoKeysPresent", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqCreate := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		server.CreateDatabase(context.Background(), reqCreate)

		req := &kvdbserver.GetKeysRequest{}
		res, err := server.GetKeys(ctx, req)
		expectedKeys := 0
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equal(t, expectedKeys, len(res.Keys), "expected keys = %d; got = %d", expectedKeys, len(res.Keys))
	})

	t.Run("KeysPresent", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqCreate := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		server.CreateDatabase(context.Background(), reqCreate)

		keys := []string{"key1", "key2", "key3"}
		for _, key := range keys {
			req := &kvdbserver.SetStringRequest{Key: key}
			_, err := server.SetString(ctx, req)
			require.NoErrorf(t, err, "expected no error; error = %v", err)
		}

		req := &kvdbserver.GetKeysRequest{}
		res, err := server.GetKeys(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, res)
		assert.Equalf(t, len(keys), len(res.Keys), "expected keys = %d; got = %d", len(keys), len(res.Keys))

		for _, key := range res.Keys {
			assert.Equalf(t, true, common.StringInSlice(key, keys), "expected key %s to be in %v", key, keys)
		}
	})
}
