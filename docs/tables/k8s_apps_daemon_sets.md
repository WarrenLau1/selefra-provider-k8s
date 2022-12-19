# Table: k8s_apps_daemon_sets

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| namespace | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| status_number_misscheduled | int | X | √ |  | 
| status_current_number_scheduled | int | X | √ |  | 
| status_desired_number_scheduled | int | X | √ |  | 
| status_observed_generation | int | X | √ |  | 
| status_number_available | int | X | √ |  | 
| status_collision_count | int | X | √ |  | 
| status_conditions | json | X | √ |  | 
| context | string | X | √ |  | 
| uid | string | √ | √ |  | 
| spec_min_ready_seconds | int | X | √ |  | 
| status_updated_number_scheduled | int | X | √ |  | 
| api_version | string | X | √ |  | 
| generation | int | X | √ |  | 
| labels | json | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| spec_selector | json | X | √ |  | 
| spec_update_strategy | json | X | √ |  | 
| status_number_ready | int | X | √ |  | 
| status_number_unavailable | int | X | √ |  | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| annotations | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_template | json | X | √ |  | 
| spec_revision_history_limit | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


