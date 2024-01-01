// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: proto/kvdbserver/db.proto

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

// DatabaseClient is the client API for Database service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseClient interface {
	// Creates a new database
	CreateDatabase(ctx context.Context, in *CreateDatabaseRequest, opts ...grpc.CallOption) (*CreateDatabaseResponse, error)
	// Gets and returns the names of all databases
	GetAllDatabases(ctx context.Context, in *GetAllDatabasesRequest, opts ...grpc.CallOption) (*GetAllDatabasesResponse, error)
	// Gets and returns info about a database
	GetDatabaseMetadata(ctx context.Context, in *GetDatabaseMetadataRequest, opts ...grpc.CallOption) (*GetDatabaseMetadataResponse, error)
}

type databaseClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseClient(cc grpc.ClientConnInterface) DatabaseClient {
	return &databaseClient{cc}
}

func (c *databaseClient) CreateDatabase(ctx context.Context, in *CreateDatabaseRequest, opts ...grpc.CallOption) (*CreateDatabaseResponse, error) {
	out := new(CreateDatabaseResponse)
	err := c.cc.Invoke(ctx, "/kvdbserverapi.Database/CreateDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) GetAllDatabases(ctx context.Context, in *GetAllDatabasesRequest, opts ...grpc.CallOption) (*GetAllDatabasesResponse, error) {
	out := new(GetAllDatabasesResponse)
	err := c.cc.Invoke(ctx, "/kvdbserverapi.Database/GetAllDatabases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) GetDatabaseMetadata(ctx context.Context, in *GetDatabaseMetadataRequest, opts ...grpc.CallOption) (*GetDatabaseMetadataResponse, error) {
	out := new(GetDatabaseMetadataResponse)
	err := c.cc.Invoke(ctx, "/kvdbserverapi.Database/GetDatabaseMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServer is the server API for Database service.
// All implementations must embed UnimplementedDatabaseServer
// for forward compatibility
type DatabaseServer interface {
	// Creates a new database
	CreateDatabase(context.Context, *CreateDatabaseRequest) (*CreateDatabaseResponse, error)
	// Gets and returns the names of all databases
	GetAllDatabases(context.Context, *GetAllDatabasesRequest) (*GetAllDatabasesResponse, error)
	// Gets and returns info about a database
	GetDatabaseMetadata(context.Context, *GetDatabaseMetadataRequest) (*GetDatabaseMetadataResponse, error)
	mustEmbedUnimplementedDatabaseServer()
}

// UnimplementedDatabaseServer must be embedded to have forward compatible implementations.
type UnimplementedDatabaseServer struct {
}

func (UnimplementedDatabaseServer) CreateDatabase(context.Context, *CreateDatabaseRequest) (*CreateDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDatabase not implemented")
}
func (UnimplementedDatabaseServer) GetAllDatabases(context.Context, *GetAllDatabasesRequest) (*GetAllDatabasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDatabases not implemented")
}
func (UnimplementedDatabaseServer) GetDatabaseMetadata(context.Context, *GetDatabaseMetadataRequest) (*GetDatabaseMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatabaseMetadata not implemented")
}
func (UnimplementedDatabaseServer) mustEmbedUnimplementedDatabaseServer() {}

// UnsafeDatabaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseServer will
// result in compilation errors.
type UnsafeDatabaseServer interface {
	mustEmbedUnimplementedDatabaseServer()
}

func RegisterDatabaseServer(s grpc.ServiceRegistrar, srv DatabaseServer) {
	s.RegisterService(&Database_ServiceDesc, srv)
}

func _Database_CreateDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).CreateDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvdbserverapi.Database/CreateDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).CreateDatabase(ctx, req.(*CreateDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_GetAllDatabases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllDatabasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).GetAllDatabases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvdbserverapi.Database/GetAllDatabases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).GetAllDatabases(ctx, req.(*GetAllDatabasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_GetDatabaseMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatabaseMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).GetDatabaseMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvdbserverapi.Database/GetDatabaseMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).GetDatabaseMetadata(ctx, req.(*GetDatabaseMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Database_ServiceDesc is the grpc.ServiceDesc for Database service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Database_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kvdbserverapi.Database",
	HandlerType: (*DatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDatabase",
			Handler:    _Database_CreateDatabase_Handler,
		},
		{
			MethodName: "GetAllDatabases",
			Handler:    _Database_GetAllDatabases_Handler,
		},
		{
			MethodName: "GetDatabaseMetadata",
			Handler:    _Database_GetDatabaseMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/kvdbserver/db.proto",
}
