// Code generated by protoc-gen-go. DO NOT EDIT.
// source: football.proto

package apiproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 比赛投注信息
type GameInfo struct {
	StartTime int64     `protobuf:"varint,1,opt,name=startTime" json:"startTime"`
	GameNo    string    `protobuf:"bytes,2,opt,name=gameNo" json:"gameNo"`
	HomeTeam  string    `protobuf:"bytes,3,opt,name=homeTeam" json:"homeTeam"`
	GuestTeam string    `protobuf:"bytes,4,opt,name=guestTeam" json:"guestTeam"`
	GameType  string    `protobuf:"bytes,5,opt,name=gameType" json:"gameType"`
	Giveball  string    `protobuf:"bytes,6,opt,name=giveball" json:"giveball"`
	EndTime   int64     `protobuf:"varint,7,opt,name=endTime" json:"endTime"`
	HomeIcon  string    `protobuf:"bytes,8,opt,name=homeIcon" json:"homeIcon"`
	GuestIcon string    `protobuf:"bytes,9,opt,name=guestIcon" json:"guestIcon"`
	Spf       []float64 `protobuf:"fixed64,10,rep,packed,name=spf" json:"spf"`
	SpfDg     bool      `protobuf:"varint,11,opt,name=spfDg" json:"spfDg"`
	Rqspf     []float64 `protobuf:"fixed64,12,rep,packed,name=rqspf" json:"rqspf"`
	RqspfDg   bool      `protobuf:"varint,13,opt,name=rqspfDg" json:"rqspfDg"`
	Zjqs      []float64 `protobuf:"fixed64,14,rep,packed,name=zjqs" json:"zjqs"`
	ZjqsDg    bool      `protobuf:"varint,15,opt,name=zjqsDg" json:"zjqsDg"`
	Bqc       []float64 `protobuf:"fixed64,16,rep,packed,name=bqc" json:"bqc"`
	BqcDg     bool      `protobuf:"varint,17,opt,name=bqcDg" json:"bqcDg"`
	Bf        []float64 `protobuf:"fixed64,18,rep,packed,name=bf" json:"bf"`
	BfDg      bool      `protobuf:"varint,19,opt,name=bfDg" json:"bfDg"`
}

func (m *GameInfo) Reset()                    { *m = GameInfo{} }
func (m *GameInfo) String() string            { return proto.CompactTextString(m) }
func (*GameInfo) ProtoMessage()               {}
func (*GameInfo) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *GameInfo) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *GameInfo) GetGameNo() string {
	if m != nil {
		return m.GameNo
	}
	return ""
}

func (m *GameInfo) GetHomeTeam() string {
	if m != nil {
		return m.HomeTeam
	}
	return ""
}

func (m *GameInfo) GetGuestTeam() string {
	if m != nil {
		return m.GuestTeam
	}
	return ""
}

func (m *GameInfo) GetGameType() string {
	if m != nil {
		return m.GameType
	}
	return ""
}

func (m *GameInfo) GetGiveball() string {
	if m != nil {
		return m.Giveball
	}
	return ""
}

func (m *GameInfo) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *GameInfo) GetHomeIcon() string {
	if m != nil {
		return m.HomeIcon
	}
	return ""
}

func (m *GameInfo) GetGuestIcon() string {
	if m != nil {
		return m.GuestIcon
	}
	return ""
}

func (m *GameInfo) GetSpf() []float64 {
	if m != nil {
		return m.Spf
	}
	return nil
}

func (m *GameInfo) GetSpfDg() bool {
	if m != nil {
		return m.SpfDg
	}
	return false
}

func (m *GameInfo) GetRqspf() []float64 {
	if m != nil {
		return m.Rqspf
	}
	return nil
}

func (m *GameInfo) GetRqspfDg() bool {
	if m != nil {
		return m.RqspfDg
	}
	return false
}

func (m *GameInfo) GetZjqs() []float64 {
	if m != nil {
		return m.Zjqs
	}
	return nil
}

func (m *GameInfo) GetZjqsDg() bool {
	if m != nil {
		return m.ZjqsDg
	}
	return false
}

func (m *GameInfo) GetBqc() []float64 {
	if m != nil {
		return m.Bqc
	}
	return nil
}

func (m *GameInfo) GetBqcDg() bool {
	if m != nil {
		return m.BqcDg
	}
	return false
}

func (m *GameInfo) GetBf() []float64 {
	if m != nil {
		return m.Bf
	}
	return nil
}

func (m *GameInfo) GetBfDg() bool {
	if m != nil {
		return m.BfDg
	}
	return false
}

// 开奖信息
type FbGameOpencai struct {
	HomeTeam  string `protobuf:"bytes,1,opt,name=homeTeam" json:"homeTeam"`
	GuestTeam string `protobuf:"bytes,2,opt,name=guestTeam" json:"guestTeam"`
	HomeIcon  string `protobuf:"bytes,3,opt,name=homeIcon" json:"homeIcon"`
	GuestIcon string `protobuf:"bytes,4,opt,name=guestIcon" json:"guestIcon"`
	Homeball  int64  `protobuf:"varint,5,opt,name=homeball" json:"homeball"`
	Guestball int64  `protobuf:"varint,6,opt,name=guestball" json:"guestball"`
	Giveball  string `protobuf:"bytes,7,opt,name=giveball" json:"giveball"`
	Spf       string `protobuf:"bytes,8,opt,name=spf" json:"spf"`
	Rqspf     string `protobuf:"bytes,9,opt,name=rqspf" json:"rqspf"`
	Zjqs      string `protobuf:"bytes,10,opt,name=zjqs" json:"zjqs"`
	Bqc       string `protobuf:"bytes,11,opt,name=bqc" json:"bqc"`
}

func (m *FbGameOpencai) Reset()                    { *m = FbGameOpencai{} }
func (m *FbGameOpencai) String() string            { return proto.CompactTextString(m) }
func (*FbGameOpencai) ProtoMessage()               {}
func (*FbGameOpencai) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *FbGameOpencai) GetHomeTeam() string {
	if m != nil {
		return m.HomeTeam
	}
	return ""
}

func (m *FbGameOpencai) GetGuestTeam() string {
	if m != nil {
		return m.GuestTeam
	}
	return ""
}

func (m *FbGameOpencai) GetHomeIcon() string {
	if m != nil {
		return m.HomeIcon
	}
	return ""
}

func (m *FbGameOpencai) GetGuestIcon() string {
	if m != nil {
		return m.GuestIcon
	}
	return ""
}

func (m *FbGameOpencai) GetHomeball() int64 {
	if m != nil {
		return m.Homeball
	}
	return 0
}

func (m *FbGameOpencai) GetGuestball() int64 {
	if m != nil {
		return m.Guestball
	}
	return 0
}

func (m *FbGameOpencai) GetGiveball() string {
	if m != nil {
		return m.Giveball
	}
	return ""
}

func (m *FbGameOpencai) GetSpf() string {
	if m != nil {
		return m.Spf
	}
	return ""
}

func (m *FbGameOpencai) GetRqspf() string {
	if m != nil {
		return m.Rqspf
	}
	return ""
}

func (m *FbGameOpencai) GetZjqs() string {
	if m != nil {
		return m.Zjqs
	}
	return ""
}

func (m *FbGameOpencai) GetBqc() string {
	if m != nil {
		return m.Bqc
	}
	return ""
}

type QueryGameInfoRes struct {
	List []*GameInfo `protobuf:"bytes,1,rep,name=list" json:"list"`
}

func (m *QueryGameInfoRes) Reset()                    { *m = QueryGameInfoRes{} }
func (m *QueryGameInfoRes) String() string            { return proto.CompactTextString(m) }
func (*QueryGameInfoRes) ProtoMessage()               {}
func (*QueryGameInfoRes) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *QueryGameInfoRes) GetList() []*GameInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type QueryOpencaiRes struct {
	List []*FbGameOpencai `protobuf:"bytes,1,rep,name=list" json:"list"`
}

func (m *QueryOpencaiRes) Reset()                    { *m = QueryOpencaiRes{} }
func (m *QueryOpencaiRes) String() string            { return proto.CompactTextString(m) }
func (*QueryOpencaiRes) ProtoMessage()               {}
func (*QueryOpencaiRes) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *QueryOpencaiRes) GetList() []*FbGameOpencai {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*GameInfo)(nil), "apiproto.GameInfo")
	proto.RegisterType((*FbGameOpencai)(nil), "apiproto.FbGameOpencai")
	proto.RegisterType((*QueryGameInfoRes)(nil), "apiproto.QueryGameInfoRes")
	proto.RegisterType((*QueryOpencaiRes)(nil), "apiproto.QueryOpencaiRes")
}

func init() { proto.RegisterFile("football.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xcb, 0x8e, 0xd3, 0x40,
	0x10, 0x94, 0x3d, 0x79, 0x38, 0x1d, 0x36, 0x9b, 0x1d, 0x10, 0xb4, 0x10, 0x07, 0x2b, 0x07, 0x64,
	0x09, 0x29, 0x07, 0xb8, 0x71, 0xe0, 0xb0, 0x8a, 0x16, 0xed, 0x85, 0x87, 0x95, 0x1f, 0xf0, 0x98,
	0xb1, 0x31, 0x8a, 0x33, 0x76, 0xec, 0x45, 0x5a, 0x7e, 0x88, 0x3f, 0xe2, 0x7b, 0x50, 0xf7, 0xcc,
	0xd8, 0x64, 0x85, 0xf6, 0xe4, 0xa9, 0xae, 0x2a, 0x4f, 0xa9, 0x4b, 0x03, 0xab, 0xc2, 0x98, 0x5e,
	0x65, 0x87, 0xc3, 0xb6, 0x39, 0x99, 0xde, 0xc8, 0x28, 0x6b, 0x2a, 0x3e, 0x6d, 0xfe, 0x08, 0x88,
	0x3e, 0x66, 0xb5, 0xbe, 0x3d, 0x16, 0x46, 0xbe, 0x82, 0x45, 0xd7, 0x67, 0xa7, 0x7e, 0x5f, 0xd5,
	0x1a, 0x83, 0x38, 0x48, 0x44, 0x3a, 0x0e, 0xe4, 0x73, 0x98, 0x95, 0x59, 0xad, 0x3f, 0x19, 0x0c,
	0xe3, 0x20, 0x59, 0xa4, 0x0e, 0xc9, 0x97, 0x10, 0x7d, 0x37, 0xb5, 0xde, 0xeb, 0xac, 0x46, 0xc1,
	0xcc, 0x80, 0xe9, 0x8f, 0xe5, 0x9d, 0xee, 0x7a, 0x26, 0x27, 0x4c, 0x8e, 0x03, 0x72, 0xd2, 0x3f,
	0xf6, 0xf7, 0x8d, 0xc6, 0xa9, 0x75, 0x7a, 0xcc, 0x5c, 0xf5, 0x53, 0x53, 0x68, 0x9c, 0x39, 0xce,
	0x61, 0x89, 0x30, 0xd7, 0xc7, 0x6f, 0x9c, 0x72, 0xce, 0x29, 0x3d, 0xf4, 0x59, 0x6e, 0x73, 0x73,
	0xc4, 0x68, 0xcc, 0x42, 0x78, 0xc8, 0xc2, 0xe4, 0xe2, 0x9f, 0x2c, 0xcc, 0xae, 0x41, 0x74, 0x4d,
	0x81, 0x10, 0x8b, 0x24, 0x48, 0xe9, 0x28, 0x9f, 0xc1, 0xb4, 0x6b, 0x8a, 0x5d, 0x89, 0xcb, 0x38,
	0x48, 0xa2, 0xd4, 0x02, 0x9a, 0x9e, 0x5a, 0x52, 0x3e, 0x61, 0xa5, 0x05, 0x94, 0x88, 0x0f, 0xbb,
	0x12, 0x2f, 0x58, 0xed, 0xa1, 0x94, 0x30, 0xf9, 0xf5, 0xa3, 0xed, 0x70, 0xc5, 0x72, 0x3e, 0xd3,
	0x26, 0xe9, 0xbb, 0x2b, 0xf1, 0x92, 0xc5, 0x0e, 0x51, 0x06, 0xd5, 0xe6, 0xb8, 0xb6, 0x19, 0x54,
	0x9b, 0xd3, 0x6d, 0xaa, 0xcd, 0x77, 0x25, 0x5e, 0xd9, 0x0c, 0x0c, 0xe4, 0x0a, 0x42, 0x55, 0xa0,
	0x64, 0x59, 0xa8, 0x0a, 0xba, 0x43, 0xd1, 0xd5, 0x4f, 0x59, 0xc4, 0xe7, 0xcd, 0xef, 0x10, 0x2e,
	0x6e, 0x14, 0x55, 0xfb, 0xb9, 0xd1, 0xc7, 0x3c, 0xab, 0xce, 0x7a, 0x0a, 0x1e, 0xeb, 0x29, 0xfc,
	0x4f, 0x4f, 0xc3, 0x56, 0xc5, 0x63, 0x5b, 0x9d, 0x3c, 0xdc, 0xaa, 0x73, 0x72, 0x8b, 0x53, 0xae,
	0x6a, 0xc0, 0x83, 0x73, 0xa8, 0x58, 0xa4, 0xe3, 0xe0, 0xac, 0xff, 0xf9, 0x83, 0xfe, 0x5d, 0x57,
	0xb6, 0x60, 0xdf, 0x95, 0x6d, 0xc5, 0xf6, 0xea, 0x5a, 0xf1, 0xbb, 0x07, 0x1e, 0xda, 0xdd, 0xbb,
	0x1d, 0x2f, 0xad, 0x57, 0xb5, 0xf9, 0xe6, 0x3d, 0xac, 0xbf, 0xde, 0xe9, 0xd3, 0xbd, 0x7f, 0x06,
	0xa9, 0xee, 0xe4, 0x6b, 0x98, 0x1c, 0xaa, 0xae, 0xc7, 0x20, 0x16, 0xc9, 0xf2, 0xad, 0xdc, 0xfa,
	0xf7, 0xb2, 0x1d, 0x44, 0xcc, 0x6f, 0x3e, 0xc0, 0x25, 0x7b, 0xdd, 0x8e, 0xc9, 0xfa, 0xe6, 0xcc,
	0xfa, 0x62, 0xb4, 0x9e, 0xb5, 0x61, 0xfd, 0xd7, 0x57, 0x30, 0xfd, 0x42, 0xe4, 0x75, 0x74, 0xe3,
	0x5e, 0xa8, 0x9a, 0xb1, 0xfa, 0xdd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73, 0xa6, 0x32, 0x9d,
	0xb4, 0x03, 0x00, 0x00,
}
