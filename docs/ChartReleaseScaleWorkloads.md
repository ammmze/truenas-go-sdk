# ChartReleaseScaleWorkloads

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseName** | Pointer to **string** |  | [optional] 
**Workloads** | Pointer to [**[]ScaleWorkload**](ScaleWorkload.md) |  | [optional] [default to []]

## Methods

### NewChartReleaseScaleWorkloads

`func NewChartReleaseScaleWorkloads() *ChartReleaseScaleWorkloads`

NewChartReleaseScaleWorkloads instantiates a new ChartReleaseScaleWorkloads object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartReleaseScaleWorkloadsWithDefaults

`func NewChartReleaseScaleWorkloadsWithDefaults() *ChartReleaseScaleWorkloads`

NewChartReleaseScaleWorkloadsWithDefaults instantiates a new ChartReleaseScaleWorkloads object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseName

`func (o *ChartReleaseScaleWorkloads) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *ChartReleaseScaleWorkloads) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *ChartReleaseScaleWorkloads) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *ChartReleaseScaleWorkloads) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetWorkloads

`func (o *ChartReleaseScaleWorkloads) GetWorkloads() []ScaleWorkload`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *ChartReleaseScaleWorkloads) GetWorkloadsOk() (*[]ScaleWorkload, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *ChartReleaseScaleWorkloads) SetWorkloads(v []ScaleWorkload)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *ChartReleaseScaleWorkloads) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


