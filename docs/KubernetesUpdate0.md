# KubernetesUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servicelb** | Pointer to **bool** | &#x60;servicelb&#x60; is a boolean to enable or disable the integrated k3s Service Loadbalancer called \&quot;Klipper\&quot;. This can be set to disabled to enable the user to run another LoadBalancer or no LoadBalancer at all. | [optional] 
**ConfigureGpus** | Pointer to **bool** | &#x60;configure_gpus&#x60; is a boolean to enable or disable to prevent automatically loading any GPU Support into kubernetes. This includes not loading any daemonsets for Intel and NVIDIA support. | [optional] 
**Pool** | Pointer to **NullableString** | &#x60;pool&#x60; must be a valid ZFS pool configured in the system. Kubernetes service will initialise the pool by creating datasets under &#x60;pool_name/ix-applications&#x60;. | [optional] 
**ClusterCidr** | Pointer to **string** | &#x60;cluster_cidr&#x60; is the CIDR to be used for default NAT network between workloads. Specifying values for &#x60;cluster_cidr&#x60;, &#x60;service_cidr&#x60; and &#x60;cluster_dns_ip&#x60; are permanent and a subsequent change requires re-initialisation of the applications. To clarify, system will destroy old &#x60;ix-applications&#x60; dataset and any data within it when any of the values for the above configuration change. | [optional] 
**ServiceCidr** | Pointer to **string** | &#x60;service_cidr&#x60; is the CIDR to be used for kubernetes services which are an abstraction and refer to a logically set of kubernetes pods. &#x60;cluster_dns_ip&#x60; is the IP of the DNS server running for the kubernetes cluster. It must be in the range of &#x60;service_cidr&#x60;. Specifying values for &#x60;cluster_cidr&#x60;, &#x60;service_cidr&#x60; and &#x60;cluster_dns_ip&#x60; are permanent and a subsequent change requires re-initialisation of the applications. To clarify, system will destroy old &#x60;ix-applications&#x60; dataset and any data within it when any of the values for the above configuration change. | [optional] 
**ClusterDnsIp** | Pointer to **string** | &#x60;cluster_dns_ip&#x60; is the IP of the DNS server running for the kubernetes cluster. It must be in the range of &#x60;service_cidr&#x60;. Specifying values for &#x60;cluster_cidr&#x60;, &#x60;service_cidr&#x60; and &#x60;cluster_dns_ip&#x60; are permanent and a subsequent change requires re-initialisation of the applications. To clarify, system will destroy old &#x60;ix-applications&#x60; dataset and any data within it when any of the values for the above configuration change. | [optional] 
**NodeIp** | Pointer to **string** | &#x60;node_ip&#x60; is the IP address which the kubernetes cluster will assign to the TrueNAS node. It defaults to 0.0.0.0 and the cluster in this case will automatically manage which IP address to use for managing traffic for default NAT network. | [optional] 
**RouteV4Interface** | Pointer to **NullableString** | If users want to restrict traffic over a certain gateway / interface, they can specify a default route for the NAT traffic. &#x60;route_v4_interface&#x60; and &#x60;route_v4_gateway&#x60; will set a default route for the kubernetes cluster IPv4 traffic. Similarly &#x60;route_v6_interface&#x60; and &#39;route_v6_gateway&#x60; can be used to specify default route for IPv6 traffic. | [optional] 
**RouteV4Gateway** | Pointer to **NullableString** | If users want to restrict traffic over a certain gateway / interface, they can specify a default route for the NAT traffic. &#x60;route_v4_interface&#x60; and &#x60;route_v4_gateway&#x60; will set a default route for the kubernetes cluster IPv4 traffic. Similarly &#x60;route_v6_interface&#x60; and &#39;route_v6_gateway&#x60; can be used to specify default route for IPv6 traffic. | [optional] 
**RouteV6Interface** | Pointer to **NullableString** | If users want to restrict traffic over a certain gateway / interface, they can specify a default route for the NAT traffic. &#x60;route_v4_interface&#x60; and &#x60;route_v4_gateway&#x60; will set a default route for the kubernetes cluster IPv4 traffic. Similarly &#x60;route_v6_interface&#x60; and &#39;route_v6_gateway&#x60; can be used to specify default route for IPv6 traffic. | [optional] 
**RouteV6Gateway** | Pointer to **NullableString** |  | [optional] 
**MigrateApplications** | Pointer to **bool** |  | [optional] 

## Methods

### NewKubernetesUpdate0

`func NewKubernetesUpdate0() *KubernetesUpdate0`

NewKubernetesUpdate0 instantiates a new KubernetesUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesUpdate0WithDefaults

`func NewKubernetesUpdate0WithDefaults() *KubernetesUpdate0`

NewKubernetesUpdate0WithDefaults instantiates a new KubernetesUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicelb

`func (o *KubernetesUpdate0) GetServicelb() bool`

GetServicelb returns the Servicelb field if non-nil, zero value otherwise.

### GetServicelbOk

`func (o *KubernetesUpdate0) GetServicelbOk() (*bool, bool)`

GetServicelbOk returns a tuple with the Servicelb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicelb

`func (o *KubernetesUpdate0) SetServicelb(v bool)`

SetServicelb sets Servicelb field to given value.

### HasServicelb

`func (o *KubernetesUpdate0) HasServicelb() bool`

HasServicelb returns a boolean if a field has been set.

### GetConfigureGpus

`func (o *KubernetesUpdate0) GetConfigureGpus() bool`

GetConfigureGpus returns the ConfigureGpus field if non-nil, zero value otherwise.

### GetConfigureGpusOk

`func (o *KubernetesUpdate0) GetConfigureGpusOk() (*bool, bool)`

GetConfigureGpusOk returns a tuple with the ConfigureGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureGpus

`func (o *KubernetesUpdate0) SetConfigureGpus(v bool)`

SetConfigureGpus sets ConfigureGpus field to given value.

### HasConfigureGpus

`func (o *KubernetesUpdate0) HasConfigureGpus() bool`

HasConfigureGpus returns a boolean if a field has been set.

### GetPool

`func (o *KubernetesUpdate0) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *KubernetesUpdate0) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *KubernetesUpdate0) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *KubernetesUpdate0) HasPool() bool`

HasPool returns a boolean if a field has been set.

### SetPoolNil

`func (o *KubernetesUpdate0) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *KubernetesUpdate0) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetClusterCidr

`func (o *KubernetesUpdate0) GetClusterCidr() string`

GetClusterCidr returns the ClusterCidr field if non-nil, zero value otherwise.

### GetClusterCidrOk

`func (o *KubernetesUpdate0) GetClusterCidrOk() (*string, bool)`

GetClusterCidrOk returns a tuple with the ClusterCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCidr

`func (o *KubernetesUpdate0) SetClusterCidr(v string)`

SetClusterCidr sets ClusterCidr field to given value.

### HasClusterCidr

`func (o *KubernetesUpdate0) HasClusterCidr() bool`

HasClusterCidr returns a boolean if a field has been set.

### GetServiceCidr

`func (o *KubernetesUpdate0) GetServiceCidr() string`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *KubernetesUpdate0) GetServiceCidrOk() (*string, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *KubernetesUpdate0) SetServiceCidr(v string)`

SetServiceCidr sets ServiceCidr field to given value.

### HasServiceCidr

`func (o *KubernetesUpdate0) HasServiceCidr() bool`

HasServiceCidr returns a boolean if a field has been set.

### GetClusterDnsIp

`func (o *KubernetesUpdate0) GetClusterDnsIp() string`

GetClusterDnsIp returns the ClusterDnsIp field if non-nil, zero value otherwise.

### GetClusterDnsIpOk

`func (o *KubernetesUpdate0) GetClusterDnsIpOk() (*string, bool)`

GetClusterDnsIpOk returns a tuple with the ClusterDnsIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDnsIp

`func (o *KubernetesUpdate0) SetClusterDnsIp(v string)`

SetClusterDnsIp sets ClusterDnsIp field to given value.

### HasClusterDnsIp

`func (o *KubernetesUpdate0) HasClusterDnsIp() bool`

HasClusterDnsIp returns a boolean if a field has been set.

### GetNodeIp

`func (o *KubernetesUpdate0) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *KubernetesUpdate0) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *KubernetesUpdate0) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *KubernetesUpdate0) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### GetRouteV4Interface

`func (o *KubernetesUpdate0) GetRouteV4Interface() string`

GetRouteV4Interface returns the RouteV4Interface field if non-nil, zero value otherwise.

### GetRouteV4InterfaceOk

`func (o *KubernetesUpdate0) GetRouteV4InterfaceOk() (*string, bool)`

GetRouteV4InterfaceOk returns a tuple with the RouteV4Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteV4Interface

`func (o *KubernetesUpdate0) SetRouteV4Interface(v string)`

SetRouteV4Interface sets RouteV4Interface field to given value.

### HasRouteV4Interface

`func (o *KubernetesUpdate0) HasRouteV4Interface() bool`

HasRouteV4Interface returns a boolean if a field has been set.

### SetRouteV4InterfaceNil

`func (o *KubernetesUpdate0) SetRouteV4InterfaceNil(b bool)`

 SetRouteV4InterfaceNil sets the value for RouteV4Interface to be an explicit nil

### UnsetRouteV4Interface
`func (o *KubernetesUpdate0) UnsetRouteV4Interface()`

UnsetRouteV4Interface ensures that no value is present for RouteV4Interface, not even an explicit nil
### GetRouteV4Gateway

`func (o *KubernetesUpdate0) GetRouteV4Gateway() string`

GetRouteV4Gateway returns the RouteV4Gateway field if non-nil, zero value otherwise.

### GetRouteV4GatewayOk

`func (o *KubernetesUpdate0) GetRouteV4GatewayOk() (*string, bool)`

GetRouteV4GatewayOk returns a tuple with the RouteV4Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteV4Gateway

`func (o *KubernetesUpdate0) SetRouteV4Gateway(v string)`

SetRouteV4Gateway sets RouteV4Gateway field to given value.

### HasRouteV4Gateway

`func (o *KubernetesUpdate0) HasRouteV4Gateway() bool`

HasRouteV4Gateway returns a boolean if a field has been set.

### SetRouteV4GatewayNil

`func (o *KubernetesUpdate0) SetRouteV4GatewayNil(b bool)`

 SetRouteV4GatewayNil sets the value for RouteV4Gateway to be an explicit nil

### UnsetRouteV4Gateway
`func (o *KubernetesUpdate0) UnsetRouteV4Gateway()`

UnsetRouteV4Gateway ensures that no value is present for RouteV4Gateway, not even an explicit nil
### GetRouteV6Interface

`func (o *KubernetesUpdate0) GetRouteV6Interface() string`

GetRouteV6Interface returns the RouteV6Interface field if non-nil, zero value otherwise.

### GetRouteV6InterfaceOk

`func (o *KubernetesUpdate0) GetRouteV6InterfaceOk() (*string, bool)`

GetRouteV6InterfaceOk returns a tuple with the RouteV6Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteV6Interface

`func (o *KubernetesUpdate0) SetRouteV6Interface(v string)`

SetRouteV6Interface sets RouteV6Interface field to given value.

### HasRouteV6Interface

`func (o *KubernetesUpdate0) HasRouteV6Interface() bool`

HasRouteV6Interface returns a boolean if a field has been set.

### SetRouteV6InterfaceNil

`func (o *KubernetesUpdate0) SetRouteV6InterfaceNil(b bool)`

 SetRouteV6InterfaceNil sets the value for RouteV6Interface to be an explicit nil

### UnsetRouteV6Interface
`func (o *KubernetesUpdate0) UnsetRouteV6Interface()`

UnsetRouteV6Interface ensures that no value is present for RouteV6Interface, not even an explicit nil
### GetRouteV6Gateway

`func (o *KubernetesUpdate0) GetRouteV6Gateway() string`

GetRouteV6Gateway returns the RouteV6Gateway field if non-nil, zero value otherwise.

### GetRouteV6GatewayOk

`func (o *KubernetesUpdate0) GetRouteV6GatewayOk() (*string, bool)`

GetRouteV6GatewayOk returns a tuple with the RouteV6Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteV6Gateway

`func (o *KubernetesUpdate0) SetRouteV6Gateway(v string)`

SetRouteV6Gateway sets RouteV6Gateway field to given value.

### HasRouteV6Gateway

`func (o *KubernetesUpdate0) HasRouteV6Gateway() bool`

HasRouteV6Gateway returns a boolean if a field has been set.

### SetRouteV6GatewayNil

`func (o *KubernetesUpdate0) SetRouteV6GatewayNil(b bool)`

 SetRouteV6GatewayNil sets the value for RouteV6Gateway to be an explicit nil

### UnsetRouteV6Gateway
`func (o *KubernetesUpdate0) UnsetRouteV6Gateway()`

UnsetRouteV6Gateway ensures that no value is present for RouteV6Gateway, not even an explicit nil
### GetMigrateApplications

`func (o *KubernetesUpdate0) GetMigrateApplications() bool`

GetMigrateApplications returns the MigrateApplications field if non-nil, zero value otherwise.

### GetMigrateApplicationsOk

`func (o *KubernetesUpdate0) GetMigrateApplicationsOk() (*bool, bool)`

GetMigrateApplicationsOk returns a tuple with the MigrateApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateApplications

`func (o *KubernetesUpdate0) SetMigrateApplications(v bool)`

SetMigrateApplications sets MigrateApplications field to given value.

### HasMigrateApplications

`func (o *KubernetesUpdate0) HasMigrateApplications() bool`

HasMigrateApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


