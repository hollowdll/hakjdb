package main

import (
	"sync"

	kvdb "github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
)

type server struct {
	kvdbserver.UnimplementedDatabaseServer
	databases map[string]*kvdb.Database
	logger    kvdb.Logger
	mutex     sync.RWMutex
}

func newServer() *server {
	return &server{
		databases: make(map[string]*kvdb.Database),
		logger:    *kvdb.NewLogger(),
	}
}

func (s *server) databaseExists(name string) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	_, exists := s.databases[name]
	return exists
}
