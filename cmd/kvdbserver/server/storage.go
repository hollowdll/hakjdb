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

// SetString sets a string value using a key.
// Accepts database name in gRPC metadata.
func (s *Server) SetString(ctx context.Context, req *kvdbserver.SetStringRequest) (res *kvdbserver.SetStringResponse, err error) {
	s.logger.Debug("Attempt to set string value")
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to set string value: %s", err)
		} else {
			s.logger.Debug("Set string value success")
		}
	}()

	dbName, err := getDatabaseNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	err = s.databases[dbName].SetString(kvdb.DatabaseKey(req.GetKey()), kvdb.DatabaseStringValue(req.GetValue()))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s", err)
	}

	return &kvdbserver.SetStringResponse{}, nil
}

// GetString gets a string value using a key.
// Accepts database name in gRPC metadata.
func (s *Server) GetString(ctx context.Context, req *kvdbserver.GetStringRequest) (res *kvdbserver.GetStringResponse, err error) {
	s.logger.Debug("Attempt to get string value")
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to get string value: %s", err)
		} else {
			s.logger.Debug("Get string value success")
		}
	}()

	dbName, err := getDatabaseNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	value, found := s.databases[dbName].GetString(kvdb.DatabaseKey(req.GetKey()))

	return &kvdbserver.GetStringResponse{Value: string(value), Found: found}, nil
}

// DeleteKey deletes a key and its value.
// Accepts database name in gRPC metadata.
func (s *Server) DeleteKey(ctx context.Context, req *kvdbserver.DeleteKeyRequest) (res *kvdbserver.DeleteKeyResponse, err error) {
	s.logger.Debug("Attempt to delete key")
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to delete key: %s", err)
		} else {
			s.logger.Debug("Delete key success")
		}
	}()

	dbName, err := getDatabaseNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	ok := s.databases[dbName].DeleteKey(kvdb.DatabaseKey(req.GetKey()))
	if !ok {
		return &kvdbserver.DeleteKeyResponse{Ok: false}, nil
	}

	return &kvdbserver.DeleteKeyResponse{Ok: true}, nil
}

// DeleteAllKeys deletes all the keys of a database.
// Accepts database name in gRPC metadata.
func (s *Server) DeleteAllKeys(ctx context.Context, req *kvdbserver.DeleteAllKeysRequest) (res *kvdbserver.DeleteAllKeysResponse, err error) {
	s.logger.Debug("Attempt to delete all keys")
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to delete all keys: %s", err)
		} else {
			s.logger.Debug("Delete all keys success")
		}
	}()

	dbName, err := getDatabaseNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	s.databases[dbName].DeleteAllKeys()

	return &kvdbserver.DeleteAllKeysResponse{}, nil
}

// GetKeys returns all the keys of a database.
// Accepts database name in gRPC metadata.
func (s *Server) GetKeys(ctx context.Context, req *kvdbserver.GetKeysRequest) (res *kvdbserver.GetKeysResponse, err error) {
	s.logger.Debug("Attempt to get keys")
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to get keys: %s", err)
		} else {
			s.logger.Debug("Get keys success")
		}
	}()

	dbName, err := getDatabaseNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	return &kvdbserver.GetKeysResponse{Keys: s.databases[dbName].GetKeys()}, nil
}

// SetHashMap sets fields in a HashMap value using a key, overwriting previous fields.
// Accepts database name in gRPC metadata.
func (s *Server) SetHashMap(ctx context.Context, req *kvdbserver.SetHashMapRequest) (res *kvdbserver.SetHashMapResponse, err error) {
	s.logger.Debug("Attempt to set HashMap fields")
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to set HashMap fields: %s", err)
		} else {
			s.logger.Debug("Set HashMap fields success")
		}
	}()

	dbName, err := getDatabaseNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	err = s.databases[dbName].SetHashMap(kvdb.DatabaseKey(req.Key), req.Fields)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &kvdbserver.SetHashMapResponse{}, nil
}

// GetHashMapFieldValue returns a single HashMap field value using a key.
// Accepts database name in gRPC metadata.
func (s *Server) GetHashMapFieldValue(ctx context.Context, req *kvdbserver.GetHashMapFieldValueRequest) (res *kvdbserver.GetHashMapFieldValueResponse, err error) {
	s.logger.Debug("Attempt to get HashMap field value")
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to get HashMap field value: %s", err)
		} else {
			s.logger.Debug("Get HashMap field value success")
		}
	}()

	dbName, err := getDatabaseNameFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	value, ok := s.databases[dbName].GetHashMapFieldValue(kvdb.DatabaseKey(req.Key), req.Field)

	return &kvdbserver.GetHashMapFieldValueResponse{Value: value, Ok: ok}, nil
}

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
