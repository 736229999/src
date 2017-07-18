// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent_options.proto

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

// Client API for DbOptionsAgent service

type DbOptionsAgentClient interface {
	QueryLotteryOptionsList(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*HomeParams, error)
	// 创建一条Banner信息
	CreateBanner(ctx context.Context, in *Banner, opts ...grpc.CallOption) (*IntValue, error)
	// 后台查询Banner列表
	QueryBannerList(ctx context.Context, in *QueryBannerArg, opts ...grpc.CallOption) (*BannerList, error)
	// 客户端查询Banner列表
	QueryClientBannerList(ctx context.Context, in *QueryClientBannerArg, opts ...grpc.CallOption) (*BannerList, error)
	// 查询一条Banner信息
	QueryBannerById(ctx context.Context, in *BannerId, opts ...grpc.CallOption) (*Banner, error)
	// 获取一条新闻
	QueryNewsById(ctx context.Context, in *NewsId, opts ...grpc.CallOption) (*News, error)
	// 客户端获取新闻列表
	QueryNewsList(ctx context.Context, in *QueryNewsArg, opts ...grpc.CallOption) (*NewsList, error)
	// 阅读一条新闻（阅读量+1）
	ReadANews(ctx context.Context, in *NewsId, opts ...grpc.CallOption) (*Nil, error)
	// 查询常见问题列表
	QueryFaqList(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*FaqList, error)
	// 查询一条常见问题
	QueryFaqById(ctx context.Context, in *FaqId, opts ...grpc.CallOption) (*Faq, error)
	QueryContact(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*Contact, error)
	InsertFeedback(ctx context.Context, in *Feedback, opts ...grpc.CallOption) (*Nil, error)
}

type dbOptionsAgentClient struct {
	cc *grpc.ClientConn
}

func NewDbOptionsAgentClient(cc *grpc.ClientConn) DbOptionsAgentClient {
	return &dbOptionsAgentClient{cc}
}

func (c *dbOptionsAgentClient) QueryLotteryOptionsList(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*HomeParams, error) {
	out := new(HomeParams)
	err := grpc.Invoke(ctx, "/dbproto.DbOptionsAgent/QueryLotteryOptionsList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbOptionsAgentClient) CreateBanner(ctx context.Context, in *Banner, opts ...grpc.CallOption) (*IntValue, error) {
	out := new(IntValue)
	err := grpc.Invoke(ctx, "/dbproto.DbOptionsAgent/CreateBanner", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbOptionsAgentClient) QueryBannerList(ctx context.Context, in *QueryBannerArg, opts ...grpc.CallOption) (*BannerList, error) {
	out := new(BannerList)
	err := grpc.Invoke(ctx, "/dbproto.DbOptionsAgent/QueryBannerList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbOptionsAgentClient) QueryClientBannerList(ctx context.Context, in *QueryClientBannerArg, opts ...grpc.CallOption) (*BannerList, error) {
	out := new(BannerList)
	err := grpc.Invoke(ctx, "/dbproto.DbOptionsAgent/QueryClientBannerList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbOptionsAgentClient) QueryBannerById(ctx context.Context, in *BannerId, opts ...grpc.CallOption) (*Banner, error) {
	out := new(Banner)
	err := grpc.Invoke(ctx, "/dbproto.DbOptionsAgent/QueryBannerById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbOptionsAgentClient) QueryNewsById(ctx context.Context, in *NewsId, opts ...grpc.CallOption) (*News, error) {
	out := new(News)
	err := grpc.Invoke(ctx, "/dbproto.DbOptionsAgent/QueryNewsById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbOptionsAgentClient) QueryNewsList(ctx context.Context, in *QueryNewsArg, opts ...grpc.CallOption) (*NewsList, error) {
	out := new(NewsList)
	err := grpc.Invoke(ctx, "/dbproto.DbOptionsAgent/QueryNewsList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbOptionsAgentClient) ReadANews(ctx context.Context, in *NewsId, opts ...grpc.CallOption) (*Nil, error) {
	out := new(Nil)
	err := grpc.Invoke(ctx, "/dbproto.DbOptionsAgent/ReadANews", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbOptionsAgentClient) QueryFaqList(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*FaqList, error) {
	out := new(FaqList)
	err := grpc.Invoke(ctx, "/dbproto.DbOptionsAgent/QueryFaqList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbOptionsAgentClient) QueryFaqById(ctx context.Context, in *FaqId, opts ...grpc.CallOption) (*Faq, error) {
	out := new(Faq)
	err := grpc.Invoke(ctx, "/dbproto.DbOptionsAgent/QueryFaqById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbOptionsAgentClient) QueryContact(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*Contact, error) {
	out := new(Contact)
	err := grpc.Invoke(ctx, "/dbproto.DbOptionsAgent/QueryContact", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbOptionsAgentClient) InsertFeedback(ctx context.Context, in *Feedback, opts ...grpc.CallOption) (*Nil, error) {
	out := new(Nil)
	err := grpc.Invoke(ctx, "/dbproto.DbOptionsAgent/InsertFeedback", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DbOptionsAgent service

type DbOptionsAgentServer interface {
	QueryLotteryOptionsList(context.Context, *Nil) (*HomeParams, error)
	// 创建一条Banner信息
	CreateBanner(context.Context, *Banner) (*IntValue, error)
	// 后台查询Banner列表
	QueryBannerList(context.Context, *QueryBannerArg) (*BannerList, error)
	// 客户端查询Banner列表
	QueryClientBannerList(context.Context, *QueryClientBannerArg) (*BannerList, error)
	// 查询一条Banner信息
	QueryBannerById(context.Context, *BannerId) (*Banner, error)
	// 获取一条新闻
	QueryNewsById(context.Context, *NewsId) (*News, error)
	// 客户端获取新闻列表
	QueryNewsList(context.Context, *QueryNewsArg) (*NewsList, error)
	// 阅读一条新闻（阅读量+1）
	ReadANews(context.Context, *NewsId) (*Nil, error)
	// 查询常见问题列表
	QueryFaqList(context.Context, *Nil) (*FaqList, error)
	// 查询一条常见问题
	QueryFaqById(context.Context, *FaqId) (*Faq, error)
	QueryContact(context.Context, *Nil) (*Contact, error)
	InsertFeedback(context.Context, *Feedback) (*Nil, error)
}

func RegisterDbOptionsAgentServer(s *grpc.Server, srv DbOptionsAgentServer) {
	s.RegisterService(&_DbOptionsAgent_serviceDesc, srv)
}

func _DbOptionsAgent_QueryLotteryOptionsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbOptionsAgentServer).QueryLotteryOptionsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbOptionsAgent/QueryLotteryOptionsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbOptionsAgentServer).QueryLotteryOptionsList(ctx, req.(*Nil))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbOptionsAgent_CreateBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Banner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbOptionsAgentServer).CreateBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbOptionsAgent/CreateBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbOptionsAgentServer).CreateBanner(ctx, req.(*Banner))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbOptionsAgent_QueryBannerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBannerArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbOptionsAgentServer).QueryBannerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbOptionsAgent/QueryBannerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbOptionsAgentServer).QueryBannerList(ctx, req.(*QueryBannerArg))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbOptionsAgent_QueryClientBannerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClientBannerArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbOptionsAgentServer).QueryClientBannerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbOptionsAgent/QueryClientBannerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbOptionsAgentServer).QueryClientBannerList(ctx, req.(*QueryClientBannerArg))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbOptionsAgent_QueryBannerById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbOptionsAgentServer).QueryBannerById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbOptionsAgent/QueryBannerById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbOptionsAgentServer).QueryBannerById(ctx, req.(*BannerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbOptionsAgent_QueryNewsById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewsId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbOptionsAgentServer).QueryNewsById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbOptionsAgent/QueryNewsById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbOptionsAgentServer).QueryNewsById(ctx, req.(*NewsId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbOptionsAgent_QueryNewsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNewsArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbOptionsAgentServer).QueryNewsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbOptionsAgent/QueryNewsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbOptionsAgentServer).QueryNewsList(ctx, req.(*QueryNewsArg))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbOptionsAgent_ReadANews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewsId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbOptionsAgentServer).ReadANews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbOptionsAgent/ReadANews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbOptionsAgentServer).ReadANews(ctx, req.(*NewsId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbOptionsAgent_QueryFaqList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbOptionsAgentServer).QueryFaqList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbOptionsAgent/QueryFaqList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbOptionsAgentServer).QueryFaqList(ctx, req.(*Nil))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbOptionsAgent_QueryFaqById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FaqId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbOptionsAgentServer).QueryFaqById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbOptionsAgent/QueryFaqById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbOptionsAgentServer).QueryFaqById(ctx, req.(*FaqId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbOptionsAgent_QueryContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbOptionsAgentServer).QueryContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbOptionsAgent/QueryContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbOptionsAgentServer).QueryContact(ctx, req.(*Nil))
	}
	return interceptor(ctx, in, info, handler)
}

func _DbOptionsAgent_InsertFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Feedback)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbOptionsAgentServer).InsertFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbproto.DbOptionsAgent/InsertFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbOptionsAgentServer).InsertFeedback(ctx, req.(*Feedback))
	}
	return interceptor(ctx, in, info, handler)
}

var _DbOptionsAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dbproto.DbOptionsAgent",
	HandlerType: (*DbOptionsAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryLotteryOptionsList",
			Handler:    _DbOptionsAgent_QueryLotteryOptionsList_Handler,
		},
		{
			MethodName: "CreateBanner",
			Handler:    _DbOptionsAgent_CreateBanner_Handler,
		},
		{
			MethodName: "QueryBannerList",
			Handler:    _DbOptionsAgent_QueryBannerList_Handler,
		},
		{
			MethodName: "QueryClientBannerList",
			Handler:    _DbOptionsAgent_QueryClientBannerList_Handler,
		},
		{
			MethodName: "QueryBannerById",
			Handler:    _DbOptionsAgent_QueryBannerById_Handler,
		},
		{
			MethodName: "QueryNewsById",
			Handler:    _DbOptionsAgent_QueryNewsById_Handler,
		},
		{
			MethodName: "QueryNewsList",
			Handler:    _DbOptionsAgent_QueryNewsList_Handler,
		},
		{
			MethodName: "ReadANews",
			Handler:    _DbOptionsAgent_ReadANews_Handler,
		},
		{
			MethodName: "QueryFaqList",
			Handler:    _DbOptionsAgent_QueryFaqList_Handler,
		},
		{
			MethodName: "QueryFaqById",
			Handler:    _DbOptionsAgent_QueryFaqById_Handler,
		},
		{
			MethodName: "QueryContact",
			Handler:    _DbOptionsAgent_QueryContact_Handler,
		},
		{
			MethodName: "InsertFeedback",
			Handler:    _DbOptionsAgent_InsertFeedback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent_options.proto",
}

func init() { proto.RegisterFile("agent_options.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x5d, 0x4b, 0xfb, 0x30,
	0x14, 0xc6, 0x73, 0xf3, 0xff, 0x8b, 0x61, 0xeb, 0x5c, 0xc6, 0x18, 0x04, 0xbc, 0xc9, 0xb5, 0x14,
	0x75, 0x7a, 0x25, 0x08, 0xdd, 0x64, 0x58, 0x98, 0xf3, 0xe5, 0xc2, 0x5b, 0x49, 0x97, 0xc3, 0x28,
	0x76, 0xc9, 0x96, 0x64, 0xca, 0xbe, 0x98, 0x9f, 0x4f, 0x9a, 0xd6, 0x2e, 0x99, 0x2f, 0x77, 0x3d,
	0xcf, 0xf9, 0x9d, 0xf3, 0x3c, 0xa7, 0xc1, 0x3d, 0xbe, 0x00, 0x69, 0x5f, 0xd4, 0xca, 0xe6, 0x4a,
	0x9a, 0x78, 0xa5, 0x95, 0x55, 0xe4, 0x40, 0x64, 0xee, 0x83, 0xe2, 0x8c, 0x1b, 0xa8, 0x44, 0xda,
	0x2f, 0x94, 0xb5, 0xa0, 0xb7, 0x21, 0x4b, 0xdb, 0x61, 0x19, 0x89, 0xdc, 0xcc, 0xd5, 0x1b, 0xe8,
	0xaa, 0x3e, 0xff, 0xf8, 0x87, 0xa3, 0x9b, 0xec, 0xbe, 0x62, 0x92, 0xd2, 0x8b, 0x5c, 0xe3, 0xc1,
	0xe3, 0x06, 0xf4, 0x76, 0x5a, 0xed, 0xab, 0x7b, 0xd3, 0xdc, 0x58, 0xd2, 0x8a, 0x6b, 0xe7, 0x78,
	0x96, 0x17, 0xb4, 0xd7, 0x54, 0xb7, 0x6a, 0x09, 0x0f, 0x5c, 0xf3, 0xa5, 0x61, 0x88, 0x5c, 0xe0,
	0xd6, 0x58, 0x03, 0xb7, 0x30, 0xe2, 0x52, 0x82, 0x26, 0x9d, 0x06, 0xab, 0x04, 0xda, 0x6d, 0x84,
	0x54, 0xda, 0x67, 0x5e, 0x6c, 0x80, 0x21, 0x92, 0xe0, 0x8e, 0x73, 0xad, 0x18, 0xe7, 0x36, 0x68,
	0x38, 0xaf, 0x93, 0xe8, 0x85, 0x67, 0xbc, 0xa3, 0x19, 0x22, 0x77, 0xb8, 0xef, 0xc0, 0x71, 0x91,
	0x83, 0xb4, 0xde, 0xa2, 0xe3, 0x70, 0x91, 0xdf, 0xff, 0x63, 0xdd, 0x65, 0x90, 0x68, 0xb4, 0x4d,
	0x05, 0xe9, 0xee, 0x91, 0xa9, 0xa0, 0xfb, 0xd7, 0x91, 0x33, 0xdc, 0x76, 0x63, 0x33, 0x78, 0x37,
	0x6e, 0x68, 0x47, 0x94, 0x52, 0x2a, 0x68, 0x3b, 0x10, 0x18, 0x22, 0x57, 0xde, 0x88, 0x0b, 0xdc,
	0x0f, 0x03, 0x97, 0x7a, 0x19, 0xb4, 0x1b, 0x0c, 0xd6, 0x31, 0x4f, 0xf0, 0xe1, 0x13, 0x70, 0x91,
	0x94, 0xd2, 0x77, 0xaf, 0xe0, 0xc5, 0x18, 0x22, 0xa7, 0xb8, 0xe5, 0x56, 0x4e, 0xf8, 0xfa, 0x87,
	0x17, 0x3d, 0x6a, 0xaa, 0xba, 0xcf, 0x10, 0x89, 0x77, 0x13, 0xee, 0x9c, 0xc8, 0x67, 0x02, 0x87,
	0x09, 0x5f, 0x7b, 0x0e, 0x63, 0x25, 0x2d, 0x9f, 0xff, 0xee, 0x50, 0xf7, 0x19, 0x22, 0x43, 0x1c,
	0xa5, 0xd2, 0x80, 0xb6, 0x13, 0x00, 0x91, 0xf1, 0xf9, 0xab, 0xf7, 0x9f, 0xbf, 0xa4, 0xfd, 0x43,
	0xb2, 0xff, 0xae, 0x18, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xac, 0xed, 0xa8, 0x07, 0x21, 0x03,
	0x00, 0x00,
}
