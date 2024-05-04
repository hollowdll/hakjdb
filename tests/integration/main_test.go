package integration

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	kvdbs "github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserverpb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

const ctxTimeout = time.Second * 5
const defaultTlsPort = 12346

func TestMain(m *testing.M) {
	viper.SetEnvPrefix(kvdbs.EnvPrefix)
	viper.AutomaticEnv()

	fmt.Fprintln(os.Stderr, "setting up test server ...")
	server := setupServer()
	fmt.Fprintln(os.Stderr, "setting up TLS test server ...")
	tlsServer := setupTlsServer()
	defer server.Stop()
	defer tlsServer.Stop()

	code := m.Run()
	os.Exit(code)
}

func setupServer() *grpc.Server {
	server := kvdbs.NewServer()
	server.DisableLogger()
	server.CreateDefaultDatabase("default")

	viper.SetDefault("port", common.ServerDefaultPort)
	server.SetPort(viper.GetUint16("port"))

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(server.AuthInterceptor))
	kvdbserverpb.RegisterDatabaseServiceServer(grpcServer, server)
	kvdbserverpb.RegisterServerServiceServer(grpcServer, server)
	kvdbserverpb.RegisterStorageServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetUint16("port")))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to listen: %v\n", err)
	}
	fmt.Fprintf(os.Stderr, "test server listening at %v\n", listener.Addr())

	// Run in goroutine so execution won't be blocked.
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to serve gRPC: %v\n", err)
		}
	}()

	return grpcServer
}

func setupTlsServer() *grpc.Server {
	server := kvdbs.NewServer()
	server.DisableLogger()
	server.CreateDefaultDatabase("default")

	viper.SetDefault("tls_port", defaultTlsPort)
	server.SetPort(viper.GetUint16("tls_port"))

	certBytes, err := os.ReadFile("../../tls/test-cert/kvdbserver.crt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read TLS certificate: %v\n", err)
	}
	keyBytes, err := os.ReadFile("../../tls/test-cert/kvdbserver.key")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read TLS private key: %v\n", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(certBytes) {
		fmt.Fprint(os.Stderr, "Failed to parse certificate\n")
	}
	cert, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse TLS public/private key pair: %v\n", err)
	}

	creds := credentials.NewServerTLSFromCert(&cert)
	grpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(server.AuthInterceptor))

	kvdbserverpb.RegisterDatabaseServiceServer(grpcServer, server)
	kvdbserverpb.RegisterServerServiceServer(grpcServer, server)
	kvdbserverpb.RegisterStorageServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetUint16("tls_port")))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to listen: %v\n", err)
	}
	fmt.Fprintf(os.Stderr, "TLS test server listening at %v\n", listener.Addr())

	// Run in goroutine so execution won't be blocked.
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to serve gRPC: %v\n", err)
		}
	}()

	return grpcServer
}

func getServerAddress() string {
	return fmt.Sprintf("%s:%d", common.ServerDefaultHost, viper.GetUint16("port"))
}

func getTlsServerAddress() string {
	return fmt.Sprintf("%s:%d", common.ServerDefaultHost, viper.GetUint16("tls_port"))
}

func insecureConnection() (*grpc.ClientConn, error) {
	return grpc.Dial(getServerAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	return grpc.Dial(getTlsServerAddress(), grpc.WithTransportCredentials(creds))
}
