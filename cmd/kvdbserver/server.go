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
