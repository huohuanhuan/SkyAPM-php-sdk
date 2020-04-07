// Code generated by protoc-gen-go. DO NOT EDIT.
// source: profile/Profile.proto

package v3

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	v3 "github.com/SkyAPM/SkyAPM-php-sdk/reporter/network/common/v3"
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

type ProfileTaskCommandQuery struct {
	// current sniffer information
	Service         string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	ServiceInstance string `protobuf:"bytes,2,opt,name=serviceInstance,proto3" json:"serviceInstance,omitempty"`
	// last command timestamp
	LastCommandTime      int64    `protobuf:"varint,3,opt,name=lastCommandTime,proto3" json:"lastCommandTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileTaskCommandQuery) Reset()         { *m = ProfileTaskCommandQuery{} }
func (m *ProfileTaskCommandQuery) String() string { return proto.CompactTextString(m) }
func (*ProfileTaskCommandQuery) ProtoMessage()    {}
func (*ProfileTaskCommandQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed871c17f6550fe9, []int{0}
}

func (m *ProfileTaskCommandQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileTaskCommandQuery.Unmarshal(m, b)
}
func (m *ProfileTaskCommandQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileTaskCommandQuery.Marshal(b, m, deterministic)
}
func (m *ProfileTaskCommandQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileTaskCommandQuery.Merge(m, src)
}
func (m *ProfileTaskCommandQuery) XXX_Size() int {
	return xxx_messageInfo_ProfileTaskCommandQuery.Size(m)
}
func (m *ProfileTaskCommandQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileTaskCommandQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileTaskCommandQuery proto.InternalMessageInfo

func (m *ProfileTaskCommandQuery) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ProfileTaskCommandQuery) GetServiceInstance() string {
	if m != nil {
		return m.ServiceInstance
	}
	return ""
}

func (m *ProfileTaskCommandQuery) GetLastCommandTime() int64 {
	if m != nil {
		return m.LastCommandTime
	}
	return 0
}

// dumped thread snapshot
type ThreadSnapshot struct {
	// profile task id
	TaskId string `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	// dumped segment id
	TraceSegmentId string `protobuf:"bytes,2,opt,name=traceSegmentId,proto3" json:"traceSegmentId,omitempty"`
	// dump timestamp
	Time int64 `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	// snapshot dump sequence, start with zero
	Sequence int32 `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// snapshot stack
	Stack                *ThreadStack `protobuf:"bytes,5,opt,name=stack,proto3" json:"stack,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ThreadSnapshot) Reset()         { *m = ThreadSnapshot{} }
func (m *ThreadSnapshot) String() string { return proto.CompactTextString(m) }
func (*ThreadSnapshot) ProtoMessage()    {}
func (*ThreadSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed871c17f6550fe9, []int{1}
}

func (m *ThreadSnapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadSnapshot.Unmarshal(m, b)
}
func (m *ThreadSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadSnapshot.Marshal(b, m, deterministic)
}
func (m *ThreadSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadSnapshot.Merge(m, src)
}
func (m *ThreadSnapshot) XXX_Size() int {
	return xxx_messageInfo_ThreadSnapshot.Size(m)
}
func (m *ThreadSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadSnapshot proto.InternalMessageInfo

func (m *ThreadSnapshot) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *ThreadSnapshot) GetTraceSegmentId() string {
	if m != nil {
		return m.TraceSegmentId
	}
	return ""
}

func (m *ThreadSnapshot) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *ThreadSnapshot) GetSequence() int32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *ThreadSnapshot) GetStack() *ThreadStack {
	if m != nil {
		return m.Stack
	}
	return nil
}

type ThreadStack struct {
	// stack code signature list
	CodeSignatures       []string `protobuf:"bytes,1,rep,name=codeSignatures,proto3" json:"codeSignatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadStack) Reset()         { *m = ThreadStack{} }
func (m *ThreadStack) String() string { return proto.CompactTextString(m) }
func (*ThreadStack) ProtoMessage()    {}
func (*ThreadStack) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed871c17f6550fe9, []int{2}
}

func (m *ThreadStack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadStack.Unmarshal(m, b)
}
func (m *ThreadStack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadStack.Marshal(b, m, deterministic)
}
func (m *ThreadStack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadStack.Merge(m, src)
}
func (m *ThreadStack) XXX_Size() int {
	return xxx_messageInfo_ThreadStack.Size(m)
}
func (m *ThreadStack) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadStack.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadStack proto.InternalMessageInfo

func (m *ThreadStack) GetCodeSignatures() []string {
	if m != nil {
		return m.CodeSignatures
	}
	return nil
}

// profile task finished report
type ProfileTaskFinishReport struct {
	// current sniffer information
	Service         string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	ServiceInstance string `protobuf:"bytes,2,opt,name=serviceInstance,proto3" json:"serviceInstance,omitempty"`
	// profile task
	TaskId               string   `protobuf:"bytes,3,opt,name=taskId,proto3" json:"taskId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileTaskFinishReport) Reset()         { *m = ProfileTaskFinishReport{} }
func (m *ProfileTaskFinishReport) String() string { return proto.CompactTextString(m) }
func (*ProfileTaskFinishReport) ProtoMessage()    {}
func (*ProfileTaskFinishReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed871c17f6550fe9, []int{3}
}

func (m *ProfileTaskFinishReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileTaskFinishReport.Unmarshal(m, b)
}
func (m *ProfileTaskFinishReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileTaskFinishReport.Marshal(b, m, deterministic)
}
func (m *ProfileTaskFinishReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileTaskFinishReport.Merge(m, src)
}
func (m *ProfileTaskFinishReport) XXX_Size() int {
	return xxx_messageInfo_ProfileTaskFinishReport.Size(m)
}
func (m *ProfileTaskFinishReport) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileTaskFinishReport.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileTaskFinishReport proto.InternalMessageInfo

func (m *ProfileTaskFinishReport) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ProfileTaskFinishReport) GetServiceInstance() string {
	if m != nil {
		return m.ServiceInstance
	}
	return ""
}

func (m *ProfileTaskFinishReport) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func init() {
	proto.RegisterType((*ProfileTaskCommandQuery)(nil), "ProfileTaskCommandQuery")
	proto.RegisterType((*ThreadSnapshot)(nil), "ThreadSnapshot")
	proto.RegisterType((*ThreadStack)(nil), "ThreadStack")
	proto.RegisterType((*ProfileTaskFinishReport)(nil), "ProfileTaskFinishReport")
}

func init() { proto.RegisterFile("profile/Profile.proto", fileDescriptor_ed871c17f6550fe9) }

var fileDescriptor_ed871c17f6550fe9 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0xc5, 0xdb, 0xed, 0xc2, 0x4e, 0xd1, 0x16, 0x19, 0xb1, 0x44, 0x3d, 0x45, 0x39, 0xac, 0x72,
	0x72, 0xc4, 0x56, 0x7b, 0xe0, 0x84, 0x04, 0x12, 0x52, 0x2f, 0xa8, 0xa4, 0x95, 0x90, 0xb8, 0x19,
	0x67, 0x48, 0xa3, 0x24, 0x76, 0xb0, 0x9d, 0xae, 0x7a, 0xe1, 0xca, 0x8f, 0x70, 0xe2, 0x1b, 0xf8,
	0x38, 0xe4, 0xd4, 0xdd, 0x26, 0x45, 0xdc, 0x38, 0xd9, 0xf3, 0xfc, 0xec, 0x37, 0x7e, 0x33, 0x03,
	0x2f, 0x1a, 0xad, 0xbe, 0x16, 0x15, 0x26, 0xcb, 0xfd, 0xca, 0x1a, 0xad, 0xac, 0x9a, 0x3d, 0x17,
	0xaa, 0xae, 0x95, 0x4c, 0xde, 0x75, 0xcb, 0x1e, 0x8c, 0x7e, 0x10, 0x78, 0xe9, 0x69, 0x6b, 0x6e,
	0x4a, 0x77, 0xc6, 0x65, 0xf6, 0xb1, 0x45, 0xbd, 0xa3, 0x01, 0x3c, 0x36, 0xa8, 0xb7, 0x85, 0xc0,
	0x80, 0x84, 0x24, 0xbe, 0x4c, 0x0f, 0x21, 0x8d, 0x61, 0xea, 0xb7, 0x0b, 0x69, 0x2c, 0x97, 0x02,
	0x83, 0xb3, 0x8e, 0x71, 0x0a, 0x3b, 0x66, 0xc5, 0x8d, 0xf5, 0xef, 0xae, 0x8b, 0x1a, 0x83, 0x51,
	0x48, 0xe2, 0x51, 0x7a, 0x0a, 0x47, 0x3f, 0x09, 0x5c, 0xad, 0x37, 0x1a, 0x79, 0xb6, 0x92, 0xbc,
	0x31, 0x1b, 0x65, 0xe9, 0x35, 0x5c, 0x58, 0x6e, 0xca, 0x45, 0xe6, 0xf5, 0x7d, 0x44, 0x6f, 0xe0,
	0xca, 0x6a, 0x2e, 0x70, 0x85, 0x79, 0x8d, 0xd2, 0x2e, 0x32, 0xaf, 0x7e, 0x82, 0x52, 0x0a, 0xe7,
	0xf6, 0xa8, 0xd8, 0xed, 0xe9, 0x0c, 0x9e, 0x18, 0xfc, 0xd6, 0xa2, 0xcb, 0xf9, 0x3c, 0x24, 0xf1,
	0x38, 0x7d, 0x88, 0x69, 0x04, 0x63, 0x63, 0xb9, 0x28, 0x83, 0x71, 0x48, 0xe2, 0xc9, 0xed, 0x53,
	0xe6, 0xf3, 0x71, 0x58, 0xba, 0x3f, 0x8a, 0xee, 0x60, 0xd2, 0x43, 0x5d, 0x2a, 0x42, 0x65, 0xb8,
	0x2a, 0x72, 0xc9, 0x6d, 0xab, 0xd1, 0x04, 0x24, 0x1c, 0xb9, 0x54, 0x86, 0x68, 0xd4, 0x0e, 0x6c,
	0x7e, 0x5f, 0xc8, 0xc2, 0x6c, 0x52, 0x6c, 0x94, 0xb6, 0xff, 0xc5, 0xe6, 0xa3, 0x53, 0xa3, 0xbe,
	0x53, 0xb7, 0xbf, 0x09, 0x4c, 0x7a, 0xba, 0xf4, 0x0d, 0x5c, 0xe7, 0x68, 0xff, 0x2e, 0xb8, 0xa1,
	0x01, 0xfb, 0x47, 0x1b, 0xcc, 0x2e, 0xd9, 0x81, 0x14, 0x3d, 0xa2, 0xaf, 0x60, 0x2a, 0x54, 0x55,
	0xa1, 0xb0, 0x0f, 0x55, 0x9a, 0xb2, 0x61, 0xd9, 0x06, 0x17, 0x62, 0x42, 0x5f, 0xc3, 0x33, 0xdd,
	0xfd, 0xf4, 0xf8, 0xf3, 0xa1, 0x5a, 0xdf, 0x8d, 0xc1, 0xe5, 0xb7, 0xdf, 0xe1, 0x4e, 0xe9, 0x9c,
	0xf1, 0x86, 0x8b, 0x0d, 0x32, 0x53, 0xee, 0xee, 0x79, 0x55, 0x16, 0xd2, 0x21, 0x35, 0x93, 0x68,
	0xef, 0x95, 0x2e, 0x59, 0xc5, 0x65, 0xde, 0xf2, 0xbc, 0x6b, 0xf3, 0xae, 0xdd, 0xb7, 0xf3, 0x25,
	0xf9, 0x7c, 0x73, 0x64, 0x27, 0x9e, 0x99, 0x1c, 0x98, 0xc9, 0x61, 0x40, 0xb6, 0xf3, 0x5f, 0x67,
	0xb3, 0x55, 0xb9, 0xfb, 0xe4, 0x9f, 0xfd, 0xb0, 0x27, 0x2e, 0xdd, 0x68, 0x08, 0x55, 0x7d, 0xb9,
	0xe8, 0x86, 0x64, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x43, 0x33, 0x62, 0x52, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProfileTaskClient is the client API for ProfileTask service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileTaskClient interface {
	// query all sniffer need to execute profile task commands
	GetProfileTaskCommands(ctx context.Context, in *ProfileTaskCommandQuery, opts ...grpc.CallOption) (*v3.Commands, error)
	// collect dumped thread snapshot
	CollectSnapshot(ctx context.Context, opts ...grpc.CallOption) (ProfileTask_CollectSnapshotClient, error)
	// report profiling task finished
	ReportTaskFinish(ctx context.Context, in *ProfileTaskFinishReport, opts ...grpc.CallOption) (*v3.Commands, error)
}

type profileTaskClient struct {
	cc *grpc.ClientConn
}

func NewProfileTaskClient(cc *grpc.ClientConn) ProfileTaskClient {
	return &profileTaskClient{cc}
}

func (c *profileTaskClient) GetProfileTaskCommands(ctx context.Context, in *ProfileTaskCommandQuery, opts ...grpc.CallOption) (*v3.Commands, error) {
	out := new(v3.Commands)
	err := c.cc.Invoke(ctx, "/ProfileTask/getProfileTaskCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileTaskClient) CollectSnapshot(ctx context.Context, opts ...grpc.CallOption) (ProfileTask_CollectSnapshotClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProfileTask_serviceDesc.Streams[0], "/ProfileTask/collectSnapshot", opts...)
	if err != nil {
		return nil, err
	}
	x := &profileTaskCollectSnapshotClient{stream}
	return x, nil
}

type ProfileTask_CollectSnapshotClient interface {
	Send(*ThreadSnapshot) error
	CloseAndRecv() (*v3.Commands, error)
	grpc.ClientStream
}

type profileTaskCollectSnapshotClient struct {
	grpc.ClientStream
}

func (x *profileTaskCollectSnapshotClient) Send(m *ThreadSnapshot) error {
	return x.ClientStream.SendMsg(m)
}

func (x *profileTaskCollectSnapshotClient) CloseAndRecv() (*v3.Commands, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(v3.Commands)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *profileTaskClient) ReportTaskFinish(ctx context.Context, in *ProfileTaskFinishReport, opts ...grpc.CallOption) (*v3.Commands, error) {
	out := new(v3.Commands)
	err := c.cc.Invoke(ctx, "/ProfileTask/reportTaskFinish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileTaskServer is the server API for ProfileTask service.
type ProfileTaskServer interface {
	// query all sniffer need to execute profile task commands
	GetProfileTaskCommands(context.Context, *ProfileTaskCommandQuery) (*v3.Commands, error)
	// collect dumped thread snapshot
	CollectSnapshot(ProfileTask_CollectSnapshotServer) error
	// report profiling task finished
	ReportTaskFinish(context.Context, *ProfileTaskFinishReport) (*v3.Commands, error)
}

// UnimplementedProfileTaskServer can be embedded to have forward compatible implementations.
type UnimplementedProfileTaskServer struct {
}

func (*UnimplementedProfileTaskServer) GetProfileTaskCommands(ctx context.Context, req *ProfileTaskCommandQuery) (*v3.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileTaskCommands not implemented")
}
func (*UnimplementedProfileTaskServer) CollectSnapshot(srv ProfileTask_CollectSnapshotServer) error {
	return status.Errorf(codes.Unimplemented, "method CollectSnapshot not implemented")
}
func (*UnimplementedProfileTaskServer) ReportTaskFinish(ctx context.Context, req *ProfileTaskFinishReport) (*v3.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportTaskFinish not implemented")
}

func RegisterProfileTaskServer(s *grpc.Server, srv ProfileTaskServer) {
	s.RegisterService(&_ProfileTask_serviceDesc, srv)
}

func _ProfileTask_GetProfileTaskCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileTaskCommandQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileTaskServer).GetProfileTaskCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProfileTask/GetProfileTaskCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileTaskServer).GetProfileTaskCommands(ctx, req.(*ProfileTaskCommandQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileTask_CollectSnapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProfileTaskServer).CollectSnapshot(&profileTaskCollectSnapshotServer{stream})
}

type ProfileTask_CollectSnapshotServer interface {
	SendAndClose(*v3.Commands) error
	Recv() (*ThreadSnapshot, error)
	grpc.ServerStream
}

type profileTaskCollectSnapshotServer struct {
	grpc.ServerStream
}

func (x *profileTaskCollectSnapshotServer) SendAndClose(m *v3.Commands) error {
	return x.ServerStream.SendMsg(m)
}

func (x *profileTaskCollectSnapshotServer) Recv() (*ThreadSnapshot, error) {
	m := new(ThreadSnapshot)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProfileTask_ReportTaskFinish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileTaskFinishReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileTaskServer).ReportTaskFinish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProfileTask/ReportTaskFinish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileTaskServer).ReportTaskFinish(ctx, req.(*ProfileTaskFinishReport))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfileTask_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ProfileTask",
	HandlerType: (*ProfileTaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getProfileTaskCommands",
			Handler:    _ProfileTask_GetProfileTaskCommands_Handler,
		},
		{
			MethodName: "reportTaskFinish",
			Handler:    _ProfileTask_ReportTaskFinish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "collectSnapshot",
			Handler:       _ProfileTask_CollectSnapshot_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "profile/Profile.proto",
}
