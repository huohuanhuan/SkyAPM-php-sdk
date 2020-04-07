// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/trace-common.proto

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

type SpanType int32

const (
	SpanType_Entry SpanType = 0
	SpanType_Exit  SpanType = 1
	SpanType_Local SpanType = 2
)

var SpanType_name = map[int32]string{
	0: "Entry",
	1: "Exit",
	2: "Local",
}

var SpanType_value = map[string]int32{
	"Entry": 0,
	"Exit":  1,
	"Local": 2,
}

func (x SpanType) String() string {
	return proto.EnumName(SpanType_name, int32(x))
}

func (SpanType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f5d7d2e503402948, []int{0}
}

type RefType int32

const (
	RefType_CrossProcess RefType = 0
	RefType_CrossThread  RefType = 1
)

var RefType_name = map[int32]string{
	0: "CrossProcess",
	1: "CrossThread",
}

var RefType_value = map[string]int32{
	"CrossProcess": 0,
	"CrossThread":  1,
}

func (x RefType) String() string {
	return proto.EnumName(RefType_name, int32(x))
}

func (RefType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f5d7d2e503402948, []int{1}
}

type SpanLayer int32

const (
	SpanLayer_Unknown      SpanLayer = 0
	SpanLayer_Database     SpanLayer = 1
	SpanLayer_RPCFramework SpanLayer = 2
	SpanLayer_Http         SpanLayer = 3
	SpanLayer_MQ           SpanLayer = 4
	SpanLayer_Cache        SpanLayer = 5
)

var SpanLayer_name = map[int32]string{
	0: "Unknown",
	1: "Database",
	2: "RPCFramework",
	3: "Http",
	4: "MQ",
	5: "Cache",
}

var SpanLayer_value = map[string]int32{
	"Unknown":      0,
	"Database":     1,
	"RPCFramework": 2,
	"Http":         3,
	"MQ":           4,
	"Cache":        5,
}

func (x SpanLayer) String() string {
	return proto.EnumName(SpanLayer_name, int32(x))
}

func (SpanLayer) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f5d7d2e503402948, []int{2}
}

type UpstreamSegment struct {
	GlobalTraceIds       []*UniqueId `protobuf:"bytes,1,rep,name=globalTraceIds,proto3" json:"globalTraceIds,omitempty"`
	Segment              []byte      `protobuf:"bytes,2,opt,name=segment,proto3" json:"segment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpstreamSegment) Reset()         { *m = UpstreamSegment{} }
func (m *UpstreamSegment) String() string { return proto.CompactTextString(m) }
func (*UpstreamSegment) ProtoMessage()    {}
func (*UpstreamSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5d7d2e503402948, []int{0}
}

func (m *UpstreamSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSegment.Unmarshal(m, b)
}
func (m *UpstreamSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSegment.Marshal(b, m, deterministic)
}
func (m *UpstreamSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSegment.Merge(m, src)
}
func (m *UpstreamSegment) XXX_Size() int {
	return xxx_messageInfo_UpstreamSegment.Size(m)
}
func (m *UpstreamSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSegment.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSegment proto.InternalMessageInfo

func (m *UpstreamSegment) GetGlobalTraceIds() []*UniqueId {
	if m != nil {
		return m.GlobalTraceIds
	}
	return nil
}

func (m *UpstreamSegment) GetSegment() []byte {
	if m != nil {
		return m.Segment
	}
	return nil
}

type UniqueId struct {
	IdParts              []int64  `protobuf:"varint,1,rep,packed,name=idParts,proto3" json:"idParts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UniqueId) Reset()         { *m = UniqueId{} }
func (m *UniqueId) String() string { return proto.CompactTextString(m) }
func (*UniqueId) ProtoMessage()    {}
func (*UniqueId) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5d7d2e503402948, []int{1}
}

func (m *UniqueId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UniqueId.Unmarshal(m, b)
}
func (m *UniqueId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UniqueId.Marshal(b, m, deterministic)
}
func (m *UniqueId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UniqueId.Merge(m, src)
}
func (m *UniqueId) XXX_Size() int {
	return xxx_messageInfo_UniqueId.Size(m)
}
func (m *UniqueId) XXX_DiscardUnknown() {
	xxx_messageInfo_UniqueId.DiscardUnknown(m)
}

var xxx_messageInfo_UniqueId proto.InternalMessageInfo

func (m *UniqueId) GetIdParts() []int64 {
	if m != nil {
		return m.IdParts
	}
	return nil
}

func init() {
	proto.RegisterEnum("skywalking.network.protocol.common.v1.SpanType", SpanType_name, SpanType_value)
	proto.RegisterEnum("skywalking.network.protocol.common.v1.RefType", RefType_name, RefType_value)
	proto.RegisterEnum("skywalking.network.protocol.common.v1.SpanLayer", SpanLayer_name, SpanLayer_value)
	proto.RegisterType((*UpstreamSegment)(nil), "skywalking.network.protocol.common.v1.UpstreamSegment")
	proto.RegisterType((*UniqueId)(nil), "skywalking.network.protocol.common.v1.UniqueId")
}

func init() { proto.RegisterFile("common/trace-common.proto", fileDescriptor_f5d7d2e503402948) }

var fileDescriptor_f5d7d2e503402948 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0x87, 0x37, 0xd9, 0xb6, 0x9b, 0xce, 0xae, 0xa8, 0xe5, 0xd3, 0xc2, 0xa9, 0xaa, 0x8a, 0x54,
	0x45, 0x90, 0x50, 0x78, 0x03, 0x96, 0x22, 0x2a, 0x15, 0x14, 0xb2, 0x1b, 0x55, 0xe2, 0x36, 0x9b,
	0x0c, 0x69, 0x94, 0xc4, 0x0e, 0xb6, 0xbb, 0x21, 0x77, 0x9e, 0x86, 0xa7, 0x44, 0xce, 0x1f, 0x40,
	0x9c, 0x38, 0xfe, 0xc6, 0x9f, 0xbf, 0xb1, 0x67, 0xe0, 0x69, 0x2a, 0xeb, 0x5a, 0x8a, 0xd0, 0x28,
	0x4c, 0xe9, 0xe5, 0x10, 0x82, 0x46, 0x49, 0x23, 0xf9, 0x73, 0x5d, 0x76, 0x2d, 0x56, 0x65, 0x21,
	0xf2, 0x40, 0x90, 0x69, 0xa5, 0x2a, 0x87, 0x93, 0x54, 0x56, 0xc1, 0x48, 0x1e, 0xae, 0x2f, 0x7e,
	0x38, 0x70, 0x96, 0x34, 0xda, 0x28, 0xc2, 0x7a, 0x4b, 0x79, 0x4d, 0xc2, 0xf0, 0x7b, 0x78, 0x92,
	0x57, 0x72, 0x8f, 0xd5, 0xce, 0x6a, 0x6f, 0x33, 0xbd, 0x76, 0xce, 0xe7, 0x57, 0xcb, 0xd7, 0x61,
	0xf0, 0x5f, 0xce, 0x20, 0x11, 0xc5, 0xb7, 0x47, 0xba, 0xcd, 0xe2, 0x7f, 0x34, 0x7c, 0x0d, 0x0b,
	0x3d, 0xf4, 0x58, 0xbb, 0xe7, 0xce, 0xd5, 0x2a, 0x9e, 0xe2, 0xc5, 0x25, 0x78, 0xd3, 0x2d, 0x4b,
	0x15, 0x59, 0x84, 0xca, 0x0c, 0x7d, 0xe7, 0xf1, 0x14, 0x7d, 0x1f, 0xbc, 0x6d, 0x83, 0x62, 0xd7,
	0x35, 0xc4, 0x4f, 0xe1, 0xf8, 0x46, 0x18, 0xd5, 0xb1, 0x19, 0xf7, 0xe0, 0xe8, 0xe6, 0x7b, 0x61,
	0x98, 0x63, 0x8b, 0x77, 0x32, 0xc5, 0x8a, 0xb9, 0xfe, 0x0b, 0x58, 0xc4, 0xf4, 0xb5, 0x47, 0x19,
	0xac, 0x36, 0x4a, 0x6a, 0x1d, 0x29, 0x99, 0x92, 0xd6, 0x6c, 0xc6, 0xcf, 0x60, 0xd9, 0x57, 0x76,
	0x0f, 0x8a, 0x30, 0x63, 0x8e, 0x9f, 0xc0, 0xa9, 0x35, 0xdf, 0x61, 0x47, 0x8a, 0x2f, 0x61, 0x91,
	0x88, 0x52, 0xc8, 0x56, 0xb0, 0x19, 0x5f, 0x81, 0xf7, 0x0e, 0x0d, 0xee, 0x51, 0x13, 0x73, 0xac,
	0x2a, 0x8e, 0x36, 0xef, 0x15, 0xd6, 0x64, 0x7f, 0xcf, 0x5c, 0xdb, 0xfc, 0x83, 0x31, 0x0d, 0x9b,
	0xf3, 0x13, 0x70, 0x3f, 0x7e, 0x66, 0x47, 0xf6, 0x11, 0x1b, 0x4c, 0x1f, 0x88, 0x1d, 0xbf, 0x6d,
	0xe1, 0x95, 0x54, 0x79, 0x80, 0x8d, 0xcd, 0x7f, 0x4f, 0x0f, 0x9b, 0xfa, 0xf7, 0x04, 0x2b, 0x14,
	0xf9, 0x23, 0xe6, 0x14, 0x60, 0x4e, 0xc2, 0x44, 0xce, 0x97, 0xcb, 0x3f, 0x60, 0x38, 0x42, 0xe1,
	0x04, 0x85, 0x3d, 0x14, 0x1e, 0xae, 0x7f, 0xba, 0xcf, 0xb6, 0x65, 0x77, 0x3f, 0xfa, 0x3e, 0x0d,
	0x58, 0x34, 0x2e, 0x63, 0x7f, 0xd2, 0xaf, 0xe5, 0xcd, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22,
	0x0b, 0x4d, 0x09, 0x21, 0x02, 0x00, 0x00,
}
