# ReportingGetData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Graphs** | Pointer to [**[]Graph**](Graph.md) |  | [optional] [default to []]
**ReportingQuery** | Pointer to [**ReportingGetData1**](ReportingGetData1.md) |  | [optional] 

## Methods

### NewReportingGetData

`func NewReportingGetData() *ReportingGetData`

NewReportingGetData instantiates a new ReportingGetData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingGetDataWithDefaults

`func NewReportingGetDataWithDefaults() *ReportingGetData`

NewReportingGetDataWithDefaults instantiates a new ReportingGetData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraphs

`func (o *ReportingGetData) GetGraphs() []Graph`

GetGraphs returns the Graphs field if non-nil, zero value otherwise.

### GetGraphsOk

`func (o *ReportingGetData) GetGraphsOk() (*[]Graph, bool)`

GetGraphsOk returns a tuple with the Graphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphs

`func (o *ReportingGetData) SetGraphs(v []Graph)`

SetGraphs sets Graphs field to given value.

### HasGraphs

`func (o *ReportingGetData) HasGraphs() bool`

HasGraphs returns a boolean if a field has been set.

### GetReportingQuery

`func (o *ReportingGetData) GetReportingQuery() ReportingGetData1`

GetReportingQuery returns the ReportingQuery field if non-nil, zero value otherwise.

### GetReportingQueryOk

`func (o *ReportingGetData) GetReportingQueryOk() (*ReportingGetData1, bool)`

GetReportingQueryOk returns a tuple with the ReportingQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingQuery

`func (o *ReportingGetData) SetReportingQuery(v ReportingGetData1)`

SetReportingQuery sets ReportingQuery field to given value.

### HasReportingQuery

`func (o *ReportingGetData) HasReportingQuery() bool`

HasReportingQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


