package benchmark

import (
	"context"
	"testing"
	"time"

	"github.com/hollowdll/kvdb/api/v0/kvpb"
	"github.com/hollowdll/kvdb/internal/testutil"
)

const timeout = time.Second * 3

func BenchmarkSetString(b *testing.B) {
	cfg := testutil.DefaultConfig()
	_, gs, port := testutil.StartTestServer(cfg)
	defer testutil.StopTestServer(gs)
	address := testutil.GetServerAddress(port)
	conn, err := testutil.InsecureConnection(address)
	if err != nil {
		b.Fatalf("setting up connection failed: %v", err)
	}
	defer conn.Close()
	client := kvpb.NewStringKVServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := &kvpb.SetStringRequest{
				Key:   "key",
				Value: []byte("value"),
			}
			client.SetString(ctx, req)
		}
	})
}

func BenchmarkGetString(b *testing.B) {
	cfg := testutil.DefaultConfig()
	_, gs, port := testutil.StartTestServer(cfg)
	defer testutil.StopTestServer(gs)
	address := testutil.GetServerAddress(port)
	conn, err := testutil.InsecureConnection(address)
	if err != nil {
		b.Fatalf("setting up connection failed: %v", err)
	}
	defer conn.Close()
	client := kvpb.NewStringKVServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	reqSet := &kvpb.SetStringRequest{
		Key:   "key",
		Value: []byte("value"),
	}
	_, err = client.SetString(ctx, reqSet)
	if err != nil {
		b.Fatalf("failed to set string KV: %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := &kvpb.GetStringRequest{
				Key: "key",
			}
			client.GetString(ctx, req)
		}
	})
}
