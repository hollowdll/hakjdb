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
	"github.com/hollowdll/kvdb/proto/kvdbserverpb"
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
	server, port := startTestServer(1000)
	defer server.Stop()
	testServerPort = port
	tlsServer, port := startTlsTestServer(1000)
	defer tlsServer.Stop()
	tlsTestServerPort = port

	code := m.Run()
	os.Exit(code)
}

func startTestServer(maxConnections uint32) (*grpc.Server, int) {
	server := kvdbs.NewServer()
	server.DisableLogger()
	server.CreateDefaultDatabase("default")

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(server.AuthInterceptor))
	kvdbserverpb.RegisterDatabaseServiceServer(grpcServer, server)
	kvdbserverpb.RegisterServerServiceServer(grpcServer, server)
	kvdbserverpb.RegisterStorageServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to listen: %v\n", err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	server.SetPort(uint16(port))
	connListener := kvdbs.NewClientConnListener(listener, server, maxConnections)
	server.ClientConnListener = connListener
	fmt.Fprintf(os.Stderr, "startTestServer listening at %v\n", listener.Addr())

	go func() {
		if err := grpcServer.Serve(connListener); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to serve gRPC: %v\n", err)
		}
	}()

	return grpcServer, port
}

func startTlsTestServer(maxConnections uint32) (*grpc.Server, int) {
	server := kvdbs.NewServer()
	server.DisableLogger()
	server.CreateDefaultDatabase("default")

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

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to listen: %v\n", err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	server.SetPort(uint16(port))
	connListener := kvdbs.NewClientConnListener(listener, server, maxConnections)
	server.ClientConnListener = connListener
	fmt.Fprintf(os.Stderr, "startTlsTestServer listening at %v\n", listener.Addr())

	go func() {
		if err := grpcServer.Serve(connListener); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to serve gRPC: %v\n", err)
		}
	}()

	return grpcServer, port
}

func getServerAddress() string {
	return fmt.Sprintf("localhost:%d", testServerPort)
}

func getTlsServerAddress() string {
	return fmt.Sprintf("localhost:%d", tlsTestServerPort)
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
