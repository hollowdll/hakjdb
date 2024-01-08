package main

import (
	"context"
	"fmt"
	"log"

	kvdb "github.com/hollowdll/kvdb"
	kvdberrors "github.com/hollowdll/kvdb/errors"
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
		return nil, status.Errorf(codes.InvalidArgument, "%s", kvdberrors.ErrMissingMetadata)
	}

	dbName := md.Get(common.GrpcMetadataKeyDbName)
	if len(dbName) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%s (%s)", kvdberrors.ErrMissingKeyInMetadata, common.GrpcMetadataKeyDbName)
	}

	if !s.databaseExists(dbName[0]) {
		return nil, status.Errorf(codes.NotFound, "%s", kvdberrors.ErrDatabaseNotFound)
	}

	err := s.databases[dbName[0]].SetString(kvdb.DatabaseKey(req.GetKey()), kvdb.DatabaseStringValue(req.GetValue()))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s", err)
	}

	logMsg := fmt.Sprintf("set value with key '%s' in database '%s'", req.GetKey(), dbName[0])
	log.Print(logMsg)

	return &kvdbserver.SetStringResponse{}, nil
}

// GetString gets a string value using a key.
// Accepts database name in gRPC metadata.
func (s *server) GetString(ctx context.Context, req *kvdbserver.GetStringRequest) (*kvdbserver.GetStringResponse, error) {
	log.Print("attempt to get value")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "%s", kvdberrors.ErrMissingMetadata)
	}

	dbName := md.Get(common.GrpcMetadataKeyDbName)
	if len(dbName) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%s (%s)", kvdberrors.ErrMissingKeyInMetadata, common.GrpcMetadataKeyDbName)
	}

	if !s.databaseExists(dbName[0]) {
		return nil, status.Errorf(codes.NotFound, "%s", kvdberrors.ErrDatabaseNotFound)
	}

	value := s.databases[dbName[0]].GetString(kvdb.DatabaseKey(req.GetKey()))

	logMsg := fmt.Sprintf("get value with key '%s' in database '%s'", req.GetKey(), dbName[0])
	log.Print(logMsg)

	return &kvdbserver.GetStringResponse{Value: string(value)}, nil
}

// DeleteKey deletes a key and its value.
// Accepts database name in gRPC metadata.
func (s *server) DeleteKey(ctx context.Context, req *kvdbserver.DeleteKeyRequest) (*kvdbserver.DeleteKeyResponse, error) {
	log.Print("attempt to delete key")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "%s", kvdberrors.ErrMissingMetadata)
	}

	dbName := md.Get(common.GrpcMetadataKeyDbName)
	if len(dbName) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%s (%s)", kvdberrors.ErrMissingKeyInMetadata, common.GrpcMetadataKeyDbName)
	}

	if !s.databaseExists(dbName[0]) {
		return nil, status.Errorf(codes.NotFound, "%s", kvdberrors.ErrDatabaseNotFound)
	}

	success := s.databases[dbName[0]].DeleteKey(kvdb.DatabaseKey(req.GetKey()))
	if !success {
		return &kvdbserver.DeleteKeyResponse{Success: false}, nil
	}

	logMsg := fmt.Sprintf("deleted key '%s' in database '%s'", req.GetKey(), dbName[0])
	log.Print(logMsg)

	return &kvdbserver.DeleteKeyResponse{Success: true}, nil
}
