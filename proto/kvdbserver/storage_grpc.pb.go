// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: proto/kvdbserver/storage.proto

package kvdbserver

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

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageServiceClient interface {
	// SetString sets a string value using a key.
	SetString(ctx context.Context, in *SetStringRequest, opts ...grpc.CallOption) (*SetStringResponse, error)
	// GetString gets a string value using a key.
	GetString(ctx context.Context, in *GetStringRequest, opts ...grpc.CallOption) (*GetStringResponse, error)
	// DeleteKey deletes a key and its value.
	DeleteKey(ctx context.Context, in *DeleteKeyRequest, opts ...grpc.CallOption) (*DeleteKeyResponse, error)
	// DeleteAllKeys deletes all the keys of a database.
	DeleteAllKeys(ctx context.Context, in *DeleteAllKeysRequest, opts ...grpc.CallOption) (*DeleteAllKeysResponse, error)
	// GetKeys returns all the keys of a database.
	GetKeys(ctx context.Context, in *GetKeysRequest, opts ...grpc.CallOption) (*GetKeysResponse, error)
}

type storageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageServiceClient(cc grpc.ClientConnInterface) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) SetString(ctx context.Context, in *SetStringRequest, opts ...grpc.CallOption) (*SetStringResponse, error) {
	out := new(SetStringResponse)
	err := c.cc.Invoke(ctx, "/kvdbserverapi.StorageService/SetString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetString(ctx context.Context, in *GetStringRequest, opts ...grpc.CallOption) (*GetStringResponse, error) {
	out := new(GetStringResponse)
	err := c.cc.Invoke(ctx, "/kvdbserverapi.StorageService/GetString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) DeleteKey(ctx context.Context, in *DeleteKeyRequest, opts ...grpc.CallOption) (*DeleteKeyResponse, error) {
	out := new(DeleteKeyResponse)
	err := c.cc.Invoke(ctx, "/kvdbserverapi.StorageService/DeleteKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) DeleteAllKeys(ctx context.Context, in *DeleteAllKeysRequest, opts ...grpc.CallOption) (*DeleteAllKeysResponse, error) {
	out := new(DeleteAllKeysResponse)
	err := c.cc.Invoke(ctx, "/kvdbserverapi.StorageService/DeleteAllKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetKeys(ctx context.Context, in *GetKeysRequest, opts ...grpc.CallOption) (*GetKeysResponse, error) {
	out := new(GetKeysResponse)
	err := c.cc.Invoke(ctx, "/kvdbserverapi.StorageService/GetKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServiceServer is the server API for StorageService service.
// All implementations must embed UnimplementedStorageServiceServer
// for forward compatibility
type StorageServiceServer interface {
	// SetString sets a string value using a key.
	SetString(context.Context, *SetStringRequest) (*SetStringResponse, error)
	// GetString gets a string value using a key.
	GetString(context.Context, *GetStringRequest) (*GetStringResponse, error)
	// DeleteKey deletes a key and its value.
	DeleteKey(context.Context, *DeleteKeyRequest) (*DeleteKeyResponse, error)
	// DeleteAllKeys deletes all the keys of a database.
	DeleteAllKeys(context.Context, *DeleteAllKeysRequest) (*DeleteAllKeysResponse, error)
	// GetKeys returns all the keys of a database.
	GetKeys(context.Context, *GetKeysRequest) (*GetKeysResponse, error)
	mustEmbedUnimplementedStorageServiceServer()
}

// UnimplementedStorageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServiceServer struct {
}

func (UnimplementedStorageServiceServer) SetString(context.Context, *SetStringRequest) (*SetStringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetString not implemented")
}
func (UnimplementedStorageServiceServer) GetString(context.Context, *GetStringRequest) (*GetStringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetString not implemented")
}
func (UnimplementedStorageServiceServer) DeleteKey(context.Context, *DeleteKeyRequest) (*DeleteKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKey not implemented")
}
func (UnimplementedStorageServiceServer) DeleteAllKeys(context.Context, *DeleteAllKeysRequest) (*DeleteAllKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllKeys not implemented")
}
func (UnimplementedStorageServiceServer) GetKeys(context.Context, *GetKeysRequest) (*GetKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeys not implemented")
}
func (UnimplementedStorageServiceServer) mustEmbedUnimplementedStorageServiceServer() {}

// UnsafeStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServiceServer will
// result in compilation errors.
type UnsafeStorageServiceServer interface {
	mustEmbedUnimplementedStorageServiceServer()
}

func RegisterStorageServiceServer(s grpc.ServiceRegistrar, srv StorageServiceServer) {
	s.RegisterService(&StorageService_ServiceDesc, srv)
}

func _StorageService_SetString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).SetString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvdbserverapi.StorageService/SetString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).SetString(ctx, req.(*SetStringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvdbserverapi.StorageService/GetString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetString(ctx, req.(*GetStringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_DeleteKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).DeleteKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvdbserverapi.StorageService/DeleteKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).DeleteKey(ctx, req.(*DeleteKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_DeleteAllKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).DeleteAllKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvdbserverapi.StorageService/DeleteAllKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).DeleteAllKeys(ctx, req.(*DeleteAllKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvdbserverapi.StorageService/GetKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetKeys(ctx, req.(*GetKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageService_ServiceDesc is the grpc.ServiceDesc for StorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kvdbserverapi.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetString",
			Handler:    _StorageService_SetString_Handler,
		},
		{
			MethodName: "GetString",
			Handler:    _StorageService_GetString_Handler,
		},
		{
			MethodName: "DeleteKey",
			Handler:    _StorageService_DeleteKey_Handler,
		},
		{
			MethodName: "DeleteAllKeys",
			Handler:    _StorageService_DeleteAllKeys_Handler,
		},
		{
			MethodName: "GetKeys",
			Handler:    _StorageService_GetKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/kvdbserver/storage.proto",
}
