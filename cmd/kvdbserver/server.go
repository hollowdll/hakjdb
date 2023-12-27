package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	kvdb "github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type databaseServer struct {
	databases map[string]*kvdb.Database
	kvdbserver.UnimplementedDatabaseServer
	logger kvdb.Logger
	mutex  sync.RWMutex
}

func newDatabaseServer() *databaseServer {
	return &databaseServer{}
}

// Creates a new database to the database server.
// Fails if it already exists or the name is invalid.
func (s *databaseServer) CreateDatabase(ctx context.Context, req *kvdbserver.CreateDatabaseRequest) (*kvdbserver.CreateDatabaseResponse, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	log.Printf("Creating new database with name %v", req.GetName())

	errPrefix := "cannot create database"
	db, err := kvdb.CreateDatabase(req.GetName())
	if err != nil {
		errMsg := fmt.Sprintf("%s: %s", errPrefix, err)
		fmt.Fprintln(os.Stderr, "Error:", errMsg)

		return nil, status.Error(codes.InvalidArgument, errMsg)
	}

	_, exists := s.databases[db.Name]
	if exists {
		errMsg := fmt.Sprintf("%s: database already exists: %s", errPrefix, db.Name)
		fmt.Fprintln(os.Stderr, "Error:", errMsg)

		return nil, status.Error(codes.AlreadyExists, errMsg)
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.databases[db.Name] = db

	err = s.logger.LogMessage(kvdb.LogTypeInfo, fmt.Sprintf("Created database %s", db.Name))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: logging failed:", err)
	}

	return &kvdbserver.CreateDatabaseResponse{Name: db.Name}, nil
}
