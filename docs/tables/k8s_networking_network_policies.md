# Table: k8s_networking_network_policies

## Primary Keys 

```
uid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| generation | int | X | √ |  | 
| spec_ingress | json | X | √ |  | 
| name | string | X | √ |  | 
| owner_references | json | X | √ |  | 
| kind | string | X | √ |  | 
| deletion_grace_period_seconds | int | X | √ |  | 
| finalizers | string_array | X | √ |  | 
| spec_pod_selector | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| resource_version | string | X | √ |  | 
| uid | string | √ | √ |  | 
| api_version | string | X | √ |  | 
| namespace | string | X | √ |  | 
| labels | json | X | √ |  | 
| annotations | json | X | √ |  | 
| spec_egress | json | X | √ |  | 
| spec_policy_types | string_array | X | √ |  | 
| context | string | X | √ |  | 
| status_conditions | json | X | √ |  | 


