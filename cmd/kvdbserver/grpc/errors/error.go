package errors

import (
	"context"

	kvdberrors "github.com/hollowdll/kvdb/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var grpcErrorMap = map[error]error{
	kvdberrors.ErrDatabaseNotFound:     status.Error(codes.NotFound, kvdberrors.ErrDatabaseNotFound.Error()),
	kvdberrors.ErrDatabaseExists:       status.Error(codes.AlreadyExists, kvdberrors.ErrDatabaseExists.Error()),
	kvdberrors.ErrDatabaseNameRequired: status.Error(codes.InvalidArgument, kvdberrors.ErrDatabaseNameRequired.Error()),
	kvdberrors.ErrDatabaseNameTooLong:  status.Error(codes.InvalidArgument, kvdberrors.ErrDatabaseNameTooLong.Error()),
	kvdberrors.ErrDatabaseNameInvalid:  status.Error(codes.InvalidArgument, kvdberrors.ErrDatabaseNameInvalid.Error()),

	kvdberrors.ErrDatabaseKeyRequired: status.Error(codes.InvalidArgument, kvdberrors.ErrDatabaseKeyRequired.Error()),
	kvdberrors.ErrDatabaseKeyTooLong:  status.Error(codes.InvalidArgument, kvdberrors.ErrDatabaseKeyTooLong.Error()),

	kvdberrors.ErrInvalidCredentials: status.Error(codes.Unauthenticated, kvdberrors.ErrInvalidCredentials.Error()),
	kvdberrors.ErrMaxKeysReached:     status.Error(codes.FailedPrecondition, kvdberrors.ErrMaxKeysReached.Error()),
	kvdberrors.ErrLogFileNotEnabled:  status.Error(codes.FailedPrecondition, kvdberrors.ErrLogFileNotEnabled.Error()),
	kvdberrors.ErrReadLogFile:        status.Error(codes.Internal, kvdberrors.ErrReadLogFile.Error()),
	kvdberrors.ErrGetOSInfo:          status.Error(codes.Internal, kvdberrors.ErrGetOSInfo.Error()),
}

// ToGrpcError converts error to the correct gRPC error status.
func ToGrpcError(err error) error {
	// gRPC maps these under the hood
	if err == context.Canceled || err == context.DeadlineExceeded {
		return err
	}
	grpcErr, ok := grpcErrorMap[err]
	if !ok {
		return status.Error(codes.Unknown, err.Error())
	}
	return grpcErr
}
