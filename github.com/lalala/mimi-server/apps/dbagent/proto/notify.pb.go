// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notify.proto

package dbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NotifyType int32

const (
	NotifyType_Announce NotifyType = 0
	NotifyType_Remind   NotifyType = 1
	NotifyType_Message  NotifyType = 2
	NotifyType_All      NotifyType = 3
)

var NotifyType_name = map[int32]string{
	0: "Announce",
	1: "Remind",
	2: "Message",
	3: "All",
}
var NotifyType_value = map[string]int32{
	"Announce": 0,
	"Remind":   1,
	"Message":  2,
	"All":      3,
}

func (x NotifyType) String() string {
	return proto.EnumName(NotifyType_name, int32(x))
}
func (NotifyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor17, []int{0} }

type Notify struct {
	Id         int64      `protobuf:"varint,1,opt,name=id" json:"id"`
	Content    string     `protobuf:"bytes,2,opt,name=content" json:"content"`
	Type       NotifyType `protobuf:"varint,3,opt,name=type,enum=dbproto.NotifyType" json:"type"`
	Target     int64      `protobuf:"varint,4,opt,name=target" json:"target"`
	TargetType string     `protobuf:"bytes,5,opt,name=targetType" json:"targetType"`
	Action     string     `protobuf:"bytes,6,opt,name=action" json:"action"`
	Sender     int64      `protobuf:"varint,7,opt,name=sender" json:"sender"`
	Created    int64      `protobuf:"varint,8,opt,name=created" json:"created"`
	Updated    int64      `protobuf:"varint,9,opt,name=updated" json:"updated"`
}

func (m *Notify) Reset()                    { *m = Notify{} }
func (m *Notify) String() string            { return proto.CompactTextString(m) }
func (*Notify) ProtoMessage()               {}
func (*Notify) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{0} }

func (m *Notify) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Notify) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Notify) GetType() NotifyType {
	if m != nil {
		return m.Type
	}
	return NotifyType_Announce
}

func (m *Notify) GetTarget() int64 {
	if m != nil {
		return m.Target
	}
	return 0
}

func (m *Notify) GetTargetType() string {
	if m != nil {
		return m.TargetType
	}
	return ""
}

func (m *Notify) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Notify) GetSender() int64 {
	if m != nil {
		return m.Sender
	}
	return 0
}

func (m *Notify) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Notify) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type UserNotify struct {
	Account    int64      `protobuf:"varint,1,opt,name=account" json:"account"`
	Notify     int64      `protobuf:"varint,2,opt,name=notify" json:"notify"`
	IsRead     bool       `protobuf:"varint,3,opt,name=isRead" json:"isRead"`
	NotifyType NotifyType `protobuf:"varint,4,opt,name=notifyType,enum=dbproto.NotifyType" json:"notifyType"`
	Created    int64      `protobuf:"varint,5,opt,name=created" json:"created"`
}

func (m *UserNotify) Reset()                    { *m = UserNotify{} }
func (m *UserNotify) String() string            { return proto.CompactTextString(m) }
func (*UserNotify) ProtoMessage()               {}
func (*UserNotify) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{1} }

func (m *UserNotify) GetAccount() int64 {
	if m != nil {
		return m.Account
	}
	return 0
}

func (m *UserNotify) GetNotify() int64 {
	if m != nil {
		return m.Notify
	}
	return 0
}

func (m *UserNotify) GetIsRead() bool {
	if m != nil {
		return m.IsRead
	}
	return false
}

func (m *UserNotify) GetNotifyType() NotifyType {
	if m != nil {
		return m.NotifyType
	}
	return NotifyType_Announce
}

func (m *UserNotify) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

type QueryUserMissedArg struct {
	Account int64      `protobuf:"varint,1,opt,name=account" json:"account"`
	Type    NotifyType `protobuf:"varint,2,opt,name=type,enum=dbproto.NotifyType" json:"type"`
}

func (m *QueryUserMissedArg) Reset()                    { *m = QueryUserMissedArg{} }
func (m *QueryUserMissedArg) String() string            { return proto.CompactTextString(m) }
func (*QueryUserMissedArg) ProtoMessage()               {}
func (*QueryUserMissedArg) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{2} }

func (m *QueryUserMissedArg) GetAccount() int64 {
	if m != nil {
		return m.Account
	}
	return 0
}

func (m *QueryUserMissedArg) GetType() NotifyType {
	if m != nil {
		return m.Type
	}
	return NotifyType_Announce
}

type QueryUserMissedRes struct {
	Notices []*Notify `protobuf:"bytes,1,rep,name=notices" json:"notices"`
}

func (m *QueryUserMissedRes) Reset()                    { *m = QueryUserMissedRes{} }
func (m *QueryUserMissedRes) String() string            { return proto.CompactTextString(m) }
func (*QueryUserMissedRes) ProtoMessage()               {}
func (*QueryUserMissedRes) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{3} }

func (m *QueryUserMissedRes) GetNotices() []*Notify {
	if m != nil {
		return m.Notices
	}
	return nil
}

type QueryUserNotifyArg struct {
	Account    int64      `protobuf:"varint,1,opt,name=account" json:"account"`
	NotifyType NotifyType `protobuf:"varint,2,opt,name=notifyType,enum=dbproto.NotifyType" json:"notifyType"`
	Page       int64      `protobuf:"varint,3,opt,name=page" json:"page"`
	PageSize   int64      `protobuf:"varint,4,opt,name=pageSize" json:"pageSize"`
}

func (m *QueryUserNotifyArg) Reset()                    { *m = QueryUserNotifyArg{} }
func (m *QueryUserNotifyArg) String() string            { return proto.CompactTextString(m) }
func (*QueryUserNotifyArg) ProtoMessage()               {}
func (*QueryUserNotifyArg) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{4} }

func (m *QueryUserNotifyArg) GetAccount() int64 {
	if m != nil {
		return m.Account
	}
	return 0
}

func (m *QueryUserNotifyArg) GetNotifyType() NotifyType {
	if m != nil {
		return m.NotifyType
	}
	return NotifyType_Announce
}

func (m *QueryUserNotifyArg) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *QueryUserNotifyArg) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type QueryUserNotifyRes struct {
	UserNoticeInfos []*UserNotifyInfo `protobuf:"bytes,1,rep,name=userNoticeInfos" json:"userNoticeInfos"`
}

func (m *QueryUserNotifyRes) Reset()                    { *m = QueryUserNotifyRes{} }
func (m *QueryUserNotifyRes) String() string            { return proto.CompactTextString(m) }
func (*QueryUserNotifyRes) ProtoMessage()               {}
func (*QueryUserNotifyRes) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{5} }

func (m *QueryUserNotifyRes) GetUserNoticeInfos() []*UserNotifyInfo {
	if m != nil {
		return m.UserNoticeInfos
	}
	return nil
}

type ReadUserNotifyArg struct {
	NotifyId  int64 `protobuf:"varint,1,opt,name=notifyId" json:"notifyId"`
	AccountId int64 `protobuf:"varint,2,opt,name=accountId" json:"accountId"`
}

func (m *ReadUserNotifyArg) Reset()                    { *m = ReadUserNotifyArg{} }
func (m *ReadUserNotifyArg) String() string            { return proto.CompactTextString(m) }
func (*ReadUserNotifyArg) ProtoMessage()               {}
func (*ReadUserNotifyArg) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{6} }

func (m *ReadUserNotifyArg) GetNotifyId() int64 {
	if m != nil {
		return m.NotifyId
	}
	return 0
}

func (m *ReadUserNotifyArg) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

type UserNotifyInfo struct {
	UserNotify *UserNotify `protobuf:"bytes,1,opt,name=userNotify" json:"userNotify"`
	Notify     *Notify     `protobuf:"bytes,2,opt,name=notify" json:"notify"`
}

func (m *UserNotifyInfo) Reset()                    { *m = UserNotifyInfo{} }
func (m *UserNotifyInfo) String() string            { return proto.CompactTextString(m) }
func (*UserNotifyInfo) ProtoMessage()               {}
func (*UserNotifyInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{7} }

func (m *UserNotifyInfo) GetUserNotify() *UserNotify {
	if m != nil {
		return m.UserNotify
	}
	return nil
}

func (m *UserNotifyInfo) GetNotify() *Notify {
	if m != nil {
		return m.Notify
	}
	return nil
}

// 创建一条公告通知参数
type CreateAnnounceArg struct {
	Content string `protobuf:"bytes,1,opt,name=content" json:"content"`
	Sender  int64  `protobuf:"varint,2,opt,name=sender" json:"sender"`
}

func (m *CreateAnnounceArg) Reset()                    { *m = CreateAnnounceArg{} }
func (m *CreateAnnounceArg) String() string            { return proto.CompactTextString(m) }
func (*CreateAnnounceArg) ProtoMessage()               {}
func (*CreateAnnounceArg) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{8} }

func (m *CreateAnnounceArg) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CreateAnnounceArg) GetSender() int64 {
	if m != nil {
		return m.Sender
	}
	return 0
}

// 创建一条提醒记录参数
type CreateRemindArg struct {
	Target     int64  `protobuf:"varint,1,opt,name=target" json:"target"`
	TargetType string `protobuf:"bytes,2,opt,name=targetType" json:"targetType"`
	Action     string `protobuf:"bytes,3,opt,name=action" json:"action"`
	Sender     int64  `protobuf:"varint,4,opt,name=sender" json:"sender"`
	Content    string `protobuf:"bytes,5,opt,name=content" json:"content"`
}

func (m *CreateRemindArg) Reset()                    { *m = CreateRemindArg{} }
func (m *CreateRemindArg) String() string            { return proto.CompactTextString(m) }
func (*CreateRemindArg) ProtoMessage()               {}
func (*CreateRemindArg) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{9} }

func (m *CreateRemindArg) GetTarget() int64 {
	if m != nil {
		return m.Target
	}
	return 0
}

func (m *CreateRemindArg) GetTargetType() string {
	if m != nil {
		return m.TargetType
	}
	return ""
}

func (m *CreateRemindArg) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CreateRemindArg) GetSender() int64 {
	if m != nil {
		return m.Sender
	}
	return 0
}

func (m *CreateRemindArg) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// 创建一条私信记录参数
type CreateMessageArg struct {
	Content  string `protobuf:"bytes,1,opt,name=content" json:"content"`
	Sender   int64  `protobuf:"varint,2,opt,name=sender" json:"sender"`
	Receiver int64  `protobuf:"varint,3,opt,name=receiver" json:"receiver"`
}

func (m *CreateMessageArg) Reset()                    { *m = CreateMessageArg{} }
func (m *CreateMessageArg) String() string            { return proto.CompactTextString(m) }
func (*CreateMessageArg) ProtoMessage()               {}
func (*CreateMessageArg) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{10} }

func (m *CreateMessageArg) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CreateMessageArg) GetSender() int64 {
	if m != nil {
		return m.Sender
	}
	return 0
}

func (m *CreateMessageArg) GetReceiver() int64 {
	if m != nil {
		return m.Receiver
	}
	return 0
}

func init() {
	proto.RegisterType((*Notify)(nil), "dbproto.Notify")
	proto.RegisterType((*UserNotify)(nil), "dbproto.UserNotify")
	proto.RegisterType((*QueryUserMissedArg)(nil), "dbproto.QueryUserMissedArg")
	proto.RegisterType((*QueryUserMissedRes)(nil), "dbproto.QueryUserMissedRes")
	proto.RegisterType((*QueryUserNotifyArg)(nil), "dbproto.QueryUserNotifyArg")
	proto.RegisterType((*QueryUserNotifyRes)(nil), "dbproto.QueryUserNotifyRes")
	proto.RegisterType((*ReadUserNotifyArg)(nil), "dbproto.ReadUserNotifyArg")
	proto.RegisterType((*UserNotifyInfo)(nil), "dbproto.UserNotifyInfo")
	proto.RegisterType((*CreateAnnounceArg)(nil), "dbproto.CreateAnnounceArg")
	proto.RegisterType((*CreateRemindArg)(nil), "dbproto.CreateRemindArg")
	proto.RegisterType((*CreateMessageArg)(nil), "dbproto.CreateMessageArg")
	proto.RegisterEnum("dbproto.NotifyType", NotifyType_name, NotifyType_value)
}

func init() { proto.RegisterFile("notify.proto", fileDescriptor17) }

var fileDescriptor17 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x65, 0xed, 0x34, 0x76, 0x26, 0x55, 0xe2, 0x2e, 0x12, 0xac, 0x22, 0x84, 0x22, 0x5f, 0x1a,
	0x38, 0xe4, 0xd0, 0x5e, 0x91, 0x50, 0x84, 0x38, 0xe4, 0x10, 0x24, 0x16, 0x10, 0x57, 0x1c, 0x7b,
	0x1a, 0x59, 0x2a, 0x6b, 0xcb, 0xbb, 0x46, 0x0a, 0x7f, 0x83, 0xdf, 0xc0, 0x6f, 0xe4, 0x8a, 0xf6,
	0xc3, 0x5f, 0x2d, 0x0d, 0x88, 0x53, 0xfc, 0x66, 0x26, 0xb3, 0x6f, 0xde, 0xcc, 0x83, 0x73, 0x51,
	0xa8, 0xfc, 0xe6, 0xb8, 0x2e, 0xab, 0x42, 0x15, 0x34, 0xc8, 0xf6, 0xe6, 0x63, 0x01, 0xfb, 0x44,
	0xa2, 0x0d, 0xc6, 0xbf, 0x08, 0x8c, 0xdf, 0x99, 0x2a, 0x3a, 0x03, 0x2f, 0xcf, 0x18, 0x59, 0x92,
	0x95, 0xcf, 0xbd, 0x3c, 0xa3, 0x0c, 0x82, 0xb4, 0x10, 0x0a, 0x85, 0x62, 0xde, 0x92, 0xac, 0x26,
	0xbc, 0x81, 0xf4, 0x12, 0x46, 0xea, 0x58, 0x22, 0xf3, 0x97, 0x64, 0x35, 0xbb, 0x7a, 0xbc, 0x76,
	0x8d, 0xd7, 0xb6, 0xd1, 0xc7, 0x63, 0x89, 0xdc, 0x14, 0xd0, 0x27, 0x30, 0x56, 0x49, 0x75, 0x40,
	0xc5, 0x46, 0xa6, 0xad, 0x43, 0xf4, 0x39, 0x80, 0xfd, 0xd2, 0xb5, 0xec, 0xcc, 0x74, 0xef, 0x45,
	0xf4, 0xff, 0x92, 0x54, 0xe5, 0x85, 0x60, 0x63, 0x93, 0x73, 0x48, 0xc7, 0x25, 0x8a, 0x0c, 0x2b,
	0x16, 0xd8, 0x7e, 0x16, 0x19, 0xaa, 0x15, 0x26, 0x0a, 0x33, 0x16, 0x9a, 0x44, 0x03, 0x75, 0xa6,
	0x2e, 0x33, 0x93, 0x99, 0xd8, 0x8c, 0x83, 0xf1, 0x4f, 0x02, 0xf0, 0x49, 0x62, 0xe5, 0xa6, 0x67,
	0x10, 0x24, 0x69, 0x5a, 0xd4, 0x42, 0x39, 0x09, 0x1a, 0xa8, 0x1f, 0xb5, 0x3a, 0x1a, 0x19, 0x7c,
	0xee, 0x90, 0x8e, 0xe7, 0x92, 0x63, 0x92, 0x19, 0x1d, 0x42, 0xee, 0x10, 0xbd, 0x06, 0x10, 0xad,
	0x10, 0x66, 0xf0, 0x07, 0x34, 0xea, 0x95, 0xf5, 0x27, 0x38, 0x1b, 0x4c, 0x10, 0x7f, 0x06, 0xfa,
	0xbe, 0xc6, 0xea, 0xa8, 0xb9, 0xee, 0x72, 0x29, 0x31, 0xdb, 0x54, 0x87, 0x13, 0x74, 0x9b, 0xe5,
	0x78, 0x7f, 0x59, 0x4e, 0xfc, 0xfa, 0x5e, 0x63, 0x8e, 0x92, 0xbe, 0x80, 0x40, 0xd3, 0x4a, 0x51,
	0x32, 0xb2, 0xf4, 0x57, 0xd3, 0xab, 0xf9, 0x9d, 0x0e, 0xbc, 0xc9, 0xc7, 0x3f, 0x48, 0xaf, 0x83,
	0x4d, 0x9e, 0xa6, 0x36, 0x54, 0xc6, 0xfb, 0x37, 0x65, 0x28, 0x8c, 0xca, 0xe4, 0x60, 0x8f, 0xcd,
	0xe7, 0xe6, 0x9b, 0x2e, 0x20, 0xd4, 0xbf, 0x1f, 0xf2, 0xef, 0xe8, 0x2e, 0xab, 0xc5, 0x03, 0xbd,
	0x1c, 0x63, 0x94, 0x74, 0x03, 0xf3, 0xda, 0x05, 0x52, 0xdc, 0x8a, 0x9b, 0xa2, 0x19, 0xef, 0x69,
	0xfb, 0x7e, 0xf7, 0x07, 0x9d, 0xe7, 0x77, 0xeb, 0xe3, 0x1d, 0x5c, 0xe8, 0xfd, 0x0e, 0x87, 0x5d,
	0x40, 0x68, 0xb9, 0x6e, 0x1b, 0xeb, 0xb4, 0x98, 0x3e, 0x83, 0x89, 0x9b, 0x7c, 0x9b, 0xb9, 0xdb,
	0xe9, 0x02, 0xb1, 0x80, 0xd9, 0xf0, 0x45, 0x2d, 0x4f, 0xdd, 0x46, 0x4c, 0xb7, 0x69, 0x4f, 0x9e,
	0xde, 0x3c, 0xbd, 0x32, 0x7a, 0x39, 0xb8, 0xce, 0x3f, 0xac, 0xcb, 0xa5, 0xe3, 0xb7, 0x70, 0xf1,
	0xc6, 0x9c, 0xd4, 0x46, 0x88, 0xa2, 0x16, 0x29, 0xba, 0x5d, 0x35, 0x1e, 0x27, 0x43, 0x8f, 0x77,
	0x56, 0xf3, 0xfa, 0x56, 0xd3, 0x4b, 0x9f, 0xdb, 0x3e, 0x1c, 0xbf, 0xe6, 0xc2, 0x1c, 0x63, 0x67,
	0x73, 0x72, 0xc2, 0xe6, 0xde, 0x09, 0x9b, 0xfb, 0x0f, 0xd8, 0x7c, 0x74, 0xcf, 0xe6, 0x8e, 0xed,
	0xd9, 0x80, 0x6d, 0xfc, 0x05, 0x22, 0x4b, 0x6a, 0x87, 0x52, 0x26, 0x87, 0xff, 0x9b, 0x4d, 0x2f,
	0xb3, 0xc2, 0x14, 0xf3, 0x6f, 0x58, 0xb9, 0x73, 0x6b, 0xf1, 0xcb, 0x57, 0x00, 0xdd, 0x81, 0xd2,
	0x73, 0x08, 0x1b, 0x19, 0xa3, 0x47, 0x14, 0x60, 0x6c, 0xc5, 0x88, 0x08, 0x9d, 0x42, 0xe0, 0x38,
	0x44, 0x1e, 0x0d, 0xc0, 0xdf, 0xdc, 0xde, 0x46, 0xfe, 0x7e, 0x6c, 0x56, 0x72, 0xfd, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x4a, 0xab, 0x97, 0x3f, 0x92, 0x05, 0x00, 0x00,
}
