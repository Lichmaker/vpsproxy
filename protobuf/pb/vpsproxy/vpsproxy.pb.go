// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.2
// source: vpsproxy.proto

package vpsproxy

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateHttpProxyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port uint32 `protobuf:"varint,1,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *UpdateHttpProxyReq) Reset() {
	*x = UpdateHttpProxyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpsproxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHttpProxyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHttpProxyReq) ProtoMessage() {}

func (x *UpdateHttpProxyReq) ProtoReflect() protoreflect.Message {
	mi := &file_vpsproxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHttpProxyReq.ProtoReflect.Descriptor instead.
func (*UpdateHttpProxyReq) Descriptor() ([]byte, []int) {
	return file_vpsproxy_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateHttpProxyReq) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type GetHttpProxyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port uint32 `protobuf:"varint,1,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *GetHttpProxyResponse) Reset() {
	*x = GetHttpProxyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpsproxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHttpProxyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHttpProxyResponse) ProtoMessage() {}

func (x *GetHttpProxyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vpsproxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHttpProxyResponse.ProtoReflect.Descriptor instead.
func (*GetHttpProxyResponse) Descriptor() ([]byte, []int) {
	return file_vpsproxy_proto_rawDescGZIP(), []int{1}
}

func (x *GetHttpProxyResponse) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_vpsproxy_proto protoreflect.FileDescriptor

var file_vpsproxy_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x76, 0x70, 0x73, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x76, 0x70, 0x73, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x50, 0x6f, 0x72,
	0x74, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x32, 0xde, 0x01,
	0x0a, 0x0f, 0x56, 0x50, 0x53, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x46, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x76, 0x70, 0x73, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x1c, 0x2e, 0x76,
	0x70, 0x73, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x74,
	0x74, 0x70, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x3a, 0x0a, 0x08, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0b,
	0x5a, 0x09, 0x2f, 0x76, 0x70, 0x73, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_vpsproxy_proto_rawDescOnce sync.Once
	file_vpsproxy_proto_rawDescData = file_vpsproxy_proto_rawDesc
)

func file_vpsproxy_proto_rawDescGZIP() []byte {
	file_vpsproxy_proto_rawDescOnce.Do(func() {
		file_vpsproxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_vpsproxy_proto_rawDescData)
	})
	return file_vpsproxy_proto_rawDescData
}

var file_vpsproxy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vpsproxy_proto_goTypes = []interface{}{
	(*UpdateHttpProxyReq)(nil),   // 0: vpsproxy.UpdateHttpProxyReq
	(*GetHttpProxyResponse)(nil), // 1: vpsproxy.GetHttpProxyResponse
	(*emptypb.Empty)(nil),        // 2: google.protobuf.Empty
}
var file_vpsproxy_proto_depIdxs = []int32{
	2, // 0: vpsproxy.VPSProxyService.GetHttpProxy:input_type -> google.protobuf.Empty
	0, // 1: vpsproxy.VPSProxyService.UpdateHttpProxy:input_type -> vpsproxy.UpdateHttpProxyReq
	2, // 2: vpsproxy.VPSProxyService.Shutdown:input_type -> google.protobuf.Empty
	1, // 3: vpsproxy.VPSProxyService.GetHttpProxy:output_type -> vpsproxy.GetHttpProxyResponse
	2, // 4: vpsproxy.VPSProxyService.UpdateHttpProxy:output_type -> google.protobuf.Empty
	2, // 5: vpsproxy.VPSProxyService.Shutdown:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vpsproxy_proto_init() }
func file_vpsproxy_proto_init() {
	if File_vpsproxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vpsproxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHttpProxyReq); i {
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
		file_vpsproxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHttpProxyResponse); i {
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
			RawDescriptor: file_vpsproxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vpsproxy_proto_goTypes,
		DependencyIndexes: file_vpsproxy_proto_depIdxs,
		MessageInfos:      file_vpsproxy_proto_msgTypes,
	}.Build()
	File_vpsproxy_proto = out.File
	file_vpsproxy_proto_rawDesc = nil
	file_vpsproxy_proto_goTypes = nil
	file_vpsproxy_proto_depIdxs = nil
}