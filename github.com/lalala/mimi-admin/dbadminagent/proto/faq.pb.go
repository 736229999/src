// Code generated by protoc-gen-go. DO NOT EDIT.
// source: faq.proto

package dbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Faq struct {
	Id         int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	Title      string `protobuf:"bytes,2,opt,name=title" json:"title"`
	Content    string `protobuf:"bytes,3,opt,name=content" json:"content"`
	ContentUrl string `protobuf:"bytes,4,opt,name=contentUrl" json:"contentUrl"`
	Html       string `protobuf:"bytes,5,opt,name=html" json:"html"`
	IsVisible  bool   `protobuf:"varint,6,opt,name=isVisible" json:"isVisible"`
	CreateTime int64  `protobuf:"varint,7,opt,name=createTime" json:"createTime"`
	UpdateTime int64  `protobuf:"varint,8,opt,name=updateTime" json:"updateTime"`
	Sort       int32  `protobuf:"varint,9,opt,name=sort" json:"sort"`
}

func (m *Faq) Reset()                    { *m = Faq{} }
func (m *Faq) String() string            { return proto.CompactTextString(m) }
func (*Faq) ProtoMessage()               {}
func (*Faq) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *Faq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Faq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Faq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Faq) GetContentUrl() string {
	if m != nil {
		return m.ContentUrl
	}
	return ""
}

func (m *Faq) GetHtml() string {
	if m != nil {
		return m.Html
	}
	return ""
}

func (m *Faq) GetIsVisible() bool {
	if m != nil {
		return m.IsVisible
	}
	return false
}

func (m *Faq) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Faq) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *Faq) GetSort() int32 {
	if m != nil {
		return m.Sort
	}
	return 0
}

type QueryFaqArg struct {
	Title    string `protobuf:"bytes,1,opt,name=title" json:"title"`
	Page     int32  `protobuf:"varint,2,opt,name=page" json:"page"`
	PageSize int32  `protobuf:"varint,3,opt,name=pageSize" json:"pageSize"`
}

func (m *QueryFaqArg) Reset()                    { *m = QueryFaqArg{} }
func (m *QueryFaqArg) String() string            { return proto.CompactTextString(m) }
func (*QueryFaqArg) ProtoMessage()               {}
func (*QueryFaqArg) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *QueryFaqArg) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *QueryFaqArg) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *QueryFaqArg) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type FaqId struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id"`
}

func (m *FaqId) Reset()                    { *m = FaqId{} }
func (m *FaqId) String() string            { return proto.CompactTextString(m) }
func (*FaqId) ProtoMessage()               {}
func (*FaqId) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

func (m *FaqId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FaqList struct {
	List  []*Faq `protobuf:"bytes,1,rep,name=list" json:"list"`
	Total int64  `protobuf:"varint,2,opt,name=total" json:"total"`
}

func (m *FaqList) Reset()                    { *m = FaqList{} }
func (m *FaqList) String() string            { return proto.CompactTextString(m) }
func (*FaqList) ProtoMessage()               {}
func (*FaqList) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{3} }

func (m *FaqList) GetList() []*Faq {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *FaqList) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*Faq)(nil), "dbproto.Faq")
	proto.RegisterType((*QueryFaqArg)(nil), "dbproto.QueryFaqArg")
	proto.RegisterType((*FaqId)(nil), "dbproto.FaqId")
	proto.RegisterType((*FaqList)(nil), "dbproto.FaqList")
}

func init() { proto.RegisterFile("faq.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4f, 0x84, 0x30,
	0x10, 0x85, 0x53, 0xa0, 0x0b, 0xcc, 0x1a, 0x0f, 0x8d, 0x89, 0x8d, 0x31, 0xa6, 0xe1, 0xc4, 0x89,
	0x83, 0xfe, 0x82, 0xbd, 0x90, 0x98, 0x78, 0xb1, 0xab, 0xde, 0x8b, 0xd4, 0xb5, 0x49, 0x77, 0xa1,
	0xa5, 0x7b, 0xd0, 0xbf, 0xec, 0x9f, 0x30, 0x0c, 0xbb, 0x42, 0x3c, 0xf1, 0xde, 0xfb, 0x48, 0x67,
	0xde, 0x40, 0xfe, 0xa1, 0x5c, 0xd5, 0xfb, 0x2e, 0x74, 0x2c, 0x6d, 0x1b, 0x14, 0xc5, 0x0f, 0x81,
	0xb8, 0x56, 0x8e, 0x5d, 0x42, 0x64, 0x5a, 0x4e, 0x04, 0x29, 0x63, 0x19, 0x99, 0x96, 0x5d, 0x01,
	0x0d, 0x26, 0x58, 0xcd, 0x23, 0x41, 0xca, 0x5c, 0x4e, 0x86, 0x71, 0x48, 0xdf, 0xbb, 0x43, 0xd0,
	0x87, 0xc0, 0x63, 0xcc, 0xcf, 0x96, 0xdd, 0x01, 0x9c, 0xe4, 0xab, 0xb7, 0x3c, 0x41, 0xb8, 0x48,
	0x18, 0x83, 0xe4, 0x33, 0xec, 0x2d, 0xa7, 0x48, 0x50, 0xb3, 0x5b, 0xc8, 0xcd, 0xf0, 0x66, 0x06,
	0xd3, 0x58, 0xcd, 0x57, 0x82, 0x94, 0x99, 0x9c, 0x03, 0x7c, 0xd1, 0x6b, 0x15, 0xf4, 0x8b, 0xd9,
	0x6b, 0x9e, 0xe2, 0x66, 0x8b, 0x64, 0xe4, 0xc7, 0xbe, 0x3d, 0xf3, 0x6c, 0xe2, 0x73, 0x32, 0x4e,
	0x1c, 0x3a, 0x1f, 0x78, 0x2e, 0x48, 0x49, 0x25, 0xea, 0x62, 0x0b, 0xeb, 0xe7, 0xa3, 0xf6, 0x5f,
	0xb5, 0x72, 0x1b, 0xbf, 0x9b, 0x4b, 0x92, 0x65, 0x49, 0x06, 0x49, 0xaf, 0x76, 0x53, 0x73, 0x2a,
	0x51, 0xb3, 0x1b, 0xc8, 0xc6, 0xef, 0xd6, 0x7c, 0x6b, 0x6c, 0x4e, 0xe5, 0x9f, 0x2f, 0xae, 0x81,
	0xd6, 0xca, 0x3d, 0xb6, 0xff, 0x6f, 0x58, 0x6c, 0x20, 0xad, 0x95, 0x7b, 0x32, 0x43, 0x60, 0x02,
	0x12, 0x6b, 0x86, 0xc0, 0x89, 0x88, 0xcb, 0xf5, 0xfd, 0x45, 0x75, 0x3a, 0x7f, 0x55, 0x2b, 0x27,
	0x91, 0xe0, 0x2e, 0x5d, 0x50, 0x16, 0xc7, 0xc6, 0x72, 0x32, 0xcd, 0x0a, 0x7f, 0x7b, 0xf8, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xaf, 0x0a, 0x27, 0xb1, 0xbb, 0x01, 0x00, 0x00,
}