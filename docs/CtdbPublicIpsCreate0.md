# CtdbPublicIpsCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pnn** | Pointer to **int32** | &#x60;pnn&#x60; node number of record to adjust | [optional] 
**Ip** | Pointer to **string** | &#x60;ip&#x60; string representing an IP v4/v6 address | [optional] 
**Netmask** | Pointer to **int32** | &#x60;netmask&#x60; integer representing a cidr notated netmask (i.e. 16/24/48/64 etc) | [optional] 
**Interface** | Pointer to **string** |  | [optional] 

## Methods

### NewCtdbPublicIpsCreate0

`func NewCtdbPublicIpsCreate0() *CtdbPublicIpsCreate0`

NewCtdbPublicIpsCreate0 instantiates a new CtdbPublicIpsCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCtdbPublicIpsCreate0WithDefaults

`func NewCtdbPublicIpsCreate0WithDefaults() *CtdbPublicIpsCreate0`

NewCtdbPublicIpsCreate0WithDefaults instantiates a new CtdbPublicIpsCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPnn

`func (o *CtdbPublicIpsCreate0) GetPnn() int32`

GetPnn returns the Pnn field if non-nil, zero value otherwise.

### GetPnnOk

`func (o *CtdbPublicIpsCreate0) GetPnnOk() (*int32, bool)`

GetPnnOk returns a tuple with the Pnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnn

`func (o *CtdbPublicIpsCreate0) SetPnn(v int32)`

SetPnn sets Pnn field to given value.

### HasPnn

`func (o *CtdbPublicIpsCreate0) HasPnn() bool`

HasPnn returns a boolean if a field has been set.

### GetIp

`func (o *CtdbPublicIpsCreate0) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CtdbPublicIpsCreate0) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CtdbPublicIpsCreate0) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CtdbPublicIpsCreate0) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetmask

`func (o *CtdbPublicIpsCreate0) GetNetmask() int32`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *CtdbPublicIpsCreate0) GetNetmaskOk() (*int32, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *CtdbPublicIpsCreate0) SetNetmask(v int32)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *CtdbPublicIpsCreate0) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetInterface

`func (o *CtdbPublicIpsCreate0) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *CtdbPublicIpsCreate0) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *CtdbPublicIpsCreate0) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *CtdbPublicIpsCreate0) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


