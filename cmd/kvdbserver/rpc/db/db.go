package db

import (
	"context"

	"github.com/hollowdll/kvdb/api/v0/dbpb"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	kvdberrors "github.com/hollowdll/kvdb/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	createDatabaseRPCName  string = "CreateDatabase"
	getAllDatabasesRPCName string = "GetAllDatabases"
	getDatabaseInfoRPCName string = "GetDatabaseInfo"
	deleteDatabaseRPCName  string = "DeleteDatabase"
)

type DBServiceServer struct {
	dbs server.DBService
	dbpb.UnimplementedDatabaseServiceServer
}

func NewDBServiceServer(s *server.KvdbServer) dbpb.DatabaseServiceServer {
	return &DBServiceServer{dbs: s}
}

// CreateDatabase is the implementation of RPC CreateDatabase.
func (s *DBServiceServer) CreateDatabase(ctx context.Context, req *dbpb.CreateDatabaseRequest) (res *dbpb.CreateDatabaseResponse, err error) {
	logger := s.dbs.Logger()
	logger.Debugf("%s: (call) %v", createDatabaseRPCName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", createDatabaseRPCName, err)
		} else {
			logger.Infof("Created database '%s'", req.DbName)
		}
	}()

	res, err = s.dbs.CreateDatabase(ctx, req)
	if err != nil {
		var code codes.Code
		switch err {
		case kvdberrors.ErrDatabaseExists:
			code = codes.AlreadyExists
		case kvdberrors.ErrDatabaseNameRequired, kvdberrors.ErrDatabaseNameTooLong, kvdberrors.ErrDatabaseNameInvalid:
			code = codes.InvalidArgument
		default:
			code = codes.Unknown
		}
		return nil, status.Error(code, err.Error())
	}

	return res, nil
}

// DeleteDatabase is the implementation of RPC DeleteDatabase.
func (s *DBServiceServer) DeleteDatabase(ctx context.Context, req *dbpb.DeleteDatabaseRequest) (res *dbpb.DeleteDatabaseResponse, err error) {
	logger := s.dbs.Logger()
	logger.Debugf("%s: (call) %v", deleteDatabaseRPCName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", deleteDatabaseRPCName, err)
		} else {
			logger.Infof("Deleted database '%s'", req.DbName)
		}
	}()

	res, err = s.dbs.DeleteDatabase(ctx, req)
	if err != nil {
		var code codes.Code
		switch err {
		case kvdberrors.ErrDatabaseNotFound:
			code = codes.NotFound
		default:
			code = codes.Unknown
		}
		return nil, status.Error(code, err.Error())
	}

	return res, nil
}

// GetAllDatabases is the implementation of RPC GetAllDatabases.
func (s *DBServiceServer) GetAllDatabases(ctx context.Context, req *dbpb.GetAllDatabasesRequest) (res *dbpb.GetAllDatabasesResponse, err error) {
	logger := s.dbs.Logger()
	logger.Debugf("%s: (call) %v", getAllDatabasesRPCName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", getAllDatabasesRPCName, err)
		} else {
			logger.Debugf("%s: (success) %v", getAllDatabasesRPCName, req)
		}
	}()

	res, err = s.dbs.GetAllDatabases(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return res, nil
}

// GetDatabaseInfo is the implementation of RPC GetDatabaseInfo.
func (s *DBServiceServer) GetDatabaseInfo(ctx context.Context, req *dbpb.GetDatabaseInfoRequest) (res *dbpb.GetDatabaseInfoResponse, err error) {
	logger := s.dbs.Logger()
	logger.Debugf("%s: (call) %v", getDatabaseInfoRPCName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", getDatabaseInfoRPCName, err)
		} else {
			logger.Debugf("%s: (success) %v", getDatabaseInfoRPCName, req)
		}
	}()

	res, err = s.dbs.GetDatabaseInfo(ctx, req)
	if err != nil {
		var code codes.Code
		switch err {
		case kvdberrors.ErrDatabaseNotFound:
			code = codes.NotFound
		default:
			code = codes.Unknown
		}
		return nil, status.Error(code, err.Error())
	}

	return res, nil
}
