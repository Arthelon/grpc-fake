// Code generated by protoc-gen-go.
// source: FakeGenerator.proto
// DO NOT EDIT!

/*
Package fakegenerator is a generated protocol buffer package.

It is generated from these files:
	FakeGenerator.proto

It has these top-level messages:
*/
package fakegenerator

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("FakeGenerator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x76, 0x4b, 0xcc, 0x4e,
	0x75, 0x4f, 0xcd, 0x4b, 0x2d, 0x4a, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4d, 0x4b, 0xcc, 0x4e, 0x4d, 0x87, 0x09, 0x4a, 0x49, 0xa3, 0xa8, 0xf1, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0x86, 0xa8, 0x35, 0x9a, 0xc5, 0xc8, 0xc5, 0x8b, 0x22, 0x2f, 0x24, 0xc7,
	0xc5, 0xe6, 0x9e, 0x5a, 0x12, 0x5a, 0x9c, 0x2a, 0xc4, 0xab, 0xe7, 0x9a, 0x5b, 0x50, 0x52, 0x09,
	0xd5, 0x21, 0xc5, 0xaa, 0x17, 0x5a, 0x9c, 0x5a, 0x24, 0xa4, 0xca, 0xc5, 0xe5, 0x9e, 0x5a, 0xe2,
	0x98, 0x92, 0x52, 0x94, 0x5a, 0x5c, 0x8c, 0xae, 0x86, 0x43, 0x0f, 0x26, 0xa1, 0xc8, 0xc5, 0xe1,
	0x9e, 0x5a, 0xe2, 0x9a, 0x9b, 0x98, 0x99, 0x83, 0xae, 0x88, 0x4d, 0x0f, 0x22, 0x2c, 0xcf, 0xc5,
	0xee, 0x9e, 0x5a, 0xe2, 0x92, 0x58, 0x82, 0xc5, 0x2a, 0x90, 0x68, 0x12, 0x1b, 0xd8, 0x8d, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x30, 0x94, 0x35, 0xe6, 0x00, 0x00, 0x00,
}