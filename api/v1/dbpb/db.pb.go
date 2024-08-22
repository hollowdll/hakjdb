// HakjDB gRPC API
// API version: 1.0.0
//
// This package contains Protobuf definitions related to databases.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: api/v1/dbpb/db.proto

package dbpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateDBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the database.
	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	// Description of the database.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateDBRequest) Reset() {
	*x = CreateDBRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_dbpb_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDBRequest) ProtoMessage() {}

func (x *CreateDBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_dbpb_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDBRequest.ProtoReflect.Descriptor instead.
func (*CreateDBRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_dbpb_db_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDBRequest) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

func (x *CreateDBRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateDBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the created database.
	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
}

func (x *CreateDBResponse) Reset() {
	*x = CreateDBResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_dbpb_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDBResponse) ProtoMessage() {}

func (x *CreateDBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_dbpb_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDBResponse.ProtoReflect.Descriptor instead.
func (*CreateDBResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_dbpb_db_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDBResponse) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

type GetAllDBsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllDBsRequest) Reset() {
	*x = GetAllDBsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_dbpb_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDBsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDBsRequest) ProtoMessage() {}

func (x *GetAllDBsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_dbpb_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDBsRequest.ProtoReflect.Descriptor instead.
func (*GetAllDBsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_dbpb_db_proto_rawDescGZIP(), []int{2}
}

type GetAllDBsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of returned database names.
	DbNames []string `protobuf:"bytes,1,rep,name=db_names,json=dbNames,proto3" json:"db_names,omitempty"`
}

func (x *GetAllDBsResponse) Reset() {
	*x = GetAllDBsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_dbpb_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDBsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDBsResponse) ProtoMessage() {}

func (x *GetAllDBsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_dbpb_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDBsResponse.ProtoReflect.Descriptor instead.
func (*GetAllDBsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_dbpb_db_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllDBsResponse) GetDbNames() []string {
	if x != nil {
		return x.DbNames
	}
	return nil
}

type GetDBInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the database.
	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
}

func (x *GetDBInfoRequest) Reset() {
	*x = GetDBInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_dbpb_db_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDBInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDBInfoRequest) ProtoMessage() {}

func (x *GetDBInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_dbpb_db_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDBInfoRequest.ProtoReflect.Descriptor instead.
func (*GetDBInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_dbpb_db_proto_rawDescGZIP(), []int{4}
}

func (x *GetDBInfoRequest) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

type GetDBInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Information about the database.
	Data *DBInfo `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetDBInfoResponse) Reset() {
	*x = GetDBInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_dbpb_db_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDBInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDBInfoResponse) ProtoMessage() {}

func (x *GetDBInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_dbpb_db_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDBInfoResponse.ProtoReflect.Descriptor instead.
func (*GetDBInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_dbpb_db_proto_rawDescGZIP(), []int{5}
}

func (x *GetDBInfoResponse) GetData() *DBInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteDBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the database.
	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
}

func (x *DeleteDBRequest) Reset() {
	*x = DeleteDBRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_dbpb_db_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDBRequest) ProtoMessage() {}

func (x *DeleteDBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_dbpb_db_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDBRequest.ProtoReflect.Descriptor instead.
func (*DeleteDBRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_dbpb_db_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDBRequest) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

type DeleteDBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the deleted database.
	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
}

func (x *DeleteDBResponse) Reset() {
	*x = DeleteDBResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_dbpb_db_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDBResponse) ProtoMessage() {}

func (x *DeleteDBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_dbpb_db_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDBResponse.ProtoReflect.Descriptor instead.
func (*DeleteDBResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_dbpb_db_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteDBResponse) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

type ChangeDBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the requested database.
	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	// New name of the database.
	NewName string `protobuf:"bytes,2,opt,name=new_name,json=newName,proto3" json:"new_name,omitempty"`
	// If database name should be changed.
	ChangeName bool `protobuf:"varint,3,opt,name=change_name,json=changeName,proto3" json:"change_name,omitempty"`
	// New description of the database.
	NewDescription string `protobuf:"bytes,4,opt,name=new_description,json=newDescription,proto3" json:"new_description,omitempty"`
	// If database description should be changed.
	ChangeDescription bool `protobuf:"varint,5,opt,name=change_description,json=changeDescription,proto3" json:"change_description,omitempty"`
}

func (x *ChangeDBRequest) Reset() {
	*x = ChangeDBRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_dbpb_db_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeDBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeDBRequest) ProtoMessage() {}

func (x *ChangeDBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_dbpb_db_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeDBRequest.ProtoReflect.Descriptor instead.
func (*ChangeDBRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_dbpb_db_proto_rawDescGZIP(), []int{8}
}

func (x *ChangeDBRequest) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

func (x *ChangeDBRequest) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

func (x *ChangeDBRequest) GetChangeName() bool {
	if x != nil {
		return x.ChangeName
	}
	return false
}

func (x *ChangeDBRequest) GetNewDescription() string {
	if x != nil {
		return x.NewDescription
	}
	return ""
}

func (x *ChangeDBRequest) GetChangeDescription() bool {
	if x != nil {
		return x.ChangeDescription
	}
	return false
}

type ChangeDBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the changed database.
	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
}

func (x *ChangeDBResponse) Reset() {
	*x = ChangeDBResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_dbpb_db_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeDBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeDBResponse) ProtoMessage() {}

func (x *ChangeDBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_dbpb_db_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeDBResponse.ProtoReflect.Descriptor instead.
func (*ChangeDBResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_dbpb_db_proto_rawDescGZIP(), []int{9}
}

func (x *ChangeDBResponse) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

// DBInfo represents information about a database.
type DBInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the database.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// UTC timestamp when the database was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// UTC timestamp when the database was updated.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Size of the stored data in bytes.
	DataSize uint64 `protobuf:"varint,4,opt,name=data_size,json=dataSize,proto3" json:"data_size,omitempty"`
	// Number of keys in the database.
	KeyCount uint32 `protobuf:"varint,5,opt,name=key_count,json=keyCount,proto3" json:"key_count,omitempty"`
	// Description of the database.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *DBInfo) Reset() {
	*x = DBInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_dbpb_db_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBInfo) ProtoMessage() {}

func (x *DBInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_dbpb_db_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBInfo.ProtoReflect.Descriptor instead.
func (*DBInfo) Descriptor() ([]byte, []int) {
	return file_api_v1_dbpb_db_proto_rawDescGZIP(), []int{10}
}

func (x *DBInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DBInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DBInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *DBInfo) GetDataSize() uint64 {
	if x != nil {
		return x.DataSize
	}
	return 0
}

func (x *DBInfo) GetKeyCount() uint32 {
	if x != nil {
		return x.KeyCount
	}
	return 0
}

func (x *DBInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_api_v1_dbpb_db_proto protoreflect.FileDescriptor

var file_api_v1_dbpb_db_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x62, 0x70, 0x62, 0x2f, 0x64, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x64,
	0x62, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x42,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x42, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x42, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x42, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x62, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x62, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x22, 0x2b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x42, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x3c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x42, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x62, 0x70,
	0x62, 0x2e, 0x44, 0x42, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2a,
	0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xbe, 0x01, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64,
	0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x65, 0x77, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xee, 0x01, 0x0a, 0x06, 0x44, 0x42, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x88, 0x03, 0x0a, 0x09, 0x44, 0x42, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x42,
	0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x62, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x62, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x42, 0x73, 0x12, 0x1d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x62, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x44, 0x42, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x62, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x44, 0x42, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x44, 0x42, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x64, 0x62, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x42, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x64, 0x62, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x42, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x08, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x42, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x64, 0x62, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x42, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x64,
	0x62, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x44, 0x42, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x62, 0x70, 0x62,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x62, 0x70, 0x62, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x62, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_dbpb_db_proto_rawDescOnce sync.Once
	file_api_v1_dbpb_db_proto_rawDescData = file_api_v1_dbpb_db_proto_rawDesc
)

func file_api_v1_dbpb_db_proto_rawDescGZIP() []byte {
	file_api_v1_dbpb_db_proto_rawDescOnce.Do(func() {
		file_api_v1_dbpb_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_dbpb_db_proto_rawDescData)
	})
	return file_api_v1_dbpb_db_proto_rawDescData
}

var file_api_v1_dbpb_db_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_v1_dbpb_db_proto_goTypes = []interface{}{
	(*CreateDBRequest)(nil),       // 0: api.v1.dbpb.CreateDBRequest
	(*CreateDBResponse)(nil),      // 1: api.v1.dbpb.CreateDBResponse
	(*GetAllDBsRequest)(nil),      // 2: api.v1.dbpb.GetAllDBsRequest
	(*GetAllDBsResponse)(nil),     // 3: api.v1.dbpb.GetAllDBsResponse
	(*GetDBInfoRequest)(nil),      // 4: api.v1.dbpb.GetDBInfoRequest
	(*GetDBInfoResponse)(nil),     // 5: api.v1.dbpb.GetDBInfoResponse
	(*DeleteDBRequest)(nil),       // 6: api.v1.dbpb.DeleteDBRequest
	(*DeleteDBResponse)(nil),      // 7: api.v1.dbpb.DeleteDBResponse
	(*ChangeDBRequest)(nil),       // 8: api.v1.dbpb.ChangeDBRequest
	(*ChangeDBResponse)(nil),      // 9: api.v1.dbpb.ChangeDBResponse
	(*DBInfo)(nil),                // 10: api.v1.dbpb.DBInfo
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_api_v1_dbpb_db_proto_depIdxs = []int32{
	10, // 0: api.v1.dbpb.GetDBInfoResponse.data:type_name -> api.v1.dbpb.DBInfo
	11, // 1: api.v1.dbpb.DBInfo.created_at:type_name -> google.protobuf.Timestamp
	11, // 2: api.v1.dbpb.DBInfo.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 3: api.v1.dbpb.DBService.CreateDB:input_type -> api.v1.dbpb.CreateDBRequest
	2,  // 4: api.v1.dbpb.DBService.GetAllDBs:input_type -> api.v1.dbpb.GetAllDBsRequest
	4,  // 5: api.v1.dbpb.DBService.GetDBInfo:input_type -> api.v1.dbpb.GetDBInfoRequest
	6,  // 6: api.v1.dbpb.DBService.DeleteDB:input_type -> api.v1.dbpb.DeleteDBRequest
	8,  // 7: api.v1.dbpb.DBService.ChangeDB:input_type -> api.v1.dbpb.ChangeDBRequest
	1,  // 8: api.v1.dbpb.DBService.CreateDB:output_type -> api.v1.dbpb.CreateDBResponse
	3,  // 9: api.v1.dbpb.DBService.GetAllDBs:output_type -> api.v1.dbpb.GetAllDBsResponse
	5,  // 10: api.v1.dbpb.DBService.GetDBInfo:output_type -> api.v1.dbpb.GetDBInfoResponse
	7,  // 11: api.v1.dbpb.DBService.DeleteDB:output_type -> api.v1.dbpb.DeleteDBResponse
	9,  // 12: api.v1.dbpb.DBService.ChangeDB:output_type -> api.v1.dbpb.ChangeDBResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_v1_dbpb_db_proto_init() }
func file_api_v1_dbpb_db_proto_init() {
	if File_api_v1_dbpb_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_dbpb_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDBRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_dbpb_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDBResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_dbpb_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDBsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_dbpb_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDBsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_dbpb_db_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDBInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_dbpb_db_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDBInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_dbpb_db_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDBRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_dbpb_db_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDBResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_dbpb_db_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeDBRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_dbpb_db_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeDBResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_dbpb_db_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_dbpb_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_dbpb_db_proto_goTypes,
		DependencyIndexes: file_api_v1_dbpb_db_proto_depIdxs,
		MessageInfos:      file_api_v1_dbpb_db_proto_msgTypes,
	}.Build()
	File_api_v1_dbpb_db_proto = out.File
	file_api_v1_dbpb_db_proto_rawDesc = nil
	file_api_v1_dbpb_db_proto_goTypes = nil
	file_api_v1_dbpb_db_proto_depIdxs = nil
}