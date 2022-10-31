package apps

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-k8s/faker"
	"github.com/selefra/selefra-provider-k8s/k8sTesting"
	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/k8s_client/mocks"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func fakeReplicaSet(t *testing.T) appsv1.ReplicaSet {
	var rs appsv1.ReplicaSet
	if err := faker.FakeObject(&rs); err != nil {
		t.Fatal(err)
	}
	rs.Spec.Template = k8sTesting.FakePodTemplateSpec(t)
	rs.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}

	return rs
}

func createReplicaSets(t *testing.T, ctrl *gomock.Controller) k8s_client.K8sServices {
	setsClient := mocks.NewMockReplicaSetsClient(ctrl)
	setsClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&appsv1.ReplicaSetList{Items: []appsv1.ReplicaSet{fakeReplicaSet(t)}}, nil,
	)
	return k8s_client.K8sServices{
		ReplicaSets: setsClient,
	}
}

func TestAppsReplicaSets(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sAppsReplicaSetsGenerator{}), createReplicaSets, k8s_client.TestOptions{})
}
