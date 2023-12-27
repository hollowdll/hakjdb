package main

import (
	kvdb "github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
)

type databaseServer struct {
	databases []kvdb.Database
	kvdbserver.UnimplementedDatabaseServer
}

func newDatabaseServer() *databaseServer {
	return &databaseServer{}
}
