package integration

import (
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/hollowdll/kvdb/cmd/kvdbserver/config"
	grpcserver "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/internal/testutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

const ctxTimeout = time.Second * 5

var (
	testServerPort    = 0
	tlsTestServerPort = 0
)

func TestMain(m *testing.M) {
	grpcServer, port := startTestServer(defaultConfig())
	testServerPort = port
	defer grpcServer.Stop()
	tlsGrpcServer, port := startTestServer(tlsConfig())
	tlsTestServerPort = port
	defer tlsGrpcServer.Stop()
	code := m.Run()
	os.Exit(code)
}

func startTestServer(cfg config.ServerConfig) (*grpc.Server, int) {
	s := server.NewKvdbServer(cfg, testutil.DisabledLogger())
	s.CreateDefaultDatabase(cfg.DefaultDB)
	gs := grpcserver.SetupGrpcServer(s)

	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to listen: %v\n", err)
	}
	port := lis.Addr().(*net.TCPAddr).Port
	connLis := server.NewClientConnListener(lis, s, cfg.MaxClientConnections)
	s.ClientConnListener = connLis
	fmt.Fprintf(os.Stderr, "test server listening at %v\n", lis.Addr())

	go func() {
		if err := gs.Serve(connLis); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to serve gRPC: %v\n", err)
		}
	}()

	return gs, port
}

func defaultConfig() config.ServerConfig {
	return config.DefaultConfig()
}

func tlsConfig() config.ServerConfig {
	tlsCertPath, err := filepath.Abs("../../tls/test-cert/kvdbserver.crt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get TLS certificate path: %v\n", err)
	}
	tlsPrivKeyPath, err := filepath.Abs("../../tls/test-cert/kvdbserver.key")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get TLS private key path: %v\n", err)
	}
	cfg := config.DefaultConfig()
	cfg.TLSEnabled = true
	cfg.TLSCertPath = tlsCertPath
	cfg.TLSPrivKeyPath = tlsPrivKeyPath
	return cfg
}

func getServerAddress() string {
	return fmt.Sprintf("localhost:%d", testServerPort)
}

func getTlsServerAddress() string {
	return fmt.Sprintf("localhost:%d", tlsTestServerPort)
}

func insecureConnection() (*grpc.ClientConn, error) {
	return grpc.NewClient(getServerAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func secureConnection() (*grpc.ClientConn, error) {
	certBytes, err := os.ReadFile("../../tls/test-cert/kvdbserver.crt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read TLS certificate: %v\n", err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(certBytes) {
		fmt.Fprint(os.Stderr, "Failed to parse certificate\n")
	}

	creds := credentials.NewClientTLSFromCert(certPool, "")
	return grpc.NewClient(getTlsServerAddress(), grpc.WithTransportCredentials(creds))
}
