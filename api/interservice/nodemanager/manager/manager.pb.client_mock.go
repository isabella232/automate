// Code generated by MockGen. DO NOT EDIT.
// Source: nodemanager/manager/manager.pb.go

// Package manager is a generated GoMock package.
package manager

import (
	context "context"
	nodes "github.com/chef/automate/api/interservice/nodemanager/nodes"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

// MockNodeManagerServiceClient is a mock of NodeManagerServiceClient interface
type MockNodeManagerServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockNodeManagerServiceClientMockRecorder
}

// MockNodeManagerServiceClientMockRecorder is the mock recorder for MockNodeManagerServiceClient
type MockNodeManagerServiceClientMockRecorder struct {
	mock *MockNodeManagerServiceClient
}

// NewMockNodeManagerServiceClient creates a new mock instance
func NewMockNodeManagerServiceClient(ctrl *gomock.Controller) *MockNodeManagerServiceClient {
	mock := &MockNodeManagerServiceClient{ctrl: ctrl}
	mock.recorder = &MockNodeManagerServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeManagerServiceClient) EXPECT() *MockNodeManagerServiceClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockNodeManagerServiceClient) Create(ctx context.Context, in *NodeManager, opts ...grpc.CallOption) (*Ids, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*Ids)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockNodeManagerServiceClientMockRecorder) Create(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNodeManagerServiceClient)(nil).Create), varargs...)
}

// Read mocks base method
func (m *MockNodeManagerServiceClient) Read(ctx context.Context, in *Id, opts ...grpc.CallOption) (*NodeManager, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Read", varargs...)
	ret0, _ := ret[0].(*NodeManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockNodeManagerServiceClientMockRecorder) Read(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockNodeManagerServiceClient)(nil).Read), varargs...)
}

// Update mocks base method
func (m *MockNodeManagerServiceClient) Update(ctx context.Context, in *NodeManager, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockNodeManagerServiceClientMockRecorder) Update(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNodeManagerServiceClient)(nil).Update), varargs...)
}

// Delete mocks base method
func (m *MockNodeManagerServiceClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockNodeManagerServiceClientMockRecorder) Delete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNodeManagerServiceClient)(nil).Delete), varargs...)
}

// DeleteWithNodes mocks base method
func (m *MockNodeManagerServiceClient) DeleteWithNodes(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Ids, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteWithNodes", varargs...)
	ret0, _ := ret[0].(*Ids)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteWithNodes indicates an expected call of DeleteWithNodes
func (mr *MockNodeManagerServiceClientMockRecorder) DeleteWithNodes(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWithNodes", reflect.TypeOf((*MockNodeManagerServiceClient)(nil).DeleteWithNodes), varargs...)
}

// DeleteWithNodeStateStopped mocks base method
func (m *MockNodeManagerServiceClient) DeleteWithNodeStateStopped(ctx context.Context, in *Id, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteWithNodeStateStopped", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteWithNodeStateStopped indicates an expected call of DeleteWithNodeStateStopped
func (mr *MockNodeManagerServiceClientMockRecorder) DeleteWithNodeStateStopped(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWithNodeStateStopped", reflect.TypeOf((*MockNodeManagerServiceClient)(nil).DeleteWithNodeStateStopped), varargs...)
}

// DeleteWithNodeStateTerminated mocks base method
func (m *MockNodeManagerServiceClient) DeleteWithNodeStateTerminated(ctx context.Context, in *Id, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteWithNodeStateTerminated", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteWithNodeStateTerminated indicates an expected call of DeleteWithNodeStateTerminated
func (mr *MockNodeManagerServiceClientMockRecorder) DeleteWithNodeStateTerminated(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWithNodeStateTerminated", reflect.TypeOf((*MockNodeManagerServiceClient)(nil).DeleteWithNodeStateTerminated), varargs...)
}

// List mocks base method
func (m *MockNodeManagerServiceClient) List(ctx context.Context, in *Query, opts ...grpc.CallOption) (*NodeManagers, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*NodeManagers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockNodeManagerServiceClientMockRecorder) List(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNodeManagerServiceClient)(nil).List), varargs...)
}

// Connect mocks base method
func (m *MockNodeManagerServiceClient) Connect(ctx context.Context, in *NodeManager, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Connect", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect
func (mr *MockNodeManagerServiceClientMockRecorder) Connect(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockNodeManagerServiceClient)(nil).Connect), varargs...)
}

// ConnectManager mocks base method
func (m *MockNodeManagerServiceClient) ConnectManager(ctx context.Context, in *Id, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConnectManager", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectManager indicates an expected call of ConnectManager
func (mr *MockNodeManagerServiceClientMockRecorder) ConnectManager(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectManager", reflect.TypeOf((*MockNodeManagerServiceClient)(nil).ConnectManager), varargs...)
}

// SearchNodeFields mocks base method
func (m *MockNodeManagerServiceClient) SearchNodeFields(ctx context.Context, in *FieldQuery, opts ...grpc.CallOption) (*Fields, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchNodeFields", varargs...)
	ret0, _ := ret[0].(*Fields)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchNodeFields indicates an expected call of SearchNodeFields
func (mr *MockNodeManagerServiceClientMockRecorder) SearchNodeFields(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchNodeFields", reflect.TypeOf((*MockNodeManagerServiceClient)(nil).SearchNodeFields), varargs...)
}

// SearchNodes mocks base method
func (m *MockNodeManagerServiceClient) SearchNodes(ctx context.Context, in *NodeQuery, opts ...grpc.CallOption) (*Nodes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchNodes", varargs...)
	ret0, _ := ret[0].(*Nodes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchNodes indicates an expected call of SearchNodes
func (mr *MockNodeManagerServiceClientMockRecorder) SearchNodes(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchNodes", reflect.TypeOf((*MockNodeManagerServiceClient)(nil).SearchNodes), varargs...)
}

// ProcessNode mocks base method
func (m *MockNodeManagerServiceClient) ProcessNode(ctx context.Context, in *NodeMetadata, opts ...grpc.CallOption) (*ProcessNodeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProcessNode", varargs...)
	ret0, _ := ret[0].(*ProcessNodeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessNode indicates an expected call of ProcessNode
func (mr *MockNodeManagerServiceClientMockRecorder) ProcessNode(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessNode", reflect.TypeOf((*MockNodeManagerServiceClient)(nil).ProcessNode), varargs...)
}

// ChangeNodeState mocks base method
func (m *MockNodeManagerServiceClient) ChangeNodeState(ctx context.Context, in *NodeState, opts ...grpc.CallOption) (*ChangeNodeStateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeNodeState", varargs...)
	ret0, _ := ret[0].(*ChangeNodeStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeNodeState indicates an expected call of ChangeNodeState
func (mr *MockNodeManagerServiceClientMockRecorder) ChangeNodeState(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeNodeState", reflect.TypeOf((*MockNodeManagerServiceClient)(nil).ChangeNodeState), varargs...)
}

// GetNodeWithSecrets mocks base method
func (m *MockNodeManagerServiceClient) GetNodeWithSecrets(ctx context.Context, in *Id, opts ...grpc.CallOption) (*nodes.Node, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNodeWithSecrets", varargs...)
	ret0, _ := ret[0].(*nodes.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeWithSecrets indicates an expected call of GetNodeWithSecrets
func (mr *MockNodeManagerServiceClientMockRecorder) GetNodeWithSecrets(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeWithSecrets", reflect.TypeOf((*MockNodeManagerServiceClient)(nil).GetNodeWithSecrets), varargs...)
}

// SearchManagerNodes mocks base method
func (m *MockNodeManagerServiceClient) SearchManagerNodes(ctx context.Context, in *NodeQuery, opts ...grpc.CallOption) (*ManagerNodes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchManagerNodes", varargs...)
	ret0, _ := ret[0].(*ManagerNodes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchManagerNodes indicates an expected call of SearchManagerNodes
func (mr *MockNodeManagerServiceClientMockRecorder) SearchManagerNodes(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchManagerNodes", reflect.TypeOf((*MockNodeManagerServiceClient)(nil).SearchManagerNodes), varargs...)
}

// MockNodeManagerServiceServer is a mock of NodeManagerServiceServer interface
type MockNodeManagerServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockNodeManagerServiceServerMockRecorder
}

// MockNodeManagerServiceServerMockRecorder is the mock recorder for MockNodeManagerServiceServer
type MockNodeManagerServiceServerMockRecorder struct {
	mock *MockNodeManagerServiceServer
}

// NewMockNodeManagerServiceServer creates a new mock instance
func NewMockNodeManagerServiceServer(ctrl *gomock.Controller) *MockNodeManagerServiceServer {
	mock := &MockNodeManagerServiceServer{ctrl: ctrl}
	mock.recorder = &MockNodeManagerServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeManagerServiceServer) EXPECT() *MockNodeManagerServiceServerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockNodeManagerServiceServer) Create(arg0 context.Context, arg1 *NodeManager) (*Ids, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*Ids)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockNodeManagerServiceServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNodeManagerServiceServer)(nil).Create), arg0, arg1)
}

// Read mocks base method
func (m *MockNodeManagerServiceServer) Read(arg0 context.Context, arg1 *Id) (*NodeManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1)
	ret0, _ := ret[0].(*NodeManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockNodeManagerServiceServerMockRecorder) Read(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockNodeManagerServiceServer)(nil).Read), arg0, arg1)
}

// Update mocks base method
func (m *MockNodeManagerServiceServer) Update(arg0 context.Context, arg1 *NodeManager) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockNodeManagerServiceServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNodeManagerServiceServer)(nil).Update), arg0, arg1)
}

// Delete mocks base method
func (m *MockNodeManagerServiceServer) Delete(arg0 context.Context, arg1 *Id) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockNodeManagerServiceServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNodeManagerServiceServer)(nil).Delete), arg0, arg1)
}

// DeleteWithNodes mocks base method
func (m *MockNodeManagerServiceServer) DeleteWithNodes(arg0 context.Context, arg1 *Id) (*Ids, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWithNodes", arg0, arg1)
	ret0, _ := ret[0].(*Ids)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteWithNodes indicates an expected call of DeleteWithNodes
func (mr *MockNodeManagerServiceServerMockRecorder) DeleteWithNodes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWithNodes", reflect.TypeOf((*MockNodeManagerServiceServer)(nil).DeleteWithNodes), arg0, arg1)
}

// DeleteWithNodeStateStopped mocks base method
func (m *MockNodeManagerServiceServer) DeleteWithNodeStateStopped(arg0 context.Context, arg1 *Id) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWithNodeStateStopped", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteWithNodeStateStopped indicates an expected call of DeleteWithNodeStateStopped
func (mr *MockNodeManagerServiceServerMockRecorder) DeleteWithNodeStateStopped(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWithNodeStateStopped", reflect.TypeOf((*MockNodeManagerServiceServer)(nil).DeleteWithNodeStateStopped), arg0, arg1)
}

// DeleteWithNodeStateTerminated mocks base method
func (m *MockNodeManagerServiceServer) DeleteWithNodeStateTerminated(arg0 context.Context, arg1 *Id) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWithNodeStateTerminated", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteWithNodeStateTerminated indicates an expected call of DeleteWithNodeStateTerminated
func (mr *MockNodeManagerServiceServerMockRecorder) DeleteWithNodeStateTerminated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWithNodeStateTerminated", reflect.TypeOf((*MockNodeManagerServiceServer)(nil).DeleteWithNodeStateTerminated), arg0, arg1)
}

// List mocks base method
func (m *MockNodeManagerServiceServer) List(arg0 context.Context, arg1 *Query) (*NodeManagers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*NodeManagers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockNodeManagerServiceServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNodeManagerServiceServer)(nil).List), arg0, arg1)
}

// Connect mocks base method
func (m *MockNodeManagerServiceServer) Connect(arg0 context.Context, arg1 *NodeManager) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect
func (mr *MockNodeManagerServiceServerMockRecorder) Connect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockNodeManagerServiceServer)(nil).Connect), arg0, arg1)
}

// ConnectManager mocks base method
func (m *MockNodeManagerServiceServer) ConnectManager(arg0 context.Context, arg1 *Id) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectManager", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectManager indicates an expected call of ConnectManager
func (mr *MockNodeManagerServiceServerMockRecorder) ConnectManager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectManager", reflect.TypeOf((*MockNodeManagerServiceServer)(nil).ConnectManager), arg0, arg1)
}

// SearchNodeFields mocks base method
func (m *MockNodeManagerServiceServer) SearchNodeFields(arg0 context.Context, arg1 *FieldQuery) (*Fields, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchNodeFields", arg0, arg1)
	ret0, _ := ret[0].(*Fields)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchNodeFields indicates an expected call of SearchNodeFields
func (mr *MockNodeManagerServiceServerMockRecorder) SearchNodeFields(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchNodeFields", reflect.TypeOf((*MockNodeManagerServiceServer)(nil).SearchNodeFields), arg0, arg1)
}

// SearchNodes mocks base method
func (m *MockNodeManagerServiceServer) SearchNodes(arg0 context.Context, arg1 *NodeQuery) (*Nodes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchNodes", arg0, arg1)
	ret0, _ := ret[0].(*Nodes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchNodes indicates an expected call of SearchNodes
func (mr *MockNodeManagerServiceServerMockRecorder) SearchNodes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchNodes", reflect.TypeOf((*MockNodeManagerServiceServer)(nil).SearchNodes), arg0, arg1)
}

// ProcessNode mocks base method
func (m *MockNodeManagerServiceServer) ProcessNode(arg0 context.Context, arg1 *NodeMetadata) (*ProcessNodeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessNode", arg0, arg1)
	ret0, _ := ret[0].(*ProcessNodeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessNode indicates an expected call of ProcessNode
func (mr *MockNodeManagerServiceServerMockRecorder) ProcessNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessNode", reflect.TypeOf((*MockNodeManagerServiceServer)(nil).ProcessNode), arg0, arg1)
}

// ChangeNodeState mocks base method
func (m *MockNodeManagerServiceServer) ChangeNodeState(arg0 context.Context, arg1 *NodeState) (*ChangeNodeStateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeNodeState", arg0, arg1)
	ret0, _ := ret[0].(*ChangeNodeStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeNodeState indicates an expected call of ChangeNodeState
func (mr *MockNodeManagerServiceServerMockRecorder) ChangeNodeState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeNodeState", reflect.TypeOf((*MockNodeManagerServiceServer)(nil).ChangeNodeState), arg0, arg1)
}

// GetNodeWithSecrets mocks base method
func (m *MockNodeManagerServiceServer) GetNodeWithSecrets(arg0 context.Context, arg1 *Id) (*nodes.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeWithSecrets", arg0, arg1)
	ret0, _ := ret[0].(*nodes.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeWithSecrets indicates an expected call of GetNodeWithSecrets
func (mr *MockNodeManagerServiceServerMockRecorder) GetNodeWithSecrets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeWithSecrets", reflect.TypeOf((*MockNodeManagerServiceServer)(nil).GetNodeWithSecrets), arg0, arg1)
}

// SearchManagerNodes mocks base method
func (m *MockNodeManagerServiceServer) SearchManagerNodes(arg0 context.Context, arg1 *NodeQuery) (*ManagerNodes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchManagerNodes", arg0, arg1)
	ret0, _ := ret[0].(*ManagerNodes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchManagerNodes indicates an expected call of SearchManagerNodes
func (mr *MockNodeManagerServiceServerMockRecorder) SearchManagerNodes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchManagerNodes", reflect.TypeOf((*MockNodeManagerServiceServer)(nil).SearchManagerNodes), arg0, arg1)
}
