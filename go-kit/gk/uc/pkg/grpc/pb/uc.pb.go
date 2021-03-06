// Code generated by protoc-gen-go. DO NOT EDIT.
// source: uc.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	uc.proto

It has these top-level messages:
	GDeviceIDRequest
	GDeviceIDReply
*/
package pb

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

type GDeviceIDRequest struct {
}

func (m *GDeviceIDRequest) Reset()                    { *m = GDeviceIDRequest{} }
func (m *GDeviceIDRequest) String() string            { return proto.CompactTextString(m) }
func (*GDeviceIDRequest) ProtoMessage()               {}
func (*GDeviceIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GDeviceIDReply struct {
}

func (m *GDeviceIDReply) Reset()                    { *m = GDeviceIDReply{} }
func (m *GDeviceIDReply) String() string            { return proto.CompactTextString(m) }
func (*GDeviceIDReply) ProtoMessage()               {}
func (*GDeviceIDReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*GDeviceIDRequest)(nil), "pb.GDeviceIDRequest")
	proto.RegisterType((*GDeviceIDReply)(nil), "pb.GDeviceIDReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Uc service

type UcClient interface {
	GDeviceID(ctx context.Context, in *GDeviceIDRequest, opts ...grpc.CallOption) (*GDeviceIDReply, error)
}

type ucClient struct {
	cc *grpc.ClientConn
}

func NewUcClient(cc *grpc.ClientConn) UcClient {
	return &ucClient{cc}
}

func (c *ucClient) GDeviceID(ctx context.Context, in *GDeviceIDRequest, opts ...grpc.CallOption) (*GDeviceIDReply, error) {
	out := new(GDeviceIDReply)
	err := grpc.Invoke(ctx, "/pb.Uc/GDeviceID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Uc service

type UcServer interface {
	GDeviceID(context.Context, *GDeviceIDRequest) (*GDeviceIDReply, error)
}

func RegisterUcServer(s *grpc.Server, srv UcServer) {
	s.RegisterService(&_Uc_serviceDesc, srv)
}

func _Uc_GDeviceID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GDeviceIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UcServer).GDeviceID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Uc/GDeviceID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UcServer).GDeviceID(ctx, req.(*GDeviceIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Uc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Uc",
	HandlerType: (*UcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GDeviceID",
			Handler:    _Uc_GDeviceID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uc.proto",
}

func init() { proto.RegisterFile("uc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 101 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x4d, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x12, 0xe2, 0x12, 0x70, 0x77, 0x49, 0x2d,
	0xcb, 0x4c, 0x4e, 0xf5, 0x74, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x12, 0xe0, 0xe2,
	0x43, 0x12, 0x2b, 0xc8, 0xa9, 0x34, 0xb2, 0xe5, 0x62, 0x0a, 0x4d, 0x16, 0x32, 0xe7, 0xe2, 0x84,
	0x8b, 0x0b, 0x89, 0xe8, 0x15, 0x24, 0xe9, 0xa1, 0x6b, 0x95, 0x12, 0x42, 0x13, 0x2d, 0xc8, 0xa9,
	0x54, 0x62, 0x48, 0x62, 0x03, 0xdb, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x95, 0x05, 0x1a,
	0x89, 0x7b, 0x00, 0x00, 0x00,
}
