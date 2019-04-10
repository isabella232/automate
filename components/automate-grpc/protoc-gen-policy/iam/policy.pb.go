// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-grpc/protoc-gen-policy/iam/policy.proto

package iam // import "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"

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

type PolicyInfo struct {
	Resource             string   `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Action               string   `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyInfo) Reset()         { *m = PolicyInfo{} }
func (m *PolicyInfo) String() string { return proto.CompactTextString(m) }
func (*PolicyInfo) ProtoMessage()    {}
func (*PolicyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_02ec8574f8e25b14, []int{0}
}
func (m *PolicyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyInfo.Unmarshal(m, b)
}
func (m *PolicyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyInfo.Marshal(b, m, deterministic)
}
func (dst *PolicyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyInfo.Merge(dst, src)
}
func (m *PolicyInfo) XXX_Size() int {
	return xxx_messageInfo_PolicyInfo.Size(m)
}
func (m *PolicyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyInfo proto.InternalMessageInfo

func (m *PolicyInfo) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *PolicyInfo) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func init() {
	proto.RegisterType((*PolicyInfo)(nil), "chef.automate.api.iam.PolicyInfo")
}

func init() {
	proto.RegisterFile("components/automate-grpc/protoc-gen-policy/iam/policy.proto", fileDescriptor_policy_02ec8574f8e25b14)
}

var fileDescriptor_policy_02ec8574f8e25b14 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4e, 0xce, 0xcf, 0x2d,
	0xc8, 0xcf, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x4f, 0x2c, 0x2d, 0xc9, 0xcf, 0x4d, 0x2c, 0x49, 0xd5,
	0x4d, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xd6, 0x4d, 0x4f, 0xcd, 0xd3,
	0x2d, 0xc8, 0xcf, 0xc9, 0x4c, 0xae, 0xd4, 0xcf, 0x4c, 0xcc, 0xd5, 0x87, 0x30, 0xf5, 0xc0, 0x92,
	0x42, 0xa2, 0xc9, 0x19, 0xa9, 0x69, 0x7a, 0x30, 0x6d, 0x7a, 0x89, 0x05, 0x99, 0x7a, 0x99, 0x89,
	0xb9, 0x4a, 0x0e, 0x5c, 0x5c, 0x01, 0x60, 0x65, 0x9e, 0x79, 0x69, 0xf9, 0x42, 0x52, 0x5c, 0x1c,
	0x45, 0xa9, 0xc5, 0xf9, 0xa5, 0x45, 0xc9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70,
	0xbe, 0x90, 0x18, 0x17, 0x5b, 0x62, 0x72, 0x49, 0x66, 0x7e, 0x9e, 0x04, 0x13, 0x58, 0x06, 0xca,
	0x73, 0x72, 0x8e, 0x72, 0x4c, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07,
	0xd9, 0xa2, 0x9f, 0x68, 0xa4, 0x4f, 0x9a, 0x53, 0x93, 0xd8, 0xc0, 0xc2, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x3c, 0x81, 0xd4, 0x1e, 0xe3, 0x00, 0x00, 0x00,
}
