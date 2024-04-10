package client

import (
	"fmt"
	"os"
	"time"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/config"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserverpb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const (
	// CtxTimeout specifies the number of seconds to wait until operation terminates.
	CtxTimeout time.Duration = time.Second * 10
	// ValueNone is a special value for values that do not exist.
	ValueNone string = "(None)"
)

var (
	GrpcDatabaseClient   kvdbserverpb.DatabaseServiceClient
	GrpcStorageClient    kvdbserverpb.StorageServiceClient
	GrpcServerClient     kvdbserverpb.ServerServiceClient
	grpcClientConnection *grpc.ClientConn
)

// InitClient initializes the client and connections.
func InitClient() {
	address := fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetUint16("port"))
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		cobra.CheckErr(fmt.Sprintf("failed to connect to the server: %s", err))
	}

	GrpcDatabaseClient = kvdbserverpb.NewDatabaseServiceClient(conn)
	GrpcStorageClient = kvdbserverpb.NewStorageServiceClient(conn)
	GrpcServerClient = kvdbserverpb.NewServerServiceClient(conn)
	grpcClientConnection = conn
}

// CloseConnections closes all connections to the server.
func CloseConnections() {
	if grpcClientConnection != nil {
		grpcClientConnection.Close()
	}
}

// ReadPasswordFromEnv reads password from environment variable.
// The returned bool is true if it is present.
func ReadPasswordFromEnv() (string, bool) {
	return os.LookupEnv(config.EnvVarPassword)
}

// GetBaseGrpcMetadata returns base gRPC metadata for all requests.
// It can be overwritten or extended.
func GetBaseGrpcMetadata() metadata.MD {
	md := metadata.Pairs()
	password, ok := ReadPasswordFromEnv()
	if ok {
		md.Set(common.GrpcMetadataKeyPassword, password)
	}

	dbName := viper.GetString(config.ConfigKeyDatabase)
	md.Set(common.GrpcMetadataKeyDbName, dbName)

	return md
}
