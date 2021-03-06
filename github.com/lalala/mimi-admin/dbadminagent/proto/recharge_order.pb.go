// Code generated by protoc-gen-go. DO NOT EDIT.
// source: recharge_order.proto

package dbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RechargeOrder_RechargeStatus int32

const (
	RechargeOrder_WIAT    RechargeOrder_RechargeStatus = 0
	RechargeOrder_SUCCESS RechargeOrder_RechargeStatus = 1
	RechargeOrder_CLOSE   RechargeOrder_RechargeStatus = 2
)

var RechargeOrder_RechargeStatus_name = map[int32]string{
	0: "WIAT",
	1: "SUCCESS",
	2: "CLOSE",
}
var RechargeOrder_RechargeStatus_value = map[string]int32{
	"WIAT":    0,
	"SUCCESS": 1,
	"CLOSE":   2,
}

func (x RechargeOrder_RechargeStatus) String() string {
	return proto.EnumName(RechargeOrder_RechargeStatus_name, int32(x))
}
func (RechargeOrder_RechargeStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor23, []int{0, 0}
}

type RechargeOrder_RechargePaymentMethod int32

const (
	RechargeOrder_WECHAT RechargeOrder_RechargePaymentMethod = 0
	RechargeOrder_ALIPAY RechargeOrder_RechargePaymentMethod = 1
)

var RechargeOrder_RechargePaymentMethod_name = map[int32]string{
	0: "WECHAT",
	1: "ALIPAY",
}
var RechargeOrder_RechargePaymentMethod_value = map[string]int32{
	"WECHAT": 0,
	"ALIPAY": 1,
}

func (x RechargeOrder_RechargePaymentMethod) String() string {
	return proto.EnumName(RechargeOrder_RechargePaymentMethod_name, int32(x))
}
func (RechargeOrder_RechargePaymentMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor23, []int{0, 1}
}

type RechargeOrder struct {
	Id            int64                               `protobuf:"varint,1,opt,name=id" json:"id"`
	OrderNo       string                              `protobuf:"bytes,2,opt,name=order_no,json=orderNo" json:"order_no"`
	AccountId     int64                               `protobuf:"varint,3,opt,name=account_id,json=accountId" json:"account_id"`
	Money         int32                               `protobuf:"varint,4,opt,name=money" json:"money"`
	Status        RechargeOrder_RechargeStatus        `protobuf:"varint,5,opt,name=status,enum=dbproto.RechargeOrder_RechargeStatus" json:"status"`
	Os            string                              `protobuf:"bytes,6,opt,name=os" json:"os"`
	ClientReqTime int32                               `protobuf:"varint,7,opt,name=client_req_time,json=clientReqTime" json:"client_req_time"`
	CreateTime    int32                               `protobuf:"varint,8,opt,name=create_time,json=createTime" json:"create_time"`
	UcRespTime    int32                               `protobuf:"varint,9,opt,name=uc_resp_time,json=ucRespTime" json:"uc_resp_time"`
	PaymentMethod RechargeOrder_RechargePaymentMethod `protobuf:"varint,10,opt,name=payment_method,json=paymentMethod,enum=dbproto.RechargeOrder_RechargePaymentMethod" json:"payment_method"`
}

func (m *RechargeOrder) Reset()                    { *m = RechargeOrder{} }
func (m *RechargeOrder) String() string            { return proto.CompactTextString(m) }
func (*RechargeOrder) ProtoMessage()               {}
func (*RechargeOrder) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{0} }

func (m *RechargeOrder) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RechargeOrder) GetOrderNo() string {
	if m != nil {
		return m.OrderNo
	}
	return ""
}

func (m *RechargeOrder) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *RechargeOrder) GetMoney() int32 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *RechargeOrder) GetStatus() RechargeOrder_RechargeStatus {
	if m != nil {
		return m.Status
	}
	return RechargeOrder_WIAT
}

func (m *RechargeOrder) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *RechargeOrder) GetClientReqTime() int32 {
	if m != nil {
		return m.ClientReqTime
	}
	return 0
}

func (m *RechargeOrder) GetCreateTime() int32 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *RechargeOrder) GetUcRespTime() int32 {
	if m != nil {
		return m.UcRespTime
	}
	return 0
}

func (m *RechargeOrder) GetPaymentMethod() RechargeOrder_RechargePaymentMethod {
	if m != nil {
		return m.PaymentMethod
	}
	return RechargeOrder_WECHAT
}

type RechargeOrderList struct {
	Total     int32                 `protobuf:"varint,1,opt,name=total" json:"total"`
	Page      int32                 `protobuf:"varint,2,opt,name=page" json:"page"`
	Size      int32                 `protobuf:"varint,3,opt,name=size" json:"size"`
	StartTime int32                 `protobuf:"varint,4,opt,name=start_time,json=startTime" json:"start_time"`
	EndTime   int32                 `protobuf:"varint,5,opt,name=end_time,json=endTime" json:"end_time"`
	List      []*RechargeStatistics `protobuf:"bytes,6,rep,name=list" json:"list"`
}

func (m *RechargeOrderList) Reset()                    { *m = RechargeOrderList{} }
func (m *RechargeOrderList) String() string            { return proto.CompactTextString(m) }
func (*RechargeOrderList) ProtoMessage()               {}
func (*RechargeOrderList) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{1} }

func (m *RechargeOrderList) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RechargeOrderList) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *RechargeOrderList) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *RechargeOrderList) GetStartTime() int32 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *RechargeOrderList) GetEndTime() int32 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *RechargeOrderList) GetList() []*RechargeStatistics {
	if m != nil {
		return m.List
	}
	return nil
}

type RechargeStatistics struct {
	RechargeTime          int32  `protobuf:"varint,1,opt,name=recharge_time,json=rechargeTime" json:"recharge_time"`
	RechargeNum           int32  `protobuf:"varint,2,opt,name=recharge_num,json=rechargeNum" json:"recharge_num"`
	RechargeTotalAmount   string `protobuf:"bytes,3,opt,name=recharge_total_amount,json=rechargeTotalAmount" json:"recharge_total_amount"`
	WechatRechargeAmount  string `protobuf:"bytes,4,opt,name=wechat_recharge_amount,json=wechatRechargeAmount" json:"wechat_recharge_amount"`
	AlipayRechargeAmount  string `protobuf:"bytes,5,opt,name=alipay_recharge_amount,json=alipayRechargeAmount" json:"alipay_recharge_amount"`
	WechatRechargeNum     int32  `protobuf:"varint,6,opt,name=wechat_recharge_num,json=wechatRechargeNum" json:"wechat_recharge_num"`
	AlipayRechargeNum     int32  `protobuf:"varint,7,opt,name=alipay_recharge_num,json=alipayRechargeNum" json:"alipay_recharge_num"`
	AndroidRechargeAmount string `protobuf:"bytes,8,opt,name=android_recharge_amount,json=androidRechargeAmount" json:"android_recharge_amount"`
	IphoneRechargeAmount  string `protobuf:"bytes,9,opt,name=iphone_recharge_amount,json=iphoneRechargeAmount" json:"iphone_recharge_amount"`
	IphoneRechargeNum     int32  `protobuf:"varint,10,opt,name=iphone_recharge_num,json=iphoneRechargeNum" json:"iphone_recharge_num"`
	AndroidRechargeNum    int32  `protobuf:"varint,11,opt,name=android_recharge_num,json=androidRechargeNum" json:"android_recharge_num"`
}

func (m *RechargeStatistics) Reset()                    { *m = RechargeStatistics{} }
func (m *RechargeStatistics) String() string            { return proto.CompactTextString(m) }
func (*RechargeStatistics) ProtoMessage()               {}
func (*RechargeStatistics) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{2} }

func (m *RechargeStatistics) GetRechargeTime() int32 {
	if m != nil {
		return m.RechargeTime
	}
	return 0
}

func (m *RechargeStatistics) GetRechargeNum() int32 {
	if m != nil {
		return m.RechargeNum
	}
	return 0
}

func (m *RechargeStatistics) GetRechargeTotalAmount() string {
	if m != nil {
		return m.RechargeTotalAmount
	}
	return ""
}

func (m *RechargeStatistics) GetWechatRechargeAmount() string {
	if m != nil {
		return m.WechatRechargeAmount
	}
	return ""
}

func (m *RechargeStatistics) GetAlipayRechargeAmount() string {
	if m != nil {
		return m.AlipayRechargeAmount
	}
	return ""
}

func (m *RechargeStatistics) GetWechatRechargeNum() int32 {
	if m != nil {
		return m.WechatRechargeNum
	}
	return 0
}

func (m *RechargeStatistics) GetAlipayRechargeNum() int32 {
	if m != nil {
		return m.AlipayRechargeNum
	}
	return 0
}

func (m *RechargeStatistics) GetAndroidRechargeAmount() string {
	if m != nil {
		return m.AndroidRechargeAmount
	}
	return ""
}

func (m *RechargeStatistics) GetIphoneRechargeAmount() string {
	if m != nil {
		return m.IphoneRechargeAmount
	}
	return ""
}

func (m *RechargeStatistics) GetIphoneRechargeNum() int32 {
	if m != nil {
		return m.IphoneRechargeNum
	}
	return 0
}

func (m *RechargeStatistics) GetAndroidRechargeNum() int32 {
	if m != nil {
		return m.AndroidRechargeNum
	}
	return 0
}

func init() {
	proto.RegisterType((*RechargeOrder)(nil), "dbproto.RechargeOrder")
	proto.RegisterType((*RechargeOrderList)(nil), "dbproto.RechargeOrderList")
	proto.RegisterType((*RechargeStatistics)(nil), "dbproto.RechargeStatistics")
	proto.RegisterEnum("dbproto.RechargeOrder_RechargeStatus", RechargeOrder_RechargeStatus_name, RechargeOrder_RechargeStatus_value)
	proto.RegisterEnum("dbproto.RechargeOrder_RechargePaymentMethod", RechargeOrder_RechargePaymentMethod_name, RechargeOrder_RechargePaymentMethod_value)
}

func init() { proto.RegisterFile("recharge_order.proto", fileDescriptor23) }

var fileDescriptor23 = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x25, 0x6d, 0x93, 0x36, 0xb7, 0x6b, 0xe9, 0xbc, 0x0e, 0x32, 0x21, 0x44, 0x29, 0x02, 0xf5,
	0x01, 0x75, 0xa8, 0x20, 0xde, 0x78, 0xa8, 0xaa, 0x49, 0x4c, 0x1a, 0xdb, 0xe4, 0x0c, 0x4d, 0x3c,
	0x45, 0x5e, 0x62, 0x6d, 0x96, 0x9a, 0x38, 0x8b, 0x1d, 0xa1, 0xf1, 0x4d, 0x7c, 0x02, 0x7f, 0xc3,
	0x8f, 0x20, 0x5f, 0x67, 0x1d, 0x69, 0x90, 0x78, 0xcb, 0xbd, 0xe7, 0x1c, 0xdf, 0x73, 0x7c, 0x1d,
	0x18, 0x17, 0x3c, 0xbe, 0x61, 0xc5, 0x35, 0x8f, 0x64, 0x91, 0xf0, 0x62, 0x9e, 0x17, 0x52, 0x4b,
	0xd2, 0x4d, 0xae, 0xf0, 0x63, 0xfa, 0xbb, 0x0d, 0x03, 0x5a, 0x31, 0xce, 0x0c, 0x81, 0x0c, 0xa1,
	0x25, 0x92, 0xc0, 0x99, 0x38, 0xb3, 0x36, 0x6d, 0x89, 0x84, 0x1c, 0x40, 0x0f, 0x95, 0x51, 0x26,
	0x83, 0xd6, 0xc4, 0x99, 0xf9, 0xb4, 0x8b, 0xf5, 0xa9, 0x24, 0xcf, 0x01, 0x58, 0x1c, 0xcb, 0x32,
	0xd3, 0x91, 0x48, 0x82, 0x36, 0x4a, 0xfc, 0xaa, 0x73, 0x9c, 0x90, 0x31, 0xb8, 0xa9, 0xcc, 0xf8,
	0x5d, 0xd0, 0x99, 0x38, 0x33, 0x97, 0xda, 0x82, 0x7c, 0x02, 0x4f, 0x69, 0xa6, 0x4b, 0x15, 0xb8,
	0x13, 0x67, 0x36, 0x5c, 0xbc, 0x9e, 0x57, 0x5e, 0xe6, 0x35, 0x1f, 0x9b, 0x2a, 0x44, 0x32, 0xad,
	0x44, 0xc6, 0x9e, 0x54, 0x81, 0x87, 0x46, 0x5a, 0x52, 0x91, 0x37, 0xf0, 0x38, 0x5e, 0x0b, 0x9e,
	0xe9, 0xa8, 0xe0, 0xb7, 0x91, 0x16, 0x29, 0x0f, 0xba, 0x38, 0x6e, 0x60, 0xdb, 0x94, 0xdf, 0x5e,
	0x88, 0x94, 0x93, 0x17, 0xd0, 0x8f, 0x0b, 0xce, 0x34, 0xb7, 0x9c, 0x1e, 0x72, 0xc0, 0xb6, 0x90,
	0x30, 0x81, 0x9d, 0x32, 0x8e, 0x0a, 0xae, 0x72, 0xcb, 0xf0, 0x2d, 0xa3, 0x8c, 0x29, 0x57, 0x39,
	0x32, 0x42, 0x18, 0xe6, 0xec, 0x2e, 0x35, 0xb3, 0x52, 0xae, 0x6f, 0x64, 0x12, 0x00, 0x26, 0x78,
	0xfb, 0x9f, 0x04, 0xe7, 0x56, 0xf4, 0x05, 0x35, 0x74, 0x90, 0xff, 0x5d, 0x4e, 0x17, 0x30, 0xac,
	0x27, 0x25, 0x3d, 0xe8, 0x5c, 0x1e, 0x2f, 0x2f, 0x46, 0x8f, 0x48, 0x1f, 0xba, 0xe1, 0xd7, 0xd5,
	0xea, 0x28, 0x0c, 0x47, 0x0e, 0xf1, 0xc1, 0x5d, 0x9d, 0x9c, 0x85, 0x47, 0xa3, 0xd6, 0xf4, 0x10,
	0xf6, 0xff, 0x79, 0x36, 0x01, 0xf0, 0x2e, 0x8f, 0x56, 0x9f, 0x51, 0x0c, 0xe0, 0x2d, 0x4f, 0x8e,
	0xcf, 0x97, 0xdf, 0x46, 0xce, 0xf4, 0x97, 0x03, 0xbb, 0x35, 0x6f, 0x27, 0x42, 0x69, 0xb3, 0x1f,
	0x2d, 0x35, 0x5b, 0xe3, 0xb2, 0x5d, 0x6a, 0x0b, 0x42, 0xa0, 0x93, 0xb3, 0x6b, 0x8e, 0xbb, 0x76,
	0x29, 0x7e, 0x9b, 0x9e, 0x12, 0x3f, 0x38, 0xae, 0xd8, 0xa5, 0xf8, 0x6d, 0x96, 0xaf, 0x34, 0x2b,
	0xb4, 0xbd, 0x2d, 0xbb, 0x62, 0x1f, 0x3b, 0x78, 0x59, 0x07, 0xd0, 0xe3, 0x59, 0x62, 0x41, 0x17,
	0xc1, 0x2e, 0xcf, 0x12, 0x84, 0x0e, 0xa1, 0xb3, 0x16, 0x4a, 0x07, 0xde, 0xa4, 0x3d, 0xeb, 0x2f,
	0x9e, 0x35, 0x6e, 0xcf, 0xdc, 0x83, 0x50, 0x5a, 0xc4, 0x8a, 0x22, 0x71, 0xfa, 0xb3, 0x03, 0xa4,
	0x09, 0x92, 0x57, 0x30, 0xd8, 0x3c, 0x6e, 0x9c, 0x63, 0x73, 0xec, 0xdc, 0x37, 0x71, 0xd8, 0x4b,
	0xd8, 0xd4, 0x51, 0x56, 0xa6, 0x55, 0xac, 0xfe, 0x7d, 0xef, 0xb4, 0x4c, 0xc9, 0x02, 0xf6, 0x1f,
	0xce, 0x31, 0x77, 0x10, 0xb1, 0xd4, 0x3c, 0x61, 0x8c, 0xeb, 0xd3, 0xbd, 0xcd, 0x79, 0x06, 0x5b,
	0x22, 0x44, 0x3e, 0xc0, 0x93, 0xef, 0xa6, 0x6d, 0x9e, 0x5d, 0x25, 0xad, 0x44, 0x1d, 0x14, 0x8d,
	0x2d, 0x7a, 0xef, 0xfa, 0x41, 0xc5, 0xd6, 0x22, 0x67, 0x77, 0x0d, 0x95, 0x6b, 0x55, 0x16, 0xdd,
	0x52, 0xcd, 0x61, 0x6f, 0x7b, 0x96, 0x49, 0xe2, 0x61, 0x92, 0xdd, 0xfa, 0x20, 0x93, 0x67, 0x0e,
	0x7b, 0xdb, 0x53, 0x0c, 0xdf, 0xfe, 0x16, 0xbb, 0xf5, 0x11, 0x86, 0xff, 0x11, 0x9e, 0xb2, 0x2c,
	0x29, 0xa4, 0x48, 0x1a, 0xb6, 0x7a, 0x68, 0x6b, 0xbf, 0x82, 0x9b, 0x69, 0x44, 0x7e, 0x23, 0x33,
	0xde, 0x90, 0xf9, 0x36, 0x8d, 0x45, 0x9b, 0x69, 0xb6, 0x55, 0xc6, 0x1d, 0x58, 0x77, 0x75, 0x89,
	0x71, 0xf7, 0x0e, 0xc6, 0x0d, 0x77, 0x46, 0xd0, 0x47, 0x01, 0xd9, 0xb2, 0x76, 0x5a, 0xa6, 0x57,
	0x1e, 0x3e, 0xa7, 0xf7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x38, 0x0d, 0x4f, 0xbb, 0xfb, 0x04,
	0x00, 0x00,
}
