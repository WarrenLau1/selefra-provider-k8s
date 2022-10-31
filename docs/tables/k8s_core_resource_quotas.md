# Table: k8s_core_resource_quotas

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| context | string | X | √ |  | 
| kind | string | X | √ |  | 
| generation | int | X | √ |  | 
| api_version | string | X | √ |  | 
| name | string | X | √ |  | 
| spec_scope_selector | json | X | √ |  | 
| status_used | json | X | √ |  | 
| spec_scopes | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| namespace | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_hard | json | X | √ |  | 
| uid | string | √ | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| owner_references | json | X | √ |  | 
| status_hard | json | X | √ |  | 


