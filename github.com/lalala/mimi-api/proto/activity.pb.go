// Code generated by protoc-gen-go. DO NOT EDIT.
// source: activity.proto

/*
Package apiproto is a generated protocol buffer package.

It is generated from these files:
	activity.proto
	base.proto
	buycai.proto
	credits.proto
	discover.proto
	error.proto
	football.proto
	fund.proto
	gift.proto
	notify.proto
	opencai.proto
	options.proto
	recharge.proto
	sms.proto
	user.proto
	withdraw.proto

It has these top-level messages:
	Task
	Activity
	ActivityList
	ActivityAccount
	UserBaseInfo
	ActivtyDetail
	Nil
	StringList
	IntValue
	StringValue
	Response
	Balls
	BuycaiScheme
	IssueMultiple
	BuycaiOrder
	OnSellInfo
	ValidTicketsRequest
	SaleIssue
	BuycaiInfo
	VendorOrderHistoryRequest
	UserOrderHistoryRequest
	VendorOrderRecord
	UserOrderRecord
	VendorOrderHistory
	UserOrderHistory
	BuycaiIssueInfo
	BuycaiUserOrder
	BuycaiVendorOrder
	CreditsTask
	CreditsTaskInfo
	News
	QueryNewsArg
	NewsList
	NewsId
	GameInfo
	FbGameOpencai
	QueryGameInfoRes
	QueryOpencaiRes
	FundInfo
	Bankcard
	AddBankcardRequest
	BuycaiTicket
	BuycaiTicketInfo
	FundChangeRecord
	FundHistoryPage
	FundHistory
	ExchangeRequest
	ExchangeReply
	UserInviteInfoReply
	GiftItem
	RegistGiftPackage
	Notify
	UserNotify
	QueryUserNotifyArg
	ReadNotifyArg
	UserNotifyInfo
	QueryUserNotifyRes
	BonusDetail
	OpenDetail
	OpenInfo
	LotteryCollection
	LatestOpen
	DayHistory
	History
	OpenInfoDigest
	DigestHistory
	LotteryOptions
	WiningList
	Banner
	BannerList
	HomeParams
	Contact
	Faq
	FaqId
	FaqList
	RechargeRequest
	RechargeResponse
	AlipayRechargeReply
	JdPayRechargeReply
	SmsRequest
	UserInfo
	RegistRequest
	LoginRequest
	LoginReply
	ForgotPwdRequest
	ForgotPayPwdRequest
	VerifyPayPwdRequest
	SetPwdRequest
	ResetPwdRequest
	BindPhoneRequest
	ChangePhoneRequest
	BindWeixinRequest
	BindQQRequest
	SetIconRequest
	SetNicknameRequest
	SetSexRequest
	SetBoolRequest
	SetIntRequest
	AuthRealnameRequest
	DailyCheckReply
	WithdrawInfo
	WithdrawRequest
	WithdrawProgress
*/
package apiproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 参与活动要完成的任务
type Task struct {
	Id         int64   `protobuf:"varint,1,opt,name=id" json:"id"`
	Name       string  `protobuf:"bytes,2,opt,name=name" json:"name"`
	Des        string  `protobuf:"bytes,3,opt,name=des" json:"des"`
	IsComplete bool    `protobuf:"varint,4,opt,name=is_complete,json=isComplete" json:"is_complete"`
	Type       string  `protobuf:"bytes,5,opt,name=type" json:"type"`
	Money      float32 `protobuf:"fixed32,6,opt,name=money" json:"money"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Task) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetDes() string {
	if m != nil {
		return m.Des
	}
	return ""
}

func (m *Task) GetIsComplete() bool {
	if m != nil {
		return m.IsComplete
	}
	return false
}

func (m *Task) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Task) GetMoney() float32 {
	if m != nil {
		return m.Money
	}
	return 0
}

// 活动
type Activity struct {
	Id        int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	Title     string `protobuf:"bytes,2,opt,name=title" json:"title"`
	Des       string `protobuf:"bytes,3,opt,name=des" json:"des"`
	Logo      string `protobuf:"bytes,4,opt,name=logo" json:"logo"`
	Num       int64  `protobuf:"varint,5,opt,name=num" json:"num"`
	LeftNum   int64  `protobuf:"varint,6,opt,name=left_num,json=leftNum" json:"left_num"`
	Starttime int64  `protobuf:"varint,7,opt,name=starttime" json:"starttime"`
	Endtime   int64  `protobuf:"varint,8,opt,name=endtime" json:"endtime"`
	HasJoin   bool   `protobuf:"varint,9,opt,name=has_join,json=hasJoin" json:"has_join"`
}

func (m *Activity) Reset()                    { *m = Activity{} }
func (m *Activity) String() string            { return proto.CompactTextString(m) }
func (*Activity) ProtoMessage()               {}
func (*Activity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Activity) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Activity) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Activity) GetDes() string {
	if m != nil {
		return m.Des
	}
	return ""
}

func (m *Activity) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *Activity) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *Activity) GetLeftNum() int64 {
	if m != nil {
		return m.LeftNum
	}
	return 0
}

func (m *Activity) GetStarttime() int64 {
	if m != nil {
		return m.Starttime
	}
	return 0
}

func (m *Activity) GetEndtime() int64 {
	if m != nil {
		return m.Endtime
	}
	return 0
}

func (m *Activity) GetHasJoin() bool {
	if m != nil {
		return m.HasJoin
	}
	return false
}

// 活动列表
type ActivityList struct {
	List []*Activity `protobuf:"bytes,1,rep,name=list" json:"list"`
}

func (m *ActivityList) Reset()                    { *m = ActivityList{} }
func (m *ActivityList) String() string            { return proto.CompactTextString(m) }
func (*ActivityList) ProtoMessage()               {}
func (*ActivityList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ActivityList) GetList() []*Activity {
	if m != nil {
		return m.List
	}
	return nil
}

// 包含活动id和用户账户id，用于查询活动详情
type ActivityAccount struct {
	ActivityId int64 `protobuf:"varint,1,opt,name=activity_id,json=activityId" json:"activity_id"`
	AccountId  int64 `protobuf:"varint,2,opt,name=account_id,json=accountId" json:"account_id"`
}

func (m *ActivityAccount) Reset()                    { *m = ActivityAccount{} }
func (m *ActivityAccount) String() string            { return proto.CompactTextString(m) }
func (*ActivityAccount) ProtoMessage()               {}
func (*ActivityAccount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ActivityAccount) GetActivityId() int64 {
	if m != nil {
		return m.ActivityId
	}
	return 0
}

func (m *ActivityAccount) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

// 用户基本信息
type UserBaseInfo struct {
	AccountId int64  `protobuf:"varint,1,opt,name=account_id,json=accountId" json:"account_id"`
	Icon      string `protobuf:"bytes,2,opt,name=icon" json:"icon"`
	Nickname  string `protobuf:"bytes,3,opt,name=nickname" json:"nickname"`
}

func (m *UserBaseInfo) Reset()                    { *m = UserBaseInfo{} }
func (m *UserBaseInfo) String() string            { return proto.CompactTextString(m) }
func (*UserBaseInfo) ProtoMessage()               {}
func (*UserBaseInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserBaseInfo) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *UserBaseInfo) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *UserBaseInfo) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

// 活动详情
type ActivtyDetail struct {
	Activity *Activity       `protobuf:"bytes,1,opt,name=activity" json:"activity"`
	Gainers  []*UserBaseInfo `protobuf:"bytes,2,rep,name=gainers" json:"gainers"`
	TaskList []*Task         `protobuf:"bytes,3,rep,name=task_list,json=taskList" json:"task_list"`
}

func (m *ActivtyDetail) Reset()                    { *m = ActivtyDetail{} }
func (m *ActivtyDetail) String() string            { return proto.CompactTextString(m) }
func (*ActivtyDetail) ProtoMessage()               {}
func (*ActivtyDetail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ActivtyDetail) GetActivity() *Activity {
	if m != nil {
		return m.Activity
	}
	return nil
}

func (m *ActivtyDetail) GetGainers() []*UserBaseInfo {
	if m != nil {
		return m.Gainers
	}
	return nil
}

func (m *ActivtyDetail) GetTaskList() []*Task {
	if m != nil {
		return m.TaskList
	}
	return nil
}

func init() {
	proto.RegisterType((*Task)(nil), "apiproto.Task")
	proto.RegisterType((*Activity)(nil), "apiproto.Activity")
	proto.RegisterType((*ActivityList)(nil), "apiproto.ActivityList")
	proto.RegisterType((*ActivityAccount)(nil), "apiproto.ActivityAccount")
	proto.RegisterType((*UserBaseInfo)(nil), "apiproto.UserBaseInfo")
	proto.RegisterType((*ActivtyDetail)(nil), "apiproto.ActivtyDetail")
}

func init() { proto.RegisterFile("activity.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xd1, 0x8a, 0xd4, 0x30,
	0x14, 0x25, 0x6d, 0x67, 0xda, 0xde, 0x5d, 0x47, 0xb9, 0x2c, 0x12, 0x45, 0xb1, 0xf4, 0x41, 0x0a,
	0xc2, 0x20, 0x0a, 0xbe, 0xaf, 0xfa, 0x32, 0x22, 0x82, 0x41, 0x1f, 0xa5, 0xc4, 0x36, 0xbb, 0x1b,
	0xa7, 0x4d, 0x86, 0x49, 0x46, 0x98, 0x2f, 0xf0, 0x2f, 0xfc, 0x2c, 0xbf, 0x47, 0x72, 0x3b, 0x99,
	0x91, 0x75, 0xdf, 0xce, 0x3d, 0xf7, 0xdc, 0xcb, 0xc9, 0xc9, 0x85, 0x85, 0xec, 0xbc, 0xfe, 0xa9,
	0xfd, 0x7e, 0xb9, 0xd9, 0x5a, 0x6f, 0xb1, 0x90, 0x1b, 0x4d, 0xa8, 0xfe, 0xc5, 0x20, 0xfb, 0x22,
	0xdd, 0x1a, 0x17, 0x90, 0xe8, 0x9e, 0xb3, 0x8a, 0x35, 0xa9, 0x48, 0x74, 0x8f, 0x08, 0x99, 0x91,
	0xa3, 0xe2, 0x49, 0xc5, 0x9a, 0x52, 0x10, 0xc6, 0x07, 0x90, 0xf6, 0xca, 0xf1, 0x94, 0xa8, 0x00,
	0xf1, 0x19, 0x9c, 0x69, 0xd7, 0x76, 0x76, 0xdc, 0x0c, 0xca, 0x2b, 0x9e, 0x55, 0xac, 0x29, 0x04,
	0x68, 0xf7, 0xee, 0xc0, 0x84, 0x35, 0x7e, 0xbf, 0x51, 0x7c, 0x36, 0xad, 0x09, 0x18, 0x2f, 0x60,
	0x36, 0x5a, 0xa3, 0xf6, 0x7c, 0x5e, 0xb1, 0x26, 0x11, 0x53, 0x51, 0xff, 0x61, 0x50, 0x5c, 0x1e,
	0x6c, 0xfe, 0xe7, 0xe6, 0x02, 0x66, 0x5e, 0xfb, 0x21, 0xda, 0x99, 0x8a, 0x3b, 0xfc, 0x20, 0x64,
	0x83, 0xbd, 0xb6, 0x64, 0xa4, 0x14, 0x84, 0x83, 0xca, 0xec, 0x46, 0x72, 0x90, 0x8a, 0x00, 0xf1,
	0x11, 0x14, 0x83, 0xba, 0xf2, 0x6d, 0xa0, 0xe7, 0x44, 0xe7, 0xa1, 0xfe, 0xb4, 0x1b, 0xf1, 0x09,
	0x94, 0xce, 0xcb, 0xad, 0xf7, 0x7a, 0x54, 0x3c, 0xa7, 0xde, 0x89, 0x40, 0x0e, 0xb9, 0x32, 0x3d,
	0xf5, 0x8a, 0x69, 0xee, 0x50, 0x86, 0x95, 0x37, 0xd2, 0xb5, 0x3f, 0xac, 0x36, 0xbc, 0xa4, 0x14,
	0xf2, 0x1b, 0xe9, 0x3e, 0x58, 0x6d, 0xea, 0x37, 0x70, 0x1e, 0xdf, 0xf5, 0x51, 0x3b, 0x8f, 0xcf,
	0x21, 0x1b, 0xb4, 0xf3, 0x9c, 0x55, 0x69, 0x73, 0xf6, 0x0a, 0x97, 0xf1, 0x2f, 0x96, 0x51, 0x25,
	0xa8, 0x5f, 0x7f, 0x86, 0xfb, 0x91, 0xb9, 0xec, 0x3a, 0xbb, 0x33, 0x3e, 0xc4, 0x1d, 0x7f, 0xb2,
	0x3d, 0xe6, 0x03, 0x91, 0x5a, 0xf5, 0xf8, 0x14, 0x40, 0x4e, 0xda, 0xd0, 0x4f, 0x26, 0xff, 0x07,
	0x66, 0xd5, 0xd7, 0xdf, 0xe0, 0xfc, 0xab, 0x53, 0xdb, 0xb7, 0xd2, 0xa9, 0x95, 0xb9, 0xb2, 0xb7,
	0xe4, 0xec, 0x96, 0x3c, 0xa4, 0xa9, 0x3b, 0x6b, 0xe2, 0x0d, 0x04, 0x8c, 0x8f, 0xa1, 0x30, 0xba,
	0x5b, 0xd3, 0x6d, 0x4c, 0xc1, 0x1f, 0xeb, 0xfa, 0x37, 0x83, 0x7b, 0x64, 0xd9, 0xef, 0xdf, 0x2b,
	0x2f, 0xf5, 0x80, 0x4b, 0x28, 0xa2, 0x3b, 0x5a, 0x7f, 0xf7, 0x7b, 0x8f, 0x1a, 0x7c, 0x09, 0xf9,
	0xb5, 0xd4, 0x46, 0x6d, 0x1d, 0x4f, 0x28, 0x9e, 0x87, 0x27, 0xf9, 0xbf, 0xce, 0x45, 0x94, 0xe1,
	0x0b, 0x28, 0xbd, 0x74, 0xeb, 0x96, 0x22, 0x4d, 0x69, 0x66, 0x71, 0x9a, 0x09, 0xa7, 0x2d, 0x8a,
	0x20, 0x08, 0xd1, 0x7f, 0x9f, 0x13, 0xfb, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x73,
	0x2a, 0xc3, 0x10, 0x03, 0x00, 0x00,
}