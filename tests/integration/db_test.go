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

	dbName := "db0"
	request := &kvdbserver.CreateDatabaseRequest{DbName: dbName}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := databaseClient.CreateDatabase(ctx, request)

	assert.NoErrorf(t, err, "expected no error; error = %v", err)
	require.NotNil(t, response)
	assert.Equalf(t, dbName, response.DbName, "expected database name = %s; got = %s", dbName, response.DbName)
}
