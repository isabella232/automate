// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: interservice/cfgmgmt/request/node_export.proto

package request

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type NodeExport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter     []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty" toml:"filter,omitempty" mapstructure:"filter,omitempty"`
	Sorting    *Sorting `protobuf:"bytes,2,opt,name=sorting,proto3" json:"sorting,omitempty" toml:"sorting,omitempty" mapstructure:"sorting,omitempty"`
	OutputType string   `protobuf:"bytes,3,opt,name=output_type,json=outputType,proto3" json:"output_type,omitempty" toml:"output_type,omitempty" mapstructure:"output_type,omitempty"`
}

func (x *NodeExport) Reset() {
	*x = NodeExport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_cfgmgmt_request_node_export_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeExport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeExport) ProtoMessage() {}

func (x *NodeExport) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_cfgmgmt_request_node_export_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeExport.ProtoReflect.Descriptor instead.
func (*NodeExport) Descriptor() ([]byte, []int) {
	return file_interservice_cfgmgmt_request_node_export_proto_rawDescGZIP(), []int{0}
}

func (x *NodeExport) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *NodeExport) GetSorting() *Sorting {
	if x != nil {
		return x.Sorting
	}
	return nil
}

func (x *NodeExport) GetOutputType() string {
	if x != nil {
		return x.OutputType
	}
	return ""
}

type ReportExport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter     []string               `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty" toml:"filter,omitempty" mapstructure:"filter,omitempty"`
	OutputType string                 `protobuf:"bytes,2,opt,name=output_type,json=outputType,proto3" json:"output_type,omitempty" toml:"output_type,omitempty" mapstructure:"output_type,omitempty"`
	NodeId     string                 `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty" toml:"node_id,omitempty" mapstructure:"node_id,omitempty"`
	Start      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty" toml:"start,omitempty" mapstructure:"start,omitempty"`
	End        *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty" toml:"end,omitempty" mapstructure:"end,omitempty"`
}

func (x *ReportExport) Reset() {
	*x = ReportExport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_cfgmgmt_request_node_export_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportExport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportExport) ProtoMessage() {}

func (x *ReportExport) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_cfgmgmt_request_node_export_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportExport.ProtoReflect.Descriptor instead.
func (*ReportExport) Descriptor() ([]byte, []int) {
	return file_interservice_cfgmgmt_request_node_export_proto_rawDescGZIP(), []int{1}
}

func (x *ReportExport) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ReportExport) GetOutputType() string {
	if x != nil {
		return x.OutputType
	}
	return ""
}

func (x *ReportExport) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *ReportExport) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *ReportExport) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

var File_interservice_cfgmgmt_request_node_export_proto protoreflect.FileDescriptor

var file_interservice_cfgmgmt_request_node_export_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63,
	0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x24, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x47, 0x0a,
	0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xc0, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03,
	0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interservice_cfgmgmt_request_node_export_proto_rawDescOnce sync.Once
	file_interservice_cfgmgmt_request_node_export_proto_rawDescData = file_interservice_cfgmgmt_request_node_export_proto_rawDesc
)

func file_interservice_cfgmgmt_request_node_export_proto_rawDescGZIP() []byte {
	file_interservice_cfgmgmt_request_node_export_proto_rawDescOnce.Do(func() {
		file_interservice_cfgmgmt_request_node_export_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_cfgmgmt_request_node_export_proto_rawDescData)
	})
	return file_interservice_cfgmgmt_request_node_export_proto_rawDescData
}

var file_interservice_cfgmgmt_request_node_export_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_interservice_cfgmgmt_request_node_export_proto_goTypes = []interface{}{
	(*NodeExport)(nil),            // 0: chef.automate.domain.cfgmgmt.request.NodeExport
	(*ReportExport)(nil),          // 1: chef.automate.domain.cfgmgmt.request.ReportExport
	(*Sorting)(nil),               // 2: chef.automate.domain.cfgmgmt.request.Sorting
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_interservice_cfgmgmt_request_node_export_proto_depIdxs = []int32{
	2, // 0: chef.automate.domain.cfgmgmt.request.NodeExport.sorting:type_name -> chef.automate.domain.cfgmgmt.request.Sorting
	3, // 1: chef.automate.domain.cfgmgmt.request.ReportExport.start:type_name -> google.protobuf.Timestamp
	3, // 2: chef.automate.domain.cfgmgmt.request.ReportExport.end:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_interservice_cfgmgmt_request_node_export_proto_init() }
func file_interservice_cfgmgmt_request_node_export_proto_init() {
	if File_interservice_cfgmgmt_request_node_export_proto != nil {
		return
	}
	file_interservice_cfgmgmt_request_parameters_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_interservice_cfgmgmt_request_node_export_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeExport); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interservice_cfgmgmt_request_node_export_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportExport); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_interservice_cfgmgmt_request_node_export_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interservice_cfgmgmt_request_node_export_proto_goTypes,
		DependencyIndexes: file_interservice_cfgmgmt_request_node_export_proto_depIdxs,
		MessageInfos:      file_interservice_cfgmgmt_request_node_export_proto_msgTypes,
	}.Build()
	File_interservice_cfgmgmt_request_node_export_proto = out.File
	file_interservice_cfgmgmt_request_node_export_proto_rawDesc = nil
	file_interservice_cfgmgmt_request_node_export_proto_goTypes = nil
	file_interservice_cfgmgmt_request_node_export_proto_depIdxs = nil
}
