package db

import (
	"context"

	"github.com/hollowdll/kvdb/api/v0/dbpb"
	grpcerrors "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc/errors"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/validation"
)

const (
	createDatabaseRPCName  string = "CreateDatabase"
	getAllDatabasesRPCName string = "GetAllDatabases"
	getDatabaseInfoRPCName string = "GetDatabaseInfo"
	deleteDatabaseRPCName  string = "DeleteDatabase"
)

type DBServiceServer struct {
	srv server.DBService
	dbpb.UnimplementedDBServiceServer
}

func NewDBServiceServer(s *server.KvdbServer) dbpb.DBServiceServer {
	return &DBServiceServer{srv: s}
}

// CreateDB is the implementation of RPC CreateDB.
func (s *DBServiceServer) CreateDB(ctx context.Context, req *dbpb.CreateDBRequest) (res *dbpb.CreateDBResponse, err error) {
	logger := s.srv.Logger()
	logger.Debugf("%s: (call) %v", createDatabaseRPCName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", createDatabaseRPCName, err)
		} else {
			logger.Infof("Created database '%s'", req.DbName)
		}
	}()

	if err = validation.ValidateDBName(req.DbName); err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	res, err = s.srv.CreateDB(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// DeleteDB is the implementation of RPC DeleteDB.
func (s *DBServiceServer) DeleteDB(ctx context.Context, req *dbpb.DeleteDBRequest) (res *dbpb.DeleteDBResponse, err error) {
	logger := s.srv.Logger()
	logger.Debugf("%s: (call) %v", deleteDatabaseRPCName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", deleteDatabaseRPCName, err)
		} else {
			logger.Infof("Deleted database '%s'", req.DbName)
		}
	}()

	res, err = s.srv.DeleteDB(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetAllDBs is the implementation of RPC GetAllDBs.
func (s *DBServiceServer) GetAllDBs(ctx context.Context, req *dbpb.GetAllDBsRequest) (res *dbpb.GetAllDBsResponse, err error) {
	logger := s.srv.Logger()
	logger.Debugf("%s: (call) %v", getAllDatabasesRPCName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", getAllDatabasesRPCName, err)
		} else {
			logger.Debugf("%s: (success) %v", getAllDatabasesRPCName, req)
		}
	}()

	res, err = s.srv.GetAllDBs(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetDBInfo is the implementation of RPC GetDBInfo.
func (s *DBServiceServer) GetDBInfo(ctx context.Context, req *dbpb.GetDBInfoRequest) (res *dbpb.GetDBInfoResponse, err error) {
	logger := s.srv.Logger()
	logger.Debugf("%s: (call) %v", getDatabaseInfoRPCName, req)
	defer func() {
		if err != nil {
			logger.Errorf("%s: operation failed: %v", getDatabaseInfoRPCName, err)
		} else {
			logger.Debugf("%s: (success) %v", getDatabaseInfoRPCName, req)
		}
	}()

	res, err = s.srv.GetDBInfo(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}
