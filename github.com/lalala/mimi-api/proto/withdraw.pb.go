// Code generated by protoc-gen-go. DO NOT EDIT.
// source: withdraw.proto

package apiproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WithdrawType int32

const (
	WithdrawType_ToUnknown  WithdrawType = 0
	WithdrawType_ToBankcard WithdrawType = 1
	WithdrawType_ToAlipay   WithdrawType = 2
)

var WithdrawType_name = map[int32]string{
	0: "ToUnknown",
	1: "ToBankcard",
	2: "ToAlipay",
}
var WithdrawType_value = map[string]int32{
	"ToUnknown":  0,
	"ToBankcard": 1,
	"ToAlipay":   2,
}

func (x WithdrawType) String() string {
	return proto.EnumName(WithdrawType_name, int32(x))
}
func (WithdrawType) EnumDescriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

type WithdrawStatus int32

const (
	WithdrawStatus_NotRequestWithdraw WithdrawStatus = 0
	WithdrawStatus_RequestWithdraw    WithdrawStatus = 1
	WithdrawStatus_AuditWithdraw      WithdrawStatus = 2
	WithdrawStatus_BankTransfer       WithdrawStatus = 3
)

var WithdrawStatus_name = map[int32]string{
	0: "NotRequestWithdraw",
	1: "RequestWithdraw",
	2: "AuditWithdraw",
	3: "BankTransfer",
}
var WithdrawStatus_value = map[string]int32{
	"NotRequestWithdraw": 0,
	"RequestWithdraw":    1,
	"AuditWithdraw":      2,
	"BankTransfer":       3,
}

func (x WithdrawStatus) String() string {
	return proto.EnumName(WithdrawStatus_name, int32(x))
}
func (WithdrawStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

// 提现信息
type WithdrawInfo struct {
	TotalWin      float64 `protobuf:"fixed64,1,opt,name=total_win,json=totalWin" json:"total_win"`
	TotalWithdraw float64 `protobuf:"fixed64,2,opt,name=total_withdraw,json=totalWithdraw" json:"total_withdraw"`
	CurWithdraw   float64 `protobuf:"fixed64,3,opt,name=cur_withdraw,json=curWithdraw" json:"cur_withdraw"`
}

func (m *WithdrawInfo) Reset()                    { *m = WithdrawInfo{} }
func (m *WithdrawInfo) String() string            { return proto.CompactTextString(m) }
func (*WithdrawInfo) ProtoMessage()               {}
func (*WithdrawInfo) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *WithdrawInfo) GetTotalWin() float64 {
	if m != nil {
		return m.TotalWin
	}
	return 0
}

func (m *WithdrawInfo) GetTotalWithdraw() float64 {
	if m != nil {
		return m.TotalWithdraw
	}
	return 0
}

func (m *WithdrawInfo) GetCurWithdraw() float64 {
	if m != nil {
		return m.CurWithdraw
	}
	return 0
}

type WithdrawRequest struct {
	Type   WithdrawType `protobuf:"varint,1,opt,name=type,enum=apiproto.WithdrawType" json:"type"`
	Amount float64      `protobuf:"fixed64,2,opt,name=amount" json:"amount"`
}

func (m *WithdrawRequest) Reset()                    { *m = WithdrawRequest{} }
func (m *WithdrawRequest) String() string            { return proto.CompactTextString(m) }
func (*WithdrawRequest) ProtoMessage()               {}
func (*WithdrawRequest) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

func (m *WithdrawRequest) GetType() WithdrawType {
	if m != nil {
		return m.Type
	}
	return WithdrawType_ToUnknown
}

func (m *WithdrawRequest) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type WithdrawProgress struct {
	StepList   []string                   `protobuf:"bytes,1,rep,name=step_list,json=stepList" json:"step_list"`
	CurStep    int32                      `protobuf:"varint,2,opt,name=cur_step,json=curStep" json:"cur_step"`
	IsSuccess  bool                       `protobuf:"varint,3,opt,name=is_success,json=isSuccess" json:"is_success"`
	Desc       string                     `protobuf:"bytes,4,opt,name=desc" json:"desc"`
	DetailList []*WithdrawProgress_Detail `protobuf:"bytes,5,rep,name=detail_list,json=detailList" json:"detail_list"`
}

func (m *WithdrawProgress) Reset()                    { *m = WithdrawProgress{} }
func (m *WithdrawProgress) String() string            { return proto.CompactTextString(m) }
func (*WithdrawProgress) ProtoMessage()               {}
func (*WithdrawProgress) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{2} }

func (m *WithdrawProgress) GetStepList() []string {
	if m != nil {
		return m.StepList
	}
	return nil
}

func (m *WithdrawProgress) GetCurStep() int32 {
	if m != nil {
		return m.CurStep
	}
	return 0
}

func (m *WithdrawProgress) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func (m *WithdrawProgress) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *WithdrawProgress) GetDetailList() []*WithdrawProgress_Detail {
	if m != nil {
		return m.DetailList
	}
	return nil
}

type WithdrawProgress_Detail struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value"`
}

func (m *WithdrawProgress_Detail) Reset()                    { *m = WithdrawProgress_Detail{} }
func (m *WithdrawProgress_Detail) String() string            { return proto.CompactTextString(m) }
func (*WithdrawProgress_Detail) ProtoMessage()               {}
func (*WithdrawProgress_Detail) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{2, 0} }

func (m *WithdrawProgress_Detail) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *WithdrawProgress_Detail) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*WithdrawInfo)(nil), "apiproto.WithdrawInfo")
	proto.RegisterType((*WithdrawRequest)(nil), "apiproto.WithdrawRequest")
	proto.RegisterType((*WithdrawProgress)(nil), "apiproto.WithdrawProgress")
	proto.RegisterType((*WithdrawProgress_Detail)(nil), "apiproto.WithdrawProgress.Detail")
	proto.RegisterEnum("apiproto.WithdrawType", WithdrawType_name, WithdrawType_value)
	proto.RegisterEnum("apiproto.WithdrawStatus", WithdrawStatus_name, WithdrawStatus_value)
}

func init() { proto.RegisterFile("withdraw.proto", fileDescriptor15) }

var fileDescriptor15 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0x2f, 0x4d, 0x5b, 0x93, 0x69, 0x9b, 0xcb, 0x8d, 0x52, 0xaa, 0x22, 0xf4, 0x0a, 0x42,
	0xe9, 0x43, 0x91, 0xf3, 0xd1, 0xa7, 0x2b, 0xbe, 0x08, 0x22, 0xc7, 0x36, 0xc7, 0x3d, 0x96, 0xbd,
	0x64, 0xab, 0x4b, 0xe3, 0x6e, 0xdc, 0x1f, 0x96, 0xfc, 0xdd, 0xfe, 0x03, 0xb2, 0x9b, 0x2c, 0x05,
	0x7d, 0x9b, 0xf9, 0xce, 0x67, 0x76, 0xbe, 0x33, 0x0b, 0xd9, 0x99, 0x9b, 0x1f, 0x95, 0xa2, 0xe7,
	0x6d, 0xa3, 0xa4, 0x91, 0x98, 0xd0, 0x86, 0xfb, 0x68, 0x65, 0x61, 0xfa, 0xd4, 0xd7, 0xbe, 0x88,
	0xa3, 0xc4, 0xb7, 0x90, 0x1a, 0x69, 0x68, 0x7d, 0x38, 0x73, 0xb1, 0x88, 0x96, 0xd1, 0x3a, 0x22,
	0x89, 0x17, 0x9e, 0xb8, 0xc0, 0xf7, 0x90, 0x85, 0x62, 0xd7, 0xb2, 0x18, 0x78, 0x62, 0xd6, 0x13,
	0x9d, 0x88, 0xb7, 0x30, 0x2d, 0xad, 0xba, 0x40, 0xb1, 0x87, 0x26, 0xa5, 0x55, 0x01, 0x59, 0x3d,
	0xc2, 0x75, 0x88, 0x09, 0xfb, 0x65, 0x99, 0x36, 0xb8, 0x81, 0xa1, 0x69, 0x1b, 0xe6, 0x87, 0x66,
	0x77, 0xf3, 0x6d, 0xb0, 0xb8, 0x0d, 0x60, 0xd1, 0x36, 0x8c, 0x78, 0x06, 0xe7, 0x30, 0xa6, 0x3f,
	0xa5, 0x15, 0xa6, 0x37, 0xd0, 0x67, 0xab, 0x3f, 0x11, 0xe4, 0x01, 0x7f, 0x50, 0xf2, 0xbb, 0x62,
	0x5a, 0xbb, 0x95, 0xb4, 0x61, 0xcd, 0xa1, 0xe6, 0xda, 0x2c, 0xa2, 0x65, 0xbc, 0x4e, 0x49, 0xe2,
	0x84, 0xaf, 0x5c, 0x1b, 0x7c, 0x0d, 0x89, 0xf3, 0xea, 0x72, 0xff, 0xd6, 0x88, 0xbc, 0x28, 0xad,
	0xda, 0x1b, 0xd6, 0xe0, 0x3b, 0x00, 0xae, 0x0f, 0xda, 0x96, 0x25, 0xd3, 0xda, 0x2f, 0x91, 0x90,
	0x94, 0xeb, 0x7d, 0x27, 0x20, 0xc2, 0xb0, 0x62, 0xba, 0x5c, 0x0c, 0x97, 0xd1, 0x3a, 0x25, 0x3e,
	0xc6, 0x1d, 0x4c, 0x2a, 0x66, 0x28, 0xaf, 0xbb, 0x61, 0xa3, 0x65, 0xbc, 0x9e, 0xdc, 0xdd, 0xfe,
	0xbf, 0x4a, 0xf0, 0xb6, 0xfd, 0xec, 0x69, 0x02, 0x5d, 0x97, 0x73, 0xf4, 0xe6, 0x03, 0x8c, 0x3b,
	0x15, 0x73, 0x88, 0x4f, 0xac, 0xf5, 0x07, 0x49, 0x89, 0x0b, 0xf1, 0x15, 0x8c, 0x7e, 0xd3, 0xda,
	0x32, 0x6f, 0x35, 0x25, 0x5d, 0xb2, 0xf9, 0x74, 0xf9, 0x43, 0x77, 0x23, 0x9c, 0x41, 0x5a, 0xc8,
	0x47, 0x71, 0x12, 0xf2, 0x2c, 0xf2, 0x2b, 0xcc, 0x00, 0x0a, 0xb9, 0xa3, 0xe2, 0x54, 0x52, 0x55,
	0xe5, 0x11, 0x4e, 0x21, 0x29, 0xe4, 0x7d, 0xcd, 0x1b, 0xda, 0xe6, 0x83, 0xcd, 0x33, 0x64, 0xa1,
	0x79, 0x6f, 0xa8, 0xb1, 0x1a, 0xe7, 0x80, 0xdf, 0xa4, 0xe9, 0xbf, 0x25, 0xd4, 0xf2, 0x2b, 0x7c,
	0x09, 0xd7, 0xff, 0x8a, 0x11, 0xde, 0xc0, 0xec, 0xde, 0x56, 0xfc, 0x22, 0x0d, 0x30, 0x87, 0xa9,
	0x9b, 0x56, 0x28, 0x2a, 0xf4, 0x91, 0xa9, 0x3c, 0xde, 0xdd, 0xc0, 0xe8, 0xc1, 0xed, 0xbf, 0x4b,
	0x02, 0xf6, 0x3c, 0xf6, 0x07, 0xf9, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x0d, 0x20, 0x8b,
	0x9a, 0x02, 0x00, 0x00,
}
