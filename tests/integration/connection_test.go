package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/hollowdll/kvdb/proto/kvdbserverpb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestMaxClientConnections(t *testing.T) {
	var maxConnections uint32 = 5
	server, port := startTestServer(maxConnections)
	defer server.Stop()
	address := fmt.Sprintf("localhost:%d", port)
	connections := make([]*grpc.ClientConn, maxConnections+5)
	req := &kvdbserverpb.GetServerInfoRequest{}
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	for i := 0; i < len(connections); i++ {
		conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		assert.NoErrorf(t, err, "expected no error; error = %v", err)

		client := kvdbserverpb.NewServerServiceClient(conn)
		_, err = client.GetServerInfo(ctx, req)

		if i < int(maxConnections) {
			assert.NoErrorf(t, err, "expected no error; error = %v", err)
		} else if i >= int(maxConnections) {
			assert.Error(t, err)
		}
		connections[i] = conn
	}

	for _, conn := range connections {
		if conn != nil {
			conn.Close()
		}
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoErrorf(t, err, "expected no error; error = %v", err)

	client := kvdbserverpb.NewServerServiceClient(conn)
	_, err = client.GetServerInfo(ctx, req)
	assert.NoErrorf(t, err, "expected no error; error = %v", err)
}
