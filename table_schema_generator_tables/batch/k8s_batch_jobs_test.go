package batch

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-k8s/faker"
	"github.com/selefra/selefra-provider-k8s/k8sTesting"
	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/k8s_client/mocks"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createBatchJobs(t *testing.T, ctrl *gomock.Controller) k8s_client.K8sServices {
	jobs := mocks.NewMockJobsClient(ctrl)
	j := batchv1.Job{}
	if err := faker.FakeObject(&j); err != nil {
		t.Fatal(err)
	}
	j.Spec.Template = k8sTesting.FakePodTemplateSpec(t)
	j.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	j.Spec.Template.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	jobs.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&batchv1.JobList{Items: []batchv1.Job{j}}, nil,
	)
	return k8s_client.K8sServices{
		Jobs: jobs,
	}
}

func TestBatchJobs(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sBatchJobsGenerator{}), createBatchJobs, k8s_client.TestOptions{})
}
