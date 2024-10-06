// HakjDB gRPC API
// API version: 1.2.0
//
// This package contains Protobuf definitions related to key-value storage.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: api/v1/kvpb/general_kv.proto

package kvpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetAllKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllKeysRequest) Reset() {
	*x = GetAllKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_kvpb_general_kv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllKeysRequest) ProtoMessage() {}

func (x *GetAllKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_kvpb_general_kv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllKeysRequest.ProtoReflect.Descriptor instead.
func (*GetAllKeysRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_kvpb_general_kv_proto_rawDescGZIP(), []int{0}
}

type GetAllKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of returned keys.
	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *GetAllKeysResponse) Reset() {
	*x = GetAllKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_kvpb_general_kv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllKeysResponse) ProtoMessage() {}

func (x *GetAllKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_kvpb_general_kv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllKeysResponse.ProtoReflect.Descriptor instead.
func (*GetAllKeysResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_kvpb_general_kv_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllKeysResponse) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type GetKeyTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The key whose data type should be returned.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetKeyTypeRequest) Reset() {
	*x = GetKeyTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_kvpb_general_kv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyTypeRequest) ProtoMessage() {}

func (x *GetKeyTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_kvpb_general_kv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyTypeRequest.ProtoReflect.Descriptor instead.
func (*GetKeyTypeRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_kvpb_general_kv_proto_rawDescGZIP(), []int{2}
}

func (x *GetKeyTypeRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetKeyTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The data type of the key.
	KeyType string `protobuf:"bytes,1,opt,name=key_type,json=keyType,proto3" json:"key_type,omitempty"`
	// True if the key exists. False if it doesn't exist.
	Ok bool `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *GetKeyTypeResponse) Reset() {
	*x = GetKeyTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_kvpb_general_kv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyTypeResponse) ProtoMessage() {}

func (x *GetKeyTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_kvpb_general_kv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyTypeResponse.ProtoReflect.Descriptor instead.
func (*GetKeyTypeResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_kvpb_general_kv_proto_rawDescGZIP(), []int{3}
}

func (x *GetKeyTypeResponse) GetKeyType() string {
	if x != nil {
		return x.KeyType
	}
	return ""
}

func (x *GetKeyTypeResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type DeleteKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The keys to delete.
	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *DeleteKeysRequest) Reset() {
	*x = DeleteKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_kvpb_general_kv_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKeysRequest) ProtoMessage() {}

func (x *DeleteKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_kvpb_general_kv_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKeysRequest.ProtoReflect.Descriptor instead.
func (*DeleteKeysRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_kvpb_general_kv_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteKeysRequest) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type DeleteKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of keys that were deleted.
	KeysDeletedCount uint32 `protobuf:"varint,1,opt,name=keys_deleted_count,json=keysDeletedCount,proto3" json:"keys_deleted_count,omitempty"`
}

func (x *DeleteKeysResponse) Reset() {
	*x = DeleteKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_kvpb_general_kv_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKeysResponse) ProtoMessage() {}

func (x *DeleteKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_kvpb_general_kv_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKeysResponse.ProtoReflect.Descriptor instead.
func (*DeleteKeysResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_kvpb_general_kv_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteKeysResponse) GetKeysDeletedCount() uint32 {
	if x != nil {
		return x.KeysDeletedCount
	}
	return 0
}

type DeleteAllKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAllKeysRequest) Reset() {
	*x = DeleteAllKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_kvpb_general_kv_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAllKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllKeysRequest) ProtoMessage() {}

func (x *DeleteAllKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_kvpb_general_kv_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllKeysRequest.ProtoReflect.Descriptor instead.
func (*DeleteAllKeysRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_kvpb_general_kv_proto_rawDescGZIP(), []int{6}
}

type DeleteAllKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAllKeysResponse) Reset() {
	*x = DeleteAllKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_kvpb_general_kv_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAllKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllKeysResponse) ProtoMessage() {}

func (x *DeleteAllKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_kvpb_general_kv_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllKeysResponse.ProtoReflect.Descriptor instead.
func (*DeleteAllKeysResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_kvpb_general_kv_proto_rawDescGZIP(), []int{7}
}

var File_api_v1_kvpb_general_kv_proto protoreflect.FileDescriptor

var file_api_v1_kvpb_general_kv_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x76, 0x70, 0x62, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x76, 0x70, 0x62, 0x22, 0x13, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x28, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x25, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0x3f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02,
	0x6f, 0x6b, 0x22, 0x27, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x42, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6b, 0x65, 0x79, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6b,
	0x65, 0x79, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x4b, 0x65, 0x79, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x6c, 0x6c, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xdf, 0x02, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x4b, 0x56, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4b,
	0x65, 0x79, 0x73, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x76, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x76, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x76,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x76,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6b,
	0x76, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6b,
	0x76, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x6c, 0x6c, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x6b, 0x76, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c,
	0x6c, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6b, 0x76, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x6c, 0x6c, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x76, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_kvpb_general_kv_proto_rawDescOnce sync.Once
	file_api_v1_kvpb_general_kv_proto_rawDescData = file_api_v1_kvpb_general_kv_proto_rawDesc
)

func file_api_v1_kvpb_general_kv_proto_rawDescGZIP() []byte {
	file_api_v1_kvpb_general_kv_proto_rawDescOnce.Do(func() {
		file_api_v1_kvpb_general_kv_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_kvpb_general_kv_proto_rawDescData)
	})
	return file_api_v1_kvpb_general_kv_proto_rawDescData
}

var file_api_v1_kvpb_general_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_v1_kvpb_general_kv_proto_goTypes = []interface{}{
	(*GetAllKeysRequest)(nil),     // 0: api.v1.kvpb.GetAllKeysRequest
	(*GetAllKeysResponse)(nil),    // 1: api.v1.kvpb.GetAllKeysResponse
	(*GetKeyTypeRequest)(nil),     // 2: api.v1.kvpb.GetKeyTypeRequest
	(*GetKeyTypeResponse)(nil),    // 3: api.v1.kvpb.GetKeyTypeResponse
	(*DeleteKeysRequest)(nil),     // 4: api.v1.kvpb.DeleteKeysRequest
	(*DeleteKeysResponse)(nil),    // 5: api.v1.kvpb.DeleteKeysResponse
	(*DeleteAllKeysRequest)(nil),  // 6: api.v1.kvpb.DeleteAllKeysRequest
	(*DeleteAllKeysResponse)(nil), // 7: api.v1.kvpb.DeleteAllKeysResponse
}
var file_api_v1_kvpb_general_kv_proto_depIdxs = []int32{
	0, // 0: api.v1.kvpb.GeneralKVService.GetAllKeys:input_type -> api.v1.kvpb.GetAllKeysRequest
	2, // 1: api.v1.kvpb.GeneralKVService.GetKeyType:input_type -> api.v1.kvpb.GetKeyTypeRequest
	4, // 2: api.v1.kvpb.GeneralKVService.DeleteKeys:input_type -> api.v1.kvpb.DeleteKeysRequest
	6, // 3: api.v1.kvpb.GeneralKVService.DeleteAllKeys:input_type -> api.v1.kvpb.DeleteAllKeysRequest
	1, // 4: api.v1.kvpb.GeneralKVService.GetAllKeys:output_type -> api.v1.kvpb.GetAllKeysResponse
	3, // 5: api.v1.kvpb.GeneralKVService.GetKeyType:output_type -> api.v1.kvpb.GetKeyTypeResponse
	5, // 6: api.v1.kvpb.GeneralKVService.DeleteKeys:output_type -> api.v1.kvpb.DeleteKeysResponse
	7, // 7: api.v1.kvpb.GeneralKVService.DeleteAllKeys:output_type -> api.v1.kvpb.DeleteAllKeysResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_kvpb_general_kv_proto_init() }
func file_api_v1_kvpb_general_kv_proto_init() {
	if File_api_v1_kvpb_general_kv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_kvpb_general_kv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllKeysRequest); i {
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
		file_api_v1_kvpb_general_kv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllKeysResponse); i {
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
		file_api_v1_kvpb_general_kv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyTypeRequest); i {
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
		file_api_v1_kvpb_general_kv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyTypeResponse); i {
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
		file_api_v1_kvpb_general_kv_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKeysRequest); i {
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
		file_api_v1_kvpb_general_kv_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKeysResponse); i {
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
		file_api_v1_kvpb_general_kv_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAllKeysRequest); i {
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
		file_api_v1_kvpb_general_kv_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAllKeysResponse); i {
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
			RawDescriptor: file_api_v1_kvpb_general_kv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_kvpb_general_kv_proto_goTypes,
		DependencyIndexes: file_api_v1_kvpb_general_kv_proto_depIdxs,
		MessageInfos:      file_api_v1_kvpb_general_kv_proto_msgTypes,
	}.Build()
	File_api_v1_kvpb_general_kv_proto = out.File
	file_api_v1_kvpb_general_kv_proto_rawDesc = nil
	file_api_v1_kvpb_general_kv_proto_goTypes = nil
	file_api_v1_kvpb_general_kv_proto_depIdxs = nil
}
