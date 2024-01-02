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

// SetValue sets a string value using a key.
func (s *server) SetValue(ctx context.Context, req *kvdbserver.SetValueRequest) (*kvdbserver.SetValueResponse, error) {
	log.Printf("attempt to set value")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "missing metadata")
	}

	dbName := md.Get(common.GrpcMetadataKeyDbName)
	if len(dbName) == 0 {
		errMsg := fmt.Sprintf("missing key in metadata: %s", common.GrpcMetadataKeyDbName)
		return nil, status.Error(codes.InvalidArgument, errMsg)
	}

	s.databases[dbName[0]].SetString(kvdb.DatabaseKey(req.GetKey()), kvdb.DatabaseStringValue(req.GetValue()))

	logMsg := fmt.Sprintf("set value with key '%s' in database: %s", req.GetKey(), dbName[0])
	log.Print(logMsg)

	err := s.logger.LogMessage(kvdb.LogTypeInfo, logMsg)
	if err != nil {
		log.Printf("error: failed to write to log file: %s", err)
	}

	return &kvdbserver.SetValueResponse{}, nil
}
