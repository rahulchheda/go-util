// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/litmuschaos/go-util/internal/proto-files/service/grpc-service.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	domain "github.com/litmuschaos/go-util/internal/grpc/domain"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Status struct {
	CompletedStatus      bool     `protobuf:"varint,1,opt,name=completedStatus,proto3" json:"completedStatus,omitempty"`
	AnyError             *Error   `protobuf:"bytes,2,opt,name=anyError,proto3" json:"anyError,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fdb37e172298772, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCompletedStatus() bool {
	if m != nil {
		return m.CompletedStatus
	}
	return false
}

func (m *Status) GetAnyError() *Error {
	if m != nil {
		return m.AnyError
	}
	return nil
}

type Error struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fdb37e172298772, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Status)(nil), "service.Status")
	proto.RegisterType((*Error)(nil), "service.Error")
}

func init() {
	proto.RegisterFile("github.com/litmuschaos/go-util/internal/proto-files/service/grpc-service.proto", fileDescriptor_3fdb37e172298772)
}

var fileDescriptor_3fdb37e172298772 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xbf, 0x4b, 0xc4, 0x30,
	0x1c, 0xc5, 0xa9, 0xe8, 0xf5, 0x2e, 0x8a, 0x85, 0x4c, 0xe5, 0xa6, 0xe3, 0xa6, 0x22, 0xb4, 0x81,
	0xf3, 0xc7, 0xe2, 0x26, 0xb8, 0x89, 0x43, 0x6f, 0x73, 0x10, 0x72, 0xe9, 0xd7, 0x5e, 0x20, 0xc9,
	0xb7, 0xe4, 0x87, 0xe8, 0x7f, 0x2f, 0x26, 0x6d, 0x07, 0x27, 0xb9, 0xed, 0xbd, 0x97, 0xf0, 0xc9,
	0xcb, 0x23, 0xaf, 0xbd, 0xf4, 0xc7, 0x70, 0x68, 0x04, 0x6a, 0xa6, 0xa4, 0xd7, 0xc1, 0x89, 0x23,
	0x47, 0xc7, 0x7a, 0xac, 0x83, 0x97, 0x8a, 0x49, 0xe3, 0xc1, 0x1a, 0xae, 0xd8, 0x60, 0xd1, 0x63,
	0xfd, 0x21, 0x15, 0x38, 0xe6, 0xc0, 0x7e, 0x4a, 0x01, 0xac, 0xb7, 0x83, 0xa8, 0x47, 0xd3, 0xc4,
	0x0b, 0x34, 0x1f, 0xed, 0xfa, 0xe5, 0x14, 0x70, 0x87, 0x9a, 0x4b, 0x93, 0xb8, 0x49, 0x27, 0xec,
	0xf6, 0x9d, 0x2c, 0xf6, 0x9e, 0xfb, 0xe0, 0x68, 0x45, 0x0a, 0x81, 0x7a, 0x50, 0xe0, 0xa1, 0x4b,
	0x51, 0x99, 0x6d, 0xb2, 0x6a, 0xd9, 0xfe, 0x8d, 0xe9, 0x0d, 0x59, 0x72, 0xf3, 0xfd, 0x6c, 0x2d,
	0xda, 0xf2, 0x6c, 0x93, 0x55, 0x97, 0xbb, 0xeb, 0x66, 0x2a, 0x1b, 0xd3, 0x76, 0x3e, 0xdf, 0xde,
	0x93, 0x8b, 0x28, 0x28, 0x25, 0xe7, 0x02, 0x3b, 0x88, 0xcc, 0x55, 0x1b, 0x35, 0x2d, 0x49, 0xae,
	0xc1, 0x39, 0xde, 0x43, 0xe4, 0xac, 0xda, 0xc9, 0xee, 0x1e, 0x49, 0x01, 0x5f, 0x20, 0x82, 0x47,
	0xbb, 0x4f, 0x64, 0x5a, 0x91, 0x3c, 0x45, 0x40, 0xaf, 0x9a, 0xf1, 0x0f, 0x86, 0x6b, 0x58, 0x17,
	0xf3, 0xe3, 0xa9, 0xdf, 0xd3, 0xc3, 0xdb, 0xdd, 0x7f, 0x37, 0xfa, 0x1d, 0x64, 0x5a, 0xfd, 0xb0,
	0x88, 0x93, 0xdc, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x58, 0xe7, 0x06, 0xbb, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExecutorServiceClient is the client API for ExecutorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecutorServiceClient interface {
	Execute(ctx context.Context, in *domain.Name, opts ...grpc.CallOption) (*Status, error)
}

type executorServiceClient struct {
	cc *grpc.ClientConn
}

func NewExecutorServiceClient(cc *grpc.ClientConn) ExecutorServiceClient {
	return &executorServiceClient{cc}
}

func (c *executorServiceClient) Execute(ctx context.Context, in *domain.Name, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/service.executorService/execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutorServiceServer is the server API for ExecutorService service.
type ExecutorServiceServer interface {
	Execute(context.Context, *domain.Name) (*Status, error)
}

// UnimplementedExecutorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedExecutorServiceServer struct {
}

func (*UnimplementedExecutorServiceServer) Execute(ctx context.Context, req *domain.Name) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

func RegisterExecutorServiceServer(s *grpc.Server, srv ExecutorServiceServer) {
	s.RegisterService(&_ExecutorService_serviceDesc, srv)
}

func _ExecutorService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.executorService/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServiceServer).Execute(ctx, req.(*domain.Name))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExecutorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.executorService",
	HandlerType: (*ExecutorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "execute",
			Handler:    _ExecutorService_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/litmuschaos/go-util/internal/proto-files/service/grpc-service.proto",
}
