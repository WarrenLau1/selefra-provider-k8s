package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/batch/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MockJobsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockJobsClientMockRecorder
}

type MockJobsClientMockRecorder struct {
	mock *MockJobsClient
}

func NewMockJobsClient(ctrl *gomock.Controller) *MockJobsClient {
	mock := &MockJobsClient{ctrl: ctrl}
	mock.recorder = &MockJobsClientMockRecorder{mock}
	return mock
}

func (m *MockJobsClient) EXPECT() *MockJobsClientMockRecorder {
	return m.recorder
}

func (m *MockJobsClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.JobList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.JobList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockJobsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockJobsClient)(nil).List), arg0, arg1)
}
