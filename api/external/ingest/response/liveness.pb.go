// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/ingest/response/liveness.proto

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

type ProcessLivenessResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessLivenessResponse) Reset()         { *m = ProcessLivenessResponse{} }
func (m *ProcessLivenessResponse) String() string { return proto.CompactTextString(m) }
func (*ProcessLivenessResponse) ProtoMessage()    {}
func (*ProcessLivenessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_liveness_b480753225715d17, []int{0}
}
func (m *ProcessLivenessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessLivenessResponse.Unmarshal(m, b)
}
func (m *ProcessLivenessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessLivenessResponse.Marshal(b, m, deterministic)
}
func (dst *ProcessLivenessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessLivenessResponse.Merge(dst, src)
}
func (m *ProcessLivenessResponse) XXX_Size() int {
	return xxx_messageInfo_ProcessLivenessResponse.Size(m)
}
func (m *ProcessLivenessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessLivenessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessLivenessResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProcessLivenessResponse)(nil), "chef.automate.api.ingest.response.ProcessLivenessResponse")
}

func init() {
	proto.RegisterFile("api/external/ingest/response/liveness.proto", fileDescriptor_liveness_b480753225715d17)
}

var fileDescriptor_liveness_b480753225715d17 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0x2c, 0xc8, 0xd4,
	0x4f, 0xad, 0x28, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0xcf, 0xcc, 0x4b, 0x4f, 0x2d, 0x2e, 0xd1,
	0x2f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0xcf, 0xc9, 0x2c, 0x4b, 0xcd, 0x4b, 0x2d,
	0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x4c, 0xce, 0x48, 0x4d, 0xd3, 0x4b, 0x2c,
	0x2d, 0xc9, 0xcf, 0x4d, 0x2c, 0x49, 0xd5, 0x4b, 0x2c, 0xc8, 0xd4, 0x83, 0xe8, 0xd0, 0x83, 0xe9,
	0x50, 0x92, 0xe4, 0x12, 0x0f, 0x28, 0xca, 0x4f, 0x4e, 0x2d, 0x2e, 0xf6, 0x81, 0xea, 0x0d, 0x82,
	0x4a, 0x39, 0x99, 0x47, 0x99, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea,
	0x83, 0x8c, 0xd2, 0x87, 0x19, 0xa5, 0x8f, 0xcf, 0x15, 0x49, 0x6c, 0x60, 0xdb, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x12, 0x90, 0xaf, 0xd7, 0xac, 0x00, 0x00, 0x00,
}
