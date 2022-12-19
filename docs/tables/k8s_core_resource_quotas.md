# Table: k8s_core_resource_quotas

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| labels | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| status_hard | json | X | √ |  | 
| status_used | json | X | √ |  | 
| context | string | X | √ |  | 
| api_version | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| spec_hard | json | X | √ |  | 
| spec_scopes | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| uid | string | √ | √ |  | 
| name | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| generation | int | X | √ |  | 
| kind | string | X | √ |  | 
| namespace | string | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_scope_selector | json | X | √ |  | 


