// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language-agent/Downstream.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// nothing down stream from collector yet.
type Downstream struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Downstream) Reset()         { *m = Downstream{} }
func (m *Downstream) String() string { return proto.CompactTextString(m) }
func (*Downstream) ProtoMessage()    {}
func (*Downstream) Descriptor() ([]byte, []int) {
	return fileDescriptor_625677b377236e86, []int{0}
}

func (m *Downstream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Downstream.Unmarshal(m, b)
}
func (m *Downstream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Downstream.Marshal(b, m, deterministic)
}
func (m *Downstream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Downstream.Merge(m, src)
}
func (m *Downstream) XXX_Size() int {
	return xxx_messageInfo_Downstream.Size(m)
}
func (m *Downstream) XXX_DiscardUnknown() {
	xxx_messageInfo_Downstream.DiscardUnknown(m)
}

var xxx_messageInfo_Downstream proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Downstream)(nil), "Downstream")
}

func init() { proto.RegisterFile("language-agent/Downstream.proto", fileDescriptor_625677b377236e86) }

var fileDescriptor_625677b377236e86 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcf, 0x49, 0xcc, 0x4b,
	0x2f, 0x4d, 0x4c, 0x4f, 0xd5, 0x4d, 0x4c, 0x4f, 0xcd, 0x2b, 0xd1, 0x77, 0xc9, 0x2f, 0xcf, 0x2b,
	0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xe2, 0xe1, 0xe2, 0x42,
	0x88, 0x39, 0x95, 0x73, 0x19, 0xe4, 0x17, 0xa5, 0xeb, 0x25, 0x16, 0x24, 0x26, 0x67, 0xa4, 0xea,
	0x15, 0x67, 0x57, 0x96, 0x27, 0xe6, 0x64, 0x67, 0xe6, 0x81, 0x44, 0x72, 0xf5, 0xf2, 0x52, 0x4b,
	0xca, 0xf3, 0x8b, 0xb2, 0xf5, 0x60, 0x46, 0xea, 0x81, 0x8d, 0x0c, 0x60, 0x8c, 0x52, 0x41, 0x28,
	0xd4, 0x87, 0x2a, 0xd2, 0x87, 0x29, 0xd2, 0x87, 0xd8, 0x5b, 0x66, 0xb8, 0x8a, 0x49, 0x2a, 0x38,
	0xbb, 0x32, 0x1c, 0x6a, 0x9e, 0x1f, 0x44, 0x59, 0x00, 0xc8, 0x0d, 0xc9, 0xf9, 0x39, 0x49, 0x6c,
	0x60, 0xd7, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x88, 0x81, 0x76, 0xd1, 0xb0, 0x00, 0x00,
	0x00,
}
