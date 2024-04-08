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

func TestGetTypeOfKey(t *testing.T) {
	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserver.GetTypeOfKeyRequest{Key: "key1"}
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
		req := &kvdbserver.GetTypeOfKeyRequest{Key: "key1"}
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

		reqSet := &kvdbserver.SetStringRequest{Key: "key1", Value: "value1"}
		server.SetString(ctx, reqSet)

		expectedKeyType := "String"
		expectedOk := true
		req := &kvdbserver.GetTypeOfKeyRequest{Key: "key1"}
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

		reqSet := &kvdbserver.SetHashMapRequest{Key: "key1", Fields: make(map[string]string)}
		server.SetHashMap(ctx, reqSet)

		expectedKeyType := "HashMap"
		expectedOk := true
		req := &kvdbserver.GetTypeOfKeyRequest{Key: "key1"}
		res, err := server.GetTypeOfKey(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedKeyType, res.KeyType, "expected key type = %s; got = %s", expectedKeyType, res.KeyType)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})
}

func TestSetString(t *testing.T) {
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

	t.Run("InvalidKey", func(t *testing.T) {
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

	t.Run("MaxKeyLimitReached", func(t *testing.T) {
		server := server.NewServerWithOptions(&server.ServerOptions{MaxKeysPerDb: 1})
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserver.SetStringRequest{Key: "key1", Value: "value1"}
		res, err := server.SetString(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)

		req = &kvdbserver.SetStringRequest{Key: "key2", Value: "value1"}
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
		assert.Equalf(t, false, response.Ok, "expected ok = %v; got = %v", false, response.Ok)
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
		assert.Equalf(t, true, response.Ok, "expected ok = %v; got = %v", true, response.Ok)
		assert.Equalf(t, expectedValue, response.Value, "expected value = %s; got = %s", expectedValue, response.Value)
	})
}

func TestDeleteKey(t *testing.T) {
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

func TestSetHashMap(t *testing.T) {
	fields := make(map[string]string)
	fields["field1"] = "value1"
	fields["field2"] = "value2"
	fields["field3"] = "value3"

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserver.SetHashMapRequest{Key: "key1", Fields: fields}
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

		req := &kvdbserver.SetHashMapRequest{Key: "key1", Fields: fields}
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

		req1 := &kvdbserver.SetHashMapRequest{Key: "key1", Fields: fields}
		res, err := server.SetHashMap(ctx, req1)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		var expectedFieldsAdded1 uint32 = 3
		assert.Equal(t, expectedFieldsAdded1, res.FieldsAdded, "expected fields added = %d; got = %d", expectedFieldsAdded1, res.FieldsAdded)

		req2 := &kvdbserver.SetHashMapRequest{Key: "key1", Fields: fieldsOverwrite}
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

		reqSet := &kvdbserver.SetHashMapRequest{Key: "    ", Fields: fields}
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

		req := &kvdbserver.SetHashMapRequest{Key: "key1", Fields: fields}
		res, err := server.SetHashMap(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)

		req = &kvdbserver.SetHashMapRequest{Key: "key2", Fields: fields}
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

		req1 := &kvdbserver.SetHashMapRequest{Key: "key1", Fields: fields}
		res, err := server.SetHashMap(ctx, req1)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		var expectedFieldsAdded1 uint32 = 3
		assert.Equal(t, expectedFieldsAdded1, res.FieldsAdded, "expected fields added = %d; got = %d", expectedFieldsAdded1, res.FieldsAdded)

		req2 := &kvdbserver.SetHashMapRequest{Key: "key1", Fields: fields2}
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

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserver.GetHashMapFieldValueRequest{Key: "key1", Field: "field2"}
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

		reqSet := &kvdbserver.SetHashMapRequest{Key: "key1", Fields: fields}
		server.SetHashMap(ctx, reqSet)

		expectedValue := "value2"
		expectedOk := true
		reqGet := &kvdbserver.GetHashMapFieldValueRequest{Key: "key1", Field: "field2"}
		res, err := server.GetHashMapFieldValue(ctx, reqGet)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedValue, res.Value, "expected value = %s; got = %s", expectedValue, res.Value)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})

	t.Run("KeyNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		expectedValue := ""
		expectedOk := false
		req := &kvdbserver.GetHashMapFieldValueRequest{Key: "key2", Field: "field2"}
		res, err := server.GetHashMapFieldValue(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedValue, res.Value, "expected value = %s; got = %s", expectedValue, res.Value)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})

	t.Run("FieldNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "default"
		server.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvdbserver.SetHashMapRequest{Key: "key1", Fields: fields}
		server.SetHashMap(ctx, reqSet)

		expectedValue := ""
		expectedOk := false
		req := &kvdbserver.GetHashMapFieldValueRequest{Key: "key1", Field: "field123"}
		res, err := server.GetHashMapFieldValue(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedValue, res.Value, "expected value = %s; got = %s", expectedValue, res.Value)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})
}

func TestDeleteHashMapFields(t *testing.T) {
	fields := make(map[string]string)
	fields["field1"] = "value1"
	fields["field2"] = "value2"
	fields["field3"] = "value3"
	fieldsToRemove := []string{"field2", "field3"}

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserver.DeleteHashMapFieldsRequest{Key: "key1", Fields: fieldsToRemove}
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
		req := &kvdbserver.DeleteHashMapFieldsRequest{Key: "key1", Fields: fieldsToRemove}
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

		reqSet := &kvdbserver.SetHashMapRequest{Key: "key1", Fields: make(map[string]string)}
		server.SetHashMap(ctx, reqSet)

		var expectedFieldsRemoved uint32 = 0
		expectedOk := true
		req := &kvdbserver.DeleteHashMapFieldsRequest{Key: "key1", Fields: fieldsToRemove}
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

		reqSet := &kvdbserver.SetHashMapRequest{Key: "key1", Fields: fields}
		server.SetHashMap(ctx, reqSet)

		var expectedFieldsRemoved uint32 = 2
		expectedOk := true
		req := &kvdbserver.DeleteHashMapFieldsRequest{Key: "key1", Fields: fieldsToRemove}
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

		reqSet := &kvdbserver.SetHashMapRequest{Key: "key1", Fields: fields}
		server.SetHashMap(ctx, reqSet)

		var expectedFieldsRemoved uint32 = 1
		expectedOk := true
		req := &kvdbserver.DeleteHashMapFieldsRequest{Key: "key1", Fields: []string{"field3", "field3", "field3"}}
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

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()
		dbName := "db0"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvdbserver.GetAllHashMapFieldsAndValuesRequest{Key: "key1"}
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
		req := &kvdbserver.GetAllHashMapFieldsAndValuesRequest{Key: "key1"}
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

		reqSet := &kvdbserver.SetHashMapRequest{Key: "key1", Fields: fields}
		server.SetHashMap(ctx, reqSet)

		expectedElements := 3
		expectedOk := true
		req := &kvdbserver.GetAllHashMapFieldsAndValuesRequest{Key: "key1"}
		res, err := server.GetAllHashMapFieldsAndValues(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		require.NotNil(t, res.FieldValueMap)
		assert.Equalf(t, expectedElements, len(res.FieldValueMap), "expected elements = %d; got = %d", expectedElements, len(res.FieldValueMap))
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})
}
