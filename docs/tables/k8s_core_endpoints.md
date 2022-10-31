# Table: k8s_core_endpoints

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| context | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| subsets | json | X | √ |  | 
| annotations | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| kind | string | X | √ |  | 
| generation | int | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| api_version | string | X | √ |  | 
| namespace | string | X | √ |  | 
| labels | json | X | √ |  | 
| uid | string | √ | √ |  | 
| name | string | X | √ |  | 
| owner_references | json | X | √ |  | 


