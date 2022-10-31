package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MockStatefulSetsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockStatefulSetsClientMockRecorder
}

type MockStatefulSetsClientMockRecorder struct {
	mock *MockStatefulSetsClient
}

func NewMockStatefulSetsClient(ctrl *gomock.Controller) *MockStatefulSetsClient {
	mock := &MockStatefulSetsClient{ctrl: ctrl}
	mock.recorder = &MockStatefulSetsClientMockRecorder{mock}
	return mock
}

func (m *MockStatefulSetsClient) EXPECT() *MockStatefulSetsClientMockRecorder {
	return m.recorder
}

func (m *MockStatefulSetsClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.StatefulSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.StatefulSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStatefulSetsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStatefulSetsClient)(nil).List), arg0, arg1)
}
