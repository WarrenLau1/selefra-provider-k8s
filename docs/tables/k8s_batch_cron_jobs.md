# Table: k8s_batch_cron_jobs

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| uid | string | √ | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_suspend | bool | X | √ |  | 
| status_active | json | X | √ |  | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| namespace | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_schedule | string | X | √ |  | 
| status_last_schedule_time | timestamp | X | √ |  | 
| spec_successful_jobs_history_limit | int | X | √ |  | 
| status_last_successful_time | timestamp | X | √ |  | 
| context | string | X | √ |  | 
| api_version | string | X | √ |  | 
| labels | json | X | √ |  | 
| spec_time_zone | string | X | √ |  | 
| spec_starting_deadline_seconds | int | X | √ |  | 
| spec_job_template | json | X | √ |  | 
| generation | int | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_concurrency_policy | string | X | √ |  | 
| spec_failed_jobs_history_limit | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


