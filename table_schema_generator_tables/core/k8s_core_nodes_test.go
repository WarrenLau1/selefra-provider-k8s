package core

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-k8s/k8sTesting"
	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/k8s_client/mocks"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createCoreNodes(t *testing.T, ctrl *gomock.Controller) k8s_client.K8sServices {
	nodes := mocks.NewMockNodesClient(ctrl)
	nodes.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&corev1.NodeList{Items: []corev1.Node{k8sTesting.FakeNode(t)}}, nil,
	)
	return k8s_client.K8sServices{
		Nodes: nodes,
	}
}

func TestCoreNodes(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sCoreNodesGenerator{}), createCoreNodes, k8s_client.TestOptions{})
}
