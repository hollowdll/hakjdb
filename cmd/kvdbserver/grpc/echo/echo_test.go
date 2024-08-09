package echo

import (
	"context"
	"testing"

	"github.com/hollowdll/kvdb/api/v0/echopb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnaryEcho(t *testing.T) {
	t.Run("EmptyMessage", func(t *testing.T) {
		gs := NewEchoServiceServer()
		req := &echopb.UnaryEchoRequest{Msg: ""}
		resp, err := gs.UnaryEcho(context.Background(), req)

		expectedMsg := ""
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, resp)
		assert.Equalf(t, expectedMsg, resp.Msg, "expected msg = %s; got = %s", expectedMsg, resp.Msg)
	})

	t.Run("NonEmptyMessage", func(t *testing.T) {
		gs := NewEchoServiceServer()
		req := &echopb.UnaryEchoRequest{Msg: "Hello?"}
		resp, err := gs.UnaryEcho(context.Background(), req)

		expectedMsg := "Hello?"
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, resp)
		assert.Equalf(t, expectedMsg, resp.Msg, "expected msg = %s; got = %s", expectedMsg, resp.Msg)
	})
}
