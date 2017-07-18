// Code generated by protoc-gen-go.
// source: discover.proto
// DO NOT EDIT!

/*
Package discoverproto is a generated protocol buffer package.

It is generated from these files:
	discover.proto

It has these top-level messages:
	News
	QueryNewsArg
	NewsList
	IntValue
	QueryNewsOfSelect
*/
package discoverproto

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

type News struct {
	Id          int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Cover       string `protobuf:"bytes,4,opt,name=cover" json:"cover,omitempty"`
	Content     string `protobuf:"bytes,5,opt,name=content" json:"content,omitempty"`
	Html        string `protobuf:"bytes,6,opt,name=html" json:"html,omitempty"`
	Author      string `protobuf:"bytes,7,opt,name=author" json:"author,omitempty"`
	Created     int64  `protobuf:"varint,8,opt,name=created" json:"created,omitempty"`
	NewsClass   int64  `protobuf:"varint,9,opt,name=newsClass" json:"newsClass,omitempty"`
	IsVisible   bool   `protobuf:"varint,10,opt,name=isVisible" json:"isVisible,omitempty"`
	Updated     int64  `protobuf:"varint,11,opt,name=updated" json:"updated,omitempty"`
	PageViews   int64  `protobuf:"varint,12,opt,name=pageViews" json:"pageViews,omitempty"`
}

func (m *News) Reset()                    { *m = News{} }
func (m *News) String() string            { return proto.CompactTextString(m) }
func (*News) ProtoMessage()               {}
func (*News) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *News) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *News) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *News) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *News) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *News) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *News) GetHtml() string {
	if m != nil {
		return m.Html
	}
	return ""
}

func (m *News) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *News) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *News) GetNewsClass() int64 {
	if m != nil {
		return m.NewsClass
	}
	return 0
}

func (m *News) GetIsVisible() bool {
	if m != nil {
		return m.IsVisible
	}
	return false
}

func (m *News) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *News) GetPageViews() int64 {
	if m != nil {
		return m.PageViews
	}
	return 0
}

type QueryNewsArg struct {
	Title    string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Author   string `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
	Start    int64  `protobuf:"varint,3,opt,name=start" json:"start,omitempty"`
	End      int64  `protobuf:"varint,4,opt,name=end" json:"end,omitempty"`
	Class    int64  `protobuf:"varint,5,opt,name=class" json:"class,omitempty"`
	Page     int64  `protobuf:"varint,6,opt,name=page" json:"page,omitempty"`
	PageSize int64  `protobuf:"varint,7,opt,name=pageSize" json:"pageSize,omitempty"`
}

func (m *QueryNewsArg) Reset()                    { *m = QueryNewsArg{} }
func (m *QueryNewsArg) String() string            { return proto.CompactTextString(m) }
func (*QueryNewsArg) ProtoMessage()               {}
func (*QueryNewsArg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *QueryNewsArg) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *QueryNewsArg) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *QueryNewsArg) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *QueryNewsArg) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *QueryNewsArg) GetClass() int64 {
	if m != nil {
		return m.Class
	}
	return 0
}

func (m *QueryNewsArg) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *QueryNewsArg) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type NewsList struct {
	List  []*News `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
	Total int64   `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
}

func (m *NewsList) Reset()                    { *m = NewsList{} }
func (m *NewsList) String() string            { return proto.CompactTextString(m) }
func (*NewsList) ProtoMessage()               {}
func (*NewsList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NewsList) GetList() []*News {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *NewsList) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type IntValue struct {
	Value int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *IntValue) Reset()                    { *m = IntValue{} }
func (m *IntValue) String() string            { return proto.CompactTextString(m) }
func (*IntValue) ProtoMessage()               {}
func (*IntValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IntValue) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// 后台select框搜索查询接口
type QueryNewsOfSelect struct {
	KeyWord string `protobuf:"bytes,1,opt,name=keyWord" json:"keyWord,omitempty"`
}

func (m *QueryNewsOfSelect) Reset()                    { *m = QueryNewsOfSelect{} }
func (m *QueryNewsOfSelect) String() string            { return proto.CompactTextString(m) }
func (*QueryNewsOfSelect) ProtoMessage()               {}
func (*QueryNewsOfSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *QueryNewsOfSelect) GetKeyWord() string {
	if m != nil {
		return m.KeyWord
	}
	return ""
}

func init() {
	proto.RegisterType((*News)(nil), "discoverproto.News")
	proto.RegisterType((*QueryNewsArg)(nil), "discoverproto.QueryNewsArg")
	proto.RegisterType((*NewsList)(nil), "discoverproto.NewsList")
	proto.RegisterType((*IntValue)(nil), "discoverproto.IntValue")
	proto.RegisterType((*QueryNewsOfSelect)(nil), "discoverproto.QueryNewsOfSelect")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Discover service

type DiscoverClient interface {
	CreateNews(ctx context.Context, in *News, opts ...grpc.CallOption) (*IntValue, error)
	QueryNewsList(ctx context.Context, in *QueryNewsArg, opts ...grpc.CallOption) (*NewsList, error)
	QueryBakendSelectOfNews(ctx context.Context, in *QueryNewsOfSelect, opts ...grpc.CallOption) (*NewsList, error)
}

type discoverClient struct {
	cc *grpc.ClientConn
}

func NewDiscoverClient(cc *grpc.ClientConn) DiscoverClient {
	return &discoverClient{cc}
}

func (c *discoverClient) CreateNews(ctx context.Context, in *News, opts ...grpc.CallOption) (*IntValue, error) {
	out := new(IntValue)
	err := grpc.Invoke(ctx, "/discoverproto.Discover/CreateNews", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoverClient) QueryNewsList(ctx context.Context, in *QueryNewsArg, opts ...grpc.CallOption) (*NewsList, error) {
	out := new(NewsList)
	err := grpc.Invoke(ctx, "/discoverproto.Discover/QueryNewsList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoverClient) QueryBakendSelectOfNews(ctx context.Context, in *QueryNewsOfSelect, opts ...grpc.CallOption) (*NewsList, error) {
	out := new(NewsList)
	err := grpc.Invoke(ctx, "/discoverproto.Discover/QueryBakendSelectOfNews", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Discover service

type DiscoverServer interface {
	CreateNews(context.Context, *News) (*IntValue, error)
	QueryNewsList(context.Context, *QueryNewsArg) (*NewsList, error)
	QueryBakendSelectOfNews(context.Context, *QueryNewsOfSelect) (*NewsList, error)
}

func RegisterDiscoverServer(s *grpc.Server, srv DiscoverServer) {
	s.RegisterService(&_Discover_serviceDesc, srv)
}

func _Discover_CreateNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(News)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoverServer).CreateNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discoverproto.Discover/CreateNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoverServer).CreateNews(ctx, req.(*News))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discover_QueryNewsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNewsArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoverServer).QueryNewsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discoverproto.Discover/QueryNewsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoverServer).QueryNewsList(ctx, req.(*QueryNewsArg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discover_QueryBakendSelectOfNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNewsOfSelect)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoverServer).QueryBakendSelectOfNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discoverproto.Discover/QueryBakendSelectOfNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoverServer).QueryBakendSelectOfNews(ctx, req.(*QueryNewsOfSelect))
	}
	return interceptor(ctx, in, info, handler)
}

var _Discover_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discoverproto.Discover",
	HandlerType: (*DiscoverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNews",
			Handler:    _Discover_CreateNews_Handler,
		},
		{
			MethodName: "QueryNewsList",
			Handler:    _Discover_QueryNewsList_Handler,
		},
		{
			MethodName: "QueryBakendSelectOfNews",
			Handler:    _Discover_QueryBakendSelectOfNews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discover.proto",
}

func init() { proto.RegisterFile("discover.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xbd, 0x49, 0xea, 0x4c, 0xda, 0x0a, 0x96, 0x8a, 0xae, 0x0a, 0x07, 0xcb, 0x17, 0x72,
	0x21, 0x87, 0x72, 0xe5, 0x02, 0x45, 0x42, 0x95, 0x10, 0x15, 0x5b, 0x29, 0x9c, 0xdd, 0xec, 0xb4,
	0x5d, 0xd5, 0xd8, 0xd1, 0xee, 0xa4, 0x51, 0xf9, 0x49, 0xfc, 0x02, 0xfe, 0x12, 0xff, 0x02, 0xed,
	0xac, 0x1d, 0xa7, 0xe5, 0xe3, 0xe4, 0x79, 0x6f, 0x67, 0x9e, 0xdf, 0xbc, 0x5d, 0x38, 0x30, 0xd6,
	0x2f, 0x9a, 0x3b, 0x74, 0xb3, 0xa5, 0x6b, 0xa8, 0x91, 0xfb, 0x1d, 0x66, 0x58, 0xfc, 0x4c, 0x61,
	0xf0, 0x19, 0xd7, 0x5e, 0x1e, 0x40, 0x6a, 0x8d, 0x4a, 0xf2, 0x64, 0x2a, 0x74, 0x6a, 0x8d, 0x3c,
	0x84, 0x21, 0x59, 0xaa, 0x50, 0xa5, 0x79, 0x32, 0x1d, 0xeb, 0x08, 0x64, 0x0e, 0x13, 0x83, 0x7e,
	0xe1, 0xec, 0x92, 0x6c, 0x53, 0x2b, 0xc1, 0x67, 0xdb, 0x54, 0x98, 0x63, 0x79, 0x35, 0x88, 0x73,
	0x0c, 0xa4, 0x82, 0xdd, 0x45, 0x53, 0x13, 0xd6, 0xa4, 0x86, 0xcc, 0x77, 0x50, 0x4a, 0x18, 0xdc,
	0xd0, 0xb7, 0x4a, 0x8d, 0x98, 0xe6, 0x5a, 0x3e, 0x87, 0x51, 0xb9, 0xa2, 0x9b, 0xc6, 0xa9, 0x5d,
	0x66, 0x5b, 0xc4, 0x2a, 0x0e, 0x4b, 0x42, 0xa3, 0x32, 0x36, 0xda, 0x41, 0xf9, 0x12, 0xc6, 0x35,
	0xae, 0xfd, 0x69, 0x55, 0x7a, 0xaf, 0xc6, 0x7c, 0xd6, 0x13, 0xe1, 0xd4, 0xfa, 0xb9, 0xf5, 0xf6,
	0xb2, 0x42, 0x05, 0x79, 0x32, 0xcd, 0x74, 0x4f, 0x04, 0xd5, 0xd5, 0xd2, 0xb0, 0xea, 0x24, 0xaa,
	0xb6, 0x30, 0xcc, 0x2d, 0xcb, 0x6b, 0x9c, 0x5b, 0x5c, 0x7b, 0xb5, 0x17, 0x55, 0x37, 0x44, 0xf1,
	0x23, 0x81, 0xbd, 0x2f, 0x2b, 0x74, 0xf7, 0x21, 0xbf, 0x77, 0xee, 0xba, 0x8f, 0x2c, 0xd9, 0x8e,
	0xac, 0x5f, 0x26, 0x7d, 0xb0, 0xcc, 0x21, 0x0c, 0x3d, 0x95, 0x8e, 0x38, 0x44, 0xa1, 0x23, 0x90,
	0x4f, 0x40, 0x60, 0x6d, 0x38, 0x3c, 0xa1, 0x43, 0xc9, 0x81, 0xf2, 0x5a, 0xc3, 0xd8, 0xc7, 0x20,
	0xc4, 0x16, 0x9c, 0x70, 0x6c, 0x42, 0x73, 0x2d, 0x8f, 0x21, 0x0b, 0xdf, 0x0b, 0xfb, 0x1d, 0x39,
	0x38, 0xa1, 0x37, 0xb8, 0x38, 0x83, 0x2c, 0xd8, 0xfc, 0x64, 0x3d, 0xc9, 0x57, 0x30, 0xa8, 0xac,
	0x27, 0x95, 0xe4, 0x62, 0x3a, 0x39, 0x79, 0x36, 0x7b, 0xf0, 0x22, 0x66, 0xa1, 0x4d, 0x73, 0x03,
	0x2f, 0xd4, 0x50, 0x59, 0xb1, 0x73, 0xa1, 0x23, 0x28, 0x72, 0xc8, 0xce, 0x6a, 0x9a, 0x97, 0xd5,
	0x0a, 0x43, 0xc7, 0x5d, 0x28, 0xda, 0x87, 0x13, 0x41, 0xf1, 0x1a, 0x9e, 0x6e, 0x82, 0x39, 0xbf,
	0xba, 0xc0, 0x0a, 0x17, 0x14, 0x62, 0xbe, 0xc5, 0xfb, 0xaf, 0x8d, 0x33, 0x6d, 0x3e, 0x1d, 0x3c,
	0xf9, 0x95, 0x40, 0xf6, 0xa1, 0xf5, 0x20, 0xdf, 0x02, 0x9c, 0xf2, 0xa5, 0xf2, 0xab, 0xfc, 0x9b,
	0xb9, 0xe3, 0xa3, 0x47, 0x64, 0xe7, 0xa6, 0xd8, 0x91, 0x1f, 0x61, 0x7f, 0xf3, 0x67, 0xde, 0xf5,
	0xc5, 0xa3, 0xde, 0xed, 0x0b, 0xfb, 0x43, 0xa8, 0x9b, 0x2a, 0x76, 0xe4, 0x1c, 0x8e, 0xb8, 0xf5,
	0x7d, 0x79, 0x8b, 0xb5, 0x89, 0x2b, 0x9c, 0x5f, 0xb1, 0xa7, 0xfc, 0x5f, 0x92, 0xdd, 0xaa, 0xff,
	0xd1, 0xbd, 0x1c, 0x31, 0xf3, 0xe6, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xb9, 0xc8, 0x29,
	0x97, 0x03, 0x00, 0x00,
}
