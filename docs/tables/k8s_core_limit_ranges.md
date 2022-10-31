# Table: k8s_core_limit_ranges

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| context | string | X | √ |  | 
| annotations | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| spec_limits | json | X | √ |  | 
| kind | string | X | √ |  | 
| generation | int | X | √ |  | 
| labels | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| resource_version | string | X | √ |  | 
| uid | string | √ | √ |  | 
| api_version | string | X | √ |  | 
| name | string | X | √ |  | 
| namespace | string | X | √ |  | 


