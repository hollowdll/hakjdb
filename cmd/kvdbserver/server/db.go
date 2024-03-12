package server

import (
	"context"

	kvdb "github.com/hollowdll/kvdb"
	kvdberrors "github.com/hollowdll/kvdb/errors"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// databaseExists returns true if a database exists on the server
func (s *Server) databaseExists(name string) bool {
	_, exists := s.databases[name]
	return exists
}

// CreateDatabase creates a new database to the server.
// Fails if it already exists or the name is invalid.
func (s *Server) CreateDatabase(ctx context.Context, req *kvdbserver.CreateDatabaseRequest) (res *kvdbserver.CreateDatabaseResponse, err error) {
	s.logger.Debugf("Attempt to create database '%s'", req.GetDbName())
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to create database '%s': %v", req.GetDbName(), err)
		} else {
			s.logger.Infof("Created database '%s'", req.GetDbName())
		}
	}()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.databaseExists(req.GetDbName()) {
		return nil, status.Error(codes.AlreadyExists, kvdberrors.ErrDatabaseExists.Error())
	}

	if err := kvdb.ValidateDatabaseName(req.GetDbName()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	db := kvdb.CreateDatabase(req.GetDbName())
	s.databases[db.Name] = db

	return &kvdbserver.CreateDatabaseResponse{DbName: db.Name}, nil
}

// GetAllDatabases returns the names of all databases on the server.
func (s *Server) GetAllDatabases(ctx context.Context, req *kvdbserver.GetAllDatabasesRequest) (res *kvdbserver.GetAllDatabasesResponse, err error) {
	s.logger.Debug("Attempt to get all databases")
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to get all databases: %v", err)
		} else {
			s.logger.Debug("Get all databases success")
		}
	}()

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var names []string
	for key := range s.databases {
		names = append(names, key)
	}

	return &kvdbserver.GetAllDatabasesResponse{DbNames: names}, nil
}

// GetDatabaseInfo returns information about a database.
func (s *Server) GetDatabaseInfo(ctx context.Context, req *kvdbserver.GetDatabaseInfoRequest) (res *kvdbserver.GetDatabaseInfoResponse, err error) {
	s.logger.Debugf("Attempt to get info for database '%s'", req.GetDbName())
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to get info for database '%s': %v", req.GetDbName(), err)
		} else {
			s.logger.Debugf("Get info for database '%s' success", req.GetDbName())
		}
	}()

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.databaseExists(req.GetDbName()) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	db := s.databases[req.GetDbName()]
	data := &kvdbserver.DatabaseInfo{
		Name:      db.Name,
		CreatedAt: timestamppb.New(db.CreatedAt),
		UpdatedAt: timestamppb.New(db.UpdatedAt),
		KeyCount:  db.GetKeyCount(),
		DataSize:  db.GetStoredSizeBytes(),
	}

	return &kvdbserver.GetDatabaseInfoResponse{Data: data}, nil
}

// DeleteDatabase deletes a database if it exists.
func (s *Server) DeleteDatabase(ctx context.Context, req *kvdbserver.DeleteDatabaseRequest) (res *kvdbserver.DeleteDatabaseResponse, err error) {
	dbName := req.GetDbName()
	s.logger.Debugf("Attempt to delete database '%s'", dbName)
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to delete database '%s': %v", dbName, err)
		} else {
			s.logger.Infof("Deleted database '%s'", dbName)
		}
	}()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.databaseExists(dbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	delete(s.databases, dbName)

	return &kvdbserver.DeleteDatabaseResponse{DbName: dbName}, nil
}
