# Table: k8s_rbac_role_bindings

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| subjects | json | X | √ |  | 
| context | string | X | √ |  | 
| generation | int | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| role_ref | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| namespace | string | X | √ |  | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| uid | string | √ | √ |  | 
| api_version | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| owner_references | json | X | √ |  | 


