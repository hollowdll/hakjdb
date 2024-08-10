package integration

import (
	"context"
	"testing"

	"github.com/hollowdll/kvdb/api/v0/serverpb"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestPasswordAuth(t *testing.T) {
	cfg := testutil.DefaultConfig()
	password := "pass123"
	s, gs, port := testutil.StartTestServer(cfg)
	defer testutil.StopTestServer(gs)
	s.EnablePasswordProtection(password)
	address := testutil.GetServerAddress(port)

	t.Run("WithPassword", func(t *testing.T) {
		conn, err := testutil.InsecureConnection(address)
		require.NoErrorf(t, err, "expected connection but connection failed: %v", err)
		defer conn.Close()
		client := serverpb.NewServerServiceClient(conn)
		ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyPassword, password))
		ctx, cancel := context.WithTimeout(ctx, ctxTimeout)
		defer cancel()

		req := &serverpb.GetServerInfoRequest{}
		res, err := client.GetServerInfo(ctx, req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
	})

	t.Run("WithoutPassword", func(t *testing.T) {
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

		expectedOk := true
		expectedCode := codes.Unauthenticated
		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, expectedOk, ok, "expected ok")
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})

	t.Run("InvalidCredentials", func(t *testing.T) {
		conn, err := testutil.InsecureConnection(address)
		require.NoErrorf(t, err, "expected connection but connection failed: %v", err)
		defer conn.Close()
		client := serverpb.NewServerServiceClient(conn)
		ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(common.GrpcMetadataKeyPassword, "invalid"))
		ctx, cancel := context.WithTimeout(ctx, ctxTimeout)
		defer cancel()

		req := &serverpb.GetServerInfoRequest{}
		res, err := client.GetServerInfo(ctx, req)
		require.Error(t, err)
		require.Nil(t, res)

		expectedOk := true
		expectedCode := codes.Unauthenticated
		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, expectedOk, ok, "expected ok")
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})
}
