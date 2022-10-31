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

func createCoreServiceAccounts(t *testing.T, ctrl *gomock.Controller) k8s_client.K8sServices {
	serviceAccounts := mocks.NewMockServiceAccountsClient(ctrl)
	e := corev1.ServiceAccount{}
	if err := faker.FakeObject(&e); err != nil {
		t.Fatal(err)
	}
	e.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	serviceAccounts.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&corev1.ServiceAccountList{Items: []corev1.ServiceAccount{e}}, nil,
	)
	return k8s_client.K8sServices{
		ServiceAccounts: serviceAccounts,
	}
}

func TestCoreServiceAccounts(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sCoreServiceAccountsGenerator{}), createCoreServiceAccounts, k8s_client.TestOptions{})
}
