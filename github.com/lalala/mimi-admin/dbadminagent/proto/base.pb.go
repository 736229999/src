// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

package dbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Nil struct {
}

func (m *Nil) Reset()                    { *m = Nil{} }
func (m *Nil) String() string            { return proto.CompactTextString(m) }
func (*Nil) ProtoMessage()               {}
func (*Nil) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type BoolValue struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value"`
}

func (m *BoolValue) Reset()                    { *m = BoolValue{} }
func (m *BoolValue) String() string            { return proto.CompactTextString(m) }
func (*BoolValue) ProtoMessage()               {}
func (*BoolValue) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *BoolValue) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type IntValue struct {
	Value int64 `protobuf:"varint,1,opt,name=value" json:"value"`
}

func (m *IntValue) Reset()                    { *m = IntValue{} }
func (m *IntValue) String() string            { return proto.CompactTextString(m) }
func (*IntValue) ProtoMessage()               {}
func (*IntValue) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *IntValue) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type StringValue struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value"`
}

func (m *StringValue) Reset()                    { *m = StringValue{} }
func (m *StringValue) String() string            { return proto.CompactTextString(m) }
func (*StringValue) ProtoMessage()               {}
func (*StringValue) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *StringValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Nil)(nil), "dbproto.Nil")
	proto.RegisterType((*BoolValue)(nil), "dbproto.BoolValue")
	proto.RegisterType((*IntValue)(nil), "dbproto.IntValue")
	proto.RegisterType((*StringValue)(nil), "dbproto.StringValue")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x49, 0x02, 0x33, 0x94, 0x58, 0xb9, 0x98,
	0xfd, 0x32, 0x73, 0x94, 0x14, 0xb9, 0x38, 0x9d, 0xf2, 0xf3, 0x73, 0xc2, 0x12, 0x73, 0x4a, 0x53,
	0x85, 0x44, 0xb8, 0x58, 0xcb, 0x40, 0x0c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x08, 0x47,
	0x49, 0x81, 0x8b, 0xc3, 0x33, 0xaf, 0x04, 0x8b, 0x0a, 0x66, 0x98, 0x0a, 0x65, 0x2e, 0xee, 0xe0,
	0x92, 0xa2, 0xcc, 0xbc, 0x74, 0x2c, 0x8a, 0x38, 0xa1, 0x8a, 0x92, 0xd8, 0xc0, 0xf6, 0x1a, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x08, 0x37, 0xae, 0xc4, 0x8e, 0x00, 0x00, 0x00,
}