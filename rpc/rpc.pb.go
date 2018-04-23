// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	Engine
	Empty
*/
package rpc

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

type Engine struct {
	Binary   string   `protobuf:"bytes,1,opt,name=binary" json:"binary,omitempty"`
	Replicas []string `protobuf:"bytes,2,rep,name=replicas" json:"replicas,omitempty"`
}

func (m *Engine) Reset()                    { *m = Engine{} }
func (m *Engine) String() string            { return proto.CompactTextString(m) }
func (*Engine) ProtoMessage()               {}
func (*Engine) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Engine) GetBinary() string {
	if m != nil {
		return m.Binary
	}
	return ""
}

func (m *Engine) GetReplicas() []string {
	if m != nil {
		return m.Replicas
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Engine)(nil), "Engine")
	proto.RegisterType((*Empty)(nil), "Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LonghornLauncherService service

type LonghornLauncherServiceClient interface {
	UpgradeEngine(ctx context.Context, in *Engine, opts ...grpc.CallOption) (*Empty, error)
}

type longhornLauncherServiceClient struct {
	cc *grpc.ClientConn
}

func NewLonghornLauncherServiceClient(cc *grpc.ClientConn) LonghornLauncherServiceClient {
	return &longhornLauncherServiceClient{cc}
}

func (c *longhornLauncherServiceClient) UpgradeEngine(ctx context.Context, in *Engine, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/LonghornLauncherService/UpgradeEngine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LonghornLauncherService service

type LonghornLauncherServiceServer interface {
	UpgradeEngine(context.Context, *Engine) (*Empty, error)
}

func RegisterLonghornLauncherServiceServer(s *grpc.Server, srv LonghornLauncherServiceServer) {
	s.RegisterService(&_LonghornLauncherService_serviceDesc, srv)
}

func _LonghornLauncherService_UpgradeEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Engine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LonghornLauncherServiceServer).UpgradeEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LonghornLauncherService/UpgradeEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LonghornLauncherServiceServer).UpgradeEngine(ctx, req.(*Engine))
	}
	return interceptor(ctx, in, info, handler)
}

var _LonghornLauncherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LonghornLauncherService",
	HandlerType: (*LonghornLauncherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpgradeEngine",
			Handler:    _LonghornLauncherService_UpgradeEngine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2a, 0x48, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2, 0xe1, 0x62, 0x73, 0xcd, 0x4b, 0xcf, 0xcc, 0x4b, 0x15,
	0x12, 0xe3, 0x62, 0x4b, 0xca, 0xcc, 0x4b, 0x2c, 0xaa, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x82, 0xf2, 0x84, 0xa4, 0xb8, 0x38, 0x8a, 0x52, 0x0b, 0x72, 0x32, 0x93, 0x13, 0x8b, 0x25, 0x98,
	0x14, 0x98, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x25, 0x76, 0x2e, 0x56, 0xd7, 0xdc, 0x82, 0x92, 0x4a,
	0x23, 0x5b, 0x2e, 0x71, 0x9f, 0xfc, 0xbc, 0xf4, 0x8c, 0xfc, 0xa2, 0x3c, 0x9f, 0xc4, 0xd2, 0xbc,
	0xe4, 0x8c, 0xd4, 0xa2, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x25, 0x2e, 0xde, 0xd0,
	0x82, 0xf4, 0xa2, 0xc4, 0x94, 0x54, 0xa8, 0x45, 0xec, 0x7a, 0x10, 0x86, 0x14, 0x9b, 0x1e, 0x58,
	0xb3, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x31, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0x7a,
	0xb0, 0xbc, 0x99, 0x00, 0x00, 0x00,
}