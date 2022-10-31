package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MockReplicaSetsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockReplicaSetsClientMockRecorder
}

type MockReplicaSetsClientMockRecorder struct {
	mock *MockReplicaSetsClient
}

func NewMockReplicaSetsClient(ctrl *gomock.Controller) *MockReplicaSetsClient {
	mock := &MockReplicaSetsClient{ctrl: ctrl}
	mock.recorder = &MockReplicaSetsClientMockRecorder{mock}
	return mock
}

func (m *MockReplicaSetsClient) EXPECT() *MockReplicaSetsClientMockRecorder {
	return m.recorder
}

func (m *MockReplicaSetsClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ReplicaSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.ReplicaSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicaSetsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockReplicaSetsClient)(nil).List), arg0, arg1)
}
