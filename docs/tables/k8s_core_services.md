# Table: k8s_core_services

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| namespace | string | X | √ |  | 
| spec_external_traffic_policy | string | X | √ |  | 
| spec_health_check_node_port | int | X | √ |  | 
| spec_allocate_load_balancer_node_ports | bool | X | √ |  | 
| spec_external_name | string | X | √ |  | 
| spec_ip_families | string_array | X | √ |  | 
| context | string | X | √ |  | 
| spec_cluster_ip | ip | X | √ |  | 
| spec_cluster_ips | ip_array | X | √ |  | 
| owner_references | json | X | √ |  | 
| api_version | string | X | √ |  | 
| labels | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| annotations | json | X | √ |  | 
| spec_publish_not_ready_addresses | bool | X | √ |  | 
| spec_internal_traffic_policy | string | X | √ |  | 
| status_conditions | json | X | √ |  | 
| spec_load_balancer_source_ranges | string_array | X | √ |  | 
| spec_ip_family_policy | string | X | √ |  | 
| spec_load_balancer_class | string | X | √ |  | 
| spec_load_balancer_ip | ip | X | √ |  | 
| kind | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_selector | json | X | √ |  | 
| resource_version | string | X | √ |  | 
| spec_type | string | X | √ |  | 
| spec_session_affinity | string | X | √ |  | 
| spec_external_ips | ip_array | X | √ |  | 
| generation | int | X | √ |  | 
| spec_session_affinity_config | json | X | √ |  | 
| status_load_balancer | json | X | √ |  | 
| uid | string | √ | √ |  | 
| name | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| spec_ports | json | X | √ |  | 


