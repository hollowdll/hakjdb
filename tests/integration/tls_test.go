package integration

import (
	"context"
	"testing"

	"github.com/hollowdll/kvdb/proto/kvdbserverpb"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetServerInfoWithTls(t *testing.T) {
	t.Run("WithCredentials", func(t *testing.T) {
		conn, err := secureConnection()
		require.NoErrorf(t, err, "expected connection but connection failed: %v", err)
		defer conn.Close()
		client := kvdbserverpb.NewServerServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
		defer cancel()

		req := &kvdbserverpb.GetServerInfoRequest{}
		res, err := client.GetServerInfo(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("NoCredentials", func(t *testing.T) {
		conn, err := grpc.Dial(getTlsServerAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		require.NoErrorf(t, err, "expected connection but connection failed: %v", err)
		defer conn.Close()
		client := kvdbserverpb.NewServerServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
		defer cancel()

		req := &kvdbserverpb.GetServerInfoRequest{}
		res, err := client.GetServerInfo(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)
	})
}
