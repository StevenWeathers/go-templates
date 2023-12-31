// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: authors/v1/authors.proto

package authorsv1

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Bio  string `protobuf:"bytes,3,opt,name=Bio,proto3" json:"Bio,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_v1_authors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_authors_v1_authors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_authors_v1_authors_proto_rawDescGZIP(), []int{0}
}

func (x *Author) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

type Authors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author []*Author `protobuf:"bytes,1,rep,name=author,proto3" json:"author,omitempty"`
}

func (x *Authors) Reset() {
	*x = Authors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_v1_authors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authors) ProtoMessage() {}

func (x *Authors) ProtoReflect() protoreflect.Message {
	mi := &file_authors_v1_authors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authors.ProtoReflect.Descriptor instead.
func (*Authors) Descriptor() ([]byte, []int) {
	return file_authors_v1_authors_proto_rawDescGZIP(), []int{1}
}

func (x *Authors) GetAuthor() []*Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type ReadAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ReadAuthorRequest) Reset() {
	*x = ReadAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_v1_authors_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAuthorRequest) ProtoMessage() {}

func (x *ReadAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authors_v1_authors_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAuthorRequest.ProtoReflect.Descriptor instead.
func (*ReadAuthorRequest) Descriptor() ([]byte, []int) {
	return file_authors_v1_authors_proto_rawDescGZIP(), []int{2}
}

func (x *ReadAuthorRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type VoidNoParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VoidNoParam) Reset() {
	*x = VoidNoParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_v1_authors_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoidNoParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoidNoParam) ProtoMessage() {}

func (x *VoidNoParam) ProtoReflect() protoreflect.Message {
	mi := &file_authors_v1_authors_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoidNoParam.ProtoReflect.Descriptor instead.
func (*VoidNoParam) Descriptor() ([]byte, []int) {
	return file_authors_v1_authors_proto_rawDescGZIP(), []int{3}
}

var File_authors_v1_authors_proto protoreflect.FileDescriptor

var file_authors_v1_authors_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x42, 0x69, 0x6f, 0x22, 0x35, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12,
	0x2a, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x23, 0x0a, 0x11, 0x52,
	0x65, 0x61, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44,
	0x22, 0x0d, 0x0a, 0x0b, 0x76, 0x6f, 0x69, 0x64, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x32,
	0x8d, 0x02, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x12, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x76,
	0x6f, 0x69, 0x64, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x22,
	0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x12, 0x4e, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x1a, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x16, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x12, 0x59, 0x0a, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x42,
	0xb8, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x74, 0x65, 0x76, 0x65, 0x6e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x6f,
	0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_authors_v1_authors_proto_rawDescOnce sync.Once
	file_authors_v1_authors_proto_rawDescData = file_authors_v1_authors_proto_rawDesc
)

func file_authors_v1_authors_proto_rawDescGZIP() []byte {
	file_authors_v1_authors_proto_rawDescOnce.Do(func() {
		file_authors_v1_authors_proto_rawDescData = protoimpl.X.CompressGZIP(file_authors_v1_authors_proto_rawDescData)
	})
	return file_authors_v1_authors_proto_rawDescData
}

var file_authors_v1_authors_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_authors_v1_authors_proto_goTypes = []interface{}{
	(*Author)(nil),            // 0: authors.v1.Author
	(*Authors)(nil),           // 1: authors.v1.Authors
	(*ReadAuthorRequest)(nil), // 2: authors.v1.ReadAuthorRequest
	(*VoidNoParam)(nil),       // 3: authors.v1.voidNoParam
}
var file_authors_v1_authors_proto_depIdxs = []int32{
	0, // 0: authors.v1.Authors.author:type_name -> authors.v1.Author
	3, // 1: authors.v1.AuthorsService.readAuthors:input_type -> authors.v1.voidNoParam
	0, // 2: authors.v1.AuthorsService.createAuthor:input_type -> authors.v1.Author
	2, // 3: authors.v1.AuthorsService.readAuthor:input_type -> authors.v1.ReadAuthorRequest
	1, // 4: authors.v1.AuthorsService.readAuthors:output_type -> authors.v1.Authors
	0, // 5: authors.v1.AuthorsService.createAuthor:output_type -> authors.v1.Author
	0, // 6: authors.v1.AuthorsService.readAuthor:output_type -> authors.v1.Author
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_authors_v1_authors_proto_init() }
func file_authors_v1_authors_proto_init() {
	if File_authors_v1_authors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authors_v1_authors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_authors_v1_authors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authors); i {
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
		file_authors_v1_authors_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAuthorRequest); i {
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
		file_authors_v1_authors_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoidNoParam); i {
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
			RawDescriptor: file_authors_v1_authors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authors_v1_authors_proto_goTypes,
		DependencyIndexes: file_authors_v1_authors_proto_depIdxs,
		MessageInfos:      file_authors_v1_authors_proto_msgTypes,
	}.Build()
	File_authors_v1_authors_proto = out.File
	file_authors_v1_authors_proto_rawDesc = nil
	file_authors_v1_authors_proto_goTypes = nil
	file_authors_v1_authors_proto_depIdxs = nil
}
