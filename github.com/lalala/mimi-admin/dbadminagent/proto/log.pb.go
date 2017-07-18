// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log.proto

package dbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Log struct {
	Id         int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	Username   string `protobuf:"bytes,2,opt,name=username" json:"username"`
	Path       string `protobuf:"bytes,3,opt,name=path" json:"path"`
	Operating  int64  `protobuf:"varint,4,opt,name=operating" json:"operating"`
	Params     string `protobuf:"bytes,5,opt,name=params" json:"params"`
	Message    string `protobuf:"bytes,6,opt,name=message" json:"message"`
	CreateTime int64  `protobuf:"varint,7,opt,name=create_time,json=createTime" json:"create_time"`
	UserId     int64  `protobuf:"varint,8,opt,name=user_id,json=userId" json:"user_id"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{0} }

func (m *Log) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Log) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Log) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Log) GetOperating() int64 {
	if m != nil {
		return m.Operating
	}
	return 0
}

func (m *Log) GetParams() string {
	if m != nil {
		return m.Params
	}
	return ""
}

func (m *Log) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Log) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Log) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type LogReply struct {
	Account       string `protobuf:"bytes,1,opt,name=account" json:"account"`
	Log           []*Log `protobuf:"bytes,2,rep,name=log" json:"log"`
	Page          int64  `protobuf:"varint,3,opt,name=page" json:"page"`
	Size          int64  `protobuf:"varint,4,opt,name=size" json:"size"`
	Total         int64  `protobuf:"varint,5,opt,name=total" json:"total"`
	Authorization bool   `protobuf:"varint,6,opt,name=Authorization" json:"Authorization"`
	UserId        int64  `protobuf:"varint,7,opt,name=user_id,json=userId" json:"user_id"`
}

func (m *LogReply) Reset()                    { *m = LogReply{} }
func (m *LogReply) String() string            { return proto.CompactTextString(m) }
func (*LogReply) ProtoMessage()               {}
func (*LogReply) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{1} }

func (m *LogReply) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *LogReply) GetLog() []*Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *LogReply) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *LogReply) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *LogReply) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *LogReply) GetAuthorization() bool {
	if m != nil {
		return m.Authorization
	}
	return false
}

func (m *LogReply) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*Log)(nil), "dbproto.Log")
	proto.RegisterType((*LogReply)(nil), "dbproto.LogReply")
}

func init() { proto.RegisterFile("log.proto", fileDescriptor17) }

var fileDescriptor17 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4b, 0x4e, 0xf3, 0x30,
	0x10, 0xc7, 0x95, 0xb8, 0xcd, 0x63, 0xfa, 0x7d, 0x2c, 0x46, 0x08, 0x2c, 0x84, 0xa0, 0xaa, 0x58,
	0x74, 0xd5, 0x05, 0x9c, 0x80, 0x25, 0x52, 0x57, 0x16, 0xfb, 0xca, 0x6d, 0x46, 0xae, 0xa5, 0x24,
	0x8e, 0x1c, 0x67, 0x41, 0x6f, 0xc7, 0x09, 0xb8, 0x12, 0xf2, 0x24, 0x85, 0xb2, 0xfb, 0x3f, 0x34,
	0xe3, 0xf9, 0x19, 0xca, 0xda, 0x99, 0x4d, 0xe7, 0x5d, 0x70, 0x98, 0x57, 0x7b, 0x16, 0xab, 0xaf,
	0x04, 0xc4, 0xd6, 0x19, 0xbc, 0x82, 0xd4, 0x56, 0x32, 0x59, 0x26, 0x6b, 0xa1, 0x52, 0x5b, 0xe1,
	0x1d, 0x14, 0x43, 0x4f, 0xbe, 0xd5, 0x0d, 0xc9, 0x74, 0x99, 0xac, 0x4b, 0xf5, 0xe3, 0x11, 0x61,
	0xd6, 0xe9, 0x70, 0x94, 0x82, 0x73, 0xd6, 0x78, 0x0f, 0xa5, 0xeb, 0xc8, 0xeb, 0x60, 0x5b, 0x23,
	0x67, 0xbc, 0xe6, 0x37, 0xc0, 0x1b, 0xc8, 0x3a, 0xed, 0x75, 0xd3, 0xcb, 0x39, 0xcf, 0x4c, 0x0e,
	0x25, 0xe4, 0x0d, 0xf5, 0xbd, 0x36, 0x24, 0x33, 0x2e, 0xce, 0x16, 0x1f, 0x61, 0x71, 0xf0, 0xa4,
	0x03, 0xed, 0x82, 0x6d, 0x48, 0xe6, 0xbc, 0x11, 0xc6, 0xe8, 0xdd, 0x36, 0x84, 0xb7, 0x90, 0xc7,
	0x83, 0x76, 0xb6, 0x92, 0x05, 0x97, 0x59, 0xb4, 0x6f, 0xd5, 0xea, 0x33, 0x81, 0x62, 0xeb, 0x8c,
	0xa2, 0xae, 0xfe, 0x88, 0x0f, 0xe8, 0xc3, 0xc1, 0x0d, 0x6d, 0x60, 0xb6, 0x52, 0x9d, 0x2d, 0x3e,
	0x80, 0xa8, 0x9d, 0x91, 0xe9, 0x52, 0xac, 0x17, 0xcf, 0xff, 0x36, 0xd3, 0x7f, 0x6c, 0xe2, 0x64,
	0x2c, 0x46, 0x48, 0x43, 0x0c, 0x29, 0x14, 0xeb, 0x98, 0xf5, 0xf6, 0x44, 0x13, 0x1f, 0x6b, 0xbc,
	0x86, 0x79, 0x70, 0x41, 0xd7, 0x4c, 0x26, 0xd4, 0x68, 0xf0, 0x09, 0xfe, 0xbf, 0x0e, 0xe1, 0xe8,
	0xbc, 0x3d, 0xe9, 0x60, 0x5d, 0xcb, 0x78, 0x85, 0xfa, 0x1b, 0x5e, 0x32, 0xe4, 0x97, 0x0c, 0xfb,
	0x8c, 0x8f, 0x79, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xcf, 0xab, 0x28, 0xb2, 0x01, 0x00,
	0x00,
}