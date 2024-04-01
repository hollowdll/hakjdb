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

// CreateDatabase is the implementation of RPC CreateDatabase.
func (s *Server) CreateDatabase(ctx context.Context, req *kvdbserver.CreateDatabaseRequest) (res *kvdbserver.CreateDatabaseResponse, err error) {
	logPrefix := "CreateDatabase"
	s.logger.Debugf("%s: (attempt) %v", logPrefix, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", logPrefix, err)
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

// GetAllDatabases is the implementation of RPC GetAllDatabases.
func (s *Server) GetAllDatabases(ctx context.Context, req *kvdbserver.GetAllDatabasesRequest) (res *kvdbserver.GetAllDatabasesResponse, err error) {
	logPrefix := "GetAllDatabases"
	s.logger.Debugf("%s: (attempt) %v", logPrefix, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: (success) %v", logPrefix, req)
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

// GetDatabaseInfo is the implementation of RPC GetDatabaseInfo.
func (s *Server) GetDatabaseInfo(ctx context.Context, req *kvdbserver.GetDatabaseInfoRequest) (res *kvdbserver.GetDatabaseInfoResponse, err error) {
	logPrefix := "GetDatabaseInfo"
	s.logger.Debugf("%s: (attempt) %v", logPrefix, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: (success) %v", logPrefix, req)
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

// DeleteDatabase is the implementation of RPC DeleteDatabase.
func (s *Server) DeleteDatabase(ctx context.Context, req *kvdbserver.DeleteDatabaseRequest) (res *kvdbserver.DeleteDatabaseResponse, err error) {
	logPrefix := "DeleteDatabase"
	s.logger.Debugf("%s: (attempt) %v", logPrefix, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", logPrefix, err)
		} else {
			s.logger.Infof("Deleted database '%s'", req.GetDbName())
		}
	}()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.databaseExists(req.GetDbName()) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	delete(s.databases, req.GetDbName())

	return &kvdbserver.DeleteDatabaseResponse{DbName: req.GetDbName()}, nil
}
