// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/echo.proto

/*
Package Echo is a generated protocol buffer package.

It is generated from these files:
	protos/echo.proto

It has these top-level messages:
	EchoRequest
	EchoResponse
*/
package Echo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EchoRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EchoRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EchoResponse struct {
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *EchoResponse) Reset()                    { *m = EchoResponse{} }
func (m *EchoResponse) String() string            { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()               {}
func (*EchoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EchoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "Echo.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "Echo.EchoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Echo service

type EchoClient interface {
	Echo(ctx context.Context, opts ...grpc.CallOption) (Echo_EchoClient, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Echo(ctx context.Context, opts ...grpc.CallOption) (Echo_EchoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Echo_serviceDesc.Streams[0], c.cc, "/Echo.Echo/Echo", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoEchoClient{stream}
	return x, nil
}

type Echo_EchoClient interface {
	Send(*EchoRequest) error
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoEchoClient struct {
	grpc.ClientStream
}

func (x *echoEchoClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoEchoClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Echo service

type EchoServer interface {
	Echo(Echo_EchoServer) error
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Echo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServer).Echo(&echoEchoServer{stream})
}

type Echo_EchoServer interface {
	Send(*EchoResponse) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoEchoServer struct {
	grpc.ServerStream
}

func (x *echoEchoServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoEchoServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Echo",
			Handler:       _Echo_Echo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/echo.proto",
}

func init() { proto.RegisterFile("protos/echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0x4d, 0xce, 0xc8, 0xd7, 0x03, 0xb3, 0x85, 0x58, 0x5c, 0x93, 0x33, 0xf2,
	0x95, 0xd4, 0xb9, 0xb8, 0x41, 0x74, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x04, 0x17,
	0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c,
	0xab, 0xa4, 0xc1, 0xc5, 0x03, 0x51, 0x58, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x8a, 0xac, 0x92, 0x09,
	0x45, 0xa5, 0x91, 0x35, 0x17, 0xd8, 0x68, 0x21, 0x63, 0x28, 0x2d, 0xa8, 0x07, 0xa2, 0xf4, 0x90,
	0xac, 0x91, 0x12, 0x42, 0x16, 0x82, 0x18, 0xa8, 0xc4, 0xa0, 0xc1, 0x68, 0xc0, 0x98, 0xc4, 0x06,
	0x76, 0x9c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xda, 0x69, 0xed, 0xb1, 0x00, 0x00, 0x00,
}
