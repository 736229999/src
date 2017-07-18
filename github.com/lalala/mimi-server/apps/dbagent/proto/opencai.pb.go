// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opencai.proto

package dbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BonusDetail struct {
	Id    int32   `protobuf:"varint,1,opt,name=id" json:"id"`
	Num   int32   `protobuf:"varint,2,opt,name=num" json:"num"`
	Money float64 `protobuf:"fixed64,3,opt,name=money" json:"money"`
}

func (m *BonusDetail) Reset()                    { *m = BonusDetail{} }
func (m *BonusDetail) String() string            { return proto.CompactTextString(m) }
func (*BonusDetail) ProtoMessage()               {}
func (*BonusDetail) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

func (m *BonusDetail) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BonusDetail) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *BonusDetail) GetMoney() float64 {
	if m != nil {
		return m.Money
	}
	return 0
}

type OpenDetail struct {
	Sale      float64        `protobuf:"fixed64,1,opt,name=sale" json:"sale"`
	Pool      float64        `protobuf:"fixed64,2,opt,name=pool" json:"pool"`
	BonusList []*BonusDetail `protobuf:"bytes,3,rep,name=bonus_list,json=bonusList" json:"bonus_list"`
}

func (m *OpenDetail) Reset()                    { *m = OpenDetail{} }
func (m *OpenDetail) String() string            { return proto.CompactTextString(m) }
func (*OpenDetail) ProtoMessage()               {}
func (*OpenDetail) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{1} }

func (m *OpenDetail) GetSale() float64 {
	if m != nil {
		return m.Sale
	}
	return 0
}

func (m *OpenDetail) GetPool() float64 {
	if m != nil {
		return m.Pool
	}
	return 0
}

func (m *OpenDetail) GetBonusList() []*BonusDetail {
	if m != nil {
		return m.BonusList
	}
	return nil
}

type OpenInfo struct {
	Issue      string      `protobuf:"bytes,1,opt,name=issue" json:"issue"`
	OpenTime   int64       `protobuf:"varint,2,opt,name=open_time,json=openTime" json:"open_time"`
	Balls      string      `protobuf:"bytes,3,opt,name=balls" json:"balls"`
	GrabTime   int64       `protobuf:"varint,4,opt,name=grab_time,json=grabTime" json:"grab_time"`
	GrabSource string      `protobuf:"bytes,5,opt,name=grab_source,json=grabSource" json:"grab_source"`
	Detail     *OpenDetail `protobuf:"bytes,6,opt,name=detail" json:"detail"`
}

func (m *OpenInfo) Reset()                    { *m = OpenInfo{} }
func (m *OpenInfo) String() string            { return proto.CompactTextString(m) }
func (*OpenInfo) ProtoMessage()               {}
func (*OpenInfo) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{2} }

func (m *OpenInfo) GetIssue() string {
	if m != nil {
		return m.Issue
	}
	return ""
}

func (m *OpenInfo) GetOpenTime() int64 {
	if m != nil {
		return m.OpenTime
	}
	return 0
}

func (m *OpenInfo) GetBalls() string {
	if m != nil {
		return m.Balls
	}
	return ""
}

func (m *OpenInfo) GetGrabTime() int64 {
	if m != nil {
		return m.GrabTime
	}
	return 0
}

func (m *OpenInfo) GetGrabSource() string {
	if m != nil {
		return m.GrabSource
	}
	return ""
}

func (m *OpenInfo) GetDetail() *OpenDetail {
	if m != nil {
		return m.Detail
	}
	return nil
}

type OpencaiQueryArg struct {
	Code string   `protobuf:"bytes,1,opt,name=code" json:"code"`
	Args []string `protobuf:"bytes,2,rep,name=args" json:"args"`
}

func (m *OpencaiQueryArg) Reset()                    { *m = OpencaiQueryArg{} }
func (m *OpencaiQueryArg) String() string            { return proto.CompactTextString(m) }
func (*OpencaiQueryArg) ProtoMessage()               {}
func (*OpencaiQueryArg) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{3} }

func (m *OpencaiQueryArg) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *OpencaiQueryArg) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type OpencaiInsertArg struct {
	Code string    `protobuf:"bytes,1,opt,name=code" json:"code"`
	Info *OpenInfo `protobuf:"bytes,2,opt,name=info" json:"info"`
}

func (m *OpencaiInsertArg) Reset()                    { *m = OpencaiInsertArg{} }
func (m *OpencaiInsertArg) String() string            { return proto.CompactTextString(m) }
func (*OpencaiInsertArg) ProtoMessage()               {}
func (*OpencaiInsertArg) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{4} }

func (m *OpencaiInsertArg) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *OpencaiInsertArg) GetInfo() *OpenInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type OpencaiUpsertDetailArg struct {
	Code   string      `protobuf:"bytes,1,opt,name=code" json:"code"`
	Issue  string      `protobuf:"bytes,2,opt,name=issue" json:"issue"`
	Detail *OpenDetail `protobuf:"bytes,3,opt,name=detail" json:"detail"`
}

func (m *OpencaiUpsertDetailArg) Reset()                    { *m = OpencaiUpsertDetailArg{} }
func (m *OpencaiUpsertDetailArg) String() string            { return proto.CompactTextString(m) }
func (*OpencaiUpsertDetailArg) ProtoMessage()               {}
func (*OpencaiUpsertDetailArg) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{5} }

func (m *OpencaiUpsertDetailArg) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *OpencaiUpsertDetailArg) GetIssue() string {
	if m != nil {
		return m.Issue
	}
	return ""
}

func (m *OpencaiUpsertDetailArg) GetDetail() *OpenDetail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func init() {
	proto.RegisterType((*BonusDetail)(nil), "dbproto.BonusDetail")
	proto.RegisterType((*OpenDetail)(nil), "dbproto.OpenDetail")
	proto.RegisterType((*OpenInfo)(nil), "dbproto.OpenInfo")
	proto.RegisterType((*OpencaiQueryArg)(nil), "dbproto.OpencaiQueryArg")
	proto.RegisterType((*OpencaiInsertArg)(nil), "dbproto.OpencaiInsertArg")
	proto.RegisterType((*OpencaiUpsertDetailArg)(nil), "dbproto.OpencaiUpsertDetailArg")
}

func init() { proto.RegisterFile("opencai.proto", fileDescriptor18) }

var fileDescriptor18 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xc1, 0x6a, 0xe3, 0x30,
	0x10, 0x45, 0x76, 0x92, 0x8d, 0xc7, 0xec, 0x6e, 0x56, 0x1b, 0x8a, 0xa1, 0x87, 0x1a, 0x43, 0xc1,
	0x50, 0xc8, 0x21, 0x39, 0xf5, 0xd8, 0xd2, 0x1e, 0x02, 0x2d, 0xa1, 0x6a, 0x7b, 0x0e, 0x76, 0xac,
	0x04, 0x81, 0x2d, 0x19, 0xcb, 0x3e, 0xe4, 0xe7, 0xfa, 0x6d, 0x65, 0x46, 0x6e, 0x92, 0x1e, 0x42,
	0x6f, 0x4f, 0x6f, 0xfc, 0xde, 0xbc, 0x37, 0x86, 0xdf, 0xa6, 0x96, 0x7a, 0x93, 0xa9, 0x59, 0xdd,
	0x98, 0xd6, 0xf0, 0x5f, 0x45, 0x4e, 0x20, 0x79, 0x84, 0xf0, 0xde, 0xe8, 0xce, 0x3e, 0xc8, 0x36,
	0x53, 0x25, 0xff, 0x03, 0x9e, 0x2a, 0x22, 0x16, 0xb3, 0x74, 0x28, 0x3c, 0x55, 0xf0, 0x09, 0xf8,
	0xba, 0xab, 0x22, 0x8f, 0x08, 0x84, 0x7c, 0x0a, 0xc3, 0xca, 0x68, 0xb9, 0x8f, 0xfc, 0x98, 0xa5,
	0x4c, 0xb8, 0x47, 0xa2, 0x00, 0x56, 0xb5, 0xd4, 0xbd, 0x0b, 0x87, 0x81, 0xcd, 0x4a, 0x49, 0x3e,
	0x4c, 0x10, 0x46, 0xae, 0x36, 0xa6, 0x24, 0x2b, 0x26, 0x08, 0xf3, 0x05, 0x40, 0x8e, 0xcb, 0xd7,
	0xa5, 0xb2, 0x6d, 0xe4, 0xc7, 0x7e, 0x1a, 0xce, 0xa7, 0xb3, 0x3e, 0xda, 0xec, 0x24, 0x97, 0x08,
	0xe8, 0xbb, 0x27, 0x65, 0xdb, 0xe4, 0x83, 0xc1, 0x18, 0x77, 0x2d, 0xf5, 0xd6, 0x60, 0x1a, 0x65,
	0x6d, 0xe7, 0x56, 0x05, 0xc2, 0x3d, 0xf8, 0x25, 0x04, 0x58, 0x77, 0xdd, 0xaa, 0x4a, 0xd2, 0x42,
	0x5f, 0x8c, 0x91, 0x78, 0x53, 0x95, 0x44, 0x49, 0x9e, 0x95, 0xa5, 0xa5, 0x02, 0x81, 0x70, 0x0f,
	0x94, 0xec, 0x9a, 0x2c, 0x77, 0x92, 0x81, 0x93, 0x20, 0x41, 0x92, 0x2b, 0x08, 0x69, 0x68, 0x4d,
	0xd7, 0x6c, 0x64, 0x34, 0x24, 0x21, 0x20, 0xf5, 0x4a, 0x0c, 0xbf, 0x81, 0x51, 0x41, 0x41, 0xa3,
	0x51, 0xcc, 0xd2, 0x70, 0xfe, 0xff, 0x50, 0xe2, 0x78, 0x15, 0xd1, 0x7f, 0x92, 0xdc, 0xc2, 0xdf,
	0x95, 0xfb, 0x19, 0x2f, 0x9d, 0x6c, 0xf6, 0x77, 0xcd, 0x0e, 0x8f, 0xb3, 0x31, 0xc5, 0x57, 0x0b,
	0xc2, 0xc8, 0x65, 0xcd, 0xce, 0x46, 0x5e, 0xec, 0x23, 0x87, 0x38, 0x79, 0x86, 0x49, 0x2f, 0x5d,
	0x6a, 0x2b, 0x9b, 0xf6, 0x9c, 0xf6, 0x1a, 0x06, 0x4a, 0x6f, 0x0d, 0x75, 0x0f, 0xe7, 0xff, 0xbe,
	0xa5, 0xc1, 0xbb, 0x09, 0x1a, 0x27, 0x06, 0x2e, 0x7a, 0xbb, 0xf7, 0x1a, 0xed, 0x5c, 0xd0, 0x73,
	0xa6, 0x87, 0x5b, 0x7b, 0xa7, 0xb7, 0x3e, 0x56, 0xf7, 0x7f, 0xac, 0x9e, 0x8f, 0x68, 0xb2, 0xf8,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x0b, 0x68, 0x55, 0x8e, 0x02, 0x00, 0x00,
}