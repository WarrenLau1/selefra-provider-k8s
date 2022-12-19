# Table: k8s_networking_network_policies

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| api_version | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| context | string | X | √ |  | 
| kind | string | X | √ |  | 
| namespace | string | X | √ |  | 
| resource_version | string | X | √ |  | 
| generation | int | X | √ |  | 
| spec_policy_types | string_array | X | √ |  | 
| status_conditions | json | X | √ |  | 
| uid | string | √ | √ |  | 
| owner_references | json | X | √ |  | 
| spec_egress | json | X | √ |  | 
| annotations | json | X | √ |  | 
| labels | json | X | √ |  | 
| spec_pod_selector | json | X | √ |  | 
| spec_ingress | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 


