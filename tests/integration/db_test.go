package integration

import (
	"context"
	"testing"
	"time"

	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestCreateDatabase(t *testing.T) {
	conn, err := grpc.Dial(getServerAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoErrorf(t, err, "Failed to connect to the server: %v", err)
	defer conn.Close()
	databaseClient := kvdbserver.NewDatabaseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dbName := "TestCreateDatabase"

	req := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
	res, err := databaseClient.CreateDatabase(ctx, req)

	assert.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res)
	assert.Equalf(t, dbName, res.DbName, "expected database name = %s; got = %s", dbName, res.DbName)
}

func TestCreateDatabaseAndGetInfo(t *testing.T) {
	conn, err := grpc.Dial(getServerAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoErrorf(t, err, "Failed to connect to the server: %v", err)
	defer conn.Close()
	databaseClient := kvdbserver.NewDatabaseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dbName := "TestCreateDatabaseAndGetInfo"

	req1 := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
	res1, err := databaseClient.CreateDatabase(ctx, req1)
	assert.NoErrorf(t, err, "expected no error; error = %v", err)
	assert.NotNil(t, res1)

	req2 := &kvdbserver.GetDatabaseInfoRequest{DbName: dbName}
	res2, err := databaseClient.GetDatabaseInfo(ctx, req2)
	assert.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, res2)
	assert.Equalf(t, dbName, res2.Data.Name, "expected database name = %s; got = %s", dbName, res2.Data.Name)
}

func TestCreateDatabaseAlreadyExists(t *testing.T) {
	conn, err := grpc.Dial(getServerAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoErrorf(t, err, "Failed to connect to the server: %v", err)
	defer conn.Close()
	databaseClient := kvdbserver.NewDatabaseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dbName := "TestCreateDatabaseAlreadyExists"

	req1 := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
	res1, err := databaseClient.CreateDatabase(ctx, req1)
	assert.NoErrorf(t, err, "expected no error; error = %v", err)
	assert.NotNil(t, res1)

	req2 := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
	res2, err := databaseClient.CreateDatabase(ctx, req2)
	assert.Error(t, err, "expected error")
	assert.Nil(t, res2, "expected nil")
}
