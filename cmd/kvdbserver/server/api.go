package server

import (
	"context"
	"os"
	"runtime"
	"time"

	"github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/api/v0/serverpb"
	kvdberrors "github.com/hollowdll/kvdb/errors"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/version"
)

type ServerService interface {
	Logger() kvdb.Logger
	GetServerInfo(ctx context.Context, req *serverpb.GetServerInfoRequest) (res *serverpb.GetServerInfoResponse, err error)
	GetLogs(ctx context.Context, req *serverpb.GetLogsRequest) (res *serverpb.GetLogsResponse, err error)
}

type DBService interface {
}

type GeneralKeyService interface {
}

type StringKeyService interface {
}

type HashMapKeyService interface {
}

func (s *KvdbServer) GetServerInfo(ctx context.Context, req *serverpb.GetServerInfoRequest) (res *serverpb.GetServerInfoResponse, err error) {
	logger := s.Logger()
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.ClientConnListener != nil {
		s.ClientConnListener.mu.RLock()
		defer s.ClientConnListener.mu.RUnlock()
	}

	osInfo, err := getOsInfo()
	if err != nil {
		logger.Errorf("%v: %v", kvdberrors.ErrGetOSInfo, err)
		return nil, kvdberrors.ErrGetOSInfo
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	var totalKeys uint64
	for _, db := range s.databases {
		totalKeys += uint64(db.GetKeyCount())
	}

	generalInfo := &serverpb.GeneralInfo{
		KvdbVersion:     version.Version,
		GoVersion:       runtime.Version(),
		Os:              osInfo,
		Arch:            runtime.GOARCH,
		ProcessId:       uint32(os.Getpid()),
		UptimeSeconds:   uint64(time.Since(s.startTime).Seconds()),
		TcpPort:         uint32(s.Cfg.PortInUse),
		TlsEnabled:      s.Cfg.TLSEnabled,
		PasswordEnabled: s.credentialStore.IsServerPasswordEnabled(),
		LogfileEnabled:  s.Cfg.LogFileEnabled,
		DebugEnabled:    s.Cfg.DebugEnabled,
	}
	memoryInfo := &serverpb.MemoryInfo{
		MemoryAlloc:      m.Alloc,
		MemoryTotalAlloc: m.TotalAlloc,
		MemorySys:        m.Sys,
	}
	storageInfo := &serverpb.StorageInfo{
		TotalDataSize: s.getTotalDataSize(),
		TotalKeys:     totalKeys,
	}
	clientInfo := &serverpb.ClientInfo{
		ClientConnections:    s.ClientConnListener.clientConnections,
		MaxClientConnections: s.ClientConnListener.maxClientConnections,
	}
	dbInfo := &serverpb.DatabaseInfo{
		DbCount:   uint32(len(s.databases)),
		DefaultDb: s.Cfg.DefaultDB,
	}

	return &serverpb.GetServerInfoResponse{
		GeneralInfo: generalInfo,
		MemoryInfo:  memoryInfo,
		StorageInfo: storageInfo,
		ClientInfo:  clientInfo,
		DbInfo:      dbInfo,
	}, nil
}

func (s *KvdbServer) GetLogs(ctx context.Context, req *serverpb.GetLogsRequest) (res *serverpb.GetLogsResponse, err error) {
	logger := s.Logger()
	s.mu.RLock()
	defer s.mu.RUnlock()

	if !s.Cfg.LogFileEnabled {
		return nil, kvdberrors.ErrLogFileNotEnabled
	}

	logs, err := common.ReadFileLines(s.Cfg.LogFilePath)
	if err != nil {
		logger.Errorf("%v: %v", kvdberrors.ErrReadLogFile, err)
		return nil, kvdberrors.ErrReadLogFile
	}

	return &serverpb.GetLogsResponse{Logs: logs}, nil
}
