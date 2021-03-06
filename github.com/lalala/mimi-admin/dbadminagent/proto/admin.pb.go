// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admin.proto

package dbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AdminUserInfoArg struct {
	Id         int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	Email      string `protobuf:"bytes,2,opt,name=email" json:"email"`
	Username   string `protobuf:"bytes,3,opt,name=username" json:"username"`
	Password   string `protobuf:"bytes,4,opt,name=password" json:"password"`
	Salt       string `protobuf:"bytes,5,opt,name=salt" json:"salt"`
	Status     int64  `protobuf:"varint,6,opt,name=status" json:"status"`
	Mobile     string `protobuf:"bytes,7,opt,name=mobile" json:"mobile"`
	CreateTime int64  `protobuf:"varint,8,opt,name=create_time,json=createTime" json:"create_time"`
	Creator    int64  `protobuf:"varint,9,opt,name=creator" json:"creator"`
	RegisterIp string `protobuf:"bytes,10,opt,name=register_ip,json=registerIp" json:"register_ip"`
}

func (m *AdminUserInfoArg) Reset()                    { *m = AdminUserInfoArg{} }
func (m *AdminUserInfoArg) String() string            { return proto.CompactTextString(m) }
func (*AdminUserInfoArg) ProtoMessage()               {}
func (*AdminUserInfoArg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *AdminUserInfoArg) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AdminUserInfoArg) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AdminUserInfoArg) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AdminUserInfoArg) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AdminUserInfoArg) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

func (m *AdminUserInfoArg) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AdminUserInfoArg) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *AdminUserInfoArg) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *AdminUserInfoArg) GetCreator() int64 {
	if m != nil {
		return m.Creator
	}
	return 0
}

func (m *AdminUserInfoArg) GetRegisterIp() string {
	if m != nil {
		return m.RegisterIp
	}
	return ""
}

type AdminUserList struct {
	UserList []*AdminUserInfoReply `protobuf:"bytes,1,rep,name=UserList" json:"UserList"`
}

func (m *AdminUserList) Reset()                    { *m = AdminUserList{} }
func (m *AdminUserList) String() string            { return proto.CompactTextString(m) }
func (*AdminUserList) ProtoMessage()               {}
func (*AdminUserList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *AdminUserList) GetUserList() []*AdminUserInfoReply {
	if m != nil {
		return m.UserList
	}
	return nil
}

type AdminUserInfoReply struct {
	Id         int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	Email      string `protobuf:"bytes,2,opt,name=email" json:"email"`
	Username   string `protobuf:"bytes,3,opt,name=username" json:"username"`
	Password   string `protobuf:"bytes,4,opt,name=password" json:"password"`
	Salt       string `protobuf:"bytes,5,opt,name=salt" json:"salt"`
	Status     string `protobuf:"bytes,6,opt,name=status" json:"status"`
	Mobile     string `protobuf:"bytes,7,opt,name=mobile" json:"mobile"`
	CreateTime string `protobuf:"bytes,8,opt,name=create_time,json=createTime" json:"create_time"`
	Creator    string `protobuf:"bytes,9,opt,name=creator" json:"creator"`
	RegisterIp string `protobuf:"bytes,10,opt,name=register_ip,json=registerIp" json:"register_ip"`
}

func (m *AdminUserInfoReply) Reset()                    { *m = AdminUserInfoReply{} }
func (m *AdminUserInfoReply) String() string            { return proto.CompactTextString(m) }
func (*AdminUserInfoReply) ProtoMessage()               {}
func (*AdminUserInfoReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *AdminUserInfoReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AdminUserInfoReply) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AdminUserInfoReply) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AdminUserInfoReply) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AdminUserInfoReply) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

func (m *AdminUserInfoReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *AdminUserInfoReply) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *AdminUserInfoReply) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *AdminUserInfoReply) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *AdminUserInfoReply) GetRegisterIp() string {
	if m != nil {
		return m.RegisterIp
	}
	return ""
}

type AdminPrivilegesReply struct {
	Id         int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name"`
	Key        string `protobuf:"bytes,3,opt,name=key" json:"key"`
	Path       string `protobuf:"bytes,4,opt,name=path" json:"path"`
	Creator    string `protobuf:"bytes,5,opt,name=creator" json:"creator"`
	CreateTime string `protobuf:"bytes,6,opt,name=create_time,json=createTime" json:"create_time"`
}

func (m *AdminPrivilegesReply) Reset()                    { *m = AdminPrivilegesReply{} }
func (m *AdminPrivilegesReply) String() string            { return proto.CompactTextString(m) }
func (*AdminPrivilegesReply) ProtoMessage()               {}
func (*AdminPrivilegesReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *AdminPrivilegesReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AdminPrivilegesReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AdminPrivilegesReply) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AdminPrivilegesReply) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *AdminPrivilegesReply) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *AdminPrivilegesReply) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

type AdminPrivilegesList struct {
	PrivilegesList []*AdminPrivilegesReply `protobuf:"bytes,1,rep,name=PrivilegesList" json:"PrivilegesList"`
}

func (m *AdminPrivilegesList) Reset()                    { *m = AdminPrivilegesList{} }
func (m *AdminPrivilegesList) String() string            { return proto.CompactTextString(m) }
func (*AdminPrivilegesList) ProtoMessage()               {}
func (*AdminPrivilegesList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *AdminPrivilegesList) GetPrivilegesList() []*AdminPrivilegesReply {
	if m != nil {
		return m.PrivilegesList
	}
	return nil
}

type AdminPrivileges struct {
	Id         int64  `protobuf:"varint,1,opt,name=id" json:"id"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name"`
	Key        string `protobuf:"bytes,3,opt,name=key" json:"key"`
	Path       string `protobuf:"bytes,4,opt,name=path" json:"path"`
	Creator    int64  `protobuf:"varint,5,opt,name=creator" json:"creator"`
	CreateTime int64  `protobuf:"varint,6,opt,name=create_time,json=createTime" json:"create_time"`
}

func (m *AdminPrivileges) Reset()                    { *m = AdminPrivileges{} }
func (m *AdminPrivileges) String() string            { return proto.CompactTextString(m) }
func (*AdminPrivileges) ProtoMessage()               {}
func (*AdminPrivileges) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *AdminPrivileges) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AdminPrivileges) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AdminPrivileges) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AdminPrivileges) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *AdminPrivileges) GetCreator() int64 {
	if m != nil {
		return m.Creator
	}
	return 0
}

func (m *AdminPrivileges) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func init() {
	proto.RegisterType((*AdminUserInfoArg)(nil), "dbproto.AdminUserInfoArg")
	proto.RegisterType((*AdminUserList)(nil), "dbproto.AdminUserList")
	proto.RegisterType((*AdminUserInfoReply)(nil), "dbproto.AdminUserInfoReply")
	proto.RegisterType((*AdminPrivilegesReply)(nil), "dbproto.AdminPrivilegesReply")
	proto.RegisterType((*AdminPrivilegesList)(nil), "dbproto.AdminPrivilegesList")
	proto.RegisterType((*AdminPrivileges)(nil), "dbproto.AdminPrivileges")
}

func init() { proto.RegisterFile("admin.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0xd2, 0xa6, 0xcd, 0x14, 0x6b, 0x59, 0x8b, 0x2c, 0x8a, 0x58, 0x72, 0xea, 0xa9,
	0x07, 0x3d, 0x78, 0xee, 0x41, 0xb0, 0xe0, 0x41, 0x82, 0xde, 0x84, 0xb2, 0x35, 0x63, 0x5d, 0x4c,
	0x9a, 0xb0, 0xbb, 0x55, 0xfa, 0x2f, 0xc4, 0x83, 0x17, 0xff, 0xac, 0x64, 0x92, 0x94, 0x26, 0x95,
	0x82, 0x82, 0x78, 0x7b, 0x6f, 0x66, 0xb6, 0x9d, 0xf7, 0x31, 0x81, 0x8e, 0x08, 0x63, 0xb9, 0x18,
	0xa5, 0x2a, 0x31, 0x09, 0x6b, 0x85, 0x33, 0x12, 0xfe, 0x9b, 0x0d, 0xbd, 0x71, 0xd6, 0xb8, 0xd3,
	0xa8, 0x26, 0x8b, 0xc7, 0x64, 0xac, 0xe6, 0xac, 0x0b, 0xb6, 0x0c, 0xb9, 0x35, 0xb0, 0x86, 0x4e,
	0x60, 0xcb, 0x90, 0xf5, 0xa1, 0x89, 0xb1, 0x90, 0x11, 0xb7, 0x07, 0xd6, 0xd0, 0x0b, 0x72, 0xc3,
	0x8e, 0xa0, 0xbd, 0xd4, 0xa8, 0x16, 0x22, 0x46, 0xee, 0x50, 0x63, 0xed, 0xb3, 0x5e, 0x2a, 0xb4,
	0x7e, 0x4d, 0x54, 0xc8, 0x1b, 0x79, 0xaf, 0xf4, 0x8c, 0x41, 0x43, 0x8b, 0xc8, 0xf0, 0x26, 0xd5,
	0x49, 0xb3, 0x43, 0x70, 0xb5, 0x11, 0x66, 0xa9, 0xb9, 0x4b, 0xff, 0x5a, 0xb8, 0xac, 0x1e, 0x27,
	0x33, 0x19, 0x21, 0x6f, 0xd1, 0x74, 0xe1, 0xd8, 0x29, 0x74, 0x1e, 0x14, 0x0a, 0x83, 0x53, 0x23,
	0x63, 0xe4, 0x6d, 0x7a, 0x04, 0x79, 0xe9, 0x56, 0xc6, 0xc8, 0x38, 0xb4, 0xc8, 0x25, 0x8a, 0x7b,
	0xd4, 0x2c, 0x6d, 0xf6, 0x54, 0xe1, 0x5c, 0x6a, 0x83, 0x6a, 0x2a, 0x53, 0x0e, 0xf4, 0xbb, 0x50,
	0x96, 0x26, 0xa9, 0x7f, 0x05, 0x7b, 0x6b, 0x22, 0xd7, 0x52, 0x1b, 0x76, 0x01, 0xed, 0x52, 0x73,
	0x6b, 0xe0, 0x0c, 0x3b, 0x67, 0xc7, 0xa3, 0x82, 0xdf, 0xa8, 0xc2, 0x2e, 0xc0, 0x34, 0x5a, 0x05,
	0xeb, 0x61, 0xff, 0xdd, 0x06, 0xb6, 0x3d, 0xf0, 0xef, 0x78, 0xbd, 0xdf, 0xe0, 0xf5, 0x76, 0xe1,
	0xf5, 0x7e, 0x80, 0xf7, 0xd3, 0x82, 0x3e, 0x41, 0xb9, 0x51, 0xf2, 0x45, 0x46, 0x38, 0x47, 0xfd,
	0x3d, 0x16, 0x06, 0x0d, 0x0a, 0x9f, 0x53, 0x21, 0xcd, 0x7a, 0xe0, 0x3c, 0xe3, 0xaa, 0xe0, 0x91,
	0xc9, 0x6c, 0x2a, 0x15, 0xe6, 0xa9, 0xc0, 0x40, 0x7a, 0x73, 0xbb, 0xe6, 0xd6, 0x76, 0x9b, 0xc1,
	0xdc, 0x7a, 0x30, 0xff, 0x1e, 0x0e, 0x6a, 0xcb, 0xd1, 0x09, 0x5c, 0x42, 0xb7, 0x5a, 0x29, 0x0e,
	0xe1, 0xa4, 0x7a, 0x08, 0xb5, 0x48, 0x41, 0xed, 0x91, 0xff, 0x61, 0xc1, 0x7e, 0x6d, 0xf0, 0xef,
	0x62, 0x3b, 0x3b, 0x63, 0x57, 0x3e, 0x97, 0x99, 0x4b, 0x21, 0xce, 0xbf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xce, 0xaa, 0x12, 0x06, 0x25, 0x04, 0x00, 0x00,
}
