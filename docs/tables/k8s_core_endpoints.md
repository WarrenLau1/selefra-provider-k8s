# Table: k8s_core_endpoints

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| context | string | X | √ |  | 
| name | string | X | √ |  | 
| generation | int | X | √ |  | 
| owner_references | json | X | √ |  | 
| subsets | json | X | √ |  | 
| kind | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| api_version | string | X | √ |  | 
| namespace | string | X | √ |  | 
| uid | string | √ | √ |  | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


