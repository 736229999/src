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

type LotteryType int32

const (
	LotteryType_AllType  LotteryType = 0
	LotteryType_LowFreq  LotteryType = 1
	LotteryType_HighFreq LotteryType = 2
	LotteryType_Comp     LotteryType = 3
)

var LotteryType_name = map[int32]string{
	0: "AllType",
	1: "LowFreq",
	2: "HighFreq",
	3: "Comp",
}
var LotteryType_value = map[string]int32{
	"AllType":  0,
	"LowFreq":  1,
	"HighFreq": 2,
	"Comp":     3,
}

func (x LotteryType) String() string {
	return proto.EnumName(LotteryType_name, int32(x))
}
func (LotteryType) EnumDescriptor() ([]byte, []int) { return fileDescriptor19, []int{0} }

type LotteryId int32

const (
	LotteryId_AllId  LotteryId = 0
	LotteryId_Dlt    LotteryId = 1
	LotteryId_Fc3d   LotteryId = 2
	LotteryId_Ssq    LotteryId = 3
	LotteryId_Cqssc  LotteryId = 4
	LotteryId_Bjpk10 LotteryId = 5
	LotteryId_Gd11x5 LotteryId = 6
	LotteryId_Pl3    LotteryId = 7
	LotteryId_Pl5    LotteryId = 8
	LotteryId_Jczq   LotteryId = 9
	LotteryId_Jclq   LotteryId = 10
)

var LotteryId_name = map[int32]string{
	0:  "AllId",
	1:  "Dlt",
	2:  "Fc3d",
	3:  "Ssq",
	4:  "Cqssc",
	5:  "Bjpk10",
	6:  "Gd11x5",
	7:  "Pl3",
	8:  "Pl5",
	9:  "Jczq",
	10: "Jclq",
}
var LotteryId_value = map[string]int32{
	"AllId":  0,
	"Dlt":    1,
	"Fc3d":   2,
	"Ssq":    3,
	"Cqssc":  4,
	"Bjpk10": 5,
	"Gd11x5": 6,
	"Pl3":    7,
	"Pl5":    8,
	"Jczq":   9,
	"Jclq":   10,
}

func (x LotteryId) String() string {
	return proto.EnumName(LotteryId_name, int32(x))
}
func (LotteryId) EnumDescriptor() ([]byte, []int) { return fileDescriptor19, []int{1} }

type LotteryWinningNo struct {
	Lottery     string `protobuf:"bytes,1,opt,name=lottery" json:"lottery"`
	Issue       string `protobuf:"bytes,2,opt,name=issue" json:"issue"`
	StartTime   int64  `protobuf:"varint,3,opt,name=start_time,json=startTime" json:"start_time"`
	EndTime     int64  `protobuf:"varint,4,opt,name=end_time,json=endTime" json:"end_time"`
	OpenTime    int64  `protobuf:"varint,5,opt,name=open_time,json=openTime" json:"open_time"`
	OpenBalls   string `protobuf:"bytes,6,opt,name=open_balls,json=openBalls" json:"open_balls"`
	LotteryName string `protobuf:"bytes,7,opt,name=lottery_name,json=lotteryName" json:"lottery_name"`
	Id          int64  `protobuf:"varint,8,opt,name=id" json:"id"`
}

func (m *LotteryWinningNo) Reset()                    { *m = LotteryWinningNo{} }
func (m *LotteryWinningNo) String() string            { return proto.CompactTextString(m) }
func (*LotteryWinningNo) ProtoMessage()               {}
func (*LotteryWinningNo) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{0} }

func (m *LotteryWinningNo) GetLottery() string {
	if m != nil {
		return m.Lottery
	}
	return ""
}

func (m *LotteryWinningNo) GetIssue() string {
	if m != nil {
		return m.Issue
	}
	return ""
}

func (m *LotteryWinningNo) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *LotteryWinningNo) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *LotteryWinningNo) GetOpenTime() int64 {
	if m != nil {
		return m.OpenTime
	}
	return 0
}

func (m *LotteryWinningNo) GetOpenBalls() string {
	if m != nil {
		return m.OpenBalls
	}
	return ""
}

func (m *LotteryWinningNo) GetLotteryName() string {
	if m != nil {
		return m.LotteryName
	}
	return ""
}

func (m *LotteryWinningNo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type LotteryWinningNoList struct {
	List []*LotteryWinningNo `protobuf:"bytes,1,rep,name=list" json:"list"`
}

func (m *LotteryWinningNoList) Reset()                    { *m = LotteryWinningNoList{} }
func (m *LotteryWinningNoList) String() string            { return proto.CompactTextString(m) }
func (*LotteryWinningNoList) ProtoMessage()               {}
func (*LotteryWinningNoList) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{1} }

func (m *LotteryWinningNoList) GetList() []*LotteryWinningNo {
	if m != nil {
		return m.List
	}
	return nil
}

type WinningLotteryList struct {
	List []*WinningLottery `protobuf:"bytes,1,rep,name=list" json:"list"`
}

func (m *WinningLotteryList) Reset()                    { *m = WinningLotteryList{} }
func (m *WinningLotteryList) String() string            { return proto.CompactTextString(m) }
func (*WinningLotteryList) ProtoMessage()               {}
func (*WinningLotteryList) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{2} }

func (m *WinningLotteryList) GetList() []*WinningLottery {
	if m != nil {
		return m.List
	}
	return nil
}

type WinningLottery struct {
	Lottery string `protobuf:"bytes,1,opt,name=lottery" json:"lottery"`
}

func (m *WinningLottery) Reset()                    { *m = WinningLottery{} }
func (m *WinningLottery) String() string            { return proto.CompactTextString(m) }
func (*WinningLottery) ProtoMessage()               {}
func (*WinningLottery) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{3} }

func (m *WinningLottery) GetLottery() string {
	if m != nil {
		return m.Lottery
	}
	return ""
}

type WinningListByLottery struct {
	List      []*LotteryWinningNo `protobuf:"bytes,1,rep,name=list" json:"list"`
	Page      int32               `protobuf:"varint,2,opt,name=page" json:"page"`
	Size      int32               `protobuf:"varint,3,opt,name=size" json:"size"`
	Total     int32               `protobuf:"varint,4,opt,name=total" json:"total"`
	Issue     string              `protobuf:"bytes,5,opt,name=issue" json:"issue"`
	Lottery   string              `protobuf:"bytes,6,opt,name=lottery" json:"lottery"`
	StartTime int32               `protobuf:"varint,7,opt,name=start_time,json=startTime" json:"start_time"`
	EndTime   int32               `protobuf:"varint,8,opt,name=end_time,json=endTime" json:"end_time"`
}

func (m *WinningListByLottery) Reset()                    { *m = WinningListByLottery{} }
func (m *WinningListByLottery) String() string            { return proto.CompactTextString(m) }
func (*WinningListByLottery) ProtoMessage()               {}
func (*WinningListByLottery) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{4} }

func (m *WinningListByLottery) GetList() []*LotteryWinningNo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *WinningListByLottery) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *WinningListByLottery) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *WinningListByLottery) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *WinningListByLottery) GetIssue() string {
	if m != nil {
		return m.Issue
	}
	return ""
}

func (m *WinningListByLottery) GetLottery() string {
	if m != nil {
		return m.Lottery
	}
	return ""
}

func (m *WinningListByLottery) GetStartTime() int32 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *WinningListByLottery) GetEndTime() int32 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func init() {
	proto.RegisterType((*LotteryWinningNo)(nil), "dbproto.LotteryWinningNo")
	proto.RegisterType((*LotteryWinningNoList)(nil), "dbproto.LotteryWinningNoList")
	proto.RegisterType((*WinningLotteryList)(nil), "dbproto.WinningLotteryList")
	proto.RegisterType((*WinningLottery)(nil), "dbproto.WinningLottery")
	proto.RegisterType((*WinningListByLottery)(nil), "dbproto.WinningListByLottery")
	proto.RegisterEnum("dbproto.LotteryType", LotteryType_name, LotteryType_value)
	proto.RegisterEnum("dbproto.LotteryId", LotteryId_name, LotteryId_value)
}

func init() { proto.RegisterFile("opencai.proto", fileDescriptor19) }

var fileDescriptor19 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0xed, 0x7e, 0xef, 0xde, 0xd4, 0x32, 0x0c, 0x01, 0xb7, 0x48, 0x21, 0xe6, 0x29, 0x44, 0x0c,
	0xc6, 0xd0, 0x67, 0x49, 0xaa, 0xd5, 0x4a, 0x28, 0xb2, 0x16, 0x7c, 0x2c, 0x9b, 0xcc, 0x10, 0x47,
	0x67, 0xbf, 0x32, 0x23, 0x9a, 0xfa, 0x93, 0xc5, 0xff, 0x20, 0x73, 0x67, 0xa3, 0xdd, 0x42, 0x1e,
	0x7c, 0x3b, 0xf7, 0x9c, 0x93, 0x93, 0xbb, 0x73, 0x2e, 0x3c, 0xaa, 0x6a, 0x5e, 0xae, 0x73, 0x31,
	0xa9, 0xb7, 0x95, 0xae, 0x68, 0xc4, 0x56, 0x08, 0x86, 0xbf, 0x1c, 0x20, 0xcb, 0x4a, 0x6b, 0xbe,
	0xdd, 0x7d, 0x12, 0x65, 0x29, 0xca, 0xcd, 0x75, 0x45, 0x53, 0x88, 0xa4, 0xe5, 0x52, 0x67, 0xe0,
	0x8c, 0x92, 0x6c, 0x3f, 0xd2, 0x3e, 0x04, 0x42, 0xa9, 0x6f, 0x3c, 0x75, 0x91, 0xb7, 0x03, 0x3d,
	0x03, 0x50, 0x3a, 0xdf, 0xea, 0x5b, 0x2d, 0x0a, 0x9e, 0x7a, 0x03, 0x67, 0xe4, 0x65, 0x09, 0x32,
	0x37, 0xa2, 0xe0, 0xf4, 0x14, 0x62, 0x5e, 0x32, 0x2b, 0xfa, 0x28, 0x46, 0xbc, 0x64, 0x28, 0x3d,
	0x81, 0xc4, 0x2c, 0x66, 0xb5, 0x00, 0xb5, 0xd8, 0x10, 0x28, 0x9e, 0x01, 0xa0, 0xb8, 0xca, 0xa5,
	0x54, 0x69, 0x88, 0xff, 0x88, 0xf6, 0x85, 0x21, 0xe8, 0x53, 0x38, 0x6e, 0xd7, 0xba, 0x2d, 0xf3,
	0x82, 0xa7, 0x11, 0x1a, 0x7a, 0x2d, 0x77, 0x9d, 0x17, 0x9c, 0x9e, 0x80, 0x2b, 0x58, 0x1a, 0x63,
	0xae, 0x2b, 0xd8, 0xf0, 0x0d, 0xf4, 0x1f, 0x7e, 0xec, 0x52, 0x28, 0x4d, 0x9f, 0x83, 0x2f, 0x85,
	0xd2, 0xa9, 0x33, 0xf0, 0x46, 0xbd, 0x97, 0xa7, 0x93, 0xf6, 0x75, 0x26, 0x0f, 0xcd, 0x19, 0xda,
	0x86, 0x73, 0xa0, 0x2d, 0xd5, 0x1a, 0x30, 0xe4, 0x59, 0x27, 0xe4, 0xf1, 0xdf, 0x90, 0xae, 0xb5,
	0x8d, 0x18, 0xc3, 0x49, 0x97, 0x3f, 0xfc, 0xe8, 0xc3, 0xdf, 0x0e, 0xf4, 0xf7, 0x66, 0xa1, 0xf4,
	0x62, 0xb7, 0xff, 0xc9, 0xff, 0xad, 0x4d, 0x29, 0xf8, 0x75, 0xbe, 0xb1, 0xdd, 0x05, 0x19, 0x62,
	0xc3, 0x29, 0x71, 0x67, 0x4b, 0x0b, 0x32, 0xc4, 0xa6, 0x64, 0x5d, 0xe9, 0x5c, 0x62, 0x59, 0x41,
	0x66, 0x87, 0x7f, 0xd5, 0x07, 0xf7, 0xab, 0xbf, 0xb7, 0x75, 0xd8, 0x3d, 0x95, 0xee, 0x51, 0x44,
	0x18, 0x75, 0xe0, 0x28, 0x62, 0x14, 0xf7, 0x47, 0x31, 0x7e, 0x05, 0xbd, 0xf6, 0x0b, 0x6e, 0x76,
	0x35, 0xa7, 0x3d, 0x88, 0xe6, 0x52, 0x1a, 0x48, 0x8e, 0xcc, 0xb0, 0xac, 0xbe, 0x5f, 0x6e, 0x79,
	0x43, 0x1c, 0x7a, 0x0c, 0xf1, 0x3b, 0xb1, 0xf9, 0x8c, 0x93, 0x4b, 0x63, 0xf0, 0x2f, 0xaa, 0xa2,
	0x26, 0xde, 0xf8, 0x27, 0x24, 0x6d, 0xc0, 0x15, 0xa3, 0x09, 0x04, 0x73, 0x29, 0xaf, 0x18, 0x39,
	0xa2, 0x11, 0x78, 0xaf, 0xa5, 0x26, 0x8e, 0xb1, 0x5e, 0xae, 0x67, 0x8c, 0xb8, 0x86, 0xfa, 0xa8,
	0x1a, 0xe2, 0x19, 0xdb, 0x45, 0xa3, 0xd4, 0x9a, 0xf8, 0x14, 0x20, 0x5c, 0x7c, 0xa9, 0xbf, 0x4e,
	0x5f, 0x90, 0xc0, 0xe0, 0xb7, 0x6c, 0x3a, 0xfd, 0x71, 0x4e, 0x42, 0xe3, 0xfd, 0x20, 0x67, 0x24,
	0xb2, 0xe0, 0x9c, 0xc4, 0x26, 0xe7, 0xfd, 0xfa, 0xae, 0x21, 0x89, 0x45, 0xb2, 0x21, 0xb0, 0x0a,
	0xb1, 0x83, 0xd9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0xfb, 0x33, 0x8e, 0x72, 0x03, 0x00,
	0x00,
}