package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MockLimitRangesClient struct {
	ctrl		*gomock.Controller
	recorder	*MockLimitRangesClientMockRecorder
}

type MockLimitRangesClientMockRecorder struct {
	mock *MockLimitRangesClient
}

func NewMockLimitRangesClient(ctrl *gomock.Controller) *MockLimitRangesClient {
	mock := &MockLimitRangesClient{ctrl: ctrl}
	mock.recorder = &MockLimitRangesClientMockRecorder{mock}
	return mock
}

func (m *MockLimitRangesClient) EXPECT() *MockLimitRangesClientMockRecorder {
	return m.recorder
}

func (m *MockLimitRangesClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.LimitRangeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.LimitRangeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLimitRangesClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLimitRangesClient)(nil).List), arg0, arg1)
}
