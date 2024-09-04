package testutil

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/hollowdll/hakjdb"

	"github.com/hollowdll/hakjdb/cmd/hakjserver/config"
	grpcserver "github.com/hollowdll/hakjdb/cmd/hakjserver/grpc"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func StartTestServer(cfg config.ServerConfig) (*server.HakjServer, *grpc.Server, int) {
	fmt.Fprint(os.Stderr, "creating test server\n")
	s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
	s.CreateDefaultDatabase(cfg.DefaultDB)
	gs := grpcserver.SetupGrpcServer(s)

	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to listen: %v\n", err)
	}
	port := lis.Addr().(*net.TCPAddr).Port
	s.Cfg.PortInUse = uint16(port)
	connLis := server.NewClientConnListener(lis, s, cfg.MaxClientConnections)
	s.ClientConnListener = connLis
	fmt.Fprintf(os.Stderr, "test server listening at %v\n", lis.Addr())

	go func() {
		if err := gs.Serve(connLis); err != nil {
			fmt.Fprintf(os.Stderr, "failed to serve gRPC: %v\n", err)
		}
	}()

	return s, gs, port
}

func StopTestServer(gs *grpc.Server) {
	fmt.Fprint(os.Stderr, "stopping test server\n")
	gs.Stop()
}

func DefaultConfig() config.ServerConfig {
	return config.DefaultConfig()
}

func TLSConfig(clientCertAuth bool) config.ServerConfig {
	tlsCertPath, err := filepath.Abs("../../tls/test-cert/hakjserver-cert.pem")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get TLS certificate path: %v\n", err)
	}
	tlsPrivKeyPath, err := filepath.Abs("../../tls/test-cert/hakjserver-key.pem")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get TLS private key path: %v\n", err)
	}
	tlsCACertPath, err := filepath.Abs("../../tls/test-cert/ca-cert.pem")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get TLS CA certificate path: %v\n", err)
	}
	cfg := config.DefaultConfig()
	cfg.TLSEnabled = true
	cfg.TLSCertPath = tlsCertPath
	cfg.TLSPrivKeyPath = tlsPrivKeyPath
	cfg.TLSCACertPath = tlsCACertPath
	cfg.TLSClientCertAuthEnabled = clientCertAuth
	return cfg
}

func GetServerAddress(port int) string {
	return fmt.Sprintf("localhost:%d", port)
}

func InsecureConnection(address string) (*grpc.ClientConn, error) {
	return grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func SecureConnection(address string, clientCertAuth bool) (*grpc.ClientConn, error) {
	certBytes, err := os.ReadFile("../../tls/test-cert/ca-cert.pem")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read TLS CA certificate: %v\n", err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(certBytes) {
		fmt.Fprint(os.Stderr, "failed to parse TLS CA certificate\n")
	}

	var certs []tls.Certificate
	if clientCertAuth {
		clientCert, err := tls.LoadX509KeyPair("../../tls/test-cert/client-cert.pem", "../../tls/test-cert/client-key.pem")
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to load TLS client public/private key pair: %v\n", err)
		}
		certs = append(certs, clientCert)
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: certs,
		RootCAs:      certPool,
	})
	return grpc.NewClient(address, grpc.WithTransportCredentials(creds))
}
