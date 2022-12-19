# Table: k8s_core_nodes

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| spec_provider_id | string | X | √ |  | 
| status_capacity | json | X | √ |  | 
| status_addresses | json | X | √ |  | 
| status_volumes_in_use | string_array | X | √ |  | 
| uid | string | √ | √ |  | 
| spec_config_source | json | X | √ |  | 
| namespace | string | X | √ |  | 
| spec_unschedulable | bool | X | √ |  | 
| kind | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| owner_references | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| status_allocatable | json | X | √ |  | 
| status_conditions | json | X | √ |  | 
| status_node_info | json | X | √ |  | 
| api_version | string | X | √ |  | 
| annotations | json | X | √ |  | 
| status_daemon_endpoints | json | X | √ |  | 
| status_config | json | X | √ |  | 
| spec_pod_cidrs | cidr_array | X | √ |  | 
| resource_version | string | X | √ |  | 
| status_phase | string | X | √ |  | 
| status_volumes_attached | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| spec_pod_cidr | cidr | X | √ |  | 
| generation | int | X | √ |  | 
| spec_taints | json | X | √ |  | 
| context | string | X | √ |  | 
| status_images | json | X | √ |  | 
| labels | json | X | √ |  | 


