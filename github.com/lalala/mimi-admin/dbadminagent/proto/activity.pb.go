// Code generated by protoc-gen-go. DO NOT EDIT.
// source: activity.proto

/*
Package dbproto is a generated protocol buffer package.

It is generated from these files:
	activity.proto
	admin.proto
	agent_admin.proto
	banner.proto
	base.proto
	buycai.proto
	buycai_user_order.proto
	buycai_vendor_order.proto
	cdkey_batch.proto
	contact.proto
	data_statistics.proto
	discover_news.proto
	faq.proto
	feedback.proto
	football.proto
	gift_package.proto
	gift_template.proto
	log.proto
	lottery_options.proto
	opencai.proto
	order.proto
	privilege.proto
	recharge_history.proto
	recharge_order.proto
	role.proto
	usercenter.proto
	winning.proto
	withdraw.proto

It has these top-level messages:
	Activity
	ActivityList
	ActivityReply
	ActivityReplyList
	ActivityDetail
	Task
	TaskTypeList
	TaskList
	TaskReplyList
	TemplateList
	AdminUserInfoArg
	AdminUserList
	AdminUserInfoReply
	AdminPrivilegesReply
	AdminPrivilegesList
	AdminPrivileges
	Banner
	BannerId
	BannerList
	QueryBannerArg
	Nil
	BoolValue
	IntValue
	StringValue
	BuycaiOptions
	BuycaiOptionsReply
	BuycaiOptionsIssue
	BuycaiOptionsUpdateIssue
	PlayTimeSettings
	PlayTimeSettingsList
	BuycaiStatistics
	BuycaiStatisticsList
	BuycaiUserOrder
	UserOrderList
	BuycaiVendorOrder
	BuycaiScheme
	Balls
	IssueMultiple
	CdKeyBatchArg
	GiftCdkeyArg
	CdkeyListArg
	CdkeyListReply
	CdkeyListArgReply
	CdkeyBatch
	CdkeyReply
	CdkeyDetail
	Contact
	StatisticsOrderAndIncome
	News
	QueryNewsArg
	NewsList
	NewsId
	QueryNewsOfSelect
	Faq
	QueryFaqArg
	FaqId
	FaqList
	Feedback
	FeedbackList
	GameInfo
	PlayOdds
	FbGameresult
	FbLeagueInfo
	FbTeamInfo
	QueryFbGameArg
	QueryOpencaiArg
	QueryFbTeamArg
	QueryFbLeagueArg
	FbOpencai
	FbGameList
	FbLeagueList
	FbTeamList
	FbOpencaiList
	GiftPackageArg
	GiftList
	Ticket
	Gift
	GiftPackageRequest
	GiftListByType
	GiftTemplate
	GiftTemplateList
	GiftTemplateContent
	TicketBox
	CreditsBox
	Tickets
	Log
	LogReply
	LotteryOptions
	LotteryOptionsList
	LotteryWinningNo
	LotteryWinningNoList
	WinningLotteryList
	WinningLottery
	WinningListByLottery
	UserOrder
	Privilege
	PrivilegeChildrenList
	PrivilegeList
	RechargeHistory
	FundHistory
	RechargeOrder
	RechargeOrderList
	RechargeStatistics
	Role
	RoleList
	UsercenterList
	Userinfo
	UserStatisticsNum
	Winning
	WinningList
	WithdrawApply
	QueryWithdrawApplyArg
	WithdrawApplyId
	WithdrawApplyList
	UpdateWAStatusArg
	ClaimWAArg
	CheckWAArg
	CheckWARes
	WithdrawAuditAuth
	WithdrawAuditAuthList
	QueryWithdrawAuditAuthArg
	WithdrawTransfer
	WithdrawTransferList
	QueryWithdrawTransferArg
*/
package dbproto

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

type ActivityDetail_HasJoin int32

const (
	ActivityDetail_already ActivityDetail_HasJoin = 0
	ActivityDetail_has_not ActivityDetail_HasJoin = 1
)

var ActivityDetail_HasJoin_name = map[int32]string{
	0: "already",
	1: "has_not",
}
var ActivityDetail_HasJoin_value = map[string]int32{
	"already": 0,
	"has_not": 1,
}

func (x ActivityDetail_HasJoin) String() string {
	return proto.EnumName(ActivityDetail_HasJoin_name, int32(x))
}
func (ActivityDetail_HasJoin) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

// 用于添加活动
type Activity struct {
	Title       string  `protobuf:"bytes,1,opt,name=title" json:"title"`
	Des         string  `protobuf:"bytes,2,opt,name=des" json:"des"`
	Logo        string  `protobuf:"bytes,3,opt,name=logo" json:"logo"`
	Num         int64   `protobuf:"varint,4,opt,name=num" json:"num"`
	PackageId   int64   `protobuf:"varint,5,opt,name=package_id,json=packageId" json:"package_id"`
	Starttime   int64   `protobuf:"varint,6,opt,name=starttime" json:"starttime"`
	Endtime     int64   `protobuf:"varint,7,opt,name=endtime" json:"endtime"`
	TaskLists   []int64 `protobuf:"varint,8,rep,packed,name=taskLists" json:"taskLists"`
	PackageName string  `protobuf:"bytes,9,opt,name=package_name,json=packageName" json:"package_name"`
	Id          int64   `protobuf:"varint,10,opt,name=id" json:"id"`
}

func (m *Activity) Reset()                    { *m = Activity{} }
func (m *Activity) String() string            { return proto.CompactTextString(m) }
func (*Activity) ProtoMessage()               {}
func (*Activity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

func (m *Activity) GetPackageId() int64 {
	if m != nil {
		return m.PackageId
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

func (m *Activity) GetTaskLists() []int64 {
	if m != nil {
		return m.TaskLists
	}
	return nil
}

func (m *Activity) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *Activity) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// 活动列表
type ActivityList struct {
	ActivityList []*Activity `protobuf:"bytes,1,rep,name=activity_list,json=activityList" json:"activity_list"`
}

func (m *ActivityList) Reset()                    { *m = ActivityList{} }
func (m *ActivityList) String() string            { return proto.CompactTextString(m) }
func (*ActivityList) ProtoMessage()               {}
func (*ActivityList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ActivityList) GetActivityList() []*Activity {
	if m != nil {
		return m.ActivityList
	}
	return nil
}

// 返回给管理后台用的
type ActivityReply struct {
	Id    int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	Title string `protobuf:"bytes,2,opt,name=title" json:"title"`
	Des   string `protobuf:"bytes,3,opt,name=des" json:"des"`
	Logo  string `protobuf:"bytes,4,opt,name=logo" json:"logo"`
	Num   int64  `protobuf:"varint,5,opt,name=num" json:"num"`
	//    int64 left_num = 6;
	PackageName string `protobuf:"bytes,6,opt,name=package_name,json=packageName" json:"package_name"`
	Starttime   int64  `protobuf:"varint,7,opt,name=starttime" json:"starttime"`
	Endtime     int64  `protobuf:"varint,8,opt,name=endtime" json:"endtime"`
	CreateAdmin string `protobuf:"bytes,9,opt,name=create_admin,json=createAdmin" json:"create_admin"`
	CreateTime  int64  `protobuf:"varint,10,opt,name=create_time,json=createTime" json:"create_time"`
}

func (m *ActivityReply) Reset()                    { *m = ActivityReply{} }
func (m *ActivityReply) String() string            { return proto.CompactTextString(m) }
func (*ActivityReply) ProtoMessage()               {}
func (*ActivityReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ActivityReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ActivityReply) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ActivityReply) GetDes() string {
	if m != nil {
		return m.Des
	}
	return ""
}

func (m *ActivityReply) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *ActivityReply) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *ActivityReply) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *ActivityReply) GetStarttime() int64 {
	if m != nil {
		return m.Starttime
	}
	return 0
}

func (m *ActivityReply) GetEndtime() int64 {
	if m != nil {
		return m.Endtime
	}
	return 0
}

func (m *ActivityReply) GetCreateAdmin() string {
	if m != nil {
		return m.CreateAdmin
	}
	return ""
}

func (m *ActivityReply) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

// 返回活动列表
type ActivityReplyList struct {
	Total int64            `protobuf:"varint,1,opt,name=total" json:"total"`
	Size  int64            `protobuf:"varint,2,opt,name=size" json:"size"`
	Page  int64            `protobuf:"varint,3,opt,name=page" json:"page"`
	List  []*ActivityReply `protobuf:"bytes,4,rep,name=list" json:"list"`
}

func (m *ActivityReplyList) Reset()                    { *m = ActivityReplyList{} }
func (m *ActivityReplyList) String() string            { return proto.CompactTextString(m) }
func (*ActivityReplyList) ProtoMessage()               {}
func (*ActivityReplyList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ActivityReplyList) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ActivityReplyList) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ActivityReplyList) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ActivityReplyList) GetList() []*ActivityReply {
	if m != nil {
		return m.List
	}
	return nil
}

// 某个活动详情
type ActivityDetail struct {
	Id         int64    `protobuf:"varint,1,opt,name=id" json:"id"`
	Title      string   `protobuf:"bytes,2,opt,name=title" json:"title"`
	Des        string   `protobuf:"bytes,3,opt,name=des" json:"des"`
	Logo       string   `protobuf:"bytes,4,opt,name=logo" json:"logo"`
	Num        int64    `protobuf:"varint,5,opt,name=num" json:"num"`
	LeftNum    int64    `protobuf:"varint,6,opt,name=left_num,json=leftNum" json:"left_num"`
	Starttime  int64    `protobuf:"varint,7,opt,name=starttime" json:"starttime"`
	Endtime    int64    `protobuf:"varint,8,opt,name=endtime" json:"endtime"`
	TaskList   []*Task  `protobuf:"bytes,9,rep,name=task_list,json=taskList" json:"task_list"`
	GetUserImg []string `protobuf:"bytes,10,rep,name=get_user_img,json=getUserImg" json:"get_user_img"`
}

func (m *ActivityDetail) Reset()                    { *m = ActivityDetail{} }
func (m *ActivityDetail) String() string            { return proto.CompactTextString(m) }
func (*ActivityDetail) ProtoMessage()               {}
func (*ActivityDetail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ActivityDetail) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ActivityDetail) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ActivityDetail) GetDes() string {
	if m != nil {
		return m.Des
	}
	return ""
}

func (m *ActivityDetail) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *ActivityDetail) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *ActivityDetail) GetLeftNum() int64 {
	if m != nil {
		return m.LeftNum
	}
	return 0
}

func (m *ActivityDetail) GetStarttime() int64 {
	if m != nil {
		return m.Starttime
	}
	return 0
}

func (m *ActivityDetail) GetEndtime() int64 {
	if m != nil {
		return m.Endtime
	}
	return 0
}

func (m *ActivityDetail) GetTaskList() []*Task {
	if m != nil {
		return m.TaskList
	}
	return nil
}

func (m *ActivityDetail) GetGetUserImg() []string {
	if m != nil {
		return m.GetUserImg
	}
	return nil
}

// 返回给管理后台用的
type Task struct {
	Id       int64   `protobuf:"varint,1,opt,name=id" json:"id"`
	Name     string  `protobuf:"bytes,2,opt,name=name" json:"name"`
	Des      string  `protobuf:"bytes,3,opt,name=des" json:"des"`
	Addtime  string  `protobuf:"bytes,4,opt,name=addtime" json:"addtime"`
	IsFinish int64   `protobuf:"varint,5,opt,name=is_finish,json=isFinish" json:"is_finish"`
	Type     string  `protobuf:"bytes,6,opt,name=type" json:"type"`
	Money    float32 `protobuf:"fixed32,7,opt,name=money" json:"money"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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

func (m *Task) GetAddtime() string {
	if m != nil {
		return m.Addtime
	}
	return ""
}

func (m *Task) GetIsFinish() int64 {
	if m != nil {
		return m.IsFinish
	}
	return 0
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

type TaskTypeList struct {
	TypeList []string `protobuf:"bytes,1,rep,name=type_list,json=typeList" json:"type_list"`
}

func (m *TaskTypeList) Reset()                    { *m = TaskTypeList{} }
func (m *TaskTypeList) String() string            { return proto.CompactTextString(m) }
func (*TaskTypeList) ProtoMessage()               {}
func (*TaskTypeList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TaskTypeList) GetTypeList() []string {
	if m != nil {
		return m.TypeList
	}
	return nil
}

// 任务列表（不带分页）
type TaskList struct {
	List []*Task `protobuf:"bytes,1,rep,name=list" json:"list"`
}

func (m *TaskList) Reset()                    { *m = TaskList{} }
func (m *TaskList) String() string            { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()               {}
func (*TaskList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TaskList) GetList() []*Task {
	if m != nil {
		return m.List
	}
	return nil
}

// 返回任务列表（带分页）
type TaskReplyList struct {
	Total int64   `protobuf:"varint,1,opt,name=total" json:"total"`
	Size  int64   `protobuf:"varint,2,opt,name=size" json:"size"`
	Page  int64   `protobuf:"varint,3,opt,name=page" json:"page"`
	List  []*Task `protobuf:"bytes,4,rep,name=list" json:"list"`
}

func (m *TaskReplyList) Reset()                    { *m = TaskReplyList{} }
func (m *TaskReplyList) String() string            { return proto.CompactTextString(m) }
func (*TaskReplyList) ProtoMessage()               {}
func (*TaskReplyList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TaskReplyList) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *TaskReplyList) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *TaskReplyList) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *TaskReplyList) GetList() []*Task {
	if m != nil {
		return m.List
	}
	return nil
}

// 礼包模板列表（添加活动时要用）
type TemplateList struct {
	List []*GiftTemplate `protobuf:"bytes,1,rep,name=list" json:"list"`
}

func (m *TemplateList) Reset()                    { *m = TemplateList{} }
func (m *TemplateList) String() string            { return proto.CompactTextString(m) }
func (*TemplateList) ProtoMessage()               {}
func (*TemplateList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TemplateList) GetList() []*GiftTemplate {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*Activity)(nil), "dbproto.Activity")
	proto.RegisterType((*ActivityList)(nil), "dbproto.ActivityList")
	proto.RegisterType((*ActivityReply)(nil), "dbproto.ActivityReply")
	proto.RegisterType((*ActivityReplyList)(nil), "dbproto.ActivityReplyList")
	proto.RegisterType((*ActivityDetail)(nil), "dbproto.ActivityDetail")
	proto.RegisterType((*Task)(nil), "dbproto.Task")
	proto.RegisterType((*TaskTypeList)(nil), "dbproto.TaskTypeList")
	proto.RegisterType((*TaskList)(nil), "dbproto.TaskList")
	proto.RegisterType((*TaskReplyList)(nil), "dbproto.TaskReplyList")
	proto.RegisterType((*TemplateList)(nil), "dbproto.TemplateList")
	proto.RegisterEnum("dbproto.ActivityDetail_HasJoin", ActivityDetail_HasJoin_name, ActivityDetail_HasJoin_value)
}

func init() { proto.RegisterFile("activity.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xf1, 0x47, 0x6b, 0x7b, 0xea, 0x54, 0xed, 0xf2, 0xa1, 0x85, 0x82, 0x70, 0xcc, 0x25,
	0x14, 0xd1, 0x03, 0x48, 0x48, 0x1c, 0x2b, 0xa1, 0x42, 0x11, 0xea, 0xc1, 0x0a, 0x67, 0x6b, 0x5b,
	0x6f, 0xdc, 0x55, 0xfd, 0x25, 0xef, 0x16, 0xc9, 0x48, 0x48, 0x3c, 0x00, 0xaf, 0xc0, 0x5b, 0xf0,
	0x80, 0x68, 0x67, 0xbd, 0x29, 0x21, 0x11, 0x07, 0x04, 0xb7, 0x99, 0xff, 0xec, 0xcc, 0x78, 0x7e,
	0x33, 0x09, 0xec, 0xb2, 0x0b, 0x25, 0x3e, 0x09, 0x35, 0x1c, 0x75, 0x7d, 0xab, 0x5a, 0x12, 0x14,
	0xe7, 0x68, 0x3c, 0xb8, 0x5d, 0x8a, 0x85, 0xca, 0x15, 0xaf, 0xbb, 0x8a, 0x29, 0x6e, 0xa2, 0xe9,
	0x57, 0x17, 0xc2, 0xe3, 0x31, 0x81, 0xdc, 0x81, 0x2d, 0x25, 0x54, 0xc5, 0xa9, 0x93, 0x38, 0xb3,
	0x28, 0x33, 0x0e, 0xd9, 0x03, 0xaf, 0xe0, 0x92, 0xba, 0xa8, 0x69, 0x93, 0x10, 0xf0, 0xab, 0xb6,
	0x6c, 0xa9, 0x87, 0x12, 0xda, 0xfa, 0x55, 0x73, 0x5d, 0x53, 0x3f, 0x71, 0x66, 0x5e, 0xa6, 0x4d,
	0xf2, 0x08, 0xa0, 0x63, 0x17, 0x57, 0xac, 0xe4, 0xb9, 0x28, 0xe8, 0x16, 0x06, 0xa2, 0x51, 0x39,
	0x2d, 0xc8, 0x43, 0x88, 0xa4, 0x62, 0xbd, 0x52, 0xa2, 0xe6, 0x74, 0xdb, 0x44, 0x97, 0x02, 0xa1,
	0x10, 0xf0, 0xa6, 0xc0, 0x58, 0x80, 0x31, 0xeb, 0xea, 0x3c, 0xc5, 0xe4, 0xd5, 0x07, 0x21, 0x95,
	0xa4, 0x61, 0xe2, 0xe9, 0xbc, 0xa5, 0x40, 0xa6, 0x10, 0xdb, 0xa6, 0x0d, 0xab, 0x39, 0x8d, 0xf0,
	0x13, 0x77, 0x46, 0xed, 0x8c, 0xd5, 0x9c, 0xec, 0x82, 0x2b, 0x0a, 0x0a, 0x58, 0xd5, 0x15, 0x45,
	0x7a, 0x02, 0xb1, 0x25, 0xa0, 0x6b, 0x90, 0x57, 0x30, 0xb1, 0x08, 0xf3, 0x4a, 0x48, 0x45, 0x9d,
	0xc4, 0x9b, 0xed, 0xbc, 0xd8, 0x3f, 0x1a, 0x41, 0x1e, 0xd9, 0xd7, 0x59, 0xcc, 0x7e, 0xc9, 0x4b,
	0xbf, 0xb9, 0x30, 0x59, 0x86, 0x78, 0x57, 0x0d, 0x63, 0x27, 0xc7, 0x76, 0xba, 0xe1, 0xeb, 0x6e,
	0xe0, 0xeb, 0xad, 0xf3, 0xf5, 0xd7, 0xf9, 0x6e, 0xdd, 0xf0, 0xfd, 0x7d, 0xd4, 0xed, 0xf5, 0x51,
	0x57, 0x18, 0x07, 0x7f, 0x60, 0x1c, 0xae, 0x32, 0x9e, 0x42, 0x7c, 0xd1, 0x73, 0xa6, 0x78, 0xce,
	0x8a, 0x5a, 0x34, 0x96, 0xa2, 0xd1, 0x8e, 0xb5, 0x44, 0x1e, 0xc3, 0xe8, 0xe6, 0x58, 0xc0, 0xe0,
	0x04, 0x23, 0xcd, 0x45, 0xcd, 0xd3, 0x2f, 0xb0, 0xbf, 0x42, 0x03, 0xd9, 0x6a, 0x02, 0xad, 0x62,
	0xd5, 0x08, 0xc5, 0x38, 0x7a, 0x5e, 0x29, 0x3e, 0x1b, 0x2c, 0x5e, 0x86, 0xb6, 0xd6, 0x3a, 0x56,
	0x72, 0xc4, 0xe2, 0x65, 0x68, 0x93, 0x43, 0xf0, 0x71, 0x21, 0x3e, 0x2e, 0xe4, 0xde, 0xfa, 0x42,
	0x74, 0x9f, 0x0c, 0xdf, 0xa4, 0x3f, 0x5c, 0xd8, 0xb5, 0xfa, 0x1b, 0xae, 0x98, 0xa8, 0xfe, 0xf3,
	0x3a, 0xee, 0x43, 0x58, 0xf1, 0x85, 0xca, 0xb5, 0x6c, 0xce, 0x39, 0xd0, 0xfe, 0xd9, 0x75, 0xfd,
	0xd7, 0x6b, 0x38, 0x34, 0xa7, 0x6e, 0xae, 0x30, 0xc2, 0xa1, 0x27, 0xcb, 0xa1, 0xe7, 0x4c, 0x5e,
	0x65, 0xa1, 0xbd, 0x7c, 0x92, 0x40, 0x5c, 0x72, 0x95, 0x5f, 0x4b, 0xde, 0xe7, 0xa2, 0x2e, 0x29,
	0x24, 0xde, 0x2c, 0xca, 0xa0, 0xe4, 0xea, 0xa3, 0xe4, 0xfd, 0x69, 0x5d, 0xa6, 0x4f, 0x20, 0x78,
	0xc7, 0xe4, 0xfb, 0x56, 0x34, 0x64, 0x07, 0x02, 0x56, 0xf5, 0x9c, 0x15, 0xc3, 0xde, 0x2d, 0xed,
	0x5c, 0x32, 0x99, 0x37, 0xad, 0xda, 0x73, 0xd2, 0xef, 0x0e, 0xf8, 0xba, 0xf2, 0x1a, 0x2c, 0x02,
	0x3e, 0x5e, 0x99, 0x61, 0x85, 0xf6, 0x06, 0x54, 0x14, 0x02, 0x56, 0x98, 0x59, 0x0c, 0x2d, 0xeb,
	0x92, 0x03, 0x88, 0x84, 0xcc, 0x17, 0xa2, 0x11, 0xf2, 0x72, 0xc4, 0x16, 0x0a, 0x79, 0x82, 0xbe,
	0x2e, 0xae, 0x86, 0xce, 0x9e, 0x30, 0xda, 0x7a, 0x3b, 0x75, 0xdb, 0xf0, 0x01, 0x81, 0xb9, 0x99,
	0x71, 0xd2, 0x67, 0x10, 0xeb, 0xcf, 0x9b, 0x0f, 0x1d, 0xc7, 0xb1, 0x0f, 0x20, 0xd2, 0xaf, 0x6f,
	0x7e, 0xa8, 0x51, 0x16, 0xaa, 0x31, 0x98, 0x3e, 0x87, 0x70, 0x6e, 0xf9, 0x4c, 0xc7, 0xdb, 0x71,
	0x36, 0x61, 0x34, 0x27, 0xd3, 0xc1, 0x04, 0xbd, 0x7f, 0x74, 0xad, 0xd3, 0x95, 0x6b, 0xdd, 0xd8,
	0xf1, 0x35, 0xc4, 0xf3, 0xf1, 0xff, 0x18, 0x1b, 0x3e, 0x5d, 0xf9, 0xc8, 0xbb, 0xcb, 0x94, 0xb7,
	0x62, 0xa1, 0xec, 0x43, 0x93, 0x7a, 0xbe, 0x8d, 0x91, 0x97, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xcf, 0x70, 0xf6, 0x3d, 0xef, 0x05, 0x00, 0x00,
}
