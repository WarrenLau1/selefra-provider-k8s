# Table: k8s_rbac_role_bindings

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| api_version | string | X | √ |  | 
| name | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| subjects | json | X | √ |  | 
| generation | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| kind | string | X | √ |  | 
| namespace | string | X | √ |  | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| role_ref | json | X | √ |  | 
| context | string | X | √ |  | 
| uid | string | √ | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| owner_references | json | X | √ |  | 


