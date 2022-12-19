# Table: k8s_core_namespaces

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| owner_references | json | X | √ |  | 
| status_phase | string | X | √ |  | 
| uid | string | √ | √ |  | 
| kind | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| namespace | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_finalizers | string_array | X | √ |  | 
| status_conditions | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| context | string | X | √ |  | 
| api_version | string | X | √ |  | 
| generation | int | X | √ |  | 
| name | string | X | √ |  | 
| labels | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 


