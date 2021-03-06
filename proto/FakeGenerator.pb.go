// Code generated by protoc-gen-go.
// source: FakeGenerator.proto
// DO NOT EDIT!

/*
Package fakegenerator is a generated protocol buffer package.

It is generated from these files:
	FakeGenerator.proto
	FakeGeneratorMessages.proto

It has these top-level messages:
	Address
	Date
	User
	Email
	EmptyMessage
*/
package fakegenerator

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

// Ignoring public import of Address from FakeGeneratorMessages.proto

// Ignoring public import of Date from FakeGeneratorMessages.proto

// Ignoring public import of User from FakeGeneratorMessages.proto

// Ignoring public import of Email from FakeGeneratorMessages.proto

// Ignoring public import of EmptyMessage from FakeGeneratorMessages.proto

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for FakeGenerator service

type FakeGeneratorClient interface {
	GetUser(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*User, error)
	GetAddress(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*Address, error)
	GetEmail(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*Email, error)
	GetDate(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*Date, error)
}

type fakeGeneratorClient struct {
	cc *grpc.ClientConn
}

func NewFakeGeneratorClient(cc *grpc.ClientConn) FakeGeneratorClient {
	return &fakeGeneratorClient{cc}
}

func (c *fakeGeneratorClient) GetUser(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/fakegenerator.FakeGenerator/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fakeGeneratorClient) GetAddress(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := grpc.Invoke(ctx, "/fakegenerator.FakeGenerator/GetAddress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fakeGeneratorClient) GetEmail(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*Email, error) {
	out := new(Email)
	err := grpc.Invoke(ctx, "/fakegenerator.FakeGenerator/GetEmail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fakeGeneratorClient) GetDate(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*Date, error) {
	out := new(Date)
	err := grpc.Invoke(ctx, "/fakegenerator.FakeGenerator/GetDate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FakeGenerator service

type FakeGeneratorServer interface {
	GetUser(context.Context, *EmptyMessage) (*User, error)
	GetAddress(context.Context, *EmptyMessage) (*Address, error)
	GetEmail(context.Context, *EmptyMessage) (*Email, error)
	GetDate(context.Context, *EmptyMessage) (*Date, error)
}

func RegisterFakeGeneratorServer(s *grpc.Server, srv FakeGeneratorServer) {
	s.RegisterService(&_FakeGenerator_serviceDesc, srv)
}

func _FakeGenerator_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FakeGeneratorServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fakegenerator.FakeGenerator/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FakeGeneratorServer).GetUser(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _FakeGenerator_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FakeGeneratorServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fakegenerator.FakeGenerator/GetAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FakeGeneratorServer).GetAddress(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _FakeGenerator_GetEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FakeGeneratorServer).GetEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fakegenerator.FakeGenerator/GetEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FakeGeneratorServer).GetEmail(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _FakeGenerator_GetDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FakeGeneratorServer).GetDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fakegenerator.FakeGenerator/GetDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FakeGeneratorServer).GetDate(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _FakeGenerator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fakegenerator.FakeGenerator",
	HandlerType: (*FakeGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _FakeGenerator_GetUser_Handler,
		},
		{
			MethodName: "GetAddress",
			Handler:    _FakeGenerator_GetAddress_Handler,
		},
		{
			MethodName: "GetEmail",
			Handler:    _FakeGenerator_GetEmail_Handler,
		},
		{
			MethodName: "GetDate",
			Handler:    _FakeGenerator_GetDate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("FakeGenerator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x76, 0x4b, 0xcc, 0x4e,
	0x75, 0x4f, 0xcd, 0x4b, 0x2d, 0x4a, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4d, 0x4b, 0xcc, 0x4e, 0x4d, 0x87, 0x09, 0x4a, 0x49, 0xa3, 0xa8, 0xf1, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0x86, 0xa8, 0x35, 0xea, 0x66, 0xe2, 0xe2, 0x45, 0x91, 0x17, 0xb2, 0xe6,
	0x62, 0x77, 0x4f, 0x2d, 0x09, 0x2d, 0x4e, 0x2d, 0x12, 0x92, 0xd6, 0x43, 0x31, 0x49, 0xcf, 0x35,
	0xb7, 0xa0, 0xa4, 0x12, 0x6a, 0x80, 0x94, 0x30, 0x9a, 0x24, 0x58, 0x87, 0x23, 0x17, 0x97, 0x7b,
	0x6a, 0x89, 0x63, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0x31, 0x7e, 0xfd, 0x62, 0x68, 0x92, 0x30, 0x4d,
	0xb6, 0x5c, 0x1c, 0xee, 0xa9, 0x25, 0xae, 0xb9, 0x89, 0x99, 0x39, 0xf8, 0x0d, 0x10, 0xc1, 0x90,
	0x04, 0x69, 0x81, 0x38, 0xdf, 0x25, 0xb1, 0x24, 0x95, 0x34, 0xe7, 0x83, 0x74, 0x04, 0x30, 0x24,
	0xb1, 0x81, 0x83, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x49, 0x20, 0x7c, 0x59, 0x01,
	0x00, 0x00,
}
