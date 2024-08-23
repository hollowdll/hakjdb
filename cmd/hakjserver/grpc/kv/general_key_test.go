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

func TestGetKeyType(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.DefaultDB = "default"

	t.Run("DBNotSentInMetadataUseDefaultDB", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewGeneralKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)

		req := &kvpb.GetKeyTypeRequest{Key: "key1"}
		res, err := gs.GetKeyType(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DBNotFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewGeneralKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, "DBNotFound"))

		req := &kvpb.GetKeyTypeRequest{Key: "key1"}
		res, err := gs.GetKeyType(ctx, req)
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
		gs := NewGeneralKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		expectedKeyType := ""
		expectedOk := false
		req := &kvpb.GetKeyTypeRequest{Key: "key1"}
		res, err := gs.GetKeyType(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedKeyType, res.KeyType, "expected key type = %s; got = %s", expectedKeyType, res.KeyType)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})

	t.Run("StringKey", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gsGeneralKV := NewGeneralKVServiceServer(s)
		gsStringKV := NewStringKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvpb.SetStringRequest{Key: "key1", Value: []byte("val")}
		gsStringKV.SetString(ctx, reqSet)

		expectedKeyType := "String"
		expectedOk := true
		req := &kvpb.GetKeyTypeRequest{Key: "key1"}
		res, err := gsGeneralKV.GetKeyType(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedKeyType, res.KeyType, "expected key type = %s; got = %s", expectedKeyType, res.KeyType)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})

	t.Run("HashMapKey", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gsGeneralKV := NewGeneralKVServiceServer(s)
		gsHashMapKV := NewHashMapKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvpb.SetHashMapRequest{Key: "key1", FieldValueMap: make(map[string][]byte)}
		gsHashMapKV.SetHashMap(ctx, reqSet)

		expectedKeyType := "HashMap"
		expectedOk := true
		req := &kvpb.GetKeyTypeRequest{Key: "key1"}
		res, err := gsGeneralKV.GetKeyType(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedKeyType, res.KeyType, "expected key type = %s; got = %s", expectedKeyType, res.KeyType)
		assert.Equalf(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)
	})
}

func TestGetAllKeys(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.DefaultDB = "default"

	t.Run("DBNotSentInMetadataUseDefaultDB", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewGeneralKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)

		req := &kvpb.GetAllKeysRequest{}
		res, err := gs.GetAllKeys(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DBNotFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewGeneralKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)
		dbName := "db123"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.GetAllKeysRequest{}
		res, err := gs.GetAllKeys(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		expectedCode := codes.NotFound
		expectedOk := true
		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, expectedOk, ok, "expected ok")
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})

	t.Run("NoKeysPresent", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewGeneralKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.GetAllKeysRequest{}
		res, err := gs.GetAllKeys(ctx, req)
		expectedKeys := 0
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equal(t, expectedKeys, len(res.Keys), "expected keys = %d; got = %d", expectedKeys, len(res.Keys))
	})

	t.Run("KeysPresent", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gsGeneralKV := NewGeneralKVServiceServer(s)
		gsStringKV := NewStringKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		keys := []string{"key1", "key2", "key3"}
		for _, key := range keys {
			req := &kvpb.SetStringRequest{Key: key, Value: []byte("val")}
			_, err := gsStringKV.SetString(ctx, req)
			require.NoErrorf(t, err, "expected no error; error = %v", err)
		}

		req := &kvpb.GetAllKeysRequest{}
		res, err := gsGeneralKV.GetAllKeys(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, res)
		assert.Equalf(t, len(keys), len(res.Keys), "expected keys = %d; got = %d", len(keys), len(res.Keys))

		for _, key := range res.Keys {
			assert.Equalf(t, true, common.StringInSlice(key, keys), "expected key %s to be in %v", key, keys)
		}
	})
}

func TestDeleteKeys(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.DefaultDB = "default"

	t.Run("DBNotSentInMetadataUseDefaultDB", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewGeneralKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)

		req := &kvpb.DeleteKeysRequest{Keys: []string{"key1"}}
		res, err := gs.DeleteKeys(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DBNotFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewGeneralKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)
		dbName := "db123"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.DeleteKeysRequest{Keys: []string{"key1"}}
		res, err := gs.DeleteKeys(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		expectedCode := codes.NotFound
		expectedOk := true
		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, expectedOk, ok, "expected ok")
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})

	t.Run("KeyNotFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewGeneralKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.DeleteKeysRequest{Keys: []string{"key1"}}
		res, err := gs.DeleteKeys(ctx, req)
		var expectedKeysDeletedCount uint32 = 0
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedKeysDeletedCount, res.KeysDeletedCount, "expected keys deleted = %d; got = %d", expectedKeysDeletedCount, res.KeysDeletedCount)
	})

	t.Run("KeyFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gsGeneralKV := NewGeneralKVServiceServer(s)
		gsStringKV := NewStringKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvpb.SetStringRequest{Key: "key1", Value: []byte("val")}
		_, err := gsStringKV.SetString(ctx, reqSet)
		require.NoErrorf(t, err, "expected no error; error = %v", err)

		reqDel := &kvpb.DeleteKeysRequest{Keys: []string{"key1", "key2", "key3"}}
		res, err := gsGeneralKV.DeleteKeys(ctx, reqDel)
		var expectedKeysDeletedCount uint32 = 1
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		assert.Equalf(t, expectedKeysDeletedCount, res.KeysDeletedCount, "expected keys deleted = %d; got = %d", expectedKeysDeletedCount, res.KeysDeletedCount)
	})
}

func TestDeleteAllKeys(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.DefaultDB = "default"

	t.Run("DBNotSentInMetadataUseDefaultDB", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewGeneralKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)

		req := &kvpb.DeleteAllKeysRequest{}
		res, err := gs.DeleteAllKeys(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DBNotFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewGeneralKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)
		dbName := "db123"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.DeleteAllKeysRequest{}
		res, err := gs.DeleteAllKeys(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.NotFound, st.Code(), "expected status = %s; got = %s", codes.NotFound, st.Code())
	})

	t.Run("NoKeysPresent", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewGeneralKVServiceServer(s)
		dbName := "db0"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.DeleteAllKeysRequest{}
		res, err := gs.DeleteAllKeys(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, res)
	})

	t.Run("KeysPresent", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gsGeneralKV := NewGeneralKVServiceServer(s)
		gsStringKV := NewStringKVServiceServer(s)
		dbName := "db0"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvpb.SetStringRequest{Key: "key1", Value: []byte("val")}
		gsStringKV.SetString(ctx, reqSet)

		req := &kvpb.DeleteAllKeysRequest{}
		res, err := gsGeneralKV.DeleteAllKeys(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, res)
	})
}
