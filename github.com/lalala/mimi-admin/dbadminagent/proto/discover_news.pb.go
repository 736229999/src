// Code generated by protoc-gen-go. DO NOT EDIT.
// source: discover_news.proto

package dbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type News struct {
	Id          int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	Title       string `protobuf:"bytes,2,opt,name=title" json:"title"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description"`
	Cover       string `protobuf:"bytes,4,opt,name=cover" json:"cover"`
	Content     string `protobuf:"bytes,5,opt,name=content" json:"content"`
	Html        string `protobuf:"bytes,6,opt,name=html" json:"html"`
	Author      string `protobuf:"bytes,7,opt,name=author" json:"author"`
	Created     int64  `protobuf:"varint,8,opt,name=created" json:"created"`
	NewsClass   int32  `protobuf:"varint,9,opt,name=newsClass" json:"newsClass"`
	IsVisible   bool   `protobuf:"varint,10,opt,name=isVisible" json:"isVisible"`
	Updated     int64  `protobuf:"varint,11,opt,name=updated" json:"updated"`
	PageViews   int64  `protobuf:"varint,12,opt,name=pageViews" json:"pageViews"`
}

func (m *News) Reset()                    { *m = News{} }
func (m *News) String() string            { return proto.CompactTextString(m) }
func (*News) ProtoMessage()               {}
func (*News) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

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

func (m *News) GetNewsClass() int32 {
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
	Title    string `protobuf:"bytes,1,opt,name=title" json:"title"`
	Author   string `protobuf:"bytes,2,opt,name=author" json:"author"`
	Start    int64  `protobuf:"varint,3,opt,name=start" json:"start"`
	End      int64  `protobuf:"varint,4,opt,name=end" json:"end"`
	Class    int32  `protobuf:"varint,5,opt,name=class" json:"class"`
	Page     int64  `protobuf:"varint,6,opt,name=page" json:"page"`
	PageSize int64  `protobuf:"varint,7,opt,name=pageSize" json:"pageSize"`
}

func (m *QueryNewsArg) Reset()                    { *m = QueryNewsArg{} }
func (m *QueryNewsArg) String() string            { return proto.CompactTextString(m) }
func (*QueryNewsArg) ProtoMessage()               {}
func (*QueryNewsArg) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

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

func (m *QueryNewsArg) GetClass() int32 {
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
	List  []*News `protobuf:"bytes,1,rep,name=list" json:"list"`
	Total int64   `protobuf:"varint,2,opt,name=total" json:"total"`
}

func (m *NewsList) Reset()                    { *m = NewsList{} }
func (m *NewsList) String() string            { return proto.CompactTextString(m) }
func (*NewsList) ProtoMessage()               {}
func (*NewsList) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

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

type NewsId struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id"`
}

func (m *NewsId) Reset()                    { *m = NewsId{} }
func (m *NewsId) String() string            { return proto.CompactTextString(m) }
func (*NewsId) ProtoMessage()               {}
func (*NewsId) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *NewsId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryNewsOfSelect struct {
	KeyWord string `protobuf:"bytes,1,opt,name=keyWord" json:"keyWord"`
}

func (m *QueryNewsOfSelect) Reset()                    { *m = QueryNewsOfSelect{} }
func (m *QueryNewsOfSelect) String() string            { return proto.CompactTextString(m) }
func (*QueryNewsOfSelect) ProtoMessage()               {}
func (*QueryNewsOfSelect) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *QueryNewsOfSelect) GetKeyWord() string {
	if m != nil {
		return m.KeyWord
	}
	return ""
}

func init() {
	proto.RegisterType((*News)(nil), "dbproto.News")
	proto.RegisterType((*QueryNewsArg)(nil), "dbproto.QueryNewsArg")
	proto.RegisterType((*NewsList)(nil), "dbproto.NewsList")
	proto.RegisterType((*NewsId)(nil), "dbproto.NewsId")
	proto.RegisterType((*QueryNewsOfSelect)(nil), "dbproto.QueryNewsOfSelect")
}

func init() { proto.RegisterFile("discover_news.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x41, 0xae, 0xd3, 0x30,
	0x10, 0x86, 0xe5, 0x38, 0x49, 0xd3, 0xe9, 0x03, 0x81, 0x79, 0x42, 0x16, 0x62, 0x11, 0xb2, 0xca,
	0x86, 0x2e, 0xe0, 0x04, 0xe8, 0xad, 0x90, 0x10, 0x08, 0x57, 0x2a, 0x4b, 0x94, 0xc6, 0x43, 0x6b,
	0x11, 0xe2, 0xc8, 0x76, 0xa9, 0xca, 0x91, 0x38, 0x01, 0xc7, 0x43, 0x9e, 0xa4, 0x4d, 0x61, 0x95,
	0xf9, 0xff, 0x89, 0xc7, 0xf3, 0x7f, 0x32, 0x3c, 0xd3, 0xc6, 0xb7, 0xf6, 0x27, 0xba, 0xaf, 0x3d,
	0x9e, 0xfc, 0x7a, 0x70, 0x36, 0x58, 0xb1, 0xd0, 0x3b, 0x2a, 0xaa, 0x3f, 0x09, 0xa4, 0x1f, 0xf1,
	0xe4, 0xc5, 0x63, 0x48, 0x8c, 0x96, 0xac, 0x64, 0x35, 0x57, 0x89, 0xd1, 0xe2, 0x1e, 0xb2, 0x60,
	0x42, 0x87, 0x32, 0x29, 0x59, 0xbd, 0x54, 0xa3, 0x10, 0x25, 0xac, 0x34, 0xfa, 0xd6, 0x99, 0x21,
	0x18, 0xdb, 0x4b, 0x4e, 0xbd, 0x5b, 0x2b, 0x9e, 0xa3, 0xdb, 0x64, 0x3a, 0x9e, 0x23, 0x21, 0x24,
	0x2c, 0x5a, 0xdb, 0x07, 0xec, 0x83, 0xcc, 0xc8, 0xbf, 0x48, 0x21, 0x20, 0x3d, 0x84, 0x1f, 0x9d,
	0xcc, 0xc9, 0xa6, 0x5a, 0x3c, 0x87, 0xbc, 0x39, 0x86, 0x83, 0x75, 0x72, 0x41, 0xee, 0xa4, 0x68,
	0x8a, 0xc3, 0x26, 0xa0, 0x96, 0x05, 0x2d, 0x7a, 0x91, 0xe2, 0x25, 0x2c, 0x63, 0xba, 0x87, 0xae,
	0xf1, 0x5e, 0x2e, 0x4b, 0x56, 0x67, 0x6a, 0x36, 0x62, 0xd7, 0xf8, 0xad, 0xf1, 0x66, 0xd7, 0xa1,
	0x84, 0x92, 0xd5, 0x85, 0x9a, 0x8d, 0x38, 0xf5, 0x38, 0x68, 0x9a, 0xba, 0x1a, 0xa7, 0x4e, 0x32,
	0x9e, 0x1b, 0x9a, 0x3d, 0x6e, 0x0d, 0x9e, 0xbc, 0xbc, 0xa3, 0xde, 0x6c, 0x54, 0xbf, 0x19, 0xdc,
	0x7d, 0x3e, 0xa2, 0x3b, 0x47, 0x7e, 0xef, 0xdc, 0x7e, 0x46, 0xc6, 0x6e, 0x91, 0xcd, 0x61, 0x92,
	0x7f, 0xc2, 0xdc, 0x43, 0xe6, 0x43, 0xe3, 0x02, 0x41, 0xe4, 0x6a, 0x14, 0xe2, 0x09, 0x70, 0xec,
	0x35, 0xc1, 0xe3, 0x2a, 0x96, 0x04, 0x94, 0x62, 0x65, 0x14, 0x6b, 0x14, 0x11, 0x5b, 0xdc, 0x84,
	0xb0, 0x71, 0x45, 0xb5, 0x78, 0x01, 0x45, 0xfc, 0x6e, 0xcc, 0x2f, 0x24, 0x70, 0x5c, 0x5d, 0x75,
	0xf5, 0x00, 0x45, 0x5c, 0xf3, 0x83, 0xf1, 0x41, 0xbc, 0x82, 0xb4, 0x33, 0x3e, 0x48, 0x56, 0xf2,
	0x7a, 0xf5, 0xe6, 0xd1, 0x7a, 0x7a, 0x0b, 0xeb, 0xf8, 0x83, 0xa2, 0x16, 0x45, 0xb1, 0xa1, 0xe9,
	0x68, 0x67, 0xae, 0x46, 0x51, 0x49, 0xc8, 0xe3, 0x3f, 0xef, 0xf5, 0xff, 0xaf, 0xa5, 0x7a, 0x0d,
	0x4f, 0xaf, 0x28, 0x3e, 0x7d, 0xdb, 0x60, 0x87, 0x6d, 0x88, 0x60, 0xbf, 0xe3, 0xf9, 0x8b, 0x75,
	0x7a, 0x22, 0x72, 0x91, 0xbb, 0x9c, 0x2e, 0x7c, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x53,
	0xae, 0x1c, 0x9c, 0x02, 0x00, 0x00,
}
