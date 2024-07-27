package server

import (
	"context"

	kvdb "github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/api/v0/dbpb"
	kvdberrors "github.com/hollowdll/kvdb/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	createDatabaseRPCName  string = "CreateDatabase"
	getAllDatabasesRPCName string = "GetAllDatabases"
	getDatabaseInfoRPCName string = "GetDatabaseInfo"
	deleteDatabaseRPCName  string = "DeleteDatabase"
)

// databaseExists returns true if a database exists on the server
func (s *Server) databaseExists(name string) bool {
	_, exists := s.databases[name]
	return exists
}

// CreateDatabase is the implementation of RPC CreateDatabase.
func (s *Server) CreateDatabase(ctx context.Context, req *dbpb.CreateDatabaseRequest) (res *dbpb.CreateDatabaseResponse, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.logger.Debugf("%s: (call) %v", createDatabaseRPCName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", createDatabaseRPCName, err)
		} else {
			s.logger.Infof("Created database '%s'", req.DbName)
		}
	}()

	if s.databaseExists(req.DbName) {
		return nil, status.Error(codes.AlreadyExists, kvdberrors.ErrDatabaseExists.Error())
	}

	if err := kvdb.ValidateDatabaseName(req.DbName); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	db := kvdb.CreateDatabase(req.DbName)
	s.databases[db.Name] = db

	return &dbpb.CreateDatabaseResponse{DbName: db.Name}, nil
}

// GetAllDatabases is the implementation of RPC GetAllDatabases.
func (s *Server) GetAllDatabases(ctx context.Context, req *dbpb.GetAllDatabasesRequest) (res *dbpb.GetAllDatabasesResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	s.logger.Debugf("%s: (call) %v", getAllDatabasesRPCName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", getAllDatabasesRPCName, err)
		} else {
			s.logger.Debugf("%s: (success) %v", getAllDatabasesRPCName, req)
		}
	}()

	var names []string
	for key := range s.databases {
		names = append(names, key)
	}

	return &dbpb.GetAllDatabasesResponse{DbNames: names}, nil
}

// GetDatabaseInfo is the implementation of RPC GetDatabaseInfo.
func (s *Server) GetDatabaseInfo(ctx context.Context, req *dbpb.GetDatabaseInfoRequest) (res *dbpb.GetDatabaseInfoResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	s.logger.Debugf("%s: (call) %v", getDatabaseInfoRPCName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", getDatabaseInfoRPCName, err)
		} else {
			s.logger.Debugf("%s: (success) %v", getDatabaseInfoRPCName, req)
		}
	}()

	if !s.databaseExists(req.DbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	db := s.databases[req.DbName]
	data := &dbpb.DatabaseInfo{
		Name:      db.Name,
		CreatedAt: timestamppb.New(db.CreatedAt),
		UpdatedAt: timestamppb.New(db.UpdatedAt),
		KeyCount:  db.GetKeyCount(),
		DataSize:  db.GetStoredSizeBytes(),
	}

	return &dbpb.GetDatabaseInfoResponse{Data: data}, nil
}

// DeleteDatabase is the implementation of RPC DeleteDatabase.
func (s *Server) DeleteDatabase(ctx context.Context, req *dbpb.DeleteDatabaseRequest) (res *dbpb.DeleteDatabaseResponse, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.logger.Debugf("%s: (call) %v", deleteDatabaseRPCName, req)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: operation failed: %v", deleteDatabaseRPCName, err)
		} else {
			s.logger.Infof("Deleted database '%s'", req.DbName)
		}
	}()

	if !s.databaseExists(req.DbName) {
		return nil, status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error())
	}

	delete(s.databases, req.DbName)

	return &dbpb.DeleteDatabaseResponse{DbName: req.DbName}, nil
}
