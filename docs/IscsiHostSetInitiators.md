# IscsiHostSetInitiators

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Associates initiator groups &#x60;ids&#x60; with host &#x60;id&#x60;. | [optional] 
**Ids** | Pointer to **[]int32** | Associates initiator groups &#x60;ids&#x60; with host &#x60;id&#x60;. | [optional] [default to []]
**Force** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewIscsiHostSetInitiators

`func NewIscsiHostSetInitiators() *IscsiHostSetInitiators`

NewIscsiHostSetInitiators instantiates a new IscsiHostSetInitiators object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIscsiHostSetInitiatorsWithDefaults

`func NewIscsiHostSetInitiatorsWithDefaults() *IscsiHostSetInitiators`

NewIscsiHostSetInitiatorsWithDefaults instantiates a new IscsiHostSetInitiators object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IscsiHostSetInitiators) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IscsiHostSetInitiators) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IscsiHostSetInitiators) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IscsiHostSetInitiators) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIds

`func (o *IscsiHostSetInitiators) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *IscsiHostSetInitiators) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *IscsiHostSetInitiators) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *IscsiHostSetInitiators) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetForce

`func (o *IscsiHostSetInitiators) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *IscsiHostSetInitiators) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *IscsiHostSetInitiators) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *IscsiHostSetInitiators) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


