# IscsiHostCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | &#x60;ip&#x60; indicates an IP address of the host. | [optional] 
**Description** | Pointer to **string** | &#x60;description&#x60; is a human-readable name for the host. | [optional] [default to ""]
**Iqns** | Pointer to **[]string** |  | [optional] [default to []]
**AddedAutomatically** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewIscsiHostCreate0

`func NewIscsiHostCreate0() *IscsiHostCreate0`

NewIscsiHostCreate0 instantiates a new IscsiHostCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIscsiHostCreate0WithDefaults

`func NewIscsiHostCreate0WithDefaults() *IscsiHostCreate0`

NewIscsiHostCreate0WithDefaults instantiates a new IscsiHostCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *IscsiHostCreate0) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IscsiHostCreate0) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IscsiHostCreate0) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *IscsiHostCreate0) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetDescription

`func (o *IscsiHostCreate0) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IscsiHostCreate0) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IscsiHostCreate0) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IscsiHostCreate0) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIqns

`func (o *IscsiHostCreate0) GetIqns() []string`

GetIqns returns the Iqns field if non-nil, zero value otherwise.

### GetIqnsOk

`func (o *IscsiHostCreate0) GetIqnsOk() (*[]string, bool)`

GetIqnsOk returns a tuple with the Iqns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqns

`func (o *IscsiHostCreate0) SetIqns(v []string)`

SetIqns sets Iqns field to given value.

### HasIqns

`func (o *IscsiHostCreate0) HasIqns() bool`

HasIqns returns a boolean if a field has been set.

### GetAddedAutomatically

`func (o *IscsiHostCreate0) GetAddedAutomatically() bool`

GetAddedAutomatically returns the AddedAutomatically field if non-nil, zero value otherwise.

### GetAddedAutomaticallyOk

`func (o *IscsiHostCreate0) GetAddedAutomaticallyOk() (*bool, bool)`

GetAddedAutomaticallyOk returns a tuple with the AddedAutomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAutomatically

`func (o *IscsiHostCreate0) SetAddedAutomatically(v bool)`

SetAddedAutomatically sets AddedAutomatically field to given value.

### HasAddedAutomatically

`func (o *IscsiHostCreate0) HasAddedAutomatically() bool`

HasAddedAutomatically returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


