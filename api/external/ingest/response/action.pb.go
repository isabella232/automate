// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/ingest/response/action.proto

package response // import "github.com/chef/automate/api/external/ingest/response"

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

type ProcessChefActionResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessChefActionResponse) Reset()         { *m = ProcessChefActionResponse{} }
func (m *ProcessChefActionResponse) String() string { return proto.CompactTextString(m) }
func (*ProcessChefActionResponse) ProtoMessage()    {}
func (*ProcessChefActionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_action_e00c23d4daf49e7d, []int{0}
}
func (m *ProcessChefActionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessChefActionResponse.Unmarshal(m, b)
}
func (m *ProcessChefActionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessChefActionResponse.Marshal(b, m, deterministic)
}
func (dst *ProcessChefActionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessChefActionResponse.Merge(dst, src)
}
func (m *ProcessChefActionResponse) XXX_Size() int {
	return xxx_messageInfo_ProcessChefActionResponse.Size(m)
}
func (m *ProcessChefActionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessChefActionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessChefActionResponse proto.InternalMessageInfo

type ProcessNodeDeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessNodeDeleteResponse) Reset()         { *m = ProcessNodeDeleteResponse{} }
func (m *ProcessNodeDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*ProcessNodeDeleteResponse) ProtoMessage()    {}
func (*ProcessNodeDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_action_e00c23d4daf49e7d, []int{1}
}
func (m *ProcessNodeDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessNodeDeleteResponse.Unmarshal(m, b)
}
func (m *ProcessNodeDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessNodeDeleteResponse.Marshal(b, m, deterministic)
}
func (dst *ProcessNodeDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessNodeDeleteResponse.Merge(dst, src)
}
func (m *ProcessNodeDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_ProcessNodeDeleteResponse.Size(m)
}
func (m *ProcessNodeDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessNodeDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessNodeDeleteResponse proto.InternalMessageInfo

type ProcessMultipleNodeDeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessMultipleNodeDeleteResponse) Reset()         { *m = ProcessMultipleNodeDeleteResponse{} }
func (m *ProcessMultipleNodeDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*ProcessMultipleNodeDeleteResponse) ProtoMessage()    {}
func (*ProcessMultipleNodeDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_action_e00c23d4daf49e7d, []int{2}
}
func (m *ProcessMultipleNodeDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessMultipleNodeDeleteResponse.Unmarshal(m, b)
}
func (m *ProcessMultipleNodeDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessMultipleNodeDeleteResponse.Marshal(b, m, deterministic)
}
func (dst *ProcessMultipleNodeDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessMultipleNodeDeleteResponse.Merge(dst, src)
}
func (m *ProcessMultipleNodeDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_ProcessMultipleNodeDeleteResponse.Size(m)
}
func (m *ProcessMultipleNodeDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessMultipleNodeDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessMultipleNodeDeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProcessChefActionResponse)(nil), "chef.automate.api.ingest.response.ProcessChefActionResponse")
	proto.RegisterType((*ProcessNodeDeleteResponse)(nil), "chef.automate.api.ingest.response.ProcessNodeDeleteResponse")
	proto.RegisterType((*ProcessMultipleNodeDeleteResponse)(nil), "chef.automate.api.ingest.response.ProcessMultipleNodeDeleteResponse")
}

func init() {
	proto.RegisterFile("api/external/ingest/response/action.proto", fileDescriptor_action_e00c23d4daf49e7d)
}

var fileDescriptor_action_e00c23d4daf49e7d = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xbf, 0xcb, 0xc2, 0x30,
	0x14, 0x45, 0xb7, 0x6f, 0xe8, 0xf8, 0x6d, 0xe2, 0x54, 0x9d, 0x5c, 0xf2, 0x06, 0x11, 0x67, 0x7f,
	0xac, 0x8a, 0x38, 0xba, 0xa5, 0xf1, 0xb6, 0x7d, 0x90, 0xe6, 0x85, 0xe4, 0x15, 0xfc, 0xf3, 0xc5,
	0x6a, 0x5d, 0x0a, 0xce, 0xe7, 0x5c, 0xb8, 0xa7, 0x58, 0xd9, 0xc8, 0x84, 0x87, 0x22, 0x05, 0xeb,
	0x89, 0x43, 0x83, 0xac, 0x94, 0x90, 0xa3, 0x84, 0x0c, 0xb2, 0x4e, 0x59, 0x82, 0x89, 0x49, 0x54,
	0xfe, 0x4b, 0xd7, 0xa2, 0x36, 0xb6, 0x57, 0xe9, 0xac, 0xc2, 0xd8, 0xc8, 0xe6, 0xed, 0x9b, 0xd1,
	0x5f, 0xcc, 0x8b, 0xd9, 0x25, 0x89, 0x43, 0xce, 0x87, 0x16, 0xf5, 0x6e, 0x58, 0x5f, 0xa7, 0xf0,
	0x2c, 0x77, 0x1c, 0xe1, 0xa1, 0xf8, 0xc2, 0x65, 0x51, 0x7e, 0xe0, 0xa9, 0xf7, 0xca, 0xd1, 0x63,
	0x2a, 0xed, 0xb7, 0xb7, 0x4d, 0xc3, 0xda, 0xf6, 0x95, 0x71, 0xd2, 0xd1, 0xeb, 0x0e, 0x8d, 0x77,
	0xe8, 0x57, 0x47, 0xf5, 0x37, 0x14, 0xac, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xa7, 0xa6,
	0xe9, 0xee, 0x00, 0x00, 0x00,
}
