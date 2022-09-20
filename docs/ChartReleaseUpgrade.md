# ChartReleaseUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseName** | Pointer to **string** | Upgrade &#x60;release_name&#x60; chart release. System will update container images being used by &#x60;release_name&#x60; chart release as a chart release upgrade is not considered complete until the images in use have also been updated to latest versions. | [optional] 
**UpgradeOptions** | Pointer to [**ChartReleaseUpgrade1**](ChartReleaseUpgrade1.md) |  | [optional] 

## Methods

### NewChartReleaseUpgrade

`func NewChartReleaseUpgrade() *ChartReleaseUpgrade`

NewChartReleaseUpgrade instantiates a new ChartReleaseUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartReleaseUpgradeWithDefaults

`func NewChartReleaseUpgradeWithDefaults() *ChartReleaseUpgrade`

NewChartReleaseUpgradeWithDefaults instantiates a new ChartReleaseUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseName

`func (o *ChartReleaseUpgrade) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *ChartReleaseUpgrade) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *ChartReleaseUpgrade) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *ChartReleaseUpgrade) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetUpgradeOptions

`func (o *ChartReleaseUpgrade) GetUpgradeOptions() ChartReleaseUpgrade1`

GetUpgradeOptions returns the UpgradeOptions field if non-nil, zero value otherwise.

### GetUpgradeOptionsOk

`func (o *ChartReleaseUpgrade) GetUpgradeOptionsOk() (*ChartReleaseUpgrade1, bool)`

GetUpgradeOptionsOk returns a tuple with the UpgradeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOptions

`func (o *ChartReleaseUpgrade) SetUpgradeOptions(v ChartReleaseUpgrade1)`

SetUpgradeOptions sets UpgradeOptions field to given value.

### HasUpgradeOptions

`func (o *ChartReleaseUpgrade) HasUpgradeOptions() bool`

HasUpgradeOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


