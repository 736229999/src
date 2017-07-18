// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package dbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AccountChangeType int32

const (
	AccountChangeType_Create      AccountChangeType = 0
	AccountChangeType_Bind        AccountChangeType = 1
	AccountChangeType_Unbind      AccountChangeType = 2
	AccountChangeType_Password    AccountChangeType = 3
	AccountChangeType_ChangePhone AccountChangeType = 4
)

var AccountChangeType_name = map[int32]string{
	0: "Create",
	1: "Bind",
	2: "Unbind",
	3: "Password",
	4: "ChangePhone",
}
var AccountChangeType_value = map[string]int32{
	"Create":      0,
	"Bind":        1,
	"Unbind":      2,
	"Password":    3,
	"ChangePhone": 4,
}

func (x AccountChangeType) String() string {
	return proto.EnumName(AccountChangeType_name, int32(x))
}
func (AccountChangeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor22, []int{0} }

type UserType int32

const (
	UserType_Phone  UserType = 0
	UserType_QQ     UserType = 1
	UserType_Weixin UserType = 2
)

var UserType_name = map[int32]string{
	0: "Phone",
	1: "QQ",
	2: "Weixin",
}
var UserType_value = map[string]int32{
	"Phone":  0,
	"QQ":     1,
	"Weixin": 2,
}

func (x UserType) String() string {
	return proto.EnumName(UserType_name, int32(x))
}
func (UserType) EnumDescriptor() ([]byte, []int) { return fileDescriptor22, []int{1} }

// 手机用户
type PhoneUser struct {
	Phone     string `protobuf:"bytes,1,opt,name=phone" json:"phone"`
	Password  string `protobuf:"bytes,2,opt,name=password" json:"password"`
	AccountId int64  `protobuf:"varint,3,opt,name=account_id,json=accountId" json:"account_id"`
}

func (m *PhoneUser) Reset()                    { *m = PhoneUser{} }
func (m *PhoneUser) String() string            { return proto.CompactTextString(m) }
func (*PhoneUser) ProtoMessage()               {}
func (*PhoneUser) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{0} }

func (m *PhoneUser) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *PhoneUser) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *PhoneUser) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

// QQ用户
type QQUser struct {
	Openid    string `protobuf:"bytes,1,opt,name=openid" json:"openid"`
	AccountId int64  `protobuf:"varint,3,opt,name=account_id,json=accountId" json:"account_id"`
}

func (m *QQUser) Reset()                    { *m = QQUser{} }
func (m *QQUser) String() string            { return proto.CompactTextString(m) }
func (*QQUser) ProtoMessage()               {}
func (*QQUser) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{1} }

func (m *QQUser) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *QQUser) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

// 微信用户
type WeixinUser struct {
	Openid    string `protobuf:"bytes,1,opt,name=openid" json:"openid"`
	AccountId int64  `protobuf:"varint,3,opt,name=account_id,json=accountId" json:"account_id"`
}

func (m *WeixinUser) Reset()                    { *m = WeixinUser{} }
func (m *WeixinUser) String() string            { return proto.CompactTextString(m) }
func (*WeixinUser) ProtoMessage()               {}
func (*WeixinUser) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{2} }

func (m *WeixinUser) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *WeixinUser) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

// 彩金卡
type Giftcard struct {
	Id         int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	Type       int32  `protobuf:"varint,2,opt,name=type" json:"type"`
	UseBase    int32  `protobuf:"varint,3,opt,name=use_base,json=useBase" json:"use_base"`
	UseSub     int32  `protobuf:"varint,4,opt,name=use_sub,json=useSub" json:"use_sub"`
	MaxStack   int32  `protobuf:"varint,5,opt,name=max_stack,json=maxStack" json:"max_stack"`
	ValidStart int64  `protobuf:"varint,6,opt,name=valid_start,json=validStart" json:"valid_start"`
	ValidEnd   int64  `protobuf:"varint,7,opt,name=valid_end,json=validEnd" json:"valid_end"`
	GainTime   int64  `protobuf:"varint,8,opt,name=gain_time,json=gainTime" json:"gain_time"`
	GainSource string `protobuf:"bytes,9,opt,name=gain_source,json=gainSource" json:"gain_source"`
	AccountId  int64  `protobuf:"varint,10,opt,name=account_id,json=accountId" json:"account_id"`
}

func (m *Giftcard) Reset()                    { *m = Giftcard{} }
func (m *Giftcard) String() string            { return proto.CompactTextString(m) }
func (*Giftcard) ProtoMessage()               {}
func (*Giftcard) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{3} }

func (m *Giftcard) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Giftcard) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Giftcard) GetUseBase() int32 {
	if m != nil {
		return m.UseBase
	}
	return 0
}

func (m *Giftcard) GetUseSub() int32 {
	if m != nil {
		return m.UseSub
	}
	return 0
}

func (m *Giftcard) GetMaxStack() int32 {
	if m != nil {
		return m.MaxStack
	}
	return 0
}

func (m *Giftcard) GetValidStart() int64 {
	if m != nil {
		return m.ValidStart
	}
	return 0
}

func (m *Giftcard) GetValidEnd() int64 {
	if m != nil {
		return m.ValidEnd
	}
	return 0
}

func (m *Giftcard) GetGainTime() int64 {
	if m != nil {
		return m.GainTime
	}
	return 0
}

func (m *Giftcard) GetGainSource() string {
	if m != nil {
		return m.GainSource
	}
	return ""
}

func (m *Giftcard) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

// 用户信息
type UserInfo struct {
	AccountId int64  `protobuf:"varint,1,opt,name=account_id,json=accountId" json:"account_id"`
	Icon      string `protobuf:"bytes,2,opt,name=icon" json:"icon"`
	Nickname  string `protobuf:"bytes,3,opt,name=nickname" json:"nickname"`
	Realname  string `protobuf:"bytes,4,opt,name=realname" json:"realname"`
	Idno      string `protobuf:"bytes,5,opt,name=idno" json:"idno"`
	Sex       int32  `protobuf:"varint,6,opt,name=sex" json:"sex"`
	Exp       int32  `protobuf:"varint,7,opt,name=exp" json:"exp"`
	Level     int32  `protobuf:"varint,8,opt,name=level" json:"level"`
	// AccountBankcard bankcard = 9; // 银行卡
	PayPassword     bool        `protobuf:"varint,10,opt,name=pay_password,json=payPassword" json:"pay_password"`
	PayOpenPassword bool        `protobuf:"varint,11,opt,name=pay_open_password,json=payOpenPassword" json:"pay_open_password"`
	InvitationCode  string      `protobuf:"bytes,12,opt,name=invitation_code,json=invitationCode" json:"invitation_code"`
	Phone           *PhoneUser  `protobuf:"bytes,20,opt,name=phone" json:"phone"`
	Qq              *QQUser     `protobuf:"bytes,21,opt,name=qq" json:"qq"`
	Weixin          *WeixinUser `protobuf:"bytes,22,opt,name=weixin" json:"weixin"`
	DailyCheckTime  int64       `protobuf:"varint,23,opt,name=daily_check_time,json=dailyCheckTime" json:"daily_check_time"`
	ContCheckDays   int32       `protobuf:"varint,24,opt,name=cont_check_days,json=contCheckDays" json:"cont_check_days"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{4} }

func (m *UserInfo) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *UserInfo) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *UserInfo) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UserInfo) GetRealname() string {
	if m != nil {
		return m.Realname
	}
	return ""
}

func (m *UserInfo) GetIdno() string {
	if m != nil {
		return m.Idno
	}
	return ""
}

func (m *UserInfo) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *UserInfo) GetExp() int32 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func (m *UserInfo) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *UserInfo) GetPayPassword() bool {
	if m != nil {
		return m.PayPassword
	}
	return false
}

func (m *UserInfo) GetPayOpenPassword() bool {
	if m != nil {
		return m.PayOpenPassword
	}
	return false
}

func (m *UserInfo) GetInvitationCode() string {
	if m != nil {
		return m.InvitationCode
	}
	return ""
}

func (m *UserInfo) GetPhone() *PhoneUser {
	if m != nil {
		return m.Phone
	}
	return nil
}

func (m *UserInfo) GetQq() *QQUser {
	if m != nil {
		return m.Qq
	}
	return nil
}

func (m *UserInfo) GetWeixin() *WeixinUser {
	if m != nil {
		return m.Weixin
	}
	return nil
}

func (m *UserInfo) GetDailyCheckTime() int64 {
	if m != nil {
		return m.DailyCheckTime
	}
	return 0
}

func (m *UserInfo) GetContCheckDays() int32 {
	if m != nil {
		return m.ContCheckDays
	}
	return 0
}

// --------------------------------------------------------------------------------------------------------------
// 创建账户参数
type CreateAccountArg struct {
	// Types that are valid to be assigned to User:
	//	*CreateAccountArg_PhoneUser
	//	*CreateAccountArg_QqUser
	//	*CreateAccountArg_WxUser
	User     isCreateAccountArg_User `protobuf_oneof:"user"`
	UserType UserType                `protobuf:"varint,4,opt,name=user_type,json=userType,enum=dbproto.UserType" json:"user_type"`
	Ip       string                  `protobuf:"bytes,5,opt,name=ip" json:"ip"`
}

func (m *CreateAccountArg) Reset()                    { *m = CreateAccountArg{} }
func (m *CreateAccountArg) String() string            { return proto.CompactTextString(m) }
func (*CreateAccountArg) ProtoMessage()               {}
func (*CreateAccountArg) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{5} }

type isCreateAccountArg_User interface {
	isCreateAccountArg_User()
}

type CreateAccountArg_PhoneUser struct {
	PhoneUser *PhoneUser `protobuf:"bytes,1,opt,name=phone_user,json=phoneUser,oneof"`
}
type CreateAccountArg_QqUser struct {
	QqUser *QQUser `protobuf:"bytes,2,opt,name=qq_user,json=qqUser,oneof"`
}
type CreateAccountArg_WxUser struct {
	WxUser *WeixinUser `protobuf:"bytes,3,opt,name=wx_user,json=wxUser,oneof"`
}

func (*CreateAccountArg_PhoneUser) isCreateAccountArg_User() {}
func (*CreateAccountArg_QqUser) isCreateAccountArg_User()    {}
func (*CreateAccountArg_WxUser) isCreateAccountArg_User()    {}

func (m *CreateAccountArg) GetUser() isCreateAccountArg_User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateAccountArg) GetPhoneUser() *PhoneUser {
	if x, ok := m.GetUser().(*CreateAccountArg_PhoneUser); ok {
		return x.PhoneUser
	}
	return nil
}

func (m *CreateAccountArg) GetQqUser() *QQUser {
	if x, ok := m.GetUser().(*CreateAccountArg_QqUser); ok {
		return x.QqUser
	}
	return nil
}

func (m *CreateAccountArg) GetWxUser() *WeixinUser {
	if x, ok := m.GetUser().(*CreateAccountArg_WxUser); ok {
		return x.WxUser
	}
	return nil
}

func (m *CreateAccountArg) GetUserType() UserType {
	if m != nil {
		return m.UserType
	}
	return UserType_Phone
}

func (m *CreateAccountArg) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CreateAccountArg) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CreateAccountArg_OneofMarshaler, _CreateAccountArg_OneofUnmarshaler, _CreateAccountArg_OneofSizer, []interface{}{
		(*CreateAccountArg_PhoneUser)(nil),
		(*CreateAccountArg_QqUser)(nil),
		(*CreateAccountArg_WxUser)(nil),
	}
}

func _CreateAccountArg_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CreateAccountArg)
	// user
	switch x := m.User.(type) {
	case *CreateAccountArg_PhoneUser:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PhoneUser); err != nil {
			return err
		}
	case *CreateAccountArg_QqUser:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.QqUser); err != nil {
			return err
		}
	case *CreateAccountArg_WxUser:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.WxUser); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CreateAccountArg.User has unexpected type %T", x)
	}
	return nil
}

func _CreateAccountArg_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CreateAccountArg)
	switch tag {
	case 1: // user.phone_user
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PhoneUser)
		err := b.DecodeMessage(msg)
		m.User = &CreateAccountArg_PhoneUser{msg}
		return true, err
	case 2: // user.qq_user
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(QQUser)
		err := b.DecodeMessage(msg)
		m.User = &CreateAccountArg_QqUser{msg}
		return true, err
	case 3: // user.wx_user
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WeixinUser)
		err := b.DecodeMessage(msg)
		m.User = &CreateAccountArg_WxUser{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CreateAccountArg_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CreateAccountArg)
	// user
	switch x := m.User.(type) {
	case *CreateAccountArg_PhoneUser:
		s := proto.Size(x.PhoneUser)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CreateAccountArg_QqUser:
		s := proto.Size(x.QqUser)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CreateAccountArg_WxUser:
		s := proto.Size(x.WxUser)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// 修改账号绑定手机参数
type ChangePhoneArg struct {
	Phone    string `protobuf:"bytes,1,opt,name=phone" json:"phone"`
	NewPhone string `protobuf:"bytes,2,opt,name=new_phone,json=newPhone" json:"new_phone"`
	Ip       string `protobuf:"bytes,3,opt,name=ip" json:"ip"`
}

func (m *ChangePhoneArg) Reset()                    { *m = ChangePhoneArg{} }
func (m *ChangePhoneArg) String() string            { return proto.CompactTextString(m) }
func (*ChangePhoneArg) ProtoMessage()               {}
func (*ChangePhoneArg) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{6} }

func (m *ChangePhoneArg) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ChangePhoneArg) GetNewPhone() string {
	if m != nil {
		return m.NewPhone
	}
	return ""
}

func (m *ChangePhoneArg) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

// 绑定/解绑用户参数
type SetUserArg struct {
	// Types that are valid to be assigned to User:
	//	*SetUserArg_PhoneUser
	//	*SetUserArg_QqUser
	//	*SetUserArg_WxUser
	User     isSetUserArg_User `protobuf_oneof:"user"`
	UserType UserType          `protobuf:"varint,4,opt,name=user_type,json=userType,enum=dbproto.UserType" json:"user_type"`
	SetType  AccountChangeType `protobuf:"varint,5,opt,name=set_type,json=setType,enum=dbproto.AccountChangeType" json:"set_type"`
	Ip       string            `protobuf:"bytes,6,opt,name=ip" json:"ip"`
}

func (m *SetUserArg) Reset()                    { *m = SetUserArg{} }
func (m *SetUserArg) String() string            { return proto.CompactTextString(m) }
func (*SetUserArg) ProtoMessage()               {}
func (*SetUserArg) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{7} }

type isSetUserArg_User interface {
	isSetUserArg_User()
}

type SetUserArg_PhoneUser struct {
	PhoneUser *PhoneUser `protobuf:"bytes,1,opt,name=phone_user,json=phoneUser,oneof"`
}
type SetUserArg_QqUser struct {
	QqUser *QQUser `protobuf:"bytes,2,opt,name=qq_user,json=qqUser,oneof"`
}
type SetUserArg_WxUser struct {
	WxUser *WeixinUser `protobuf:"bytes,3,opt,name=wx_user,json=wxUser,oneof"`
}

func (*SetUserArg_PhoneUser) isSetUserArg_User() {}
func (*SetUserArg_QqUser) isSetUserArg_User()    {}
func (*SetUserArg_WxUser) isSetUserArg_User()    {}

func (m *SetUserArg) GetUser() isSetUserArg_User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *SetUserArg) GetPhoneUser() *PhoneUser {
	if x, ok := m.GetUser().(*SetUserArg_PhoneUser); ok {
		return x.PhoneUser
	}
	return nil
}

func (m *SetUserArg) GetQqUser() *QQUser {
	if x, ok := m.GetUser().(*SetUserArg_QqUser); ok {
		return x.QqUser
	}
	return nil
}

func (m *SetUserArg) GetWxUser() *WeixinUser {
	if x, ok := m.GetUser().(*SetUserArg_WxUser); ok {
		return x.WxUser
	}
	return nil
}

func (m *SetUserArg) GetUserType() UserType {
	if m != nil {
		return m.UserType
	}
	return UserType_Phone
}

func (m *SetUserArg) GetSetType() AccountChangeType {
	if m != nil {
		return m.SetType
	}
	return AccountChangeType_Create
}

func (m *SetUserArg) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SetUserArg) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SetUserArg_OneofMarshaler, _SetUserArg_OneofUnmarshaler, _SetUserArg_OneofSizer, []interface{}{
		(*SetUserArg_PhoneUser)(nil),
		(*SetUserArg_QqUser)(nil),
		(*SetUserArg_WxUser)(nil),
	}
}

func _SetUserArg_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SetUserArg)
	// user
	switch x := m.User.(type) {
	case *SetUserArg_PhoneUser:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PhoneUser); err != nil {
			return err
		}
	case *SetUserArg_QqUser:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.QqUser); err != nil {
			return err
		}
	case *SetUserArg_WxUser:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.WxUser); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SetUserArg.User has unexpected type %T", x)
	}
	return nil
}

func _SetUserArg_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SetUserArg)
	switch tag {
	case 1: // user.phone_user
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PhoneUser)
		err := b.DecodeMessage(msg)
		m.User = &SetUserArg_PhoneUser{msg}
		return true, err
	case 2: // user.qq_user
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(QQUser)
		err := b.DecodeMessage(msg)
		m.User = &SetUserArg_QqUser{msg}
		return true, err
	case 3: // user.wx_user
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WeixinUser)
		err := b.DecodeMessage(msg)
		m.User = &SetUserArg_WxUser{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SetUserArg_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SetUserArg)
	// user
	switch x := m.User.(type) {
	case *SetUserArg_PhoneUser:
		s := proto.Size(x.PhoneUser)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SetUserArg_QqUser:
		s := proto.Size(x.QqUser)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SetUserArg_WxUser:
		s := proto.Size(x.WxUser)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PhonePassword struct {
	Phone    string `protobuf:"bytes,1,opt,name=phone" json:"phone"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password"`
	Ip       string `protobuf:"bytes,3,opt,name=ip" json:"ip"`
}

func (m *PhonePassword) Reset()                    { *m = PhonePassword{} }
func (m *PhonePassword) String() string            { return proto.CompactTextString(m) }
func (*PhonePassword) ProtoMessage()               {}
func (*PhonePassword) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{8} }

func (m *PhonePassword) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *PhonePassword) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *PhonePassword) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type UserInfoArg struct {
	AccountId int64  `protobuf:"varint,1,opt,name=account_id,json=accountId" json:"account_id"`
	Icon      string `protobuf:"bytes,2,opt,name=icon" json:"icon"`
	Nickname  string `protobuf:"bytes,3,opt,name=nickname" json:"nickname"`
	Sex       int32  `protobuf:"varint,4,opt,name=sex" json:"sex"`
}

func (m *UserInfoArg) Reset()                    { *m = UserInfoArg{} }
func (m *UserInfoArg) String() string            { return proto.CompactTextString(m) }
func (*UserInfoArg) ProtoMessage()               {}
func (*UserInfoArg) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{9} }

func (m *UserInfoArg) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *UserInfoArg) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *UserInfoArg) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UserInfoArg) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

type AccountIdcard struct {
	AccountId int64  `protobuf:"varint,1,opt,name=account_id,json=accountId" json:"account_id"`
	Idno      string `protobuf:"bytes,2,opt,name=idno" json:"idno"`
	Realname  string `protobuf:"bytes,3,opt,name=realname" json:"realname"`
	AddTime   int64  `protobuf:"varint,4,opt,name=add_time,json=addTime" json:"add_time"`
}

func (m *AccountIdcard) Reset()                    { *m = AccountIdcard{} }
func (m *AccountIdcard) String() string            { return proto.CompactTextString(m) }
func (*AccountIdcard) ProtoMessage()               {}
func (*AccountIdcard) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{10} }

func (m *AccountIdcard) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *AccountIdcard) GetIdno() string {
	if m != nil {
		return m.Idno
	}
	return ""
}

func (m *AccountIdcard) GetRealname() string {
	if m != nil {
		return m.Realname
	}
	return ""
}

func (m *AccountIdcard) GetAddTime() int64 {
	if m != nil {
		return m.AddTime
	}
	return 0
}

type PayPasswordArg struct {
	AccountId int64  `protobuf:"varint,1,opt,name=account_id,json=accountId" json:"account_id"`
	Password  string `protobuf:"bytes,2,opt,name=password" json:"password"`
	Open      bool   `protobuf:"varint,3,opt,name=open" json:"open"`
}

func (m *PayPasswordArg) Reset()                    { *m = PayPasswordArg{} }
func (m *PayPasswordArg) String() string            { return proto.CompactTextString(m) }
func (*PayPasswordArg) ProtoMessage()               {}
func (*PayPasswordArg) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{11} }

func (m *PayPasswordArg) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *PayPasswordArg) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *PayPasswordArg) GetOpen() bool {
	if m != nil {
		return m.Open
	}
	return false
}

type PaySettings struct {
	Password string `protobuf:"bytes,1,opt,name=password" json:"password"`
	Open     bool   `protobuf:"varint,2,opt,name=open" json:"open"`
}

func (m *PaySettings) Reset()                    { *m = PaySettings{} }
func (m *PaySettings) String() string            { return proto.CompactTextString(m) }
func (*PaySettings) ProtoMessage()               {}
func (*PaySettings) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{12} }

func (m *PaySettings) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *PaySettings) GetOpen() bool {
	if m != nil {
		return m.Open
	}
	return false
}

type DailyCheckArg struct {
	AccountId     int64 `protobuf:"varint,1,opt,name=account_id,json=accountId" json:"account_id"`
	Exp           int32 `protobuf:"varint,2,opt,name=exp" json:"exp"`
	Level         int32 `protobuf:"varint,3,opt,name=level" json:"level"`
	ContCheckDays int32 `protobuf:"varint,4,opt,name=cont_check_days,json=contCheckDays" json:"cont_check_days"`
}

func (m *DailyCheckArg) Reset()                    { *m = DailyCheckArg{} }
func (m *DailyCheckArg) String() string            { return proto.CompactTextString(m) }
func (*DailyCheckArg) ProtoMessage()               {}
func (*DailyCheckArg) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{13} }

func (m *DailyCheckArg) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *DailyCheckArg) GetExp() int32 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func (m *DailyCheckArg) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *DailyCheckArg) GetContCheckDays() int32 {
	if m != nil {
		return m.ContCheckDays
	}
	return 0
}

func init() {
	proto.RegisterType((*PhoneUser)(nil), "dbproto.PhoneUser")
	proto.RegisterType((*QQUser)(nil), "dbproto.QQUser")
	proto.RegisterType((*WeixinUser)(nil), "dbproto.WeixinUser")
	proto.RegisterType((*Giftcard)(nil), "dbproto.Giftcard")
	proto.RegisterType((*UserInfo)(nil), "dbproto.UserInfo")
	proto.RegisterType((*CreateAccountArg)(nil), "dbproto.CreateAccountArg")
	proto.RegisterType((*ChangePhoneArg)(nil), "dbproto.ChangePhoneArg")
	proto.RegisterType((*SetUserArg)(nil), "dbproto.SetUserArg")
	proto.RegisterType((*PhonePassword)(nil), "dbproto.PhonePassword")
	proto.RegisterType((*UserInfoArg)(nil), "dbproto.UserInfoArg")
	proto.RegisterType((*AccountIdcard)(nil), "dbproto.AccountIdcard")
	proto.RegisterType((*PayPasswordArg)(nil), "dbproto.PayPasswordArg")
	proto.RegisterType((*PaySettings)(nil), "dbproto.PaySettings")
	proto.RegisterType((*DailyCheckArg)(nil), "dbproto.DailyCheckArg")
	proto.RegisterEnum("dbproto.AccountChangeType", AccountChangeType_name, AccountChangeType_value)
	proto.RegisterEnum("dbproto.UserType", UserType_name, UserType_value)
}

func init() { proto.RegisterFile("user.proto", fileDescriptor22) }

var fileDescriptor22 = []byte{
	// 931 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x55, 0xcf, 0x6e, 0xdb, 0xc6,
	0x13, 0xb6, 0x28, 0x8a, 0x22, 0x47, 0xb6, 0x24, 0xef, 0x2f, 0xbf, 0x84, 0x49, 0x50, 0xc4, 0xd5,
	0xa1, 0x55, 0x5d, 0xc0, 0x87, 0x04, 0x3d, 0x16, 0x85, 0xad, 0x14, 0x4d, 0x4e, 0xb5, 0x28, 0x07,
	0xbd, 0x14, 0x20, 0x56, 0xdc, 0x8d, 0xbd, 0xb0, 0xb4, 0xa4, 0xb8, 0x94, 0x25, 0x5e, 0x7a, 0xeb,
	0x4b, 0x14, 0x7d, 0xc3, 0xbe, 0x44, 0x31, 0xb3, 0x14, 0xad, 0x2a, 0x4a, 0xeb, 0xa2, 0x3d, 0xf5,
	0x36, 0x7f, 0xbe, 0x19, 0x72, 0x3e, 0xce, 0x7c, 0x04, 0x58, 0x1a, 0x99, 0x9f, 0x65, 0x79, 0x5a,
	0xa4, 0xac, 0x2d, 0xa6, 0x64, 0x0c, 0x7e, 0x84, 0xe0, 0xf2, 0x26, 0xd5, 0xf2, 0x9d, 0x91, 0x39,
	0x7b, 0x04, 0xad, 0x0c, 0x9d, 0xb0, 0x71, 0xd2, 0x18, 0x06, 0x91, 0x75, 0xd8, 0x33, 0xf0, 0x33,
	0x6e, 0xcc, 0x2a, 0xcd, 0x45, 0xe8, 0x50, 0xa2, 0xf6, 0xd9, 0x27, 0x00, 0x3c, 0x49, 0xd2, 0xa5,
	0x2e, 0x62, 0x25, 0xc2, 0xe6, 0x49, 0x63, 0xd8, 0x8c, 0x82, 0x2a, 0xf2, 0x56, 0x0c, 0xbe, 0x01,
	0x6f, 0x3c, 0xa6, 0xd6, 0x8f, 0xc1, 0x4b, 0x33, 0xa9, 0x95, 0xa8, 0x7a, 0x57, 0xde, 0x5f, 0x35,
	0x18, 0x01, 0xfc, 0x20, 0xd5, 0x5a, 0xe9, 0x7f, 0xd2, 0xe4, 0x17, 0x07, 0xfc, 0xef, 0xd4, 0xfb,
	0x22, 0xe1, 0xb9, 0x60, 0x5d, 0x70, 0xaa, 0xfa, 0x66, 0xe4, 0x28, 0xc1, 0x18, 0xb8, 0x45, 0x99,
	0x49, 0x9a, 0xac, 0x15, 0x91, 0xcd, 0x9e, 0x82, 0xbf, 0x34, 0x32, 0x9e, 0x72, 0x23, 0xa9, 0x5b,
	0x2b, 0x6a, 0x2f, 0x8d, 0xbc, 0xe0, 0x46, 0xb2, 0x27, 0x80, 0x66, 0x6c, 0x96, 0xd3, 0xd0, 0xa5,
	0x8c, 0xb7, 0x34, 0x72, 0xb2, 0x9c, 0xb2, 0xe7, 0x10, 0xcc, 0xf9, 0x3a, 0x36, 0x05, 0x4f, 0x6e,
	0xc3, 0x16, 0xa5, 0xfc, 0x39, 0x5f, 0x4f, 0xd0, 0x67, 0x2f, 0xa0, 0x73, 0xc7, 0x67, 0x4a, 0x60,
	0x3a, 0x2f, 0x42, 0x8f, 0x9e, 0x0e, 0x14, 0x9a, 0x60, 0x04, 0xab, 0x2d, 0x40, 0x6a, 0x11, 0xb6,
	0x29, 0xed, 0x53, 0xe0, 0x5b, 0x2d, 0x30, 0x79, 0xcd, 0x95, 0x8e, 0x0b, 0x35, 0x97, 0xa1, 0x6f,
	0x93, 0x18, 0xb8, 0x52, 0x73, 0x89, 0xad, 0x29, 0x69, 0xd2, 0x65, 0x9e, 0xc8, 0x30, 0x20, 0x62,
	0x00, 0x43, 0x13, 0x8a, 0xec, 0x90, 0x03, 0xbb, 0xe4, 0xfc, 0xec, 0x82, 0x8f, 0xe4, 0xbe, 0xd5,
	0xef, 0xd3, 0x1d, 0x6c, 0x63, 0x07, 0x8b, 0x5c, 0xa9, 0x24, 0xd5, 0xd5, 0x16, 0x90, 0x8d, 0xdb,
	0xa1, 0x55, 0x72, 0xab, 0xf9, 0xdc, 0x72, 0x15, 0x44, 0xb5, 0x8f, 0xb9, 0x5c, 0xf2, 0x19, 0xe5,
	0x5c, 0x9b, 0xdb, 0xf8, 0xd4, 0x4b, 0xe8, 0x94, 0xa8, 0xc2, 0x5e, 0x42, 0xa7, 0xac, 0x0f, 0x4d,
	0x23, 0xd7, 0x44, 0x4f, 0x2b, 0x42, 0x13, 0x23, 0x72, 0x9d, 0x11, 0x23, 0xad, 0x08, 0x4d, 0xdc,
	0xd1, 0x99, 0xbc, 0x93, 0x33, 0x22, 0xa2, 0x15, 0x59, 0x87, 0x7d, 0x0a, 0x87, 0x19, 0x2f, 0xe3,
	0x7a, 0x4f, 0x71, 0x4c, 0x3f, 0xea, 0x64, 0xbc, 0xbc, 0xdc, 0xac, 0xea, 0x29, 0x1c, 0x23, 0x04,
	0x57, 0xe6, 0x1e, 0xd7, 0x21, 0x5c, 0x2f, 0xe3, 0xe5, 0xf7, 0x99, 0xd4, 0x35, 0xf6, 0x73, 0xe8,
	0x29, 0x7d, 0xa7, 0x0a, 0x5e, 0xa8, 0x54, 0xc7, 0x49, 0x2a, 0x64, 0x78, 0x48, 0xef, 0xd9, 0xbd,
	0x0f, 0x8f, 0x52, 0x21, 0xd9, 0x70, 0x73, 0x31, 0x8f, 0x4e, 0x1a, 0xc3, 0xce, 0x4b, 0x76, 0x56,
	0xdd, 0xd5, 0x59, 0x7d, 0x54, 0x9b, 0x2b, 0x7a, 0x01, 0xce, 0x62, 0x11, 0xfe, 0x9f, 0x60, 0xbd,
	0x1a, 0x66, 0xaf, 0x23, 0x72, 0x16, 0x0b, 0xf6, 0x25, 0x78, 0x2b, 0x5a, 0xf5, 0xf0, 0x31, 0x81,
	0xfe, 0x57, 0x83, 0xee, 0x2f, 0x20, 0xaa, 0x20, 0x6c, 0x08, 0x7d, 0xc1, 0xd5, 0xac, 0x8c, 0x93,
	0x1b, 0x99, 0xdc, 0xda, 0xcd, 0x78, 0x42, 0x9f, 0xab, 0x4b, 0xf1, 0x11, 0x86, 0x69, 0x3f, 0x3e,
	0x83, 0x5e, 0x92, 0xea, 0xa2, 0x02, 0x0a, 0x5e, 0x9a, 0x30, 0x24, 0xe6, 0x8e, 0x30, 0x4c, 0xb8,
	0xd7, 0xbc, 0x34, 0x83, 0xdf, 0x1a, 0xd0, 0x1f, 0xe5, 0x92, 0x17, 0xf2, 0xdc, 0x7e, 0xef, 0xf3,
	0xfc, 0x9a, 0xbd, 0x02, 0xa0, 0xb7, 0x8f, 0x51, 0x3a, 0x68, 0x1f, 0xf6, 0xce, 0xf8, 0xe6, 0x20,
	0x0a, 0xb2, 0x5a, 0x45, 0x4e, 0xa1, 0xbd, 0x58, 0xd8, 0x0a, 0x67, 0xef, 0xb8, 0x6f, 0x0e, 0x22,
	0x6f, 0xb1, 0x20, 0xec, 0x19, 0xb4, 0x57, 0x6b, 0x8b, 0x6d, 0x7e, 0x74, 0x6a, 0xc4, 0xaf, 0xd6,
	0x15, 0x3e, 0x40, 0x70, 0x4c, 0x27, 0x8b, 0x2b, 0xd5, 0x7d, 0x79, 0x5c, 0x57, 0x20, 0xe2, 0xaa,
	0xcc, 0x64, 0x84, 0xd7, 0x4b, 0x16, 0x5d, 0x7b, 0x56, 0xed, 0x98, 0xa3, 0xb2, 0x0b, 0x0f, 0x5c,
	0xcc, 0x0d, 0x26, 0xd0, 0x1d, 0xdd, 0x70, 0x7d, 0x2d, 0x69, 0x06, 0x1c, 0x75, 0xbf, 0xf6, 0x3d,
	0x87, 0x40, 0xcb, 0x55, 0x6c, 0x33, 0x95, 0xf8, 0x69, 0xb9, 0xa2, 0xaa, 0xaa, 0x79, 0x73, 0xd3,
	0x7c, 0xf0, 0xab, 0x03, 0x30, 0x91, 0x05, 0xbe, 0xc6, 0x7f, 0x82, 0xbc, 0xaf, 0xc0, 0x37, 0xb2,
	0xb0, 0xf0, 0x16, 0xc1, 0x9f, 0xd5, 0xf0, 0x6a, 0x49, 0x2c, 0x89, 0x54, 0xd7, 0x36, 0xb2, 0xd8,
	0xe2, 0xdc, 0xfb, 0x80, 0xf3, 0x31, 0x1c, 0xd1, 0xd0, 0xf5, 0x95, 0xfd, 0xfd, 0xdf, 0xcd, 0x2e,
	0xe3, 0x1a, 0x3a, 0x1b, 0xed, 0x42, 0xc6, 0xff, 0x65, 0xf9, 0xaa, 0xe4, 0xc8, 0xad, 0xe5, 0x68,
	0x50, 0xc2, 0xd1, 0xf9, 0xa6, 0x1d, 0xfd, 0x4d, 0x1e, 0xf0, 0x44, 0x14, 0x39, 0x67, 0x4b, 0xe4,
	0xb6, 0x45, 0xb1, 0xb9, 0x23, 0x8a, 0x4f, 0xc1, 0xe7, 0x42, 0xd8, 0x73, 0x76, 0xa9, 0x59, 0x9b,
	0x0b, 0x81, 0x77, 0x3c, 0x88, 0xa1, 0x7b, 0x79, 0xaf, 0x66, 0x0f, 0x98, 0xf6, 0xcf, 0x78, 0x64,
	0xe0, 0xa2, 0x0e, 0xd2, 0xf3, 0xfd, 0x88, 0xec, 0xc1, 0xd7, 0xd0, 0xb9, 0xe4, 0xe5, 0x44, 0x16,
	0x85, 0xd2, 0xd7, 0xe6, 0x0f, 0xe5, 0x8d, 0x8f, 0x94, 0x3b, 0x5b, 0xe5, 0x3f, 0xc1, 0xd1, 0xeb,
	0x5a, 0x79, 0x1e, 0xf0, 0x7a, 0x95, 0xb2, 0x3b, 0x7b, 0x94, 0xbd, 0xb9, 0xad, 0xec, 0x7b, 0xf4,
	0xcb, 0xdd, 0xa3, 0x5f, 0xa7, 0x57, 0x70, 0xfc, 0xc1, 0x4e, 0x32, 0x00, 0xcf, 0x6a, 0x5a, 0xff,
	0x80, 0xf9, 0xe0, 0x5e, 0x28, 0x2d, 0xfa, 0x0d, 0x8c, 0xbe, 0xd3, 0x53, 0xb4, 0x1d, 0x76, 0x08,
	0xfe, 0x86, 0xd3, 0x7e, 0x93, 0xf5, 0xa0, 0xb3, 0x25, 0x0b, 0x7d, 0xf7, 0xf4, 0x0b, 0xfb, 0x73,
	0xa4, 0x66, 0x01, 0xb4, 0x6c, 0xf8, 0x80, 0x79, 0xe0, 0x8c, 0xc7, 0xb6, 0x93, 0xbd, 0xb4, 0xbe,
	0x33, 0xf5, 0xe8, 0x30, 0x5e, 0xfd, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x79, 0x89, 0x19, 0x47, 0x67,
	0x09, 0x00, 0x00,
}
