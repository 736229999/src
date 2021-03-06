// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lottery_options.proto

package dbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LotteryOptions struct {
	Id          int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	LotteryName string `protobuf:"bytes,2,opt,name=lottery_name,json=lotteryName" json:"lottery_name"`
	IsPlusAward bool   `protobuf:"varint,3,opt,name=is_plus_award,json=isPlusAward" json:"is_plus_award"`
	Info        string `protobuf:"bytes,4,opt,name=info" json:"info"`
	StopSale    bool   `protobuf:"varint,5,opt,name=stop_sale,json=stopSale" json:"stop_sale"`
	CreateTime  int64  `protobuf:"varint,6,opt,name=create_time,json=createTime" json:"create_time"`
	UpdateTime  int64  `protobuf:"varint,7,opt,name=update_time,json=updateTime" json:"update_time"`
}

func (m *LotteryOptions) Reset()                    { *m = LotteryOptions{} }
func (m *LotteryOptions) String() string            { return proto.CompactTextString(m) }
func (*LotteryOptions) ProtoMessage()               {}
func (*LotteryOptions) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

func (m *LotteryOptions) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LotteryOptions) GetLotteryName() string {
	if m != nil {
		return m.LotteryName
	}
	return ""
}

func (m *LotteryOptions) GetIsPlusAward() bool {
	if m != nil {
		return m.IsPlusAward
	}
	return false
}

func (m *LotteryOptions) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *LotteryOptions) GetStopSale() bool {
	if m != nil {
		return m.StopSale
	}
	return false
}

func (m *LotteryOptions) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *LotteryOptions) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type LotteryOptionsList struct {
	List []*LotteryOptions `protobuf:"bytes,2,rep,name=list" json:"list"`
}

func (m *LotteryOptionsList) Reset()                    { *m = LotteryOptionsList{} }
func (m *LotteryOptionsList) String() string            { return proto.CompactTextString(m) }
func (*LotteryOptionsList) ProtoMessage()               {}
func (*LotteryOptionsList) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{1} }

func (m *LotteryOptionsList) GetList() []*LotteryOptions {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*LotteryOptions)(nil), "dbproto.LotteryOptions")
	proto.RegisterType((*LotteryOptionsList)(nil), "dbproto.LotteryOptionsList")
}

func init() { proto.RegisterFile("lottery_options.proto", fileDescriptor18) }

var fileDescriptor18 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x86, 0xe9, 0x87, 0xfb, 0x31, 0xd5, 0x3d, 0x0c, 0x88, 0x01, 0x0f, 0xd6, 0x9e, 0x0a, 0x42,
	0x0f, 0xfa, 0x0b, 0xf6, 0xbe, 0xa8, 0x54, 0xef, 0x21, 0x6b, 0x22, 0x0c, 0xa4, 0x4d, 0x68, 0xa6,
	0x88, 0x7f, 0xd5, 0x5f, 0x23, 0x4d, 0x56, 0x61, 0x6f, 0x2f, 0x4f, 0x9e, 0x37, 0xcc, 0x0c, 0x5c,
	0x5b, 0xc7, 0x6c, 0xa6, 0x6f, 0xe9, 0x3c, 0x93, 0x1b, 0x43, 0xe7, 0x27, 0xc7, 0x0e, 0xd7, 0xfa,
	0x18, 0x43, 0xf3, 0x93, 0xc1, 0xee, 0x90, 0x94, 0x97, 0x64, 0xe0, 0x0e, 0x72, 0xd2, 0x22, 0xab,
	0xb3, 0xb6, 0xe8, 0x73, 0xd2, 0x78, 0x0f, 0x97, 0x7f, 0x9f, 0x8c, 0x6a, 0x30, 0x22, 0xaf, 0xb3,
	0x76, 0xdb, 0x57, 0x27, 0xf6, 0xac, 0x06, 0x83, 0x0d, 0x5c, 0x51, 0x90, 0xde, 0xce, 0x41, 0xaa,
	0x2f, 0x35, 0x69, 0x51, 0xd4, 0x59, 0xbb, 0xe9, 0x2b, 0x0a, 0xaf, 0x76, 0x0e, 0xfb, 0x05, 0x21,
	0x42, 0x49, 0xe3, 0xa7, 0x13, 0x65, 0xac, 0xc7, 0x8c, 0xb7, 0xb0, 0x0d, 0xec, 0xbc, 0x0c, 0xca,
	0x1a, 0x71, 0x11, 0x3b, 0x9b, 0x05, 0xbc, 0x29, 0x6b, 0xf0, 0x0e, 0xaa, 0x8f, 0xc9, 0x28, 0x36,
	0x92, 0x69, 0x30, 0x62, 0x15, 0x07, 0x82, 0x84, 0xde, 0x69, 0x88, 0xc2, 0xec, 0xf5, 0xbf, 0xb0,
	0x4e, 0x42, 0x42, 0x8b, 0xd0, 0xec, 0x01, 0xcf, 0x77, 0x3b, 0x50, 0x60, 0x7c, 0x80, 0xd2, 0x52,
	0x60, 0x91, 0xd7, 0x45, 0x5b, 0x3d, 0xde, 0x74, 0xa7, 0x53, 0x74, 0xe7, 0x6a, 0x1f, 0xa5, 0xe3,
	0x2a, 0xbe, 0x3d, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x64, 0xc0, 0x62, 0x4d, 0x48, 0x01, 0x00,
	0x00,
}
