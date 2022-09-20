# ChartReleaseScale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseName** | Pointer to **string** | Scale a &#x60;release_name&#x60; chart release to &#x60;scale_options.replica_count&#x60; specified. | [optional] 
**ScaleOptions** | Pointer to [**ChartReleaseScale1**](ChartReleaseScale1.md) |  | [optional] 

## Methods

### NewChartReleaseScale

`func NewChartReleaseScale() *ChartReleaseScale`

NewChartReleaseScale instantiates a new ChartReleaseScale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartReleaseScaleWithDefaults

`func NewChartReleaseScaleWithDefaults() *ChartReleaseScale`

NewChartReleaseScaleWithDefaults instantiates a new ChartReleaseScale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseName

`func (o *ChartReleaseScale) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *ChartReleaseScale) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *ChartReleaseScale) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *ChartReleaseScale) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetScaleOptions

`func (o *ChartReleaseScale) GetScaleOptions() ChartReleaseScale1`

GetScaleOptions returns the ScaleOptions field if non-nil, zero value otherwise.

### GetScaleOptionsOk

`func (o *ChartReleaseScale) GetScaleOptionsOk() (*ChartReleaseScale1, bool)`

GetScaleOptionsOk returns a tuple with the ScaleOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleOptions

`func (o *ChartReleaseScale) SetScaleOptions(v ChartReleaseScale1)`

SetScaleOptions sets ScaleOptions field to given value.

### HasScaleOptions

`func (o *ChartReleaseScale) HasScaleOptions() bool`

HasScaleOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


