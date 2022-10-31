package rbac

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-k8s/faker"
	"github.com/selefra/selefra-provider-k8s/k8sTesting"
	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/k8s_client/mocks"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func fakeRoleBinding(t *testing.T) *v1.RoleBinding {
	r := v1.RoleBinding{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}
	r.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	return &r
}

func createRbacRoleBindings(t *testing.T, ctrl *gomock.Controller) k8s_client.K8sServices {
	roles := mocks.NewMockRoleBindingsClient(ctrl)
	roles.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&v1.RoleBindingList{Items: []v1.RoleBinding{*fakeRoleBinding(t)}}, nil,
	)
	return k8s_client.K8sServices{
		RoleBindings: roles,
	}
}

func TestRbacRoleBindings(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sRbacRoleBindingsGenerator{}), createRbacRoleBindings, k8s_client.TestOptions{})
}
