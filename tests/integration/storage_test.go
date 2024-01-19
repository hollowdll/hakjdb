package integration

import (
	"context"
	"testing"
	"time"

	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func TestSetAndGetString(t *testing.T) {
	conn, err := grpc.Dial(getServerAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoErrorf(t, err, "Failed to connect to the server: %v", err)
	defer conn.Close()
	databaseClient := kvdbserver.NewDatabaseServiceClient(conn)
	storageClient := kvdbserver.NewStorageServiceClient(conn)
	dbName := "TestSetAndGetString"
	key := "key1"
	value := "value1"
	ctxMd := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))
	ctx, cancel := context.WithTimeout(ctxMd, 5*time.Second)
	defer cancel()

	req1 := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
	res1, err := databaseClient.CreateDatabase(ctx, req1)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res1)

	req2 := &kvdbserver.SetStringRequest{Key: key, Value: value}
	res2, err := storageClient.SetString(ctx, req2)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res2)

	req3 := &kvdbserver.GetStringRequest{Key: key}
	res3, err := storageClient.GetString(ctx, req3)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res3)
	assert.Equal(t, value, res3.Value, "expected value = %s; got = %s", value, res3.Value)
	assert.Equal(t, true, res3.Found, "expected found = %v; got = %v", true, res3.Found)
}

func TestSetGetDeleteString(t *testing.T) {
	conn, err := grpc.Dial(getServerAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoErrorf(t, err, "Failed to connect to the server: %v", err)
	defer conn.Close()
	databaseClient := kvdbserver.NewDatabaseServiceClient(conn)
	storageClient := kvdbserver.NewStorageServiceClient(conn)
	dbName := "TestSetGetDeleteString"
	key := "key1"
	value := "value1"
	ctxMd := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))
	ctx, cancel := context.WithTimeout(ctxMd, 5*time.Second)
	defer cancel()

	req1 := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
	res1, err := databaseClient.CreateDatabase(ctx, req1)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res1)

	req2 := &kvdbserver.SetStringRequest{Key: key, Value: value}
	res2, err := storageClient.SetString(ctx, req2)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res2)

	req3 := &kvdbserver.GetStringRequest{Key: key}
	res3, err := storageClient.GetString(ctx, req3)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res3)
	assert.Equal(t, value, res3.Value, "expected value = %s; got = %s", value, res3.Value)
	assert.Equal(t, true, res3.Found, "expected found = %v; got = %v", true, res3.Found)

	req4 := &kvdbserver.DeleteKeyRequest{Key: key}
	res4, err := storageClient.DeleteKey(ctx, req4)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res4)
	assert.Equal(t, true, res4.Success, "expected success = %v; got = %v", true, res4.Success)

	req5 := &kvdbserver.GetStringRequest{Key: key}
	res5, err := storageClient.GetString(ctx, req5)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res5)
	assert.Equal(t, "", res5.Value, "expected empty string; got = %s", value, res5.Value)
	assert.Equal(t, false, res5.Found, "expected found = %v; got = %v", false, res5.Found)
}
