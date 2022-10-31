# Table: k8s_core_namespaces

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resource_version | string | X | √ |  | 
| labels | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| uid | string | √ | √ |  | 
| namespace | string | X | √ |  | 
| spec_finalizers | string_array | X | √ |  | 
| name | string | X | √ |  | 
| generation | int | X | √ |  | 
| api_version | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| context | string | X | √ |  | 
| kind | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| status_phase | string | X | √ |  | 
| status_conditions | json | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| annotations | json | X | √ |  | 


