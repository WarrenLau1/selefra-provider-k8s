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
	"k8s.io/apimachinery/pkg/util/intstr"
)

func fakeAppsDeployment(t *testing.T) appsv1.Deployment {
	var deployment appsv1.Deployment
	if err := faker.FakeObject(&deployment); err != nil {
		t.Fatal(err)
	}

	intOrStr1 := intstr.FromInt(100)
	intOrStr2 := intstr.FromInt(100)
	deployment.Spec.Strategy.RollingUpdate.MaxSurge = &intOrStr1
	deployment.Spec.Strategy.RollingUpdate.MaxUnavailable = &intOrStr2

	deployment.Spec.Template = k8sTesting.FakePodTemplateSpec(t)
	deployment.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	return deployment
}

func createDeployments(t *testing.T, ctrl *gomock.Controller) k8s_client.K8sServices {
	deploymentsClient := mocks.NewMockDeploymentsClient(ctrl)
	deploymentsClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&appsv1.DeploymentList{Items: []appsv1.Deployment{fakeAppsDeployment(t)}}, nil,
	)
	return k8s_client.K8sServices{
		Deployments: deploymentsClient,
	}
}

func TestDeployments(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sAppsDeploymentsGenerator{}), createDeployments, k8s_client.TestOptions{})
}
