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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func fakeStatefulSet(t *testing.T) appsv1.StatefulSet {
	var rs appsv1.StatefulSet
	if err := faker.FakeObject(&rs); err != nil {
		t.Fatal(err)
	}
	intOrStr := intstr.FromInt(100)
	rs.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable = &intOrStr
	rs.Spec.VolumeClaimTemplates = []corev1.PersistentVolumeClaim{*k8sTesting.FakePersistentVolumeClaim(t)}
	rs.Spec.Selector = k8sTesting.FakeSelector(t)
	rs.Spec.Template = k8sTesting.FakePodTemplateSpec(t)
	rs.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	return rs
}

func createStatefulSets(t *testing.T, ctrl *gomock.Controller) k8s_client.K8sServices {
	setsClient := mocks.NewMockStatefulSetsClient(ctrl)

	setsClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&appsv1.StatefulSetList{Items: []appsv1.StatefulSet{fakeStatefulSet(t)}}, nil,
	)
	return k8s_client.K8sServices{
		StatefulSets: setsClient,
	}
}

func TestAppsStatefulSets(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sAppsStatefulSetsGenerator{}), createStatefulSets, k8s_client.TestOptions{})
}
