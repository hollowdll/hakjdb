package server

import (
	"context"

	kvdb "github.com/hollowdll/kvdb"
	kvdberrors "github.com/hollowdll/kvdb/errors"
	"github.com/hollowdll/kvdb/proto/kvdbserverpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetTypeOfKey is the implementation of RPC GetTypeOfKey.
func (s *Server) GetTypeOfKey(ctx context.Context, req *kvdbserverpb.GetTypeOfKeyRequest) (res *kvdbserverpb.GetTypeOfKeyResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	logPrefix := "GetTypeOfKey"
	dbName := s.getDatabaseNameFromContext(ctx)

	s.logger.Debugf("%s: (attempt) db = %s %v", logPrefix, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", logPrefix, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	keyType, ok := s.databases[dbName].GetTypeOfKey(kvdb.DatabaseKey(req.Key))

	return &kvdbserverpb.GetTypeOfKeyResponse{KeyType: keyType, Ok: ok}, nil
}

// SetString is the implementation of RPC SetString.
func (s *Server) SetString(ctx context.Context, req *kvdbserverpb.SetStringRequest) (res *kvdbserverpb.SetStringResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	logPrefix := "SetString"
	dbName := s.getDatabaseNameFromContext(ctx)

	s.logger.Debugf("%s: (attempt) db = %s %v", logPrefix, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", logPrefix, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	if err := kvdb.ValidateDatabaseKey(kvdb.DatabaseKey(req.GetKey())); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if s.DbMaxKeysReached(s.databases[dbName]) {
		return nil, status.Error(codes.FailedPrecondition, kvdberrors.ErrMaxKeysReached.Error())
	}

	s.databases[dbName].SetString(kvdb.DatabaseKey(req.GetKey()), kvdb.DatabaseStringValue(req.GetValue()))

	return &kvdbserverpb.SetStringResponse{}, nil
}

// GetString is the implementation of RPC GetString.
func (s *Server) GetString(ctx context.Context, req *kvdbserverpb.GetStringRequest) (res *kvdbserverpb.GetStringResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	logPrefix := "GetString"
	dbName := s.getDatabaseNameFromContext(ctx)

	s.logger.Debugf("%s: (attempt) db = %s %v", logPrefix, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", logPrefix, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	value, ok := s.databases[dbName].GetString(kvdb.DatabaseKey(req.GetKey()))

	return &kvdbserverpb.GetStringResponse{Value: string(value), Ok: ok}, nil
}

// DeleteKey is the implementation of RPC DeleteKey.
func (s *Server) DeleteKey(ctx context.Context, req *kvdbserverpb.DeleteKeyRequest) (res *kvdbserverpb.DeleteKeyResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	logPrefix := "DeleteKey"
	dbName := s.getDatabaseNameFromContext(ctx)

	s.logger.Debugf("%s: (attempt) db = %s %v", logPrefix, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", logPrefix, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	ok := s.databases[dbName].DeleteKey(kvdb.DatabaseKey(req.GetKey()))
	if !ok {
		return &kvdbserverpb.DeleteKeyResponse{Ok: false}, nil
	}

	return &kvdbserverpb.DeleteKeyResponse{Ok: true}, nil
}

// DeleteAllKeys is the implementation of RPC DeleteAllKeys.
func (s *Server) DeleteAllKeys(ctx context.Context, req *kvdbserverpb.DeleteAllKeysRequest) (res *kvdbserverpb.DeleteAllKeysResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	logPrefix := "DeleteAllKeys"
	dbName := s.getDatabaseNameFromContext(ctx)

	s.logger.Debugf("%s: (attempt) db = %s %v", logPrefix, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", logPrefix, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	s.databases[dbName].DeleteAllKeys()

	return &kvdbserverpb.DeleteAllKeysResponse{}, nil
}

// GetKeys is the implementation of RPC GetKeys.
func (s *Server) GetKeys(ctx context.Context, req *kvdbserverpb.GetKeysRequest) (res *kvdbserverpb.GetKeysResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	logPrefix := "GetKeys"
	dbName := s.getDatabaseNameFromContext(ctx)

	s.logger.Debugf("%s: (attempt) db = %s %v", logPrefix, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", logPrefix, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	return &kvdbserverpb.GetKeysResponse{Keys: s.databases[dbName].GetKeys()}, nil
}

// SetHashMap is the implementation of RPC SetHashMap.
func (s *Server) SetHashMap(ctx context.Context, req *kvdbserverpb.SetHashMapRequest) (res *kvdbserverpb.SetHashMapResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	logPrefix := "SetHashMap"
	dbName := s.getDatabaseNameFromContext(ctx)

	s.logger.Debugf("%s: (attempt) db = %s %v", logPrefix, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", logPrefix, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	if err := kvdb.ValidateDatabaseKey(kvdb.DatabaseKey(req.GetKey())); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if s.DbMaxKeysReached(s.databases[dbName]) {
		return nil, status.Error(codes.FailedPrecondition, kvdberrors.ErrMaxKeysReached.Error())
	}

	fieldsAdded := s.databases[dbName].SetHashMap(kvdb.DatabaseKey(req.Key), req.Fields, s.maxHashMapFields)

	return &kvdbserverpb.SetHashMapResponse{FieldsAdded: fieldsAdded}, nil
}

// GetHashMapFieldValue is the implementation of RPC GetHashMapFieldValue.
func (s *Server) GetHashMapFieldValue(ctx context.Context, req *kvdbserverpb.GetHashMapFieldValueRequest) (res *kvdbserverpb.GetHashMapFieldValueResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	logPrefix := "GetHashMapFieldValue"
	dbName := s.getDatabaseNameFromContext(ctx)

	s.logger.Debugf("%s: (attempt) db = %s %v", logPrefix, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", logPrefix, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	value, ok := s.databases[dbName].GetHashMapFieldValue(kvdb.DatabaseKey(req.Key), req.Field)

	return &kvdbserverpb.GetHashMapFieldValueResponse{Value: value, Ok: ok}, nil
}

// DeleteHashMapFields is the implementation of RPC DeleteHashMapFields.
func (s *Server) DeleteHashMapFields(ctx context.Context, req *kvdbserverpb.DeleteHashMapFieldsRequest) (res *kvdbserverpb.DeleteHashMapFieldsResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	logPrefix := "DeleteHashMapFields"
	dbName := s.getDatabaseNameFromContext(ctx)

	s.logger.Debugf("%s: (attempt) db = %s %v", logPrefix, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", logPrefix, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	fieldsRemoved, ok := s.databases[dbName].DeleteHashMapFields(kvdb.DatabaseKey(req.Key), req.Fields)

	return &kvdbserverpb.DeleteHashMapFieldsResponse{FieldsRemoved: fieldsRemoved, Ok: ok}, nil
}

// GetAllHashMapFieldsAndValues is the implementation of RPC GetAllHashMapFieldsAndValues.
func (s *Server) GetAllHashMapFieldsAndValues(ctx context.Context, req *kvdbserverpb.GetAllHashMapFieldsAndValuesRequest) (res *kvdbserverpb.GetAllHashMapFieldsAndValuesResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	logPrefix := "GetAllHashMapFieldsAndValues"
	dbName := s.getDatabaseNameFromContext(ctx)

	s.logger.Debugf("%s: (attempt) db = %s %v", logPrefix, dbName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: (success) db = %s %v", logPrefix, dbName, req)
		}
	}()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	fieldValueMap, ok := s.databases[dbName].GetAllHashMapFieldsAndValues(kvdb.DatabaseKey(req.Key))

	return &kvdbserverpb.GetAllHashMapFieldsAndValuesResponse{FieldValueMap: fieldValueMap, Ok: ok}, nil
}
