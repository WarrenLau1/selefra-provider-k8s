package k8s_client

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type K8sServices struct {
	Client *kubernetes.Clientset

	CronJobs        CronJobsClient
	DaemonSets      DaemonSetsClient
	Deployments     DeploymentsClient
	Endpoints       EndpointsClient
	Jobs            JobsClient
	LimitRanges     LimitRangesClient
	Namespaces      NamespacesClient
	NetworkPolicies NetworkPoliciesClient
	Nodes           NodesClient
	Pods            PodsClient
	ReplicaSets     ReplicaSetsClient
	ResourceQuotas  ResourceQuotasClient
	RoleBindings    RoleBindingsClient
	Roles           RolesClient
	ServiceAccounts ServiceAccountsClient
	K8sServices     ServicesClient
	StatefulSets    StatefulSetsClient
	//PolicyV1beta1   *policyv1beta1.PolicyV1beta1Client
}

type CronJobsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*batchv1.CronJobList, error)
}

type DaemonSetsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*appsv1.DaemonSetList, error)
}

type DeploymentsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*appsv1.DeploymentList, error)
}

type EndpointsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.EndpointsList, error)
}

type JobsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*batchv1.JobList, error)
}

type LimitRangesClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.LimitRangeList, error)
}

type NamespacesClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.NamespaceList, error)
}

type NetworkPoliciesClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*networkingv1.NetworkPolicyList, error)
}

type NodesClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.NodeList, error)
}

type PodsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.PodList, error)
}

type ReplicaSetsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*appsv1.ReplicaSetList, error)
}

type ResourceQuotasClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.ResourceQuotaList, error)
}

type RoleBindingsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*rbacv1.RoleBindingList, error)
}

type RolesClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*rbacv1.RoleList, error)
}

type ServiceAccountsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.ServiceAccountList, error)
}

type ServicesClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.ServiceList, error)
}

type StatefulSetsClient interface {
	List(ctx context.Context, opts metav1.ListOptions) (*appsv1.StatefulSetList, error)
}
