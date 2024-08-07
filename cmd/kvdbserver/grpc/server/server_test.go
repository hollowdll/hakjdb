package server

import (
	"context"
	"testing"

	"github.com/hollowdll/kvdb/api/v0/serverpb"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/config"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/internal/testutil"

	"github.com/hollowdll/kvdb/proto/kvdbserverpb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetServerInfo(t *testing.T) {
	cfg := config.ServerConfig{
		DefaultDB:            "default",
		MaxClientConnections: 1000,
	}

	t.Run("Success", func(t *testing.T) {
		s := server.NewKvdbServer(cfg, testutil.DisabledLogger())
		connLis := server.NewClientConnListener(nil, s, cfg.MaxClientConnections)
		s.ClientConnListener = connLis

		req := &serverpb.GetServerInfoRequest{}
		resp, err := s.GetServerInfo(context.Background(), req)
		assert.NoErrorf(t, err, "expected no error; error = %v", err)
		assert.NotNil(t, resp)
	})
}

func TestGetLogs(t *testing.T) {
	t.Run("LogFileNotEnabled", func(t *testing.T) {
		server := server.NewServer()
		server.DisableLogger()

		req := &kvdbserverpb.GetLogsRequest{}
		res, err := server.GetLogs(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, res)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.FailedPrecondition, st.Code(), "expected status = %s; got = %s", codes.FailedPrecondition, st.Code())
	})

	t.Run("MultipleLogs", func(t *testing.T) {
		logFilePath := "testdata/multiline_log.testlog"
		server := server.NewServer()
		server.SetLogFilePath(logFilePath)
		server.EnableLogFile()
		server.DisableLogger()

		lines, err := common.ReadFileLines(logFilePath)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, lines)

		req := &kvdbserverpb.GetLogsRequest{}
		res, err := server.GetLogs(context.Background(), req)
		expectedLogs := 4
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		require.Equal(t, expectedLogs, len(res.Logs), "expected logs = %d; got = %d", expectedLogs, len(res.Logs))
		require.Equal(t, true, res.LogfileEnabled, "expected LogfileEnabled = %v; got = %v", true, res.LogfileEnabled)

		for i, log := range res.Logs {
			assert.Equal(t, lines[i], log, "expected log = %s; got = %s", lines[i], log)
		}
	})

	t.Run("NoLogs", func(t *testing.T) {
		logFilePath := "testdata/empty_log.testlog"
		server := server.NewServer()
		server.SetLogFilePath(logFilePath)
		server.EnableLogFile()
		server.DisableLogger()

		req := &kvdbserverpb.GetLogsRequest{}
		res, err := server.GetLogs(context.Background(), req)
		expectedLogs := 0
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, res)
		require.Equal(t, expectedLogs, len(res.Logs), "expected logs = %d; got = %d", expectedLogs, len(res.Logs))
		require.Equal(t, true, res.LogfileEnabled, "expected LogfileEnabled = %v; got = %v", true, res.LogfileEnabled)
	})
}
