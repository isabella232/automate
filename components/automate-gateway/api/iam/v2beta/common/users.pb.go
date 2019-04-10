// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/common/users.proto

package common // import "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/common"

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
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// TODO (tc): Will remove post-GA
	MembershipId         string   `protobuf:"bytes,3,opt,name=membership_id,json=membershipId,proto3" json:"membership_id,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_15db1476fcf1f29d, []int{0}
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

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Deprecated: Do not use.
func (m *User) GetMembershipId() string {
	if m != nil {
		return m.MembershipId
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "chef.automate.api.iam.v2beta.User")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2beta/common/users.proto", fileDescriptor_users_15db1476fcf1f29d)
}

var fileDescriptor_users_15db1476fcf1f29d = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0x31, 0x4b, 0x05, 0x31,
	0x0c, 0x80, 0xb9, 0xf3, 0x21, 0x58, 0xd4, 0xa1, 0x53, 0x07, 0x07, 0x71, 0xd1, 0xc5, 0x06, 0x74,
	0x76, 0x79, 0x9b, 0xe2, 0xf4, 0xc4, 0xc5, 0x45, 0xd2, 0x36, 0xbe, 0xcb, 0x90, 0xa6, 0xb4, 0x3d,
	0xc5, 0x7f, 0x2f, 0x77, 0xa2, 0xce, 0x6e, 0x21, 0x1f, 0x1f, 0xc9, 0x67, 0xee, 0xa2, 0x4a, 0xd1,
	0x4c, 0xb9, 0x37, 0xc0, 0xb9, 0xab, 0x60, 0xa7, 0xeb, 0x3d, 0x76, 0xfa, 0xc0, 0x4f, 0xc0, 0xc2,
	0xc0, 0x28, 0xf0, 0x7e, 0x13, 0xa8, 0x23, 0x44, 0x15, 0xd1, 0x0c, 0x73, 0xa3, 0xda, 0x7c, 0xa9,
	0xda, 0xd5, 0x9e, 0xc5, 0x89, 0xde, 0xfc, 0x8f, 0xe8, 0xb1, 0xb0, 0x67, 0x14, 0xff, 0x2d, 0x5c,
	0x3c, 0x99, 0xcd, 0x73, 0xa3, 0x6a, 0xad, 0xd9, 0x64, 0x14, 0x72, 0xc3, 0xf9, 0x70, 0x75, 0xb4,
	0x5b, 0x67, 0x7b, 0x6a, 0x46, 0x4e, 0x6e, 0x5c, 0x37, 0x23, 0x27, 0x7b, 0x69, 0x4e, 0x84, 0x24,
	0x50, 0x6d, 0x13, 0x97, 0x57, 0x4e, 0xee, 0x60, 0x41, 0xdb, 0xd1, 0x0d, 0xbb, 0xe3, 0x3f, 0x70,
	0x9f, 0xb6, 0x8f, 0x2f, 0x0f, 0x7b, 0xee, 0xd3, 0x1c, 0x7c, 0x54, 0x81, 0xe5, 0xfe, 0xef, 0xe3,
	0xf0, 0xef, 0x98, 0x70, 0xb8, 0x76, 0xdc, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x11, 0x3d, 0x30,
	0x6f, 0x08, 0x01, 0x00, 0x00,
}
