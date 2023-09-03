// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: v1/echo.proto

package echov1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type EchoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EchoMessage) Reset() {
	*x = EchoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoMessage) ProtoMessage() {}

func (x *EchoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_v1_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoMessage.ProtoReflect.Descriptor instead.
func (*EchoMessage) Descriptor() ([]byte, []int) {
	return file_v1_echo_proto_rawDescGZIP(), []int{0}
}

func (x *EchoMessage) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_v1_echo_proto protoreflect.FileDescriptor

var file_v1_echo_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0b,
	0x45, 0x63, 0x68, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x32, 0x62, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x53, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x63, 0x68,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x63, 0x68, 0x6f, 0x42, 0xb5, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x45, 0x63, 0x68,
	0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x65, 0x76, 0x65, 0x6e, 0x77, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x65, 0x63, 0x68, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x45, 0x58, 0xaa, 0x02, 0x0d,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x45, 0x63, 0x68, 0x6f, 0x5f, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x45, 0x63, 0x68, 0x6f, 0x5f, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x45, 0x63, 0x68, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_echo_proto_rawDescOnce sync.Once
	file_v1_echo_proto_rawDescData = file_v1_echo_proto_rawDesc
)

func file_v1_echo_proto_rawDescGZIP() []byte {
	file_v1_echo_proto_rawDescOnce.Do(func() {
		file_v1_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_echo_proto_rawDescData)
	})
	return file_v1_echo_proto_rawDescData
}

var file_v1_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v1_echo_proto_goTypes = []interface{}{
	(*EchoMessage)(nil), // 0: proto.echo.v1.EchoMessage
}
var file_v1_echo_proto_depIdxs = []int32{
	0, // 0: proto.echo.v1.EchoService.Echo:input_type -> proto.echo.v1.EchoMessage
	0, // 1: proto.echo.v1.EchoService.Echo:output_type -> proto.echo.v1.EchoMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_echo_proto_init() }
func file_v1_echo_proto_init() {
	if File_v1_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoMessage); i {
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
			RawDescriptor: file_v1_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_echo_proto_goTypes,
		DependencyIndexes: file_v1_echo_proto_depIdxs,
		MessageInfos:      file_v1_echo_proto_msgTypes,
	}.Build()
	File_v1_echo_proto = out.File
	file_v1_echo_proto_rawDesc = nil
	file_v1_echo_proto_goTypes = nil
	file_v1_echo_proto_depIdxs = nil
}
