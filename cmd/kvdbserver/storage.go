package main

import (
	"context"
	"fmt"
	"log"

	kvdb "github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// SetString sets a string value using a key.
// Accepts database name in gRPC metadata.
func (s *server) SetString(ctx context.Context, req *kvdbserver.SetStringRequest) (*kvdbserver.SetStringResponse, error) {
	log.Print("attempt to set value")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "missing metadata")
	}

	dbName := md.Get(common.GrpcMetadataKeyDbName)
	if len(dbName) == 0 {
		errMsg := fmt.Sprintf("missing key in metadata: %s", common.GrpcMetadataKeyDbName)
		return nil, status.Error(codes.InvalidArgument, errMsg)
	}

	if !s.databaseExists(dbName[0]) {
		errMsg := "database doesn't exist"
		return nil, status.Error(codes.NotFound, errMsg)
	}

	err := s.databases[dbName[0]].SetString(kvdb.DatabaseKey(req.GetKey()), kvdb.DatabaseStringValue(req.GetValue()))
	if err != nil {
		errMsg := fmt.Sprintf("%s", err)
		return nil, status.Error(codes.InvalidArgument, errMsg)
	}

	logMsg := fmt.Sprintf("set value with key '%s' in database '%s'", req.GetKey(), dbName[0])
	log.Print(logMsg)

	err = s.logger.LogMessage(kvdb.LogTypeInfo, logMsg)
	if err != nil {
		log.Printf("error: failed to write to log file: %s", err)
	}

	return &kvdbserver.SetStringResponse{}, nil
}

// GetString gets a string value using a key.
// Accepts database name in gRPC metadata.
func (s *server) GetString(ctx context.Context, req *kvdbserver.GetStringRequest) (*kvdbserver.GetStringResponse, error) {
	log.Print("attempt to get value")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "missing metadata")
	}

	dbName := md.Get(common.GrpcMetadataKeyDbName)
	if len(dbName) == 0 {
		errMsg := fmt.Sprintf("missing key in metadata: %s", common.GrpcMetadataKeyDbName)
		return nil, status.Error(codes.InvalidArgument, errMsg)
	}

	if !s.databaseExists(dbName[0]) {
		errMsg := "database doesn't exist"
		return nil, status.Error(codes.NotFound, errMsg)
	}

	value := s.databases[dbName[0]].GetString(kvdb.DatabaseKey(req.GetKey()))

	logMsg := fmt.Sprintf("Get value with key '%s' in database '%s'", req.GetKey(), dbName[0])
	log.Print(logMsg)

	err := s.logger.LogMessage(kvdb.LogTypeInfo, logMsg)
	if err != nil {
		log.Printf("error: failed to write to log file: %s", err)
	}

	return &kvdbserver.GetStringResponse{Value: string(value)}, nil
}

// DeleteKey deletes a key and its value.
// Accepts database name in gRPC metadata.
func (s *server) DeleteKey(ctx context.Context, req *kvdbserver.DeleteKeyRequest) (*kvdbserver.DeleteKeyResponse, error) {
	log.Print("attempt to delete key")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "missing metadata")
	}

	dbName := md.Get(common.GrpcMetadataKeyDbName)
	if len(dbName) == 0 {
		errMsg := fmt.Sprintf("missing key in metadata: %s", common.GrpcMetadataKeyDbName)
		return nil, status.Error(codes.InvalidArgument, errMsg)
	}

	if !s.databaseExists(dbName[0]) {
		errMsg := "database doesn't exist"
		return nil, status.Error(codes.NotFound, errMsg)
	}

	success := s.databases[dbName[0]].DeleteKey(kvdb.DatabaseKey(req.GetKey()))
	if !success {
		return &kvdbserver.DeleteKeyResponse{Success: false}, nil
	}

	logMsg := fmt.Sprintf("deleted key '%s' in database '%s'", req.GetKey(), dbName[0])
	log.Print(logMsg)

	err := s.logger.LogMessage(kvdb.LogTypeInfo, logMsg)
	if err != nil {
		log.Printf("error: failed to write to log file: %s", err)
	}

	return &kvdbserver.DeleteKeyResponse{Success: true}, nil
}
