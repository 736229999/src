// Code generated by protoc-gen-go. DO NOT EDIT.
// source: buycai_vendor_order.proto

package dbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BuycaiVendorOrder struct {
	Id             int64           `protobuf:"varint,1,opt,name=id" json:"id"`
	AccountId      int64           `protobuf:"varint,2,opt,name=account_id,json=accountId" json:"account_id"`
	UserOrderId    int64           `protobuf:"varint,3,opt,name=user_order_id,json=userOrderId" json:"user_order_id"`
	LotteryId      int32           `protobuf:"varint,4,opt,name=lottery_id,json=lotteryId" json:"lottery_id"`
	Issue          string          `protobuf:"bytes,5,opt,name=issue" json:"issue"`
	SumNum         int32           `protobuf:"varint,6,opt,name=sum_num,json=sumNum" json:"sum_num"`
	Multiple       int32           `protobuf:"varint,7,opt,name=multiple" json:"multiple"`
	Money          float64         `protobuf:"fixed64,8,opt,name=money" json:"money"`
	Cai            float64         `protobuf:"fixed64,9,opt,name=cai" json:"cai"`
	Balance        float64         `protobuf:"fixed64,10,opt,name=balance" json:"balance"`
	ChaseNo        int32           `protobuf:"varint,11,opt,name=chase_no,json=chaseNo" json:"chase_no"`
	Vendor         string          `protobuf:"bytes,12,opt,name=vendor" json:"vendor"`
	SchemeList     []*BuycaiScheme `protobuf:"bytes,13,rep,name=scheme_list,json=schemeList" json:"scheme_list"`
	AddTime        int64           `protobuf:"varint,14,opt,name=add_time,json=addTime" json:"add_time"`
	Status         int32           `protobuf:"varint,15,opt,name=status" json:"status"`
	StatusDesc     string          `protobuf:"bytes,16,opt,name=status_desc,json=statusDesc" json:"status_desc"`
	VendorReqTime  int64           `protobuf:"varint,17,opt,name=vendor_req_time,json=vendorReqTime" json:"vendor_req_time"`
	VendorRespTime int64           `protobuf:"varint,18,opt,name=vendor_resp_time,json=vendorRespTime" json:"vendor_resp_time"`
	VendorRespId   string          `protobuf:"bytes,19,opt,name=vendor_resp_id,json=vendorRespId" json:"vendor_resp_id"`
	WinMoney       float64         `protobuf:"fixed64,20,opt,name=win_money,json=winMoney" json:"win_money"`
}

func (m *BuycaiVendorOrder) Reset()                    { *m = BuycaiVendorOrder{} }
func (m *BuycaiVendorOrder) String() string            { return proto.CompactTextString(m) }
func (*BuycaiVendorOrder) ProtoMessage()               {}
func (*BuycaiVendorOrder) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *BuycaiVendorOrder) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BuycaiVendorOrder) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *BuycaiVendorOrder) GetUserOrderId() int64 {
	if m != nil {
		return m.UserOrderId
	}
	return 0
}

func (m *BuycaiVendorOrder) GetLotteryId() int32 {
	if m != nil {
		return m.LotteryId
	}
	return 0
}

func (m *BuycaiVendorOrder) GetIssue() string {
	if m != nil {
		return m.Issue
	}
	return ""
}

func (m *BuycaiVendorOrder) GetSumNum() int32 {
	if m != nil {
		return m.SumNum
	}
	return 0
}

func (m *BuycaiVendorOrder) GetMultiple() int32 {
	if m != nil {
		return m.Multiple
	}
	return 0
}

func (m *BuycaiVendorOrder) GetMoney() float64 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *BuycaiVendorOrder) GetCai() float64 {
	if m != nil {
		return m.Cai
	}
	return 0
}

func (m *BuycaiVendorOrder) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *BuycaiVendorOrder) GetChaseNo() int32 {
	if m != nil {
		return m.ChaseNo
	}
	return 0
}

func (m *BuycaiVendorOrder) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *BuycaiVendorOrder) GetSchemeList() []*BuycaiScheme {
	if m != nil {
		return m.SchemeList
	}
	return nil
}

func (m *BuycaiVendorOrder) GetAddTime() int64 {
	if m != nil {
		return m.AddTime
	}
	return 0
}

func (m *BuycaiVendorOrder) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *BuycaiVendorOrder) GetStatusDesc() string {
	if m != nil {
		return m.StatusDesc
	}
	return ""
}

func (m *BuycaiVendorOrder) GetVendorReqTime() int64 {
	if m != nil {
		return m.VendorReqTime
	}
	return 0
}

func (m *BuycaiVendorOrder) GetVendorRespTime() int64 {
	if m != nil {
		return m.VendorRespTime
	}
	return 0
}

func (m *BuycaiVendorOrder) GetVendorRespId() string {
	if m != nil {
		return m.VendorRespId
	}
	return ""
}

func (m *BuycaiVendorOrder) GetWinMoney() float64 {
	if m != nil {
		return m.WinMoney
	}
	return 0
}

type BuycaiScheme struct {
	Type        string            `protobuf:"bytes,1,opt,name=type" json:"type"`
	SelectBalls map[string]*Balls `protobuf:"bytes,2,rep,name=select_balls,json=selectBalls" json:"select_balls" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Num         int32             `protobuf:"varint,3,opt,name=num" json:"num"`
	Money       float64           `protobuf:"fixed64,4,opt,name=money" json:"money"`
}

func (m *BuycaiScheme) Reset()                    { *m = BuycaiScheme{} }
func (m *BuycaiScheme) String() string            { return proto.CompactTextString(m) }
func (*BuycaiScheme) ProtoMessage()               {}
func (*BuycaiScheme) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *BuycaiScheme) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *BuycaiScheme) GetSelectBalls() map[string]*Balls {
	if m != nil {
		return m.SelectBalls
	}
	return nil
}

func (m *BuycaiScheme) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *BuycaiScheme) GetMoney() float64 {
	if m != nil {
		return m.Money
	}
	return 0
}

type Balls struct {
	List []int32 `protobuf:"varint,1,rep,packed,name=list" json:"list"`
}

func (m *Balls) Reset()                    { *m = Balls{} }
func (m *Balls) String() string            { return proto.CompactTextString(m) }
func (*Balls) ProtoMessage()               {}
func (*Balls) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *Balls) GetList() []int32 {
	if m != nil {
		return m.List
	}
	return nil
}

type IssueMultiple struct {
	Issue    string `protobuf:"bytes,1,opt,name=issue" json:"issue"`
	Multiple int32  `protobuf:"varint,2,opt,name=multiple" json:"multiple"`
}

func (m *IssueMultiple) Reset()                    { *m = IssueMultiple{} }
func (m *IssueMultiple) String() string            { return proto.CompactTextString(m) }
func (*IssueMultiple) ProtoMessage()               {}
func (*IssueMultiple) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *IssueMultiple) GetIssue() string {
	if m != nil {
		return m.Issue
	}
	return ""
}

func (m *IssueMultiple) GetMultiple() int32 {
	if m != nil {
		return m.Multiple
	}
	return 0
}

func init() {
	proto.RegisterType((*BuycaiVendorOrder)(nil), "dbproto.BuycaiVendorOrder")
	proto.RegisterType((*BuycaiScheme)(nil), "dbproto.BuycaiScheme")
	proto.RegisterType((*Balls)(nil), "dbproto.Balls")
	proto.RegisterType((*IssueMultiple)(nil), "dbproto.IssueMultiple")
}

func init() { proto.RegisterFile("buycai_vendor_order.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x65, 0x3b, 0x8e, 0xe3, 0x71, 0x92, 0xa6, 0x4b, 0x81, 0x6d, 0x2b, 0x84, 0x15, 0x55,
	0x95, 0x4f, 0x39, 0x14, 0x09, 0x21, 0x6e, 0x54, 0x70, 0xb0, 0x44, 0x83, 0xe4, 0x22, 0xae, 0x96,
	0xe3, 0x5d, 0xa9, 0x2b, 0xfc, 0x27, 0xf5, 0xee, 0xb6, 0xf2, 0x07, 0x86, 0xcf, 0x81, 0x76, 0xd6,
	0x6d, 0x0c, 0xe2, 0x36, 0xf3, 0xf2, 0xdb, 0x99, 0xc9, 0xf3, 0x83, 0xd3, 0x9d, 0xee, 0xcb, 0x42,
	0xe4, 0x0f, 0xbc, 0x61, 0x6d, 0x97, 0xb7, 0x1d, 0xe3, 0xdd, 0x66, 0xdf, 0xb5, 0xaa, 0x25, 0x01,
	0xdb, 0x61, 0xb1, 0xfe, 0x3d, 0x81, 0xe3, 0x6b, 0xc4, 0x7e, 0x20, 0xf5, 0xcd, 0x40, 0x64, 0x09,
	0xae, 0x60, 0xd4, 0x89, 0x9d, 0xc4, 0xcb, 0x5c, 0xc1, 0xc8, 0x1b, 0x80, 0xa2, 0x2c, 0x5b, 0xdd,
	0xa8, 0x5c, 0x30, 0xea, 0xa2, 0x1e, 0x0e, 0x4a, 0xca, 0xc8, 0x1a, 0x16, 0x5a, 0xf2, 0x61, 0x83,
	0x21, 0x3c, 0x24, 0x22, 0x23, 0xe2, 0xc0, 0x14, 0x47, 0x54, 0xad, 0x52, 0xbc, 0xeb, 0x0d, 0x30,
	0x89, 0x9d, 0xc4, 0xcf, 0xc2, 0x41, 0x49, 0x19, 0x39, 0x01, 0x5f, 0x48, 0xa9, 0x39, 0xf5, 0x63,
	0x27, 0x09, 0x33, 0xdb, 0x90, 0xd7, 0x10, 0x48, 0x5d, 0xe7, 0x8d, 0xae, 0xe9, 0x14, 0x5f, 0x4c,
	0xa5, 0xae, 0xb7, 0xba, 0x26, 0x67, 0x30, 0xab, 0x75, 0xa5, 0xc4, 0xbe, 0xe2, 0x34, 0xc0, 0x5f,
	0x9e, 0x7b, 0x33, 0xaa, 0x6e, 0x1b, 0xde, 0xd3, 0x59, 0xec, 0x24, 0x4e, 0x66, 0x1b, 0xb2, 0x02,
	0xaf, 0x2c, 0x04, 0x0d, 0x51, 0x33, 0x25, 0xa1, 0x10, 0xec, 0x8a, 0xaa, 0x68, 0x4a, 0x4e, 0x01,
	0xd5, 0xa7, 0x96, 0x9c, 0xc2, 0xac, 0xbc, 0x2b, 0x24, 0xcf, 0x9b, 0x96, 0x46, 0x38, 0x3d, 0xc0,
	0x7e, 0xdb, 0x92, 0x57, 0x30, 0xb5, 0x76, 0xd2, 0x39, 0x1e, 0x3a, 0x74, 0xe4, 0x3d, 0x44, 0xb2,
	0xbc, 0xe3, 0x35, 0xcf, 0x2b, 0x21, 0x15, 0x5d, 0xc4, 0x5e, 0x12, 0x5d, 0xbd, 0xdc, 0x0c, 0x36,
	0x6f, 0xac, 0xc5, 0xb7, 0x48, 0x64, 0x60, 0xc9, 0xaf, 0x42, 0x2a, 0xb3, 0xaa, 0x60, 0x2c, 0x57,
	0xa2, 0xe6, 0x74, 0x89, 0xae, 0x05, 0x05, 0x63, 0xdf, 0x45, 0xcd, 0xcd, 0x2a, 0xa9, 0x0a, 0xa5,
	0x25, 0x3d, 0x1a, 0xfe, 0x3b, 0x76, 0xe4, 0x2d, 0x44, 0xb6, 0xca, 0x19, 0x97, 0x25, 0x5d, 0xe1,
	0x1d, 0x60, 0xa5, 0xcf, 0x5c, 0x96, 0xe4, 0x12, 0x8e, 0x86, 0x4f, 0xde, 0xf1, 0x7b, 0x3b, 0xfa,
	0x18, 0x47, 0x2f, 0xac, 0x9c, 0xf1, 0x7b, 0x5c, 0x90, 0xc0, 0xea, 0x99, 0x93, 0x7b, 0x0b, 0x12,
	0x04, 0x97, 0x4f, 0xa0, 0xdc, 0x23, 0x79, 0x01, 0xcb, 0x31, 0x29, 0x18, 0x7d, 0x81, 0x5b, 0xe7,
	0x07, 0x2e, 0x65, 0xe4, 0x1c, 0xc2, 0x47, 0xd1, 0xe4, 0xd6, 0xfc, 0x13, 0xb4, 0x74, 0xf6, 0x28,
	0x9a, 0x1b, 0xd3, 0xaf, 0x7f, 0x39, 0x30, 0x1f, 0xbb, 0x40, 0x08, 0x4c, 0x54, 0xbf, 0xe7, 0x98,
	0xb2, 0x30, 0xc3, 0x9a, 0xa4, 0x30, 0x97, 0xbc, 0xe2, 0xa5, 0xca, 0x77, 0x45, 0x55, 0x49, 0xea,
	0xa2, 0x8d, 0x97, 0xff, 0xb5, 0x71, 0x73, 0x8b, 0xe4, 0xb5, 0x01, 0xbf, 0x34, 0xaa, 0xeb, 0xb3,
	0x48, 0x1e, 0x14, 0xf3, 0xbd, 0x4d, 0x6c, 0x3c, 0xb4, 0xce, 0x94, 0x87, 0x5c, 0x4c, 0x46, 0xb9,
	0x38, 0xdb, 0xc2, 0xea, 0xdf, 0x41, 0xe6, 0xed, 0x4f, 0xde, 0x0f, 0x97, 0x99, 0x92, 0x5c, 0x80,
	0xff, 0x50, 0x54, 0x9a, 0x63, 0xf6, 0xa3, 0xab, 0xe5, 0xe1, 0x22, 0xf3, 0x2a, 0xb3, 0x3f, 0x7e,
	0x74, 0x3f, 0x38, 0xeb, 0x73, 0xf0, 0xed, 0x01, 0x04, 0x26, 0x18, 0x05, 0x27, 0xf6, 0x12, 0x3f,
	0xc3, 0x7a, 0xfd, 0x09, 0x16, 0xa9, 0x09, 0xf6, 0xcd, 0x28, 0xab, 0x36, 0xf6, 0xce, 0x38, 0xf6,
	0xe3, 0x74, 0xbb, 0x7f, 0xa7, 0x7b, 0x37, 0xc5, 0xbd, 0xef, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x62, 0xbe, 0x6e, 0x8a, 0xdd, 0x03, 0x00, 0x00,
}