// Code generated by protoc-gen-go.
// source: mysqlctl.proto
// DO NOT EDIT!

/*
Package mysqlctl is a generated protocol buffer package.

It is generated from these files:
	mysqlctl.proto

It has these top-level messages:
	StartRequest
	StartResponse
	ShutdownRequest
	ShutdownResponse
	RunMysqlUpgradeRequest
	RunMysqlUpgradeResponse
*/
package mysqlctl

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

type StartRequest struct {
}

func (m *StartRequest) Reset()                    { *m = StartRequest{} }
func (m *StartRequest) String() string            { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()               {}
func (*StartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StartResponse struct {
}

func (m *StartResponse) Reset()                    { *m = StartResponse{} }
func (m *StartResponse) String() string            { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()               {}
func (*StartResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ShutdownRequest struct {
	WaitForMysqld bool `protobuf:"varint,1,opt,name=wait_for_mysqld,json=waitForMysqld" json:"wait_for_mysqld,omitempty"`
}

func (m *ShutdownRequest) Reset()                    { *m = ShutdownRequest{} }
func (m *ShutdownRequest) String() string            { return proto.CompactTextString(m) }
func (*ShutdownRequest) ProtoMessage()               {}
func (*ShutdownRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ShutdownResponse struct {
}

func (m *ShutdownResponse) Reset()                    { *m = ShutdownResponse{} }
func (m *ShutdownResponse) String() string            { return proto.CompactTextString(m) }
func (*ShutdownResponse) ProtoMessage()               {}
func (*ShutdownResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type RunMysqlUpgradeRequest struct {
}

func (m *RunMysqlUpgradeRequest) Reset()                    { *m = RunMysqlUpgradeRequest{} }
func (m *RunMysqlUpgradeRequest) String() string            { return proto.CompactTextString(m) }
func (*RunMysqlUpgradeRequest) ProtoMessage()               {}
func (*RunMysqlUpgradeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type RunMysqlUpgradeResponse struct {
}

func (m *RunMysqlUpgradeResponse) Reset()                    { *m = RunMysqlUpgradeResponse{} }
func (m *RunMysqlUpgradeResponse) String() string            { return proto.CompactTextString(m) }
func (*RunMysqlUpgradeResponse) ProtoMessage()               {}
func (*RunMysqlUpgradeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*StartRequest)(nil), "mysqlctl.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "mysqlctl.StartResponse")
	proto.RegisterType((*ShutdownRequest)(nil), "mysqlctl.ShutdownRequest")
	proto.RegisterType((*ShutdownResponse)(nil), "mysqlctl.ShutdownResponse")
	proto.RegisterType((*RunMysqlUpgradeRequest)(nil), "mysqlctl.RunMysqlUpgradeRequest")
	proto.RegisterType((*RunMysqlUpgradeResponse)(nil), "mysqlctl.RunMysqlUpgradeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for MysqlCtl service

type MysqlCtlClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
	RunMysqlUpgrade(ctx context.Context, in *RunMysqlUpgradeRequest, opts ...grpc.CallOption) (*RunMysqlUpgradeResponse, error)
}

type mysqlCtlClient struct {
	cc *grpc.ClientConn
}

func NewMysqlCtlClient(cc *grpc.ClientConn) MysqlCtlClient {
	return &mysqlCtlClient{cc}
}

func (c *mysqlCtlClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := grpc.Invoke(ctx, "/mysqlctl.MysqlCtl/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mysqlCtlClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := grpc.Invoke(ctx, "/mysqlctl.MysqlCtl/Shutdown", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mysqlCtlClient) RunMysqlUpgrade(ctx context.Context, in *RunMysqlUpgradeRequest, opts ...grpc.CallOption) (*RunMysqlUpgradeResponse, error) {
	out := new(RunMysqlUpgradeResponse)
	err := grpc.Invoke(ctx, "/mysqlctl.MysqlCtl/RunMysqlUpgrade", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MysqlCtl service

type MysqlCtlServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
	RunMysqlUpgrade(context.Context, *RunMysqlUpgradeRequest) (*RunMysqlUpgradeResponse, error)
}

func RegisterMysqlCtlServer(s *grpc.Server, srv MysqlCtlServer) {
	s.RegisterService(&_MysqlCtl_serviceDesc, srv)
}

func _MysqlCtl_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlCtlServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mysqlctl.MysqlCtl/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlCtlServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MysqlCtl_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlCtlServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mysqlctl.MysqlCtl/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlCtlServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MysqlCtl_RunMysqlUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunMysqlUpgradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlCtlServer).RunMysqlUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mysqlctl.MysqlCtl/RunMysqlUpgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlCtlServer).RunMysqlUpgrade(ctx, req.(*RunMysqlUpgradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MysqlCtl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mysqlctl.MysqlCtl",
	HandlerType: (*MysqlCtlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _MysqlCtl_Start_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _MysqlCtl_Shutdown_Handler,
		},
		{
			MethodName: "RunMysqlUpgrade",
			Handler:    _MysqlCtl_RunMysqlUpgrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("mysqlctl.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xad, 0x2c, 0x2e,
	0xcc, 0x49, 0x2e, 0xc9, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xf8,
	0xb8, 0x78, 0x82, 0x4b, 0x12, 0x8b, 0x4a, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0xf8,
	0xb9, 0x78, 0xa1, 0xfc, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x25, 0x4b, 0x2e, 0xfe, 0xe0, 0x8c,
	0xd2, 0x92, 0x94, 0xfc, 0xf2, 0x3c, 0xa8, 0x1a, 0x21, 0x35, 0x2e, 0xfe, 0xf2, 0xc4, 0xcc, 0x92,
	0xf8, 0xb4, 0xfc, 0xa2, 0x78, 0xb0, 0x41, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0xbc,
	0x20, 0x61, 0xb7, 0xfc, 0x22, 0x5f, 0xb0, 0xa0, 0x92, 0x10, 0x97, 0x00, 0x42, 0x2b, 0xd4, 0x38,
	0x09, 0x2e, 0xb1, 0xa0, 0xd2, 0x3c, 0xb0, 0x82, 0xd0, 0x82, 0xf4, 0xa2, 0xc4, 0x94, 0x54, 0x98,
	0xcd, 0x92, 0x5c, 0xe2, 0x18, 0x32, 0x10, 0x4d, 0x46, 0x4f, 0x19, 0xb9, 0x38, 0xc0, 0x12, 0xce,
	0x25, 0x39, 0x42, 0x56, 0x5c, 0xac, 0x60, 0x17, 0x0a, 0x89, 0xe9, 0xc1, 0x7d, 0x85, 0xec, 0x05,
	0x29, 0x71, 0x0c, 0x71, 0xa8, 0xdd, 0x0c, 0x42, 0xce, 0x5c, 0x1c, 0x30, 0x17, 0x09, 0x49, 0x22,
	0x29, 0x43, 0xf5, 0xa0, 0x94, 0x14, 0x36, 0x29, 0xb8, 0x21, 0x11, 0x5c, 0xfc, 0x68, 0x0e, 0x15,
	0x52, 0x40, 0x68, 0xc0, 0xee, 0x3b, 0x29, 0x45, 0x3c, 0x2a, 0x60, 0x26, 0x27, 0xb1, 0x81, 0x63,
	0xc7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x7f, 0x74, 0x52, 0xaf, 0x01, 0x00, 0x00,
}
