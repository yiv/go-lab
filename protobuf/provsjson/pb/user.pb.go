// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	UserInfo
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserInfo struct {
	Uid         int64  `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Unionid     string `protobuf:"bytes,2,opt,name=unionid" json:"unionid,omitempty"`
	Uuid        string `protobuf:"bytes,3,opt,name=uuid" json:"uuid,omitempty"`
	Username    string `protobuf:"bytes,4,opt,name=username" json:"username,omitempty"`
	Password    string `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
	Nick        string `protobuf:"bytes,6,opt,name=nick" json:"nick,omitempty"`
	Gender      bool   `protobuf:"varint,7,opt,name=gender" json:"gender,omitempty"`
	Addr        string `protobuf:"bytes,8,opt,name=addr" json:"addr,omitempty"`
	Avatar      string `protobuf:"bytes,9,opt,name=avatar" json:"avatar,omitempty"`
	Isguest     bool   `protobuf:"varint,10,opt,name=isguest" json:"isguest,omitempty"`
	Condays     int32  `protobuf:"varint,11,opt,name=condays" json:"condays,omitempty"`
	Signdate    int64  `protobuf:"varint,12,opt,name=signdate" json:"signdate,omitempty"`
	Vipsigndate int64  `protobuf:"varint,13,opt,name=vipsigndate" json:"vipsigndate,omitempty"`
	Status      bool   `protobuf:"varint,14,opt,name=status" json:"status,omitempty"`
	Mtime       int64  `protobuf:"varint,15,opt,name=mtime" json:"mtime,omitempty"`
	Ctime       int64  `protobuf:"varint,16,opt,name=ctime" json:"ctime,omitempty"`
	Token       string `protobuf:"bytes,17,opt,name=token" json:"token,omitempty"`
	Bankpwd     string `protobuf:"bytes,18,opt,name=bankpwd" json:"bankpwd,omitempty"`
	Forbid      string `protobuf:"bytes,19,opt,name=forbid" json:"forbid,omitempty"`
	Imsi        string `protobuf:"bytes,20,opt,name=imsi" json:"imsi,omitempty"`
	Imei        string `protobuf:"bytes,21,opt,name=imei" json:"imei,omitempty"`
	Mac         string `protobuf:"bytes,22,opt,name=mac" json:"mac,omitempty"`
	Did         string `protobuf:"bytes,23,opt,name=did" json:"did,omitempty"`
	Psystem     string `protobuf:"bytes,24,opt,name=psystem" json:"psystem,omitempty"`
	Pmodel      string `protobuf:"bytes,25,opt,name=pmodel" json:"pmodel,omitempty"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserInfo) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UserInfo) GetUnionid() string {
	if m != nil {
		return m.Unionid
	}
	return ""
}

func (m *UserInfo) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *UserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserInfo) GetNick() string {
	if m != nil {
		return m.Nick
	}
	return ""
}

func (m *UserInfo) GetGender() bool {
	if m != nil {
		return m.Gender
	}
	return false
}

func (m *UserInfo) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *UserInfo) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UserInfo) GetIsguest() bool {
	if m != nil {
		return m.Isguest
	}
	return false
}

func (m *UserInfo) GetCondays() int32 {
	if m != nil {
		return m.Condays
	}
	return 0
}

func (m *UserInfo) GetSigndate() int64 {
	if m != nil {
		return m.Signdate
	}
	return 0
}

func (m *UserInfo) GetVipsigndate() int64 {
	if m != nil {
		return m.Vipsigndate
	}
	return 0
}

func (m *UserInfo) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *UserInfo) GetMtime() int64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *UserInfo) GetCtime() int64 {
	if m != nil {
		return m.Ctime
	}
	return 0
}

func (m *UserInfo) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UserInfo) GetBankpwd() string {
	if m != nil {
		return m.Bankpwd
	}
	return ""
}

func (m *UserInfo) GetForbid() string {
	if m != nil {
		return m.Forbid
	}
	return ""
}

func (m *UserInfo) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *UserInfo) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *UserInfo) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *UserInfo) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *UserInfo) GetPsystem() string {
	if m != nil {
		return m.Psystem
	}
	return ""
}

func (m *UserInfo) GetPmodel() string {
	if m != nil {
		return m.Pmodel
	}
	return ""
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "pb.UserInfo")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0x4b, 0x6e, 0xe3, 0x30,
	0x0c, 0x40, 0xe1, 0x7c, 0x1d, 0x65, 0x3e, 0x19, 0x4d, 0x9a, 0xb2, 0x5d, 0x19, 0x5d, 0x79, 0xd5,
	0x4d, 0x4f, 0xd1, 0xad, 0x81, 0x1e, 0x40, 0xb6, 0x94, 0x80, 0x48, 0x2d, 0x09, 0x92, 0x9c, 0x20,
	0x17, 0xed, 0x79, 0x0a, 0x8a, 0x76, 0xd0, 0x1d, 0xdf, 0x23, 0x4d, 0xd2, 0x92, 0x84, 0x18, 0xa2,
	0x09, 0xaf, 0x3e, 0xb8, 0xe4, 0xe4, 0xcc, 0xb7, 0x2f, 0x5f, 0x0b, 0x51, 0x7e, 0x44, 0x13, 0xde,
	0xed, 0xd1, 0xc9, 0x9d, 0x98, 0x0f, 0xa8, 0xa1, 0xa8, 0x8a, 0x7a, 0xde, 0x50, 0x28, 0x41, 0xac,
	0x07, 0x8b, 0xce, 0xa2, 0x86, 0x59, 0x55, 0xd4, 0x9b, 0x66, 0x42, 0x29, 0xc5, 0x62, 0xa0, 0xe2,
	0x79, 0xd6, 0x39, 0x96, 0xcf, 0xa2, 0xa4, 0xf6, 0x56, 0xf5, 0x06, 0x16, 0xd9, 0xdf, 0x99, 0x72,
	0x5e, 0xc5, 0x78, 0x75, 0x41, 0xc3, 0x92, 0x73, 0x13, 0x53, 0x2f, 0x8b, 0xdd, 0x19, 0x56, 0xdc,
	0x8b, 0x62, 0x79, 0x10, 0xab, 0x93, 0xb1, 0xda, 0x04, 0x58, 0x57, 0x45, 0x5d, 0x36, 0x23, 0x51,
	0xad, 0xd2, 0x3a, 0x40, 0xc9, 0xb5, 0x14, 0x53, 0xad, 0xba, 0xa8, 0xa4, 0x02, 0x6c, 0xb2, 0x1d,
	0x89, 0xb6, 0xc7, 0x78, 0x1a, 0x4c, 0x4c, 0x20, 0x72, 0x93, 0x09, 0x29, 0xd3, 0x39, 0xab, 0xd5,
	0x2d, 0xc2, 0xb6, 0x2a, 0xea, 0x65, 0x33, 0x21, 0xed, 0x19, 0xf1, 0x64, 0xb5, 0x4a, 0x06, 0x7e,
	0xe5, 0x83, 0xb8, 0xb3, 0xac, 0xc4, 0xf6, 0x82, 0xfe, 0x9e, 0xfe, 0x9d, 0xd3, 0x3f, 0x15, 0x6d,
	0x12, 0x93, 0x4a, 0x43, 0x84, 0x3f, 0xbc, 0x35, 0x93, 0xdc, 0x8b, 0x65, 0x9f, 0xb0, 0x37, 0xf0,
	0x37, 0x7f, 0xc3, 0x40, 0xb6, 0xcb, 0x76, 0xc7, 0xb6, 0x9b, 0x6c, 0x72, 0x67, 0x63, 0xe1, 0x5f,
	0xfe, 0x19, 0x06, 0xda, 0xb8, 0x55, 0xf6, 0xec, 0xaf, 0x1a, 0x24, 0xdf, 0xc4, 0x88, 0x34, 0xf3,
	0xe8, 0x42, 0x8b, 0x1a, 0xfe, 0xf3, 0xdf, 0x33, 0xd1, 0x49, 0x61, 0x1f, 0x11, 0xf6, 0x7c, 0x52,
	0x14, 0xb3, 0x33, 0x08, 0x0f, 0x93, 0x33, 0x48, 0xb7, 0xde, 0xab, 0x0e, 0x0e, 0x59, 0x51, 0x48,
	0x46, 0xa3, 0x86, 0x47, 0x36, 0x9a, 0xdf, 0x81, 0x8f, 0xb7, 0x98, 0x4c, 0x0f, 0xc0, 0xd3, 0x47,
	0xa4, 0xe9, 0xbe, 0x77, 0xda, 0x7c, 0xc2, 0x13, 0x4f, 0x67, 0x6a, 0x57, 0xf9, 0x8d, 0xbd, 0x7d,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x77, 0x37, 0x5e, 0x72, 0x71, 0x02, 0x00, 0x00,
}