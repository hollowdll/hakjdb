package main

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
func (s *server) databaseExists(name string) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	_, exists := s.databases[name]
	return exists
}

// CreateDatabase creates a new database to the server.
// Fails if it already exists or the name is invalid.
func (s *server) CreateDatabase(ctx context.Context, req *kvdbserver.CreateDatabaseRequest) (res *kvdbserver.CreateDatabaseResponse, err error) {
	s.logger.Debugf("Attempt to create database '%s'", req.GetDbName())
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to create database '%s': %s", req.GetDbName(), err)
		} else {
			s.logger.Infof("Created database '%s'", req.GetDbName())
		}
	}()

	db, err := kvdb.CreateDatabase(req.GetDbName())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s", err)
	}

	if s.databaseExists(db.Name) {
		return nil, status.Errorf(codes.AlreadyExists, "%s", kvdberrors.ErrDatabaseExists)
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.databases[db.Name] = db

	return &kvdbserver.CreateDatabaseResponse{DbName: db.Name}, nil
}

// GetAllDatabases returns the names of all databases on the server.
func (s *server) GetAllDatabases(ctx context.Context, req *kvdbserver.GetAllDatabasesRequest) (res *kvdbserver.GetAllDatabasesResponse, err error) {
	s.logger.Debug("Attempt to get all databases")
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to get all databases: %s", err)
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

func (s *server) GetDatabaseInfo(ctx context.Context, req *kvdbserver.GetDatabaseInfoRequest) (res *kvdbserver.GetDatabaseInfoResponse, err error) {
	s.logger.Debugf("Attempt to get info for database '%s'", req.GetDbName())
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to get info for database '%s': %s", req.GetDbName(), err)
		} else {
			s.logger.Debugf("Get info for database '%s' success", req.GetDbName())
		}
	}()

	if !s.databaseExists(req.GetDbName()) {
		return nil, status.Errorf(codes.NotFound, "%s", kvdberrors.ErrDatabaseNotFound)
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

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
