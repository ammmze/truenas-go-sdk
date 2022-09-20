# CatalogGetItemDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemName** | Pointer to **string** |  | [optional] 
**ItemVersionDetails** | Pointer to [**CatalogGetItemDetails1**](CatalogGetItemDetails1.md) |  | [optional] 

## Methods

### NewCatalogGetItemDetails

`func NewCatalogGetItemDetails() *CatalogGetItemDetails`

NewCatalogGetItemDetails instantiates a new CatalogGetItemDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogGetItemDetailsWithDefaults

`func NewCatalogGetItemDetailsWithDefaults() *CatalogGetItemDetails`

NewCatalogGetItemDetailsWithDefaults instantiates a new CatalogGetItemDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemName

`func (o *CatalogGetItemDetails) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *CatalogGetItemDetails) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *CatalogGetItemDetails) SetItemName(v string)`

SetItemName sets ItemName field to given value.

### HasItemName

`func (o *CatalogGetItemDetails) HasItemName() bool`

HasItemName returns a boolean if a field has been set.

### GetItemVersionDetails

`func (o *CatalogGetItemDetails) GetItemVersionDetails() CatalogGetItemDetails1`

GetItemVersionDetails returns the ItemVersionDetails field if non-nil, zero value otherwise.

### GetItemVersionDetailsOk

`func (o *CatalogGetItemDetails) GetItemVersionDetailsOk() (*CatalogGetItemDetails1, bool)`

GetItemVersionDetailsOk returns a tuple with the ItemVersionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersionDetails

`func (o *CatalogGetItemDetails) SetItemVersionDetails(v CatalogGetItemDetails1)`

SetItemVersionDetails sets ItemVersionDetails field to given value.

### HasItemVersionDetails

`func (o *CatalogGetItemDetails) HasItemVersionDetails() bool`

HasItemVersionDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


