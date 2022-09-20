# ChartReleaseCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to **map[string]interface{}** |  | [optional] 
**Catalog** | Pointer to **string** | &#x60;catalog&#x60; is a valid catalog id where system will look for catalog &#x60;item&#x60; details. &#x60;train&#x60; is which train to look for under &#x60;catalog&#x60; i.e stable / testing etc. | [optional] 
**Item** | Pointer to **string** | &#x60;catalog&#x60; is a valid catalog id where system will look for catalog &#x60;item&#x60; details. &#x60;version&#x60; specifies the catalog &#x60;item&#x60; version. | [optional] 
**ReleaseName** | Pointer to **string** | &#x60;release_name&#x60; is the name which will be used to identify the created chart release. | [optional] 
**Train** | Pointer to **string** | &#x60;train&#x60; is which train to look for under &#x60;catalog&#x60; i.e stable / testing etc. | [optional] [default to "charts"]
**Version** | Pointer to **string** | &#x60;version&#x60; specifies the catalog &#x60;item&#x60; version. | [optional] [default to "latest"]

## Methods

### NewChartReleaseCreate0

`func NewChartReleaseCreate0() *ChartReleaseCreate0`

NewChartReleaseCreate0 instantiates a new ChartReleaseCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartReleaseCreate0WithDefaults

`func NewChartReleaseCreate0WithDefaults() *ChartReleaseCreate0`

NewChartReleaseCreate0WithDefaults instantiates a new ChartReleaseCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *ChartReleaseCreate0) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ChartReleaseCreate0) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ChartReleaseCreate0) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *ChartReleaseCreate0) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetCatalog

`func (o *ChartReleaseCreate0) GetCatalog() string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *ChartReleaseCreate0) GetCatalogOk() (*string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *ChartReleaseCreate0) SetCatalog(v string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *ChartReleaseCreate0) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetItem

`func (o *ChartReleaseCreate0) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ChartReleaseCreate0) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ChartReleaseCreate0) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *ChartReleaseCreate0) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetReleaseName

`func (o *ChartReleaseCreate0) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *ChartReleaseCreate0) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *ChartReleaseCreate0) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *ChartReleaseCreate0) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetTrain

`func (o *ChartReleaseCreate0) GetTrain() string`

GetTrain returns the Train field if non-nil, zero value otherwise.

### GetTrainOk

`func (o *ChartReleaseCreate0) GetTrainOk() (*string, bool)`

GetTrainOk returns a tuple with the Train field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrain

`func (o *ChartReleaseCreate0) SetTrain(v string)`

SetTrain sets Train field to given value.

### HasTrain

`func (o *ChartReleaseCreate0) HasTrain() bool`

HasTrain returns a boolean if a field has been set.

### GetVersion

`func (o *ChartReleaseCreate0) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ChartReleaseCreate0) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ChartReleaseCreate0) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ChartReleaseCreate0) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


