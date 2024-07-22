// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: api/v0/storagepb/general.proto

package storagepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GeneralKeyServiceClient is the client API for GeneralKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeneralKeyServiceClient interface {
	// GetAllKeys returns all the keys.
	// Uses the database sent in gRPC metadata or the default database if omitted.
	GetAllKeys(ctx context.Context, in *GetAllKeysRequest, opts ...grpc.CallOption) (*GetAllKeysResponse, error)
	// GetKeyType returns the data type of the value a key is holding.
	// Uses the database sent in gRPC metadata or the default database if omitted.
	GetKeyType(ctx context.Context, in *GetKeyTypeRequest, opts ...grpc.CallOption) (*GetKeyTypeResponse, error)
	// DeleteKeys deletes the specified keys and the values they are holding.
	// Ignores keys that do not exist.
	// Uses the database sent in gRPC metadata or the default database if omitted.
	DeleteKeys(ctx context.Context, in *DeleteKeysRequest, opts ...grpc.CallOption) (*DeleteKeysResponse, error)
	// DeleteAllKeys deletes all the keys.
	// Uses the database sent in gRPC metadata or the default database if omitted.
	DeleteAllKeys(ctx context.Context, in *DeleteAllKeysRequest, opts ...grpc.CallOption) (*DeleteAllKeysResponse, error)
}

type generalKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGeneralKeyServiceClient(cc grpc.ClientConnInterface) GeneralKeyServiceClient {
	return &generalKeyServiceClient{cc}
}

func (c *generalKeyServiceClient) GetAllKeys(ctx context.Context, in *GetAllKeysRequest, opts ...grpc.CallOption) (*GetAllKeysResponse, error) {
	out := new(GetAllKeysResponse)
	err := c.cc.Invoke(ctx, "/api.v0.storagepb.GeneralKeyService/GetAllKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generalKeyServiceClient) GetKeyType(ctx context.Context, in *GetKeyTypeRequest, opts ...grpc.CallOption) (*GetKeyTypeResponse, error) {
	out := new(GetKeyTypeResponse)
	err := c.cc.Invoke(ctx, "/api.v0.storagepb.GeneralKeyService/GetKeyType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generalKeyServiceClient) DeleteKeys(ctx context.Context, in *DeleteKeysRequest, opts ...grpc.CallOption) (*DeleteKeysResponse, error) {
	out := new(DeleteKeysResponse)
	err := c.cc.Invoke(ctx, "/api.v0.storagepb.GeneralKeyService/DeleteKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generalKeyServiceClient) DeleteAllKeys(ctx context.Context, in *DeleteAllKeysRequest, opts ...grpc.CallOption) (*DeleteAllKeysResponse, error) {
	out := new(DeleteAllKeysResponse)
	err := c.cc.Invoke(ctx, "/api.v0.storagepb.GeneralKeyService/DeleteAllKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeneralKeyServiceServer is the server API for GeneralKeyService service.
// All implementations must embed UnimplementedGeneralKeyServiceServer
// for forward compatibility
type GeneralKeyServiceServer interface {
	// GetAllKeys returns all the keys.
	// Uses the database sent in gRPC metadata or the default database if omitted.
	GetAllKeys(context.Context, *GetAllKeysRequest) (*GetAllKeysResponse, error)
	// GetKeyType returns the data type of the value a key is holding.
	// Uses the database sent in gRPC metadata or the default database if omitted.
	GetKeyType(context.Context, *GetKeyTypeRequest) (*GetKeyTypeResponse, error)
	// DeleteKeys deletes the specified keys and the values they are holding.
	// Ignores keys that do not exist.
	// Uses the database sent in gRPC metadata or the default database if omitted.
	DeleteKeys(context.Context, *DeleteKeysRequest) (*DeleteKeysResponse, error)
	// DeleteAllKeys deletes all the keys.
	// Uses the database sent in gRPC metadata or the default database if omitted.
	DeleteAllKeys(context.Context, *DeleteAllKeysRequest) (*DeleteAllKeysResponse, error)
	mustEmbedUnimplementedGeneralKeyServiceServer()
}

// UnimplementedGeneralKeyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGeneralKeyServiceServer struct {
}

func (UnimplementedGeneralKeyServiceServer) GetAllKeys(context.Context, *GetAllKeysRequest) (*GetAllKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllKeys not implemented")
}
func (UnimplementedGeneralKeyServiceServer) GetKeyType(context.Context, *GetKeyTypeRequest) (*GetKeyTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeyType not implemented")
}
func (UnimplementedGeneralKeyServiceServer) DeleteKeys(context.Context, *DeleteKeysRequest) (*DeleteKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKeys not implemented")
}
func (UnimplementedGeneralKeyServiceServer) DeleteAllKeys(context.Context, *DeleteAllKeysRequest) (*DeleteAllKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllKeys not implemented")
}
func (UnimplementedGeneralKeyServiceServer) mustEmbedUnimplementedGeneralKeyServiceServer() {}

// UnsafeGeneralKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeneralKeyServiceServer will
// result in compilation errors.
type UnsafeGeneralKeyServiceServer interface {
	mustEmbedUnimplementedGeneralKeyServiceServer()
}

func RegisterGeneralKeyServiceServer(s grpc.ServiceRegistrar, srv GeneralKeyServiceServer) {
	s.RegisterService(&GeneralKeyService_ServiceDesc, srv)
}

func _GeneralKeyService_GetAllKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneralKeyServiceServer).GetAllKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v0.storagepb.GeneralKeyService/GetAllKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneralKeyServiceServer).GetAllKeys(ctx, req.(*GetAllKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeneralKeyService_GetKeyType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeyTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneralKeyServiceServer).GetKeyType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v0.storagepb.GeneralKeyService/GetKeyType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneralKeyServiceServer).GetKeyType(ctx, req.(*GetKeyTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeneralKeyService_DeleteKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneralKeyServiceServer).DeleteKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v0.storagepb.GeneralKeyService/DeleteKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneralKeyServiceServer).DeleteKeys(ctx, req.(*DeleteKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeneralKeyService_DeleteAllKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneralKeyServiceServer).DeleteAllKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v0.storagepb.GeneralKeyService/DeleteAllKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneralKeyServiceServer).DeleteAllKeys(ctx, req.(*DeleteAllKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GeneralKeyService_ServiceDesc is the grpc.ServiceDesc for GeneralKeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeneralKeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v0.storagepb.GeneralKeyService",
	HandlerType: (*GeneralKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllKeys",
			Handler:    _GeneralKeyService_GetAllKeys_Handler,
		},
		{
			MethodName: "GetKeyType",
			Handler:    _GeneralKeyService_GetKeyType_Handler,
		},
		{
			MethodName: "DeleteKeys",
			Handler:    _GeneralKeyService_DeleteKeys_Handler,
		},
		{
			MethodName: "DeleteAllKeys",
			Handler:    _GeneralKeyService_DeleteAllKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v0/storagepb/general.proto",
}
