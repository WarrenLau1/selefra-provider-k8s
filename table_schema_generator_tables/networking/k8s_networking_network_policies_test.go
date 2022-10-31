package networking

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-k8s/faker"
	"github.com/selefra/selefra-provider-k8s/k8sTesting"
	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/k8s_client/mocks"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func createNetworkingNetworkPolicies(t *testing.T, ctrl *gomock.Controller) k8s_client.K8sServices {
	s := mocks.NewMockNetworkPoliciesClient(ctrl)
	var networkPolicy networkingv1.NetworkPolicy
	if err := faker.FakeObject(&networkPolicy); err != nil {
		t.Fatal(err)
	}
	networkPolicy.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	intStr1 := intstr.FromInt(100)
	intStr2 := intstr.FromInt(200)
	networkPolicy.Spec.Ingress[0].Ports[0].Port = &intStr1
	networkPolicy.Spec.Egress[0].Ports[0].Port = &intStr2

	s.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&networkingv1.NetworkPolicyList{Items: []networkingv1.NetworkPolicy{networkPolicy}}, nil,
	)
	return k8s_client.K8sServices{
		NetworkPolicies: s,
	}
}

func TestNetworkingNetworkPolicies(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sNetworkingNetworkPoliciesGenerator{}), createNetworkingNetworkPolicies, k8s_client.TestOptions{})
}
