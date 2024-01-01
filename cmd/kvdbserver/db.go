package main

import (
	"context"
	"fmt"
	"log"

	kvdb "github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
		errMsg := fmt.Sprintf("%s", err)
		return nil, status.Error(codes.InvalidArgument, errMsg)
	}

	if s.databaseExists(db.Name) {
		errMsg := fmt.Sprintf("database already exists: %s", db.Name)
		return nil, status.Error(codes.AlreadyExists, errMsg)
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.databases[db.Name] = db
	log.Printf("created database: %s", db.Name)

	err = s.logger.LogMessage(kvdb.LogTypeInfo, fmt.Sprintf("Created database: %s", db.Name))
	if err != nil {
		log.Printf("error: failed to write to log file: %s", err)
	}

	return &kvdbserver.CreateDatabaseResponse{Name: db.Name}, nil
}

// GetAllDatabases returns the names of all databases on the server.
func (s *server) GetAllDatabases(ctx context.Context, req *kvdbserver.Empty) (*kvdbserver.GetAllDatabasesResponse, error) {
	log.Printf("attempt to get all databases")

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var names []string
	for key := range s.databases {
		names = append(names, key)
	}
	log.Printf("get all databases")

	err := s.logger.LogMessage(kvdb.LogTypeInfo, "Get all databases")
	if err != nil {
		log.Printf("error: failed to write to log file: %s", err)
	}

	return &kvdbserver.GetAllDatabasesResponse{Names: names}, nil
}