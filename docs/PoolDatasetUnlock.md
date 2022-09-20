# PoolDatasetUnlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unlock &#x60;id&#x60; dataset. If &#x60;id&#x60; dataset is not encrypted an exception will be raised. There is one exception: when &#x60;id&#x60; is a root dataset and &#x60;unlock_options.recursive&#x60; is specified, encryption validation will not be performed for &#x60;id&#x60;. This allow unlocking encrypted children the &#x60;id&#x60; pool. | [optional] 
**UnlockOptions** | Pointer to [**PoolDatasetUnlock1**](PoolDatasetUnlock1.md) |  | [optional] 

## Methods

### NewPoolDatasetUnlock

`func NewPoolDatasetUnlock() *PoolDatasetUnlock`

NewPoolDatasetUnlock instantiates a new PoolDatasetUnlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolDatasetUnlockWithDefaults

`func NewPoolDatasetUnlockWithDefaults() *PoolDatasetUnlock`

NewPoolDatasetUnlockWithDefaults instantiates a new PoolDatasetUnlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PoolDatasetUnlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolDatasetUnlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolDatasetUnlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PoolDatasetUnlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUnlockOptions

`func (o *PoolDatasetUnlock) GetUnlockOptions() PoolDatasetUnlock1`

GetUnlockOptions returns the UnlockOptions field if non-nil, zero value otherwise.

### GetUnlockOptionsOk

`func (o *PoolDatasetUnlock) GetUnlockOptionsOk() (*PoolDatasetUnlock1, bool)`

GetUnlockOptionsOk returns a tuple with the UnlockOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockOptions

`func (o *PoolDatasetUnlock) SetUnlockOptions(v PoolDatasetUnlock1)`

SetUnlockOptions sets UnlockOptions field to given value.

### HasUnlockOptions

`func (o *PoolDatasetUnlock) HasUnlockOptions() bool`

HasUnlockOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


