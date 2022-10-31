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

func createCoreResourceQuotas(t *testing.T, ctrl *gomock.Controller) k8s_client.K8sServices {
	resourceQuotas := mocks.NewMockResourceQuotasClient(ctrl)
	e := corev1.ResourceQuota{}
	if err := faker.FakeObject(&e); err != nil {
		t.Fatal(err)
	}
	ss := corev1.ScopeSelector{}
	if err := faker.FakeObject(&ss); err != nil {
		t.Fatal(err)
	}
	rqsp := corev1.ResourceQuotaSpec{
		Hard:		*k8sTesting.FakeResourceList(t),
		Scopes:		[]corev1.ResourceQuotaScope{corev1.ResourceQuotaScopeBestEffort},
		ScopeSelector:	&ss,
	}
	rqst := corev1.ResourceQuotaStatus{
		Hard:	*k8sTesting.FakeResourceList(t),
		Used:	*k8sTesting.FakeResourceList(t),
	}
	e.Spec = rqsp
	e.Status = rqst
	e.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	resourceQuotas.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&corev1.ResourceQuotaList{Items: []corev1.ResourceQuota{e}}, nil,
	)
	return k8s_client.K8sServices{
		ResourceQuotas: resourceQuotas,
	}
}

func TestCoreResourceQuotas(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sCoreResourceQuotasGenerator{}), createCoreResourceQuotas, k8s_client.TestOptions{})
}
