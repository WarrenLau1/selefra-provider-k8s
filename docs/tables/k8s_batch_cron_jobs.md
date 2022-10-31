# Table: k8s_batch_cron_jobs

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| context | string | X | √ |  | 
| labels | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_job_template | json | X | √ |  | 
| status_last_successful_time | timestamp | X | √ |  | 
| spec_failed_jobs_history_limit | int | X | √ |  | 
| status_last_schedule_time | timestamp | X | √ |  | 
| uid | string | √ | √ |  | 
| api_version | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_time_zone | string | X | √ |  | 
| spec_starting_deadline_seconds | int | X | √ |  | 
| status_active | json | X | √ |  | 
| spec_suspend | bool | X | √ |  | 
| spec_successful_jobs_history_limit | int | X | √ |  | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| namespace | string | X | √ |  | 
| generation | int | X | √ |  | 
| spec_schedule | string | X | √ |  | 
| spec_concurrency_policy | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


