// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/auth/teams/teams.proto

package teams // import "github.com/chef/automate/components/automate-gateway/api/auth/teams"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import version "github.com/chef/automate/api/external/common/version"
import request "github.com/chef/automate/components/automate-gateway/api/auth/teams/request"
import response "github.com/chef/automate/components/automate-gateway/api/auth/teams/response"
import _ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/api"
import _ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TeamsClient is the client API for Teams service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TeamsClient interface {
	GetVersion(ctx context.Context, in *version.VersionInfoRequest, opts ...grpc.CallOption) (*version.VersionInfo, error)
	GetTeams(ctx context.Context, in *request.GetTeamsReq, opts ...grpc.CallOption) (*response.Teams, error)
	GetTeam(ctx context.Context, in *request.GetTeamReq, opts ...grpc.CallOption) (*response.GetTeamResp, error)
	CreateTeam(ctx context.Context, in *request.CreateTeamReq, opts ...grpc.CallOption) (*response.CreateTeamResp, error)
	UpdateTeam(ctx context.Context, in *request.UpdateTeamReq, opts ...grpc.CallOption) (*response.UpdateTeamResp, error)
	DeleteTeam(ctx context.Context, in *request.DeleteTeamReq, opts ...grpc.CallOption) (*response.DeleteTeamResp, error)
	GetUsers(ctx context.Context, in *request.GetUsersReq, opts ...grpc.CallOption) (*response.GetUsersResp, error)
	AddUsers(ctx context.Context, in *request.AddUsersReq, opts ...grpc.CallOption) (*response.AddUsersResp, error)
	RemoveUsers(ctx context.Context, in *request.RemoveUsersReq, opts ...grpc.CallOption) (*response.RemoveUsersResp, error)
	GetTeamsForUser(ctx context.Context, in *request.GetTeamsForUserReq, opts ...grpc.CallOption) (*response.GetTeamsForUserResp, error)
}

type teamsClient struct {
	cc *grpc.ClientConn
}

func NewTeamsClient(cc *grpc.ClientConn) TeamsClient {
	return &teamsClient{cc}
}

func (c *teamsClient) GetVersion(ctx context.Context, in *version.VersionInfoRequest, opts ...grpc.CallOption) (*version.VersionInfo, error) {
	out := new(version.VersionInfo)
	err := c.cc.Invoke(ctx, "/chef.automate.api.teams.Teams/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) GetTeams(ctx context.Context, in *request.GetTeamsReq, opts ...grpc.CallOption) (*response.Teams, error) {
	out := new(response.Teams)
	err := c.cc.Invoke(ctx, "/chef.automate.api.teams.Teams/GetTeams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) GetTeam(ctx context.Context, in *request.GetTeamReq, opts ...grpc.CallOption) (*response.GetTeamResp, error) {
	out := new(response.GetTeamResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.teams.Teams/GetTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) CreateTeam(ctx context.Context, in *request.CreateTeamReq, opts ...grpc.CallOption) (*response.CreateTeamResp, error) {
	out := new(response.CreateTeamResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.teams.Teams/CreateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) UpdateTeam(ctx context.Context, in *request.UpdateTeamReq, opts ...grpc.CallOption) (*response.UpdateTeamResp, error) {
	out := new(response.UpdateTeamResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.teams.Teams/UpdateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) DeleteTeam(ctx context.Context, in *request.DeleteTeamReq, opts ...grpc.CallOption) (*response.DeleteTeamResp, error) {
	out := new(response.DeleteTeamResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.teams.Teams/DeleteTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) GetUsers(ctx context.Context, in *request.GetUsersReq, opts ...grpc.CallOption) (*response.GetUsersResp, error) {
	out := new(response.GetUsersResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.teams.Teams/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) AddUsers(ctx context.Context, in *request.AddUsersReq, opts ...grpc.CallOption) (*response.AddUsersResp, error) {
	out := new(response.AddUsersResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.teams.Teams/AddUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) RemoveUsers(ctx context.Context, in *request.RemoveUsersReq, opts ...grpc.CallOption) (*response.RemoveUsersResp, error) {
	out := new(response.RemoveUsersResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.teams.Teams/RemoveUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) GetTeamsForUser(ctx context.Context, in *request.GetTeamsForUserReq, opts ...grpc.CallOption) (*response.GetTeamsForUserResp, error) {
	out := new(response.GetTeamsForUserResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.teams.Teams/GetTeamsForUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamsServer is the server API for Teams service.
type TeamsServer interface {
	GetVersion(context.Context, *version.VersionInfoRequest) (*version.VersionInfo, error)
	GetTeams(context.Context, *request.GetTeamsReq) (*response.Teams, error)
	GetTeam(context.Context, *request.GetTeamReq) (*response.GetTeamResp, error)
	CreateTeam(context.Context, *request.CreateTeamReq) (*response.CreateTeamResp, error)
	UpdateTeam(context.Context, *request.UpdateTeamReq) (*response.UpdateTeamResp, error)
	DeleteTeam(context.Context, *request.DeleteTeamReq) (*response.DeleteTeamResp, error)
	GetUsers(context.Context, *request.GetUsersReq) (*response.GetUsersResp, error)
	AddUsers(context.Context, *request.AddUsersReq) (*response.AddUsersResp, error)
	RemoveUsers(context.Context, *request.RemoveUsersReq) (*response.RemoveUsersResp, error)
	GetTeamsForUser(context.Context, *request.GetTeamsForUserReq) (*response.GetTeamsForUserResp, error)
}

func RegisterTeamsServer(s *grpc.Server, srv TeamsServer) {
	s.RegisterService(&_Teams_serviceDesc, srv)
}

func _Teams_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(version.VersionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.teams.Teams/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).GetVersion(ctx, req.(*version.VersionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_GetTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetTeamsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).GetTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.teams.Teams/GetTeams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).GetTeams(ctx, req.(*request.GetTeamsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_GetTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).GetTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.teams.Teams/GetTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).GetTeam(ctx, req.(*request.GetTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_CreateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).CreateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.teams.Teams/CreateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).CreateTeam(ctx, req.(*request.CreateTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_UpdateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).UpdateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.teams.Teams/UpdateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).UpdateTeam(ctx, req.(*request.UpdateTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_DeleteTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).DeleteTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.teams.Teams/DeleteTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).DeleteTeam(ctx, req.(*request.DeleteTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.teams.Teams/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).GetUsers(ctx, req.(*request.GetUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_AddUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.AddUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).AddUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.teams.Teams/AddUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).AddUsers(ctx, req.(*request.AddUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_RemoveUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.RemoveUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).RemoveUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.teams.Teams/RemoveUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).RemoveUsers(ctx, req.(*request.RemoveUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_GetTeamsForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetTeamsForUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).GetTeamsForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.teams.Teams/GetTeamsForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).GetTeamsForUser(ctx, req.(*request.GetTeamsForUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Teams_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.teams.Teams",
	HandlerType: (*TeamsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Teams_GetVersion_Handler,
		},
		{
			MethodName: "GetTeams",
			Handler:    _Teams_GetTeams_Handler,
		},
		{
			MethodName: "GetTeam",
			Handler:    _Teams_GetTeam_Handler,
		},
		{
			MethodName: "CreateTeam",
			Handler:    _Teams_CreateTeam_Handler,
		},
		{
			MethodName: "UpdateTeam",
			Handler:    _Teams_UpdateTeam_Handler,
		},
		{
			MethodName: "DeleteTeam",
			Handler:    _Teams_DeleteTeam_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Teams_GetUsers_Handler,
		},
		{
			MethodName: "AddUsers",
			Handler:    _Teams_AddUsers_Handler,
		},
		{
			MethodName: "RemoveUsers",
			Handler:    _Teams_RemoveUsers_Handler,
		},
		{
			MethodName: "GetTeamsForUser",
			Handler:    _Teams_GetTeamsForUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/auth/teams/teams.proto",
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/auth/teams/teams.proto", fileDescriptor_teams_ee016a0f3d1d0810)
}

var fileDescriptor_teams_ee016a0f3d1d0810 = []byte{
	// 788 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0x3f, 0x73, 0xd3, 0x3e,
	0x18, 0xc7, 0xcf, 0xbf, 0x3f, 0x21, 0xa8, 0x94, 0xa6, 0x6a, 0x9a, 0x8a, 0xc0, 0xe4, 0x85, 0xa3,
	0x10, 0x1b, 0x28, 0x30, 0x78, 0xe0, 0x5f, 0x81, 0x1e, 0x1b, 0xd7, 0x6b, 0x19, 0xb8, 0x03, 0xce,
	0x75, 0x9e, 0xa6, 0xbe, 0x8b, 0x2d, 0xd5, 0x52, 0x02, 0x85, 0x63, 0xc9, 0x98, 0x95, 0x57, 0xc0,
	0x6b, 0xe0, 0xfc, 0x12, 0x18, 0x99, 0xba, 0x72, 0x07, 0x03, 0x0b, 0x03, 0x13, 0x6f, 0x80, 0x93,
	0x64, 0xc7, 0x76, 0xd3, 0x24, 0xce, 0xd2, 0x5e, 0xac, 0xef, 0x57, 0x7e, 0x3e, 0xfa, 0x3e, 0xb2,
	0x84, 0xee, 0x78, 0x34, 0x60, 0x34, 0x84, 0x50, 0x70, 0xdb, 0xed, 0x09, 0x1a, 0xb8, 0x02, 0x5a,
	0x1d, 0x57, 0xc0, 0x1b, 0xf7, 0xc8, 0x76, 0x99, 0x2f, 0x1f, 0x1e, 0xd8, 0x02, 0xdc, 0x80, 0xeb,
	0xbf, 0x16, 0x8b, 0xa8, 0xa0, 0x78, 0xcd, 0x3b, 0x80, 0x7d, 0x2b, 0x75, 0x58, 0x2e, 0xf3, 0x2d,
	0x35, 0xdc, 0xbc, 0xd4, 0xa1, 0xb4, 0xd3, 0x05, 0xed, 0x0d, 0x43, 0x2a, 0x5c, 0xe1, 0xd3, 0x30,
	0xb1, 0x35, 0xaf, 0xc8, 0xc7, 0xf0, 0x56, 0x40, 0x14, 0xba, 0x5d, 0xdb, 0xa3, 0x41, 0x40, 0x43,
	0xbb, 0x0f, 0x11, 0xf7, 0xb3, 0xff, 0x89, 0xf4, 0xee, 0x1c, 0x95, 0x45, 0x70, 0xd8, 0x03, 0x2e,
	0xf2, 0x15, 0x36, 0xef, 0xcd, 0xe5, 0xe7, 0x8c, 0x86, 0x1c, 0x0a, 0x13, 0xdc, 0x3f, 0x75, 0x82,
	0x88, 0x79, 0xb6, 0x1a, 0xf7, 0x5a, 0x1d, 0x08, 0x5b, 0x8c, 0x76, 0x7d, 0xef, 0x68, 0x02, 0xed,
	0x3c, 0x33, 0xf8, 0x6e, 0x30, 0x3e, 0xc3, 0xcd, 0x4f, 0xcb, 0xe8, 0xff, 0x1d, 0x59, 0x13, 0xfe,
	0x6d, 0x20, 0xb4, 0x05, 0xe2, 0xb9, 0x5e, 0x23, 0x7c, 0xcb, 0x1a, 0x0f, 0x40, 0x2f, 0xa7, 0x95,
	0x2e, 0x63, 0x22, 0x7d, 0x1a, 0xee, 0xd3, 0x6d, 0xbd, 0x3e, 0xcd, 0xd6, 0x5c, 0x2e, 0xf3, 0xdd,
	0x20, 0x26, 0x0d, 0x54, 0xe7, 0x10, 0xf5, 0x7d, 0x0f, 0x5e, 0xfb, 0xe1, 0x3e, 0x75, 0x12, 0xdd,
	0x20, 0x26, 0x15, 0xfc, 0x5f, 0x04, 0x6e, 0x7b, 0x18, 0x13, 0x82, 0x1a, 0xfc, 0x88, 0x0b, 0x08,
	0x9c, 0x44, 0x9a, 0xaa, 0x86, 0x31, 0xb9, 0x88, 0x2f, 0x14, 0xc7, 0x92, 0x17, 0x38, 0x1d, 0x10,
	0x83, 0xe3, 0x9f, 0x1f, 0xff, 0x59, 0xc5, 0x2b, 0xf9, 0x28, 0x12, 0x27, 0xfe, 0x6c, 0xa0, 0xea,
	0x16, 0x08, 0x0d, 0x7f, 0xcd, 0x9a, 0xd0, 0x6e, 0x56, 0x92, 0xbc, 0x95, 0x4a, 0xb7, 0xe1, 0xb0,
	0x79, 0x79, 0x8a, 0x5a, 0xe7, 0x6c, 0x29, 0xad, 0xf9, 0x6c, 0x10, 0x93, 0x73, 0x08, 0xc9, 0xb7,
	0x3b, 0x4a, 0x50, 0xa0, 0x5a, 0x40, 0x67, 0x7d, 0x37, 0xd0, 0x03, 0xc3, 0x98, 0xd4, 0xf0, 0xf9,
	0xd1, 0x4f, 0xa7, 0xeb, 0x73, 0x5d, 0xfd, 0x22, 0x5e, 0xc8, 0x55, 0x8f, 0xbf, 0x18, 0xe8, 0x4c,
	0x52, 0x0a, 0xbe, 0x5a, 0xb6, 0x68, 0x59, 0x73, 0x6b, 0x76, 0xcd, 0x23, 0x35, 0x67, 0xe6, 0xab,
	0x41, 0x4c, 0x96, 0xd1, 0x52, 0x56, 0xb9, 0xf3, 0xde, 0x6f, 0x7f, 0x28, 0x94, 0x5f, 0x43, 0xb9,
	0x7a, 0xe5, 0xe8, 0x30, 0x26, 0x4b, 0x78, 0x31, 0x7b, 0x96, 0x06, 0x80, 0x71, 0x2d, 0x1f, 0x80,
	0x94, 0xe2, 0xaf, 0x06, 0x42, 0x9b, 0x11, 0xb8, 0x02, 0x14, 0x8a, 0x35, 0x13, 0x25, 0x13, 0x4b,
	0x9a, 0xeb, 0xb3, 0x69, 0xf2, 0x06, 0xce, 0xcc, 0x97, 0xa7, 0x44, 0x51, 0xc5, 0x15, 0x4f, 0xa9,
	0xc6, 0xc3, 0xc0, 0xb8, 0x96, 0x81, 0x68, 0x91, 0x8e, 0xc3, 0xcc, 0xc7, 0xa1, 0x1e, 0xfd, 0xeb,
	0x18, 0xeb, 0xf8, 0x9b, 0x81, 0xd0, 0x2e, 0x6b, 0x97, 0xe7, 0xc9, 0xc4, 0x25, 0x79, 0xf2, 0x06,
	0xce, 0xcc, 0xee, 0xa4, 0x80, 0xaa, 0xb8, 0xd2, 0x53, 0xd2, 0x09, 0x11, 0x15, 0xc8, 0xb4, 0x52,
	0xa7, 0xd4, 0x1c, 0x4b, 0x29, 0xc3, 0x3b, 0x36, 0x10, 0x7a, 0x04, 0x5d, 0x28, 0x8d, 0x97, 0x89,
	0x4b, 0xe2, 0xe5, 0x0d, 0x9c, 0x99, 0xde, 0x14, 0xbc, 0xb6, 0x92, 0x96, 0xc1, 0xd3, 0x4a, 0x8d,
	0xb7, 0x3e, 0xde, 0x84, 0xdf, 0xf5, 0x27, 0x60, 0x97, 0x43, 0x54, 0xf2, 0x13, 0xa0, 0xa4, 0x92,
	0xc8, 0x2a, 0xb5, 0x9d, 0x12, 0x39, 0x67, 0x66, 0x6f, 0x10, 0x93, 0x35, 0xb4, 0x7a, 0x82, 0xc7,
	0xe9, 0x49, 0x41, 0x61, 0x57, 0x35, 0x50, 0xbd, 0xc8, 0xa4, 0x35, 0xc3, 0x98, 0xd4, 0x31, 0x4e,
	0x47, 0xd4, 0xbc, 0xd9, 0x37, 0x82, 0xe0, 0xc6, 0x49, 0x36, 0x5b, 0xb9, 0xf0, 0x0f, 0x03, 0x55,
	0x1f, 0xb4, 0xdb, 0x65, 0x09, 0x53, 0x69, 0x49, 0xc2, 0x4c, 0xce, 0x99, 0xd9, 0x9f, 0x92, 0xd8,
	0x68, 0x97, 0x9d, 0x96, 0x58, 0x03, 0xd7, 0x8b, 0x5c, 0xb9, 0xed, 0x46, 0xcc, 0x09, 0x64, 0x59,
	0x6b, 0xfe, 0x32, 0xd0, 0xc2, 0x36, 0x04, 0xb4, 0x0f, 0x9a, 0xd2, 0x9e, 0x49, 0x99, 0x53, 0x4b,
	0xd0, 0x1b, 0xb3, 0x41, 0x0b, 0x8e, 0x19, 0xac, 0x53, 0xbb, 0x73, 0x8c, 0x35, 0xd7, 0xa1, 0xa4,
	0x39, 0x93, 0xf5, 0x8f, 0x81, 0x96, 0xd2, 0x83, 0xe8, 0x09, 0x8d, 0xe4, 0x14, 0x78, 0xa3, 0xf4,
	0xd1, 0x95, 0x38, 0x24, 0xf3, 0xed, 0xd2, 0xa7, 0x41, 0xe6, 0xe2, 0xcc, 0x14, 0xb9, 0x2e, 0x56,
	0xa5, 0xe9, 0x0e, 0x1d, 0x3f, 0xda, 0x92, 0x2e, 0x3e, 0xa9, 0x19, 0xc6, 0x64, 0x05, 0x2f, 0xa7,
	0x23, 0x3b, 0x85, 0x53, 0x62, 0xd4, 0xc4, 0xca, 0xa4, 0xf1, 0x95, 0xe9, 0xe1, 0xe3, 0x17, 0x9b,
	0x1d, 0x5f, 0x1c, 0xf4, 0xf6, 0xe4, 0x55, 0xc2, 0x96, 0x85, 0x8f, 0x2e, 0x3b, 0x76, 0xf9, 0x3b,
	0xd8, 0x5e, 0x45, 0xdd, 0x78, 0x36, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x76, 0xf3, 0xb1,
	0x92, 0x0a, 0x00, 0x00,
}
