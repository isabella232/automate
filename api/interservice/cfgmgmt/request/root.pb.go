// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/cfgmgmt/request/root.proto

package request // import "github.com/chef/automate/api/interservice/cfgmgmt/request"

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

type VersionInfo struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *VersionInfo) Reset()         { *m = VersionInfo{} }
func (m *VersionInfo) String() string { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()    {}
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_root_a0e1a0d188312b20, []int{0}
}
func (m *VersionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionInfo.Unmarshal(m, b)
}
func (m *VersionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionInfo.Marshal(b, m, deterministic)
}
func (dst *VersionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionInfo.Merge(dst, src)
}
func (m *VersionInfo) XXX_Size() int {
	return xxx_messageInfo_VersionInfo.Size(m)
}
func (m *VersionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VersionInfo proto.InternalMessageInfo

type Health struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Health) Reset()         { *m = Health{} }
func (m *Health) String() string { return proto.CompactTextString(m) }
func (*Health) ProtoMessage()    {}
func (*Health) Descriptor() ([]byte, []int) {
	return fileDescriptor_root_a0e1a0d188312b20, []int{1}
}
func (m *Health) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Health.Unmarshal(m, b)
}
func (m *Health) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Health.Marshal(b, m, deterministic)
}
func (dst *Health) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Health.Merge(dst, src)
}
func (m *Health) XXX_Size() int {
	return xxx_messageInfo_Health.Size(m)
}
func (m *Health) XXX_DiscardUnknown() {
	xxx_messageInfo_Health.DiscardUnknown(m)
}

var xxx_messageInfo_Health proto.InternalMessageInfo

type Organizations struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Organizations) Reset()         { *m = Organizations{} }
func (m *Organizations) String() string { return proto.CompactTextString(m) }
func (*Organizations) ProtoMessage()    {}
func (*Organizations) Descriptor() ([]byte, []int) {
	return fileDescriptor_root_a0e1a0d188312b20, []int{2}
}
func (m *Organizations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Organizations.Unmarshal(m, b)
}
func (m *Organizations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Organizations.Marshal(b, m, deterministic)
}
func (dst *Organizations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Organizations.Merge(dst, src)
}
func (m *Organizations) XXX_Size() int {
	return xxx_messageInfo_Organizations.Size(m)
}
func (m *Organizations) XXX_DiscardUnknown() {
	xxx_messageInfo_Organizations.DiscardUnknown(m)
}

var xxx_messageInfo_Organizations proto.InternalMessageInfo

type SourceFQDNS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *SourceFQDNS) Reset()         { *m = SourceFQDNS{} }
func (m *SourceFQDNS) String() string { return proto.CompactTextString(m) }
func (*SourceFQDNS) ProtoMessage()    {}
func (*SourceFQDNS) Descriptor() ([]byte, []int) {
	return fileDescriptor_root_a0e1a0d188312b20, []int{3}
}
func (m *SourceFQDNS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceFQDNS.Unmarshal(m, b)
}
func (m *SourceFQDNS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceFQDNS.Marshal(b, m, deterministic)
}
func (dst *SourceFQDNS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceFQDNS.Merge(dst, src)
}
func (m *SourceFQDNS) XXX_Size() int {
	return xxx_messageInfo_SourceFQDNS.Size(m)
}
func (m *SourceFQDNS) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceFQDNS.DiscardUnknown(m)
}

var xxx_messageInfo_SourceFQDNS proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VersionInfo)(nil), "chef.automate.domain.cfgmgmt.request.VersionInfo")
	proto.RegisterType((*Health)(nil), "chef.automate.domain.cfgmgmt.request.Health")
	proto.RegisterType((*Organizations)(nil), "chef.automate.domain.cfgmgmt.request.Organizations")
	proto.RegisterType((*SourceFQDNS)(nil), "chef.automate.domain.cfgmgmt.request.SourceFQDNS")
}

func init() {
	proto.RegisterFile("api/interservice/cfgmgmt/request/root.proto", fileDescriptor_root_a0e1a0d188312b20)
}

var fileDescriptor_root_a0e1a0d188312b20 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8e, 0xb1, 0x0a, 0xc2, 0x30,
	0x10, 0x40, 0xb7, 0x22, 0x91, 0x22, 0xf8, 0x09, 0xc5, 0x4d, 0x48, 0x06, 0x27, 0x71, 0x13, 0x11,
	0x5d, 0x14, 0x29, 0x38, 0xb8, 0xa5, 0xf1, 0x9a, 0x1e, 0x98, 0x5c, 0xbd, 0x5c, 0x1c, 0xfc, 0x7a,
	0xa9, 0xd4, 0xd9, 0xfd, 0x3d, 0xde, 0x53, 0x4b, 0xdb, 0xa3, 0xc1, 0x28, 0xc0, 0x09, 0xf8, 0x85,
	0x0e, 0x8c, 0x6b, 0x7d, 0xf0, 0x41, 0x0c, 0xc3, 0x33, 0x43, 0x12, 0xc3, 0x44, 0xa2, 0x7b, 0x26,
	0xa1, 0xf9, 0xc2, 0x75, 0xd0, 0x6a, 0x9b, 0x85, 0x82, 0x15, 0xd0, 0x77, 0x0a, 0x16, 0xa3, 0x1e,
	0x05, 0x3d, 0x0a, 0x55, 0xa9, 0xa6, 0x57, 0xe0, 0x84, 0x14, 0x8f, 0xb1, 0xa5, 0x6a, 0xa2, 0x8a,
	0x03, 0xd8, 0x87, 0x74, 0xd5, 0x4c, 0x95, 0x67, 0xf6, 0x36, 0xe2, 0xdb, 0x0a, 0x52, 0x4c, 0x03,
	0x59, 0x53, 0x66, 0x07, 0xfb, 0xcb, 0xee, 0x54, 0x6f, 0x37, 0xb7, 0xb5, 0x47, 0xe9, 0x72, 0xa3,
	0x1d, 0x05, 0x33, 0xb4, 0xcc, 0xaf, 0x65, 0xfe, 0x6d, 0x36, 0xc5, 0x77, 0x71, 0xf5, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x9e, 0xdd, 0x96, 0xfe, 0xd1, 0x00, 0x00, 0x00,
}
