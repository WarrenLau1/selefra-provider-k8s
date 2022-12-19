# Table: k8s_core_service_accounts

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| context | string | X | √ |  | 
| kind | string | X | √ |  | 
| api_version | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| namespace | string | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| image_pull_secrets | json | X | √ |  | 
| generation | int | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| annotations | json | X | √ |  | 
| secrets | json | X | √ |  | 
| automount_service_account_token | bool | X | √ |  | 
| uid | string | √ | √ |  | 
| resource_version | string | X | √ |  | 
| labels | json | X | √ |  | 
| owner_references | json | X | √ |  | 


