# FailoverUnlock0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pools** | Pointer to [**[]PoolKeys**](PoolKeys.md) |  | [optional] [default to []]
**Datasets** | Pointer to [**[]DatasetKeys**](DatasetKeys.md) |  | [optional] [default to []]

## Methods

### NewFailoverUnlock0

`func NewFailoverUnlock0() *FailoverUnlock0`

NewFailoverUnlock0 instantiates a new FailoverUnlock0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailoverUnlock0WithDefaults

`func NewFailoverUnlock0WithDefaults() *FailoverUnlock0`

NewFailoverUnlock0WithDefaults instantiates a new FailoverUnlock0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPools

`func (o *FailoverUnlock0) GetPools() []PoolKeys`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *FailoverUnlock0) GetPoolsOk() (*[]PoolKeys, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *FailoverUnlock0) SetPools(v []PoolKeys)`

SetPools sets Pools field to given value.

### HasPools

`func (o *FailoverUnlock0) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetDatasets

`func (o *FailoverUnlock0) GetDatasets() []DatasetKeys`

GetDatasets returns the Datasets field if non-nil, zero value otherwise.

### GetDatasetsOk

`func (o *FailoverUnlock0) GetDatasetsOk() (*[]DatasetKeys, bool)`

GetDatasetsOk returns a tuple with the Datasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasets

`func (o *FailoverUnlock0) SetDatasets(v []DatasetKeys)`

SetDatasets sets Datasets field to given value.

### HasDatasets

`func (o *FailoverUnlock0) HasDatasets() bool`

HasDatasets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


