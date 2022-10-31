# Table: k8s_apps_stateful_sets

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| status_available_replicas | int | X | √ |  | 
| spec_update_strategy | json | X | √ |  | 
| status_current_replicas | int | X | √ |  | 
| status_updated_replicas | int | X | √ |  | 
| spec_selector | json | X | √ |  | 
| spec_revision_history_limit | int | X | √ |  | 
| spec_min_ready_seconds | int | X | √ |  | 
| status_update_revision | string | X | √ |  | 
| status_conditions | json | X | √ |  | 
| labels | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_replicas | int | X | √ |  | 
| namespace | string | X | √ |  | 
| generation | int | X | √ |  | 
| status_ready_replicas | int | X | √ |  | 
| annotations | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_service_name | string | X | √ |  | 
| status_replicas | int | X | √ |  | 
| status_current_revision | string | X | √ |  | 
| kind | string | X | √ |  | 
| api_version | string | X | √ |  | 
| name | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| status_collision_count | int | X | √ |  | 
| uid | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| spec_template | json | X | √ |  | 
| spec_pod_management_policy | string | X | √ |  | 
| status_observed_generation | int | X | √ |  | 
| context | string | X | √ |  | 
| spec_volume_claim_templates | json | X | √ |  | 
| spec_persistent_volume_claim_retention_policy | json | X | √ |  | 


