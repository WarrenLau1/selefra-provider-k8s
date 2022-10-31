# Table: k8s_core_nodes

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| labels | json | X | √ |  | 
| status_phase | string | X | √ |  | 
| status_conditions | json | X | √ |  | 
| status_images | json | X | √ |  | 
| status_config | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| uid | string | √ | √ |  | 
| generation | int | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| context | string | X | √ |  | 
| spec_provider_id | string | X | √ |  | 
| status_allocatable | json | X | √ |  | 
| api_version | string | X | √ |  | 
| namespace | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| status_capacity | json | X | √ |  | 
| status_daemon_endpoints | json | X | √ |  | 
| spec_unschedulable | bool | X | √ |  | 
| status_volumes_in_use | string_array | X | √ |  | 
| spec_pod_cidr | cidr | X | √ |  | 
| kind | string | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_taints | json | X | √ |  | 
| resource_version | string | X | √ |  | 
| spec_config_source | json | X | √ |  | 
| status_node_info | json | X | √ |  | 
| spec_pod_cidrs | cidr_array | X | √ |  | 
| name | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| status_addresses | json | X | √ |  | 
| status_volumes_attached | json | X | √ |  | 


