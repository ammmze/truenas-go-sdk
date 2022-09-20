# ChartReleasePullContainerImages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseName** | Pointer to **string** | Update container images being used by &#x60;release_name&#x60; chart release. | [optional] 
**PullContainerImagesOptions** | Pointer to [**ChartReleasePullContainerImages1**](ChartReleasePullContainerImages1.md) |  | [optional] 

## Methods

### NewChartReleasePullContainerImages

`func NewChartReleasePullContainerImages() *ChartReleasePullContainerImages`

NewChartReleasePullContainerImages instantiates a new ChartReleasePullContainerImages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartReleasePullContainerImagesWithDefaults

`func NewChartReleasePullContainerImagesWithDefaults() *ChartReleasePullContainerImages`

NewChartReleasePullContainerImagesWithDefaults instantiates a new ChartReleasePullContainerImages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseName

`func (o *ChartReleasePullContainerImages) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *ChartReleasePullContainerImages) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *ChartReleasePullContainerImages) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *ChartReleasePullContainerImages) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetPullContainerImagesOptions

`func (o *ChartReleasePullContainerImages) GetPullContainerImagesOptions() ChartReleasePullContainerImages1`

GetPullContainerImagesOptions returns the PullContainerImagesOptions field if non-nil, zero value otherwise.

### GetPullContainerImagesOptionsOk

`func (o *ChartReleasePullContainerImages) GetPullContainerImagesOptionsOk() (*ChartReleasePullContainerImages1, bool)`

GetPullContainerImagesOptionsOk returns a tuple with the PullContainerImagesOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullContainerImagesOptions

`func (o *ChartReleasePullContainerImages) SetPullContainerImagesOptions(v ChartReleasePullContainerImages1)`

SetPullContainerImagesOptions sets PullContainerImagesOptions field to given value.

### HasPullContainerImagesOptions

`func (o *ChartReleasePullContainerImages) HasPullContainerImagesOptions() bool`

HasPullContainerImagesOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


