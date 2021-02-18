// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: examples/uuid.proto

package validator_examples

import (
	_ "github.com/OneSignal/go-proto-validators"
	v1 "github.com/OneSignal/schema.protobuf/gen/go/uuid/v1"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UUIDMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user_id must be a valid version 4 UUID.
	UserId string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Xyz    *v1.UUID `protobuf:"bytes,2,opt,name=xyz,proto3" json:"xyz,omitempty"`
	Abc    []string `protobuf:"bytes,3,rep,name=abc,proto3" json:"abc,omitempty"`
}

func (x *UUIDMsg) Reset() {
	*x = UUIDMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_uuid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUIDMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUIDMsg) ProtoMessage() {}

func (x *UUIDMsg) ProtoReflect() protoreflect.Message {
	mi := &file_examples_uuid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUIDMsg.ProtoReflect.Descriptor instead.
func (*UUIDMsg) Descriptor() ([]byte, []int) {
	return file_examples_uuid_proto_rawDescGZIP(), []int{0}
}

func (x *UUIDMsg) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UUIDMsg) GetXyz() *v1.UUID {
	if x != nil {
		return x.Xyz
	}
	return nil
}

func (x *UUIDMsg) GetAbc() []string {
	if x != nil {
		return x.Abc
	}
	return nil
}

type Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Foo) Reset() {
	*x = Foo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_uuid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_examples_uuid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_examples_uuid_proto_rawDescGZIP(), []int{1}
}

var File_examples_uuid_proto protoreflect.FileDescriptor

var file_examples_uuid_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x1a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4f, 0x6e, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2f,
	0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x75, 0x75, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x75, 0x69,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x07, 0x55, 0x55, 0x49, 0x44, 0x4d,
	0x73, 0x67, 0x12, 0x24, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0xdf, 0x1f, 0x07, 0x58, 0x01, 0x92, 0x01, 0x02, 0x08, 0x04,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x03, 0x78, 0x79, 0x7a, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x55, 0x49, 0x44, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x98, 0x01, 0x01, 0x52, 0x03, 0x78,
	0x79, 0x7a, 0x12, 0x18, 0x0a, 0x03, 0x61, 0x62, 0x63, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x03, 0x61, 0x62, 0x63, 0x22, 0x05, 0x0a, 0x03,
	0x46, 0x6f, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_examples_uuid_proto_rawDescOnce sync.Once
	file_examples_uuid_proto_rawDescData = file_examples_uuid_proto_rawDesc
)

func file_examples_uuid_proto_rawDescGZIP() []byte {
	file_examples_uuid_proto_rawDescOnce.Do(func() {
		file_examples_uuid_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_uuid_proto_rawDescData)
	})
	return file_examples_uuid_proto_rawDescData
}

var file_examples_uuid_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_examples_uuid_proto_goTypes = []interface{}{
	(*UUIDMsg)(nil), // 0: validator.examples.UUIDMsg
	(*Foo)(nil),     // 1: validator.examples.Foo
	(*v1.UUID)(nil), // 2: uuid.v1.UUID
}
var file_examples_uuid_proto_depIdxs = []int32{
	2, // 0: validator.examples.UUIDMsg.xyz:type_name -> uuid.v1.UUID
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_examples_uuid_proto_init() }
func file_examples_uuid_proto_init() {
	if File_examples_uuid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examples_uuid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUIDMsg); i {
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
		file_examples_uuid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foo); i {
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
			RawDescriptor: file_examples_uuid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_examples_uuid_proto_goTypes,
		DependencyIndexes: file_examples_uuid_proto_depIdxs,
		MessageInfos:      file_examples_uuid_proto_msgTypes,
	}.Build()
	File_examples_uuid_proto = out.File
	file_examples_uuid_proto_rawDesc = nil
	file_examples_uuid_proto_goTypes = nil
	file_examples_uuid_proto_depIdxs = nil
}
