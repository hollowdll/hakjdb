package grpc

import (
	"context"

	"github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/version"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// newMetadataUnaryInterceptor returns unary interceptor that sends metadata back to the client.
func newHeaderUnaryInterceptor(s *server.KvdbServer) grpc.UnaryServerInterceptor {
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
func newAuthUnaryInterceptor(s *server.KvdbServer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := s.AuthorizeIncomingRpcCall(ctx); err != nil {
			logger := s.Logger()
			logger.Debugf("Failed to authorize request: %v", err)
			return nil, err
		}
		return handler(ctx, req)
	}
}

// newUnaryAuthInterceptor returns unary interceptor to handle RPC call logging.
func newLogUnaryInterceptor(s *server.KvdbServer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger := s.Logger()
		dbName := s.GetDBNameFromContext(ctx)
		logRequestCall(logger, s.Cfg.VerboseLogsEnabled, info.FullMethod, dbName, req)
		resp, err := handler(ctx, req)
		if err != nil {
			logRequestFailed(logger, s.Cfg.VerboseLogsEnabled, info.FullMethod, dbName, req, err)
		} else {
			logRequestSuccess(logger, s.Cfg.VerboseLogsEnabled, info.FullMethod, dbName, req, resp)
		}
		return resp, err
	}
}

func logRequestCall(logger kvdb.Logger, verbose bool, fullMethod string, dbName string, req any) {
	if verbose {
		logger.Debugf("(call) %s: db = %s; req = %v", fullMethod, dbName, req)
	} else {
		logger.Debugf("(call) %s", fullMethod)
	}
}

func logRequestFailed(logger kvdb.Logger, verbose bool, fullMethod string, dbName string, req any, err error) {
	if verbose {
		logger.Debugf("(failed) %s: db = %s; req = %v; error = %v", fullMethod, dbName, req, err)
	} else {
		logger.Debugf("(failed) %s: %v", fullMethod, err)
	}
}

func logRequestSuccess(logger kvdb.Logger, verbose bool, fullMethod string, dbName string, req any, resp any) {
	if verbose {
		logger.Debugf("(success) %s: db = %s; req = %v; resp = %v", fullMethod, dbName, req, resp)
	} else {
		logger.Debugf("(success) %s", fullMethod)
	}
}
