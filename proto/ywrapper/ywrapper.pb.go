// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aristanetworks/ygot/proto/ywrapper/ywrapper.proto

package ywrapper

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// BytesValue is used to store a value which is a byte array, particularly
// the YANG binary type.
type BytesValue struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BytesValue) Reset()         { *m = BytesValue{} }
func (m *BytesValue) String() string { return proto.CompactTextString(m) }
func (*BytesValue) ProtoMessage()    {}
func (*BytesValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_ywrapper_bddc4fc8b83b6037, []int{0}
}
func (m *BytesValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesValue.Unmarshal(m, b)
}
func (m *BytesValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesValue.Marshal(b, m, deterministic)
}
func (dst *BytesValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesValue.Merge(dst, src)
}
func (m *BytesValue) XXX_Size() int {
	return xxx_messageInfo_BytesValue.Size(m)
}
func (m *BytesValue) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesValue.DiscardUnknown(m)
}

var xxx_messageInfo_BytesValue proto.InternalMessageInfo

func (m *BytesValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// BoolValue is used to store a value which is a boolean, particularly the
// YANG bool and empty types.
type BoolValue struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoolValue) Reset()         { *m = BoolValue{} }
func (m *BoolValue) String() string { return proto.CompactTextString(m) }
func (*BoolValue) ProtoMessage()    {}
func (*BoolValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_ywrapper_bddc4fc8b83b6037, []int{1}
}
func (m *BoolValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolValue.Unmarshal(m, b)
}
func (m *BoolValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolValue.Marshal(b, m, deterministic)
}
func (dst *BoolValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolValue.Merge(dst, src)
}
func (m *BoolValue) XXX_Size() int {
	return xxx_messageInfo_BoolValue.Size(m)
}
func (m *BoolValue) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolValue.DiscardUnknown(m)
}

var xxx_messageInfo_BoolValue proto.InternalMessageInfo

func (m *BoolValue) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

// Decimal64Value is used to store a value which is a decimal64, split into
// a digits field, and a precision field. The precision indicates the number
// of digits that occur after the decimal point in the digits field.
type Decimal64Value struct {
	Digits               int64    `protobuf:"varint,1,opt,name=digits,proto3" json:"digits,omitempty"`
	Precision            uint32   `protobuf:"varint,2,opt,name=precision,proto3" json:"precision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Decimal64Value) Reset()         { *m = Decimal64Value{} }
func (m *Decimal64Value) String() string { return proto.CompactTextString(m) }
func (*Decimal64Value) ProtoMessage()    {}
func (*Decimal64Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_ywrapper_bddc4fc8b83b6037, []int{2}
}
func (m *Decimal64Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Decimal64Value.Unmarshal(m, b)
}
func (m *Decimal64Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Decimal64Value.Marshal(b, m, deterministic)
}
func (dst *Decimal64Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Decimal64Value.Merge(dst, src)
}
func (m *Decimal64Value) XXX_Size() int {
	return xxx_messageInfo_Decimal64Value.Size(m)
}
func (m *Decimal64Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Decimal64Value.DiscardUnknown(m)
}

var xxx_messageInfo_Decimal64Value proto.InternalMessageInfo

func (m *Decimal64Value) GetDigits() int64 {
	if m != nil {
		return m.Digits
	}
	return 0
}

func (m *Decimal64Value) GetPrecision() uint32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

// IntValue stores a value which a signed integer, particularly the YANG
// int8, int16, int32, and int64 types.
type IntValue struct {
	Value                int64    `protobuf:"zigzag64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntValue) Reset()         { *m = IntValue{} }
func (m *IntValue) String() string { return proto.CompactTextString(m) }
func (*IntValue) ProtoMessage()    {}
func (*IntValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_ywrapper_bddc4fc8b83b6037, []int{3}
}
func (m *IntValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntValue.Unmarshal(m, b)
}
func (m *IntValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntValue.Marshal(b, m, deterministic)
}
func (dst *IntValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntValue.Merge(dst, src)
}
func (m *IntValue) XXX_Size() int {
	return xxx_messageInfo_IntValue.Size(m)
}
func (m *IntValue) XXX_DiscardUnknown() {
	xxx_messageInfo_IntValue.DiscardUnknown(m)
}

var xxx_messageInfo_IntValue proto.InternalMessageInfo

func (m *IntValue) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// StringValue stores a value which is a string, particularly the YANG
// string type.
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
	return fileDescriptor_ywrapper_bddc4fc8b83b6037, []int{4}
}
func (m *StringValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringValue.Unmarshal(m, b)
}
func (m *StringValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringValue.Marshal(b, m, deterministic)
}
func (dst *StringValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringValue.Merge(dst, src)
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

// UintVal is used to store a value which an unsigned integer, particularly
// the YANG uint8, uint16, uint32 and uint64 types.
type UintValue struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UintValue) Reset()         { *m = UintValue{} }
func (m *UintValue) String() string { return proto.CompactTextString(m) }
func (*UintValue) ProtoMessage()    {}
func (*UintValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_ywrapper_bddc4fc8b83b6037, []int{5}
}
func (m *UintValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UintValue.Unmarshal(m, b)
}
func (m *UintValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UintValue.Marshal(b, m, deterministic)
}
func (dst *UintValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UintValue.Merge(dst, src)
}
func (m *UintValue) XXX_Size() int {
	return xxx_messageInfo_UintValue.Size(m)
}
func (m *UintValue) XXX_DiscardUnknown() {
	xxx_messageInfo_UintValue.DiscardUnknown(m)
}

var xxx_messageInfo_UintValue proto.InternalMessageInfo

func (m *UintValue) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*BytesValue)(nil), "ywrapper.BytesValue")
	proto.RegisterType((*BoolValue)(nil), "ywrapper.BoolValue")
	proto.RegisterType((*Decimal64Value)(nil), "ywrapper.Decimal64Value")
	proto.RegisterType((*IntValue)(nil), "ywrapper.IntValue")
	proto.RegisterType((*StringValue)(nil), "ywrapper.StringValue")
	proto.RegisterType((*UintValue)(nil), "ywrapper.UintValue")
}

func init() {
	proto.RegisterFile("github.com/aristanetworks/ygot/proto/ywrapper/ywrapper.proto", fileDescriptor_ywrapper_bddc4fc8b83b6037)
}

var fileDescriptor_ywrapper_bddc4fc8b83b6037 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2c, 0xca, 0x2c, 0x2e, 0x49, 0xcc, 0x4b, 0x2d,
	0x29, 0xcf, 0x2f, 0xca, 0x2e, 0xd6, 0xaf, 0x4c, 0xcf, 0x2f, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9,
	0xd7, 0xaf, 0x2c, 0x2f, 0x4a, 0x2c, 0x28, 0x48, 0x2d, 0x82, 0x33, 0xf4, 0xc0, 0xe2, 0x42, 0x1c,
	0x30, 0xbe, 0x92, 0x12, 0x17, 0x97, 0x53, 0x65, 0x49, 0x6a, 0x71, 0x58, 0x62, 0x4e, 0x69, 0xaa,
	0x90, 0x08, 0x17, 0x6b, 0x19, 0x88, 0x21, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0xe1, 0x28,
	0x29, 0x72, 0x71, 0x3a, 0xe5, 0xe7, 0xe7, 0x60, 0x51, 0xc2, 0x01, 0x53, 0xe2, 0xc6, 0xc5, 0xe7,
	0x92, 0x9a, 0x9c, 0x99, 0x9b, 0x98, 0x63, 0x66, 0x02, 0x51, 0x27, 0xc6, 0xc5, 0x96, 0x92, 0x99,
	0x9e, 0x59, 0x52, 0x0c, 0x56, 0xc8, 0x1c, 0x04, 0xe5, 0x09, 0xc9, 0x70, 0x71, 0x16, 0x14, 0xa5,
	0x26, 0x67, 0x16, 0x67, 0xe6, 0xe7, 0x49, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x06, 0x21, 0x04, 0x94,
	0x14, 0xb8, 0x38, 0x3c, 0xf3, 0x4a, 0xb0, 0xd8, 0x24, 0x04, 0xb3, 0x49, 0x99, 0x8b, 0x3b, 0xb8,
	0xa4, 0x28, 0x33, 0x2f, 0x1d, 0x8b, 0x22, 0x4e, 0x24, 0x17, 0x87, 0x66, 0x62, 0x35, 0x87, 0x05,
	0xaa, 0x24, 0x89, 0x0d, 0x1c, 0x12, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x88, 0xd4,
	0xb3, 0x49, 0x01, 0x00, 0x00,
}
