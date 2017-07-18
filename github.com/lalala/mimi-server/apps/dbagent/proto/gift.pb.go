// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gift.proto

package dbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GiftPackageType int32

const (
	GiftPackageType_cdkey       GiftPackageType = 0
	GiftPackageType_type_invite GiftPackageType = 1
	GiftPackageType_private     GiftPackageType = 2
	GiftPackageType_register    GiftPackageType = 3
)

var GiftPackageType_name = map[int32]string{
	0: "cdkey",
	1: "type_invite",
	2: "private",
	3: "register",
}
var GiftPackageType_value = map[string]int32{
	"cdkey":       0,
	"type_invite": 1,
	"private":     2,
	"register":    3,
}

func (x GiftPackageType) String() string {
	return proto.EnumName(GiftPackageType_name, int32(x))
}
func (GiftPackageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

type GiftPackage struct {
	Title            string `protobuf:"bytes,1,opt,name=title" json:"title"`
	ContentDesc      string `protobuf:"bytes,2,opt,name=content_desc,json=contentDesc" json:"content_desc"`
	Content          string `protobuf:"bytes,3,opt,name=content" json:"content"`
	TotalExchangeNum int64  `protobuf:"varint,4,opt,name=total_exchange_num,json=totalExchangeNum" json:"total_exchange_num"`
	GiftType         int64  `protobuf:"varint,5,opt,name=gift_type,json=giftType" json:"gift_type"`
	AddTime          int64  `protobuf:"varint,6,opt,name=add_time,json=addTime" json:"add_time"`
}

func (m *GiftPackage) Reset()                    { *m = GiftPackage{} }
func (m *GiftPackage) String() string            { return proto.CompactTextString(m) }
func (*GiftPackage) ProtoMessage()               {}
func (*GiftPackage) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *GiftPackage) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GiftPackage) GetContentDesc() string {
	if m != nil {
		return m.ContentDesc
	}
	return ""
}

func (m *GiftPackage) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *GiftPackage) GetTotalExchangeNum() int64 {
	if m != nil {
		return m.TotalExchangeNum
	}
	return 0
}

func (m *GiftPackage) GetGiftType() int64 {
	if m != nil {
		return m.GiftType
	}
	return 0
}

func (m *GiftPackage) GetAddTime() int64 {
	if m != nil {
		return m.AddTime
	}
	return 0
}

type GiftArg struct {
	TicketList []*BuycaiTicket       `protobuf:"bytes,1,rep,name=TicketList" json:"TicketList"`
	Credits    *ChangeVirtualFundArg `protobuf:"bytes,2,opt,name=Credits" json:"Credits"`
	CdKey      string                `protobuf:"bytes,3,opt,name=CdKey" json:"CdKey"`
	AccountId  int64                 `protobuf:"varint,4,opt,name=AccountId" json:"AccountId"`
	Batch      int64                 `protobuf:"varint,5,opt,name=Batch" json:"Batch"`
}

func (m *GiftArg) Reset()                    { *m = GiftArg{} }
func (m *GiftArg) String() string            { return proto.CompactTextString(m) }
func (*GiftArg) ProtoMessage()               {}
func (*GiftArg) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *GiftArg) GetTicketList() []*BuycaiTicket {
	if m != nil {
		return m.TicketList
	}
	return nil
}

func (m *GiftArg) GetCredits() *ChangeVirtualFundArg {
	if m != nil {
		return m.Credits
	}
	return nil
}

func (m *GiftArg) GetCdKey() string {
	if m != nil {
		return m.CdKey
	}
	return ""
}

func (m *GiftArg) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *GiftArg) GetBatch() int64 {
	if m != nil {
		return m.Batch
	}
	return 0
}

type CdKeyBatchArg struct {
	Id            int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	MaxExchange   int64  `protobuf:"varint,2,opt,name=Max_exchange,json=MaxExchange" json:"Max_exchange"`
	Title         string `protobuf:"bytes,3,opt,name=title" json:"title"`
	GiftPackageId int64  `protobuf:"varint,4,opt,name=Gift_package_id,json=GiftPackageId" json:"Gift_package_id"`
	ValidStart    int64  `protobuf:"varint,5,opt,name=Valid_start,json=ValidStart" json:"Valid_start"`
	ValidEnd      int64  `protobuf:"varint,6,opt,name=Valid_end,json=ValidEnd" json:"Valid_end"`
	AddTime       int64  `protobuf:"varint,7,opt,name=add_time,json=addTime" json:"add_time"`
}

func (m *CdKeyBatchArg) Reset()                    { *m = CdKeyBatchArg{} }
func (m *CdKeyBatchArg) String() string            { return proto.CompactTextString(m) }
func (*CdKeyBatchArg) ProtoMessage()               {}
func (*CdKeyBatchArg) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *CdKeyBatchArg) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CdKeyBatchArg) GetMaxExchange() int64 {
	if m != nil {
		return m.MaxExchange
	}
	return 0
}

func (m *CdKeyBatchArg) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CdKeyBatchArg) GetGiftPackageId() int64 {
	if m != nil {
		return m.GiftPackageId
	}
	return 0
}

func (m *CdKeyBatchArg) GetValidStart() int64 {
	if m != nil {
		return m.ValidStart
	}
	return 0
}

func (m *CdKeyBatchArg) GetValidEnd() int64 {
	if m != nil {
		return m.ValidEnd
	}
	return 0
}

func (m *CdKeyBatchArg) GetAddTime() int64 {
	if m != nil {
		return m.AddTime
	}
	return 0
}

type GiftCdkeyArg struct {
	Gift  *GiftPackage   `protobuf:"bytes,1,opt,name=gift" json:"gift"`
	Cdkey *CdKeyBatchArg `protobuf:"bytes,2,opt,name=cdkey" json:"cdkey"`
}

func (m *GiftCdkeyArg) Reset()                    { *m = GiftCdkeyArg{} }
func (m *GiftCdkeyArg) String() string            { return proto.CompactTextString(m) }
func (*GiftCdkeyArg) ProtoMessage()               {}
func (*GiftCdkeyArg) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

func (m *GiftCdkeyArg) GetGift() *GiftPackage {
	if m != nil {
		return m.Gift
	}
	return nil
}

func (m *GiftCdkeyArg) GetCdkey() *CdKeyBatchArg {
	if m != nil {
		return m.Cdkey
	}
	return nil
}

type InviteArg struct {
	Code      string `protobuf:"bytes,1,opt,name=Code" json:"Code"`
	InviteeId int64  `protobuf:"varint,2,opt,name=InviteeId" json:"InviteeId"`
}

func (m *InviteArg) Reset()                    { *m = InviteArg{} }
func (m *InviteArg) String() string            { return proto.CompactTextString(m) }
func (*InviteArg) ProtoMessage()               {}
func (*InviteArg) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

func (m *InviteArg) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *InviteArg) GetInviteeId() int64 {
	if m != nil {
		return m.InviteeId
	}
	return 0
}

type InviteRelationArg struct {
	Inviter int64 `protobuf:"varint,1,opt,name=inviter" json:"inviter"`
	Invitee int64 `protobuf:"varint,2,opt,name=invitee" json:"invitee"`
}

func (m *InviteRelationArg) Reset()                    { *m = InviteRelationArg{} }
func (m *InviteRelationArg) String() string            { return proto.CompactTextString(m) }
func (*InviteRelationArg) ProtoMessage()               {}
func (*InviteRelationArg) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{5} }

func (m *InviteRelationArg) GetInviter() int64 {
	if m != nil {
		return m.Inviter
	}
	return 0
}

func (m *InviteRelationArg) GetInvitee() int64 {
	if m != nil {
		return m.Invitee
	}
	return 0
}

type CreditsArg struct {
	Credits   int64 `protobuf:"varint,1,opt,name=Credits" json:"Credits"`
	AccountId int64 `protobuf:"varint,2,opt,name=AccountId" json:"AccountId"`
}

func (m *CreditsArg) Reset()                    { *m = CreditsArg{} }
func (m *CreditsArg) String() string            { return proto.CompactTextString(m) }
func (*CreditsArg) ProtoMessage()               {}
func (*CreditsArg) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{6} }

func (m *CreditsArg) GetCredits() int64 {
	if m != nil {
		return m.Credits
	}
	return 0
}

func (m *CreditsArg) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

type Gift struct {
	Credits    int64           `protobuf:"varint,1,opt,name=Credits" json:"Credits"`
	TicketList []*BuycaiTicket `protobuf:"bytes,2,rep,name=TicketList" json:"TicketList"`
	Inviter    int64           `protobuf:"varint,3,opt,name=Inviter" json:"Inviter"`
	Invitee    int64           `protobuf:"varint,4,opt,name=invitee" json:"invitee"`
}

func (m *Gift) Reset()                    { *m = Gift{} }
func (m *Gift) String() string            { return proto.CompactTextString(m) }
func (*Gift) ProtoMessage()               {}
func (*Gift) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{7} }

func (m *Gift) GetCredits() int64 {
	if m != nil {
		return m.Credits
	}
	return 0
}

func (m *Gift) GetTicketList() []*BuycaiTicket {
	if m != nil {
		return m.TicketList
	}
	return nil
}

func (m *Gift) GetInviter() int64 {
	if m != nil {
		return m.Inviter
	}
	return 0
}

func (m *Gift) GetInvitee() int64 {
	if m != nil {
		return m.Invitee
	}
	return 0
}

type GiftPackageArg struct {
	Tickets       []*BuycaiTicket `protobuf:"bytes,1,rep,name=Tickets" json:"Tickets"`
	Credits       int64           `protobuf:"varint,2,opt,name=Credits" json:"Credits"`
	GiftPackageId int64           `protobuf:"varint,3,opt,name=gift_package_id,json=giftPackageId" json:"gift_package_id"`
}

func (m *GiftPackageArg) Reset()                    { *m = GiftPackageArg{} }
func (m *GiftPackageArg) String() string            { return proto.CompactTextString(m) }
func (*GiftPackageArg) ProtoMessage()               {}
func (*GiftPackageArg) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{8} }

func (m *GiftPackageArg) GetTickets() []*BuycaiTicket {
	if m != nil {
		return m.Tickets
	}
	return nil
}

func (m *GiftPackageArg) GetCredits() int64 {
	if m != nil {
		return m.Credits
	}
	return 0
}

func (m *GiftPackageArg) GetGiftPackageId() int64 {
	if m != nil {
		return m.GiftPackageId
	}
	return 0
}

type UserInviteInfo struct {
	InviteNum        int64 `protobuf:"varint,1,opt,name=invite_num,json=inviteNum" json:"invite_num"`
	Credits          int64 `protobuf:"varint,2,opt,name=credits" json:"credits"`
	TicketsNum       int64 `protobuf:"varint,3,opt,name=tickets_num,json=ticketsNum" json:"tickets_num"`
	TicketsMoney     int64 `protobuf:"varint,4,opt,name=tickets_money,json=ticketsMoney" json:"tickets_money"`
	UserInviteStatus bool  `protobuf:"varint,5,opt,name=UserInviteStatus" json:"UserInviteStatus"`
}

func (m *UserInviteInfo) Reset()                    { *m = UserInviteInfo{} }
func (m *UserInviteInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInviteInfo) ProtoMessage()               {}
func (*UserInviteInfo) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{9} }

func (m *UserInviteInfo) GetInviteNum() int64 {
	if m != nil {
		return m.InviteNum
	}
	return 0
}

func (m *UserInviteInfo) GetCredits() int64 {
	if m != nil {
		return m.Credits
	}
	return 0
}

func (m *UserInviteInfo) GetTicketsNum() int64 {
	if m != nil {
		return m.TicketsNum
	}
	return 0
}

func (m *UserInviteInfo) GetTicketsMoney() int64 {
	if m != nil {
		return m.TicketsMoney
	}
	return 0
}

func (m *UserInviteInfo) GetUserInviteStatus() bool {
	if m != nil {
		return m.UserInviteStatus
	}
	return false
}

type PhoneRegistGift struct {
	AccountId int64           `protobuf:"varint,1,opt,name=account_id,json=accountId" json:"account_id"`
	Credits   int32           `protobuf:"varint,2,opt,name=credits" json:"credits"`
	List      []*BuycaiTicket `protobuf:"bytes,3,rep,name=list" json:"list"`
}

func (m *PhoneRegistGift) Reset()                    { *m = PhoneRegistGift{} }
func (m *PhoneRegistGift) String() string            { return proto.CompactTextString(m) }
func (*PhoneRegistGift) ProtoMessage()               {}
func (*PhoneRegistGift) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{10} }

func (m *PhoneRegistGift) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *PhoneRegistGift) GetCredits() int32 {
	if m != nil {
		return m.Credits
	}
	return 0
}

func (m *PhoneRegistGift) GetList() []*BuycaiTicket {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*GiftPackage)(nil), "dbproto.GiftPackage")
	proto.RegisterType((*GiftArg)(nil), "dbproto.GiftArg")
	proto.RegisterType((*CdKeyBatchArg)(nil), "dbproto.CdKeyBatchArg")
	proto.RegisterType((*GiftCdkeyArg)(nil), "dbproto.GiftCdkeyArg")
	proto.RegisterType((*InviteArg)(nil), "dbproto.InviteArg")
	proto.RegisterType((*InviteRelationArg)(nil), "dbproto.InviteRelationArg")
	proto.RegisterType((*CreditsArg)(nil), "dbproto.CreditsArg")
	proto.RegisterType((*Gift)(nil), "dbproto.Gift")
	proto.RegisterType((*GiftPackageArg)(nil), "dbproto.GiftPackageArg")
	proto.RegisterType((*UserInviteInfo)(nil), "dbproto.UserInviteInfo")
	proto.RegisterType((*PhoneRegistGift)(nil), "dbproto.PhoneRegistGift")
	proto.RegisterEnum("dbproto.GiftPackageType", GiftPackageType_name, GiftPackageType_value)
}

func init() { proto.RegisterFile("gift.proto", fileDescriptor14) }

var fileDescriptor14 = []byte{
	// 755 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0xbe, 0x8e, 0x13, 0x9c, 0x1c, 0x27, 0x24, 0x77, 0xc4, 0xbd, 0xf2, 0xbd, 0x2d, 0x2a, 0x75,
	0x25, 0x94, 0x22, 0x44, 0x25, 0xaa, 0xaa, 0xab, 0x2e, 0x20, 0x50, 0x6a, 0xb5, 0x54, 0xc8, 0x50,
	0xb6, 0xd6, 0xe0, 0x99, 0x98, 0x11, 0xc9, 0x38, 0xb2, 0xc7, 0x88, 0xac, 0xbb, 0xec, 0x23, 0x75,
	0x59, 0xa9, 0xef, 0xd0, 0xb7, 0xa9, 0xe6, 0xc7, 0xb1, 0x0d, 0x12, 0xea, 0xce, 0xe7, 0x9b, 0x33,
	0xe3, 0xef, 0xe7, 0xcc, 0x00, 0x24, 0x6c, 0x2a, 0xf6, 0x16, 0x59, 0x2a, 0x52, 0xe4, 0x90, 0x2b,
	0xf5, 0xf1, 0x3f, 0x4c, 0x0b, 0x4e, 0x34, 0xe8, 0xff, 0xb4, 0xc0, 0x3d, 0x61, 0x53, 0x71, 0x86,
	0xe3, 0x1b, 0x9c, 0x50, 0xb4, 0x01, 0x1d, 0xc1, 0xc4, 0x8c, 0x7a, 0xd6, 0x96, 0x35, 0xee, 0x85,
	0xba, 0x40, 0xcf, 0xa1, 0x1f, 0xa7, 0x5c, 0x50, 0x2e, 0x22, 0x42, 0xf3, 0xd8, 0x6b, 0xa9, 0x45,
	0xd7, 0x60, 0x47, 0x34, 0x8f, 0x91, 0x07, 0x8e, 0x29, 0x3d, 0x5b, 0xad, 0x96, 0x25, 0xda, 0x05,
	0x24, 0x52, 0x81, 0x67, 0x11, 0xbd, 0x8b, 0xaf, 0x31, 0x4f, 0x68, 0xc4, 0x8b, 0xb9, 0xd7, 0xde,
	0xb2, 0xc6, 0x76, 0x38, 0x52, 0x2b, 0xc7, 0x66, 0xe1, 0x73, 0x31, 0x47, 0x4f, 0xa0, 0x27, 0x39,
	0x47, 0x62, 0xb9, 0xa0, 0x5e, 0x47, 0x35, 0x75, 0x25, 0x70, 0xb1, 0x5c, 0x50, 0xf4, 0x1f, 0x74,
	0x31, 0x21, 0x91, 0x60, 0x73, 0xea, 0xad, 0xa9, 0x35, 0x07, 0x13, 0x72, 0xc1, 0xe6, 0xd4, 0xff,
	0x61, 0x81, 0x23, 0x85, 0x1c, 0x64, 0x09, 0x7a, 0x03, 0x70, 0xc1, 0xe2, 0x1b, 0x2a, 0x3e, 0xb1,
	0x5c, 0x78, 0xd6, 0x96, 0x3d, 0x76, 0xf7, 0xff, 0xd9, 0x33, 0xf2, 0xf7, 0x0e, 0x8b, 0x65, 0x8c,
	0x99, 0x6e, 0x08, 0x6b, 0x8d, 0xe8, 0x2d, 0x38, 0x93, 0x8c, 0x12, 0x26, 0x72, 0x25, 0xd0, 0xdd,
	0xdf, 0x5c, 0xed, 0x99, 0x28, 0x7e, 0x97, 0x2c, 0x13, 0x05, 0x9e, 0xbd, 0x2f, 0x38, 0x39, 0xc8,
	0x92, 0xb0, 0xec, 0x96, 0xa6, 0x4d, 0xc8, 0x47, 0xba, 0x34, 0xca, 0x75, 0x81, 0x9e, 0x42, 0xef,
	0x20, 0x8e, 0xd3, 0x82, 0x8b, 0x80, 0x18, 0xb9, 0x15, 0x20, 0xf7, 0x1c, 0x62, 0x11, 0x5f, 0x1b,
	0x8d, 0xba, 0xf0, 0x7f, 0x59, 0x30, 0x50, 0xbb, 0x55, 0x29, 0xb5, 0xac, 0x43, 0x8b, 0x11, 0x95,
	0x86, 0x1d, 0xb6, 0x18, 0x91, 0x51, 0x9c, 0xe2, 0xbb, 0x95, 0x97, 0x8a, 0xa9, 0x1d, 0xba, 0xa7,
	0xf8, 0xae, 0x74, 0xb1, 0xca, 0xd0, 0xae, 0x67, 0xb8, 0x0d, 0x43, 0xe9, 0x4f, 0xb4, 0xd0, 0x49,
	0x47, 0xac, 0x24, 0x35, 0xa8, 0xe5, 0x1f, 0x10, 0xf4, 0x0c, 0xdc, 0x4b, 0x3c, 0x63, 0x24, 0xca,
	0x05, 0xce, 0x84, 0xa1, 0x07, 0x0a, 0x3a, 0x97, 0x88, 0x4c, 0x48, 0x37, 0x50, 0x4e, 0x4c, 0x0a,
	0x5d, 0x05, 0x1c, 0x73, 0xd2, 0x48, 0xc8, 0x69, 0x26, 0x34, 0x85, 0xbe, 0xfc, 0xd3, 0x84, 0xdc,
	0xd0, 0xa5, 0x54, 0x36, 0x86, 0xb6, 0x0c, 0x56, 0x69, 0x73, 0xf7, 0x37, 0x56, 0x5e, 0xd7, 0xe8,
	0x84, 0xaa, 0x03, 0xed, 0x42, 0x27, 0x96, 0xbb, 0x4c, 0x2c, 0xff, 0x56, 0xb1, 0xd4, 0xad, 0x0a,
	0x75, 0x93, 0xff, 0x0e, 0x7a, 0x01, 0xbf, 0x65, 0x82, 0xca, 0x9f, 0x20, 0x68, 0x4f, 0x52, 0x52,
	0x8e, 0xb3, 0xfa, 0x96, 0xc1, 0xe8, 0x06, 0x1a, 0x10, 0xe3, 0x5f, 0x05, 0xf8, 0x27, 0xf0, 0xb7,
	0x2e, 0x42, 0x3a, 0xc3, 0x82, 0xa5, 0x5c, 0x1e, 0xe3, 0x81, 0xc3, 0x14, 0x98, 0x99, 0x28, 0xca,
	0xb2, 0x5a, 0x29, 0xa3, 0x28, 0x4b, 0xff, 0x08, 0xc0, 0x0c, 0x88, 0x39, 0xa1, 0x1c, 0x2e, 0x73,
	0x42, 0x39, 0x3d, 0x8d, 0x39, 0x69, 0xdd, 0x9b, 0x13, 0xff, 0x9b, 0x05, 0x6d, 0xe9, 0xc8, 0x23,
	0x07, 0x34, 0xc7, 0xbd, 0xf5, 0xa7, 0xe3, 0xee, 0x81, 0x13, 0x18, 0x4d, 0xb6, 0x3e, 0x30, 0x78,
	0xa8, 0xa9, 0xdd, 0xd4, 0xf4, 0xd5, 0x82, 0xf5, 0x5a, 0x3e, 0x52, 0xd8, 0x2b, 0x70, 0xf4, 0xa1,
	0xf9, 0xe3, 0x37, 0xad, 0xec, 0xaa, 0x0b, 0x69, 0x35, 0x85, 0x6c, 0xc3, 0x30, 0xb9, 0x37, 0xa2,
	0x9a, 0xd9, 0x20, 0xa9, 0x8f, 0xa8, 0xff, 0xdd, 0x82, 0xf5, 0x2f, 0x39, 0xcd, 0x34, 0xdf, 0x80,
	0x4f, 0x53, 0xb4, 0x09, 0xa0, 0x39, 0xaa, 0xc7, 0x45, 0x1b, 0xd4, 0xd3, 0x88, 0x7c, 0x55, 0xe4,
	0xeb, 0xd4, 0xfc, 0xa7, 0x29, 0xe5, 0xb8, 0x0b, 0x4d, 0x4c, 0xed, 0xd4, 0xff, 0x03, 0x03, 0xc9,
	0xad, 0x2f, 0x60, 0x50, 0x36, 0xcc, 0x53, 0x4e, 0x97, 0xc6, 0x92, 0xbe, 0x01, 0x4f, 0x25, 0x86,
	0x76, 0x60, 0x54, 0x11, 0x3a, 0x17, 0x58, 0x14, 0xb9, 0xba, 0x39, 0xdd, 0xf0, 0x01, 0xee, 0x17,
	0x30, 0x3c, 0xbb, 0x4e, 0x39, 0x0d, 0x69, 0xc2, 0x72, 0xa1, 0xb2, 0xdd, 0x04, 0xc0, 0x3a, 0xf1,
	0x68, 0x75, 0xd9, 0x7b, 0x78, 0xf5, 0x56, 0xdc, 0x63, 0xdf, 0xa9, 0xd8, 0xbf, 0x84, 0xf6, 0x4c,
	0x86, 0x6e, 0x3f, 0xe6, 0xbc, 0x6a, 0xd9, 0xf9, 0xa0, 0xef, 0xbf, 0x71, 0x51, 0x3d, 0xa7, 0x3d,
	0x73, 0xaf, 0x46, 0x7f, 0xa1, 0x21, 0xb8, 0xf2, 0xc5, 0x8d, 0xb4, 0x65, 0x23, 0x0b, 0xb9, 0xe0,
	0x2c, 0x32, 0x76, 0x8b, 0x05, 0x1d, 0xb5, 0x50, 0x1f, 0xba, 0x99, 0x62, 0x4b, 0xb3, 0x91, 0x7d,
	0xb5, 0xa6, 0xfe, 0xf1, 0xfa, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xb5, 0x61, 0x5d, 0x5d,
	0x06, 0x00, 0x00,
}