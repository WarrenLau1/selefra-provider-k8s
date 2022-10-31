# Table: k8s_apps_deployments

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| context | string | X | √ |  | 
| spec_paused | bool | X | √ |  | 
| spec_progress_deadline_seconds | int | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_revision_history_limit | int | X | √ |  | 
| api_version | string | X | √ |  | 
| annotations | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| status_replicas | int | X | √ |  | 
| status_available_replicas | int | X | √ |  | 
| status_collision_count | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| resource_version | string | X | √ |  | 
| generation | int | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| labels | json | X | √ |  | 
| spec_selector | json | X | √ |  | 
| spec_template | json | X | √ |  | 
| status_observed_generation | int | X | √ |  | 
| status_ready_replicas | int | X | √ |  | 
| status_unavailable_replicas | int | X | √ |  | 
| status_conditions | json | X | √ |  | 
| uid | string | √ | √ |  | 
| spec_replicas | int | X | √ |  | 
| spec_strategy | json | X | √ |  | 
| spec_min_ready_seconds | int | X | √ |  | 
| namespace | string | X | √ |  | 
| status_updated_replicas | int | X | √ |  | 


