package kv

import (
	"context"
	"testing"

	"github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/api/v0/kvpb"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/config"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/internal/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestSetString(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.DefaultDB = "default"

	t.Run("DBNotSentInMetadataUseDefaultDB", func(t *testing.T) {
		s := server.NewKvdbServer(cfg, kvdb.DisabledLogger())
		gs := NewStringKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)

		req := &kvpb.SetStringRequest{Key: "key1", Value: []byte("val")}
		res, err := gs.SetString(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DBNotFound", func(t *testing.T) {
		s := server.NewKvdbServer(cfg, kvdb.DisabledLogger())
		gs := NewStringKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)
		dbName := "db123"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.SetStringRequest{Key: "key1", Value: []byte("val")}
		resp, err := gs.SetString(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.NotFound, st.Code(), "expected status = %s; got = %s", codes.NotFound, st.Code())
	})

	t.Run("Success", func(t *testing.T) {
		s := server.NewKvdbServer(cfg, kvdb.DisabledLogger())
		gs := NewStringKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.SetStringRequest{Key: "key1", Value: []byte("val")}
		resp, err := gs.SetString(ctx, req)
		assert.NoErrorf(t, err, "expected no error; error = %s", err)
		assert.NotNil(t, resp)
	})

	t.Run("InvalidKey", func(t *testing.T) {
		s := server.NewKvdbServer(cfg, kvdb.DisabledLogger())
		gs := NewStringKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.SetStringRequest{Key: "      ", Value: []byte("val")}
		resp, err := gs.SetString(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})

	t.Run("MaxKeyLimitReached", func(t *testing.T) {
		cfg := cfg
		cfg.MaxKeysPerDB = 1
		s := server.NewKvdbServer(cfg, kvdb.DisabledLogger())
		gs := NewStringKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.SetStringRequest{Key: "key1", Value: []byte("val")}
		resp, err := gs.SetString(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, resp)

		req = &kvpb.SetStringRequest{Key: "key2", Value: []byte("val")}
		resp, err = gs.SetString(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)

		expectedOk := true
		expectedCode := codes.FailedPrecondition
		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equalf(t, expectedOk, ok, "expected ok = %v; got = %v", expectedOk, ok)
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})
}

func TestGetString(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.DefaultDB = "default"

	t.Run("DBNotSentInMetadataUseDefaultDB", func(t *testing.T) {
		s := server.NewKvdbServer(cfg, kvdb.DisabledLogger())
		gs := NewStringKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)

		req := &kvpb.GetStringRequest{Key: "key1"}
		res, err := gs.GetString(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("DBNotFound", func(t *testing.T) {
		s := server.NewKvdbServer(cfg, kvdb.DisabledLogger())
		gs := NewStringKVServiceServer(s)
		s.CreateDefaultDatabase(cfg.DefaultDB)
		dbName := "db123"
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		req := &kvpb.GetStringRequest{Key: "key1"}
		res, err := gs.GetString(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.NotFound, st.Code(), "expected status = %s; got = %s", codes.NotFound, st.Code())
	})

	t.Run("KeyNotFound", func(t *testing.T) {
		s := server.NewKvdbServer(cfg, kvdb.DisabledLogger())
		gs := NewStringKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		expectedValue := []byte(nil)
		req := &kvpb.GetStringRequest{Key: "key1"}
		res, err := gs.GetString(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, res)
		assert.Equalf(t, false, res.Ok, "expected ok = %v; got = %v", false, res.Ok)
		assert.Equalf(t, expectedValue, res.Value, "expected value = %s; got = %s", expectedValue, res.Value)
	})

	t.Run("KeyFound", func(t *testing.T) {
		s := server.NewKvdbServer(cfg, kvdb.DisabledLogger())
		gs := NewStringKVServiceServer(s)
		dbName := "db123"
		s.CreateDefaultDatabase(dbName)
		expectedValue := []byte("val")
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))

		reqSet := &kvpb.SetStringRequest{Key: "key1", Value: expectedValue}
		_, err := gs.SetString(ctx, reqSet)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		reqGet := &kvpb.GetStringRequest{Key: "key1"}
		res, err := gs.GetString(ctx, reqGet)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, res)
		assert.Equalf(t, true, res.Ok, "expected ok = %v; got = %v", true, res.Ok)
		assert.Equalf(t, expectedValue, res.Value, "expected value = %s; got = %s", expectedValue, res.Value)
	})
}
