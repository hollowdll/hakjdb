package integration

import (
	"context"
	"testing"

	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserverpb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func TestOverwriteAndGetString(t *testing.T) {
	conn, err := insecureConnection()
	require.NoErrorf(t, err, "Failed to connect to the server: %v", err)
	defer conn.Close()
	databaseClient := kvdbserverpb.NewDatabaseServiceClient(conn)
	storageClient := kvdbserverpb.NewStorageServiceClient(conn)

	dbName := "TestSetAndGetString"
	key := "key1"
	value1 := "value1"
	value2 := "value2"
	ctxMd := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))
	ctx, cancel := context.WithTimeout(ctxMd, ctxTimeout)
	defer cancel()

	req1 := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
	res1, err := databaseClient.CreateDatabase(ctx, req1)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res1)

	req2 := &kvdbserverpb.SetStringRequest{Key: key, Value: value1}
	res2, err := storageClient.SetString(ctx, req2)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res2)

	req3 := &kvdbserverpb.SetStringRequest{Key: key, Value: value2}
	res3, err := storageClient.SetString(ctx, req3)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res3)

	req4 := &kvdbserverpb.GetStringRequest{Key: key}
	res4, err := storageClient.GetString(ctx, req4)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res4)
	assert.Equal(t, value2, res4.Value, "expected value = %s; got = %s", value2, res4.Value)
	assert.Equal(t, true, res4.Ok, "expected ok = %v; got = %v", true, res4.Ok)
}

func TestSetGetDeleteString(t *testing.T) {
	conn, err := insecureConnection()
	require.NoErrorf(t, err, "Failed to connect to the server: %v", err)
	defer conn.Close()
	databaseClient := kvdbserverpb.NewDatabaseServiceClient(conn)
	storageClient := kvdbserverpb.NewStorageServiceClient(conn)

	dbName := "TestSetGetDeleteString"
	key := "key1"
	value := "value1"
	ctxMd := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))
	ctx, cancel := context.WithTimeout(ctxMd, ctxTimeout)
	defer cancel()

	req1 := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
	res1, err := databaseClient.CreateDatabase(ctx, req1)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res1)

	req2 := &kvdbserverpb.SetStringRequest{Key: key, Value: value}
	res2, err := storageClient.SetString(ctx, req2)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res2)

	req3 := &kvdbserverpb.GetStringRequest{Key: key}
	res3, err := storageClient.GetString(ctx, req3)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res3)
	assert.Equal(t, value, res3.Value, "expected value = %s; got = %s", value, res3.Value)
	assert.Equal(t, true, res3.Ok, "expected ok = %v; got = %v", true, res3.Ok)

	req4 := &kvdbserverpb.DeleteKeyRequest{Key: key}
	res4, err := storageClient.DeleteKey(ctx, req4)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res4)
	assert.Equal(t, true, res4.Ok, "expected ok = %v; got = %v", true, res4.Ok)

	req5 := &kvdbserverpb.GetStringRequest{Key: key}
	res5, err := storageClient.GetString(ctx, req5)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res5)
	assert.Equal(t, "", res5.Value, "expected empty string; got = %s", value, res5.Value)
	assert.Equal(t, false, res5.Ok, "expected ok = %v; got = %v", false, res5.Ok)

	req6 := &kvdbserverpb.DeleteKeyRequest{Key: key}
	res6, err := storageClient.DeleteKey(ctx, req6)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res6)
	assert.Equal(t, false, res6.Ok, "expected ok = %v; got = %v", false, res6.Ok)
}

func TestDeleteAllKeys(t *testing.T) {
	conn, err := insecureConnection()
	require.NoErrorf(t, err, "Failed to connect to the server: %v", err)
	defer conn.Close()
	databaseClient := kvdbserverpb.NewDatabaseServiceClient(conn)
	storageClient := kvdbserverpb.NewStorageServiceClient(conn)

	dbName := "TestDeleteAllKeys"
	ctxMd := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))
	ctx, cancel := context.WithTimeout(ctxMd, ctxTimeout)
	defer cancel()

	reqCreate := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
	databaseClient.CreateDatabase(ctx, reqCreate)

	reqSet1 := &kvdbserverpb.SetStringRequest{Key: "key1", Value: "value"}
	storageClient.SetString(ctx, reqSet1)
	reqSet2 := &kvdbserverpb.SetStringRequest{Key: "key2", Value: "value"}
	storageClient.SetString(ctx, reqSet2)
	reqSet3 := &kvdbserverpb.SetStringRequest{Key: "key3", Value: "value"}
	storageClient.SetString(ctx, reqSet3)

	expectedKeys := uint32(3)
	reqGet := &kvdbserverpb.GetDatabaseInfoRequest{DbName: dbName}
	resGet1, err := databaseClient.GetDatabaseInfo(ctx, reqGet)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, resGet1)
	assert.Equalf(t, expectedKeys, resGet1.Data.KeyCount, "expected keys = %d; got = %d", expectedKeys, resGet1.Data.KeyCount)

	reqDelete := &kvdbserverpb.DeleteAllKeysRequest{}
	storageClient.DeleteAllKeys(ctx, reqDelete)

	expectedKeys = uint32(0)
	resGet2, err := databaseClient.GetDatabaseInfo(ctx, reqGet)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, resGet2)
	assert.Equalf(t, expectedKeys, resGet2.Data.KeyCount, "expected keys = %d; got = %d", expectedKeys, resGet2.Data.KeyCount)
}

func TestSetGetDeleteHashMap(t *testing.T) {
	conn, err := insecureConnection()
	require.NoErrorf(t, err, "Failed to connect to the server: %v", err)
	defer conn.Close()
	databaseClient := kvdbserverpb.NewDatabaseServiceClient(conn)
	storageClient := kvdbserverpb.NewStorageServiceClient(conn)

	fields := make(map[string]string)
	fields["field1"] = "value1"
	fields["field2"] = "value2"
	fields["field3"] = "value3"

	dbName := "TestSetGetDeleteHashMap"
	ctxMd := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyDbName, dbName))
	ctx, cancel := context.WithTimeout(ctxMd, ctxTimeout)
	defer cancel()

	reqCreate := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
	databaseClient.CreateDatabase(ctx, reqCreate)

	reqSet := &kvdbserverpb.SetHashMapRequest{Key: "key1", Fields: fields}
	storageClient.SetHashMap(ctx, reqSet)

	expectedValue := "value2"
	expectedOk := true
	reqGet := &kvdbserverpb.GetHashMapFieldValueRequest{Key: "key1", Field: "field2"}
	res, _ := storageClient.GetHashMapFieldValue(ctx, reqGet)
	require.NotNil(t, res)
	assert.Equal(t, expectedValue, res.Value, "expected value = %s; got = %s", expectedValue, res.Value)
	assert.Equal(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)

	reqDelete := &kvdbserverpb.DeleteHashMapFieldsRequest{Key: "key1", Fields: []string{"field2"}}
	storageClient.DeleteHashMapFields(ctx, reqDelete)

	expectedValue = ""
	expectedOk = false
	res, _ = storageClient.GetHashMapFieldValue(ctx, reqGet)
	require.NotNil(t, res)
	assert.Equal(t, expectedValue, res.Value, "expected value = %s; got = %s", expectedValue, res.Value)
	assert.Equal(t, expectedOk, res.Ok, "expected ok = %v; got = %v", expectedOk, res.Ok)

	expectedKeys := uint32(1)
	reqInfo := &kvdbserverpb.GetDatabaseInfoRequest{DbName: dbName}
	resInfo, _ := databaseClient.GetDatabaseInfo(ctx, reqInfo)
	require.NotNil(t, resInfo)
	assert.Equalf(t, expectedKeys, resInfo.Data.KeyCount, "expected keys = %d; got = %d", expectedKeys, resInfo.Data.KeyCount)
}
