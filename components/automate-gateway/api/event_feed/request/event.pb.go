// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/event_feed/request/event.proto

package request // import "github.com/chef/automate/components/automate-gateway/api/event_feed/request"

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

type EventFilter struct {
	Filter               []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	Start                int64    `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End                  int64    `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	PageSize             int32    `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	After                int64    `protobuf:"varint,5,opt,name=after,proto3" json:"after,omitempty"`
	Before               int64    `protobuf:"varint,6,opt,name=before,proto3" json:"before,omitempty"`
	Cursor               string   `protobuf:"bytes,7,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Collapse             bool     `protobuf:"varint,8,opt,name=collapse,proto3" json:"collapse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventFilter) Reset()         { *m = EventFilter{} }
func (m *EventFilter) String() string { return proto.CompactTextString(m) }
func (*EventFilter) ProtoMessage()    {}
func (*EventFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_b185a937b9c8fea7, []int{0}
}
func (m *EventFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventFilter.Unmarshal(m, b)
}
func (m *EventFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventFilter.Marshal(b, m, deterministic)
}
func (dst *EventFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFilter.Merge(dst, src)
}
func (m *EventFilter) XXX_Size() int {
	return xxx_messageInfo_EventFilter.Size(m)
}
func (m *EventFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFilter.DiscardUnknown(m)
}

var xxx_messageInfo_EventFilter proto.InternalMessageInfo

func (m *EventFilter) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *EventFilter) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *EventFilter) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *EventFilter) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *EventFilter) GetAfter() int64 {
	if m != nil {
		return m.After
	}
	return 0
}

func (m *EventFilter) GetBefore() int64 {
	if m != nil {
		return m.Before
	}
	return 0
}

func (m *EventFilter) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func (m *EventFilter) GetCollapse() bool {
	if m != nil {
		return m.Collapse
	}
	return false
}

type EventCountsFilter struct {
	Filter               []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	Start                int64    `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End                  int64    `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventCountsFilter) Reset()         { *m = EventCountsFilter{} }
func (m *EventCountsFilter) String() string { return proto.CompactTextString(m) }
func (*EventCountsFilter) ProtoMessage()    {}
func (*EventCountsFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_b185a937b9c8fea7, []int{1}
}
func (m *EventCountsFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventCountsFilter.Unmarshal(m, b)
}
func (m *EventCountsFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventCountsFilter.Marshal(b, m, deterministic)
}
func (dst *EventCountsFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCountsFilter.Merge(dst, src)
}
func (m *EventCountsFilter) XXX_Size() int {
	return xxx_messageInfo_EventCountsFilter.Size(m)
}
func (m *EventCountsFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCountsFilter.DiscardUnknown(m)
}

var xxx_messageInfo_EventCountsFilter proto.InternalMessageInfo

func (m *EventCountsFilter) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *EventCountsFilter) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *EventCountsFilter) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func init() {
	proto.RegisterType((*EventFilter)(nil), "chef.automate.api.event_feed.request.EventFilter")
	proto.RegisterType((*EventCountsFilter)(nil), "chef.automate.api.event_feed.request.EventCountsFilter")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/event_feed/request/event.proto", fileDescriptor_event_b185a937b9c8fea7)
}

var fileDescriptor_event_b185a937b9c8fea7 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x31, 0x4f, 0xf3, 0x30,
	0x10, 0x86, 0xe5, 0x2f, 0x5f, 0x43, 0x6a, 0x16, 0xb0, 0x10, 0xb2, 0x60, 0x89, 0x2a, 0x86, 0x2c,
	0xd8, 0x03, 0x3b, 0x03, 0x08, 0x16, 0xc4, 0x92, 0x6e, 0x2c, 0x95, 0x93, 0x5e, 0x52, 0x4b, 0x49,
	0x6c, 0xec, 0x0b, 0x88, 0xfe, 0x46, 0x7e, 0x14, 0xb2, 0x13, 0xca, 0x8c, 0xc4, 0x76, 0xcf, 0x23,
	0xdd, 0x9b, 0x9c, 0x5f, 0x7a, 0x5b, 0x9b, 0xde, 0x9a, 0x01, 0x06, 0xf4, 0x52, 0x8d, 0x68, 0x7a,
	0x85, 0x70, 0xdd, 0x2a, 0x84, 0x77, 0xf5, 0x21, 0x95, 0xd5, 0x12, 0xde, 0x60, 0xc0, 0x4d, 0x03,
	0xb0, 0x95, 0x0e, 0x5e, 0x47, 0xf0, 0x38, 0x29, 0x61, 0x9d, 0x41, 0xc3, 0xae, 0xea, 0x1d, 0x34,
	0xe2, 0x7b, 0x53, 0x28, 0xab, 0xc5, 0xcf, 0x86, 0x98, 0x37, 0x56, 0x9f, 0x84, 0x1e, 0x3f, 0x04,
	0xfd, 0xa8, 0x3b, 0x04, 0xc7, 0xce, 0x69, 0xda, 0xc4, 0x89, 0x93, 0x3c, 0x29, 0x96, 0xe5, 0x4c,
	0xec, 0x8c, 0x2e, 0x3c, 0x2a, 0x87, 0xfc, 0x5f, 0x4e, 0x8a, 0xa4, 0x9c, 0x80, 0x9d, 0xd0, 0x04,
	0x86, 0x2d, 0x4f, 0xa2, 0x0b, 0x23, 0xbb, 0xa4, 0x4b, 0xab, 0x5a, 0xd8, 0x78, 0xbd, 0x07, 0xfe,
	0x3f, 0x27, 0xc5, 0xa2, 0xcc, 0x82, 0x58, 0xeb, 0x3d, 0x84, 0x10, 0xd5, 0x84, 0xec, 0xc5, 0x14,
	0x12, 0x21, 0x7c, 0xb2, 0x82, 0xc6, 0x38, 0xe0, 0x69, 0xd4, 0x33, 0x05, 0x5f, 0x8f, 0xce, 0x1b,
	0xc7, 0x8f, 0x72, 0x12, 0x7e, 0x65, 0x22, 0x76, 0x41, 0xb3, 0xda, 0x74, 0x9d, 0xb2, 0x1e, 0x78,
	0x96, 0x93, 0x22, 0x2b, 0x0f, 0xbc, 0x5a, 0xd3, 0xd3, 0x78, 0xcd, 0xbd, 0x19, 0x07, 0xf4, 0x7f,
	0x73, 0xd3, 0xdd, 0xf3, 0xcb, 0x53, 0xab, 0x71, 0x37, 0x56, 0xa2, 0x36, 0xbd, 0x0c, 0xcf, 0x7a,
	0x28, 0x44, 0xfe, 0xbe, 0xa4, 0x2a, 0x8d, 0xfd, 0xdc, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8d,
	0x25, 0x88, 0xad, 0xe1, 0x01, 0x00, 0x00,
}
