package integration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestMaxClientConnections(t *testing.T) {
	var maxConnections uint32 = 5
	server, port := startMaxClientConnectionsTestServer(maxConnections)
	defer server.Stop()
	address := fmt.Sprintf("localhost:%d", port)
	connections := make([]*grpc.ClientConn, maxConnections+5)

	for i := 0; i < len(connections); i++ {
		conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
}
