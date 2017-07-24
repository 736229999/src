// Code generated by protoc-gen-go. DO NOT EDIT.
// source: access.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	access.proto
	idgen.proto
	logic.proto
	manager.proto
	notify.proto
	register.proto

It has these top-level messages:
	ASSendP2PMsgReq
	ASSendP2PMsgRes
	ASSendP2PMsgFromJobReq
	ASSendP2PMsgFromJobRes
	ASSendNotifyReq
	ASSendNotifyRes
	Snowflake
	LoginReq
	LoginRes
	PingReq
	PingRes
	SendP2PMsgReq
	SendP2PMsgRes
	SyncMsgReq
	SyncMsgRes
	AcceptP2PMsgAckReq
	AcceptP2PMsgAckRes
	SendGroupMsgReq
	SendGroupMsgRes
	MGExceptionMsgReq
	MGExceptionMsgRes
	MGSyncMsgReq
	MGSyncMsgRes
	NFNotifyMsgReq
	NFNotifyMsgRes
	RGRegisterReq
	RGRegisterRes
	RGLoginReq
	RGLoginRes
	RGAccessReq
	RGAccessRes
	RGAuthReq
	RGAuthRes
	RGPingReq
	RGPingRes
	RGOnlineReq
	RGOnlineRes
	RGGetUsersByGroupIDReq
	RGGetUsersByGroupIDRes
	RGCreateGroupReq
	RGCreateGroupRes
	RGJoinGroupReq
	RGJoinGroupRes
	RGQuitGroupReq
	RGQuitGroupRes
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

type ASSendP2PMsgReq struct {
	SourceUID int64  `protobuf:"varint,1,opt,name=sourceUID" json:"sourceUID,omitempty"`
	TargetUID int64  `protobuf:"varint,2,opt,name=targetUID" json:"targetUID,omitempty"`
	MsgID     string `protobuf:"bytes,3,opt,name=msgID" json:"msgID,omitempty"`
	Msg       string `protobuf:"bytes,4,opt,name=msg" json:"msg,omitempty"`
}

func (m *ASSendP2PMsgReq) Reset()                    { *m = ASSendP2PMsgReq{} }
func (m *ASSendP2PMsgReq) String() string            { return proto.CompactTextString(m) }
func (*ASSendP2PMsgReq) ProtoMessage()               {}
func (*ASSendP2PMsgReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ASSendP2PMsgReq) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *ASSendP2PMsgReq) GetTargetUID() int64 {
	if m != nil {
		return m.TargetUID
	}
	return 0
}

func (m *ASSendP2PMsgReq) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *ASSendP2PMsgReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ASSendP2PMsgRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *ASSendP2PMsgRes) Reset()                    { *m = ASSendP2PMsgRes{} }
func (m *ASSendP2PMsgRes) String() string            { return proto.CompactTextString(m) }
func (*ASSendP2PMsgRes) ProtoMessage()               {}
func (*ASSendP2PMsgRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ASSendP2PMsgRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *ASSendP2PMsgRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

type ASSendP2PMsgFromJobReq struct {
	SourceUID        int64  `protobuf:"varint,1,opt,name=sourceUID" json:"sourceUID,omitempty"`
	TargetUID        int64  `protobuf:"varint,2,opt,name=targetUID" json:"targetUID,omitempty"`
	MsgID            string `protobuf:"bytes,3,opt,name=msgID" json:"msgID,omitempty"`
	Msg              string `protobuf:"bytes,4,opt,name=msg" json:"msg,omitempty"`
	AccessServerAddr string `protobuf:"bytes,5,opt,name=accessServerAddr" json:"accessServerAddr,omitempty"`
}

func (m *ASSendP2PMsgFromJobReq) Reset()                    { *m = ASSendP2PMsgFromJobReq{} }
func (m *ASSendP2PMsgFromJobReq) String() string            { return proto.CompactTextString(m) }
func (*ASSendP2PMsgFromJobReq) ProtoMessage()               {}
func (*ASSendP2PMsgFromJobReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ASSendP2PMsgFromJobReq) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *ASSendP2PMsgFromJobReq) GetTargetUID() int64 {
	if m != nil {
		return m.TargetUID
	}
	return 0
}

func (m *ASSendP2PMsgFromJobReq) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *ASSendP2PMsgFromJobReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ASSendP2PMsgFromJobReq) GetAccessServerAddr() string {
	if m != nil {
		return m.AccessServerAddr
	}
	return ""
}

type ASSendP2PMsgFromJobRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *ASSendP2PMsgFromJobRes) Reset()                    { *m = ASSendP2PMsgFromJobRes{} }
func (m *ASSendP2PMsgFromJobRes) String() string            { return proto.CompactTextString(m) }
func (*ASSendP2PMsgFromJobRes) ProtoMessage()               {}
func (*ASSendP2PMsgFromJobRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ASSendP2PMsgFromJobRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *ASSendP2PMsgFromJobRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

type ASSendNotifyReq struct {
	UID              int64  `protobuf:"varint,1,opt,name=UID" json:"UID,omitempty"`
	CurrentID        int64  `protobuf:"varint,2,opt,name=currentID" json:"currentID,omitempty"`
	TotalID          int64  `protobuf:"varint,3,opt,name=totalID" json:"totalID,omitempty"`
	AccessServerAddr string `protobuf:"bytes,4,opt,name=accessServerAddr" json:"accessServerAddr,omitempty"`
}

func (m *ASSendNotifyReq) Reset()                    { *m = ASSendNotifyReq{} }
func (m *ASSendNotifyReq) String() string            { return proto.CompactTextString(m) }
func (*ASSendNotifyReq) ProtoMessage()               {}
func (*ASSendNotifyReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ASSendNotifyReq) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *ASSendNotifyReq) GetCurrentID() int64 {
	if m != nil {
		return m.CurrentID
	}
	return 0
}

func (m *ASSendNotifyReq) GetTotalID() int64 {
	if m != nil {
		return m.TotalID
	}
	return 0
}

func (m *ASSendNotifyReq) GetAccessServerAddr() string {
	if m != nil {
		return m.AccessServerAddr
	}
	return ""
}

type ASSendNotifyRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *ASSendNotifyRes) Reset()                    { *m = ASSendNotifyRes{} }
func (m *ASSendNotifyRes) String() string            { return proto.CompactTextString(m) }
func (*ASSendNotifyRes) ProtoMessage()               {}
func (*ASSendNotifyRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ASSendNotifyRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *ASSendNotifyRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

func init() {
	proto.RegisterType((*ASSendP2PMsgReq)(nil), "rpc.ASSendP2PMsgReq")
	proto.RegisterType((*ASSendP2PMsgRes)(nil), "rpc.ASSendP2PMsgRes")
	proto.RegisterType((*ASSendP2PMsgFromJobReq)(nil), "rpc.ASSendP2PMsgFromJobReq")
	proto.RegisterType((*ASSendP2PMsgFromJobRes)(nil), "rpc.ASSendP2PMsgFromJobRes")
	proto.RegisterType((*ASSendNotifyReq)(nil), "rpc.ASSendNotifyReq")
	proto.RegisterType((*ASSendNotifyRes)(nil), "rpc.ASSendNotifyRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AccessServerRPC service

type AccessServerRPCClient interface {
	SendP2PMsgFromJob(ctx context.Context, in *ASSendP2PMsgFromJobReq, opts ...grpc.CallOption) (*ASSendP2PMsgFromJobRes, error)
	SendNotify(ctx context.Context, in *ASSendNotifyReq, opts ...grpc.CallOption) (*ASSendNotifyRes, error)
}

type accessServerRPCClient struct {
	cc *grpc.ClientConn
}

func NewAccessServerRPCClient(cc *grpc.ClientConn) AccessServerRPCClient {
	return &accessServerRPCClient{cc}
}

func (c *accessServerRPCClient) SendP2PMsgFromJob(ctx context.Context, in *ASSendP2PMsgFromJobReq, opts ...grpc.CallOption) (*ASSendP2PMsgFromJobRes, error) {
	out := new(ASSendP2PMsgFromJobRes)
	err := grpc.Invoke(ctx, "/rpc.AccessServerRPC/SendP2PMsgFromJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessServerRPCClient) SendNotify(ctx context.Context, in *ASSendNotifyReq, opts ...grpc.CallOption) (*ASSendNotifyRes, error) {
	out := new(ASSendNotifyRes)
	err := grpc.Invoke(ctx, "/rpc.AccessServerRPC/SendNotify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccessServerRPC service

type AccessServerRPCServer interface {
	SendP2PMsgFromJob(context.Context, *ASSendP2PMsgFromJobReq) (*ASSendP2PMsgFromJobRes, error)
	SendNotify(context.Context, *ASSendNotifyReq) (*ASSendNotifyRes, error)
}

func RegisterAccessServerRPCServer(s *grpc.Server, srv AccessServerRPCServer) {
	s.RegisterService(&_AccessServerRPC_serviceDesc, srv)
}

func _AccessServerRPC_SendP2PMsgFromJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ASSendP2PMsgFromJobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServerRPCServer).SendP2PMsgFromJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.AccessServerRPC/SendP2PMsgFromJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServerRPCServer).SendP2PMsgFromJob(ctx, req.(*ASSendP2PMsgFromJobReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessServerRPC_SendNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ASSendNotifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServerRPCServer).SendNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.AccessServerRPC/SendNotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServerRPCServer).SendNotify(ctx, req.(*ASSendNotifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccessServerRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.AccessServerRPC",
	HandlerType: (*AccessServerRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendP2PMsgFromJob",
			Handler:    _AccessServerRPC_SendP2PMsgFromJob_Handler,
		},
		{
			MethodName: "SendNotify",
			Handler:    _AccessServerRPC_SendNotify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "access.proto",
}

func init() { proto.RegisterFile("access.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0xa9, 0x03, 0x18, 0x6e, 0x34, 0xe2, 0x84, 0x90, 0x06, 0x5d, 0x90, 0x59, 0x11, 0x17,
	0x2c, 0x70, 0xe7, 0x8e, 0x94, 0x98, 0x40, 0xa2, 0x92, 0x69, 0x7c, 0x00, 0x98, 0x8e, 0x8d, 0x89,
	0x30, 0xf5, 0xce, 0xd4, 0xc4, 0x27, 0xf0, 0x2d, 0x5c, 0xfb, 0x98, 0x66, 0xfa, 0x1f, 0x28, 0x2e,
	0x58, 0xb8, 0xeb, 0x3d, 0xa7, 0xb9, 0xe7, 0x7e, 0xa7, 0x29, 0x9c, 0xad, 0x84, 0x90, 0x5a, 0x8f,
	0x23, 0x54, 0x46, 0x51, 0x82, 0x91, 0x60, 0x31, 0x5c, 0x4c, 0x7d, 0x5f, 0x6e, 0x83, 0xe5, 0x64,
	0xf9, 0xa0, 0x43, 0x2e, 0xdf, 0xe9, 0x35, 0x74, 0xb4, 0x8a, 0x51, 0xc8, 0xe7, 0xf9, 0xcc, 0x75,
	0x86, 0xce, 0x88, 0xf0, 0x52, 0xb0, 0xae, 0x59, 0x61, 0x28, 0x8d, 0x75, 0x4f, 0x52, 0xb7, 0x10,
	0x68, 0x0f, 0x5a, 0x1b, 0x1d, 0xce, 0x67, 0x2e, 0x19, 0x3a, 0xa3, 0x0e, 0x4f, 0x07, 0xda, 0x05,
	0xb2, 0xd1, 0xa1, 0xdb, 0x4c, 0x34, 0xfb, 0xc8, 0xbc, 0xdd, 0x58, 0x4d, 0x5d, 0x38, 0x95, 0x88,
	0x9e, 0x0a, 0x64, 0x12, 0x7a, 0xce, 0xf3, 0x91, 0xf6, 0xa1, 0x2d, 0x11, 0x7d, 0x83, 0x49, 0x5e,
	0x87, 0x67, 0x13, 0xfb, 0x71, 0xa0, 0x5f, 0xdd, 0x72, 0x8f, 0x6a, 0xb3, 0x50, 0xeb, 0x7f, 0x62,
	0xa0, 0x37, 0xd0, 0x4d, 0xfb, 0xf4, 0x25, 0x7e, 0x48, 0x9c, 0x06, 0x01, 0xba, 0xad, 0xc4, 0xde,
	0xd3, 0xd9, 0xe2, 0xc0, 0xa5, 0xc7, 0x60, 0x7f, 0x39, 0x79, 0x79, 0x8f, 0xca, 0xbc, 0xbe, 0x7c,
	0x5a, 0xde, 0x2e, 0x90, 0x92, 0x94, 0x64, 0x8c, 0x22, 0x46, 0x94, 0x5b, 0x53, 0x32, 0x16, 0x82,
	0x4d, 0x35, 0xca, 0xac, 0xde, 0x32, 0x4a, 0xc2, 0xf3, 0xb1, 0x96, 0xaa, 0x79, 0x80, 0xca, 0xdb,
	0x3d, 0xe4, 0x08, 0x9c, 0xc9, 0xb7, 0xc5, 0xa9, 0x6c, 0xe6, 0x4b, 0x8f, 0x3e, 0xc1, 0xe5, 0x5e,
	0x59, 0xf4, 0x6a, 0x8c, 0x91, 0x18, 0xd7, 0x7f, 0xf0, 0xc1, 0x1f, 0xa6, 0x66, 0x0d, 0x7a, 0x07,
	0x50, 0xde, 0x49, 0x7b, 0x95, 0x97, 0x8b, 0x0e, 0x07, 0x75, 0xaa, 0x66, 0x8d, 0x75, 0x3b, 0xf9,
	0x5d, 0x6e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x20, 0x20, 0x31, 0x3e, 0x03, 0x00, 0x00,
}