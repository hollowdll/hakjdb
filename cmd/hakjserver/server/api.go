package server

import (
	"context"
	"os"
	"runtime"
	"time"

	"github.com/hollowdll/hakjdb"
	"github.com/hollowdll/hakjdb/api/v1/authpb"
	"github.com/hollowdll/hakjdb/api/v1/dbpb"
	"github.com/hollowdll/hakjdb/api/v1/kvpb"
	"github.com/hollowdll/hakjdb/api/v1/serverpb"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/auth"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/validation"
	hakjerrors "github.com/hollowdll/hakjdb/errors"
	"github.com/hollowdll/hakjdb/internal/common"
	"github.com/hollowdll/hakjdb/version"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServerService interface {
	GetServerInfo(ctx context.Context, req *serverpb.GetServerInfoRequest) (*serverpb.GetServerInfoResponse, error)
	GetLogs(ctx context.Context, req *serverpb.GetLogsRequest) (*serverpb.GetLogsResponse, error)
	ReloadConfig(ctx context.Context, req *serverpb.ReloadConfigRequest) (*serverpb.ReloadConfigResponse, error)
}

type DBService interface {
	CreateDB(ctx context.Context, req *dbpb.CreateDBRequest) (*dbpb.CreateDBResponse, error)
	DeleteDB(ctx context.Context, req *dbpb.DeleteDBRequest) (*dbpb.DeleteDBResponse, error)
	GetAllDBs(ctx context.Context, req *dbpb.GetAllDBsRequest) (*dbpb.GetAllDBsResponse, error)
	GetDBInfo(ctx context.Context, req *dbpb.GetDBInfoRequest) (*dbpb.GetDBInfoResponse, error)
	ChangeDB(ctx context.Context, req *dbpb.ChangeDBRequest) (*dbpb.ChangeDBResponse, error)
}

type GeneralKVService interface {
	GetAllKeys(ctx context.Context, req *kvpb.GetAllKeysRequest) (*kvpb.GetAllKeysResponse, error)
	GetKeyType(ctx context.Context, req *kvpb.GetKeyTypeRequest) (*kvpb.GetKeyTypeResponse, error)
	DeleteKeys(ctx context.Context, req *kvpb.DeleteKeysRequest) (*kvpb.DeleteKeysResponse, error)
	DeleteAllKeys(ctx context.Context, req *kvpb.DeleteAllKeysRequest) (*kvpb.DeleteAllKeysResponse, error)
}

type StringKVService interface {
	SetString(ctx context.Context, req *kvpb.SetStringRequest) (*kvpb.SetStringResponse, error)
	GetString(ctx context.Context, req *kvpb.GetStringRequest) (*kvpb.GetStringResponse, error)
}

type HashMapKVService interface {
	SetHashMap(ctx context.Context, req *kvpb.SetHashMapRequest) (*kvpb.SetHashMapResponse, error)
	GetHashMapFieldValues(ctx context.Context, req *kvpb.GetHashMapFieldValuesRequest) (*kvpb.GetHashMapFieldValuesResponse, error)
	GetAllHashMapFieldsAndValues(ctx context.Context, req *kvpb.GetAllHashMapFieldsAndValuesRequest) (*kvpb.GetAllHashMapFieldsAndValuesResponse, error)
	DeleteHashMapFields(ctx context.Context, req *kvpb.DeleteHashMapFieldsRequest) (*kvpb.DeleteHashMapFieldsResponse, error)
}

type AuthService interface {
	Authenticate(ctx context.Context, req *authpb.AuthenticateRequest) (*authpb.AuthenticateResponse, error)
}

func (s *HakjServer) GetServerInfo(ctx context.Context, req *serverpb.GetServerInfoRequest) (*serverpb.GetServerInfoResponse, error) {
	lg := s.Logger()
	cfg := s.Config()
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.ClientConnListener != nil {
		s.ClientConnListener.mu.RLock()
		defer s.ClientConnListener.mu.RUnlock()
	}

	osInfo, err := getOsInfo()
	if err != nil {
		lg.Errorf("%v: %v", hakjerrors.ErrGetOSInfo, err)
		return nil, hakjerrors.ErrGetOSInfo
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	var totalKeys uint64
	for _, db := range s.dbs {
		totalKeys += uint64(db.GetKeyCount())
	}

	generalInfo := &serverpb.GeneralInfo{
		ServerVersion:            version.Version,
		GoVersion:                runtime.Version(),
		Os:                       osInfo,
		Arch:                     runtime.GOARCH,
		ProcessId:                uint32(os.Getpid()),
		UptimeSeconds:            uint64(time.Since(s.startTime).Seconds()),
		TcpPort:                  uint32(cfg.PortInUse),
		TlsEnabled:               cfg.TLSEnabled,
		TlsClientCertAuthEnabled: cfg.TLSClientCertAuthEnabled,
		AuthEnabled:              cfg.AuthEnabled,
		LogfileEnabled:           cfg.LogFileEnabled,
		DebugEnabled:             cfg.DebugEnabled,
		ApiVersion:               version.APIVersion,
	}
	memoryInfo := &serverpb.MemoryInfo{
		MemoryAlloc:      m.Alloc,
		MemoryTotalAlloc: m.TotalAlloc,
		MemorySys:        m.Sys,
	}
	storageInfo := &serverpb.StorageInfo{
		TotalDataSize: s.totalStoredDataSize(),
		TotalKeys:     totalKeys,
	}
	clientInfo := &serverpb.ClientInfo{
		ClientConnections:    s.ClientConnListener.clientConnectionsCount,
		MaxClientConnections: s.ClientConnListener.maxClientConnections,
	}
	dbInfo := &serverpb.DatabaseInfo{
		DbCount:   uint32(len(s.dbs)),
		DefaultDb: cfg.DefaultDB,
	}

	return &serverpb.GetServerInfoResponse{
		GeneralInfo: generalInfo,
		MemoryInfo:  memoryInfo,
		StorageInfo: storageInfo,
		ClientInfo:  clientInfo,
		DbInfo:      dbInfo,
	}, nil
}

func (s *HakjServer) GetLogs(ctx context.Context, req *serverpb.GetLogsRequest) (*serverpb.GetLogsResponse, error) {
	lg := s.Logger()
	cfg := s.Config()
	s.mu.RLock()
	defer s.mu.RUnlock()

	if !cfg.LogFileEnabled {
		lg.Debug("Logs were requested but the log file is not enabled. Consider enabling it.")
		return nil, hakjerrors.ErrLogFileNotEnabled
	}

	logs, err := common.ReadFileLines(cfg.LogFilePath)
	if err != nil {
		lg.Errorf("%v: %v", hakjerrors.ErrReadLogFile, err)
		return nil, hakjerrors.ErrReadLogFile
	}

	return &serverpb.GetLogsResponse{Logs: logs}, nil
}

func (s *HakjServer) ReloadConfig(ctx context.Context, req *serverpb.ReloadConfigRequest) (*serverpb.ReloadConfigResponse, error) {
	lg := s.Logger()
	s.cfgMu.Lock()
	defer s.cfgMu.Unlock()

	s.Cfg.Reload(lg)
	s.ProcessConfigReload(&s.Cfg)

	return &serverpb.ReloadConfigResponse{}, nil
}

func (s *HakjServer) CreateDB(ctx context.Context, req *dbpb.CreateDBRequest) (*dbpb.CreateDBResponse, error) {
	if err := validateCreateDB(req); err != nil {
		return nil, err
	}

	lg := s.Logger()
	cfg := s.Config()
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.dbExists(req.DbName) {
		logDBAlreadyExists(lg, req.DbName)
		return nil, hakjerrors.ErrDatabaseExists
	}

	dbConfig := hakjdb.DBConfig{MaxHashMapFields: cfg.MaxHashMapFields}
	db := hakjdb.NewDB(req.DbName, req.Description, dbConfig)
	s.dbs[db.Name()] = db
	lg.Infof("Created database '%s'", db.Name())

	return &dbpb.CreateDBResponse{DbName: db.Name()}, nil
}

func (s *HakjServer) DeleteDB(ctx context.Context, req *dbpb.DeleteDBRequest) (*dbpb.DeleteDBResponse, error) {
	lg := s.Logger()
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.dbExists(req.DbName) {
		logDBNotFound(lg, req.DbName)
		return nil, hakjerrors.ErrDatabaseNotFound
	}

	delete(s.dbs, req.DbName)
	lg.Infof("Deleted database '%s'", req.DbName)

	return &dbpb.DeleteDBResponse{DbName: req.DbName}, nil
}

func (s *HakjServer) GetAllDBs(ctx context.Context, req *dbpb.GetAllDBsRequest) (*dbpb.GetAllDBsResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var dbNames []string
	for dbName := range s.dbs {
		dbNames = append(dbNames, dbName)
	}

	return &dbpb.GetAllDBsResponse{DbNames: dbNames}, nil
}

func (s *HakjServer) GetDBInfo(ctx context.Context, req *dbpb.GetDBInfoRequest) (*dbpb.GetDBInfoResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if !s.dbExists(req.DbName) {
		lg := s.Logger()
		logDBNotFound(lg, req.DbName)
		return nil, hakjerrors.ErrDatabaseNotFound
	}

	db := s.dbs[req.DbName]
	data := &dbpb.DBInfo{
		Name:        db.Name(),
		Description: db.Description(),
		CreatedAt:   timestamppb.New(db.CreatedAt()),
		UpdatedAt:   timestamppb.New(db.UpdatedAt()),
		KeyCount:    uint32(db.GetKeyCount()),
		DataSize:    db.GetEstimatedStorageSizeBytes(),
	}

	return &dbpb.GetDBInfoResponse{Data: data}, nil
}

func (s *HakjServer) ChangeDB(ctx context.Context, req *dbpb.ChangeDBRequest) (*dbpb.ChangeDBResponse, error) {
	if err := validateChangeDB(req); err != nil {
		return nil, err
	}

	lg := s.Logger()
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.dbExists(req.DbName) {
		logDBNotFound(lg, req.DbName)
		return nil, hakjerrors.ErrDatabaseNotFound
	}

	db := s.dbs[req.DbName]
	dbName := req.DbName
	if req.ChangeName {
		db.ChangeName(req.NewName)
		dbName = req.NewName
		delete(s.dbs, req.DbName)
		lg.Infof("Changed name of database '%s', new name: '%s'", req.DbName, dbName)
	}
	if req.ChangeDescription {
		db.ChangeDescription(req.NewDescription)
		lg.Infof("Changed description of database '%s'", req.DbName)
	}
	s.dbs[dbName] = db

	return &dbpb.ChangeDBResponse{DbName: dbName}, nil
}

func (s *HakjServer) GetAllKeys(ctx context.Context, req *kvpb.GetAllKeysRequest) (*kvpb.GetAllKeysResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		lg := s.Logger()
		logDBNotFound(lg, dbName)
		return nil, hakjerrors.ErrDatabaseNotFound
	}

	return &kvpb.GetAllKeysResponse{Keys: s.dbs[dbName].GetAllKeys()}, nil
}

func (s *HakjServer) GetKeyType(ctx context.Context, req *kvpb.GetKeyTypeRequest) (*kvpb.GetKeyTypeResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		lg := s.Logger()
		logDBNotFound(lg, dbName)
		return nil, hakjerrors.ErrDatabaseNotFound
	}

	keyType, ok := s.dbs[dbName].GetKeyType(req.Key)

	return &kvpb.GetKeyTypeResponse{KeyType: keyType.String(), Ok: ok}, nil
}

func (s *HakjServer) DeleteKeys(ctx context.Context, req *kvpb.DeleteKeysRequest) (*kvpb.DeleteKeysResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		lg := s.Logger()
		logDBNotFound(lg, dbName)
		return nil, hakjerrors.ErrDatabaseNotFound
	}

	keysDeletedCount := s.dbs[dbName].DeleteKeys(req.Keys)

	return &kvpb.DeleteKeysResponse{KeysDeletedCount: keysDeletedCount}, nil
}

func (s *HakjServer) DeleteAllKeys(ctx context.Context, req *kvpb.DeleteAllKeysRequest) (*kvpb.DeleteAllKeysResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		lg := s.Logger()
		logDBNotFound(lg, dbName)
		return nil, hakjerrors.ErrDatabaseNotFound
	}

	s.dbs[dbName].DeleteAllKeys()

	return &kvpb.DeleteAllKeysResponse{}, nil
}

func (s *HakjServer) SetString(ctx context.Context, req *kvpb.SetStringRequest) (*kvpb.SetStringResponse, error) {
	if err := validation.ValidateDBKey(req.Key); err != nil {
		return nil, err
	}

	lg := s.Logger()
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		logDBNotFound(lg, dbName)
		return nil, hakjerrors.ErrDatabaseNotFound
	}

	if s.DBMaxKeysReached(s.dbs[dbName]) {
		logMaxKeysReached(lg, dbName)
		return nil, hakjerrors.ErrMaxKeysReached
	}

	s.dbs[dbName].SetString(req.Key, req.Value)

	return &kvpb.SetStringResponse{}, nil
}

func (s *HakjServer) GetString(ctx context.Context, req *kvpb.GetStringRequest) (*kvpb.GetStringResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		lg := s.Logger()
		logDBNotFound(lg, dbName)
		return nil, hakjerrors.ErrDatabaseNotFound
	}

	kv, ok := s.dbs[dbName].GetStringKey(req.Key)

	return &kvpb.GetStringResponse{Value: kv.Value, Ok: ok}, nil
}

func (s *HakjServer) SetHashMap(ctx context.Context, req *kvpb.SetHashMapRequest) (*kvpb.SetHashMapResponse, error) {
	if err := validation.ValidateDBKey(req.Key); err != nil {
		return nil, err
	}

	lg := s.Logger()
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		logDBNotFound(lg, dbName)
		return nil, hakjerrors.ErrDatabaseNotFound
	}

	if s.DBMaxKeysReached(s.dbs[dbName]) {
		logMaxKeysReached(lg, dbName)
		return nil, hakjerrors.ErrMaxKeysReached
	}

	fieldsAddedCount := s.dbs[dbName].SetHashMap(req.Key, req.FieldValueMap)

	return &kvpb.SetHashMapResponse{FieldsAddedCount: fieldsAddedCount}, nil
}

func (s *HakjServer) GetHashMapFieldValues(ctx context.Context, req *kvpb.GetHashMapFieldValuesRequest) (res *kvpb.GetHashMapFieldValuesResponse, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		lg := s.Logger()
		logDBNotFound(lg, dbName)
		return nil, hakjerrors.ErrDatabaseNotFound
	}

	result, ok := s.dbs[dbName].GetHashMapFieldValues(req.Key, req.Fields)

	var fieldValueMap = make(map[string]*kvpb.HashMapFieldValue)
	for field, value := range result {
		fieldValueMap[field] = &kvpb.HashMapFieldValue{
			Value: value.FieldValue.Value,
			Ok:    value.Ok,
		}
	}

	return &kvpb.GetHashMapFieldValuesResponse{FieldValueMap: fieldValueMap, Ok: ok}, nil
}

func (s *HakjServer) GetAllHashMapFieldsAndValues(ctx context.Context, req *kvpb.GetAllHashMapFieldsAndValuesRequest) (res *kvpb.GetAllHashMapFieldsAndValuesResponse, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		lg := s.Logger()
		logDBNotFound(lg, dbName)
		return nil, hakjerrors.ErrDatabaseNotFound
	}

	kv, ok := s.dbs[dbName].GetHashMapKey(req.Key)
	fieldValueMap := make(map[string][]byte)
	for field, value := range kv.Value {
		fieldValueMap[field] = value.Value
	}

	return &kvpb.GetAllHashMapFieldsAndValuesResponse{FieldValueMap: fieldValueMap, Ok: ok}, nil
}

func (s *HakjServer) DeleteHashMapFields(ctx context.Context, req *kvpb.DeleteHashMapFieldsRequest) (res *kvpb.DeleteHashMapFieldsResponse, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dbName := s.GetDBNameFromContext(ctx)
	if !s.dbExists(dbName) {
		lg := s.Logger()
		logDBNotFound(lg, dbName)
		return nil, hakjerrors.ErrDatabaseNotFound
	}

	fieldsRemovedCount, ok := s.dbs[dbName].DeleteHashMapFields(req.Key, req.Fields)

	return &kvpb.DeleteHashMapFieldsResponse{FieldsRemovedCount: fieldsRemovedCount, Ok: ok}, nil
}

func (s *HakjServer) Authenticate(ctx context.Context, req *authpb.AuthenticateRequest) (*authpb.AuthenticateResponse, error) {
	lg := s.Logger()
	cfg := s.Config()
	s.mu.RLock()
	defer s.mu.RUnlock()

	defer func() {
		if req != nil {
			req.Password = ""
		}
	}()

	if !cfg.AuthEnabled {
		return nil, hakjerrors.ErrAuthNotEnabled
	}

	username := auth.RootUserName
	err := s.credentialStore.IsCorrectPassword(username, []byte(req.Password))
	if err != nil {
		lg.Debugf("%v: %v", hakjerrors.ErrInvalidCredentials, err)
		return nil, hakjerrors.ErrInvalidCredentials
	}

	opts := &auth.JWTOptions{
		SignKey: cfg.AuthTokenSecretKey,
		TTL:     time.Duration(cfg.AuthTokenTTL) * time.Second,
	}
	lg.Debugf("JWT token TTL: %s", opts.TTL)
	token, err := auth.GenerateJWT(opts, username)
	if err != nil {
		lg.Debugf("failed to generate JWT token: %v", err)
		return nil, hakjerrors.ErrAuthFailed
	}
	lg.Debugf("created a new JWT token for user %s", username)

	return &authpb.AuthenticateResponse{AuthToken: token}, nil
}

func logDBNotFound(lg hakjdb.Logger, dbName string) {
	lg.Warningf("database '%s' not found", dbName)
}

func logDBAlreadyExists(lg hakjdb.Logger, dbName string) {
	lg.Warningf("database '%s' already exists", dbName)
}

func logMaxKeysReached(lg hakjdb.Logger, dbName string) {
	lg.Warningf("maximum number of keys reached in database '%s'", dbName)
}

func validateCreateDB(req *dbpb.CreateDBRequest) error {
	if err := validation.ValidateDBName(req.DbName); err != nil {
		return err
	}
	if err := validation.ValidateDBDesc(req.Description); err != nil {
		return err
	}
	return nil
}

func validateChangeDB(req *dbpb.ChangeDBRequest) error {
	if req.ChangeName {
		if err := validation.ValidateDBName(req.NewName); err != nil {
			return err
		}
	}
	if req.ChangeDescription {
		if err := validation.ValidateDBDesc(req.NewDescription); err != nil {
			return err
		}
	}
	return nil
}
