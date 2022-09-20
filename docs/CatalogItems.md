# CatalogItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Retrieve item details for &#x60;label&#x60; catalog. &#x60;options.cache&#x60; is a boolean which when set will try to get items details for &#x60;label&#x60; catalog from cache if available. | [optional] 
**Options** | Pointer to [**CatalogItems1**](CatalogItems1.md) |  | [optional] 

## Methods

### NewCatalogItems

`func NewCatalogItems() *CatalogItems`

NewCatalogItems instantiates a new CatalogItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemsWithDefaults

`func NewCatalogItemsWithDefaults() *CatalogItems`

NewCatalogItemsWithDefaults instantiates a new CatalogItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CatalogItems) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CatalogItems) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CatalogItems) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CatalogItems) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetOptions

`func (o *CatalogItems) GetOptions() CatalogItems1`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CatalogItems) GetOptionsOk() (*CatalogItems1, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CatalogItems) SetOptions(v CatalogItems1)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CatalogItems) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


