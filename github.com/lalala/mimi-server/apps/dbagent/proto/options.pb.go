// Code generated by protoc-gen-go. DO NOT EDIT.
// source: options.proto

package dbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FeedbackStatus int32

const (
	FeedbackStatus_wait       FeedbackStatus = 0
	FeedbackStatus_examine    FeedbackStatus = 1
	FeedbackStatus_process_ok FeedbackStatus = 2
	FeedbackStatus_ignore     FeedbackStatus = 3
)

var FeedbackStatus_name = map[int32]string{
	0: "wait",
	1: "examine",
	2: "process_ok",
	3: "ignore",
}
var FeedbackStatus_value = map[string]int32{
	"wait":       0,
	"examine":    1,
	"process_ok": 2,
	"ignore":     3,
}

func (x FeedbackStatus) String() string {
	return proto.EnumName(FeedbackStatus_name, int32(x))
}
func (FeedbackStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor19, []int{0} }

type Banner_TargetType int32

const (
	Banner_TargetType_All      Banner_TargetType = 0
	Banner_TargetType_News     Banner_TargetType = 1
	Banner_TargetType_Activity Banner_TargetType = 2
	Banner_TargetType_Link     Banner_TargetType = 3
)

var Banner_TargetType_name = map[int32]string{
	0: "TargetType_All",
	1: "TargetType_News",
	2: "TargetType_Activity",
	3: "TargetType_Link",
}
var Banner_TargetType_value = map[string]int32{
	"TargetType_All":      0,
	"TargetType_News":     1,
	"TargetType_Activity": 2,
	"TargetType_Link":     3,
}

func (x Banner_TargetType) String() string {
	return proto.EnumName(Banner_TargetType_name, int32(x))
}
func (Banner_TargetType) EnumDescriptor() ([]byte, []int) { return fileDescriptor19, []int{0, 0} }

type Banner_Location int32

const (
	Banner_Location_All      Banner_Location = 0
	Banner_Location_Home     Banner_Location = 1
	Banner_Location_Discover Banner_Location = 2
)

var Banner_Location_name = map[int32]string{
	0: "Location_All",
	1: "Location_Home",
	2: "Location_Discover",
}
var Banner_Location_value = map[string]int32{
	"Location_All":      0,
	"Location_Home":     1,
	"Location_Discover": 2,
}

func (x Banner_Location) String() string {
	return proto.EnumName(Banner_Location_name, int32(x))
}
func (Banner_Location) EnumDescriptor() ([]byte, []int) { return fileDescriptor19, []int{0, 1} }

type QueryClientBannerArg_Location int32

const (
	QueryClientBannerArg_Location_All      QueryClientBannerArg_Location = 0
	QueryClientBannerArg_Location_Home     QueryClientBannerArg_Location = 1
	QueryClientBannerArg_Location_Discover QueryClientBannerArg_Location = 2
)

var QueryClientBannerArg_Location_name = map[int32]string{
	0: "Location_All",
	1: "Location_Home",
	2: "Location_Discover",
}
var QueryClientBannerArg_Location_value = map[string]int32{
	"Location_All":      0,
	"Location_Home":     1,
	"Location_Discover": 2,
}

func (x QueryClientBannerArg_Location) String() string {
	return proto.EnumName(QueryClientBannerArg_Location_name, int32(x))
}
func (QueryClientBannerArg_Location) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor19, []int{3, 0}
}

type Banner struct {
	Id          int64             `protobuf:"varint,1,opt,name=id" json:"id"`
	Url         string            `protobuf:"bytes,2,opt,name=url" json:"url"`
	TargetLink  string            `protobuf:"bytes,3,opt,name=targetLink" json:"targetLink"`
	IsVisible   bool              `protobuf:"varint,4,opt,name=isVisible" json:"isVisible"`
	Description string            `protobuf:"bytes,5,opt,name=description" json:"description"`
	Created     int64             `protobuf:"varint,6,opt,name=created" json:"created"`
	Sort        int64             `protobuf:"varint,7,opt,name=sort" json:"sort"`
	TargetId    int64             `protobuf:"varint,8,opt,name=targetId" json:"targetId"`
	TargetType  Banner_TargetType `protobuf:"varint,9,opt,name=targetType,enum=dbproto.Banner_TargetType" json:"targetType"`
	Location    Banner_Location   `protobuf:"varint,10,opt,name=location,enum=dbproto.Banner_Location" json:"location"`
	Updated     int64             `protobuf:"varint,11,opt,name=updated" json:"updated"`
}

func (m *Banner) Reset()                    { *m = Banner{} }
func (m *Banner) String() string            { return proto.CompactTextString(m) }
func (*Banner) ProtoMessage()               {}
func (*Banner) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{0} }

func (m *Banner) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Banner) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Banner) GetTargetLink() string {
	if m != nil {
		return m.TargetLink
	}
	return ""
}

func (m *Banner) GetIsVisible() bool {
	if m != nil {
		return m.IsVisible
	}
	return false
}

func (m *Banner) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Banner) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Banner) GetSort() int64 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *Banner) GetTargetId() int64 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *Banner) GetTargetType() Banner_TargetType {
	if m != nil {
		return m.TargetType
	}
	return Banner_TargetType_All
}

func (m *Banner) GetLocation() Banner_Location {
	if m != nil {
		return m.Location
	}
	return Banner_Location_All
}

func (m *Banner) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type BannerId struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id"`
}

func (m *BannerId) Reset()                    { *m = BannerId{} }
func (m *BannerId) String() string            { return proto.CompactTextString(m) }
func (*BannerId) ProtoMessage()               {}
func (*BannerId) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{1} }

func (m *BannerId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type BannerList struct {
	List  []*Banner `protobuf:"bytes,1,rep,name=list" json:"list"`
	Total int64     `protobuf:"varint,2,opt,name=total" json:"total"`
}

func (m *BannerList) Reset()                    { *m = BannerList{} }
func (m *BannerList) String() string            { return proto.CompactTextString(m) }
func (*BannerList) ProtoMessage()               {}
func (*BannerList) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{2} }

func (m *BannerList) GetList() []*Banner {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *BannerList) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type QueryClientBannerArg struct {
	Location QueryClientBannerArg_Location `protobuf:"varint,1,opt,name=location,enum=dbproto.QueryClientBannerArg_Location" json:"location"`
}

func (m *QueryClientBannerArg) Reset()                    { *m = QueryClientBannerArg{} }
func (m *QueryClientBannerArg) String() string            { return proto.CompactTextString(m) }
func (*QueryClientBannerArg) ProtoMessage()               {}
func (*QueryClientBannerArg) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{3} }

func (m *QueryClientBannerArg) GetLocation() QueryClientBannerArg_Location {
	if m != nil {
		return m.Location
	}
	return QueryClientBannerArg_Location_All
}

type QueryBannerArg struct {
	Location Banner_Location `protobuf:"varint,1,opt,name=location,enum=dbproto.Banner_Location" json:"location"`
}

func (m *QueryBannerArg) Reset()                    { *m = QueryBannerArg{} }
func (m *QueryBannerArg) String() string            { return proto.CompactTextString(m) }
func (*QueryBannerArg) ProtoMessage()               {}
func (*QueryBannerArg) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{4} }

func (m *QueryBannerArg) GetLocation() Banner_Location {
	if m != nil {
		return m.Location
	}
	return Banner_Location_All
}

type Contact struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	Qq       string `protobuf:"bytes,2,opt,name=qq" json:"qq"`
	Wechat   string `protobuf:"bytes,3,opt,name=wechat" json:"wechat"`
	Email    string `protobuf:"bytes,4,opt,name=email" json:"email"`
	Telphone string `protobuf:"bytes,5,opt,name=telphone" json:"telphone"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{5} }

func (m *Contact) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Contact) GetQq() string {
	if m != nil {
		return m.Qq
	}
	return ""
}

func (m *Contact) GetWechat() string {
	if m != nil {
		return m.Wechat
	}
	return ""
}

func (m *Contact) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Contact) GetTelphone() string {
	if m != nil {
		return m.Telphone
	}
	return ""
}

type Feedback struct {
	Id         int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	Email      string `protobuf:"bytes,2,opt,name=email" json:"email"`
	Name       string `protobuf:"bytes,3,opt,name=name" json:"name"`
	Content    string `protobuf:"bytes,4,opt,name=content" json:"content"`
	Status     int32  `protobuf:"varint,5,opt,name=status" json:"status"`
	CreateTime int32  `protobuf:"varint,6,opt,name=create_time,json=createTime" json:"create_time"`
}

func (m *Feedback) Reset()                    { *m = Feedback{} }
func (m *Feedback) String() string            { return proto.CompactTextString(m) }
func (*Feedback) ProtoMessage()               {}
func (*Feedback) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{6} }

func (m *Feedback) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Feedback) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Feedback) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feedback) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Feedback) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Feedback) GetCreateTime() int32 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

type Faq struct {
	Id         int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	Title      string `protobuf:"bytes,2,opt,name=title" json:"title"`
	Content    string `protobuf:"bytes,3,opt,name=content" json:"content"`
	ContentUrl string `protobuf:"bytes,4,opt,name=contentUrl" json:"contentUrl"`
	Html       string `protobuf:"bytes,5,opt,name=html" json:"html"`
	IsVisible  bool   `protobuf:"varint,6,opt,name=isVisible" json:"isVisible"`
	CreateTime int64  `protobuf:"varint,7,opt,name=createTime" json:"createTime"`
	UpdateTime int64  `protobuf:"varint,8,opt,name=updateTime" json:"updateTime"`
}

func (m *Faq) Reset()                    { *m = Faq{} }
func (m *Faq) String() string            { return proto.CompactTextString(m) }
func (*Faq) ProtoMessage()               {}
func (*Faq) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{7} }

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

type FaqId struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id"`
}

func (m *FaqId) Reset()                    { *m = FaqId{} }
func (m *FaqId) String() string            { return proto.CompactTextString(m) }
func (*FaqId) ProtoMessage()               {}
func (*FaqId) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{8} }

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
func (*FaqList) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{9} }

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
	proto.RegisterType((*Banner)(nil), "dbproto.Banner")
	proto.RegisterType((*BannerId)(nil), "dbproto.BannerId")
	proto.RegisterType((*BannerList)(nil), "dbproto.BannerList")
	proto.RegisterType((*QueryClientBannerArg)(nil), "dbproto.QueryClientBannerArg")
	proto.RegisterType((*QueryBannerArg)(nil), "dbproto.QueryBannerArg")
	proto.RegisterType((*Contact)(nil), "dbproto.Contact")
	proto.RegisterType((*Feedback)(nil), "dbproto.Feedback")
	proto.RegisterType((*Faq)(nil), "dbproto.Faq")
	proto.RegisterType((*FaqId)(nil), "dbproto.FaqId")
	proto.RegisterType((*FaqList)(nil), "dbproto.FaqList")
	proto.RegisterEnum("dbproto.FeedbackStatus", FeedbackStatus_name, FeedbackStatus_value)
	proto.RegisterEnum("dbproto.Banner_TargetType", Banner_TargetType_name, Banner_TargetType_value)
	proto.RegisterEnum("dbproto.Banner_Location", Banner_Location_name, Banner_Location_value)
	proto.RegisterEnum("dbproto.QueryClientBannerArg_Location", QueryClientBannerArg_Location_name, QueryClientBannerArg_Location_value)
}

func init() { proto.RegisterFile("options.proto", fileDescriptor19) }

var fileDescriptor19 = []byte{
	// 682 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x4f, 0xdc, 0x30,
	0x10, 0x25, 0xc9, 0x7e, 0x64, 0x67, 0x21, 0x04, 0x43, 0x4b, 0x84, 0x2a, 0x1a, 0xa5, 0x52, 0xb5,
	0xea, 0x61, 0x0f, 0x94, 0x53, 0x6f, 0x40, 0x95, 0x16, 0x09, 0x55, 0x6a, 0x4a, 0x7b, 0x45, 0xde,
	0xc4, 0x02, 0x8b, 0x24, 0xde, 0xd8, 0x5e, 0x28, 0x3f, 0xa2, 0xe7, 0xde, 0xfb, 0xa7, 0xfa, 0x77,
	0xaa, 0xd8, 0x4e, 0x36, 0x2c, 0xab, 0xaa, 0x87, 0xde, 0xe6, 0xbd, 0xb1, 0x3d, 0xcf, 0xe3, 0x79,
	0x86, 0x2d, 0x36, 0x97, 0x94, 0x95, 0x62, 0x3a, 0xe7, 0x4c, 0x32, 0x34, 0xcc, 0x66, 0x2a, 0x88,
	0x7e, 0xf4, 0x60, 0x70, 0x8a, 0xcb, 0x92, 0x70, 0xe4, 0x81, 0x4d, 0xb3, 0xc0, 0x0a, 0xad, 0x89,
	0x93, 0xd8, 0x34, 0x43, 0x3e, 0x38, 0x0b, 0x9e, 0x07, 0x76, 0x68, 0x4d, 0x46, 0x49, 0x1d, 0xa2,
	0x43, 0x00, 0x89, 0xf9, 0x35, 0x91, 0x17, 0xb4, 0xbc, 0x0d, 0x1c, 0x95, 0xe8, 0x30, 0xe8, 0x05,
	0x8c, 0xa8, 0xf8, 0x46, 0x05, 0x9d, 0xe5, 0x24, 0xe8, 0x85, 0xd6, 0xc4, 0x4d, 0x96, 0x04, 0x0a,
	0x61, 0x9c, 0x11, 0x91, 0x72, 0xaa, 0x94, 0x04, 0x7d, 0xb5, 0xbd, 0x4b, 0xa1, 0x00, 0x86, 0x29,
	0x27, 0x58, 0x92, 0x2c, 0x18, 0x28, 0x19, 0x0d, 0x44, 0x08, 0x7a, 0x82, 0x71, 0x19, 0x0c, 0x15,
	0xad, 0x62, 0x74, 0x00, 0xae, 0xae, 0x7d, 0x9e, 0x05, 0xae, 0xe2, 0x5b, 0x8c, 0xde, 0x35, 0x4a,
	0x2f, 0x1f, 0xe6, 0x24, 0x18, 0x85, 0xd6, 0xc4, 0x3b, 0x3a, 0x98, 0x9a, 0x4b, 0x4f, 0xf5, 0x85,
	0xa7, 0x97, 0xed, 0x8a, 0xa4, 0xb3, 0x1a, 0x1d, 0x83, 0x9b, 0xb3, 0x14, 0x2b, 0x91, 0xa0, 0x76,
	0x06, 0xab, 0x3b, 0x2f, 0x4c, 0x3e, 0x69, 0x57, 0xd6, 0xda, 0x17, 0xf3, 0x4c, 0x69, 0x1f, 0x6b,
	0xed, 0x06, 0x46, 0x29, 0xc0, 0xb2, 0x12, 0x42, 0xe0, 0x2d, 0xd1, 0xd5, 0x49, 0x9e, 0xfb, 0x1b,
	0x68, 0x17, 0xb6, 0x3b, 0xdc, 0x27, 0x72, 0x2f, 0x7c, 0x0b, 0xed, 0xc3, 0x6e, 0x77, 0x61, 0x2a,
	0xe9, 0x1d, 0x95, 0x0f, 0xbe, 0xbd, 0xb2, 0xba, 0x6e, 0xbc, 0xef, 0x44, 0x31, 0xb8, 0x8d, 0x28,
	0xe4, 0xc3, 0x66, 0x13, 0x9b, 0x02, 0x3b, 0xb0, 0xd5, 0x32, 0x1f, 0x59, 0x41, 0x7c, 0x0b, 0x3d,
	0x83, 0x9d, 0x96, 0x7a, 0x4f, 0x45, 0xca, 0xee, 0x08, 0xf7, 0xed, 0xe8, 0x00, 0x5c, 0x7d, 0xc7,
	0xf3, 0x6c, 0x75, 0x20, 0xa2, 0x0f, 0x00, 0x3a, 0x77, 0x41, 0x85, 0x44, 0xaf, 0xa0, 0x97, 0x53,
	0x21, 0x03, 0x2b, 0x74, 0x26, 0xe3, 0xa3, 0xed, 0x95, 0x16, 0x25, 0x2a, 0x89, 0xf6, 0xa0, 0x2f,
	0x99, 0xc4, 0x7a, 0x8a, 0x9c, 0x44, 0x83, 0xe8, 0x97, 0x05, 0x7b, 0x9f, 0x17, 0x84, 0x3f, 0x9c,
	0xe5, 0x94, 0x94, 0x52, 0xef, 0x38, 0xe1, 0xd7, 0xe8, 0xb4, 0xd3, 0x7a, 0x4b, 0xb5, 0xfe, 0x75,
	0x7b, 0xee, 0xba, 0x0d, 0x6b, 0x1e, 0xe2, 0xbf, 0x75, 0x22, 0x06, 0x4f, 0x95, 0x5c, 0xaa, 0x3b,
	0x7e, 0xa2, 0xee, 0x1f, 0x06, 0x23, 0x12, 0x30, 0x3c, 0x63, 0xa5, 0xc4, 0xa9, 0x7c, 0xe2, 0x30,
	0x0f, 0xec, 0xaa, 0x32, 0x06, 0xb3, 0xab, 0x0a, 0x3d, 0x87, 0xc1, 0x3d, 0x49, 0x6f, 0xb0, 0x34,
	0xde, 0x32, 0xa8, 0xee, 0x22, 0x29, 0x30, 0xcd, 0x95, 0xa7, 0x46, 0x89, 0x06, 0x6a, 0xfe, 0x49,
	0x3e, 0xbf, 0x61, 0x25, 0x31, 0x66, 0x6a, 0x71, 0xf4, 0xd3, 0x02, 0x37, 0x26, 0x24, 0x9b, 0xe1,
	0xf4, 0xf6, 0x49, 0xd9, 0xf6, 0x38, 0xbb, 0x7b, 0x1c, 0x82, 0x5e, 0x89, 0x0b, 0x62, 0x4a, 0xab,
	0x58, 0x19, 0x92, 0x95, 0x92, 0x94, 0xd2, 0x94, 0x6e, 0x60, 0x2d, 0x55, 0x48, 0x2c, 0x17, 0x42,
	0x95, 0xee, 0x27, 0x06, 0xa1, 0x97, 0x30, 0xd6, 0x9e, 0xbd, 0x92, 0xb4, 0x20, 0xca, 0xc6, 0xfd,
	0x04, 0x34, 0x75, 0x49, 0x0b, 0x12, 0xfd, 0xb6, 0xc0, 0x89, 0x71, 0xb5, 0x4e, 0x94, 0xa4, 0x32,
	0x27, 0x8d, 0x28, 0x05, 0xba, 0x02, 0x9c, 0xc7, 0x02, 0x0e, 0x01, 0x4c, 0xf8, 0x95, 0x37, 0x8d,
	0xe9, 0x30, 0xf5, 0x75, 0x6e, 0x64, 0x91, 0x9b, 0xce, 0xa8, 0xf8, 0xf1, 0xff, 0x34, 0x58, 0xfd,
	0x9f, 0xea, 0x13, 0x5b, 0x9d, 0xe6, 0xa7, 0xe9, 0x30, 0x75, 0x5e, 0x5b, 0x5a, 0xe5, 0xf5, 0x8f,
	0xd3, 0x61, 0xa2, 0x7d, 0xe8, 0xc7, 0xb8, 0x5a, 0xe3, 0x9b, 0x13, 0x18, 0xc6, 0xb8, 0x52, 0xa6,
	0x09, 0x1f, 0x99, 0x66, 0xb3, 0x1d, 0x9f, 0x18, 0x57, 0x7f, 0x73, 0xcc, 0x9b, 0x33, 0xf0, 0x9a,
	0xe7, 0xfc, 0xa2, 0x1b, 0xed, 0x42, 0xef, 0x1e, 0x53, 0xe9, 0x6f, 0xa0, 0x31, 0x0c, 0xc9, 0x77,
	0x5c, 0xd0, 0xb2, 0x1e, 0x66, 0x0f, 0x60, 0xce, 0x59, 0x4a, 0x84, 0xb8, 0x62, 0xb7, 0xbe, 0x8d,
	0x00, 0x06, 0xf4, 0xba, 0x64, 0x9c, 0xf8, 0xce, 0x6c, 0xa0, 0x6a, 0xbd, 0xfd, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x91, 0xbd, 0x28, 0xb0, 0x0c, 0x06, 0x00, 0x00,
}
