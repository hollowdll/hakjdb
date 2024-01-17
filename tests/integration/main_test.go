package integration

import (
	"fmt"
	"net"
	"os"
	"testing"

	kvdbs "github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func TestMain(m *testing.M) {
	server := setupServer()
	defer server.Stop()
	code := m.Run()
	os.Exit(code)
}

func setupServer() *grpc.Server {
	server := kvdbs.NewServer()
	server.DisableLogger()

	viper.SetDefault(kvdbs.ConfigKeyPort, common.ServerDefaultPort)
	viper.SetEnvPrefix(kvdbs.EnvPrefix)
	viper.AutomaticEnv()

	grpcServer := grpc.NewServer()
	kvdbserver.RegisterDatabaseServiceServer(grpcServer, server)
	kvdbserver.RegisterServerServiceServer(grpcServer, server)
	kvdbserver.RegisterStorageServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetUint16(kvdbs.ConfigKeyPort)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to listen: %v\n", err)
	}

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to serve gRPC: %v\n", err)
		}
	}()

	return grpcServer
}
