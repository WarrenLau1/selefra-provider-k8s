package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MockDeploymentsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockDeploymentsClientMockRecorder
}

type MockDeploymentsClientMockRecorder struct {
	mock *MockDeploymentsClient
}

func NewMockDeploymentsClient(ctrl *gomock.Controller) *MockDeploymentsClient {
	mock := &MockDeploymentsClient{ctrl: ctrl}
	mock.recorder = &MockDeploymentsClientMockRecorder{mock}
	return mock
}

func (m *MockDeploymentsClient) EXPECT() *MockDeploymentsClientMockRecorder {
	return m.recorder
}

func (m *MockDeploymentsClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.DeploymentList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.DeploymentList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDeploymentsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDeploymentsClient)(nil).List), arg0, arg1)
}
