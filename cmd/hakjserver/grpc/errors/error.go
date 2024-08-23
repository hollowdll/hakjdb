package errors

import (
	"context"

	errors "github.com/hollowdll/hakjdb/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var grpcErrorMap = map[error]error{
	errors.ErrDatabaseNotFound:           status.Error(codes.NotFound, errors.ErrDatabaseNotFound.Error()),
	errors.ErrDatabaseExists:             status.Error(codes.AlreadyExists, errors.ErrDatabaseExists.Error()),
	errors.ErrDatabaseNameRequired:       status.Error(codes.InvalidArgument, errors.ErrDatabaseNameRequired.Error()),
	errors.ErrDatabaseNameTooLong:        status.Error(codes.InvalidArgument, errors.ErrDatabaseNameTooLong.Error()),
	errors.ErrDatabaseNameInvalid:        status.Error(codes.InvalidArgument, errors.ErrDatabaseNameInvalid.Error()),
	errors.ErrDatabaseDescriptionTooLong: status.Error(codes.InvalidArgument, errors.ErrDatabaseDescriptionTooLong.Error()),

	errors.ErrDatabaseKeyRequired: status.Error(codes.InvalidArgument, errors.ErrDatabaseKeyRequired.Error()),
	errors.ErrDatabaseKeyTooLong:  status.Error(codes.InvalidArgument, errors.ErrDatabaseKeyTooLong.Error()),

	errors.ErrInvalidCredentials: status.Error(codes.InvalidArgument, errors.ErrInvalidCredentials.Error()),
	errors.ErrInvalidAuthToken:   status.Error(codes.Unauthenticated, errors.ErrInvalidAuthToken.Error()),
	errors.ErrAuthFailed:         status.Error(codes.InvalidArgument, errors.ErrAuthFailed.Error()),
	errors.ErrAuthNotEnabled:     status.Error(codes.FailedPrecondition, errors.ErrAuthNotEnabled.Error()),

	errors.ErrMaxKeysReached:    status.Error(codes.FailedPrecondition, errors.ErrMaxKeysReached.Error()),
	errors.ErrLogFileNotEnabled: status.Error(codes.FailedPrecondition, errors.ErrLogFileNotEnabled.Error()),
	errors.ErrReadLogFile:       status.Error(codes.Internal, errors.ErrReadLogFile.Error()),
	errors.ErrGetOSInfo:         status.Error(codes.Internal, errors.ErrGetOSInfo.Error()),
	errors.ErrMissingMetadata:   status.Error(codes.InvalidArgument, errors.ErrMissingMetadata.Error()),
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
