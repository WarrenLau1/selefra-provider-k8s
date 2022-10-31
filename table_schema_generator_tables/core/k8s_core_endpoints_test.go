package core

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-k8s/faker"
	"github.com/selefra/selefra-provider-k8s/k8sTesting"
	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/k8s_client/mocks"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createCoreEndpoints(t *testing.T, ctrl *gomock.Controller) k8s_client.K8sServices {
	endpoints := mocks.NewMockEndpointsClient(ctrl)
	e := corev1.Endpoints{}
	if err := faker.FakeObject(&e); err != nil {
		t.Fatal(err)
	}
	e.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	subset := corev1.EndpointSubset{}
	if err := faker.FakeObject(&subset); err != nil {
		t.Fatal(err)
	}
	address := corev1.EndpointAddress{}
	if err := faker.FakeObject(&address); err != nil {
		t.Fatal(err)
	}
	address.IP = "127.0.0.1"
	subset.Addresses = []corev1.EndpointAddress{address}
	subset.NotReadyAddresses = []corev1.EndpointAddress{address}
	e.Subsets = []corev1.EndpointSubset{subset}
	endpoints.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&corev1.EndpointsList{Items: []corev1.Endpoints{e}}, nil,
	)
	return k8s_client.K8sServices{
		Endpoints: endpoints,
	}
}

func TestCoreEndpoints(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sCoreEndpointsGenerator{}), createCoreEndpoints, k8s_client.TestOptions{})
}
