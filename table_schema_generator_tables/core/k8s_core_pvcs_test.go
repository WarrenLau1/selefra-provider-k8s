package core







import (


	"github.com/selefra/selefra-provider-k8s/constants"




	"testing"

	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-k8s/faker"
	"github.com/selefra/selefra-provider-k8s/k8s_client"


	"github.com/selefra/selefra-provider-k8s/mocks"


	resourcemock "github.com/selefra/selefra-provider-k8s/mocks/core/v1"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	resource "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"




	"k8s.io/client-go/kubernetes"
)





func createPvcs(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.PersistentVolumeClaim{}


	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)


	}



	resourceClient := resourcemock.NewMockPersistentVolumeClaimInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&resource.PersistentVolumeClaimList{Items: []resource.PersistentVolumeClaim{r}}, nil,


	)









	serviceClient := resourcemock.NewMockCoreV1Interface(ctrl)



	serviceClient.EXPECT().PersistentVolumeClaims(constants.Constants_34).AnyTimes().Return(resourceClient)



	cl := mocks.NewMockInterface(ctrl)


	cl.EXPECT().CoreV1().AnyTimes().Return(serviceClient)





	return cl
}





func TestPvcs(t *testing.T) {




	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sCorePvcsGenerator{}), createPvcs, k8s_client.TestOptions{})


}


