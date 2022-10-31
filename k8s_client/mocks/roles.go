package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/rbac/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MockRolesClient struct {
	ctrl		*gomock.Controller
	recorder	*MockRolesClientMockRecorder
}

type MockRolesClientMockRecorder struct {
	mock *MockRolesClient
}

func NewMockRolesClient(ctrl *gomock.Controller) *MockRolesClient {
	mock := &MockRolesClient{ctrl: ctrl}
	mock.recorder = &MockRolesClientMockRecorder{mock}
	return mock
}

func (m *MockRolesClient) EXPECT() *MockRolesClientMockRecorder {
	return m.recorder
}

func (m *MockRolesClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.RoleList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.RoleList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRolesClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRolesClient)(nil).List), arg0, arg1)
}
