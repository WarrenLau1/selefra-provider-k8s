# Table: k8s_rbac_roles

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| uid | string | √ | √ |  | 
| api_version | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| labels | json | X | √ |  | 
| owner_references | json | X | √ |  | 
| rules | json | X | √ |  | 
| kind | string | X | √ |  | 
| name | string | X | √ |  | 
| namespace | string | X | √ |  | 
| generation | int | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| finalizers | string_array | X | √ |  | 
| context | string | X | √ |  | 
| annotations | json | X | √ |  | 


