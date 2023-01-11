package v1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/networking/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/networking/v1"
	v12 "k8s.io/client-go/kubernetes/typed/networking/v1"
)

type MockIngressClassesGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockIngressClassesGetterMockRecorder
}

type MockIngressClassesGetterMockRecorder struct {
	mock *MockIngressClassesGetter
}

func NewMockIngressClassesGetter(ctrl *gomock.Controller) *MockIngressClassesGetter {
	mock := &MockIngressClassesGetter{ctrl: ctrl}
	mock.recorder = &MockIngressClassesGetterMockRecorder{mock}
	return mock
}

func (m *MockIngressClassesGetter) EXPECT() *MockIngressClassesGetterMockRecorder {
	return m.recorder
}

func (m *MockIngressClassesGetter) IngressClasses() v12.IngressClassInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.IngressClasses)
	ret0, _ := ret[0].(v12.IngressClassInterface)
	return ret0
}

func (mr *MockIngressClassesGetterMockRecorder) IngressClasses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.IngressClasses, reflect.TypeOf((*MockIngressClassesGetter)(nil).IngressClasses))
}

type MockIngressClassInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockIngressClassInterfaceMockRecorder
}

type MockIngressClassInterfaceMockRecorder struct {
	mock *MockIngressClassInterface
}

func NewMockIngressClassInterface(ctrl *gomock.Controller) *MockIngressClassInterface {
	mock := &MockIngressClassInterface{ctrl: ctrl}
	mock.recorder = &MockIngressClassInterfaceMockRecorder{mock}
	return mock
}

func (m *MockIngressClassInterface) EXPECT() *MockIngressClassInterfaceMockRecorder {
	return m.recorder
}

func (m *MockIngressClassInterface) Apply(arg0 context.Context, arg1 *v11.IngressClassApplyConfiguration, arg2 v10.ApplyOptions) (*v1.IngressClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.IngressClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIngressClassInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockIngressClassInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockIngressClassInterface) Create(arg0 context.Context, arg1 *v1.IngressClass, arg2 v10.CreateOptions) (*v1.IngressClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.IngressClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIngressClassInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockIngressClassInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockIngressClassInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockIngressClassInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockIngressClassInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockIngressClassInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockIngressClassInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockIngressClassInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockIngressClassInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.IngressClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.IngressClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIngressClassInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockIngressClassInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockIngressClassInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.IngressClassList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.IngressClassList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIngressClassInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockIngressClassInterface)(nil).List), arg0, arg1)
}

func (m *MockIngressClassInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.IngressClass, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.IngressClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIngressClassInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockIngressClassInterface)(nil).Patch), varargs...)
}

func (m *MockIngressClassInterface) Update(arg0 context.Context, arg1 *v1.IngressClass, arg2 v10.UpdateOptions) (*v1.IngressClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.IngressClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIngressClassInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockIngressClassInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockIngressClassInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIngressClassInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockIngressClassInterface)(nil).Watch), arg0, arg1)
}
