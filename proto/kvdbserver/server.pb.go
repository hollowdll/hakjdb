// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: proto/kvdbserver/server.proto

package kvdbserver

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

type GetServerInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetServerInfoRequest) Reset() {
	*x = GetServerInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserver_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServerInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServerInfoRequest) ProtoMessage() {}

func (x *GetServerInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserver_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServerInfoRequest.ProtoReflect.Descriptor instead.
func (*GetServerInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserver_server_proto_rawDescGZIP(), []int{0}
}

type GetServerInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Information about the server.
	Data *ServerInfo `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetServerInfoResponse) Reset() {
	*x = GetServerInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserver_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServerInfoResponse) ProtoMessage() {}

func (x *GetServerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserver_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServerInfoResponse.ProtoReflect.Descriptor instead.
func (*GetServerInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserver_server_proto_rawDescGZIP(), []int{1}
}

func (x *GetServerInfoResponse) GetData() *ServerInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

// ServerInfo represents information about the server. Will have more fields in future versions.
type ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version of kvdb.
	KvdbVersion string `protobuf:"bytes,1,opt,name=kvdb_version,json=kvdbVersion,proto3" json:"kvdb_version,omitempty"`
	// Version of go used to compile the server.
	GoVersion string `protobuf:"bytes,2,opt,name=go_version,json=goVersion,proto3" json:"go_version,omitempty"`
	// Number of databases.
	DbCount uint32 `protobuf:"varint,3,opt,name=db_count,json=dbCount,proto3" json:"db_count,omitempty"`
	// Total amount of stored data in bytes.
	TotalDataSize uint64 `protobuf:"varint,4,opt,name=total_data_size,json=totalDataSize,proto3" json:"total_data_size,omitempty"`
	// Server operating system.
	Os string `protobuf:"bytes,5,opt,name=os,proto3" json:"os,omitempty"`
	// Architecture which can be 32 or 64 bits.
	Arch string `protobuf:"bytes,6,opt,name=arch,proto3" json:"arch,omitempty"`
	// PID of the server process.
	ProcessId uint32 `protobuf:"varint,7,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	// Server process uptime in seconds.
	UptimeSeconds uint64 `protobuf:"varint,8,opt,name=uptime_seconds,json=uptimeSeconds,proto3" json:"uptime_seconds,omitempty"`
	// Server TCP/IP port.
	TcpPort uint32 `protobuf:"varint,9,opt,name=tcp_port,json=tcpPort,proto3" json:"tcp_port,omitempty"`
}

func (x *ServerInfo) Reset() {
	*x = ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserver_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfo) ProtoMessage() {}

func (x *ServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserver_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfo.ProtoReflect.Descriptor instead.
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserver_server_proto_rawDescGZIP(), []int{2}
}

func (x *ServerInfo) GetKvdbVersion() string {
	if x != nil {
		return x.KvdbVersion
	}
	return ""
}

func (x *ServerInfo) GetGoVersion() string {
	if x != nil {
		return x.GoVersion
	}
	return ""
}

func (x *ServerInfo) GetDbCount() uint32 {
	if x != nil {
		return x.DbCount
	}
	return 0
}

func (x *ServerInfo) GetTotalDataSize() uint64 {
	if x != nil {
		return x.TotalDataSize
	}
	return 0
}

func (x *ServerInfo) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *ServerInfo) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *ServerInfo) GetProcessId() uint32 {
	if x != nil {
		return x.ProcessId
	}
	return 0
}

func (x *ServerInfo) GetUptimeSeconds() uint64 {
	if x != nil {
		return x.UptimeSeconds
	}
	return 0
}

func (x *ServerInfo) GetTcpPort() uint32 {
	if x != nil {
		return x.TcpPort
	}
	return 0
}

// This will have fields in the future to do filtering.
// For example to return only the latest 10 logs.
type GetLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLogsRequest) Reset() {
	*x = GetLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserver_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogsRequest) ProtoMessage() {}

func (x *GetLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserver_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogsRequest.ProtoReflect.Descriptor instead.
func (*GetLogsRequest) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserver_server_proto_rawDescGZIP(), []int{3}
}

type GetLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of returned logs.
	Logs []string `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
	// True if log file is enabled. Otherwise false.
	LogfileEnabled bool `protobuf:"varint,2,opt,name=logfile_enabled,json=logfileEnabled,proto3" json:"logfile_enabled,omitempty"`
}

func (x *GetLogsResponse) Reset() {
	*x = GetLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserver_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogsResponse) ProtoMessage() {}

func (x *GetLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserver_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogsResponse.ProtoReflect.Descriptor instead.
func (*GetLogsResponse) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserver_server_proto_rawDescGZIP(), []int{4}
}

func (x *GetLogsResponse) GetLogs() []string {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *GetLogsResponse) GetLogfileEnabled() bool {
	if x != nil {
		return x.LogfileEnabled
	}
	return false
}

var File_proto_kvdbserver_server_proto protoreflect.FileDescriptor

var file_proto_kvdbserver_server_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69, 0x22, 0x16,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x96,
	0x02, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a,
	0x0c, 0x6b, 0x76, 0x64, 0x62, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6b, 0x76, 0x64, 0x62, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x6f, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x6f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x19, 0x0a, 0x08, 0x64, 0x62, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x64, 0x62, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x6f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x75,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x63, 0x70, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x74, 0x63, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x66, 0x69,
	0x6c, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x32, 0xb9, 0x01, 0x0a, 0x0d, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x2e, 0x6b,
	0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1d, 0x2e, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b,
	0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_kvdbserver_server_proto_rawDescOnce sync.Once
	file_proto_kvdbserver_server_proto_rawDescData = file_proto_kvdbserver_server_proto_rawDesc
)

func file_proto_kvdbserver_server_proto_rawDescGZIP() []byte {
	file_proto_kvdbserver_server_proto_rawDescOnce.Do(func() {
		file_proto_kvdbserver_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_kvdbserver_server_proto_rawDescData)
	})
	return file_proto_kvdbserver_server_proto_rawDescData
}

var file_proto_kvdbserver_server_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_kvdbserver_server_proto_goTypes = []interface{}{
	(*GetServerInfoRequest)(nil),  // 0: kvdbserverapi.GetServerInfoRequest
	(*GetServerInfoResponse)(nil), // 1: kvdbserverapi.GetServerInfoResponse
	(*ServerInfo)(nil),            // 2: kvdbserverapi.ServerInfo
	(*GetLogsRequest)(nil),        // 3: kvdbserverapi.GetLogsRequest
	(*GetLogsResponse)(nil),       // 4: kvdbserverapi.GetLogsResponse
}
var file_proto_kvdbserver_server_proto_depIdxs = []int32{
	2, // 0: kvdbserverapi.GetServerInfoResponse.data:type_name -> kvdbserverapi.ServerInfo
	0, // 1: kvdbserverapi.ServerService.GetServerInfo:input_type -> kvdbserverapi.GetServerInfoRequest
	3, // 2: kvdbserverapi.ServerService.GetLogs:input_type -> kvdbserverapi.GetLogsRequest
	1, // 3: kvdbserverapi.ServerService.GetServerInfo:output_type -> kvdbserverapi.GetServerInfoResponse
	4, // 4: kvdbserverapi.ServerService.GetLogs:output_type -> kvdbserverapi.GetLogsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_kvdbserver_server_proto_init() }
func file_proto_kvdbserver_server_proto_init() {
	if File_proto_kvdbserver_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_kvdbserver_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServerInfoRequest); i {
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
		file_proto_kvdbserver_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServerInfoResponse); i {
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
		file_proto_kvdbserver_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfo); i {
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
		file_proto_kvdbserver_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogsRequest); i {
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
		file_proto_kvdbserver_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogsResponse); i {
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
			RawDescriptor: file_proto_kvdbserver_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_kvdbserver_server_proto_goTypes,
		DependencyIndexes: file_proto_kvdbserver_server_proto_depIdxs,
		MessageInfos:      file_proto_kvdbserver_server_proto_msgTypes,
	}.Build()
	File_proto_kvdbserver_server_proto = out.File
	file_proto_kvdbserver_server_proto_rawDesc = nil
	file_proto_kvdbserver_server_proto_goTypes = nil
	file_proto_kvdbserver_server_proto_depIdxs = nil
}
