# Table: k8s_batch_jobs

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| labels | json | X | √ |  | 
| status_start_time | timestamp | X | √ |  | 
| spec_pod_failure_policy | json | X | √ |  | 
| spec_selector | json | X | √ |  | 
| spec_template | json | X | √ |  | 
| spec_ttl_seconds_after_finished | int | X | √ |  | 
| spec_completion_mode | string | X | √ |  | 
| context | string | X | √ |  | 
| generation | int | X | √ |  | 
| spec_parallelism | int | X | √ |  | 
| status_uncounted_terminated_pods | json | X | √ |  | 
| namespace | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_backoff_limit | int | X | √ |  | 
| spec_manual_selector | bool | X | √ |  | 
| kind | string | X | √ |  | 
| api_version | string | X | √ |  | 
| status_completion_time | timestamp | X | √ |  | 
| uid | string | √ | √ |  | 
| name | string | X | √ |  | 
| status_completed_indexes | string | X | √ |  | 
| spec_suspend | bool | X | √ |  | 
| status_failed | int | X | √ |  | 
| status_succeeded | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| resource_version | string | X | √ |  | 
| spec_active_deadline_seconds | int | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_completions | int | X | √ |  | 
| status_conditions | json | X | √ |  | 
| status_active | int | X | √ |  | 
| status_ready | int | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| annotations | json | X | √ |  | 


