// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validator.proto

package validator

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type FieldValidator struct {
	// Uses a Golang RE2-syntax regex to match the field contents.
	Regex *StringValue `protobuf:"bytes,1,opt,name=regex,proto3" json:"regex,omitempty"`
	// Field value of integer strictly greater than this value.
	IntGt *Int64Value `protobuf:"bytes,2,opt,name=int_gt,json=intGt,proto3" json:"int_gt,omitempty"`
	// Field value of integer strictly smaller than this value.
	IntLt *Int64Value `protobuf:"bytes,3,opt,name=int_lt,json=intLt,proto3" json:"int_lt,omitempty"`
	// Used for nested message types, requires that the message type exists.
	MsgExists bool `protobuf:"varint,4,opt,name=msg_exists,json=msgExists,proto3" json:"msg_exists,omitempty"`
	// Human error specifies a user-customizable error that is visible to the user.
	HumanError string `protobuf:"bytes,5,opt,name=human_error,json=humanError,proto3" json:"human_error,omitempty"`
	// Field value of double strictly greater than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatGt *DoubleValue `protobuf:"bytes,6,opt,name=float_gt,json=floatGt,proto3" json:"float_gt,omitempty"`
	// Field value of double strictly smaller than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatLt *DoubleValue `protobuf:"bytes,7,opt,name=float_lt,json=floatLt,proto3" json:"float_lt,omitempty"`
	// Field value of double describing the epsilon within which
	// any comparison should be considered to be true. For example,
	// when using float_gt = 0.35, using a float_epsilon of 0.05
	// would mean that any value above 0.30 is acceptable. It can be
	// thought of as a {float_value_condition} +- {float_epsilon}.
	// If unset, no correction for floating point inaccuracies in
	// comparisons will be attempted.
	FloatEpsilon *DoubleValue `protobuf:"bytes,8,opt,name=float_epsilon,json=floatEpsilon,proto3" json:"float_epsilon,omitempty"`
	// Floating-point value compared to which the field content should be greater or equal.
	FloatGte *DoubleValue `protobuf:"bytes,9,opt,name=float_gte,json=floatGte,proto3" json:"float_gte,omitempty"`
	// Floating-point value compared to which the field content should be smaller or equal.
	FloatLte *DoubleValue `protobuf:"bytes,10,opt,name=float_lte,json=floatLte,proto3" json:"float_lte,omitempty"`
	// Used for string fields, requires the string to be not empty (i.e different from "").
	StringNotEmpty bool `protobuf:"varint,11,opt,name=string_not_empty,json=stringNotEmpty,proto3" json:"string_not_empty,omitempty"`
	// Repeated field with at least this number of elements.
	RepeatedCountMin *Int64Value `protobuf:"bytes,12,opt,name=repeated_count_min,json=repeatedCountMin,proto3" json:"repeated_count_min,omitempty"`
	// Repeated field with at most this number of elements.
	RepeatedCountMax *Int64Value `protobuf:"bytes,13,opt,name=repeated_count_max,json=repeatedCountMax,proto3" json:"repeated_count_max,omitempty"`
	// Field value of length greater than this value.
	LengthGt *Int64Value `protobuf:"bytes,14,opt,name=length_gt,json=lengthGt,proto3" json:"length_gt,omitempty"`
	// Field value of length smaller than this value.
	LengthLt *Int64Value `protobuf:"bytes,15,opt,name=length_lt,json=lengthLt,proto3" json:"length_lt,omitempty"`
	// Field value of length strictly equal to this value.
	LengthEq *Int64Value `protobuf:"bytes,16,opt,name=length_eq,json=lengthEq,proto3" json:"length_eq,omitempty"`
	// Requires that the value is in the enum.
	IsInEnum bool `protobuf:"varint,17,opt,name=is_in_enum,json=isInEnum,proto3" json:"is_in_enum,omitempty"`
	// Ensures that a string value is in UUID format.
	// uuid_ver specifies the valid UUID versions. Valid values are: 0-5.
	// If uuid_ver is 0 all UUID versions are accepted.
	UuidVer *Int32Value `protobuf:"bytes,18,opt,name=uuid_ver,json=uuidVer,proto3" json:"uuid_ver,omitempty"`
	// Ensure that OneSignal's internal UUID wrapper `uuid.v1.UUID` contains a valid UUID
	OnesignalUuidV1      bool     `protobuf:"varint,19,opt,name=onesignal_uuid_v1,json=onesignalUuidV1,proto3" json:"onesignal_uuid_v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldValidator) Reset()         { *m = FieldValidator{} }
func (m *FieldValidator) String() string { return proto.CompactTextString(m) }
func (*FieldValidator) ProtoMessage()    {}
func (*FieldValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{0}
}
func (m *FieldValidator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldValidator.Unmarshal(m, b)
}
func (m *FieldValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldValidator.Marshal(b, m, deterministic)
}
func (m *FieldValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldValidator.Merge(m, src)
}
func (m *FieldValidator) XXX_Size() int {
	return xxx_messageInfo_FieldValidator.Size(m)
}
func (m *FieldValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldValidator.DiscardUnknown(m)
}

var xxx_messageInfo_FieldValidator proto.InternalMessageInfo

func (m *FieldValidator) GetRegex() *StringValue {
	if m != nil {
		return m.Regex
	}
	return nil
}

func (m *FieldValidator) GetIntGt() *Int64Value {
	if m != nil {
		return m.IntGt
	}
	return nil
}

func (m *FieldValidator) GetIntLt() *Int64Value {
	if m != nil {
		return m.IntLt
	}
	return nil
}

func (m *FieldValidator) GetMsgExists() bool {
	if m != nil {
		return m.MsgExists
	}
	return false
}

func (m *FieldValidator) GetHumanError() string {
	if m != nil {
		return m.HumanError
	}
	return ""
}

func (m *FieldValidator) GetFloatGt() *DoubleValue {
	if m != nil {
		return m.FloatGt
	}
	return nil
}

func (m *FieldValidator) GetFloatLt() *DoubleValue {
	if m != nil {
		return m.FloatLt
	}
	return nil
}

func (m *FieldValidator) GetFloatEpsilon() *DoubleValue {
	if m != nil {
		return m.FloatEpsilon
	}
	return nil
}

func (m *FieldValidator) GetFloatGte() *DoubleValue {
	if m != nil {
		return m.FloatGte
	}
	return nil
}

func (m *FieldValidator) GetFloatLte() *DoubleValue {
	if m != nil {
		return m.FloatLte
	}
	return nil
}

func (m *FieldValidator) GetStringNotEmpty() bool {
	if m != nil {
		return m.StringNotEmpty
	}
	return false
}

func (m *FieldValidator) GetRepeatedCountMin() *Int64Value {
	if m != nil {
		return m.RepeatedCountMin
	}
	return nil
}

func (m *FieldValidator) GetRepeatedCountMax() *Int64Value {
	if m != nil {
		return m.RepeatedCountMax
	}
	return nil
}

func (m *FieldValidator) GetLengthGt() *Int64Value {
	if m != nil {
		return m.LengthGt
	}
	return nil
}

func (m *FieldValidator) GetLengthLt() *Int64Value {
	if m != nil {
		return m.LengthLt
	}
	return nil
}

func (m *FieldValidator) GetLengthEq() *Int64Value {
	if m != nil {
		return m.LengthEq
	}
	return nil
}

func (m *FieldValidator) GetIsInEnum() bool {
	if m != nil {
		return m.IsInEnum
	}
	return false
}

func (m *FieldValidator) GetUuidVer() *Int32Value {
	if m != nil {
		return m.UuidVer
	}
	return nil
}

func (m *FieldValidator) GetOnesignalUuidV1() bool {
	if m != nil {
		return m.OnesignalUuidV1
	}
	return false
}

type Int64Value struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64Value) Reset()         { *m = Int64Value{} }
func (m *Int64Value) String() string { return proto.CompactTextString(m) }
func (*Int64Value) ProtoMessage()    {}
func (*Int64Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{1}
}
func (m *Int64Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64Value.Unmarshal(m, b)
}
func (m *Int64Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64Value.Marshal(b, m, deterministic)
}
func (m *Int64Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64Value.Merge(m, src)
}
func (m *Int64Value) XXX_Size() int {
	return xxx_messageInfo_Int64Value.Size(m)
}
func (m *Int64Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64Value.DiscardUnknown(m)
}

var xxx_messageInfo_Int64Value proto.InternalMessageInfo

func (m *Int64Value) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Int32Value struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32Value) Reset()         { *m = Int32Value{} }
func (m *Int32Value) String() string { return proto.CompactTextString(m) }
func (*Int32Value) ProtoMessage()    {}
func (*Int32Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{2}
}
func (m *Int32Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32Value.Unmarshal(m, b)
}
func (m *Int32Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32Value.Marshal(b, m, deterministic)
}
func (m *Int32Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32Value.Merge(m, src)
}
func (m *Int32Value) XXX_Size() int {
	return xxx_messageInfo_Int32Value.Size(m)
}
func (m *Int32Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32Value.DiscardUnknown(m)
}

var xxx_messageInfo_Int32Value proto.InternalMessageInfo

func (m *Int32Value) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type StringValue struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringValue) Reset()         { *m = StringValue{} }
func (m *StringValue) String() string { return proto.CompactTextString(m) }
func (*StringValue) ProtoMessage()    {}
func (*StringValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{3}
}
func (m *StringValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringValue.Unmarshal(m, b)
}
func (m *StringValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringValue.Marshal(b, m, deterministic)
}
func (m *StringValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringValue.Merge(m, src)
}
func (m *StringValue) XXX_Size() int {
	return xxx_messageInfo_StringValue.Size(m)
}
func (m *StringValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StringValue.DiscardUnknown(m)
}

var xxx_messageInfo_StringValue proto.InternalMessageInfo

func (m *StringValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type DoubleValue struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleValue) Reset()         { *m = DoubleValue{} }
func (m *DoubleValue) String() string { return proto.CompactTextString(m) }
func (*DoubleValue) ProtoMessage()    {}
func (*DoubleValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{4}
}
func (m *DoubleValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleValue.Unmarshal(m, b)
}
func (m *DoubleValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleValue.Marshal(b, m, deterministic)
}
func (m *DoubleValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleValue.Merge(m, src)
}
func (m *DoubleValue) XXX_Size() int {
	return xxx_messageInfo_DoubleValue.Size(m)
}
func (m *DoubleValue) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleValue.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleValue proto.InternalMessageInfo

func (m *DoubleValue) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type OneofValidator struct {
	// Require that one of the oneof fields is set.
	Required             bool     `protobuf:"varint,1,opt,name=required,proto3" json:"required,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OneofValidator) Reset()         { *m = OneofValidator{} }
func (m *OneofValidator) String() string { return proto.CompactTextString(m) }
func (*OneofValidator) ProtoMessage()    {}
func (*OneofValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{5}
}
func (m *OneofValidator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OneofValidator.Unmarshal(m, b)
}
func (m *OneofValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OneofValidator.Marshal(b, m, deterministic)
}
func (m *OneofValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OneofValidator.Merge(m, src)
}
func (m *OneofValidator) XXX_Size() int {
	return xxx_messageInfo_OneofValidator.Size(m)
}
func (m *OneofValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_OneofValidator.DiscardUnknown(m)
}

var xxx_messageInfo_OneofValidator proto.InternalMessageInfo

func (m *OneofValidator) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*FieldValidator)(nil),
	Field:         65020,
	Name:          "validator.field",
	Tag:           "bytes,65020,opt,name=field",
	Filename:      "validator.proto",
}

var E_Oneof = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.OneofOptions)(nil),
	ExtensionType: (*OneofValidator)(nil),
	Field:         65021,
	Name:          "validator.oneof",
	Tag:           "bytes,65021,opt,name=oneof",
	Filename:      "validator.proto",
}

func init() {
	proto.RegisterType((*FieldValidator)(nil), "validator.FieldValidator")
	proto.RegisterType((*Int64Value)(nil), "validator.Int64Value")
	proto.RegisterType((*Int32Value)(nil), "validator.Int32Value")
	proto.RegisterType((*StringValue)(nil), "validator.StringValue")
	proto.RegisterType((*DoubleValue)(nil), "validator.DoubleValue")
	proto.RegisterType((*OneofValidator)(nil), "validator.OneofValidator")
	proto.RegisterExtension(E_Field)
	proto.RegisterExtension(E_Oneof)
}

func init() { proto.RegisterFile("validator.proto", fileDescriptor_bf1c6ec7c0d80dd5) }

var fileDescriptor_bf1c6ec7c0d80dd5 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x5f, 0x4f, 0xdb, 0x3c,
	0x14, 0xc6, 0xd5, 0x97, 0xb7, 0x90, 0x9c, 0x42, 0x01, 0xef, 0x8f, 0x3c, 0x34, 0xb4, 0x8a, 0xdd,
	0x54, 0x13, 0xb4, 0xa3, 0xa0, 0x5d, 0x8c, 0xbb, 0xb1, 0x0e, 0x21, 0x75, 0x63, 0x0a, 0x1a, 0x17,
	0xbb, 0x89, 0x52, 0x7a, 0x1a, 0x2c, 0x39, 0x76, 0x71, 0x4e, 0x50, 0xf7, 0xb9, 0xf6, 0xf1, 0xb6,
	0x49, 0x53, 0xec, 0xa6, 0x4d, 0xb5, 0xaa, 0xeb, 0xee, 0xe2, 0xe3, 0xdf, 0xf3, 0x9c, 0x13, 0x1f,
	0xfb, 0xc0, 0xf6, 0x43, 0x24, 0xc5, 0x20, 0x22, 0x6d, 0x5a, 0x23, 0xa3, 0x49, 0x33, 0x7f, 0x1a,
	0xd8, 0x6b, 0xc4, 0x5a, 0xc7, 0x12, 0xdb, 0x76, 0xa3, 0x9f, 0x0d, 0xdb, 0x03, 0x4c, 0x6f, 0x8d,
	0x18, 0x4d, 0xe1, 0x83, 0xef, 0x1b, 0x50, 0xff, 0x20, 0x50, 0x0e, 0x6e, 0x0a, 0x11, 0x3b, 0x84,
	0xaa, 0xc1, 0x18, 0xc7, 0xbc, 0xd2, 0xa8, 0x34, 0x6b, 0x9d, 0xa7, 0xad, 0x59, 0x82, 0x6b, 0x32,
	0x42, 0xc5, 0x37, 0x91, 0xcc, 0x30, 0x70, 0x10, 0x3b, 0x84, 0x75, 0xa1, 0x28, 0x8c, 0x89, 0xff,
	0x67, 0xf1, 0x27, 0x25, 0xfc, 0x52, 0xd1, 0x9b, 0xd3, 0x09, 0x2d, 0x14, 0x5d, 0x50, 0x41, 0x4b,
	0xe2, 0x6b, 0x7f, 0xa3, 0x7b, 0xc4, 0xf6, 0x01, 0x92, 0x34, 0x0e, 0x71, 0x2c, 0x52, 0x4a, 0xf9,
	0xff, 0x8d, 0x4a, 0xd3, 0x0b, 0xfc, 0x24, 0x8d, 0xbb, 0x36, 0xc0, 0x5e, 0x40, 0xed, 0x2e, 0x4b,
	0x22, 0x15, 0xa2, 0x31, 0xda, 0xf0, 0x6a, 0xa3, 0xd2, 0xf4, 0x03, 0xb0, 0xa1, 0x6e, 0x1e, 0x61,
	0xc7, 0xe0, 0x0d, 0xa5, 0x8e, 0x6c, 0x75, 0xeb, 0x7f, 0xfc, 0xcc, 0x7b, 0x9d, 0xf5, 0x25, 0xba,
	0x84, 0x1b, 0x96, 0xbb, 0xa0, 0x99, 0x44, 0x12, 0xdf, 0x58, 0x41, 0xd2, 0x23, 0x76, 0x06, 0x5b,
	0x4e, 0x82, 0xa3, 0x54, 0x48, 0xad, 0xb8, 0xb7, 0x54, 0xb7, 0x69, 0xe1, 0xae, 0x63, 0xd9, 0x09,
	0xf8, 0x45, 0x89, 0xc8, 0xfd, 0xa5, 0x42, 0x6f, 0x52, 0x23, 0xce, 0x44, 0x92, 0x90, 0xc3, 0x0a,
	0xa2, 0x1e, 0x21, 0x6b, 0xc2, 0x4e, 0x6a, 0xdb, 0x17, 0x2a, 0x4d, 0x21, 0x26, 0x23, 0xfa, 0xc6,
	0x6b, 0xf6, 0x48, 0xeb, 0x2e, 0xfe, 0x49, 0x53, 0x37, 0x8f, 0xb2, 0x73, 0x60, 0x06, 0x47, 0x18,
	0x11, 0x0e, 0xc2, 0x5b, 0x9d, 0x29, 0x0a, 0x13, 0xa1, 0xf8, 0xe6, 0xb2, 0x86, 0xed, 0x14, 0x82,
	0xf3, 0x9c, 0xff, 0x28, 0xd4, 0x22, 0x93, 0x68, 0xcc, 0xb7, 0xfe, 0xc1, 0x24, 0x1a, 0xb3, 0x0e,
	0xf8, 0x12, 0x55, 0x4c, 0x77, 0x79, 0x07, 0xeb, 0xcb, 0xb4, 0x9e, 0xe3, 0x2e, 0xa8, 0xa4, 0x91,
	0xc4, 0xb7, 0x57, 0xd0, 0xf4, 0xca, 0x1a, 0xbc, 0xe7, 0x3b, 0x2b, 0x68, 0xba, 0xf7, 0xec, 0x39,
	0x80, 0x48, 0x43, 0xa1, 0x42, 0x54, 0x59, 0xc2, 0x77, 0xed, 0x49, 0x7a, 0x22, 0xbd, 0x54, 0x5d,
	0x95, 0x25, 0xec, 0x35, 0x78, 0x59, 0x26, 0x06, 0xe1, 0x03, 0x1a, 0xce, 0x16, 0x19, 0x9e, 0x74,
	0x26, 0xd7, 0x28, 0xc7, 0x6e, 0xd0, 0xb0, 0x57, 0xb0, 0xab, 0x15, 0xa6, 0x22, 0x56, 0x91, 0x0c,
	0x9d, 0xf6, 0x98, 0x3f, 0xb2, 0xb6, 0xdb, 0xd3, 0x8d, 0x2f, 0x39, 0x7c, 0x7c, 0x70, 0x00, 0x30,
	0xab, 0x89, 0x3d, 0x86, 0xea, 0x43, 0xfe, 0x61, 0x1f, 0xec, 0x5a, 0xe0, 0x16, 0x13, 0x66, 0x92,
	0x66, 0x9e, 0xa9, 0x16, 0xcc, 0x4b, 0xa8, 0x95, 0x9e, 0xf4, 0x3c, 0xe4, 0x97, 0xa0, 0xd2, 0x8d,
	0x9a, 0x87, 0x2a, 0x05, 0x74, 0x08, 0xf5, 0x2b, 0x85, 0x7a, 0x38, 0x1b, 0x23, 0x7b, 0xe0, 0x19,
	0xbc, 0xcf, 0x84, 0xc1, 0x81, 0x45, 0xbd, 0x60, 0xba, 0x7e, 0xfb, 0x19, 0xaa, 0xc3, 0x7c, 0xe8,
	0xb0, 0xfd, 0x96, 0x9b, 0x50, 0xad, 0x62, 0x42, 0xb5, 0xec, 0x30, 0xba, 0x1a, 0x91, 0xd0, 0x2a,
	0xe5, 0x3f, 0x7f, 0xb8, 0x31, 0xf1, 0xac, 0x74, 0x76, 0xf3, 0xd3, 0x2a, 0x70, 0x46, 0xb9, 0xa3,
	0xce, 0xf3, 0x2f, 0x70, 0xb4, 0x75, 0x15, 0x8e, 0xbf, 0x16, 0x38, 0xce, 0x17, 0x1e, 0x38, 0xa3,
	0x77, 0xa7, 0x5f, 0x3b, 0xb1, 0xa0, 0xbb, 0xac, 0xdf, 0xba, 0xd5, 0x49, 0xfb, 0x4a, 0xe1, 0xb5,
	0xed, 0x40, 0x3b, 0xd6, 0x47, 0xd6, 0xfa, 0x68, 0x6a, 0x90, 0x9e, 0x4d, 0x3f, 0xfb, 0xeb, 0x76,
	0xef, 0xe4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x93, 0x58, 0x5e, 0x96, 0x05, 0x00, 0x00,
}
