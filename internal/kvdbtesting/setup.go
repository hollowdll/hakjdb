package kvdbtesting

import (
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
)

// SetupServerWithLoggerDisabled returns a new server with logger disabled
func SetupServerWithLoggerDisabled() *server.Server {
	server := server.NewServer()
	server.DisableLogger()

	return server
}
