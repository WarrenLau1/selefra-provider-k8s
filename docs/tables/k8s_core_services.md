# Table: k8s_core_services

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| spec_session_affinity | string | X | √ |  | 
| spec_external_name | string | X | √ |  | 
| spec_ip_families | string_array | X | √ |  | 
| spec_allocate_load_balancer_node_ports | bool | X | √ |  | 
| kind | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| resource_version | string | X | √ |  | 
| spec_health_check_node_port | int | X | √ |  | 
| spec_ip_family_policy | string | X | √ |  | 
| spec_external_ips | ip_array | X | √ |  | 
| api_version | string | X | √ |  | 
| spec_external_traffic_policy | string | X | √ |  | 
| spec_load_balancer_class | string | X | √ |  | 
| spec_load_balancer_ip | ip | X | √ |  | 
| labels | json | X | √ |  | 
| spec_selector | json | X | √ |  | 
| status_conditions | json | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_internal_traffic_policy | string | X | √ |  | 
| namespace | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_session_affinity_config | json | X | √ |  | 
| context | string | X | √ |  | 
| uid | string | √ | √ |  | 
| name | string | X | √ |  | 
| spec_ports | json | X | √ |  | 
| spec_publish_not_ready_addresses | bool | X | √ |  | 
| status_load_balancer | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| spec_cluster_ip | ip | X | √ |  | 
| spec_cluster_ips | ip_array | X | √ |  | 
| spec_load_balancer_source_ranges | string_array | X | √ |  | 
| generation | int | X | √ |  | 
| spec_type | string | X | √ |  | 


