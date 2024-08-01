package grpc

import (
	"context"

	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"google.golang.org/grpc"
)

// newAuthUnaryInterceptor returns unary interceptor to handle RPC call authorization.
func newAuthUnaryInterceptor(s *server.KvdbServer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := s.AuthorizeIncomingRpcCall(ctx); err != nil {
			logger := s.Logger()
			logger.Errorf("Failed to authorize request: %v", err)

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
		logger.Debugf("%s: (call) db=%s; req=%v", info.FullMethod, dbName, req)
		resp, err := handler(ctx, req)
		if err != nil {
			logger.Errorf("%s: (fail) %v", info.FullMethod, err)
		}
		logger.Debugf("%s: (success) db=%s; req=%v; resp=%v", info.FullMethod, dbName, req, resp)
		return resp, err
	}
}