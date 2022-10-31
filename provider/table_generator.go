package provider

import (
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/apps"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/batch"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/core"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/networking"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator_tables/rbac"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&apps.TableK8sAppsReplicaSetsGenerator{}),
		table_schema_generator.GenTableSchema(&apps.TableK8sAppsStatefulSetsGenerator{}),
		table_schema_generator.GenTableSchema(&apps.TableK8sAppsDaemonSetsGenerator{}),
		table_schema_generator.GenTableSchema(&apps.TableK8sAppsDeploymentsGenerator{}),
		table_schema_generator.GenTableSchema(&batch.TableK8sBatchCronJobsGenerator{}),
		table_schema_generator.GenTableSchema(&batch.TableK8sBatchJobsGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreLimitRangesGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreResourceQuotasGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreNamespacesGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreNodesGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreServiceAccountsGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreServicesGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCorePodsGenerator{}),
		table_schema_generator.GenTableSchema(&core.TableK8sCoreEndpointsGenerator{}),
		table_schema_generator.GenTableSchema(&networking.TableK8sNetworkingNetworkPoliciesGenerator{}),
		table_schema_generator.GenTableSchema(&rbac.TableK8sRbacRoleBindingsGenerator{}),
		table_schema_generator.GenTableSchema(&rbac.TableK8sRbacRolesGenerator{}),
	}
}
