# StatsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Dataset** | Pointer to **string** |  | [optional] 
**Cf** | Pointer to **string** |  | [optional] [default to "AVERAGE"]

## Methods

### NewStatsData

`func NewStatsData() *StatsData`

NewStatsData instantiates a new StatsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsDataWithDefaults

`func NewStatsDataWithDefaults() *StatsData`

NewStatsDataWithDefaults instantiates a new StatsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *StatsData) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StatsData) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StatsData) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StatsData) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *StatsData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StatsData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StatsData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StatsData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDataset

`func (o *StatsData) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *StatsData) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *StatsData) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *StatsData) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetCf

`func (o *StatsData) GetCf() string`

GetCf returns the Cf field if non-nil, zero value otherwise.

### GetCfOk

`func (o *StatsData) GetCfOk() (*string, bool)`

GetCfOk returns a tuple with the Cf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCf

`func (o *StatsData) SetCf(v string)`

SetCf sets Cf field to given value.

### HasCf

`func (o *StatsData) HasCf() bool`

HasCf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


