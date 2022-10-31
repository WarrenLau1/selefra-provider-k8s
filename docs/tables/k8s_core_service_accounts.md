# Table: k8s_core_service_accounts

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| context | string | X | √ |  | 
| api_version | string | X | √ |  | 
| generation | int | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| labels | json | X | √ |  | 
| image_pull_secrets | json | X | √ |  | 
| name | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| secrets | json | X | √ |  | 
| automount_service_account_token | bool | X | √ |  | 
| uid | string | √ | √ |  | 
| kind | string | X | √ |  | 
| namespace | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| annotations | json | X | √ |  | 


