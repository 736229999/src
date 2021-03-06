// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent_buycai.proto

package dbproto

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DbBuycaiAgent service

type DbBuycaiAgentClient interface {
	BuycaiQuerySaleList(ctx context.Context, in *StringValue, opts ...grpc.CallOption) (*BuycaiSaleList, error)
	BuycaiQueryIssue(ctx context.Context, in *BuycaiQueryIssueArg, opts ...grpc.CallOption) (*BuycaiSaleIssue, error)
	BuycaiUpsertIssue(ctx context.Context, in *BuycaiUpsertIssueArg, opts ...grpc.CallOption) (*Nil, error)
	BuycaiUpdateOpenBalls(ctx context.Context, in *BuycaiUpsertIssueArg, opts ...grpc.CallOption) (*Nil, error)
}

type dbBuycaiAgentClient struct {
	cc *grpc.ClientConn
}

func NewDbBuycaiAgentClient(cc *grpc.ClientConn) DbBuycaiAgentClient {
	return &dbBuycaiAgentClient{cc}
}

func (c *dbBuycaiAgentClient) BuycaiQuerySaleList(ctx context.Context, in *StringValue, opts ...grpc.CallOption) (*BuycaiSaleList, error) {
	out := new(BuycaiSaleList)
	err := grpc.Invoke(ctx, "/dbproto.DbBuycaiAgent/BuycaiQuerySaleList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbBuycaiAgentClient) BuycaiQueryIssue(ctx context.Context, in *BuycaiQueryIssueArg, opts ...grpc.CallOption) (*BuycaiSaleIssue, error) {
	out := new(BuycaiSaleIssue)
	err := grpc.Invoke(ctx, "/dbproto.DbBuycaiAgent/BuycaiQueryIssue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbBuycaiAgentClient) BuycaiUpsertIssue(ctx context.Context, in *BuycaiUpsertIssueArg, opts ...grpc.CallOption) (*Nil, error) {
	out := new(Nil)
	err := grpc.Invoke(ctx, "/dbproto.DbBuycaiAgent/BuycaiUpsertIssue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbBuycaiAgentClient) BuycaiUpdateOpenBalls(ctx context.Context, in *BuycaiUpsertIssueArg, opts ...grpc.CallOption) (*Nil, error) {
	out := new(Nil)
	err := grpc.Invoke(ctx, "/dbproto.DbBuycaiAgent/BuycaiUpdateOpenBalls", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DbBuycaiAgent service

type DbBuycaiAgentServer interface {
	BuycaiQuerySaleList(context.Context, *StringValue) (*BuycaiSaleList, error)
	BuycaiQueryIssue(context.Context, *BuycaiQueryIssueArg) (*BuycaiSaleIssue, error)
	BuycaiUpsertIssue(context.Context, *BuycaiUpsertIssueArg) (*Nil, error)
	BuycaiUpdateOpenBalls(context.Context, *BuycaiUpsertIssueArg) (*Nil, error)
}

func RegisterDbBuycaiAgentServer(s *grpc.Server, srv DbBuycaiAgentServer) {
	s.RegisterService(&_DbBuycaiAgent_serviceDesc, srv)
}

func _DbBuycaiAgent_BuycaiQuerySaleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbBuycaiAgentServer).BuycaiQuerySaleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbBuycaiAgent/BuycaiQuerySaleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbBuycaiAgentServer).BuycaiQuerySaleList(ctx, req.(*StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbBuycaiAgent_BuycaiQueryIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuycaiQueryIssueArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbBuycaiAgentServer).BuycaiQueryIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbBuycaiAgent/BuycaiQueryIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbBuycaiAgentServer).BuycaiQueryIssue(ctx, req.(*BuycaiQueryIssueArg))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbBuycaiAgent_BuycaiUpsertIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuycaiUpsertIssueArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbBuycaiAgentServer).BuycaiUpsertIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbBuycaiAgent/BuycaiUpsertIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbBuycaiAgentServer).BuycaiUpsertIssue(ctx, req.(*BuycaiUpsertIssueArg))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbBuycaiAgent_BuycaiUpdateOpenBalls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuycaiUpsertIssueArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbBuycaiAgentServer).BuycaiUpdateOpenBalls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbBuycaiAgent/BuycaiUpdateOpenBalls",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbBuycaiAgentServer).BuycaiUpdateOpenBalls(ctx, req.(*BuycaiUpsertIssueArg))
	}
	return interceptor(ctx, in, info, handler)
}

var _DbBuycaiAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dbproto.DbBuycaiAgent",
	HandlerType: (*DbBuycaiAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuycaiQuerySaleList",
			Handler:    _DbBuycaiAgent_BuycaiQuerySaleList_Handler,
		},
		{
			MethodName: "BuycaiQueryIssue",
			Handler:    _DbBuycaiAgent_BuycaiQueryIssue_Handler,
		},
		{
			MethodName: "BuycaiUpsertIssue",
			Handler:    _DbBuycaiAgent_BuycaiUpsertIssue_Handler,
		},
		{
			MethodName: "BuycaiUpdateOpenBalls",
			Handler:    _DbBuycaiAgent_BuycaiUpdateOpenBalls_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent_buycai.proto",
}

func init() { proto.RegisterFile("agent_buycai.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x4c, 0x4f, 0xcd,
	0x2b, 0x89, 0x4f, 0x2a, 0xad, 0x4c, 0x4e, 0xcc, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x4f, 0x49, 0x02, 0x33, 0xa4, 0xb8, 0x92, 0x12, 0x8b, 0x53, 0x21, 0x82, 0x52, 0x3c, 0xc8, 0x4a,
	0x8c, 0x36, 0x32, 0x71, 0xf1, 0xba, 0x24, 0x39, 0x81, 0x85, 0x1c, 0x41, 0x26, 0x08, 0xb9, 0x71,
	0x09, 0x43, 0xb8, 0x81, 0xa5, 0xa9, 0x45, 0x95, 0xc1, 0x89, 0x39, 0xa9, 0x3e, 0x99, 0xc5, 0x25,
	0x42, 0x22, 0x7a, 0x50, 0xc3, 0xf4, 0x82, 0x4b, 0x8a, 0x32, 0xf3, 0xd2, 0xc3, 0x12, 0x73, 0x4a,
	0x53, 0xa5, 0xc4, 0xe1, 0xa2, 0x10, 0x3d, 0x30, 0xe5, 0x4a, 0x0c, 0x42, 0x3e, 0x5c, 0x02, 0x48,
	0xe6, 0x78, 0x16, 0x17, 0x97, 0xa6, 0x0a, 0xc9, 0xa0, 0x29, 0x47, 0x48, 0x39, 0x16, 0xa5, 0x4b,
	0x49, 0x60, 0x31, 0x0c, 0x2c, 0xa9, 0xc4, 0x20, 0xe4, 0xc4, 0x25, 0x08, 0x11, 0x0c, 0x2d, 0x28,
	0x4e, 0x2d, 0x2a, 0x81, 0x18, 0x27, 0x8b, 0xa6, 0x01, 0x49, 0x0e, 0x64, 0x1e, 0x0f, 0x5c, 0xda,
	0x2f, 0x33, 0x47, 0x89, 0x41, 0xc8, 0x8d, 0x4b, 0x14, 0xa6, 0x2e, 0x25, 0xb1, 0x24, 0xd5, 0xbf,
	0x20, 0x35, 0xcf, 0x29, 0x31, 0x27, 0xa7, 0x98, 0x44, 0x73, 0x92, 0xd8, 0xc0, 0x1c, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0xf6, 0x06, 0x37, 0x73, 0x01, 0x00, 0x00,
}
