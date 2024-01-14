package main_test

import (
	"context"
	"testing"
	"time"

	main "github.com/hollowdll/kvdb/cmd/kvdbserver"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestSetString(t *testing.T) {
	t.Run("MissingMetadata", func(t *testing.T) {
		server := main.NewServer()
		server.DisableLogger()

		request := &kvdbserver.SetStringRequest{Key: "key1", Value: "value1"}
		response, err := server.SetString(context.Background(), request)
		assert.Error(t, err, "expected error")
		assert.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		assert.NotNil(t, st, "expected status to be non-nil")
		assert.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})

	t.Run("MissingDatabaseInMetadata", func(t *testing.T) {
		server := main.NewServer()
		server.DisableLogger()
		dbName := "db0"

		md := metadata.Pairs("wrong-key", dbName)
		ctx := metadata.NewIncomingContext(context.Background(), md)
		ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
		defer cancel()

		request := &kvdbserver.SetStringRequest{Key: "key1", Value: "value1"}
		response, err := server.SetString(ctx, request)
		assert.Error(t, err, "expected error")
		assert.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		assert.NotNil(t, st, "expected status to be non-nil")
		assert.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		server := main.NewServer()
		server.DisableLogger()
		dbName := "db0"

		md := metadata.Pairs(common.GrpcMetadataKeyDbName, dbName)
		ctx := metadata.NewIncomingContext(context.Background(), md)
		ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
		defer cancel()

		request := &kvdbserver.SetStringRequest{Key: "key1", Value: "value1"}
		response, err := server.SetString(ctx, request)
		assert.Error(t, err, "expected error")
		assert.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		assert.NotNil(t, st, "expected status to be non-nil")
		assert.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.NotFound, st.Code(), "expected status = %s; got = %s", codes.NotFound, st.Code())
	})

	t.Run("DatabaseExists", func(t *testing.T) {
		server := main.NewServer()
		server.DisableLogger()
		dbName := "db0"

		requestCreate := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), requestCreate)
		assert.NoErrorf(t, err, "expected no error; error = %s", err)

		md := metadata.Pairs(common.GrpcMetadataKeyDbName, dbName)
		ctx := metadata.NewIncomingContext(context.Background(), md)
		ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
		defer cancel()

		request := &kvdbserver.SetStringRequest{Key: "key1", Value: "value1"}
		response, err := server.SetString(ctx, request)
		assert.NoErrorf(t, err, "expected no error; error = %s", err)
		assert.NotNil(t, response, "expected response to be non-nil")
	})

	t.Run("InvalidInput", func(t *testing.T) {
		server := main.NewServer()
		server.DisableLogger()
		dbName := "db0"

		requestCreate := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
		_, err := server.CreateDatabase(context.Background(), requestCreate)
		assert.NoErrorf(t, err, "expected no error; error = %s", err)

		md := metadata.Pairs(common.GrpcMetadataKeyDbName, dbName)
		ctx := metadata.NewIncomingContext(context.Background(), md)
		ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
		defer cancel()

		request := &kvdbserver.SetStringRequest{Key: "      ", Value: "value1"}
		response, err := server.SetString(ctx, request)
		assert.Error(t, err, "expected error")
		assert.Nil(t, response, "expected response to be nil")

		st, ok := status.FromError(err)
		assert.NotNil(t, st, "expected status to be non-nil")
		assert.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})
}
