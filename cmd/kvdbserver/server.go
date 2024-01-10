package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"

	kvdb "github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/hollowdll/kvdb/version"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	kvdbserver.UnimplementedDatabaseServiceServer
	kvdbserver.UnimplementedServerServiceServer
	kvdbserver.UnimplementedStorageServiceServer
	startTime time.Time
	databases map[string]*kvdb.Database
	logger    kvdb.Logger
	mutex     sync.RWMutex
}

func newServer() *server {
	return &server{
		startTime: time.Now(),
		databases: make(map[string]*kvdb.Database),
		logger:    kvdb.NewDefaultLogger(),
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

// getOsInfo returns information about the server's operating system.
func getOsInfo() (string, error) {
	osInfo := runtime.GOOS

	switch osInfo {
	case "linux":
		cmd := exec.Command("uname", "-r", "-m")
		output, err := cmd.Output()

		if err != nil {
			return "", err
		}

		return "Linux " + strings.TrimSpace(string(output)), nil
	case "windows":
		cmd := exec.Command("cmd", "/c", "ver")
		output, err := cmd.Output()

		if err != nil {
			return "", err
		}

		return strings.TrimSpace(string(output)), nil
	default:
		return osInfo, nil
	}
}

// GetServerInfo returns information about the server.
func (s *server) GetServerInfo(ctx context.Context, req *kvdbserver.GetServerInfoRequest) (res *kvdbserver.GetServerInfoResponse, err error) {
	s.logger.Debug("Attempt to get server info")
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to get server info: %s", err)
		} else {
			s.logger.Debug("Get server info success")
		}
	}()

	osInfo, err := getOsInfo()
	if err != nil {
		errMsg := fmt.Sprintf("%s", err)
		return nil, status.Error(codes.Internal, errMsg)
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	info := &kvdbserver.ServerInfo{
		Version:       version.Version,
		GoVersion:     runtime.Version(),
		DbCount:       uint32(len(s.databases)),
		TotalDataSize: s.getTotalDataSize(),
		Os:            osInfo,
		Arch:          runtime.GOARCH,
		ProcessId:     uint32(os.Getpid()),
		UptimeSeconds: uint64(time.Since(s.startTime).Seconds()),
		TcpPort:       uint32(common.ServerDefaultPort),
	}

	return &kvdbserver.GetServerInfoResponse{Data: info}, nil
}
