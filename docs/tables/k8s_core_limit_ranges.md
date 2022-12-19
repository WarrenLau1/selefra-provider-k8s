# Table: k8s_core_limit_ranges

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| generation | int | X | √ |  | 
| labels | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| spec_limits | json | X | √ |  | 
| context | string | X | √ |  | 
| uid | string | √ | √ |  | 
| namespace | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| kind | string | X | √ |  | 
| api_version | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| annotations | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


