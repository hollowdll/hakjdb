package client

import (
	"crypto/x509"
	"fmt"
	"os"
	"time"

	"github.com/hollowdll/kvdb/api/v0/dbpb"
	"github.com/hollowdll/kvdb/api/v0/serverpb"
	"github.com/hollowdll/kvdb/api/v0/storagepb"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/config"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	GrpcServerClient     serverpb.ServerServiceClient
	GrpcDatabaseClient   dbpb.DatabaseServiceClient
	GrpcGeneralKeyClient storagepb.GeneralKeyServiceClient
	GrpcStringKeyClient  storagepb.StringKeyServiceClient
	GrpcHashMapKeyClient storagepb.HashMapKeyServiceClient
	grpcClientConn       *grpc.ClientConn
)

// InitClient initializes the client and connections.
func InitClient() {
	var dialOption grpc.DialOption = nil
	if viper.GetBool(config.ConfigKeyTlsEnabled) {
		certBytes, err := os.ReadFile(viper.GetString(config.ConfigKeyTlsCertPath))
		if err != nil {
			cobra.CheckErr(fmt.Sprintf("failed to read TLS certificate: %v", err))
		}
		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(certBytes) {
			cobra.CheckErr("failed to parse TLS certificate")
		}

		creds := credentials.NewClientTLSFromCert(certPool, "")
		dialOption = grpc.WithTransportCredentials(creds)
	} else {
		dialOption = grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	address := fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetUint16("port"))
	conn, err := grpc.Dial(address, dialOption)
	if err != nil {
		cobra.CheckErr(fmt.Sprintf("failed to connect to the server: %s", err))
	}

	GrpcServerClient = serverpb.NewServerServiceClient(conn)
	GrpcDatabaseClient = dbpb.NewDatabaseServiceClient(conn)
	GrpcGeneralKeyClient = storagepb.NewGeneralKeyServiceClient(conn)
	GrpcStringKeyClient = storagepb.NewStringKeyServiceClient(conn)
	GrpcHashMapKeyClient = storagepb.NewHashMapKeyServiceClient(conn)
	grpcClientConn = conn
}

// CloseConnections closes all connections to the server.
func CloseConnections() {
	if grpcClientConn != nil {
		if err := grpcClientConn.Close(); err != nil {
			cobra.CheckErr(fmt.Sprintf("failed to close connections: %v", err))
		}
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
