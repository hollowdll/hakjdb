package server

import (
	"context"

	kvdb "github.com/hollowdll/kvdb"
	kvdberrors "github.com/hollowdll/kvdb/errors"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// GetTypeOfKey is the implementation of RPC GetTypeOfKey.
func (s *Server) GetTypeOfKey(ctx context.Context, req *kvdbserver.GetTypeOfKeyRequest) (res *kvdbserver.GetTypeOfKeyResponse, err error) {
	logPrefix := "GetTypeOfKey"
	s.logger.Debugf("%s: attempt to get key data type", logPrefix)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: failed to get the data type of a key: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: get key data type success", logPrefix)
		}
	}()

	dbName, err := getDatabaseNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	keyType, ok := s.databases[dbName].GetTypeOfKey(kvdb.DatabaseKey(req.Key))

	return &kvdbserver.GetTypeOfKeyResponse{KeyType: keyType, Ok: ok}, nil
}

// SetString is the implementation of RPC SetString.
func (s *Server) SetString(ctx context.Context, req *kvdbserver.SetStringRequest) (res *kvdbserver.SetStringResponse, err error) {
	logPrefix := "SetString"
	s.logger.Debugf("%s: attempt to set string value", logPrefix)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: failed to set string value: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: set string value success", logPrefix)
		}
	}()

	dbName, err := getDatabaseNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

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

	return &kvdbserver.SetStringResponse{}, nil
}

// GetString is the implementation of RPC GetString.
func (s *Server) GetString(ctx context.Context, req *kvdbserver.GetStringRequest) (res *kvdbserver.GetStringResponse, err error) {
	logPrefix := "GetString"
	s.logger.Debugf("%s: attempt to get string value", logPrefix)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: failed to get string value: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: get string value success", logPrefix)
		}
	}()

	dbName, err := getDatabaseNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	value, ok := s.databases[dbName].GetString(kvdb.DatabaseKey(req.GetKey()))

	return &kvdbserver.GetStringResponse{Value: string(value), Ok: ok}, nil
}

// DeleteKey is the implementation of RPC DeleteKey.
func (s *Server) DeleteKey(ctx context.Context, req *kvdbserver.DeleteKeyRequest) (res *kvdbserver.DeleteKeyResponse, err error) {
	logPrefix := "DeleteKey"
	s.logger.Debugf("%s: attempt to delete key", logPrefix)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: failed to delete key: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: delete key success", logPrefix)
		}
	}()

	dbName, err := getDatabaseNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	ok := s.databases[dbName].DeleteKey(kvdb.DatabaseKey(req.GetKey()))
	if !ok {
		return &kvdbserver.DeleteKeyResponse{Ok: false}, nil
	}

	return &kvdbserver.DeleteKeyResponse{Ok: true}, nil
}

// DeleteAllKeys is the implementation of RPC DeleteAllKeys.
func (s *Server) DeleteAllKeys(ctx context.Context, req *kvdbserver.DeleteAllKeysRequest) (res *kvdbserver.DeleteAllKeysResponse, err error) {
	logPrefix := "DeleteAllKeys"
	s.logger.Debugf("%s: attempt to delete all keys", logPrefix)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: failed to delete all keys: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: delete all keys success", logPrefix)
		}
	}()

	dbName, err := getDatabaseNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	s.databases[dbName].DeleteAllKeys()

	return &kvdbserver.DeleteAllKeysResponse{}, nil
}

// GetKeys is the implementation of RPC GetKeys.
func (s *Server) GetKeys(ctx context.Context, req *kvdbserver.GetKeysRequest) (res *kvdbserver.GetKeysResponse, err error) {
	logPrefix := "GetKeys"
	s.logger.Debugf("%s: attempt to get keys", logPrefix)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: failed to get keys: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: get keys success", logPrefix)
		}
	}()

	dbName, err := getDatabaseNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	return &kvdbserver.GetKeysResponse{Keys: s.databases[dbName].GetKeys()}, nil
}

// SetHashMap is the implementation of RPC SetHashMap.
func (s *Server) SetHashMap(ctx context.Context, req *kvdbserver.SetHashMapRequest) (res *kvdbserver.SetHashMapResponse, err error) {
	logPrefix := "SetHashMap"
	s.logger.Debugf("%s: attempt to set HashMap fields", logPrefix)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: failed to set HashMap fields: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: set HashMap fields success", logPrefix)
		}
	}()

	dbName, err := getDatabaseNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	if err := kvdb.ValidateDatabaseKey(kvdb.DatabaseKey(req.GetKey())); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if s.DbMaxKeysReached(s.databases[dbName]) {
		return nil, status.Error(codes.FailedPrecondition, kvdberrors.ErrMaxKeysReached.Error())
	}

	s.databases[dbName].SetHashMap(kvdb.DatabaseKey(req.Key), req.Fields)

	return &kvdbserver.SetHashMapResponse{}, nil
}

// GetHashMapFieldValue is the implementation of RPC GetHashMapFieldValue.
func (s *Server) GetHashMapFieldValue(ctx context.Context, req *kvdbserver.GetHashMapFieldValueRequest) (res *kvdbserver.GetHashMapFieldValueResponse, err error) {
	logPrefix := "GetHashMapFieldValue"
	s.logger.Debugf("%s: attempt to get HashMap field value", logPrefix)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: failed to get HashMap field value: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: get HashMap field value success", logPrefix)
		}
	}()

	dbName, err := getDatabaseNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	value, ok := s.databases[dbName].GetHashMapFieldValue(kvdb.DatabaseKey(req.Key), req.Field)

	return &kvdbserver.GetHashMapFieldValueResponse{Value: value, Ok: ok}, nil
}

// getDatabaseNameFromContext gets the database name from the received gRPC metadata.
func getDatabaseNameFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.InvalidArgument, kvdberrors.ErrMissingMetadata.Error())
	}

	dbName := md.Get(common.GrpcMetadataKeyDbName)
	if len(dbName) < 1 {
		return "", status.Errorf(codes.InvalidArgument, "%s (%s)", kvdberrors.ErrMissingKeyInMetadata, common.GrpcMetadataKeyDbName)
	}

	return dbName[0], nil
}
