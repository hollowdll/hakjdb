package integration

import (
	"context"
	"testing"

	"github.com/hollowdll/kvdb/proto/kvdbserverpb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDatabase(t *testing.T) {
	conn, err := insecureConnection()
	require.NoErrorf(t, err, "Failed to connect to the server: %v", err)
	defer conn.Close()
	databaseClient := kvdbserverpb.NewDatabaseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	dbName := "TestCreateDatabase"

	req := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
	res, err := databaseClient.CreateDatabase(ctx, req)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res)
	assert.Equalf(t, dbName, res.DbName, "expected database name = %s; got = %s", dbName, res.DbName)
}

func TestCreateDatabaseAndGetInfo(t *testing.T) {
	conn, err := insecureConnection()
	require.NoErrorf(t, err, "Failed to connect to the server: %v", err)
	defer conn.Close()
	databaseClient := kvdbserverpb.NewDatabaseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	dbName := "TestCreateDatabaseAndGetInfo"

	req1 := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
	res1, err := databaseClient.CreateDatabase(ctx, req1)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res1)

	req2 := &kvdbserverpb.GetDatabaseInfoRequest{DbName: dbName}
	res2, err := databaseClient.GetDatabaseInfo(ctx, req2)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res2)
	assert.Equalf(t, dbName, res2.Data.Name, "expected database name = %s; got = %s", dbName, res2.Data.Name)
}

func TestCreateDatabaseAlreadyExists(t *testing.T) {
	conn, err := insecureConnection()
	require.NoErrorf(t, err, "Failed to connect to the server: %v", err)
	defer conn.Close()
	databaseClient := kvdbserverpb.NewDatabaseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	dbName := "TestCreateDatabaseAlreadyExists"

	req1 := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
	res1, err := databaseClient.CreateDatabase(ctx, req1)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res1)

	req2 := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
	res2, err := databaseClient.CreateDatabase(ctx, req2)
	assert.Error(t, err, "expected error")
	assert.Nil(t, res2, "expected nil")
}

func TestDefaultDatabaseCreated(t *testing.T) {
	conn, err := insecureConnection()
	require.NoErrorf(t, err, "Failed to connect to the server: %v", err)
	defer conn.Close()
	databaseClient := kvdbserverpb.NewDatabaseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	dbName := "default"

	req := &kvdbserverpb.GetDatabaseInfoRequest{DbName: dbName}
	res, err := databaseClient.GetDatabaseInfo(ctx, req)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res)
	assert.Equalf(t, dbName, res.Data.Name, "expected database name = %s; got = %s", dbName, res.Data.Name)
}

func TestCreateAndDeleteDatabase(t *testing.T) {
	conn, err := insecureConnection()
	require.NoErrorf(t, err, "Failed to connect to the server: %v", err)
	defer conn.Close()
	databaseClient := kvdbserverpb.NewDatabaseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	dbName := "TestCreateAndDeleteDatabase"

	req1 := &kvdbserverpb.CreateDatabaseRequest{DbName: dbName}
	res1, err := databaseClient.CreateDatabase(ctx, req1)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res1)

	req2 := &kvdbserverpb.DeleteDatabaseRequest{DbName: dbName}
	res2, err := databaseClient.DeleteDatabase(ctx, req2)
	require.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res2)

	req3 := &kvdbserverpb.GetDatabaseInfoRequest{DbName: dbName}
	res3, err := databaseClient.GetDatabaseInfo(ctx, req3)
	require.Error(t, err, "expected error")
	require.Nil(t, res3)
}
