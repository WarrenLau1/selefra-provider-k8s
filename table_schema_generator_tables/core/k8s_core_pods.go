package core

import (
	"context"
	"reflect"

	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"github.com/selefra/selefra-provider-k8s/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-utils/pkg/reflect_util"
	"github.com/songzhibin97/go-ognl"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TableK8sCorePodsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableK8sCorePodsGenerator{}

func (x *TableK8sCorePodsGenerator) GetTableName() string {
	return "k8s_core_pods"
}

func (x *TableK8sCorePodsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableK8sCorePodsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableK8sCorePodsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"uid",
		},
	}
}

func (x *TableK8sCorePodsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			pods := client.(*k8s_client.Client).K8sServices().Pods
			result, err := pods.List(ctx, metav1.ListOptions{})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- result.Items
			return nil
		},
	}
}

func (x *TableK8sCorePodsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return k8s_client.ContextExpand()
}

func (x *TableK8sCorePodsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("spec_active_deadline_seconds").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("Spec.ActiveDeadlineSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_share_process_namespace").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.ShareProcessNamespace")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_image_pull_secrets").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.ImagePullSecrets")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status.Message")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_grace_period_seconds").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("DeletionGracePeriodSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_restart_policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.RestartPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_termination_grace_period_seconds").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("Spec.TerminationGracePeriodSeconds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_containers").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.Containers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_phase").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status.Phase")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_conditions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status.Conditions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_pod_ips").ColumnType(schema.ColumnTypeIpArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_host_network").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.HostNetwork")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_hostname").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.Hostname")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_nominated_node_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status.NominatedNodeName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("APIVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_init_containers").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.InitContainers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_overhead").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.Overhead")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_ephemeral_containers").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.EphemeralContainers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_automount_service_account_token").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.AutomountServiceAccountToken")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_init_container_statuses").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status.InitContainerStatuses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Namespace")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generation").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("Generation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_priority").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("Spec.Priority")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_topology_spread_constraints").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.TopologySpreadConstraints")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status.Reason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_tolerations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.Tolerations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_priority_class_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.PriorityClassName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_scheduler_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.SchedulerName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_preemption_policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.PreemptionPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("annotations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Annotations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_container_statuses").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status.ContainerStatuses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_readiness_gates").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.ReadinessGates")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("context").ColumnType(schema.ColumnTypeString).
			Extractor(k8s_client.ContextExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_pod_ip").ColumnType(schema.ColumnTypeIp).
			Extractor(column_value_extractor.StructSelector("Status.PodIP")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_subdomain").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.Subdomain")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_runtime_class_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.RuntimeClassName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_security_context").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.SecurityContext")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_host_aliases").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.HostAliases")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_dns_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.DNSConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_host_ipc").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.HostIPC")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_os").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.OS")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_qos_class").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status.QOSClass")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_references").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OwnerReferences")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_host_pid").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.HostPID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("finalizers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Finalizers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_enable_service_links").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.EnableServiceLinks")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_set_hostname_as_fqdn").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.SetHostnameAsFQDN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_host_users").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Spec.HostUsers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_ephemeral_container_statuses").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status.EphemeralContainerStatuses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_service_account_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.ServiceAccountName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_affinity").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.Affinity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_node_selector").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.NodeSelector")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_node_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.NodeName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_start_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				if reflect_util.IsNil(result) {
					return nil, nil
				}
				value := ognl.Get(result, "Status.StartTime").Value()
				if reflect_util.IsNil(result) {
					return nil, nil
				}
				of := reflect.ValueOf(value)
				if !of.IsValid() {
					return value, nil
				}
				if of.Kind() == reflect.Pointer {
					of = of.Elem()
					if !of.IsValid() {
						return value, nil
					}
				}
				name := of.MethodByName("Format")
				if !name.IsValid() {
					return value, nil
				}
				call := name.Call([]reflect.Value{reflect.ValueOf("2006-01-02 15:04:05")})
				if len(call) == 0 {
					return nil, nil
				}
				return call[0].Interface(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_host_ip").ColumnType(schema.ColumnTypeIp).
			Extractor(column_value_extractor.StructSelector("Status.HostIP")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_volumes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Spec.Volumes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("spec_dns_policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Spec.DNSPolicy")).Build(),
	}
}

func (x *TableK8sCorePodsGenerator) GetSubTables() []*schema.Table {
	return nil
}
