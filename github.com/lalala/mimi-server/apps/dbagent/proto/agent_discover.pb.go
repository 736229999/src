// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent_discover.proto

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

// Client API for DbDiscoveragent service

type DbDiscoveragentClient interface {
	// 获取一条新闻
	GetANews(ctx context.Context, in *NewsId, opts ...grpc.CallOption) (*News, error)
	// 客户端获取新闻列表
	QueryNewsListClient(ctx context.Context, in *QueryNewsArg, opts ...grpc.CallOption) (*NewsList, error)
	// 阅读一条新闻（阅读量+1）
	ReadANews(ctx context.Context, in *NewsId, opts ...grpc.CallOption) (*Nil, error)
}

type dbDiscoveragentClient struct {
	cc *grpc.ClientConn
}

func NewDbDiscoveragentClient(cc *grpc.ClientConn) DbDiscoveragentClient {
	return &dbDiscoveragentClient{cc}
}

func (c *dbDiscoveragentClient) GetANews(ctx context.Context, in *NewsId, opts ...grpc.CallOption) (*News, error) {
	out := new(News)
	err := grpc.Invoke(ctx, "/dbproto.DbDiscoveragent/GetANews", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbDiscoveragentClient) QueryNewsListClient(ctx context.Context, in *QueryNewsArg, opts ...grpc.CallOption) (*NewsList, error) {
	out := new(NewsList)
	err := grpc.Invoke(ctx, "/dbproto.DbDiscoveragent/QueryNewsListClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbDiscoveragentClient) ReadANews(ctx context.Context, in *NewsId, opts ...grpc.CallOption) (*Nil, error) {
	out := new(Nil)
	err := grpc.Invoke(ctx, "/dbproto.DbDiscoveragent/ReadANews", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DbDiscoveragent service

type DbDiscoveragentServer interface {
	// 获取一条新闻
	GetANews(context.Context, *NewsId) (*News, error)
	// 客户端获取新闻列表
	QueryNewsListClient(context.Context, *QueryNewsArg) (*NewsList, error)
	// 阅读一条新闻（阅读量+1）
	ReadANews(context.Context, *NewsId) (*Nil, error)
}

func RegisterDbDiscoveragentServer(s *grpc.Server, srv DbDiscoveragentServer) {
	s.RegisterService(&_DbDiscoveragent_serviceDesc, srv)
}

func _DbDiscoveragent_GetANews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewsId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbDiscoveragentServer).GetANews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbDiscoveragent/GetANews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbDiscoveragentServer).GetANews(ctx, req.(*NewsId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbDiscoveragent_QueryNewsListClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNewsArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbDiscoveragentServer).QueryNewsListClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbDiscoveragent/QueryNewsListClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbDiscoveragentServer).QueryNewsListClient(ctx, req.(*QueryNewsArg))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbDiscoveragent_ReadANews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewsId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbDiscoveragentServer).ReadANews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbDiscoveragent/ReadANews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbDiscoveragentServer).ReadANews(ctx, req.(*NewsId))
	}
	return interceptor(ctx, in, info, handler)
}

var _DbDiscoveragent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dbproto.DbDiscoveragent",
	HandlerType: (*DbDiscoveragentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetANews",
			Handler:    _DbDiscoveragent_GetANews_Handler,
		},
		{
			MethodName: "QueryNewsListClient",
			Handler:    _DbDiscoveragent_QueryNewsListClient_Handler,
		},
		{
			MethodName: "ReadANews",
			Handler:    _DbDiscoveragent_ReadANews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent_discover.proto",
}

func init() { proto.RegisterFile("agent_discover.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x4c, 0x4f, 0xcd,
	0x2b, 0x89, 0x4f, 0xc9, 0x2c, 0x4e, 0xce, 0x2f, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x4f, 0x49, 0x02, 0x33, 0xa4, 0xb8, 0x92, 0x12, 0x8b, 0x53, 0x21, 0x82, 0x52, 0x7c,
	0xa8, 0x8a, 0x8c, 0x36, 0x30, 0x72, 0xf1, 0xbb, 0x24, 0xb9, 0x40, 0x05, 0xc1, 0xe6, 0x08, 0xe9,
	0x70, 0x71, 0xb8, 0xa7, 0x96, 0x38, 0xfa, 0xa5, 0x96, 0x17, 0x0b, 0xf1, 0xeb, 0x41, 0x4d, 0xd1,
	0x03, 0x71, 0x3d, 0x53, 0xa4, 0x78, 0x51, 0x04, 0x94, 0x18, 0x84, 0x1c, 0xb9, 0x84, 0x03, 0x4b,
	0x53, 0x8b, 0x2a, 0x41, 0x5c, 0x9f, 0xcc, 0xe2, 0x12, 0xe7, 0x9c, 0x4c, 0x90, 0x21, 0xa2, 0x70,
	0x75, 0x70, 0x59, 0xc7, 0xa2, 0x74, 0x29, 0x41, 0x14, 0xed, 0x20, 0xf5, 0x4a, 0x0c, 0x42, 0x3a,
	0x5c, 0x9c, 0x41, 0xa9, 0x89, 0x29, 0x38, 0x6c, 0xe4, 0x41, 0x08, 0x64, 0xe6, 0x28, 0x31, 0x24,
	0xb1, 0x81, 0x39, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x59, 0xb7, 0x5c, 0xf6, 0x00,
	0x00, 0x00,
}
