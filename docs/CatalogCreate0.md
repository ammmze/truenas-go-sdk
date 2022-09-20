# CatalogCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**PreferredTrains** | Pointer to **[]interface{}** |  | [optional] [default to []]
**Force** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewCatalogCreate0

`func NewCatalogCreate0() *CatalogCreate0`

NewCatalogCreate0 instantiates a new CatalogCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogCreate0WithDefaults

`func NewCatalogCreate0WithDefaults() *CatalogCreate0`

NewCatalogCreate0WithDefaults instantiates a new CatalogCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CatalogCreate0) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CatalogCreate0) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CatalogCreate0) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CatalogCreate0) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRepository

`func (o *CatalogCreate0) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CatalogCreate0) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CatalogCreate0) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *CatalogCreate0) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetBranch

`func (o *CatalogCreate0) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CatalogCreate0) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CatalogCreate0) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *CatalogCreate0) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetPreferredTrains

`func (o *CatalogCreate0) GetPreferredTrains() []interface{}`

GetPreferredTrains returns the PreferredTrains field if non-nil, zero value otherwise.

### GetPreferredTrainsOk

`func (o *CatalogCreate0) GetPreferredTrainsOk() (*[]interface{}, bool)`

GetPreferredTrainsOk returns a tuple with the PreferredTrains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTrains

`func (o *CatalogCreate0) SetPreferredTrains(v []interface{})`

SetPreferredTrains sets PreferredTrains field to given value.

### HasPreferredTrains

`func (o *CatalogCreate0) HasPreferredTrains() bool`

HasPreferredTrains returns a boolean if a field has been set.

### GetForce

`func (o *CatalogCreate0) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *CatalogCreate0) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *CatalogCreate0) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *CatalogCreate0) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


