package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MockNodesClient struct {
	ctrl		*gomock.Controller
	recorder	*MockNodesClientMockRecorder
}

type MockNodesClientMockRecorder struct {
	mock *MockNodesClient
}

func NewMockNodesClient(ctrl *gomock.Controller) *MockNodesClient {
	mock := &MockNodesClient{ctrl: ctrl}
	mock.recorder = &MockNodesClientMockRecorder{mock}
	return mock
}

func (m *MockNodesClient) EXPECT() *MockNodesClientMockRecorder {
	return m.recorder
}

func (m *MockNodesClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.NodeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.NodeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNodesClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNodesClient)(nil).List), arg0, arg1)
}
