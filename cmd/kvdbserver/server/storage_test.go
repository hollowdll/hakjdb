package server_test

import (
	"context"
	"testing"

	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserverpb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestGetTypeOfKey(t *testing.T) {
	t.Run("MetadataNotSent", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		server.CreateDefaultDatabase("default")

		req := &kvdbserverpb.GetTypeOfKeyRequest{Key: "key1"}
		res, err := server.GetTypeOfKey(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotInMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("wrong-key", dbName))

		req := &kvdbserverpb.GetTypeOfKeyRequest{Key: "key1"}
		res, err := server.GetTypeOfKey(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserverpb.GetTypeOfKeyRequest{Key: "key1"}
		res, err := server.GetTypeOfKey(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		expectedOk := true
		expectedCode := codes.NotFound
		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equalf(t, expectedOk, ok, "expected ok = %v; got = %v", expectedOk, ok)
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})

	t.Run("KeyNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		expectedKeyType := ""
		expectedOk := false
		req := &kvdbserverpb.GetTypeOfKeyRequest{Key: "key1"}
		res, err := server.GetTypeOfKey(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedKeyType, res.KeyType, "expected key type = %s; got = %s", expectedKeyType, res.KeyType)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})

	t.Run("String", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvdbserverpb.SetStringRequest{Key: "key1", Value: "value1"}
		server.SetString(ctx, reqSet)

		expectedKeyType := "String"
		expectedOk := true
		req := &kvdbserverpb.GetTypeOfKeyRequest{Key: "key1"}
		res, err := server.GetTypeOfKey(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedKeyType, res.KeyType, "expected key type = %s; got = %s", expectedKeyType, res.KeyType)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})

	t.Run("HashMap", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: make(map[string]string)}
		server.SetHashMap(ctx, reqSet)

		expectedKeyType := "HashMap"
		expectedOk := true
		req := &kvdbserverpb.GetTypeOfKeyRequest{Key: "key1"}
		res, err := server.GetTypeOfKey(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedKeyType, res.KeyType, "expected key type = %s; got = %s", expectedKeyType, res.KeyType)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})
}

func TestSetString(t *testing.T) {
	t.Run("MetadataNotSent", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		server.CreateDefaultDatabase("default")

		req := &kvdbserverpb.SetStringRequest{Key: "key1", Value: "value1"}
		res, err := server.SetString(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotInMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("wrong-key", dbName))

		req := &kvdbserverpb.SetStringRequest{Key: "key1", Value: "value1"}
		res, err := server.SetString(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		request := &kvdbserverpb.SetStringRequest{Key: "key1", Value: "value1"}
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

		requestCreate := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), requestCreate)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		request := &kvdbserverpb.SetStringRequest{Key: "key1", Value: "value1"}
		response, err := server.SetString(ctxMd, request)
		assert.NoErrorf(t, err, "expected no error; error = %s", err)
		assert.NotNil(t, response, "expected response to be non-nil")
	})

	t.Run("InvalidKey", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		requestCreate := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), requestCreate)
		assert.NoErrorf(t, err, "expected no error; error = %s", err)

		request := &kvdbserverpb.SetStringRequest{Key: "      ", Value: "value1"}
		response, err := server.SetString(ctxMd, request)
		require.Error(t, err, "expected error")
		require.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})

	t.Run("MaxKeyLimitReached", func(t *testing.T) {
		server := server.NewServerWithOptions(&server.ServerOptions{MaxKeysPerDb: 1})
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserverpb.SetStringRequest{Key: "key1", Value: "value1"}
		res, err := server.SetString(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)

		req = &kvdbserverpb.SetStringRequest{Key: "key2", Value: "value1"}
		res, err = server.SetString(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		expectedOk := true
		expectedCode := codes.FailedPrecondition
		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equalf(t, expectedOk, ok, "expected ok = %v; got = %v", expectedOk, ok)
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})
}

func TestGetString(t *testing.T) {
	t.Run("MetadataNotSent", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		server.CreateDefaultDatabase("default")

		req := &kvdbserverpb.GetStringRequest{Key: "key1"}
		res, err := server.GetString(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotInMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("wrong-key", dbName))

		req := &kvdbserverpb.GetStringRequest{Key: "key1"}
		res, err := server.GetString(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		request := &kvdbserverpb.GetStringRequest{Key: "key1"}
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

		requestCreate := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), requestCreate)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		request := &kvdbserverpb.GetStringRequest{Key: "key1"}
		response, err := server.GetString(ctxMd, request)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, response, "expected response to be non-nil")
		assert.Equalf(t, false, response.Ok, "expected ok = %v; got = %v", false, response.Ok)
		assert.Equalf(t, "", response.Value, "expected empty string; got = %s", response.Value)
	})

	t.Run("SuccessKeyFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		expectedValue := "value1"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		requestCreate := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), requestCreate)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		requestSet := &kvdbserverpb.SetStringRequest{Key: "key1", Value: expectedValue}
		_, err = server.SetString(ctxMd, requestSet)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		requestGet := &kvdbserverpb.GetStringRequest{Key: "key1"}
		response, err := server.GetString(ctxMd, requestGet)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, response, "expected response to be non-nil")
		assert.Equalf(t, true, response.Ok, "expected ok = %v; got = %v", true, response.Ok)
		assert.Equalf(t, expectedValue, response.Value, "expected value = %s; got = %s", expectedValue, response.Value)
	})
}

func TestDeleteKey(t *testing.T) {
	t.Run("MetadataNotSent", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		server.CreateDefaultDatabase("default")

		req := &kvdbserverpb.DeleteKeyRequest{Keys: []string{"key1"}}
		res, err := server.DeleteKey(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotInMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("wrong-key", dbName))

		req := &kvdbserverpb.DeleteKeyRequest{Keys: []string{"key1"}}
		res, err := server.DeleteKey(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserverpb.DeleteKeyRequest{Keys: []string{"key1"}}
		res, err := server.DeleteKey(ctxMd, req)
		require.Error(t, err, "expected error")
		require.Nil(t, res)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.NotFound, st.Code(), "expected status = %s; got = %s", codes.NotFound, st.Code())
	})

	t.Run("SuccessKeyNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqCreate := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), reqCreate)
		require.NoErrorf(t, err, "expected no error; error = %v", err)

		req := &kvdbserverpb.DeleteKeyRequest{Keys: []string{"key1"}}
		res, err := server.DeleteKey(ctxMd, req)
		var expected uint32 = 0
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expected, res.KeysDeleted, "expected keys deleted = %d; got = %d", expected, res.KeysDeleted)
	})

	t.Run("SuccessKeyFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqCreate := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), reqCreate)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		reqSet := &kvdbserverpb.SetStringRequest{Key: "key1", Value: "v"}
		_, err = server.SetString(ctxMd, reqSet)
		require.NoErrorf(t, err, "expected no error; error = %v", err)

		reqGet := &kvdbserverpb.DeleteKeyRequest{Keys: []string{"key1", "key2", "key3"}}
		res, err := server.DeleteKey(ctxMd, reqGet)
		var expected uint32 = 1
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expected, res.KeysDeleted, "expected keys deleted = %d; got = %d", expected, res.KeysDeleted)
	})
}

func TestDeleteAllKeys(t *testing.T) {
	t.Run("MetadataNotSent", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		server.CreateDefaultDatabase("default")

		req := &kvdbserverpb.DeleteAllKeysRequest{}
		res, err := server.DeleteAllKeys(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotInMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("wrong-key", dbName))

		req := &kvdbserverpb.DeleteAllKeysRequest{}
		res, err := server.DeleteAllKeys(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserverpb.DeleteAllKeysRequest{}
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

		reqCreate := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
		server.CreateDatabase(context.Background(), reqCreate)

		req := &kvdbserverpb.DeleteAllKeysRequest{}
		response, err := server.DeleteAllKeys(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, response)
	})

	t.Run("KeysPresent", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqCreate := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
		server.CreateDatabase(context.Background(), reqCreate)

		reqSet := &kvdbserverpb.SetStringRequest{Key: "key1", Value: "v"}
		server.SetString(ctx, reqSet)

		req := &kvdbserverpb.DeleteAllKeysRequest{}
		response, err := server.DeleteAllKeys(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, response)
	})
}

func TestGetKeys(t *testing.T) {
	t.Run("MetadataNotSent", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		server.CreateDefaultDatabase("default")

		req := &kvdbserverpb.GetKeysRequest{}
		res, err := server.GetKeys(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotInMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("wrong-key", dbName))

		req := &kvdbserverpb.GetKeysRequest{}
		res, err := server.GetKeys(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserverpb.GetKeysRequest{}
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

		reqCreate := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
		server.CreateDatabase(context.Background(), reqCreate)

		req := &kvdbserverpb.GetKeysRequest{}
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

		reqCreate := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
		server.CreateDatabase(context.Background(), reqCreate)

		keys := []string{"key1", "key2", "key3"}
		for _, key := range keys {
			req := &kvdbserverpb.SetStringRequest{Key: key}
			_, err := server.SetString(ctx, req)
			require.NoErrorf(t, err, "expected no error; error = %v", err)
		}

		req := &kvdbserverpb.GetKeysRequest{}
		res, err := server.GetKeys(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, res)
		assert.Equalf(t, len(keys), len(res.Keys), "expected keys = %d; got = %d", len(keys), len(res.Keys))

		for _, key := range res.Keys {
			assert.Equalf(t, true, common.StringInSlice(key, keys), "expected key %s to be in %v", key, keys)
		}
	})
}

func TestSetHashMap(t *testing.T) {
	fields := make(map[string]string)
	fields["field1"] = "value1"
	fields["field2"] = "value2"
	fields["field3"] = "value3"

	t.Run("MetadataNotSent", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		server.CreateDefaultDatabase("default")

		req := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: fields}
		res, err := server.SetHashMap(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotInMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("wrong-key", dbName))

		req := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: fields}
		res, err := server.SetHashMap(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: fields}
		res, err := server.SetHashMap(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		expectedOk := true
		expectedCode := codes.NotFound
		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equalf(t, expectedOk, ok, "expected ok = %v; got = %v", expectedOk, ok)
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})

	t.Run("FieldsAdded", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: fields}
		res, err := server.SetHashMap(ctx, req)

		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		var expectedFieldsAdded uint32 = 3
		assert.Equal(t, expectedFieldsAdded, res.FieldsAdded, "expected fields added = %d; got = %d", expectedFieldsAdded, res.FieldsAdded)
	})

	t.Run("OverwriteFields", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))
		fieldsOverwrite := make(map[string]string)
		fieldsOverwrite["field1"] = "a"
		fieldsOverwrite["field2"] = "b"
		fieldsOverwrite["field3"] = "c"
		fieldsOverwrite["new_field"] = "d"

		req1 := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: fields}
		res, err := server.SetHashMap(ctx, req1)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		var expectedFieldsAdded1 uint32 = 3
		assert.Equal(t, expectedFieldsAdded1, res.FieldsAdded, "expected fields added = %d; got = %d", expectedFieldsAdded1, res.FieldsAdded)

		req2 := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: fieldsOverwrite}
		res, err = server.SetHashMap(ctx, req2)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		var expectedFieldsAdded2 uint32 = 1
		assert.Equal(t, expectedFieldsAdded2, res.FieldsAdded, "expected fields added = %d; got = %d", expectedFieldsAdded2, res.FieldsAdded)
	})

	t.Run("InvalidKey", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctxMd := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvdbserverpb.SetHashMapRequest{Key: "    ", Fields: fields}
		res, err := server.SetHashMap(ctxMd, reqSet)
		require.Error(t, err)
		require.Nil(t, res)

		expectedOk := true
		expectedCode := codes.InvalidArgument
		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equalf(t, expectedOk, ok, "expected ok = %v; got = %v", expectedOk, ok)
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})

	t.Run("MaxKeyLimitReached", func(t *testing.T) {
		server := server.NewServerWithOptions(&server.ServerOptions{MaxKeysPerDb: 1})
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: fields}
		res, err := server.SetHashMap(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)

		req = &kvdbserverpb.SetHashMapRequest{Key: "key2", Fields: fields}
		res, err = server.SetHashMap(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		expectedOk := true
		expectedCode := codes.FailedPrecondition
		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equalf(t, expectedOk, ok, "expected ok = %v; got = %v", expectedOk, ok)
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})

	t.Run("MaxFieldLimitReached", func(t *testing.T) {
		server := server.NewServerWithOptions(&server.ServerOptions{MaxHashMapFields: 4})
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))
		fields2 := make(map[string]string)
		fields2["field4"] = "val4"
		fields2["field5"] = "val5"
		fields2["field6"] = "val6"
		fields2["field7"] = "val7"

		req1 := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: fields}
		res, err := server.SetHashMap(ctx, req1)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		var expectedFieldsAdded1 uint32 = 3
		assert.Equal(t, expectedFieldsAdded1, res.FieldsAdded, "expected fields added = %d; got = %d", expectedFieldsAdded1, res.FieldsAdded)

		req2 := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: fields2}
		res, err = server.SetHashMap(ctx, req2)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		var expectedFieldsAdded2 uint32 = 1
		assert.Equal(t, expectedFieldsAdded2, res.FieldsAdded, "expected fields added = %d; got = %d", expectedFieldsAdded2, res.FieldsAdded)
	})
}

func TestGetHashMapFieldValue(t *testing.T) {
	fields := make(map[string]string)
	fields["field1"] = "value1"
	fields["field2"] = "value2"
	fields["field3"] = "value3"

	t.Run("MetadataNotSent", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		server.CreateDefaultDatabase("default")

		req := &kvdbserverpb.GetHashMapFieldValueRequest{Key: "key1", Fields: []string{"field2"}}
		res, err := server.GetHashMapFieldValue(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotInMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("wrong-key", dbName))

		req := &kvdbserverpb.GetHashMapFieldValueRequest{Key: "key1", Fields: []string{"field2"}}
		res, err := server.GetHashMapFieldValue(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserverpb.GetHashMapFieldValueRequest{Key: "key1", Fields: []string{"field2"}}
		res, err := server.GetHashMapFieldValue(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		expectedOk := true
		expectedCode := codes.NotFound
		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equalf(t, expectedOk, ok, "expected ok = %v; got = %v", expectedOk, ok)
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})

	t.Run("KeyAndFieldFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: fields}
		server.SetHashMap(ctx, reqSet)

		expectedValue := "value2"
		expectedOk := true
		expectedKeyFound := true
		reqGet := &kvdbserverpb.GetHashMapFieldValueRequest{Key: "key1", Fields: []string{"field2"}}
		res, err := server.GetHashMapFieldValue(ctx, reqGet)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedValue, res.FieldValueMap["field2"].Value, "expected value = %s; got = %s", expectedValue, res.FieldValueMap["field2"].Value)
		assert.Equalf(t, expectedOk, res.FieldValueMap["field2"].Ok, "expected ok = %v; got = %v", expectedOk, res.FieldValueMap["field2"].Ok)
		assert.Equalf(t, expectedKeyFound, res.Ok, "expected ok = %v; got = %v", expectedKeyFound, res.Ok)
	})

	t.Run("MultipleFieldsFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: fields}
		server.SetHashMap(ctx, reqSet)

		reqGet := &kvdbserverpb.GetHashMapFieldValueRequest{Key: "key1", Fields: []string{"field1", "field2", "field3"}}
		res, err := server.GetHashMapFieldValue(ctx, reqGet)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)

		for field, expectedValue := range fields {
			assert.Equalf(t, expectedValue, res.FieldValueMap[field].Value, "expected value = %s; got = %s", expectedValue, res.FieldValueMap[field].Value)
			assert.Equalf(t, true, res.FieldValueMap[field].Ok, "expected ok = %v; got = %v", true, res.FieldValueMap[field].Ok)
		}
	})

	t.Run("KeyNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		expectedOk := false
		req := &kvdbserverpb.GetHashMapFieldValueRequest{Key: "key2", Fields: []string{"field3"}}
		res, err := server.GetHashMapFieldValue(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})

	t.Run("FieldNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: fields}
		server.SetHashMap(ctx, reqSet)

		expectedValue := ""
		expectedOk := false
		req := &kvdbserverpb.GetHashMapFieldValueRequest{Key: "key1", Fields: []string{"field123"}}
		res, err := server.GetHashMapFieldValue(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedValue, res.FieldValueMap["field123"].Value, "expected value = %s; got = %s", expectedValue, res.FieldValueMap["field123"].Value)
		assert.Equalf(t, expectedOk, res.FieldValueMap["field123"].Ok, "expected ok = %v; got = %v", expectedOk, res.FieldValueMap["field123"].Ok)
	})
}

func TestDeleteHashMapFields(t *testing.T) {
	fields := make(map[string]string)
	fields["field1"] = "value1"
	fields["field2"] = "value2"
	fields["field3"] = "value3"
	fieldsToRemove := []string{"field2", "field3"}

	t.Run("MetadataNotSent", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		server.CreateDefaultDatabase("default")

		req := &kvdbserverpb.DeleteHashMapFieldsRequest{Key: "key1", Fields: fieldsToRemove}
		res, err := server.DeleteHashMapFields(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotInMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("wrong-key", dbName))

		req := &kvdbserverpb.DeleteHashMapFieldsRequest{Key: "key1", Fields: fieldsToRemove}
		res, err := server.DeleteHashMapFields(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserverpb.DeleteHashMapFieldsRequest{Key: "key1", Fields: fieldsToRemove}
		res, err := server.DeleteHashMapFields(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		expectedOk := true
		expectedCode := codes.NotFound
		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equalf(t, expectedOk, ok, "expected ok = %v; got = %v", expectedOk, ok)
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})

	t.Run("KeyNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		var expectedFieldsRemoved uint32 = 0
		expectedOk := false
		req := &kvdbserverpb.DeleteHashMapFieldsRequest{Key: "key1", Fields: fieldsToRemove}
		res, err := server.DeleteHashMapFields(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedFieldsRemoved, res.FieldsRemoved, "expected fields removed = %d; got = %d", expectedFieldsRemoved, res.FieldsRemoved)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})

	t.Run("FieldsNotExist", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: make(map[string]string)}
		server.SetHashMap(ctx, reqSet)

		var expectedFieldsRemoved uint32 = 0
		expectedOk := true
		req := &kvdbserverpb.DeleteHashMapFieldsRequest{Key: "key1", Fields: fieldsToRemove}
		res, err := server.DeleteHashMapFields(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedFieldsRemoved, res.FieldsRemoved, "expected fields removed = %d; got = %d", expectedFieldsRemoved, res.FieldsRemoved)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})

	t.Run("FieldsExist", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: fields}
		server.SetHashMap(ctx, reqSet)

		var expectedFieldsRemoved uint32 = 2
		expectedOk := true
		req := &kvdbserverpb.DeleteHashMapFieldsRequest{Key: "key1", Fields: fieldsToRemove}
		res, err := server.DeleteHashMapFields(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedFieldsRemoved, res.FieldsRemoved, "expected fields removed = %d; got = %d", expectedFieldsRemoved, res.FieldsRemoved)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})

	t.Run("DuplicateFields", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: fields}
		server.SetHashMap(ctx, reqSet)

		var expectedFieldsRemoved uint32 = 1
		expectedOk := true
		req := &kvdbserverpb.DeleteHashMapFieldsRequest{Key: "key1", Fields: []string{"field3", "field3", "field3"}}
		res, err := server.DeleteHashMapFields(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedFieldsRemoved, res.FieldsRemoved, "expected fields removed = %d; got = %d", expectedFieldsRemoved, res.FieldsRemoved)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})
}

func TestGetAllHashMapFieldsAndValues(t *testing.T) {
	fields := make(map[string]string)
	fields["field1"] = "value1"
	fields["field2"] = "value2"
	fields["field3"] = "value3"

	t.Run("MetadataNotSent", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		server.CreateDefaultDatabase("default")

		req := &kvdbserverpb.GetAllHashMapFieldsAndValuesRequest{Key: "key1"}
		res, err := server.GetAllHashMapFieldsAndValues(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotInMetadata", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("wrong-key", dbName))

		req := &kvdbserverpb.GetAllHashMapFieldsAndValuesRequest{Key: "key1"}
		res, err := server.GetAllHashMapFieldsAndValues(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserverpb.GetAllHashMapFieldsAndValuesRequest{Key: "key1"}
		res, err := server.GetAllHashMapFieldsAndValues(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		expectedOk := true
		expectedCode := codes.NotFound
		st, ok := status.FromError(err)
		require.NotNil(t, st, "expected status to be non-nil")
		require.Equalf(t, expectedOk, ok, "expected ok = %v; got = %v", expectedOk, ok)
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})

	t.Run("KeyNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		expectedElements := 0
		expectedOk := false
		req := &kvdbserverpb.GetAllHashMapFieldsAndValuesRequest{Key: "key1"}
		res, err := server.GetAllHashMapFieldsAndValues(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		require.NotNil(t, res.FieldValueMap)
		assert.Equalf(t, expectedElements, len(res.FieldValueMap), "expected elements = %d; got = %d", expectedElements, len(res.FieldValueMap))
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})

	t.Run("KeyFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: fields}
		server.SetHashMap(ctx, reqSet)

		expectedElements := 3
		expectedOk := true
		req := &kvdbserverpb.GetAllHashMapFieldsAndValuesRequest{Key: "key1"}
		res, err := server.GetAllHashMapFieldsAndValues(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		require.NotNil(t, res.FieldValueMap)
		assert.Equalf(t, expectedElements, len(res.FieldValueMap), "expected elements = %d; got = %d", expectedElements, len(res.FieldValueMap))
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})
}
