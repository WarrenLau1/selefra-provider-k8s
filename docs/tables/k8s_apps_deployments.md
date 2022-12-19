# Table: k8s_apps_deployments

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| namespace | string | X | √ |  | 
| annotations | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_min_ready_seconds | int | X | √ |  | 
| status_collision_count | int | X | √ |  | 
| context | string | X | √ |  | 
| kind | string | X | √ |  | 
| api_version | string | X | √ |  | 
| spec_replicas | int | X | √ |  | 
| spec_template | json | X | √ |  | 
| status_replicas | int | X | √ |  | 
| status_conditions | json | X | √ |  | 
| generation | int | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_strategy | json | X | √ |  | 
| status_unavailable_replicas | int | X | √ |  | 
| name | string | X | √ |  | 
| spec_revision_history_limit | int | X | √ |  | 
| spec_paused | bool | X | √ |  | 
| status_ready_replicas | int | X | √ |  | 
| labels | json | X | √ |  | 
| status_observed_generation | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| uid | string | √ | √ |  | 
| resource_version | string | X | √ |  | 
| spec_selector | json | X | √ |  | 
| status_updated_replicas | int | X | √ |  | 
| status_available_replicas | int | X | √ |  | 
| spec_progress_deadline_seconds | int | X | √ |  | 


