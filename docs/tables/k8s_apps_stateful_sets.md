# Table: k8s_apps_stateful_sets

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| annotations | json | X | √ |  | 
| name | string | X | √ |  | 
| uid | string | √ | √ |  | 
| namespace | string | X | √ |  | 
| generation | int | X | √ |  | 
| spec_volume_claim_templates | json | X | √ |  | 
| status_observed_generation | int | X | √ |  | 
| context | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_template | json | X | √ |  | 
| spec_pod_management_policy | string | X | √ |  | 
| spec_revision_history_limit | int | X | √ |  | 
| status_ready_replicas | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| kind | string | X | √ |  | 
| spec_service_name | string | X | √ |  | 
| status_current_replicas | int | X | √ |  | 
| status_update_revision | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| spec_persistent_volume_claim_retention_policy | json | X | √ |  | 
| labels | json | X | √ |  | 
| spec_replicas | int | X | √ |  | 
| spec_selector | json | X | √ |  | 
| spec_update_strategy | json | X | √ |  | 
| status_updated_replicas | int | X | √ |  | 
| status_conditions | json | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| spec_min_ready_seconds | int | X | √ |  | 
| status_collision_count | int | X | √ |  | 
| status_available_replicas | int | X | √ |  | 
| api_version | string | X | √ |  | 
| status_replicas | int | X | √ |  | 
| status_current_revision | string | X | √ |  | 
| owner_references | json | X | √ |  | 


