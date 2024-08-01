package server

import (
	"context"
	"os"
	"runtime"
	"time"

	"github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/api/v0/dbpb"
	"github.com/hollowdll/kvdb/api/v0/kvpb"
	"github.com/hollowdll/kvdb/api/v0/serverpb"
	kvdberrors "github.com/hollowdll/kvdb/errors"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/version"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServerService interface {
	Logger() kvdb.Logger
	GetServerInfo(ctx context.Context, req *serverpb.GetServerInfoRequest) (*serverpb.GetServerInfoResponse, error)
	GetLogs(ctx context.Context, req *serverpb.GetLogsRequest) (*serverpb.GetLogsResponse, error)
}

type DBService interface {
	Logger() kvdb.Logger
	CreateDB(ctx context.Context, req *dbpb.CreateDBRequest) (*dbpb.CreateDBResponse, error)
	DeleteDB(ctx context.Context, req *dbpb.DeleteDBRequest) (*dbpb.DeleteDBResponse, error)
	GetAllDBs(ctx context.Context, req *dbpb.GetAllDBsRequest) (*dbpb.GetAllDBsResponse, error)
	GetDBInfo(ctx context.Context, req *dbpb.GetDBInfoRequest) (*dbpb.GetDBInfoResponse, error)
}

type GeneralKVService interface {
	Logger() kvdb.Logger
	GetDBNameFromContext(ctx context.Context) string
	GetAllKeys(ctx context.Context, req *kvpb.GetAllKeysRequest) (*kvpb.GetAllKeysResponse, error)
	GetKeyType(ctx context.Context, req *kvpb.GetKeyTypeRequest) (*kvpb.GetKeyTypeResponse, error)
	DeleteKeys(ctx context.Context, req *kvpb.DeleteKeysRequest) (*kvpb.DeleteKeysResponse, error)
	DeleteAllKeys(ctx context.Context, req *kvpb.DeleteAllKeysRequest) (*kvpb.DeleteAllKeysResponse, error)
}

type StringKVService interface {
	Logger() kvdb.Logger
	GetDBNameFromContext(ctx context.Context) string
	SetString(ctx context.Context, req *kvpb.SetStringRequest) (*kvpb.SetStringResponse, error)
	GetString(ctx context.Context, req *kvpb.GetStringRequest) (*kvpb.GetStringResponse, error)
}

type HashMapKVService interface {
	Logger() kvdb.Logger
	GetDBNameFromContext(ctx context.Context) string
	SetHashMap(ctx context.Context, req *kvpb.SetHashMapRequest) (*kvpb.SetHashMapResponse, error)
	GetHashMapFieldValues(ctx context.Context, req *kvpb.GetHashMapFieldValuesRequest) (*kvpb.GetHashMapFieldValuesResponse, error)
	GetAllHashMapFieldsAndValues(ctx context.Context, req *kvpb.GetAllHashMapFieldsAndValuesRequest) (*kvpb.GetAllHashMapFieldsAndValuesResponse, error)
	DeleteHashMapFields(ctx context.Context, req *kvpb.DeleteHashMapFieldsRequest) (*kvpb.DeleteHashMapFieldsResponse, error)
}

func (s *KvdbServer) GetServerInfo(ctx context.Context, req *serverpb.GetServerInfoRequest) (*serverpb.GetServerInfoResponse, error) {
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
		ClientConnections:    s.ClientConnListener.clientConnectionsCount,
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

func (s *KvdbServer) GetLogs(ctx context.Context, req *serverpb.GetLogsRequest) (*serverpb.GetLogsResponse, error) {
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

func (s *KvdbServer) CreateDB(ctx context.Context, req *dbpb.CreateDBRequest) (*dbpb.CreateDBResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.dbExists(req.DbName) {
		return nil, kvdberrors.ErrDatabaseExists
	}

	db := kvdb.CreateDatabase(req.DbName)
	s.databases[db.Name] = db

	return &dbpb.CreateDBResponse{DbName: db.Name}, nil
}

func (s *KvdbServer) DeleteDB(ctx context.Context, req *dbpb.DeleteDBRequest) (*dbpb.DeleteDBResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.dbExists(req.DbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	delete(s.databases, req.DbName)

	return &dbpb.DeleteDBResponse{DbName: req.DbName}, nil
}

func (s *KvdbServer) GetAllDBs(ctx context.Context, req *dbpb.GetAllDBsRequest) (*dbpb.GetAllDBsResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var names []string
	for key := range s.databases {
		names = append(names, key)
	}

	return &dbpb.GetAllDBsResponse{DbNames: names}, nil
}

func (s *KvdbServer) GetDBInfo(ctx context.Context, req *dbpb.GetDBInfoRequest) (*dbpb.GetDBInfoResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if !s.dbExists(req.DbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	db := s.databases[req.DbName]
	data := &dbpb.DBInfo{
		Name:      db.Name,
		CreatedAt: timestamppb.New(db.CreatedAt),
		UpdatedAt: timestamppb.New(db.UpdatedAt),
		KeyCount:  db.GetKeyCount(),
		DataSize:  db.GetStoredSizeBytes(),
	}

	return &dbpb.GetDBInfoResponse{Data: data}, nil
}

func (s *KvdbServer) GetAllKeys(ctx context.Context, req *kvpb.GetAllKeysRequest) (*kvpb.GetAllKeysResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	return &kvpb.GetAllKeysResponse{Keys: s.databases[dbName].GetKeys()}, nil
}

func (s *KvdbServer) GetKeyType(ctx context.Context, req *kvpb.GetKeyTypeRequest) (*kvpb.GetKeyTypeResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	keyType, ok := s.databases[dbName].GetKeyType(req.Key)

	return &kvpb.GetKeyTypeResponse{KeyType: keyType.String(), Ok: ok}, nil
}

func (s *KvdbServer) DeleteKeys(ctx context.Context, req *kvpb.DeleteKeysRequest) (*kvpb.DeleteKeysResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	keysDeletedCount := s.databases[dbName].DeleteKeys(req.Keys)

	return &kvpb.DeleteKeysResponse{KeysDeletedCount: keysDeletedCount}, nil
}

func (s *KvdbServer) DeleteAllKeys(ctx context.Context, req *kvpb.DeleteAllKeysRequest) (*kvpb.DeleteAllKeysResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	s.databases[dbName].DeleteAllKeys()

	return &kvpb.DeleteAllKeysResponse{}, nil
}

func (s *KvdbServer) SetString(ctx context.Context, req *kvpb.SetStringRequest) (*kvpb.SetStringResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	if s.DBMaxKeysReached(s.databases[dbName]) {
		return nil, kvdberrors.ErrMaxKeysReached
	}

	s.databases[dbName].SetString(req.Key, req.Value)

	return &kvpb.SetStringResponse{}, nil
}

func (s *KvdbServer) GetString(ctx context.Context, req *kvpb.GetStringRequest) (*kvpb.GetStringResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	value, ok := s.databases[dbName].GetString(req.Key)

	return &kvpb.GetStringResponse{Value: value, Ok: ok}, nil
}

func (s *KvdbServer) SetHashMap(ctx context.Context, req *kvpb.SetHashMapRequest) (*kvpb.SetHashMapResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	if s.DBMaxKeysReached(s.databases[dbName]) {
		return nil, kvdberrors.ErrMaxKeysReached
	}

	fieldsAddedCount := s.databases[dbName].SetHashMap(req.Key, req.FieldValueMap, s.Cfg.MaxHashMapFields)

	return &kvpb.SetHashMapResponse{FieldsAddedCount: fieldsAddedCount}, nil
}

func (s *KvdbServer) GetHashMapFieldValues(ctx context.Context, req *kvpb.GetHashMapFieldValuesRequest) (res *kvpb.GetHashMapFieldValuesResponse, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	result, ok := s.databases[dbName].GetHashMapFieldValues(req.Key, req.Fields)

	var fieldValueMap = make(map[string]*kvpb.HashMapFieldValue)
	for field, value := range result {
		fieldValueMap[field] = &kvpb.HashMapFieldValue{
			Value: value.Value,
			Ok:    value.Ok,
		}
	}

	return &kvpb.GetHashMapFieldValuesResponse{FieldValueMap: fieldValueMap, Ok: ok}, nil
}

func (s *KvdbServer) GetAllHashMapFieldsAndValues(ctx context.Context, req *kvpb.GetAllHashMapFieldsAndValuesRequest) (res *kvpb.GetAllHashMapFieldsAndValuesResponse, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	fieldValueMap, ok := s.databases[dbName].GetAllHashMapFieldsAndValues(req.Key)

	return &kvpb.GetAllHashMapFieldsAndValuesResponse{FieldValueMap: fieldValueMap, Ok: ok}, nil
}

func (s *KvdbServer) DeleteHashMapFields(ctx context.Context, req *kvpb.DeleteHashMapFieldsRequest) (res *kvpb.DeleteHashMapFieldsResponse, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	fieldsRemovedCount, ok := s.databases[dbName].DeleteHashMapFields(req.Key, req.Fields)

	return &kvpb.DeleteHashMapFieldsResponse{FieldsRemovedCount: fieldsRemovedCount, Ok: ok}, nil
}
