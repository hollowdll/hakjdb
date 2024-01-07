package main

import (
	"context"
	"fmt"
	"log"

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
func (s *server) CreateDatabase(ctx context.Context, req *kvdbserver.CreateDatabaseRequest) (*kvdbserver.CreateDatabaseResponse, error) {
	log.Printf("attempt to create database: %s", req.GetName())

	db, err := kvdb.CreateDatabase(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s", err)
	}

	if s.databaseExists(db.Name) {
		return nil, status.Errorf(codes.AlreadyExists, "%s", kvdberrors.ErrDatabaseExists)
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.databases[db.Name] = db

	logMsg := fmt.Sprintf("created database: %s", db.Name)
	log.Print(logMsg)

	err = s.logger.LogMessage(kvdb.LogTypeInfo, logMsg)
	if err != nil {
		log.Printf("%s: %s", kvdberrors.ErrWriteLogFile, err)
	}

	return &kvdbserver.CreateDatabaseResponse{Name: db.Name}, nil
}

// GetAllDatabases returns the names of all databases on the server.
func (s *server) GetAllDatabases(ctx context.Context, req *kvdbserver.GetAllDatabasesRequest) (*kvdbserver.GetAllDatabasesResponse, error) {
	log.Printf("attempt to get all databases")

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var names []string
	for key := range s.databases {
		names = append(names, key)
	}

	logMsg := "get all databases"
	log.Print(logMsg)

	err := s.logger.LogMessage(kvdb.LogTypeInfo, logMsg)
	if err != nil {
		log.Printf("%s: %s", kvdberrors.ErrWriteLogFile, err)
	}

	return &kvdbserver.GetAllDatabasesResponse{Names: names}, nil
}

func (s *server) GetDatabaseMetadata(ctx context.Context, req *kvdbserver.GetDatabaseMetadataRequest) (*kvdbserver.GetDatabaseMetadataResponse, error) {
	if !s.databaseExists(req.GetDbName()) {
		return nil, status.Errorf(codes.NotFound, "%s", kvdberrors.ErrDatabaseNotFound)
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	db := s.databases[req.GetDbName()]
	data := &kvdbserver.DatabaseMetadata{
		Name:      db.Name,
		CreatedAt: timestamppb.New(db.CreatedAt),
		UpdatedAt: timestamppb.New(db.UpdatedAt),
		KeyCount:  db.GetKeyCount(),
		DataSize:  db.GetStoredSizeBytes(),
	}

	return &kvdbserver.GetDatabaseMetadataResponse{Data: data}, nil
}
