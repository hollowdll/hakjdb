package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"

	kvdb "github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/hollowdll/kvdb/version"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	kvdbserver.UnimplementedDatabaseServiceServer
	kvdbserver.UnimplementedServerServiceServer
	kvdbserver.UnimplementedStorageServiceServer
	databases map[string]*kvdb.Database
	logger    kvdb.Logger
	mutex     sync.RWMutex
}

func newServer() *server {
	return &server{
		databases: make(map[string]*kvdb.Database),
		logger:    *kvdb.NewLogger(),
	}
}

// getTotalDataSize returns the total amount of stored data on this server in bytes.
func (s *server) getTotalDataSize() uint64 {
	var sum uint64
	for _, db := range s.databases {
		sum += db.GetStoredSizeBytes()
	}

	return sum
}

func (s *server) GetServerInfo(ctx context.Context, req *kvdbserver.GetServerInfoRequest) (*kvdbserver.GetServerInfoResponse, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var osInfo string
	osInfo += runtime.GOOS + " " + runtime.GOARCH

	hostname, err := os.Hostname()
	if err != nil {
		errMsg := fmt.Sprintf("%s", err)
		return nil, status.Error(codes.Internal, errMsg)
	}
	osInfo += " " + hostname

	info := &kvdbserver.ServerInfo{
		Version:       version.Version,
		GoVersion:     runtime.Version(),
		DbCount:       uint32(len(s.databases)),
		TotalDataSize: s.getTotalDataSize(),
		Os:            osInfo,
	}

	return &kvdbserver.GetServerInfoResponse{Info: info}, nil
}
