package grpc

import (
	"context"

	"github.com/hollowdll/hakjdb"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/server"
	"github.com/hollowdll/hakjdb/internal/common"
	"github.com/hollowdll/hakjdb/version"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// newMetadataUnaryInterceptor returns unary interceptor that sends metadata back to the client.
func newHeaderUnaryInterceptor(s *server.HakjServer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		lg := s.Logger()
		md := metadata.Pairs(common.GrpcMetadataKeyAPIVersion, version.APIVersion)
		lg.Debugf("metadata to be sent to the client: %v", md)
		if err := grpc.SendHeader(ctx, md); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

// newAuthUnaryInterceptor returns unary interceptor to handle RPC call authorization.
func newAuthUnaryInterceptor(s *server.HakjServer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if !bypassAuthorization(info.FullMethod) {
			if err := s.AuthorizeIncomingRpcCall(ctx); err != nil {
				logger := s.Logger()
				logger.Debugf("Failed to authorize request: %v", err)
				return nil, err
			}
		}
		return handler(ctx, req)
	}
}

// newUnaryAuthInterceptor returns unary interceptor to handle RPC call logging.
func newLogUnaryInterceptor(s *server.HakjServer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger := s.Logger()
		cfg := s.Config()
		dbName := s.GetDBNameFromContext(ctx)
		logRequestCall(logger, cfg.VerboseLogsEnabled, info.FullMethod, dbName, req)
		resp, err := handler(ctx, req)
		if err != nil {
			logRequestFailed(logger, cfg.VerboseLogsEnabled, info.FullMethod, dbName, req, err)
		} else {
			logRequestSuccess(logger, cfg.VerboseLogsEnabled, info.FullMethod, dbName, req, resp)
		}
		return resp, err
	}
}

func logRequestCall(logger hakjdb.Logger, verbose bool, fullMethod string, dbName string, req any) {
	if verbose {
		if bypassDetailedLog(fullMethod) {
			logger.Debugf("(call) %s: db = %s", fullMethod, dbName)
		} else {
			logger.Debugf("(call) %s: db = %s; req = %v", fullMethod, dbName, req)
		}
	} else {
		logger.Debugf("(call) %s", fullMethod)
	}
}

func logRequestFailed(logger hakjdb.Logger, verbose bool, fullMethod string, dbName string, req any, err error) {
	if verbose {
		if bypassDetailedLog(fullMethod) {
			logger.Debugf("(failed) %s: db = %s; error = %v", fullMethod, dbName, err)
		} else {
			logger.Debugf("(failed) %s: db = %s; req = %v; error = %v", fullMethod, dbName, req, err)
		}
	} else {
		logger.Debugf("(failed) %s: %v", fullMethod, err)
	}
}

func logRequestSuccess(logger hakjdb.Logger, verbose bool, fullMethod string, dbName string, req any, resp any) {
	if verbose {
		if bypassDetailedLog(fullMethod) {
			logger.Debugf("(success) %s: db = %s", fullMethod, dbName)
		} else if bypassResponseDataLog(fullMethod) {
			logger.Debugf("(success) %s: db = %s; req = %v", fullMethod, dbName, req)
		} else {
			logger.Debugf("(success) %s: db = %s; req = %v; resp = %v", fullMethod, dbName, req, resp)
		}
	} else {
		logger.Debugf("(success) %s", fullMethod)
	}
}

func bypassAuthorization(fullMethod string) bool {
	return fullMethod == "/api.v1.authpb.AuthService/Authenticate"
}

// bypassDetailedLog is used to check if request and response data are logged.
// Mainly used to prevent logging sensitive data like passwords.
func bypassDetailedLog(fullMethod string) bool {
	return fullMethod == "/api.v1.authpb.AuthService/Authenticate"
}

func bypassResponseDataLog(fullMethod string) bool {
	return fullMethod == "/api.v1.serverpb.ServerService/GetLogs"
}
