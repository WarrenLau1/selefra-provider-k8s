package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MockNamespacesClient struct {
	ctrl		*gomock.Controller
	recorder	*MockNamespacesClientMockRecorder
}

type MockNamespacesClientMockRecorder struct {
	mock *MockNamespacesClient
}

func NewMockNamespacesClient(ctrl *gomock.Controller) *MockNamespacesClient {
	mock := &MockNamespacesClient{ctrl: ctrl}
	mock.recorder = &MockNamespacesClientMockRecorder{mock}
	return mock
}

func (m *MockNamespacesClient) EXPECT() *MockNamespacesClientMockRecorder {
	return m.recorder
}

func (m *MockNamespacesClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.NamespaceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.NamespaceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNamespacesClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNamespacesClient)(nil).List), arg0, arg1)
}
