package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	kvdb "github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type databaseServer struct {
	kvdbserver.UnimplementedDatabaseServer
	databases map[string]*kvdb.Database
	logger    kvdb.Logger
	mutex     sync.RWMutex
}

func newDatabaseServer() *databaseServer {
	return &databaseServer{
		databases: make(map[string]*kvdb.Database),
		logger:    *kvdb.NewLogger(),
	}
}

func (s *databaseServer) databaseExists(name string) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	_, exists := s.databases[name]
	return exists
}

// Creates a new database to the database server.
// Fails if it already exists or the name is invalid.
func (s *databaseServer) CreateDatabase(ctx context.Context, req *kvdbserver.CreateDatabaseRequest) (*kvdbserver.CreateDatabaseResponse, error) {
	log.Printf("Creating new database: %s", req.GetName())

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

	err = s.logger.LogMessage(kvdb.LogTypeInfo, fmt.Sprintf("Created database: %s", db.Name))
	if err != nil {
		log.Printf("error: failed to write to log file: %s", err)
	}
	log.Printf("Created new database: %s", db.Name)

	return &kvdbserver.CreateDatabaseResponse{Name: db.Name}, nil
}
