// Code generated by protoc-gen-go. DO NOT EDIT.
// source: buycai.proto

package dbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BuycaiOptions struct {
	Id        int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	Issue     string `protobuf:"bytes,2,opt,name=issue" json:"issue"`
	StartTime int64  `protobuf:"varint,3,opt,name=start_time,json=startTime" json:"start_time"`
	EndTime   int64  `protobuf:"varint,4,opt,name=end_time,json=endTime" json:"end_time"`
	OpenTime  int64  `protobuf:"varint,5,opt,name=open_time,json=openTime" json:"open_time"`
	OpenBalls string `protobuf:"bytes,6,opt,name=open_balls,json=openBalls" json:"open_balls"`
}

func (m *BuycaiOptions) Reset()                    { *m = BuycaiOptions{} }
func (m *BuycaiOptions) String() string            { return proto.CompactTextString(m) }
func (*BuycaiOptions) ProtoMessage()               {}
func (*BuycaiOptions) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *BuycaiOptions) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BuycaiOptions) GetIssue() string {
	if m != nil {
		return m.Issue
	}
	return ""
}

func (m *BuycaiOptions) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *BuycaiOptions) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *BuycaiOptions) GetOpenTime() int64 {
	if m != nil {
		return m.OpenTime
	}
	return 0
}

func (m *BuycaiOptions) GetOpenBalls() string {
	if m != nil {
		return m.OpenBalls
	}
	return ""
}

type BuycaiOptionsReply struct {
	Size    int64            `protobuf:"varint,1,opt,name=size" json:"size"`
	Page    int64            `protobuf:"varint,2,opt,name=page" json:"page"`
	Total   int64            `protobuf:"varint,3,opt,name=total" json:"total"`
	Lottery string           `protobuf:"bytes,4,opt,name=lottery" json:"lottery"`
	Buycai  []*BuycaiOptions `protobuf:"bytes,5,rep,name=buycai" json:"buycai"`
}

func (m *BuycaiOptionsReply) Reset()                    { *m = BuycaiOptionsReply{} }
func (m *BuycaiOptionsReply) String() string            { return proto.CompactTextString(m) }
func (*BuycaiOptionsReply) ProtoMessage()               {}
func (*BuycaiOptionsReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *BuycaiOptionsReply) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *BuycaiOptionsReply) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *BuycaiOptionsReply) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *BuycaiOptionsReply) GetLottery() string {
	if m != nil {
		return m.Lottery
	}
	return ""
}

func (m *BuycaiOptionsReply) GetBuycai() []*BuycaiOptions {
	if m != nil {
		return m.Buycai
	}
	return nil
}

type BuycaiOptionsIssue struct {
	Num      int64            `protobuf:"varint,1,opt,name=num" json:"num"`
	Issue    string           `protobuf:"bytes,2,opt,name=issue" json:"issue"`
	Buycai   []*BuycaiOptions `protobuf:"bytes,3,rep,name=buycai" json:"buycai"`
	DayMaxNo int64            `protobuf:"varint,4,opt,name=day_max_no,json=dayMaxNo" json:"day_max_no"`
	Lottery  string           `protobuf:"bytes,5,opt,name=lottery" json:"lottery"`
	Id       int64            `protobuf:"varint,6,opt,name=id" json:"id"`
}

func (m *BuycaiOptionsIssue) Reset()                    { *m = BuycaiOptionsIssue{} }
func (m *BuycaiOptionsIssue) String() string            { return proto.CompactTextString(m) }
func (*BuycaiOptionsIssue) ProtoMessage()               {}
func (*BuycaiOptionsIssue) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *BuycaiOptionsIssue) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *BuycaiOptionsIssue) GetIssue() string {
	if m != nil {
		return m.Issue
	}
	return ""
}

func (m *BuycaiOptionsIssue) GetBuycai() []*BuycaiOptions {
	if m != nil {
		return m.Buycai
	}
	return nil
}

func (m *BuycaiOptionsIssue) GetDayMaxNo() int64 {
	if m != nil {
		return m.DayMaxNo
	}
	return 0
}

func (m *BuycaiOptionsIssue) GetLottery() string {
	if m != nil {
		return m.Lottery
	}
	return ""
}

func (m *BuycaiOptionsIssue) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type BuycaiOptionsUpdateIssue struct {
	Id      int64          `protobuf:"varint,1,opt,name=id" json:"id"`
	Lottery string         `protobuf:"bytes,2,opt,name=lottery" json:"lottery"`
	Buycai  *BuycaiOptions `protobuf:"bytes,3,opt,name=buycai" json:"buycai"`
}

func (m *BuycaiOptionsUpdateIssue) Reset()                    { *m = BuycaiOptionsUpdateIssue{} }
func (m *BuycaiOptionsUpdateIssue) String() string            { return proto.CompactTextString(m) }
func (*BuycaiOptionsUpdateIssue) ProtoMessage()               {}
func (*BuycaiOptionsUpdateIssue) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *BuycaiOptionsUpdateIssue) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BuycaiOptionsUpdateIssue) GetLottery() string {
	if m != nil {
		return m.Lottery
	}
	return ""
}

func (m *BuycaiOptionsUpdateIssue) GetBuycai() *BuycaiOptions {
	if m != nil {
		return m.Buycai
	}
	return nil
}

type PlayTimeSettings struct {
	Id             int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	LotteryId      int64  `protobuf:"varint,2,opt,name=lottery_id,json=lotteryId" json:"lottery_id"`
	StartTime      int64  `protobuf:"varint,3,opt,name=start_time,json=startTime" json:"start_time"`
	EndTime        int64  `protobuf:"varint,4,opt,name=end_time,json=endTime" json:"end_time"`
	ChaseStartTime int64  `protobuf:"varint,5,opt,name=chase_start_time,json=chaseStartTime" json:"chase_start_time"`
	Name           string `protobuf:"bytes,6,opt,name=name" json:"name"`
}

func (m *PlayTimeSettings) Reset()                    { *m = PlayTimeSettings{} }
func (m *PlayTimeSettings) String() string            { return proto.CompactTextString(m) }
func (*PlayTimeSettings) ProtoMessage()               {}
func (*PlayTimeSettings) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *PlayTimeSettings) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PlayTimeSettings) GetLotteryId() int64 {
	if m != nil {
		return m.LotteryId
	}
	return 0
}

func (m *PlayTimeSettings) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *PlayTimeSettings) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *PlayTimeSettings) GetChaseStartTime() int64 {
	if m != nil {
		return m.ChaseStartTime
	}
	return 0
}

func (m *PlayTimeSettings) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PlayTimeSettingsList struct {
	List []*PlayTimeSettings `protobuf:"bytes,1,rep,name=list" json:"list"`
}

func (m *PlayTimeSettingsList) Reset()                    { *m = PlayTimeSettingsList{} }
func (m *PlayTimeSettingsList) String() string            { return proto.CompactTextString(m) }
func (*PlayTimeSettingsList) ProtoMessage()               {}
func (*PlayTimeSettingsList) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *PlayTimeSettingsList) GetList() []*PlayTimeSettings {
	if m != nil {
		return m.List
	}
	return nil
}

type BuycaiStatistics struct {
	BuycaiTime          int32  `protobuf:"varint,1,opt,name=buycai_time,json=buycaiTime" json:"buycai_time"`
	BuycaiNum           int32  `protobuf:"varint,2,opt,name=buycai_num,json=buycaiNum" json:"buycai_num"`
	BuycaiAmount        string `protobuf:"bytes,3,opt,name=buycai_amount,json=buycaiAmount" json:"buycai_amount"`
	IosBuycaiNum        int32  `protobuf:"varint,4,opt,name=ios_buycai_num,json=iosBuycaiNum" json:"ios_buycai_num"`
	IosBuycaiAmount     string `protobuf:"bytes,5,opt,name=ios_buycai_amount,json=iosBuycaiAmount" json:"ios_buycai_amount"`
	AndroidBuycaiNum    int32  `protobuf:"varint,6,opt,name=android_buycai_num,json=androidBuycaiNum" json:"android_buycai_num"`
	AndroidBuycaiAmount string `protobuf:"bytes,7,opt,name=android_buycai_amount,json=androidBuycaiAmount" json:"android_buycai_amount"`
}

func (m *BuycaiStatistics) Reset()                    { *m = BuycaiStatistics{} }
func (m *BuycaiStatistics) String() string            { return proto.CompactTextString(m) }
func (*BuycaiStatistics) ProtoMessage()               {}
func (*BuycaiStatistics) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *BuycaiStatistics) GetBuycaiTime() int32 {
	if m != nil {
		return m.BuycaiTime
	}
	return 0
}

func (m *BuycaiStatistics) GetBuycaiNum() int32 {
	if m != nil {
		return m.BuycaiNum
	}
	return 0
}

func (m *BuycaiStatistics) GetBuycaiAmount() string {
	if m != nil {
		return m.BuycaiAmount
	}
	return ""
}

func (m *BuycaiStatistics) GetIosBuycaiNum() int32 {
	if m != nil {
		return m.IosBuycaiNum
	}
	return 0
}

func (m *BuycaiStatistics) GetIosBuycaiAmount() string {
	if m != nil {
		return m.IosBuycaiAmount
	}
	return ""
}

func (m *BuycaiStatistics) GetAndroidBuycaiNum() int32 {
	if m != nil {
		return m.AndroidBuycaiNum
	}
	return 0
}

func (m *BuycaiStatistics) GetAndroidBuycaiAmount() string {
	if m != nil {
		return m.AndroidBuycaiAmount
	}
	return ""
}

type BuycaiStatisticsList struct {
	StartTime int32               `protobuf:"varint,1,opt,name=start_time,json=startTime" json:"start_time"`
	EndTime   int32               `protobuf:"varint,2,opt,name=end_time,json=endTime" json:"end_time"`
	Page      int32               `protobuf:"varint,3,opt,name=page" json:"page"`
	Size      int32               `protobuf:"varint,4,opt,name=size" json:"size"`
	Total     int32               `protobuf:"varint,5,opt,name=total" json:"total"`
	List      []*BuycaiStatistics `protobuf:"bytes,6,rep,name=list" json:"list"`
}

func (m *BuycaiStatisticsList) Reset()                    { *m = BuycaiStatisticsList{} }
func (m *BuycaiStatisticsList) String() string            { return proto.CompactTextString(m) }
func (*BuycaiStatisticsList) ProtoMessage()               {}
func (*BuycaiStatisticsList) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *BuycaiStatisticsList) GetStartTime() int32 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *BuycaiStatisticsList) GetEndTime() int32 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *BuycaiStatisticsList) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *BuycaiStatisticsList) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *BuycaiStatisticsList) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *BuycaiStatisticsList) GetList() []*BuycaiStatistics {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*BuycaiOptions)(nil), "dbproto.BuycaiOptions")
	proto.RegisterType((*BuycaiOptionsReply)(nil), "dbproto.BuycaiOptionsReply")
	proto.RegisterType((*BuycaiOptionsIssue)(nil), "dbproto.BuycaiOptionsIssue")
	proto.RegisterType((*BuycaiOptionsUpdateIssue)(nil), "dbproto.BuycaiOptionsUpdateIssue")
	proto.RegisterType((*PlayTimeSettings)(nil), "dbproto.PlayTimeSettings")
	proto.RegisterType((*PlayTimeSettingsList)(nil), "dbproto.PlayTimeSettingsList")
	proto.RegisterType((*BuycaiStatistics)(nil), "dbproto.BuycaiStatistics")
	proto.RegisterType((*BuycaiStatisticsList)(nil), "dbproto.BuycaiStatisticsList")
}

func init() { proto.RegisterFile("buycai.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0x95, 0xed, 0xd8, 0xa9, 0xa7, 0x6d, 0xbe, 0x7c, 0x4b, 0x40, 0xae, 0x20, 0x22, 0x32, 0x1c,
	0x22, 0x04, 0x39, 0x94, 0x5f, 0x40, 0x24, 0x0e, 0x95, 0xa0, 0x20, 0x07, 0xce, 0xd6, 0x26, 0xbb,
	0x2a, 0x2b, 0xd9, 0x5e, 0x2b, 0xbb, 0x91, 0x6a, 0x7e, 0x0b, 0xff, 0xa0, 0x57, 0x2e, 0x1c, 0xf8,
	0x6f, 0xc8, 0xb3, 0x6b, 0xd7, 0x36, 0x8a, 0x40, 0xe2, 0xb6, 0x7e, 0x6f, 0xfc, 0xf6, 0xf9, 0xcd,
	0x78, 0xe0, 0x6c, 0x7b, 0xa8, 0x76, 0x54, 0xac, 0xca, 0xbd, 0xd4, 0x92, 0x8c, 0xd9, 0x16, 0x0f,
	0xf1, 0x9d, 0x03, 0xe7, 0x6b, 0x64, 0x3e, 0x94, 0x5a, 0xc8, 0x42, 0x91, 0x09, 0xb8, 0x82, 0x45,
	0xce, 0xc2, 0x59, 0x7a, 0x89, 0x2b, 0x18, 0x99, 0x81, 0x2f, 0x94, 0x3a, 0xf0, 0xc8, 0x5d, 0x38,
	0xcb, 0x30, 0x31, 0x0f, 0x64, 0x0e, 0xa0, 0x34, 0xdd, 0xeb, 0x54, 0x8b, 0x9c, 0x47, 0x1e, 0x56,
	0x87, 0x88, 0x7c, 0x12, 0x39, 0x27, 0x17, 0x70, 0xc2, 0x0b, 0x66, 0xc8, 0x11, 0x92, 0x63, 0x5e,
	0x30, 0xa4, 0x1e, 0x43, 0x28, 0x4b, 0x5e, 0x18, 0xce, 0x47, 0xee, 0xa4, 0x06, 0x90, 0x9c, 0x03,
	0x20, 0xb9, 0xa5, 0x59, 0xa6, 0xa2, 0x00, 0x6f, 0xc4, 0xf2, 0x75, 0x0d, 0xc4, 0xdf, 0x1c, 0x20,
	0x3d, 0xb7, 0x09, 0x2f, 0xb3, 0x8a, 0x10, 0x18, 0x29, 0xf1, 0x95, 0x5b, 0xd3, 0x78, 0xae, 0xb1,
	0x92, 0xde, 0x18, 0xd7, 0x5e, 0x82, 0xe7, 0xfa, 0x53, 0xb4, 0xd4, 0x34, 0xb3, 0x7e, 0xcd, 0x03,
	0x89, 0x60, 0x9c, 0x49, 0xad, 0xf9, 0xbe, 0x42, 0xab, 0x61, 0xd2, 0x3c, 0x92, 0x15, 0x04, 0x26,
	0xb5, 0xc8, 0x5f, 0x78, 0xcb, 0xd3, 0xcb, 0x47, 0x2b, 0x1b, 0xdb, 0xaa, 0x6f, 0xc2, 0x56, 0xc5,
	0xdf, 0x87, 0xf6, 0xae, 0x30, 0xab, 0x29, 0x78, 0xc5, 0x21, 0xb7, 0xee, 0xea, 0xe3, 0x91, 0x4c,
	0xef, 0xaf, 0xf3, 0xfe, 0xe6, 0x3a, 0xf2, 0x04, 0x80, 0xd1, 0x2a, 0xcd, 0xe9, 0x6d, 0x5a, 0x48,
	0x1b, 0xf3, 0x09, 0xa3, 0xd5, 0x7b, 0x7a, 0x7b, 0x2d, 0xbb, 0x9f, 0xe5, 0xf7, 0x3f, 0xcb, 0x74,
	0x38, 0x68, 0x3a, 0x1c, 0x6b, 0x88, 0x7a, 0x17, 0x7c, 0x2e, 0x19, 0xd5, 0xdc, 0x78, 0x1f, 0x4e,
	0x43, 0x47, 0xd5, 0x3d, 0x16, 0x56, 0x9d, 0xee, 0x9f, 0xc3, 0xfa, 0xe1, 0xc0, 0xf4, 0x63, 0x46,
	0xab, 0xba, 0xef, 0x1b, 0xae, 0xb5, 0x28, 0x6e, 0x7e, 0x1f, 0xbe, 0x39, 0x80, 0xd5, 0x4f, 0x05,
	0xb3, 0xbd, 0x0c, 0x2d, 0x72, 0xc5, 0xfe, 0x61, 0x0a, 0x97, 0x30, 0xdd, 0x7d, 0xa1, 0x8a, 0xa7,
	0x9d, 0xf7, 0xcd, 0x30, 0x4e, 0x10, 0xdf, 0xb4, 0x22, 0x04, 0x46, 0x05, 0xcd, 0xb9, 0x1d, 0x46,
	0x3c, 0xc7, 0x6f, 0x61, 0x36, 0xb4, 0xfe, 0x4e, 0x28, 0x4d, 0x5e, 0xc1, 0x28, 0x13, 0x4a, 0x47,
	0x0e, 0xf6, 0xef, 0xa2, 0x4d, 0x60, 0x58, 0x9c, 0x60, 0x59, 0x7c, 0xe7, 0xc2, 0xd4, 0x84, 0xb3,
	0xd1, 0x54, 0x0b, 0xa5, 0xc5, 0x4e, 0x91, 0xa7, 0x70, 0x6a, 0x12, 0x32, 0xa6, 0xea, 0x2c, 0xfc,
	0x04, 0x0c, 0xd4, 0xfc, 0x23, 0xb6, 0xa0, 0x9e, 0x2a, 0x17, 0xf9, 0xd0, 0x20, 0xd7, 0x87, 0x9c,
	0x3c, 0x83, 0x73, 0x4b, 0xd3, 0x5c, 0x1e, 0x0a, 0x8d, 0xb1, 0x84, 0x89, 0xfd, 0xff, 0xdf, 0x20,
	0x46, 0x9e, 0xc3, 0x44, 0x48, 0x95, 0x76, 0x74, 0x46, 0xa8, 0x73, 0x26, 0xa4, 0x5a, 0xb7, 0x52,
	0x2f, 0xe0, 0xff, 0x4e, 0x95, 0x95, 0x33, 0xc3, 0xf4, 0x5f, 0x5b, 0x68, 0x15, 0x5f, 0x02, 0xa1,
	0x05, 0xdb, 0x4b, 0xc1, 0xba, 0xaa, 0x01, 0xaa, 0x4e, 0x2d, 0x73, 0xaf, 0x7c, 0x09, 0x0f, 0x07,
	0xd5, 0x56, 0x7d, 0x8c, 0xea, 0x0f, 0x7a, 0x2f, 0x98, 0x1b, 0xe2, 0x9f, 0x0e, 0xcc, 0x86, 0x69,
	0x61, 0xea, 0xfd, 0x29, 0x30, 0x81, 0x1d, 0x99, 0x02, 0x93, 0x56, 0x3b, 0x05, 0xcd, 0x92, 0xf0,
	0x10, 0x36, 0x4b, 0xa2, 0x59, 0x26, 0x26, 0x10, 0xb3, 0x4c, 0xda, 0xc5, 0xe1, 0x23, 0x68, 0x17,
	0x47, 0xd3, 0xed, 0x60, 0xd0, 0xed, 0xa1, 0x49, 0xd3, 0xed, 0x6d, 0x80, 0xec, 0xeb, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x29, 0x3f, 0xf1, 0x4f, 0x8a, 0x05, 0x00, 0x00,
}