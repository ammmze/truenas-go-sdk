# ChartReleasePodLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseName** | Pointer to **string** | Export logs of &#x60;options.container_name&#x60; container in &#x60;options.pod_name&#x60; pod in &#x60;release_name&#x60; chart release. | [optional] 
**Options** | Pointer to [**ChartReleasePodLogs1**](ChartReleasePodLogs1.md) |  | [optional] 

## Methods

### NewChartReleasePodLogs

`func NewChartReleasePodLogs() *ChartReleasePodLogs`

NewChartReleasePodLogs instantiates a new ChartReleasePodLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartReleasePodLogsWithDefaults

`func NewChartReleasePodLogsWithDefaults() *ChartReleasePodLogs`

NewChartReleasePodLogsWithDefaults instantiates a new ChartReleasePodLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseName

`func (o *ChartReleasePodLogs) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *ChartReleasePodLogs) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *ChartReleasePodLogs) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *ChartReleasePodLogs) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetOptions

`func (o *ChartReleasePodLogs) GetOptions() ChartReleasePodLogs1`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ChartReleasePodLogs) GetOptionsOk() (*ChartReleasePodLogs1, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ChartReleasePodLogs) SetOptions(v ChartReleasePodLogs1)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ChartReleasePodLogs) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


