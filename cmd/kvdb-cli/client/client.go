package client

import (
	"time"

	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"google.golang.org/grpc"
)

var (
	GrpcDatabaseClient kvdbserver.DatabaseClient
	// ClientCtxTimeout specifies how long to wait until operation terminates.
	ClientCtxTimeout = time.Second * 10
)

func InitClient(conn *grpc.ClientConn) {
	GrpcDatabaseClient = kvdbserver.NewDatabaseClient(conn)
}
