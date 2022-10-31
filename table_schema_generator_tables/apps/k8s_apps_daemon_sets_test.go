package apps

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-k8s/k8sTesting"
	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/k8s_client/mocks"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createAppsDaemonSets(t *testing.T, ctrl *gomock.Controller) k8s_client.K8sServices {
	daemonSetsClient := mocks.NewMockDaemonSetsClient(ctrl)
	daemonSetsClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&appsv1.DaemonSetList{Items: []appsv1.DaemonSet{k8sTesting.FakeDaemonSet(t)}}, nil,
	)
	return k8s_client.K8sServices{
		DaemonSets: daemonSetsClient,
	}
}

func TestDaemonSets(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sAppsDaemonSetsGenerator{}), createAppsDaemonSets, k8s_client.TestOptions{})
}
