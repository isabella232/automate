// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/response/tokens.proto

package response // import "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/response"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateTokenResp struct {
	Token                *common.Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateTokenResp) Reset()         { *m = CreateTokenResp{} }
func (m *CreateTokenResp) String() string { return proto.CompactTextString(m) }
func (*CreateTokenResp) ProtoMessage()    {}
func (*CreateTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_tokens_19ce22a4e22cc0d4, []int{0}
}
func (m *CreateTokenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTokenResp.Unmarshal(m, b)
}
func (m *CreateTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTokenResp.Marshal(b, m, deterministic)
}
func (dst *CreateTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTokenResp.Merge(dst, src)
}
func (m *CreateTokenResp) XXX_Size() int {
	return xxx_messageInfo_CreateTokenResp.Size(m)
}
func (m *CreateTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTokenResp proto.InternalMessageInfo

func (m *CreateTokenResp) GetToken() *common.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type GetTokenResp struct {
	Token                *common.Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetTokenResp) Reset()         { *m = GetTokenResp{} }
func (m *GetTokenResp) String() string { return proto.CompactTextString(m) }
func (*GetTokenResp) ProtoMessage()    {}
func (*GetTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_tokens_19ce22a4e22cc0d4, []int{1}
}
func (m *GetTokenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenResp.Unmarshal(m, b)
}
func (m *GetTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenResp.Marshal(b, m, deterministic)
}
func (dst *GetTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenResp.Merge(dst, src)
}
func (m *GetTokenResp) XXX_Size() int {
	return xxx_messageInfo_GetTokenResp.Size(m)
}
func (m *GetTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenResp proto.InternalMessageInfo

func (m *GetTokenResp) GetToken() *common.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type UpdateTokenResp struct {
	Token                *common.Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateTokenResp) Reset()         { *m = UpdateTokenResp{} }
func (m *UpdateTokenResp) String() string { return proto.CompactTextString(m) }
func (*UpdateTokenResp) ProtoMessage()    {}
func (*UpdateTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_tokens_19ce22a4e22cc0d4, []int{2}
}
func (m *UpdateTokenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTokenResp.Unmarshal(m, b)
}
func (m *UpdateTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTokenResp.Marshal(b, m, deterministic)
}
func (dst *UpdateTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTokenResp.Merge(dst, src)
}
func (m *UpdateTokenResp) XXX_Size() int {
	return xxx_messageInfo_UpdateTokenResp.Size(m)
}
func (m *UpdateTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTokenResp proto.InternalMessageInfo

func (m *UpdateTokenResp) GetToken() *common.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type DeleteTokenResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTokenResp) Reset()         { *m = DeleteTokenResp{} }
func (m *DeleteTokenResp) String() string { return proto.CompactTextString(m) }
func (*DeleteTokenResp) ProtoMessage()    {}
func (*DeleteTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_tokens_19ce22a4e22cc0d4, []int{3}
}
func (m *DeleteTokenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTokenResp.Unmarshal(m, b)
}
func (m *DeleteTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTokenResp.Marshal(b, m, deterministic)
}
func (dst *DeleteTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTokenResp.Merge(dst, src)
}
func (m *DeleteTokenResp) XXX_Size() int {
	return xxx_messageInfo_DeleteTokenResp.Size(m)
}
func (m *DeleteTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTokenResp proto.InternalMessageInfo

type ListTokensResp struct {
	Tokens               []*common.Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListTokensResp) Reset()         { *m = ListTokensResp{} }
func (m *ListTokensResp) String() string { return proto.CompactTextString(m) }
func (*ListTokensResp) ProtoMessage()    {}
func (*ListTokensResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_tokens_19ce22a4e22cc0d4, []int{4}
}
func (m *ListTokensResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTokensResp.Unmarshal(m, b)
}
func (m *ListTokensResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTokensResp.Marshal(b, m, deterministic)
}
func (dst *ListTokensResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTokensResp.Merge(dst, src)
}
func (m *ListTokensResp) XXX_Size() int {
	return xxx_messageInfo_ListTokensResp.Size(m)
}
func (m *ListTokensResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTokensResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListTokensResp proto.InternalMessageInfo

func (m *ListTokensResp) GetTokens() []*common.Token {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTokenResp)(nil), "chef.automate.api.iam.v2beta.CreateTokenResp")
	proto.RegisterType((*GetTokenResp)(nil), "chef.automate.api.iam.v2beta.GetTokenResp")
	proto.RegisterType((*UpdateTokenResp)(nil), "chef.automate.api.iam.v2beta.UpdateTokenResp")
	proto.RegisterType((*DeleteTokenResp)(nil), "chef.automate.api.iam.v2beta.DeleteTokenResp")
	proto.RegisterType((*ListTokensResp)(nil), "chef.automate.api.iam.v2beta.ListTokensResp")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2beta/response/tokens.proto", fileDescriptor_tokens_19ce22a4e22cc0d4)
}

var fileDescriptor_tokens_19ce22a4e22cc0d4 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x31, 0x4b, 0x03, 0x41,
	0x10, 0x85, 0x39, 0xc4, 0x14, 0xa3, 0x78, 0x98, 0x2a, 0x88, 0x45, 0x38, 0x1b, 0x1b, 0x67, 0x20,
	0x56, 0x22, 0x88, 0xa8, 0x20, 0xc2, 0x69, 0x21, 0xda, 0xd8, 0xcd, 0x9d, 0x63, 0xb2, 0xe8, 0xee,
	0x2c, 0xb7, 0x13, 0xc5, 0x7f, 0x2f, 0xb9, 0x55, 0xd1, 0x46, 0x54, 0xd2, 0x2e, 0xfb, 0x7d, 0xf3,
	0x98, 0x37, 0x70, 0xdc, 0xaa, 0x8f, 0x1a, 0x24, 0x58, 0x22, 0x9e, 0x9b, 0x7a, 0x36, 0xd9, 0x9b,
	0xb2, 0xc9, 0x0b, 0xbf, 0x12, 0x47, 0x47, 0x8e, 0x3d, 0x3d, 0x4f, 0x1a, 0x31, 0xa6, 0x4e, 0x52,
	0xd4, 0x90, 0x84, 0x4c, 0x1f, 0x25, 0x24, 0x8c, 0x9d, 0x9a, 0x0e, 0xb7, 0xdb, 0x99, 0x3c, 0xe0,
	0x07, 0x8b, 0x1c, 0x1d, 0x3a, 0xf6, 0x98, 0x99, 0xad, 0xa3, 0x3f, 0xf8, 0x5b, 0xf5, 0x5e, 0xc3,
	0x37, 0x7b, 0x55, 0x43, 0x79, 0xda, 0x09, 0x9b, 0xdc, 0x2c, 0x5e, 0xaf, 0x25, 0xc5, 0xe1, 0x01,
	0xac, 0xf6, 0x5f, 0x46, 0xc5, 0xb8, 0xd8, 0x5d, 0x9b, 0xec, 0xe0, 0x4f, 0x01, 0x30, 0x73, 0x99,
	0xa8, 0x2e, 0x60, 0xfd, 0x5c, 0x6c, 0x29, 0xaa, 0x1a, 0xca, 0xdb, 0x78, 0xbf, 0xac, 0x60, 0x9b,
	0x50, 0x9e, 0xc9, 0x93, 0x7c, 0xb1, 0x55, 0x97, 0xb0, 0x51, 0xbb, 0x94, 0xc3, 0xa6, 0xde, 0x7f,
	0x08, 0x83, 0xbc, 0x9b, 0x51, 0x31, 0x5e, 0xf9, 0xed, 0x80, 0x77, 0xe4, 0xe4, 0xea, 0xae, 0x9e,
	0x3a, 0x9b, 0xcd, 0x1b, 0x6c, 0xd5, 0xd3, 0x02, 0xfc, 0xec, 0x83, 0xfe, 0x71, 0x03, 0xcd, 0xa0,
	0xef, 0x67, 0xff, 0x2d, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x74, 0xd1, 0x9f, 0x41, 0x02, 0x00, 0x00,
}
