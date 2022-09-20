# ChartReleaseUpgradeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseName** | Pointer to **string** | Retrieve upgrade summary for &#x60;release_name&#x60; which will include which container images will be updated and changelog for &#x60;options.item_version&#x60; chart version specified if applicable. If only container images need to be updated, changelog will be &#x60;null&#x60;. | [optional] 
**Options** | Pointer to [**ChartReleaseUpgradeSummary1**](ChartReleaseUpgradeSummary1.md) |  | [optional] 

## Methods

### NewChartReleaseUpgradeSummary

`func NewChartReleaseUpgradeSummary() *ChartReleaseUpgradeSummary`

NewChartReleaseUpgradeSummary instantiates a new ChartReleaseUpgradeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartReleaseUpgradeSummaryWithDefaults

`func NewChartReleaseUpgradeSummaryWithDefaults() *ChartReleaseUpgradeSummary`

NewChartReleaseUpgradeSummaryWithDefaults instantiates a new ChartReleaseUpgradeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseName

`func (o *ChartReleaseUpgradeSummary) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *ChartReleaseUpgradeSummary) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *ChartReleaseUpgradeSummary) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *ChartReleaseUpgradeSummary) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetOptions

`func (o *ChartReleaseUpgradeSummary) GetOptions() ChartReleaseUpgradeSummary1`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ChartReleaseUpgradeSummary) GetOptionsOk() (*ChartReleaseUpgradeSummary1, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ChartReleaseUpgradeSummary) SetOptions(v ChartReleaseUpgradeSummary1)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ChartReleaseUpgradeSummary) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


