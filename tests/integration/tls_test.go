package integration

import (
	"context"
	"testing"

	"github.com/hollowdll/hakjdb/api/v1/echopb"
	"github.com/hollowdll/hakjdb/api/v1/serverpb"
	"github.com/hollowdll/hakjdb/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestGetServerInfoWithTLS(t *testing.T) {
	cfg := testutil.TLSConfig(false)
	_, gs, port := testutil.StartTestServer(cfg)
	defer testutil.StopTestServer(gs)
	address := testutil.GetServerAddress(port)

	t.Run("WithCredentials", func(t *testing.T) {
		conn, err := testutil.SecureConnection(address, false)
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

func TestTLSClientCertAuth(t *testing.T) {
	cfg := testutil.TLSConfig(true)
	_, gs, port := testutil.StartTestServer(cfg)
	defer testutil.StopTestServer(gs)
	address := testutil.GetServerAddress(port)

	t.Run("WithClientCert", func(t *testing.T) {
		conn, err := testutil.SecureConnection(address, true)
		require.NoErrorf(t, err, "expected connection but connection failed: %v", err)
		defer conn.Close()
		client := echopb.NewEchoServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
		defer cancel()

		req := &echopb.UnaryEchoRequest{}
		res, err := client.UnaryEcho(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("WithoutClientCert", func(t *testing.T) {
		conn, err := testutil.SecureConnection(address, false)
		require.NoErrorf(t, err, "expected connection but connection failed: %v", err)
		defer conn.Close()
		client := echopb.NewEchoServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
		defer cancel()

		req := &echopb.UnaryEchoRequest{}
		res, err := client.UnaryEcho(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)
	})
}
