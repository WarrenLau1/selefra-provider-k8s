# Table: k8s_batch_jobs

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| spec_ttl_seconds_after_finished | int | X | √ |  | 
| api_version | string | X | √ |  | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_parallelism | int | X | √ |  | 
| context | string | X | √ |  | 
| spec_template | json | X | √ |  | 
| status_succeeded | int | X | √ |  | 
| kind | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| spec_backoff_limit | int | X | √ |  | 
| status_failed | int | X | √ |  | 
| resource_version | string | X | √ |  | 
| spec_pod_failure_policy | json | X | √ |  | 
| spec_completion_mode | string | X | √ |  | 
| status_conditions | json | X | √ |  | 
| generation | int | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_suspend | bool | X | √ |  | 
| status_active | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| spec_completions | int | X | √ |  | 
| spec_active_deadline_seconds | int | X | √ |  | 
| spec_manual_selector | bool | X | √ |  | 
| status_start_time | timestamp | X | √ |  | 
| uid | string | √ | √ |  | 
| name | string | X | √ |  | 
| spec_selector | json | X | √ |  | 
| status_uncounted_terminated_pods | json | X | √ |  | 
| status_ready | int | X | √ |  | 
| namespace | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| status_completion_time | timestamp | X | √ |  | 
| status_completed_indexes | string | X | √ |  | 


