package server

import (
	"context"

	kvdb "github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/api/v0/storagepb"
	kvdberrors "github.com/hollowdll/kvdb/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	getKeyTypeRPCName                   string = "GetKeyType"
	getAllKeysRPCName                   string = "GetAllKeys"
	deleteKeysRPCName                   string = "DeleteKeys"
	deleteAllKeysRPCName                string = "DeleteAllKeys"
	setStringRPCName                    string = "SetString"
	getStringRPCName                    string = "GetString"
	setHashMapRPCName                   string = "SetHashMap"
	getHashMapFieldValuesRPCName        string = "GetHashMapFieldValues"
	getAllHashMapFieldsAndValuesRPCName string = "GetAllHashMapFieldsAndValues"
	deleteHashMapFieldsRPCName          string = "DeleteHashMapFields"
)

// GetKeyType is the implementation of RPC GetKeyType.
func (s *Server) GetKeyType(ctx context.Context, req *storagepb.GetKeyTypeRequest) (res *storagepb.GetKeyTypeResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	dbName := s.getDatabaseNameFromContext(ctx)
	s.logger.Debugf("%s: (call) db = %s %v", getKeyTypeRPCName, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", getKeyTypeRPCName, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", getKeyTypeRPCName, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	keyType, ok := s.databases[dbName].GetKeyType(req.Key)

	return &storagepb.GetKeyTypeResponse{KeyType: keyType, Ok: ok}, nil
}

// SetString is the implementation of RPC SetString.
func (s *Server) SetString(ctx context.Context, req *storagepb.SetStringRequest) (res *storagepb.SetStringResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	dbName := s.getDatabaseNameFromContext(ctx)
	s.logger.Debugf("%s: (call) db = %s %v", setStringRPCName, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", setStringRPCName, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", setStringRPCName, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	if err := kvdb.ValidateDatabaseKey(req.Key); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if s.DbMaxKeysReached(s.databases[dbName]) {
		return nil, status.Error(codes.FailedPrecondition, kvdberrors.ErrMaxKeysReached.Error())
	}

	s.databases[dbName].SetString(req.Key, req.GetValue())

	return &storagepb.SetStringResponse{}, nil
}

// GetString is the implementation of RPC GetString.
func (s *Server) GetString(ctx context.Context, req *storagepb.GetStringRequest) (res *storagepb.GetStringResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	dbName := s.getDatabaseNameFromContext(ctx)
	s.logger.Debugf("%s: (call) db = %s %v", getStringRPCName, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", getStringRPCName, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", getStringRPCName, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	value, ok := s.databases[dbName].GetString(req.Key)

	return &storagepb.GetStringResponse{Value: value, Ok: ok}, nil
}

// DeleteKeys is the implementation of RPC DeleteKeys.
func (s *Server) DeleteKeys(ctx context.Context, req *storagepb.DeleteKeysRequest) (res *storagepb.DeleteKeysResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	dbName := s.getDatabaseNameFromContext(ctx)
	s.logger.Debugf("%s: (call) db = %s %v", deleteKeysRPCName, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", deleteKeysRPCName, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", deleteKeysRPCName, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	keysDeletedCount := s.databases[dbName].DeleteKeys(req.Keys)

	return &storagepb.DeleteKeysResponse{KeysDeletedCount: keysDeletedCount}, nil
}

// DeleteAllKeys is the implementation of RPC DeleteAllKeys.
func (s *Server) DeleteAllKeys(ctx context.Context, req *storagepb.DeleteAllKeysRequest) (res *storagepb.DeleteAllKeysResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	dbName := s.getDatabaseNameFromContext(ctx)
	s.logger.Debugf("%s: (call) db = %s %v", deleteAllKeysRPCName, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", deleteAllKeysRPCName, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", deleteAllKeysRPCName, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	s.databases[dbName].DeleteAllKeys()

	return &storagepb.DeleteAllKeysResponse{}, nil
}

// GetAllKeys is the implementation of RPC GetAllKeys.
func (s *Server) GetAllKeys(ctx context.Context, req *storagepb.GetAllKeysRequest) (res *storagepb.GetAllKeysResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	dbName := s.getDatabaseNameFromContext(ctx)
	s.logger.Debugf("%s: (call) db = %s %v", getAllKeysRPCName, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", getAllKeysRPCName, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", getAllKeysRPCName, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	return &storagepb.GetAllKeysResponse{Keys: s.databases[dbName].GetKeys()}, nil
}

// SetHashMap is the implementation of RPC SetHashMap.
func (s *Server) SetHashMap(ctx context.Context, req *storagepb.SetHashMapRequest) (res *storagepb.SetHashMapResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	dbName := s.getDatabaseNameFromContext(ctx)
	s.logger.Debugf("%s: (call) db = %s %v", setHashMapRPCName, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", setHashMapRPCName, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", setHashMapRPCName, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	if err := kvdb.ValidateDatabaseKey(req.Key); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if s.DbMaxKeysReached(s.databases[dbName]) {
		return nil, status.Error(codes.FailedPrecondition, kvdberrors.ErrMaxKeysReached.Error())
	}

	fieldsAddedCount := s.databases[dbName].SetHashMap(req.Key, req.FieldValueMap, s.maxHashMapFields)

	return &storagepb.SetHashMapResponse{FieldsAddedCount: fieldsAddedCount}, nil
}

// GetHashMapFieldValues is the implementation of RPC GetHashMapFieldValues.
func (s *Server) GetHashMapFieldValues(ctx context.Context, req *storagepb.GetHashMapFieldValueRequest) (res *storagepb.GetHashMapFieldValueResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	dbName := s.getDatabaseNameFromContext(ctx)
	s.logger.Debugf("%s: (call) db = %s %v", getHashMapFieldValuesRPCName, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", getHashMapFieldValuesRPCName, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", getHashMapFieldValuesRPCName, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
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

// DeleteHashMapFields is the implementation of RPC DeleteHashMapFields.
func (s *Server) DeleteHashMapFields(ctx context.Context, req *storagepb.DeleteHashMapFieldsRequest) (res *storagepb.DeleteHashMapFieldsResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	dbName := s.getDatabaseNameFromContext(ctx)
	s.logger.Debugf("%s: (call) db = %s %v", deleteHashMapFieldsRPCName, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", deleteHashMapFieldsRPCName, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", deleteHashMapFieldsRPCName, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	fieldsRemovedCount, ok := s.databases[dbName].DeleteHashMapFields(req.Key, req.Fields)

	return &storagepb.DeleteHashMapFieldsResponse{FieldsRemovedCount: fieldsRemovedCount, Ok: ok}, nil
}

// GetAllHashMapFieldsAndValues is the implementation of RPC GetAllHashMapFieldsAndValues.
func (s *Server) GetAllHashMapFieldsAndValues(ctx context.Context, req *storagepb.GetAllHashMapFieldsAndValuesRequest) (res *storagepb.GetAllHashMapFieldsAndValuesResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	dbName := s.getDatabaseNameFromContext(ctx)
	s.logger.Debugf("%s: (call) db = %s %v", getAllHashMapFieldsAndValuesRPCName, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", getAllHashMapFieldsAndValuesRPCName, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", getAllHashMapFieldsAndValuesRPCName, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	fieldValueMap, ok := s.databases[dbName].GetAllHashMapFieldsAndValues(req.Key)

	return &storagepb.GetAllHashMapFieldsAndValuesResponse{FieldValueMap: fieldValueMap, Ok: ok}, nil
}
