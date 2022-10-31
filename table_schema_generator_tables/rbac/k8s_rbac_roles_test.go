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

func fakeRole(t *testing.T) *v1.Role {
	r := v1.Role{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}
	r.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	return &r
}

func createRbacRoles(t *testing.T, ctrl *gomock.Controller) k8s_client.K8sServices {
	roles := mocks.NewMockRolesClient(ctrl)
	roles.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&v1.RoleList{Items: []v1.Role{*fakeRole(t)}}, nil,
	)
	return k8s_client.K8sServices{
		Roles: roles,
	}
}

func TestRbacRoles(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sRbacRolesGenerator{}), createRbacRoles, k8s_client.TestOptions{})
}
