# Table: k8s_core_pods

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| spec_active_deadline_seconds | int | X | √ |  | 
| spec_share_process_namespace | bool | X | √ |  | 
| spec_image_pull_secrets | json | X | √ |  | 
| status_message | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| spec_restart_policy | string | X | √ |  | 
| spec_termination_grace_period_seconds | int | X | √ |  | 
| uid | string | √ | √ |  | 
| spec_containers | json | X | √ |  | 
| status_phase | string | X | √ |  | 
| status_conditions | json | X | √ |  | 
| status_pod_ips | ip_array | X | √ |  | 
| spec_host_network | bool | X | √ |  | 
| spec_hostname | string | X | √ |  | 
| status_nominated_node_name | string | X | √ |  | 
| api_version | string | X | √ |  | 
| spec_init_containers | json | X | √ |  | 
| spec_overhead | json | X | √ |  | 
| spec_ephemeral_containers | json | X | √ |  | 
| spec_automount_service_account_token | bool | X | √ |  | 
| status_init_container_statuses | json | X | √ |  | 
| namespace | string | X | √ |  | 
| generation | int | X | √ |  | 
| labels | json | X | √ |  | 
| spec_priority | int | X | √ |  | 
| spec_topology_spread_constraints | json | X | √ |  | 
| status_reason | string | X | √ |  | 
| name | string | X | √ |  | 
| spec_tolerations | json | X | √ |  | 
| spec_priority_class_name | string | X | √ |  | 
| spec_scheduler_name | string | X | √ |  | 
| spec_preemption_policy | string | X | √ |  | 
| annotations | json | X | √ |  | 
| status_container_statuses | json | X | √ |  | 
| spec_readiness_gates | json | X | √ |  | 
| context | string | X | √ |  | 
| status_pod_ip | ip | X | √ |  | 
| spec_subdomain | string | X | √ |  | 
| spec_runtime_class_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| spec_security_context | json | X | √ |  | 
| spec_host_aliases | json | X | √ |  | 
| spec_dns_config | json | X | √ |  | 
| spec_host_ipc | bool | X | √ |  | 
| spec_os | json | X | √ |  | 
| status_qos_class | string | X | √ |  | 
| kind | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_host_pid | bool | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_enable_service_links | bool | X | √ |  | 
| spec_set_hostname_as_fqdn | bool | X | √ |  | 
| spec_host_users | bool | X | √ |  | 
| status_ephemeral_container_statuses | json | X | √ |  | 
| resource_version | string | X | √ |  | 
| spec_service_account_name | string | X | √ |  | 
| spec_affinity | json | X | √ |  | 
| spec_node_selector | json | X | √ |  | 
| spec_node_name | string | X | √ |  | 
| status_start_time | timestamp | X | √ |  | 
| status_host_ip | ip | X | √ |  | 
| spec_volumes | json | X | √ |  | 
| spec_dns_policy | string | X | √ |  | 


