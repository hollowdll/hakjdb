package main_test

import (
	"context"
	"testing"

	main "github.com/hollowdll/kvdb/cmd/kvdbserver"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/stretchr/testify/assert"
)

func TestGetServerInfo(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		server := main.NewServer()
		server.DisableLogger()

		request := &kvdbserver.GetServerInfoRequest{}
		response, err := server.GetServerInfo(context.Background(), request)
		assert.NoErrorf(t, err, "expected no error; error = %s", err)
		assert.NotNil(t, response, "expected response to be non-nil")
	})
}
