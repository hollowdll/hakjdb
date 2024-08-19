package integration

import (
	"context"
	"testing"

	"github.com/hollowdll/kvdb/api/v0/authpb"
	"github.com/hollowdll/kvdb/api/v0/echopb"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestAuth(t *testing.T) {
	cfg := testutil.DefaultConfig()
	cfg.AuthEnabled = true
	cfg.AuthTokenSecretKey = "test-key"
	cfg.AuthTokenTTL = 30
	password := "pass123"
	s, gs, port := testutil.StartTestServer(cfg)
	defer testutil.StopTestServer(gs)
	s.EnableAuth(password)
	address := testutil.GetServerAddress(port)

	t.Run("AuthenticateWithValidCredentials", func(t *testing.T) {
		conn, err := testutil.InsecureConnection(address)
		require.NoErrorf(t, err, "expected connection but connection failed: %v", err)
		defer conn.Close()
		authClient := authpb.NewAuthServiceClient(conn)
		echoClient := echopb.NewEchoServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
		defer cancel()

		authReq := &authpb.AuthenticateRequest{Password: password}
		authRes, err := authClient.Authenticate(ctx, authReq)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, authRes)
		assert.NotEqual(t, "", authRes.AuthToken, "expected auth token not to be empty")

		ctx, cancel = context.WithTimeout(
			metadata.NewOutgoingContext(
				context.Background(),
				metadata.Pairs(common.GrpcMetadataKeyAuthToken, authRes.AuthToken),
			),
			ctxTimeout,
		)
		defer cancel()
		echoReq := &echopb.UnaryEchoRequest{}
		echoRes, err := echoClient.UnaryEcho(ctx, echoReq)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, echoRes)
	})

	t.Run("AuthenticateWithInvalidCredentials", func(t *testing.T) {
		conn, err := testutil.InsecureConnection(address)
		require.NoErrorf(t, err, "expected connection but connection failed: %v", err)
		defer conn.Close()
		authClient := authpb.NewAuthServiceClient(conn)
		echoClient := echopb.NewEchoServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
		defer cancel()

		authReq := &authpb.AuthenticateRequest{Password: "invalid-password"}
		authRes, err := authClient.Authenticate(ctx, authReq)
		require.Error(t, err)
		require.Nil(t, authRes)

		echoReq := &echopb.UnaryEchoRequest{}
		echoRes, err := echoClient.UnaryEcho(ctx, echoReq)
		require.Error(t, err)
		require.Nil(t, echoRes)

		expectedOk := true
		expectedCode := codes.Unauthenticated
		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, expectedOk, ok, "expected ok")
		assert.Equal(t, expectedCode, st.Code(), "expected status = %s; got = %s", expectedCode, st.Code())
	})
}
