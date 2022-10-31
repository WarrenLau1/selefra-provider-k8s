package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MockDaemonSetsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockDaemonSetsClientMockRecorder
}

type MockDaemonSetsClientMockRecorder struct {
	mock *MockDaemonSetsClient
}

func NewMockDaemonSetsClient(ctrl *gomock.Controller) *MockDaemonSetsClient {
	mock := &MockDaemonSetsClient{ctrl: ctrl}
	mock.recorder = &MockDaemonSetsClientMockRecorder{mock}
	return mock
}

func (m *MockDaemonSetsClient) EXPECT() *MockDaemonSetsClientMockRecorder {
	return m.recorder
}

func (m *MockDaemonSetsClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.DaemonSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.DaemonSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaemonSetsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDaemonSetsClient)(nil).List), arg0, arg1)
}
