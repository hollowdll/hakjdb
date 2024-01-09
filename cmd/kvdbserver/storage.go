package main

import (
	"context"

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
func (s *server) SetString(ctx context.Context, req *kvdbserver.SetStringRequest) (res *kvdbserver.SetStringResponse, err error) {
	s.logger.Debug("Attempt to set string value")
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to set string value: %s", err)
		} else {
			s.logger.Debug("Set string value success")
		}
	}()

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

	err = s.databases[dbName[0]].SetString(kvdb.DatabaseKey(req.GetKey()), kvdb.DatabaseStringValue(req.GetValue()))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s", err)
	}

	return &kvdbserver.SetStringResponse{}, nil
}

// GetString gets a string value using a key.
// Accepts database name in gRPC metadata.
func (s *server) GetString(ctx context.Context, req *kvdbserver.GetStringRequest) (res *kvdbserver.GetStringResponse, err error) {
	s.logger.Debug("Attempt to get string value")
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to get string value: %s", err)
		} else {
			s.logger.Debug("Get string value success")
		}
	}()

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
	value, found := s.databases[dbName[0]].GetString(kvdb.DatabaseKey(req.GetKey()))

	return &kvdbserver.GetStringResponse{Value: string(value), Found: found}, nil
}

// DeleteKey deletes a key and its value.
// Accepts database name in gRPC metadata.
func (s *server) DeleteKey(ctx context.Context, req *kvdbserver.DeleteKeyRequest) (res *kvdbserver.DeleteKeyResponse, err error) {
	s.logger.Debug("Attempt to delete key")
	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to delete key: %s", err)
		} else {
			s.logger.Debug("Delete key success")
		}
	}()

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

	return &kvdbserver.DeleteKeyResponse{Success: true}, nil
}
