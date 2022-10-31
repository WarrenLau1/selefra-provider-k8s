package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/rbac/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MockRoleBindingsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockRoleBindingsClientMockRecorder
}

type MockRoleBindingsClientMockRecorder struct {
	mock *MockRoleBindingsClient
}

func NewMockRoleBindingsClient(ctrl *gomock.Controller) *MockRoleBindingsClient {
	mock := &MockRoleBindingsClient{ctrl: ctrl}
	mock.recorder = &MockRoleBindingsClientMockRecorder{mock}
	return mock
}

func (m *MockRoleBindingsClient) EXPECT() *MockRoleBindingsClientMockRecorder {
	return m.recorder
}

func (m *MockRoleBindingsClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.RoleBindingList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.RoleBindingList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoleBindingsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoleBindingsClient)(nil).List), arg0, arg1)
}
