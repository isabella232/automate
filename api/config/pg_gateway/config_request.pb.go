// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/config/pg_gateway/config_request.proto

package pg_gateway // import "github.com/chef/automate/api/config/pg_gateway"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import shared "github.com/chef/automate/api/config/shared"
import _ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConfigRequest struct {
	V1                   *ConfigRequest_V1 `protobuf:"bytes,1,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32             `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_1d1676e5f08e1cca, []int{0}
}
func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(dst, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if m != nil {
		return m.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	Sys                  *ConfigRequest_V1_System `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1) Reset()         { *m = ConfigRequest_V1{} }
func (m *ConfigRequest_V1) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1) ProtoMessage()    {}
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_1d1676e5f08e1cca, []int{0, 0}
}
func (m *ConfigRequest_V1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1.Unmarshal(m, b)
}
func (m *ConfigRequest_V1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1.Merge(dst, src)
}
func (m *ConfigRequest_V1) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1.Size(m)
}
func (m *ConfigRequest_V1) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1 proto.InternalMessageInfo

func (m *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if m != nil {
		return m.Sys
	}
	return nil
}

type ConfigRequest_V1_System struct {
	Mlsa                 *shared.Mlsa                      `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Service              *ConfigRequest_V1_System_Service  `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Tls                  *shared.TLSCredentials            `protobuf:"bytes,3,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Log                  *shared.Log                       `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	Timeouts             *ConfigRequest_V1_System_Timeouts `protobuf:"bytes,5,opt,name=timeouts,proto3" json:"timeouts,omitempty" toml:"timeouts,omitempty" mapstructure:"timeouts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                             `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_1d1676e5f08e1cca, []int{0, 0, 0}
}
func (m *ConfigRequest_V1_System) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System.Merge(dst, src)
}
func (m *ConfigRequest_V1_System) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System.Size(m)
}
func (m *ConfigRequest_V1_System) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System proto.InternalMessageInfo

func (m *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if m != nil {
		return m.Mlsa
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if m != nil {
		return m.Tls
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetLog() *shared.Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetTimeouts() *ConfigRequest_V1_System_Timeouts {
	if m != nil {
		return m.Timeouts
	}
	return nil
}

type ConfigRequest_V1_System_Service struct {
	Host                 *wrappers.StringValue       `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"`
	Port                 *wrappers.Int32Value        `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	ExternalPostgresql   *shared.External_Postgresql `protobuf:"bytes,4,opt,name=external_postgresql,json=externalPostgresql,proto3" json:"external_postgresql,omitempty" toml:"external_postgresql,omitempty" mapstructure:"external_postgresql,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                      `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                       `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Service) Reset()         { *m = ConfigRequest_V1_System_Service{} }
func (m *ConfigRequest_V1_System_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_1d1676e5f08e1cca, []int{0, 0, 0, 0}
}
func (m *ConfigRequest_V1_System_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Service.Merge(dst, src)
}
func (m *ConfigRequest_V1_System_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Size(m)
}
func (m *ConfigRequest_V1_System_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Service proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Service) GetHost() *wrappers.StringValue {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetPort() *wrappers.Int32Value {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetExternalPostgresql() *shared.External_Postgresql {
	if m != nil {
		return m.ExternalPostgresql
	}
	return nil
}

type ConfigRequest_V1_System_Timeouts struct {
	Connect              *wrappers.Int32Value `protobuf:"bytes,1,opt,name=connect,proto3" json:"connect,omitempty" toml:"connect,omitempty" mapstructure:"connect,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Timeouts) Reset()         { *m = ConfigRequest_V1_System_Timeouts{} }
func (m *ConfigRequest_V1_System_Timeouts) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Timeouts) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Timeouts) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_1d1676e5f08e1cca, []int{0, 0, 0, 1}
}
func (m *ConfigRequest_V1_System_Timeouts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Timeouts.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Timeouts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Timeouts.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System_Timeouts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Timeouts.Merge(dst, src)
}
func (m *ConfigRequest_V1_System_Timeouts) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Timeouts.Size(m)
}
func (m *ConfigRequest_V1_System_Timeouts) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Timeouts.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Timeouts proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Timeouts) GetConnect() *wrappers.Int32Value {
	if m != nil {
		return m.Connect
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.infra.pg_gateway.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.infra.pg_gateway.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.infra.pg_gateway.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Service)(nil), "chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Service")
	proto.RegisterType((*ConfigRequest_V1_System_Timeouts)(nil), "chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Timeouts")
}

func init() {
	proto.RegisterFile("api/config/pg_gateway/config_request.proto", fileDescriptor_config_request_1d1676e5f08e1cca)
}

var fileDescriptor_config_request_1d1676e5f08e1cca = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xe1, 0x6a, 0xdb, 0x30,
	0x10, 0x80, 0x69, 0xec, 0x26, 0x41, 0x63, 0xd0, 0xa9, 0x14, 0x3c, 0x77, 0x64, 0x65, 0xbf, 0x46,
	0x21, 0x72, 0x92, 0x6e, 0x0c, 0xc6, 0x60, 0x5d, 0xcb, 0x06, 0x2d, 0x1d, 0x1b, 0x4a, 0x09, 0x6c,
	0x0c, 0x32, 0xc5, 0x55, 0x14, 0x83, 0x2c, 0xb9, 0xd2, 0x39, 0x5d, 0x1e, 0x65, 0x3f, 0xf7, 0x2a,
	0x7d, 0x8c, 0xbd, 0x42, 0x7f, 0xf6, 0x05, 0x86, 0x6d, 0x39, 0xdd, 0x68, 0x09, 0xac, 0x3f, 0xa5,
	0xbb, 0xef, 0xd3, 0xdd, 0x59, 0x32, 0xda, 0x65, 0x59, 0x12, 0xc5, 0x5a, 0x4d, 0x13, 0x11, 0x65,
	0x62, 0x2c, 0x18, 0xf0, 0x0b, 0xb6, 0x70, 0x3b, 0x63, 0xc3, 0xcf, 0x73, 0x6e, 0x81, 0x64, 0x46,
	0x83, 0xc6, 0x9d, 0x78, 0xc6, 0xa7, 0x84, 0xe5, 0xa0, 0x53, 0x06, 0x9c, 0x24, 0x6a, 0x6a, 0x18,
	0xb9, 0x81, 0xc2, 0xce, 0x5f, 0x2e, 0x3b, 0x63, 0x86, 0x9f, 0x45, 0x42, 0xea, 0x09, 0x93, 0x15,
	0x1f, 0x6e, 0xdf, 0x8e, 0x83, 0xb4, 0x2e, 0x78, 0x1c, 0xeb, 0x34, 0xd3, 0x8a, 0x2b, 0xb0, 0x51,
	0x7d, 0x44, 0x57, 0x98, 0x2c, 0x8e, 0xca, 0x78, 0xdc, 0x15, 0x5c, 0x75, 0xd9, 0xa0, 0xeb, 0xf8,
	0x42, 0xc5, 0x06, 0xc5, 0x22, 0x62, 0x4a, 0x69, 0x60, 0x90, 0x68, 0x55, 0xbb, 0x3a, 0x42, 0x6b,
	0x21, 0x79, 0x45, 0x4e, 0xf2, 0x69, 0x74, 0x61, 0x58, 0x96, 0x71, 0xe3, 0xe2, 0xcf, 0x7e, 0x37,
	0xd1, 0xc3, 0xc3, 0xd2, 0x43, 0xab, 0x06, 0xf1, 0x3e, 0x6a, 0xcc, 0xfb, 0xc1, 0xda, 0xce, 0xda,
	0xf3, 0x07, 0x83, 0x1e, 0x59, 0xdd, 0x27, 0xf9, 0x07, 0x25, 0xa3, 0x3e, 0x6d, 0xcc, 0xfb, 0xe1,
	0xd5, 0x3a, 0x6a, 0x8c, 0xfa, 0xf8, 0x08, 0x79, 0x76, 0x61, 0x9d, 0xe9, 0xd5, 0xff, 0x9a, 0xc8,
	0x70, 0x61, 0x81, 0xa7, 0xb4, 0x70, 0x84, 0x3f, 0xd7, 0x51, 0xb3, 0x5a, 0xe3, 0x17, 0xc8, 0x4f,
	0xa5, 0x65, 0x4e, 0xbb, 0x73, 0xa7, 0xb6, 0x1a, 0x0c, 0xf9, 0x28, 0x2d, 0xa3, 0x65, 0x36, 0xfe,
	0x82, 0x5a, 0x96, 0x9b, 0x79, 0x12, 0xf3, 0xa0, 0x51, 0x82, 0x6f, 0xef, 0x59, 0x0f, 0x19, 0x56,
	0x1a, 0x5a, 0xfb, 0xf0, 0x1b, 0xe4, 0x81, 0xb4, 0x81, 0x57, 0x6a, 0x77, 0x57, 0xd5, 0x73, 0x7a,
	0x32, 0x3c, 0x34, 0xfc, 0x8c, 0x2b, 0x48, 0x98, 0xb4, 0xb4, 0xc0, 0x70, 0x1f, 0x79, 0x52, 0x8b,
	0xc0, 0x2f, 0xe9, 0xa7, 0xab, 0xe8, 0x13, 0x2d, 0x68, 0x91, 0x8b, 0xbf, 0xa1, 0x36, 0x24, 0x29,
	0xd7, 0x39, 0xd8, 0x60, 0xbd, 0xe4, 0xf6, 0xef, 0xdb, 0xcc, 0xa9, 0xf3, 0xd0, 0xa5, 0x31, 0xbc,
	0x5a, 0x43, 0x2d, 0xd7, 0x23, 0xee, 0x21, 0x7f, 0xa6, 0x2d, 0xb8, 0x59, 0x3f, 0x21, 0xd5, 0x5d,
	0x22, 0xf5, 0x5d, 0x22, 0x43, 0x30, 0x89, 0x12, 0x23, 0x26, 0x73, 0x4e, 0xcb, 0x4c, 0xfc, 0x01,
	0xf9, 0x99, 0x36, 0xe0, 0x86, 0xbc, 0x7d, 0x8b, 0x38, 0x52, 0xb0, 0x37, 0x28, 0x81, 0x83, 0xad,
	0xcb, 0xeb, 0xe0, 0xd1, 0xf2, 0xb3, 0x6c, 0xfc, 0xfa, 0x14, 0x7a, 0x10, 0x67, 0xb4, 0xe4, 0xf1,
	0x77, 0xb4, 0xc9, 0x7f, 0x00, 0x37, 0x8a, 0xc9, 0x71, 0xa6, 0x2d, 0x08, 0xc3, 0xed, 0xb9, 0x74,
	0x63, 0x8a, 0x56, 0x8d, 0xe9, 0xbd, 0xc3, 0xc8, 0xe7, 0x25, 0x46, 0x71, 0xed, 0xba, 0xd9, 0x3b,
	0xf6, 0xdb, 0xde, 0x86, 0x1f, 0xbe, 0x43, 0xed, 0x7a, 0x06, 0xf8, 0x25, 0x6a, 0xc5, 0x5a, 0x29,
	0x1e, 0xd7, 0x0d, 0xaf, 0x2a, 0x9f, 0xd6, 0xb9, 0xaf, 0x1f, 0x5f, 0x5e, 0x07, 0x5b, 0x68, 0x73,
	0xf9, 0x50, 0x33, 0xd1, 0x75, 0x63, 0x3f, 0xe8, 0x7d, 0x25, 0x22, 0x81, 0x59, 0x3e, 0x21, 0xb1,
	0x4e, 0xa3, 0xa2, 0xe8, 0xe5, 0x7b, 0x8e, 0xee, 0xfc, 0xd9, 0x4c, 0x9a, 0xe5, 0x51, 0x7b, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x01, 0x77, 0x2d, 0x03, 0x8c, 0x04, 0x00, 0x00,
}
