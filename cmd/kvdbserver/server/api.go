package server

import (
	"context"
	"os"
	"runtime"
	"time"

	"github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/api/v0/dbpb"
	"github.com/hollowdll/kvdb/api/v0/serverpb"
	"github.com/hollowdll/kvdb/api/v0/storagepb"
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
	CreateDatabase(ctx context.Context, req *dbpb.CreateDatabaseRequest) (*dbpb.CreateDatabaseResponse, error)
	DeleteDatabase(ctx context.Context, req *dbpb.DeleteDatabaseRequest) (*dbpb.DeleteDatabaseResponse, error)
	GetAllDatabases(ctx context.Context, req *dbpb.GetAllDatabasesRequest) (*dbpb.GetAllDatabasesResponse, error)
	GetDatabaseInfo(ctx context.Context, req *dbpb.GetDatabaseInfoRequest) (*dbpb.GetDatabaseInfoResponse, error)
}

type GeneralKeyService interface {
	Logger() kvdb.Logger
	GetDBNameFromContext(ctx context.Context) string
	GetAllKeys(ctx context.Context, req *storagepb.GetAllKeysRequest) (*storagepb.GetAllKeysResponse, error)
	GetKeyType(ctx context.Context, req *storagepb.GetKeyTypeRequest) (*storagepb.GetKeyTypeResponse, error)
	DeleteKeys(ctx context.Context, req *storagepb.DeleteKeysRequest) (*storagepb.DeleteKeysResponse, error)
	DeleteAllKeys(ctx context.Context, req *storagepb.DeleteAllKeysRequest) (*storagepb.DeleteAllKeysResponse, error)
}

type StringKeyService interface {
	Logger() kvdb.Logger
	GetDBNameFromContext(ctx context.Context) string
	SetString(ctx context.Context, req *storagepb.SetStringRequest) (*storagepb.SetStringResponse, error)
	GetString(ctx context.Context, req *storagepb.GetStringRequest) (*storagepb.GetStringResponse, error)
}

type HashMapKeyService interface {
	Logger() kvdb.Logger
	GetDBNameFromContext(ctx context.Context) string
	SetHashMap(ctx context.Context, req *storagepb.SetHashMapRequest) (*storagepb.SetHashMapResponse, error)
	GetHashMapFieldValues(ctx context.Context, req *storagepb.GetHashMapFieldValueRequest) (*storagepb.GetHashMapFieldValueResponse, error)
	GetAllHashMapFieldsAndValues(ctx context.Context, req *storagepb.GetAllHashMapFieldsAndValuesRequest) (*storagepb.GetAllHashMapFieldsAndValuesResponse, error)
	DeleteHashMapFields(ctx context.Context, req *storagepb.DeleteHashMapFieldsRequest) (*storagepb.DeleteHashMapFieldsResponse, error)
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

func (s *KvdbServer) CreateDatabase(ctx context.Context, req *dbpb.CreateDatabaseRequest) (*dbpb.CreateDatabaseResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.dbExists(req.DbName) {
		return nil, kvdberrors.ErrDatabaseExists
	}

	if err := kvdb.ValidateDatabaseName(req.DbName); err != nil {
		return nil, err
	}

	db := kvdb.CreateDatabase(req.DbName)
	s.databases[db.Name] = db

	return &dbpb.CreateDatabaseResponse{DbName: db.Name}, nil
}

func (s *KvdbServer) DeleteDatabase(ctx context.Context, req *dbpb.DeleteDatabaseRequest) (*dbpb.DeleteDatabaseResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.dbExists(req.DbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	delete(s.databases, req.DbName)

	return &dbpb.DeleteDatabaseResponse{DbName: req.DbName}, nil
}

func (s *KvdbServer) GetAllDatabases(ctx context.Context, req *dbpb.GetAllDatabasesRequest) (*dbpb.GetAllDatabasesResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var names []string
	for key := range s.databases {
		names = append(names, key)
	}

	return &dbpb.GetAllDatabasesResponse{DbNames: names}, nil
}

func (s *KvdbServer) GetDatabaseInfo(ctx context.Context, req *dbpb.GetDatabaseInfoRequest) (*dbpb.GetDatabaseInfoResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if !s.dbExists(req.DbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	db := s.databases[req.DbName]
	data := &dbpb.DatabaseInfo{
		Name:      db.Name,
		CreatedAt: timestamppb.New(db.CreatedAt),
		UpdatedAt: timestamppb.New(db.UpdatedAt),
		KeyCount:  db.GetKeyCount(),
		DataSize:  db.GetStoredSizeBytes(),
	}

	return &dbpb.GetDatabaseInfoResponse{Data: data}, nil
}

func (s *KvdbServer) GetAllKeys(ctx context.Context, req *storagepb.GetAllKeysRequest) (*storagepb.GetAllKeysResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	return &storagepb.GetAllKeysResponse{Keys: s.databases[dbName].GetKeys()}, nil
}

func (s *KvdbServer) GetKeyType(ctx context.Context, req *storagepb.GetKeyTypeRequest) (*storagepb.GetKeyTypeResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	keyType, ok := s.databases[dbName].GetKeyType(req.Key)

	return &storagepb.GetKeyTypeResponse{KeyType: keyType.String(), Ok: ok}, nil
}

func (s *KvdbServer) DeleteKeys(ctx context.Context, req *storagepb.DeleteKeysRequest) (*storagepb.DeleteKeysResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	keysDeletedCount := s.databases[dbName].DeleteKeys(req.Keys)

	return &storagepb.DeleteKeysResponse{KeysDeletedCount: keysDeletedCount}, nil
}

func (s *KvdbServer) DeleteAllKeys(ctx context.Context, req *storagepb.DeleteAllKeysRequest) (*storagepb.DeleteAllKeysResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	s.databases[dbName].DeleteAllKeys()

	return &storagepb.DeleteAllKeysResponse{}, nil
}

func (s *KvdbServer) SetString(ctx context.Context, req *storagepb.SetStringRequest) (*storagepb.SetStringResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	if err := kvdb.ValidateDatabaseKey(req.Key); err != nil {
		return nil, err
	}

	if s.DBMaxKeysReached(s.databases[dbName]) {
		return nil, kvdberrors.ErrMaxKeysReached
	}

	s.databases[dbName].SetString(req.Key, req.Value)

	return &storagepb.SetStringResponse{}, nil
}

func (s *KvdbServer) GetString(ctx context.Context, req *storagepb.GetStringRequest) (*storagepb.GetStringResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	value, ok := s.databases[dbName].GetString(req.Key)

	return &storagepb.GetStringResponse{Value: value, Ok: ok}, nil
}

func (s *KvdbServer) SetHashMap(ctx context.Context, req *storagepb.SetHashMapRequest) (*storagepb.SetHashMapResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	if err := kvdb.ValidateDatabaseKey(req.Key); err != nil {
		return nil, err
	}

	if s.DBMaxKeysReached(s.databases[dbName]) {
		return nil, kvdberrors.ErrMaxKeysReached
	}

	fieldsAddedCount := s.databases[dbName].SetHashMap(req.Key, req.FieldValueMap, s.Cfg.MaxHashMapFields)

	return &storagepb.SetHashMapResponse{FieldsAddedCount: fieldsAddedCount}, nil
}

func (s *KvdbServer) GetHashMapFieldValues(ctx context.Context, req *storagepb.GetHashMapFieldValueRequest) (res *storagepb.GetHashMapFieldValueResponse, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	result, ok := s.databases[dbName].GetHashMapFieldValues(req.Key, req.Fields)

	var fieldValueMap = make(map[string]*storagepb.HashMapFieldValue)
	for field, value := range result {
		fieldValueMap[field] = &storagepb.HashMapFieldValue{
			Value: value.Value,
			Ok:    value.Ok,
		}
	}

	return &storagepb.GetHashMapFieldValueResponse{FieldValueMap: fieldValueMap, Ok: ok}, nil
}

func (s *KvdbServer) GetAllHashMapFieldsAndValues(ctx context.Context, req *storagepb.GetAllHashMapFieldsAndValuesRequest) (res *storagepb.GetAllHashMapFieldsAndValuesResponse, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	fieldValueMap, ok := s.databases[dbName].GetAllHashMapFieldsAndValues(req.Key)

	return &storagepb.GetAllHashMapFieldsAndValuesResponse{FieldValueMap: fieldValueMap, Ok: ok}, nil
}

func (s *KvdbServer) DeleteHashMapFields(ctx context.Context, req *storagepb.DeleteHashMapFieldsRequest) (res *storagepb.DeleteHashMapFieldsResponse, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		return nil, kvdberrors.ErrDatabaseNotFound
	}

	fieldsRemovedCount, ok := s.databases[dbName].DeleteHashMapFields(req.Key, req.Fields)

	return &storagepb.DeleteHashMapFieldsResponse{FieldsRemovedCount: fieldsRemovedCount, Ok: ok}, nil
}
