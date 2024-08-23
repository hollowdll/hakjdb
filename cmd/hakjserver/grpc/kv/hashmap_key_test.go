package kv

import (
	"context"
	"testing"

	"github.com/hollowdll/hakjdb"
	"github.com/hollowdll/hakjdb/api/v1/kvpb"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/config"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/server"
	"github.com/hollowdll/hakjdb/internal/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestSetHashMap(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.DefaultDB = "default"
	fields := make(map[string][]byte)
	fields["field1"] = []byte("value1")
	fields["field2"] = []byte("value2")
	fields["field3"] = []byte("value3")

	t.Run("DBNotSentInMetadataUseDefaultDB", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)

		req := &kvpb.SetHashMapRequest{Key: "key1", FieldValueMap: fields}
		res, err := gs.SetHashMap(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DBNotFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)
		dbName := "db123"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.SetHashMapRequest{Key: "key1", FieldValueMap: fields}
		res, err := gs.SetHashMap(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		expectedOk := true
		expectedCode := codes.NotFound
		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equalf(t, expectedOk, ok, "expected ok = %v; got = %v", expectedOk, ok)
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})

	t.Run("FieldsAdded", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.SetHashMapRequest{Key: "key1", FieldValueMap: fields}
		res, err := gs.SetHashMap(ctx, req)

		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		var expectedFieldsAdded uint32 = 3
		assert.Equal(t, expectedFieldsAdded, res.FieldsAddedCount, "expected fields added = %d; got = %d", expectedFieldsAdded, res.FieldsAddedCount)
	})

	t.Run("OverwriteFields", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		fieldsOverwrite := make(map[string][]byte)
		fieldsOverwrite["field1"] = []byte("a")
		fieldsOverwrite["field2"] = []byte("b")
		fieldsOverwrite["field3"] = []byte("c")
		fieldsOverwrite["new_field"] = []byte("d")

		req1 := &kvpb.SetHashMapRequest{Key: "key1", FieldValueMap: fields}
		res, err := gs.SetHashMap(ctx, req1)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		var expectedFieldsAdded1 uint32 = 3
		assert.Equal(t, expectedFieldsAdded1, res.FieldsAddedCount, "expected fields added = %d; got = %d", expectedFieldsAdded1, res.FieldsAddedCount)

		req2 := &kvpb.SetHashMapRequest{Key: "key1", FieldValueMap: fieldsOverwrite}
		res, err = gs.SetHashMap(ctx, req2)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		var expectedFieldsAdded2 uint32 = 1
		assert.Equal(t, expectedFieldsAdded2, res.FieldsAddedCount, "expected fields added = %d; got = %d", expectedFieldsAdded2, res.FieldsAddedCount)
	})

	t.Run("InvalidKey", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.SetHashMapRequest{Key: "    ", FieldValueMap: fields}
		res, err := gs.SetHashMap(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		expectedOk := true
		expectedCode := codes.InvalidArgument
		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equalf(t, expectedOk, ok, "expected ok = %v; got = %v", expectedOk, ok)
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})

	t.Run("MaxKeyLimitReached", func(t *testing.T) {
		cfg := cfg
		cfg.MaxKeysPerDB = 1
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.SetHashMapRequest{Key: "key1", FieldValueMap: fields}
		res, err := gs.SetHashMap(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)

		req = &kvpb.SetHashMapRequest{Key: "key2", FieldValueMap: fields}
		res, err = gs.SetHashMap(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		expectedOk := true
		expectedCode := codes.FailedPrecondition
		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equalf(t, expectedOk, ok, "expected ok = %v; got = %v", expectedOk, ok)
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})

	t.Run("MaxFieldLimitReached", func(t *testing.T) {
		cfg := cfg
		cfg.MaxHashMapFields = 4
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		fields2 := make(map[string][]byte)
		fields2["field4"] = []byte("val4")
		fields2["field5"] = []byte("val5")
		fields2["field6"] = []byte("val6")
		fields2["field7"] = []byte("val7")

		req1 := &kvpb.SetHashMapRequest{Key: "key1", FieldValueMap: fields}
		res, err := gs.SetHashMap(ctx, req1)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		var expectedFieldsAdded1 uint32 = 3
		assert.Equal(t, expectedFieldsAdded1, res.FieldsAddedCount, "expected fields added = %d; got = %d", expectedFieldsAdded1, res.FieldsAddedCount)

		req2 := &kvpb.SetHashMapRequest{Key: "key1", FieldValueMap: fields2}
		res, err = gs.SetHashMap(ctx, req2)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		var expectedFieldsAdded2 uint32 = 1
		assert.Equal(t, expectedFieldsAdded2, res.FieldsAddedCount, "expected fields added = %d; got = %d", expectedFieldsAdded2, res.FieldsAddedCount)
	})
}

func TestGetHashMapFieldValues(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.DefaultDB = "default"
	fields := make(map[string][]byte)
	fields["field1"] = []byte("value1")
	fields["field2"] = []byte("value2")
	fields["field3"] = []byte("value3")

	t.Run("DBNotSentInMetadataUseDefaultDB", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)

		req := &kvpb.GetHashMapFieldValuesRequest{Key: "key1", Fields: []string{"field2"}}
		res, err := gs.GetHashMapFieldValues(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DBNotFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)
		dbName := "db123"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.GetHashMapFieldValuesRequest{Key: "key1", Fields: []string{"field2"}}
		res, err := gs.GetHashMapFieldValues(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		expectedOk := true
		expectedCode := codes.NotFound
		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equalf(t, expectedOk, ok, "expected ok = %v; got = %v", expectedOk, ok)
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})

	t.Run("KeyAndFieldFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvpb.SetHashMapRequest{Key: "key1", FieldValueMap: fields}
		gs.SetHashMap(ctx, reqSet)

		expectedValue := []byte("value2")
		expectedOk := true
		expectedKeyFound := true
		reqGet := &kvpb.GetHashMapFieldValuesRequest{Key: "key1", Fields: []string{"field2"}}
		res, err := gs.GetHashMapFieldValues(ctx, reqGet)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedValue, res.FieldValueMap["field2"].Value, "expected value = %s; got = %s", expectedValue, res.FieldValueMap["field2"].Value)
		assert.Equalf(t, expectedOk, res.FieldValueMap["field2"].Ok, "expected ok = %v; got = %v", expectedOk, res.FieldValueMap["field2"].Ok)
		assert.Equalf(t, expectedKeyFound, res.Ok, "expected key found = %v; got = %v", expectedKeyFound, res.Ok)
	})

	t.Run("MultipleFieldsFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvpb.SetHashMapRequest{Key: "key1", FieldValueMap: fields}
		gs.SetHashMap(ctx, reqSet)

		reqGet := &kvpb.GetHashMapFieldValuesRequest{Key: "key1", Fields: []string{"field1", "field2", "field3"}}
		res, err := gs.GetHashMapFieldValues(ctx, reqGet)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)

		for field, expectedValue := range fields {
			assert.Equalf(t, expectedValue, res.FieldValueMap[field].Value, "expected value = %s; got = %s", expectedValue, res.FieldValueMap[field].Value)
			assert.Equalf(t, true, res.FieldValueMap[field].Ok, "expected ok = %v; got = %v", true, res.FieldValueMap[field].Ok)
		}
	})

	t.Run("KeyNotFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		expectedOk := false
		req := &kvpb.GetHashMapFieldValuesRequest{Key: "key2", Fields: []string{"field3"}}
		res, err := gs.GetHashMapFieldValues(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})

	t.Run("FieldNotFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvpb.SetHashMapRequest{Key: "key1", FieldValueMap: fields}
		gs.SetHashMap(ctx, reqSet)

		expectedValue := []byte(nil)
		expectedOk := false
		req := &kvpb.GetHashMapFieldValuesRequest{Key: "key1", Fields: []string{"field123"}}
		res, err := gs.GetHashMapFieldValues(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedValue, res.FieldValueMap["field123"].Value, "expected value = %s; got = %s", expectedValue, res.FieldValueMap["field123"].Value)
		assert.Equalf(t, expectedOk, res.FieldValueMap["field123"].Ok, "expected ok = %v; got = %v", expectedOk, res.FieldValueMap["field123"].Ok)
	})
}

func TestDeleteHashMapFields(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.DefaultDB = "default"
	fields := make(map[string][]byte)
	fields["field1"] = []byte("value1")
	fields["field2"] = []byte("value2")
	fields["field3"] = []byte("value3")
	fieldsToRemove := []string{"field2", "field3"}

	t.Run("DBNotSentInMetadataUseDefaultDB", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)

		req := &kvpb.DeleteHashMapFieldsRequest{Key: "key1", Fields: fieldsToRemove}
		res, err := gs.DeleteHashMapFields(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DBNotFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)
		dbName := "db123"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.DeleteHashMapFieldsRequest{Key: "key1", Fields: fieldsToRemove}
		res, err := gs.DeleteHashMapFields(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		expectedOk := true
		expectedCode := codes.NotFound
		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equalf(t, expectedOk, ok, "expected ok = %v; got = %v", expectedOk, ok)
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})

	t.Run("KeyNotFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		var expectedFieldsRemoved uint32 = 0
		expectedOk := false
		req := &kvpb.DeleteHashMapFieldsRequest{Key: "key1", Fields: fieldsToRemove}
		res, err := gs.DeleteHashMapFields(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedFieldsRemoved, res.FieldsRemovedCount, "expected fields removed = %d; got = %d", expectedFieldsRemoved, res.FieldsRemovedCount)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})

	t.Run("FieldsNotExist", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvpb.SetHashMapRequest{Key: "key1", FieldValueMap: make(map[string][]byte)}
		gs.SetHashMap(ctx, reqSet)

		var expectedFieldsRemoved uint32 = 0
		expectedOk := true
		req := &kvpb.DeleteHashMapFieldsRequest{Key: "key1", Fields: fieldsToRemove}
		res, err := gs.DeleteHashMapFields(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedFieldsRemoved, res.FieldsRemovedCount, "expected fields removed = %d; got = %d", expectedFieldsRemoved, res.FieldsRemovedCount)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})

	t.Run("FieldsExist", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvpb.SetHashMapRequest{Key: "key1", FieldValueMap: fields}
		gs.SetHashMap(ctx, reqSet)

		var expectedFieldsRemoved uint32 = 2
		expectedOk := true
		req := &kvpb.DeleteHashMapFieldsRequest{Key: "key1", Fields: fieldsToRemove}
		res, err := gs.DeleteHashMapFields(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedFieldsRemoved, res.FieldsRemovedCount, "expected fields removed = %d; got = %d", expectedFieldsRemoved, res.FieldsRemovedCount)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})

	t.Run("DuplicateFields", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvpb.SetHashMapRequest{Key: "key1", FieldValueMap: fields}
		gs.SetHashMap(ctx, reqSet)

		var expectedFieldsRemoved uint32 = 1
		expectedOk := true
		req := &kvpb.DeleteHashMapFieldsRequest{Key: "key1", Fields: []string{"field3", "field3", "field3"}}
		res, err := gs.DeleteHashMapFields(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedFieldsRemoved, res.FieldsRemovedCount, "expected fields removed = %d; got = %d", expectedFieldsRemoved, res.FieldsRemovedCount)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})
}

func TestGetAllHashMapFieldsAndValues(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.DefaultDB = "default"
	fields := make(map[string][]byte)
	fields["field1"] = []byte("value1")
	fields["field2"] = []byte("value2")
	fields["field3"] = []byte("value3")

	t.Run("DBNotSentInMetadataUseDefaultDB", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)

		req := &kvpb.GetAllHashMapFieldsAndValuesRequest{Key: "key1"}
		res, err := gs.GetAllHashMapFieldsAndValues(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DBNotFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)
		dbName := "db123"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.GetAllHashMapFieldsAndValuesRequest{Key: "key1"}
		res, err := gs.GetAllHashMapFieldsAndValues(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		expectedOk := true
		expectedCode := codes.NotFound
		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equalf(t, expectedOk, ok, "expected ok = %v; got = %v", expectedOk, ok)
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})

	t.Run("KeyNotFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		expectedElements := 0
		expectedOk := false
		req := &kvpb.GetAllHashMapFieldsAndValuesRequest{Key: "key1"}
		res, err := gs.GetAllHashMapFieldsAndValues(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		require.NotNil(t, res.FieldValueMap)
		assert.Equalf(t, expectedElements, len(res.FieldValueMap), "expected elements = %d; got = %d", expectedElements, len(res.FieldValueMap))
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})

	t.Run("KeyFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewHashMapKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvpb.SetHashMapRequest{Key: "key1", FieldValueMap: fields}
		gs.SetHashMap(ctx, reqSet)

		expectedElements := 3
		expectedOk := true
		req := &kvpb.GetAllHashMapFieldsAndValuesRequest{Key: "key1"}
		res, err := gs.GetAllHashMapFieldsAndValues(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		require.NotNil(t, res.FieldValueMap)
		assert.Equalf(t, expectedElements, len(res.FieldValueMap), "expected elements = %d; got = %d", expectedElements, len(res.FieldValueMap))
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})
}
