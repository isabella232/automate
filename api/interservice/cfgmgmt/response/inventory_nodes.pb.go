// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/cfgmgmt/response/inventory_nodes.proto

package response // import "github.com/chef/automate/api/interservice/cfgmgmt/response"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InventoryNodes struct {
	Nodes                []*InventoryNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty" toml:"nodes,omitempty" mapstructure:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32            `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *InventoryNodes) Reset()         { *m = InventoryNodes{} }
func (m *InventoryNodes) String() string { return proto.CompactTextString(m) }
func (*InventoryNodes) ProtoMessage()    {}
func (*InventoryNodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_inventory_nodes_659aa2ae1f39340d, []int{0}
}
func (m *InventoryNodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InventoryNodes.Unmarshal(m, b)
}
func (m *InventoryNodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InventoryNodes.Marshal(b, m, deterministic)
}
func (dst *InventoryNodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InventoryNodes.Merge(dst, src)
}
func (m *InventoryNodes) XXX_Size() int {
	return xxx_messageInfo_InventoryNodes.Size(m)
}
func (m *InventoryNodes) XXX_DiscardUnknown() {
	xxx_messageInfo_InventoryNodes.DiscardUnknown(m)
}

var xxx_messageInfo_InventoryNodes proto.InternalMessageInfo

func (m *InventoryNodes) GetNodes() []*InventoryNode {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type InventoryNode struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	Checkin              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=checkin,proto3" json:"checkin,omitempty" toml:"checkin,omitempty" mapstructure:"checkin,omitempty"`
	Organization         string               `protobuf:"bytes,3,opt,name=organization,proto3" json:"organization,omitempty" toml:"organization,omitempty" mapstructure:"organization,omitempty"`
	Platform             string               `protobuf:"bytes,4,opt,name=platform,proto3" json:"platform,omitempty" toml:"platform,omitempty" mapstructure:"platform,omitempty"`
	PlatformFamily       string               `protobuf:"bytes,5,opt,name=platform_family,json=platformFamily,proto3" json:"platform_family,omitempty" toml:"platform_family,omitempty" mapstructure:"platform_family,omitempty"`
	PlatformVersion      string               `protobuf:"bytes,6,opt,name=platform_version,json=platformVersion,proto3" json:"platform_version,omitempty" toml:"platform_version,omitempty" mapstructure:"platform_version,omitempty"`
	ClientVersion        string               `protobuf:"bytes,7,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty" toml:"client_version,omitempty" mapstructure:"client_version,omitempty"`
	Ec2InstanceId        string               `protobuf:"bytes,8,opt,name=ec2_instance_id,json=ec2InstanceId,proto3" json:"ec2_instance_id,omitempty" toml:"ec2_instance_id,omitempty" mapstructure:"ec2_instance_id,omitempty"`
	Ec2InstanceType      string               `protobuf:"bytes,9,opt,name=ec2_instance_type,json=ec2InstanceType,proto3" json:"ec2_instance_type,omitempty" toml:"ec2_instance_type,omitempty" mapstructure:"ec2_instance_type,omitempty"`
	LastCcrReceived      *timestamp.Timestamp `protobuf:"bytes,10,opt,name=last_ccr_received,json=lastCcrReceived,proto3" json:"last_ccr_received,omitempty" toml:"last_ccr_received,omitempty" mapstructure:"last_ccr_received,omitempty"`
	Name                 string               `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	Fqdn                 string               `protobuf:"bytes,12,opt,name=fqdn,proto3" json:"fqdn,omitempty" toml:"fqdn,omitempty" mapstructure:"fqdn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *InventoryNode) Reset()         { *m = InventoryNode{} }
func (m *InventoryNode) String() string { return proto.CompactTextString(m) }
func (*InventoryNode) ProtoMessage()    {}
func (*InventoryNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_inventory_nodes_659aa2ae1f39340d, []int{1}
}
func (m *InventoryNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InventoryNode.Unmarshal(m, b)
}
func (m *InventoryNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InventoryNode.Marshal(b, m, deterministic)
}
func (dst *InventoryNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InventoryNode.Merge(dst, src)
}
func (m *InventoryNode) XXX_Size() int {
	return xxx_messageInfo_InventoryNode.Size(m)
}
func (m *InventoryNode) XXX_DiscardUnknown() {
	xxx_messageInfo_InventoryNode.DiscardUnknown(m)
}

var xxx_messageInfo_InventoryNode proto.InternalMessageInfo

func (m *InventoryNode) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InventoryNode) GetCheckin() *timestamp.Timestamp {
	if m != nil {
		return m.Checkin
	}
	return nil
}

func (m *InventoryNode) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func (m *InventoryNode) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *InventoryNode) GetPlatformFamily() string {
	if m != nil {
		return m.PlatformFamily
	}
	return ""
}

func (m *InventoryNode) GetPlatformVersion() string {
	if m != nil {
		return m.PlatformVersion
	}
	return ""
}

func (m *InventoryNode) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

func (m *InventoryNode) GetEc2InstanceId() string {
	if m != nil {
		return m.Ec2InstanceId
	}
	return ""
}

func (m *InventoryNode) GetEc2InstanceType() string {
	if m != nil {
		return m.Ec2InstanceType
	}
	return ""
}

func (m *InventoryNode) GetLastCcrReceived() *timestamp.Timestamp {
	if m != nil {
		return m.LastCcrReceived
	}
	return nil
}

func (m *InventoryNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InventoryNode) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func init() {
	proto.RegisterType((*InventoryNodes)(nil), "chef.automate.domain.cfgmgmt.response.InventoryNodes")
	proto.RegisterType((*InventoryNode)(nil), "chef.automate.domain.cfgmgmt.response.InventoryNode")
}

func init() {
	proto.RegisterFile("api/interservice/cfgmgmt/response/inventory_nodes.proto", fileDescriptor_inventory_nodes_659aa2ae1f39340d)
}

var fileDescriptor_inventory_nodes_659aa2ae1f39340d = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x8b, 0xdb, 0x30,
	0x10, 0xc5, 0x49, 0xb2, 0x7f, 0x27, 0xbb, 0x49, 0x57, 0x27, 0x91, 0x4b, 0x43, 0x60, 0xdb, 0xb4,
	0x07, 0x09, 0xd2, 0x85, 0x42, 0xe9, 0xa9, 0x85, 0x85, 0xf4, 0xd0, 0x43, 0x58, 0x7a, 0x28, 0x05,
	0xa3, 0xc8, 0x63, 0x67, 0xa8, 0x25, 0xb9, 0xb2, 0x12, 0x48, 0x3f, 0x67, 0x3f, 0x50, 0xb1, 0x1c,
	0x85, 0xe6, 0xb4, 0xbd, 0x8d, 0x9f, 0x7e, 0xef, 0xbd, 0xc1, 0x0c, 0xbc, 0x57, 0x35, 0x49, 0xb2,
	0x01, 0x7d, 0x83, 0x7e, 0x47, 0x1a, 0xa5, 0x2e, 0x4a, 0x53, 0x9a, 0x20, 0x3d, 0x36, 0xb5, 0xb3,
	0x0d, 0x4a, 0xb2, 0x3b, 0xb4, 0xc1, 0xf9, 0x7d, 0x66, 0x5d, 0x8e, 0x8d, 0xa8, 0xbd, 0x0b, 0x8e,
	0xdd, 0xeb, 0x0d, 0x16, 0x42, 0x6d, 0x83, 0x33, 0x2a, 0xa0, 0xc8, 0x9d, 0x51, 0x64, 0xc5, 0xc1,
	0x2c, 0x92, 0x79, 0xf2, 0xb2, 0x74, 0xae, 0xac, 0x50, 0x46, 0xd3, 0x7a, 0x5b, 0xc8, 0x40, 0x06,
	0x9b, 0xa0, 0x4c, 0xdd, 0xe5, 0xcc, 0x7e, 0xc0, 0x68, 0x99, 0x0a, 0xbe, 0xb6, 0xf9, 0xec, 0x0b,
	0x9c, 0xc7, 0x22, 0xde, 0x9b, 0x0e, 0xe6, 0xc3, 0xc5, 0x83, 0xf8, 0xaf, 0x26, 0x71, 0x92, 0xb2,
	0xea, 0x22, 0x66, 0x7f, 0x06, 0x70, 0x7b, 0xf2, 0xc0, 0x46, 0xd0, 0xa7, 0x9c, 0xf7, 0xa6, 0xbd,
	0xf9, 0xf5, 0xaa, 0x4f, 0x39, 0x7b, 0x80, 0x4b, 0xbd, 0x41, 0xfd, 0x93, 0x2c, 0xef, 0x4f, 0x7b,
	0xf3, 0xe1, 0x62, 0x22, 0xba, 0x95, 0x45, 0x5a, 0x59, 0x3c, 0xa5, 0x95, 0x57, 0x09, 0x65, 0x33,
	0xb8, 0x71, 0xbe, 0x54, 0x96, 0x7e, 0xab, 0x40, 0xce, 0xf2, 0x41, 0xcc, 0x3b, 0xd1, 0xd8, 0x04,
	0xae, 0xea, 0x4a, 0x85, 0xc2, 0x79, 0xc3, 0xcf, 0xe2, 0xfb, 0xf1, 0x9b, 0xbd, 0x86, 0x71, 0x9a,
	0xb3, 0x42, 0x19, 0xaa, 0xf6, 0xfc, 0x3c, 0x22, 0xa3, 0x24, 0x3f, 0x46, 0x95, 0xbd, 0x81, 0x17,
	0x47, 0x70, 0x87, 0xbe, 0x69, 0xcb, 0x2e, 0x22, 0x79, 0x0c, 0xf8, 0xd6, 0xc9, 0xec, 0x1e, 0x46,
	0xba, 0x22, 0xb4, 0xe1, 0x08, 0x5e, 0x46, 0xf0, 0xb6, 0x53, 0x13, 0xf6, 0x0a, 0xc6, 0xa8, 0x17,
	0x19, 0xd9, 0x26, 0x28, 0xab, 0x31, 0xa3, 0x9c, 0x5f, 0x75, 0x1c, 0xea, 0xc5, 0xf2, 0xa0, 0x2e,
	0x73, 0xf6, 0x16, 0xee, 0x4e, 0xb8, 0xb0, 0xaf, 0x91, 0x5f, 0x77, 0xd5, 0xff, 0x90, 0x4f, 0xfb,
	0x1a, 0xd9, 0x23, 0xdc, 0x55, 0xaa, 0x09, 0x99, 0xd6, 0x3e, 0xf3, 0xa8, 0x91, 0x76, 0x98, 0x73,
	0x78, 0xf6, 0x77, 0x8e, 0x5b, 0xd3, 0x67, 0xed, 0x57, 0x07, 0x0b, 0x63, 0x70, 0x66, 0x95, 0x41,
	0x3e, 0x8c, 0x35, 0x71, 0x6e, 0xb5, 0xe2, 0x57, 0x6e, 0xf9, 0x4d, 0xa7, 0xb5, 0xf3, 0xa7, 0x8f,
	0xdf, 0x3f, 0x94, 0x14, 0x36, 0xdb, 0xb5, 0xd0, 0xce, 0xc8, 0xf6, 0x3e, 0x64, 0xba, 0x0f, 0xf9,
	0xec, 0x41, 0xaf, 0x2f, 0xe2, 0x2a, 0xef, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xda, 0xb2, 0xcf,
	0x23, 0xfc, 0x02, 0x00, 0x00,
}
