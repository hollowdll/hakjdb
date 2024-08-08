package integration

import (
	"context"
	"testing"

	"github.com/hollowdll/kvdb/api/v0/serverpb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestMaxClientConnections(t *testing.T) {
	var maxConnections uint32 = 5
	cfg := defaultConfig()
	cfg.MaxClientConnections = maxConnections
	_, gs, port := startTestServer(cfg)
	defer gs.Stop()

	address := getServerAddress(port)
	connections := make([]*grpc.ClientConn, maxConnections+5)
	req := &serverpb.GetServerInfoRequest{}
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	for i := 0; i < len(connections); i++ {
		conn, err := insecureConnection(address)
		assert.NoErrorf(t, err, "expected no error; error = %v", err)

		client := serverpb.NewServerServiceClient(conn)
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

	conn, err := insecureConnection(address)
	assert.NoErrorf(t, err, "expected no error; error = %v", err)

	client := serverpb.NewServerServiceClient(conn)
	_, err = client.GetServerInfo(ctx, req)
	assert.NoErrorf(t, err, "expected no error; error = %v", err)
}
