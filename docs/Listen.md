# Listen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] [default to 3260]

## Methods

### NewListen

`func NewListen() *Listen`

NewListen instantiates a new Listen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListenWithDefaults

`func NewListenWithDefaults() *Listen`

NewListenWithDefaults instantiates a new Listen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *Listen) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Listen) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Listen) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Listen) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPort

`func (o *Listen) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Listen) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Listen) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Listen) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


