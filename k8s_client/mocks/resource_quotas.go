package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MockResourceQuotasClient struct {
	ctrl		*gomock.Controller
	recorder	*MockResourceQuotasClientMockRecorder
}

type MockResourceQuotasClientMockRecorder struct {
	mock *MockResourceQuotasClient
}

func NewMockResourceQuotasClient(ctrl *gomock.Controller) *MockResourceQuotasClient {
	mock := &MockResourceQuotasClient{ctrl: ctrl}
	mock.recorder = &MockResourceQuotasClientMockRecorder{mock}
	return mock
}

func (m *MockResourceQuotasClient) EXPECT() *MockResourceQuotasClientMockRecorder {
	return m.recorder
}

func (m *MockResourceQuotasClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ResourceQuotaList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.ResourceQuotaList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourceQuotasClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockResourceQuotasClient)(nil).List), arg0, arg1)
}
