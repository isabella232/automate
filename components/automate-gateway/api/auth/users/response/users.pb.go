// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/auth/users/response/users.proto

package response // import "github.com/chef/automate/components/automate-gateway/api/auth/users/response"

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"` // Deprecated: Do not use.
	Username             string   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_3288f3e102f3a7e6, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Deprecated: Do not use.
func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_3288f3e102f3a7e6, []int{1}
}
func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (dst *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(dst, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type DeleteUserResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserResp) Reset()         { *m = DeleteUserResp{} }
func (m *DeleteUserResp) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResp) ProtoMessage()    {}
func (*DeleteUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_3288f3e102f3a7e6, []int{2}
}
func (m *DeleteUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResp.Unmarshal(m, b)
}
func (m *DeleteUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResp.Marshal(b, m, deterministic)
}
func (dst *DeleteUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResp.Merge(dst, src)
}
func (m *DeleteUserResp) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResp.Size(m)
}
func (m *DeleteUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "chef.automate.api.users.response.User")
	proto.RegisterType((*Users)(nil), "chef.automate.api.users.response.Users")
	proto.RegisterType((*DeleteUserResp)(nil), "chef.automate.api.users.response.DeleteUserResp")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/auth/users/response/users.proto", fileDescriptor_users_3288f3e102f3a7e6)
}

var fileDescriptor_users_3288f3e102f3a7e6 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc6, 0x69, 0xaf, 0x15, 0x7d, 0xc2, 0x21, 0x99, 0x82, 0x53, 0xe9, 0x20, 0xb7, 0x98, 0x80,
	0xae, 0x82, 0x70, 0xe8, 0x26, 0x0e, 0x05, 0x17, 0x27, 0xdf, 0xf5, 0x9e, 0xd7, 0xc0, 0xa5, 0x09,
	0x7d, 0x29, 0xe2, 0x7f, 0x2f, 0x49, 0xb8, 0xae, 0xc2, 0x6d, 0x79, 0xbf, 0xe4, 0xf7, 0xbe, 0xf0,
	0xc1, 0x73, 0xef, 0xac, 0x77, 0x23, 0x8d, 0x81, 0x35, 0xce, 0xc1, 0x59, 0x0c, 0x74, 0x7f, 0xc0,
	0x40, 0x3f, 0xf8, 0xab, 0xd1, 0x9b, 0x08, 0x07, 0x3d, 0x33, 0x4d, 0xac, 0x27, 0x62, 0xef, 0x46,
	0xa6, 0x3c, 0x2a, 0x3f, 0xb9, 0xe0, 0x44, 0xd3, 0x0f, 0xf4, 0xad, 0x4e, 0xaa, 0x42, 0x6f, 0x54,
	0xbe, 0x3e, 0xbd, 0x6e, 0xbf, 0xa0, 0xfa, 0x60, 0x9a, 0xc4, 0x1a, 0x4a, 0xb3, 0x97, 0x45, 0x53,
	0x6c, 0xae, 0xba, 0xd2, 0xec, 0x85, 0x80, 0x6a, 0x44, 0x4b, 0xb2, 0x4c, 0x24, 0x9d, 0x85, 0x84,
	0x9a, 0x2c, 0x9a, 0xa3, 0x5c, 0x45, 0xb8, 0x2d, 0x65, 0xd1, 0x65, 0x20, 0x6e, 0xe1, 0x32, 0xee,
	0x4d, 0x46, 0x95, 0x8c, 0x65, 0x6e, 0x5f, 0xa1, 0x8e, 0x09, 0x2c, 0x9e, 0xa0, 0x4e, 0xe1, 0xb2,
	0x68, 0x56, 0x9b, 0xeb, 0x87, 0x3b, 0xf5, 0xdf, 0xe7, 0x54, 0xf4, 0xba, 0x2c, 0xb5, 0x37, 0xb0,
	0x7e, 0xa1, 0x23, 0x05, 0x4a, 0x90, 0xd8, 0x6f, 0xdf, 0x3f, 0xdf, 0x0e, 0x26, 0x0c, 0xf3, 0x4e,
	0xf5, 0xce, 0xea, 0xb8, 0x6c, 0x29, 0x49, 0x9f, 0x51, 0xdc, 0xee, 0x22, 0x75, 0xf6, 0xf8, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x94, 0x9b, 0x7e, 0xd3, 0x76, 0x01, 0x00, 0x00,
}
