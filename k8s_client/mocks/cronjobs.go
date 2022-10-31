package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/batch/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MockCronJobsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCronJobsClientMockRecorder
}

type MockCronJobsClientMockRecorder struct {
	mock *MockCronJobsClient
}

func NewMockCronJobsClient(ctrl *gomock.Controller) *MockCronJobsClient {
	mock := &MockCronJobsClient{ctrl: ctrl}
	mock.recorder = &MockCronJobsClientMockRecorder{mock}
	return mock
}

func (m *MockCronJobsClient) EXPECT() *MockCronJobsClientMockRecorder {
	return m.recorder
}

func (m *MockCronJobsClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.CronJobList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.CronJobList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCronJobsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCronJobsClient)(nil).List), arg0, arg1)
}
