// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package apiproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

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
func (UserType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

type Sex int32

const (
	Sex_Unknown Sex = 0
	Sex_Male    Sex = 1
	Sex_Female  Sex = 2
)

var Sex_name = map[int32]string{
	0: "Unknown",
	1: "Male",
	2: "Female",
}
var Sex_value = map[string]int32{
	"Unknown": 0,
	"Male":    1,
	"Female":  2,
}

func (x Sex) String() string {
	return proto.EnumName(Sex_name, int32(x))
}
func (Sex) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

type UserInfo struct {
	AccountId        int64  `protobuf:"varint,1,opt,name=account_id,json=accountId" json:"account_id"`
	Nickname         string `protobuf:"bytes,2,opt,name=nickname" json:"nickname"`
	Icon             string `protobuf:"bytes,3,opt,name=icon" json:"icon"`
	Idno             string `protobuf:"bytes,4,opt,name=idno" json:"idno"`
	Sex              Sex    `protobuf:"varint,5,opt,name=sex,enum=apiproto.Sex" json:"sex"`
	Exp              int32  `protobuf:"varint,6,opt,name=exp" json:"exp"`
	Level            int32  `protobuf:"varint,7,opt,name=level" json:"level"`
	Realname         string `protobuf:"bytes,8,opt,name=realname" json:"realname"`
	Bindphone        string `protobuf:"bytes,10,opt,name=bindphone" json:"bindphone"`
	Bindqq           string `protobuf:"bytes,11,opt,name=bindqq" json:"bindqq"`
	Bindwx           string `protobuf:"bytes,12,opt,name=bindwx" json:"bindwx"`
	InvitationCode   string `protobuf:"bytes,13,opt,name=invitation_code,json=invitationCode" json:"invitation_code"`
	PayPassword      bool   `protobuf:"varint,14,opt,name=pay_password,json=payPassword" json:"pay_password"`
	PayOpenPassword  bool   `protobuf:"varint,15,opt,name=pay_open_password,json=payOpenPassword" json:"pay_open_password"`
	LoginPassword    bool   `protobuf:"varint,16,opt,name=login_password,json=loginPassword" json:"login_password"`
	UserInviteStatus bool   `protobuf:"varint,17,opt,name=UserInviteStatus" json:"UserInviteStatus"`
	IsDailyCheck     bool   `protobuf:"varint,18,opt,name=is_daily_check,json=isDailyCheck" json:"is_daily_check"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *UserInfo) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *UserInfo) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UserInfo) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *UserInfo) GetIdno() string {
	if m != nil {
		return m.Idno
	}
	return ""
}

func (m *UserInfo) GetSex() Sex {
	if m != nil {
		return m.Sex
	}
	return Sex_Unknown
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

func (m *UserInfo) GetRealname() string {
	if m != nil {
		return m.Realname
	}
	return ""
}

func (m *UserInfo) GetBindphone() string {
	if m != nil {
		return m.Bindphone
	}
	return ""
}

func (m *UserInfo) GetBindqq() string {
	if m != nil {
		return m.Bindqq
	}
	return ""
}

func (m *UserInfo) GetBindwx() string {
	if m != nil {
		return m.Bindwx
	}
	return ""
}

func (m *UserInfo) GetInvitationCode() string {
	if m != nil {
		return m.InvitationCode
	}
	return ""
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

func (m *UserInfo) GetLoginPassword() bool {
	if m != nil {
		return m.LoginPassword
	}
	return false
}

func (m *UserInfo) GetUserInviteStatus() bool {
	if m != nil {
		return m.UserInviteStatus
	}
	return false
}

func (m *UserInfo) GetIsDailyCheck() bool {
	if m != nil {
		return m.IsDailyCheck
	}
	return false
}

type RegistRequest struct {
	Phone    string `protobuf:"bytes,1,opt,name=phone" json:"phone"`
	Code     string `protobuf:"bytes,2,opt,name=code" json:"code"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password"`
}

func (m *RegistRequest) Reset()                    { *m = RegistRequest{} }
func (m *RegistRequest) String() string            { return proto.CompactTextString(m) }
func (*RegistRequest) ProtoMessage()               {}
func (*RegistRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *RegistRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RegistRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *RegistRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginRequest struct {
	Type        UserType `protobuf:"varint,1,opt,name=type,enum=apiproto.UserType" json:"type"`
	Phone       string   `protobuf:"bytes,2,opt,name=phone" json:"phone"`
	Password    string   `protobuf:"bytes,3,opt,name=password" json:"password"`
	Openid      string   `protobuf:"bytes,4,opt,name=openid" json:"openid"`
	AccessToken string   `protobuf:"bytes,5,opt,name=access_token,json=accessToken" json:"access_token"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *LoginRequest) GetType() UserType {
	if m != nil {
		return m.Type
	}
	return UserType_Phone
}

func (m *LoginRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *LoginRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type LoginReply struct {
	User     *UserInfo `protobuf:"bytes,1,opt,name=user" json:"user"`
	Fund     *FundInfo `protobuf:"bytes,2,opt,name=fund" json:"fund"`
	Bankcard *Bankcard `protobuf:"bytes,3,opt,name=bankcard" json:"bankcard"`
}

func (m *LoginReply) Reset()                    { *m = LoginReply{} }
func (m *LoginReply) String() string            { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()               {}
func (*LoginReply) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

func (m *LoginReply) GetUser() *UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *LoginReply) GetFund() *FundInfo {
	if m != nil {
		return m.Fund
	}
	return nil
}

func (m *LoginReply) GetBankcard() *Bankcard {
	if m != nil {
		return m.Bankcard
	}
	return nil
}

type ForgotPwdRequest struct {
	Phone    string `protobuf:"bytes,1,opt,name=phone" json:"phone"`
	Code     string `protobuf:"bytes,2,opt,name=code" json:"code"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password"`
}

func (m *ForgotPwdRequest) Reset()                    { *m = ForgotPwdRequest{} }
func (m *ForgotPwdRequest) String() string            { return proto.CompactTextString(m) }
func (*ForgotPwdRequest) ProtoMessage()               {}
func (*ForgotPwdRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

func (m *ForgotPwdRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ForgotPwdRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ForgotPwdRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ForgotPayPwdRequest struct {
	Phone         string `protobuf:"bytes,1,opt,name=phone" json:"phone"`
	LoginPassword string `protobuf:"bytes,2,opt,name=login_password,json=loginPassword" json:"login_password"`
	Code          string `protobuf:"bytes,3,opt,name=code" json:"code"`
}

func (m *ForgotPayPwdRequest) Reset()                    { *m = ForgotPayPwdRequest{} }
func (m *ForgotPayPwdRequest) String() string            { return proto.CompactTextString(m) }
func (*ForgotPayPwdRequest) ProtoMessage()               {}
func (*ForgotPayPwdRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{5} }

func (m *ForgotPayPwdRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ForgotPayPwdRequest) GetLoginPassword() string {
	if m != nil {
		return m.LoginPassword
	}
	return ""
}

func (m *ForgotPayPwdRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type VerifyPayPwdRequest struct {
	Password string `protobuf:"bytes,1,opt,name=password" json:"password"`
	Api      string `protobuf:"bytes,2,opt,name=api" json:"api"`
}

func (m *VerifyPayPwdRequest) Reset()                    { *m = VerifyPayPwdRequest{} }
func (m *VerifyPayPwdRequest) String() string            { return proto.CompactTextString(m) }
func (*VerifyPayPwdRequest) ProtoMessage()               {}
func (*VerifyPayPwdRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{6} }

func (m *VerifyPayPwdRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *VerifyPayPwdRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

type SetPwdRequest struct {
	Password string `protobuf:"bytes,1,opt,name=password" json:"password"`
	Token    string `protobuf:"bytes,2,opt,name=token" json:"token"`
}

func (m *SetPwdRequest) Reset()                    { *m = SetPwdRequest{} }
func (m *SetPwdRequest) String() string            { return proto.CompactTextString(m) }
func (*SetPwdRequest) ProtoMessage()               {}
func (*SetPwdRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{7} }

func (m *SetPwdRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SetPwdRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ResetPwdRequest struct {
	Password    string `protobuf:"bytes,1,opt,name=password" json:"password"`
	NewPassword string `protobuf:"bytes,2,opt,name=new_password,json=newPassword" json:"new_password"`
}

func (m *ResetPwdRequest) Reset()                    { *m = ResetPwdRequest{} }
func (m *ResetPwdRequest) String() string            { return proto.CompactTextString(m) }
func (*ResetPwdRequest) ProtoMessage()               {}
func (*ResetPwdRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{8} }

func (m *ResetPwdRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ResetPwdRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type BindPhoneRequest struct {
	Phone string `protobuf:"bytes,1,opt,name=phone" json:"phone"`
	Code  string `protobuf:"bytes,2,opt,name=code" json:"code"`
}

func (m *BindPhoneRequest) Reset()                    { *m = BindPhoneRequest{} }
func (m *BindPhoneRequest) String() string            { return proto.CompactTextString(m) }
func (*BindPhoneRequest) ProtoMessage()               {}
func (*BindPhoneRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{9} }

func (m *BindPhoneRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *BindPhoneRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type ChangePhoneRequest struct {
	Phone    string `protobuf:"bytes,1,opt,name=phone" json:"phone"`
	NewPhone string `protobuf:"bytes,2,opt,name=new_phone,json=newPhone" json:"new_phone"`
	Code     string `protobuf:"bytes,3,opt,name=code" json:"code"`
}

func (m *ChangePhoneRequest) Reset()                    { *m = ChangePhoneRequest{} }
func (m *ChangePhoneRequest) String() string            { return proto.CompactTextString(m) }
func (*ChangePhoneRequest) ProtoMessage()               {}
func (*ChangePhoneRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{10} }

func (m *ChangePhoneRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ChangePhoneRequest) GetNewPhone() string {
	if m != nil {
		return m.NewPhone
	}
	return ""
}

func (m *ChangePhoneRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type BindWeixinRequest struct {
	Openid string `protobuf:"bytes,1,opt,name=openid" json:"openid"`
}

func (m *BindWeixinRequest) Reset()                    { *m = BindWeixinRequest{} }
func (m *BindWeixinRequest) String() string            { return proto.CompactTextString(m) }
func (*BindWeixinRequest) ProtoMessage()               {}
func (*BindWeixinRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{11} }

func (m *BindWeixinRequest) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

type BindQQRequest struct {
	Openid string `protobuf:"bytes,1,opt,name=openid" json:"openid"`
}

func (m *BindQQRequest) Reset()                    { *m = BindQQRequest{} }
func (m *BindQQRequest) String() string            { return proto.CompactTextString(m) }
func (*BindQQRequest) ProtoMessage()               {}
func (*BindQQRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{12} }

func (m *BindQQRequest) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

type SetIconRequest struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value"`
}

func (m *SetIconRequest) Reset()                    { *m = SetIconRequest{} }
func (m *SetIconRequest) String() string            { return proto.CompactTextString(m) }
func (*SetIconRequest) ProtoMessage()               {}
func (*SetIconRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{13} }

func (m *SetIconRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetNicknameRequest struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value"`
}

func (m *SetNicknameRequest) Reset()                    { *m = SetNicknameRequest{} }
func (m *SetNicknameRequest) String() string            { return proto.CompactTextString(m) }
func (*SetNicknameRequest) ProtoMessage()               {}
func (*SetNicknameRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{14} }

func (m *SetNicknameRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetSexRequest struct {
	Value Sex `protobuf:"varint,1,opt,name=value,enum=apiproto.Sex" json:"value"`
}

func (m *SetSexRequest) Reset()                    { *m = SetSexRequest{} }
func (m *SetSexRequest) String() string            { return proto.CompactTextString(m) }
func (*SetSexRequest) ProtoMessage()               {}
func (*SetSexRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{15} }

func (m *SetSexRequest) GetValue() Sex {
	if m != nil {
		return m.Value
	}
	return Sex_Unknown
}

type SetBoolRequest struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value"`
}

func (m *SetBoolRequest) Reset()                    { *m = SetBoolRequest{} }
func (m *SetBoolRequest) String() string            { return proto.CompactTextString(m) }
func (*SetBoolRequest) ProtoMessage()               {}
func (*SetBoolRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{16} }

func (m *SetBoolRequest) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type SetIntRequest struct {
	Value int64 `protobuf:"varint,1,opt,name=value" json:"value"`
}

func (m *SetIntRequest) Reset()                    { *m = SetIntRequest{} }
func (m *SetIntRequest) String() string            { return proto.CompactTextString(m) }
func (*SetIntRequest) ProtoMessage()               {}
func (*SetIntRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{17} }

func (m *SetIntRequest) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// 实名认证
type AuthRealnameRequest struct {
	Realname string `protobuf:"bytes,1,opt,name=realname" json:"realname"`
	Idno     string `protobuf:"bytes,2,opt,name=idno" json:"idno"`
	Phone    string `protobuf:"bytes,3,opt,name=phone" json:"phone"`
	SmsCode  string `protobuf:"bytes,4,opt,name=sms_code,json=smsCode" json:"sms_code"`
}

func (m *AuthRealnameRequest) Reset()                    { *m = AuthRealnameRequest{} }
func (m *AuthRealnameRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthRealnameRequest) ProtoMessage()               {}
func (*AuthRealnameRequest) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{18} }

func (m *AuthRealnameRequest) GetRealname() string {
	if m != nil {
		return m.Realname
	}
	return ""
}

func (m *AuthRealnameRequest) GetIdno() string {
	if m != nil {
		return m.Idno
	}
	return ""
}

func (m *AuthRealnameRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *AuthRealnameRequest) GetSmsCode() string {
	if m != nil {
		return m.SmsCode
	}
	return ""
}

// 每日签到
type DailyCheckReply struct {
	ContCheckDays int32 `protobuf:"varint,1,opt,name=cont_check_days,json=contCheckDays" json:"cont_check_days"`
	Exp           int32 `protobuf:"varint,2,opt,name=exp" json:"exp"`
}

func (m *DailyCheckReply) Reset()                    { *m = DailyCheckReply{} }
func (m *DailyCheckReply) String() string            { return proto.CompactTextString(m) }
func (*DailyCheckReply) ProtoMessage()               {}
func (*DailyCheckReply) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{19} }

func (m *DailyCheckReply) GetContCheckDays() int32 {
	if m != nil {
		return m.ContCheckDays
	}
	return 0
}

func (m *DailyCheckReply) GetExp() int32 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "apiproto.UserInfo")
	proto.RegisterType((*RegistRequest)(nil), "apiproto.RegistRequest")
	proto.RegisterType((*LoginRequest)(nil), "apiproto.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "apiproto.LoginReply")
	proto.RegisterType((*ForgotPwdRequest)(nil), "apiproto.ForgotPwdRequest")
	proto.RegisterType((*ForgotPayPwdRequest)(nil), "apiproto.ForgotPayPwdRequest")
	proto.RegisterType((*VerifyPayPwdRequest)(nil), "apiproto.VerifyPayPwdRequest")
	proto.RegisterType((*SetPwdRequest)(nil), "apiproto.SetPwdRequest")
	proto.RegisterType((*ResetPwdRequest)(nil), "apiproto.ResetPwdRequest")
	proto.RegisterType((*BindPhoneRequest)(nil), "apiproto.BindPhoneRequest")
	proto.RegisterType((*ChangePhoneRequest)(nil), "apiproto.ChangePhoneRequest")
	proto.RegisterType((*BindWeixinRequest)(nil), "apiproto.BindWeixinRequest")
	proto.RegisterType((*BindQQRequest)(nil), "apiproto.BindQQRequest")
	proto.RegisterType((*SetIconRequest)(nil), "apiproto.SetIconRequest")
	proto.RegisterType((*SetNicknameRequest)(nil), "apiproto.SetNicknameRequest")
	proto.RegisterType((*SetSexRequest)(nil), "apiproto.SetSexRequest")
	proto.RegisterType((*SetBoolRequest)(nil), "apiproto.SetBoolRequest")
	proto.RegisterType((*SetIntRequest)(nil), "apiproto.SetIntRequest")
	proto.RegisterType((*AuthRealnameRequest)(nil), "apiproto.AuthRealnameRequest")
	proto.RegisterType((*DailyCheckReply)(nil), "apiproto.DailyCheckReply")
	proto.RegisterEnum("apiproto.UserType", UserType_name, UserType_value)
	proto.RegisterEnum("apiproto.Sex", Sex_name, Sex_value)
}

func init() { proto.RegisterFile("user.proto", fileDescriptor14) }

var fileDescriptor14 = []byte{
	// 893 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdb, 0x6e, 0xdc, 0x36,
	0x10, 0x8d, 0xf6, 0x62, 0x6b, 0x67, 0x6f, 0x32, 0x1d, 0x14, 0x6a, 0xda, 0xa2, 0x1b, 0xb5, 0x49,
	0xb6, 0x2e, 0xe0, 0x87, 0xb4, 0x8f, 0x7d, 0xc9, 0x3a, 0x30, 0x60, 0xf4, 0xb6, 0xd6, 0xc6, 0x6d,
	0x81, 0x3e, 0x2c, 0x68, 0x89, 0xb6, 0x89, 0x95, 0x49, 0x79, 0xc9, 0xbd, 0xe8, 0x1f, 0xfa, 0xd2,
	0x6f, 0xe8, 0x8f, 0x16, 0x43, 0xea, 0xe6, 0x75, 0x90, 0xb8, 0x40, 0xdf, 0x38, 0x87, 0x47, 0xc3,
	0x23, 0x9e, 0x19, 0x0e, 0xc0, 0x4a, 0xb1, 0xe5, 0x71, 0xba, 0x94, 0x5a, 0x12, 0x97, 0xa6, 0xdc,
	0xac, 0x9e, 0xc1, 0xd5, 0x4a, 0xc4, 0x16, 0x0d, 0xfe, 0x6e, 0x81, 0x7b, 0xa1, 0xd8, 0xf2, 0x4c,
	0x5c, 0x49, 0xf2, 0x05, 0x00, 0x8d, 0x22, 0xb9, 0x12, 0x7a, 0xce, 0x63, 0xdf, 0x19, 0x39, 0xe3,
	0x66, 0xd8, 0xc9, 0x91, 0xb3, 0x98, 0x3c, 0x03, 0x57, 0xf0, 0x68, 0x21, 0xe8, 0x2d, 0xf3, 0x1b,
	0x23, 0x67, 0xdc, 0x09, 0xcb, 0x98, 0x10, 0x68, 0xf1, 0x48, 0x0a, 0xbf, 0x69, 0x70, 0xb3, 0x36,
	0x58, 0x2c, 0xa4, 0xdf, 0xca, 0xb1, 0x58, 0x48, 0xf2, 0x25, 0x34, 0x15, 0xdb, 0xfa, 0xed, 0x91,
	0x33, 0x1e, 0xbc, 0xee, 0x1f, 0x17, 0x9a, 0x8e, 0x67, 0x6c, 0x1b, 0xe2, 0x0e, 0xf1, 0xa0, 0xc9,
	0xb6, 0xa9, 0xbf, 0x37, 0x72, 0xc6, 0xed, 0x10, 0x97, 0xe4, 0x29, 0xb4, 0x13, 0xb6, 0x66, 0x89,
	0xbf, 0x6f, 0x30, 0x1b, 0xa0, 0x98, 0x25, 0xa3, 0x89, 0x11, 0xe3, 0x5a, 0x31, 0x45, 0x4c, 0x3e,
	0x87, 0xce, 0x25, 0x17, 0x71, 0x7a, 0x23, 0x05, 0xf3, 0xc1, 0x6c, 0x56, 0x00, 0xf9, 0x04, 0xf6,
	0x30, 0xb8, 0xbb, 0xf3, 0xbb, 0x66, 0x2b, 0x8f, 0x0a, 0x7c, 0xb3, 0xf5, 0x7b, 0x15, 0xbe, 0xd9,
	0x92, 0x57, 0x30, 0xe4, 0x62, 0xcd, 0x35, 0xd5, 0x5c, 0x8a, 0x79, 0x24, 0x63, 0xe6, 0xf7, 0x0d,
	0x61, 0x50, 0xc1, 0x27, 0x32, 0x66, 0xe4, 0x39, 0xf4, 0x52, 0x9a, 0xcd, 0x53, 0xaa, 0xd4, 0x46,
	0x2e, 0x63, 0x7f, 0x30, 0x72, 0xc6, 0x6e, 0xd8, 0x4d, 0x69, 0x36, 0xcd, 0x21, 0x72, 0x04, 0x07,
	0x48, 0x91, 0x29, 0x13, 0x15, 0x6f, 0x68, 0x78, 0xc3, 0x94, 0x66, 0xbf, 0xa6, 0x4c, 0x94, 0xdc,
	0x17, 0x30, 0x48, 0xe4, 0x35, 0xaf, 0x11, 0x3d, 0x43, 0xec, 0x1b, 0xb4, 0x96, 0xd2, 0xb3, 0x06,
	0xae, 0xb9, 0x66, 0x33, 0x4d, 0xf5, 0x4a, 0xf9, 0x07, 0x86, 0xf8, 0x00, 0x27, 0x5f, 0xc3, 0x80,
	0xab, 0x79, 0x4c, 0x79, 0x92, 0xcd, 0xa3, 0x1b, 0x16, 0x2d, 0x7c, 0x62, 0x98, 0x3d, 0xae, 0xde,
	0x22, 0x78, 0x82, 0x58, 0x70, 0x01, 0xfd, 0x90, 0x5d, 0x73, 0xa5, 0x43, 0x76, 0xb7, 0x62, 0x4a,
	0xa3, 0x03, 0xf6, 0x2e, 0x1d, 0xf3, 0xdf, 0x36, 0x40, 0x7b, 0xcd, 0x65, 0xd8, 0x52, 0x30, 0x6b,
	0x74, 0xa5, 0x54, 0x6b, 0x4b, 0xa1, 0x8c, 0x83, 0x7f, 0x1c, 0xe8, 0xfd, 0x84, 0xd2, 0x8b, 0xb4,
	0x2f, 0xa1, 0xa5, 0xb3, 0xd4, 0x66, 0x1d, 0xbc, 0x26, 0x55, 0x31, 0xa0, 0xee, 0x77, 0x59, 0xca,
	0x42, 0xb3, 0x5f, 0x1d, 0xdf, 0xa8, 0x1f, 0xff, 0x81, 0xa3, 0xd0, 0x4a, 0xbc, 0x62, 0x1e, 0xe7,
	0xb5, 0x97, 0x47, 0xe8, 0x10, 0x8d, 0x22, 0xa6, 0xd4, 0x5c, 0xcb, 0x05, 0x13, 0xa6, 0x0c, 0x3b,
	0x61, 0xd7, 0x62, 0xef, 0x10, 0x0a, 0xfe, 0x72, 0x00, 0x72, 0x95, 0x69, 0x92, 0xa1, 0x46, 0xec,
	0x21, 0xa3, 0xb1, 0xbb, 0xab, 0x11, 0x9b, 0x26, 0x34, 0xfb, 0xc8, 0xc3, 0xae, 0x32, 0x12, 0xef,
	0xf1, 0x4e, 0x57, 0x22, 0xb6, 0x3c, 0xdc, 0x27, 0xc7, 0xe0, 0x5e, 0x52, 0xb1, 0x88, 0x68, 0xae,
	0xfa, 0x1e, 0x77, 0x92, 0xef, 0x84, 0x25, 0x27, 0xf8, 0x03, 0xbc, 0x53, 0xb9, 0xbc, 0x96, 0x7a,
	0xba, 0x89, 0xff, 0x5f, 0x3b, 0xae, 0xe0, 0x30, 0xcf, 0x4c, 0xb3, 0x8f, 0x26, 0x7f, 0x58, 0x8b,
	0xf6, 0x98, 0x9d, 0x5a, 0x2c, 0x34, 0x34, 0x2b, 0x0d, 0xc1, 0x09, 0x1c, 0xfe, 0xc6, 0x96, 0xfc,
	0x2a, 0xbb, 0x7f, 0x4e, 0x5d, 0x9a, 0xb3, 0x63, 0x9f, 0x07, 0x4d, 0x9a, 0xf2, 0xfc, 0x08, 0x5c,
	0x06, 0x6f, 0xa0, 0x3f, 0x63, 0xfa, 0x91, 0x9f, 0x3f, 0x85, 0xb6, 0xb5, 0x37, 0xaf, 0x17, 0x13,
	0x04, 0x53, 0x18, 0x86, 0x4c, 0x3d, 0x3a, 0xc9, 0x73, 0xe8, 0x09, 0xb6, 0xd9, 0xfd, 0xdf, 0xae,
	0x60, 0x9b, 0xe2, 0x6f, 0x83, 0x1f, 0xc0, 0x9b, 0x70, 0x11, 0x4f, 0xf1, 0x86, 0xfe, 0xb3, 0x37,
	0xc1, 0x9f, 0x40, 0x4e, 0x6e, 0xa8, 0xb8, 0x66, 0x8f, 0xf8, 0xfe, 0x33, 0xe8, 0x18, 0x31, 0xb5,
	0x2e, 0x70, 0x51, 0xc9, 0xbd, 0xe4, 0xf5, 0x4b, 0xff, 0x16, 0x0e, 0x50, 0xda, 0xef, 0x8c, 0x6f,
	0xab, 0x7e, 0xab, 0xba, 0xc2, 0xa9, 0x77, 0x45, 0xf0, 0x0a, 0xfa, 0x48, 0x3e, 0x3f, 0xff, 0x18,
	0xf1, 0x25, 0x0c, 0x66, 0x4c, 0x9f, 0x45, 0x52, 0xd4, 0xe4, 0xae, 0x69, 0xb2, 0x2a, 0xe5, 0x9a,
	0x20, 0x38, 0x02, 0x32, 0x63, 0xfa, 0x97, 0x7c, 0x36, 0x7c, 0x98, 0xfb, 0xbd, 0x71, 0x16, 0x9f,
	0xff, 0x9c, 0xf6, 0x55, 0x9d, 0xf6, 0x60, 0x46, 0xe4, 0x5f, 0x59, 0x25, 0x13, 0x29, 0x93, 0xf7,
	0x66, 0x77, 0x0b, 0xde, 0x0b, 0x93, 0xfd, 0x4c, 0xe8, 0xf7, 0xd2, 0x9a, 0x05, 0x6d, 0x0d, 0x87,
	0x6f, 0x56, 0xfa, 0x26, 0xcc, 0x07, 0x48, 0xad, 0x3e, 0xca, 0x19, 0xe3, 0xec, 0xcc, 0x98, 0x62,
	0xb8, 0x35, 0x6a, 0xc3, 0xad, 0x34, 0xaf, 0x59, 0x37, 0xef, 0x53, 0x70, 0xd5, 0xad, 0xb2, 0x83,
	0xc3, 0x3e, 0x47, 0xfb, 0xea, 0x56, 0xe1, 0xc4, 0x08, 0x7e, 0x84, 0x61, 0xf5, 0xee, 0x16, 0x0f,
	0xce, 0x30, 0x92, 0x42, 0xdb, 0xe7, 0x79, 0x1e, 0xd3, 0x4c, 0x99, 0xa3, 0xdb, 0x61, 0x1f, 0x61,
	0x43, 0x7c, 0x4b, 0x33, 0x55, 0xcc, 0xc9, 0x46, 0x39, 0x27, 0x8f, 0xbe, 0xb1, 0x93, 0x1c, 0x1f,
	0x4e, 0xd2, 0x81, 0xb6, 0x29, 0x0e, 0xef, 0x09, 0xd9, 0x83, 0xc6, 0xf9, 0xb9, 0xe7, 0x10, 0x80,
	0x3d, 0x5b, 0x0e, 0x5e, 0xe3, 0x68, 0x0c, 0xcd, 0x19, 0xdb, 0x92, 0x2e, 0xec, 0x5f, 0x88, 0x85,
	0x90, 0x1b, 0xe1, 0x3d, 0x21, 0x2e, 0xb4, 0x7e, 0xa6, 0x09, 0xb3, 0xcc, 0x53, 0x76, 0x8b, 0xeb,
	0xc6, 0xa4, 0x0f, 0xed, 0x29, 0x5e, 0xfe, 0xa4, 0x85, 0xb9, 0x2f, 0xf7, 0x8c, 0x13, 0xdf, 0xfd,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0xae, 0x04, 0x16, 0x71, 0x59, 0x08, 0x00, 0x00,
}
