# Table: k8s_apps_daemon_sets

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| spec_template | json | X | √ |  | 
| status_desired_number_scheduled | int | X | √ |  | 
| status_number_available | int | X | √ |  | 
| status_number_unavailable | int | X | √ |  | 
| annotations | json | X | √ |  | 
| status_number_misscheduled | int | X | √ |  | 
| resource_version | string | X | √ |  | 
| spec_update_strategy | json | X | √ |  | 
| status_current_number_scheduled | int | X | √ |  | 
| context | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_revision_history_limit | int | X | √ |  | 
| api_version | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| status_conditions | json | X | √ |  | 
| status_collision_count | int | X | √ |  | 
| name | string | X | √ |  | 
| namespace | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_selector | json | X | √ |  | 
| status_number_ready | int | X | √ |  | 
| status_observed_generation | int | X | √ |  | 
| status_updated_number_scheduled | int | X | √ |  | 
| uid | string | √ | √ |  | 
| labels | json | X | √ |  | 
| spec_min_ready_seconds | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| kind | string | X | √ |  | 
| generation | int | X | √ |  | 


