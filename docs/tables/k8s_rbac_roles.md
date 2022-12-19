# Table: k8s_rbac_roles

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| uid | string | √ | √ |  | 
| annotations | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| api_version | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| labels | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| context | string | X | √ |  | 
| name | string | X | √ |  | 
| generation | int | X | √ |  | 
| rules | json | X | √ |  | 
| kind | string | X | √ |  | 
| namespace | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


