package integration

import (
	"context"
	"testing"

	"github.com/hollowdll/kvdb/api/v0/serverpb"
	"github.com/hollowdll/kvdb/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestGetServerInfoWithTLS(t *testing.T) {
	cfg := testutil.TLSConfig()
	_, gs, port := testutil.StartTestServer(cfg)
	defer testutil.StopTestServer(gs)
	address := testutil.GetServerAddress(port)

	t.Run("WithCredentials", func(t *testing.T) {
		conn, err := testutil.SecureConnection(address)
		require.NoErrorf(t, err, "expected connection but connection failed: %v", err)
		defer conn.Close()
		client := serverpb.NewServerServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
		defer cancel()

		req := &serverpb.GetServerInfoRequest{}
		res, err := client.GetServerInfo(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("WithoutCredentials", func(t *testing.T) {
		conn, err := testutil.InsecureConnection(address)
		require.NoErrorf(t, err, "expected connection but connection failed: %v", err)
		defer conn.Close()
		client := serverpb.NewServerServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
		defer cancel()

		req := &serverpb.GetServerInfoRequest{}
		res, err := client.GetServerInfo(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)
	})
}
