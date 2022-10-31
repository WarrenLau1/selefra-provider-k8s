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

func fakeCronJob(t *testing.T) batchv1.CronJob {
	var job batchv1.CronJob
	if err := faker.FakeObject(&job); err != nil {
		t.Fatal(err)
	}
	job.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	job.Spec.JobTemplate.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	job.Spec.JobTemplate.Spec.Template = k8sTesting.FakePodTemplateSpec(t)
	return job
}

func createBatchCronJobs(t *testing.T, ctrl *gomock.Controller) k8s_client.K8sServices {
	cronJobs := mocks.NewMockCronJobsClient(ctrl)

	cronJobs.EXPECT().List(gomock.Any(), metav1.ListOptions{}).AnyTimes().Return(
		&batchv1.CronJobList{Items: []batchv1.CronJob{fakeCronJob(t)}},
		nil,
	)
	return k8s_client.K8sServices{
		CronJobs: cronJobs,
	}
}

func TestBatchCronJobs(t *testing.T) {
	k8s_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableK8sBatchCronJobsGenerator{}), createBatchCronJobs, k8s_client.TestOptions{})
}
