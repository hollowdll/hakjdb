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

// Setups server for tests.
func setupServer() *grpc.Server {
	server := kvdbs.NewServer()
	server.DisableLogger()

	viper.SetDefault("port", common.ServerDefaultPort)
	viper.SetDefault("host", common.ServerDefaultHost)
	viper.SetEnvPrefix(kvdbs.EnvPrefix)
	viper.AutomaticEnv()

	grpcServer := grpc.NewServer()
	kvdbserver.RegisterDatabaseServiceServer(grpcServer, server)
	kvdbserver.RegisterServerServiceServer(grpcServer, server)
	kvdbserver.RegisterStorageServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetUint16("port")))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to listen: %v\n", err)
	}

	// Run in goroutine so execution won't be blocked.
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to serve gRPC: %v\n", err)
		}
	}()

	return grpcServer
}

func getServerAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetUint16("port"))
}
