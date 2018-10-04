// Code generated by protoc-gen-go. DO NOT EDIT.
// source: key_value.proto

package keyvalue

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

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_key_value_6105a7c652e4cf0c, []int{0}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type GetResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_key_value_6105a7c652e4cf0c, []int{1}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetRequest)(nil), "keyvalue.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "keyvalue.GetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeyValueClient is the client API for KeyValue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyValueClient interface {
	Foo(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Bar(ctx context.Context, opts ...grpc.CallOption) (KeyValue_BarClient, error)
	Baz(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (KeyValue_BazClient, error)
	Qux(ctx context.Context, opts ...grpc.CallOption) (KeyValue_QuxClient, error)
}

type keyValueClient struct {
	cc *grpc.ClientConn
}

func NewKeyValueClient(cc *grpc.ClientConn) KeyValueClient {
	return &keyValueClient{cc}
}

func (c *keyValueClient) Foo(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/keyvalue.KeyValue/Foo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueClient) Bar(ctx context.Context, opts ...grpc.CallOption) (KeyValue_BarClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KeyValue_serviceDesc.Streams[0], "/keyvalue.KeyValue/Bar", opts...)
	if err != nil {
		return nil, err
	}
	x := &keyValueBarClient{stream}
	return x, nil
}

type KeyValue_BarClient interface {
	Send(*GetRequest) error
	CloseAndRecv() (*GetResponse, error)
	grpc.ClientStream
}

type keyValueBarClient struct {
	grpc.ClientStream
}

func (x *keyValueBarClient) Send(m *GetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *keyValueBarClient) CloseAndRecv() (*GetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *keyValueClient) Baz(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (KeyValue_BazClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KeyValue_serviceDesc.Streams[1], "/keyvalue.KeyValue/Baz", opts...)
	if err != nil {
		return nil, err
	}
	x := &keyValueBazClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KeyValue_BazClient interface {
	Recv() (*GetResponse, error)
	grpc.ClientStream
}

type keyValueBazClient struct {
	grpc.ClientStream
}

func (x *keyValueBazClient) Recv() (*GetResponse, error) {
	m := new(GetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *keyValueClient) Qux(ctx context.Context, opts ...grpc.CallOption) (KeyValue_QuxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KeyValue_serviceDesc.Streams[2], "/keyvalue.KeyValue/Qux", opts...)
	if err != nil {
		return nil, err
	}
	x := &keyValueQuxClient{stream}
	return x, nil
}

type KeyValue_QuxClient interface {
	Send(*GetRequest) error
	Recv() (*GetResponse, error)
	grpc.ClientStream
}

type keyValueQuxClient struct {
	grpc.ClientStream
}

func (x *keyValueQuxClient) Send(m *GetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *keyValueQuxClient) Recv() (*GetResponse, error) {
	m := new(GetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KeyValueServer is the server API for KeyValue service.
type KeyValueServer interface {
	Foo(context.Context, *GetRequest) (*GetResponse, error)
	Bar(KeyValue_BarServer) error
	Baz(*GetRequest, KeyValue_BazServer) error
	Qux(KeyValue_QuxServer) error
}

func RegisterKeyValueServer(s *grpc.Server, srv KeyValueServer) {
	s.RegisterService(&_KeyValue_serviceDesc, srv)
}

func _KeyValue_Foo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServer).Foo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyvalue.KeyValue/Foo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServer).Foo(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValue_Bar_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KeyValueServer).Bar(&keyValueBarServer{stream})
}

type KeyValue_BarServer interface {
	SendAndClose(*GetResponse) error
	Recv() (*GetRequest, error)
	grpc.ServerStream
}

type keyValueBarServer struct {
	grpc.ServerStream
}

func (x *keyValueBarServer) SendAndClose(m *GetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *keyValueBarServer) Recv() (*GetRequest, error) {
	m := new(GetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _KeyValue_Baz_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KeyValueServer).Baz(m, &keyValueBazServer{stream})
}

type KeyValue_BazServer interface {
	Send(*GetResponse) error
	grpc.ServerStream
}

type keyValueBazServer struct {
	grpc.ServerStream
}

func (x *keyValueBazServer) Send(m *GetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KeyValue_Qux_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KeyValueServer).Qux(&keyValueQuxServer{stream})
}

type KeyValue_QuxServer interface {
	Send(*GetResponse) error
	Recv() (*GetRequest, error)
	grpc.ServerStream
}

type keyValueQuxServer struct {
	grpc.ServerStream
}

func (x *keyValueQuxServer) Send(m *GetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *keyValueQuxServer) Recv() (*GetRequest, error) {
	m := new(GetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _KeyValue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "keyvalue.KeyValue",
	HandlerType: (*KeyValueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Foo",
			Handler:    _KeyValue_Foo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Bar",
			Handler:       _KeyValue_Bar_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Baz",
			Handler:       _KeyValue_Baz_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Qux",
			Handler:       _KeyValue_Qux_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "key_value.proto",
}

func init() { proto.RegisterFile("key_value.proto", fileDescriptor_key_value_6105a7c652e4cf0c) }

var fileDescriptor_key_value_6105a7c652e4cf0c = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4e, 0xad, 0x8c,
	0x2f, 0x4b, 0xcc, 0x29, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x4e, 0xad,
	0x04, 0xf3, 0x95, 0x78, 0xb8, 0xb8, 0xdc, 0x53, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x94, 0x78, 0xb9, 0xb8, 0xc1, 0xbc, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xa3, 0x47, 0x8c, 0x5c,
	0x1c, 0xde, 0xa9, 0x95, 0x61, 0x20, 0x95, 0x42, 0x46, 0x5c, 0xcc, 0x6e, 0xf9, 0xf9, 0x42, 0x22,
	0x7a, 0x30, 0xbd, 0x7a, 0x08, 0x8d, 0x52, 0xa2, 0x68, 0xa2, 0x10, 0x03, 0x84, 0x4c, 0xb8, 0x98,
	0x9d, 0x12, 0x8b, 0x48, 0xd2, 0xa3, 0xc1, 0x08, 0xd1, 0x55, 0x45, 0x92, 0x2e, 0x03, 0x46, 0x21,
	0x33, 0x2e, 0xe6, 0xc0, 0xd2, 0x0a, 0x12, 0xed, 0x32, 0x60, 0x4c, 0x62, 0x03, 0x07, 0x89, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xd7, 0x65, 0x96, 0x25, 0x01, 0x00, 0x00,
}