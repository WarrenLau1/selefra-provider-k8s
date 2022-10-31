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

func createCoreNamespace(t *testing.T, ctrl *gomock.Controller) k8s_client.K8sServices {
	s := mocks.NewMockNamespacesClient(ctrl)
	var namespace corev1.Namespace
	if err := faker.FakeObject(&namespace); err != nil {
		t.Fatal(err)
	}
	namespace.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	s.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&corev1.NamespaceList{Items: []corev1.Namespace{namespace}}, nil,
	)
	return k8s_client.K8sServices{
		Namespaces: s,
	}
}

func TestCoreNamespaces(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sCoreNamespacesGenerator{}), createCoreNamespace, k8s_client.TestOptions{})
}
