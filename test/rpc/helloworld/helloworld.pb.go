// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/rpc/helloworld/helloworld.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	test/rpc/helloworld/helloworld.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package helloworld

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

func init() { proto.RegisterFile("test/rpc/helloworld/helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x49, 0x2d, 0x2e,
	0xd1, 0x2f, 0x2a, 0x48, 0xd6, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0x41,
	0x62, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x71, 0x21, 0x44, 0x94, 0x94, 0xb8, 0x78, 0x3c,
	0x40, 0xbc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc,
	0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x8d, 0x8b, 0x0b, 0xaa, 0xa6,
	0x20, 0xa7, 0x52, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x1d, 0xa6, 0x08, 0xc6,
	0x35, 0xf2, 0xe4, 0x62, 0x77, 0x2f, 0x4a, 0x4d, 0x2d, 0x49, 0x2d, 0x12, 0xb2, 0xe3, 0xe2, 0x08,
	0x4e, 0xac, 0x04, 0xeb, 0x12, 0x92, 0xd0, 0x43, 0x72, 0x01, 0xb2, 0x65, 0x52, 0x62, 0x58, 0x64,
	0x0a, 0x72, 0x2a, 0x95, 0x18, 0x9c, 0xd8, 0x16, 0x31, 0x31, 0x7b, 0xf8, 0x84, 0x27, 0xb1, 0x81,
	0x5d, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x6f, 0x82, 0x7d, 0xd9, 0x00, 0x00, 0x00,
}
