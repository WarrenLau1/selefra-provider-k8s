# Table: k8s_apps_replica_sets

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| namespace | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_replicas | int | X | √ |  | 
| spec_selector | json | X | √ |  | 
| status_observed_generation | int | X | √ |  | 
| context | string | X | √ |  | 
| api_version | string | X | √ |  | 
| name | string | X | √ |  | 
| labels | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_template | json | X | √ |  | 
| status_replicas | int | X | √ |  | 
| status_ready_replicas | int | X | √ |  | 
| status_conditions | json | X | √ |  | 
| uid | string | √ | √ |  | 
| kind | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_min_ready_seconds | int | X | √ |  | 
| status_fully_labeled_replicas | int | X | √ |  | 
| status_available_replicas | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| resource_version | string | X | √ |  | 
| generation | int | X | √ |  | 


