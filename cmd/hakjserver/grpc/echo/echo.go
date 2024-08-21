package echo

import (
	"context"

	"github.com/hollowdll/hakjdb/api/v1/echopb"
)

type EchoServiceServer struct {
	echopb.UnimplementedEchoServiceServer
}

func NewEchoServiceServer() echopb.EchoServiceServer {
	return &EchoServiceServer{}
}

func (s *EchoServiceServer) UnaryEcho(ctx context.Context, req *echopb.UnaryEchoRequest) (*echopb.UnaryEchoResponse, error) {
	return &echopb.UnaryEchoResponse{Msg: req.Msg}, nil
}
