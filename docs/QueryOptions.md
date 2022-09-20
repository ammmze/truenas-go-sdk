# QueryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Relationships** | Pointer to **bool** |  | [optional] [default to true]
**Extend** | Pointer to **NullableString** |  | [optional] 
**ExtendContext** | Pointer to **NullableString** |  | [optional] 
**Prefix** | Pointer to **NullableString** |  | [optional] 
**Extra** | Pointer to **map[string]interface{}** |  | [optional] 
**OrderBy** | Pointer to **[]interface{}** |  | [optional] [default to []]
**Select** | Pointer to **[]interface{}** |  | [optional] [default to []]
**Count** | Pointer to **bool** |  | [optional] [default to false]
**Get** | Pointer to **bool** |  | [optional] [default to false]
**Offset** | Pointer to **int32** |  | [optional] [default to 0]
**Limit** | Pointer to **int32** |  | [optional] [default to 0]
**ForceSqlFilters** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewQueryOptions

`func NewQueryOptions() *QueryOptions`

NewQueryOptions instantiates a new QueryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryOptionsWithDefaults

`func NewQueryOptionsWithDefaults() *QueryOptions`

NewQueryOptionsWithDefaults instantiates a new QueryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelationships

`func (o *QueryOptions) GetRelationships() bool`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *QueryOptions) GetRelationshipsOk() (*bool, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *QueryOptions) SetRelationships(v bool)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *QueryOptions) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetExtend

`func (o *QueryOptions) GetExtend() string`

GetExtend returns the Extend field if non-nil, zero value otherwise.

### GetExtendOk

`func (o *QueryOptions) GetExtendOk() (*string, bool)`

GetExtendOk returns a tuple with the Extend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtend

`func (o *QueryOptions) SetExtend(v string)`

SetExtend sets Extend field to given value.

### HasExtend

`func (o *QueryOptions) HasExtend() bool`

HasExtend returns a boolean if a field has been set.

### SetExtendNil

`func (o *QueryOptions) SetExtendNil(b bool)`

 SetExtendNil sets the value for Extend to be an explicit nil

### UnsetExtend
`func (o *QueryOptions) UnsetExtend()`

UnsetExtend ensures that no value is present for Extend, not even an explicit nil
### GetExtendContext

`func (o *QueryOptions) GetExtendContext() string`

GetExtendContext returns the ExtendContext field if non-nil, zero value otherwise.

### GetExtendContextOk

`func (o *QueryOptions) GetExtendContextOk() (*string, bool)`

GetExtendContextOk returns a tuple with the ExtendContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendContext

`func (o *QueryOptions) SetExtendContext(v string)`

SetExtendContext sets ExtendContext field to given value.

### HasExtendContext

`func (o *QueryOptions) HasExtendContext() bool`

HasExtendContext returns a boolean if a field has been set.

### SetExtendContextNil

`func (o *QueryOptions) SetExtendContextNil(b bool)`

 SetExtendContextNil sets the value for ExtendContext to be an explicit nil

### UnsetExtendContext
`func (o *QueryOptions) UnsetExtendContext()`

UnsetExtendContext ensures that no value is present for ExtendContext, not even an explicit nil
### GetPrefix

`func (o *QueryOptions) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *QueryOptions) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *QueryOptions) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *QueryOptions) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *QueryOptions) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *QueryOptions) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetExtra

`func (o *QueryOptions) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *QueryOptions) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *QueryOptions) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *QueryOptions) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetOrderBy

`func (o *QueryOptions) GetOrderBy() []interface{}`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *QueryOptions) GetOrderByOk() (*[]interface{}, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *QueryOptions) SetOrderBy(v []interface{})`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *QueryOptions) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetSelect

`func (o *QueryOptions) GetSelect() []interface{}`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *QueryOptions) GetSelectOk() (*[]interface{}, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *QueryOptions) SetSelect(v []interface{})`

SetSelect sets Select field to given value.

### HasSelect

`func (o *QueryOptions) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### GetCount

`func (o *QueryOptions) GetCount() bool`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *QueryOptions) GetCountOk() (*bool, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *QueryOptions) SetCount(v bool)`

SetCount sets Count field to given value.

### HasCount

`func (o *QueryOptions) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetGet

`func (o *QueryOptions) GetGet() bool`

GetGet returns the Get field if non-nil, zero value otherwise.

### GetGetOk

`func (o *QueryOptions) GetGetOk() (*bool, bool)`

GetGetOk returns a tuple with the Get field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGet

`func (o *QueryOptions) SetGet(v bool)`

SetGet sets Get field to given value.

### HasGet

`func (o *QueryOptions) HasGet() bool`

HasGet returns a boolean if a field has been set.

### GetOffset

`func (o *QueryOptions) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QueryOptions) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QueryOptions) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QueryOptions) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *QueryOptions) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QueryOptions) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QueryOptions) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QueryOptions) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetForceSqlFilters

`func (o *QueryOptions) GetForceSqlFilters() bool`

GetForceSqlFilters returns the ForceSqlFilters field if non-nil, zero value otherwise.

### GetForceSqlFiltersOk

`func (o *QueryOptions) GetForceSqlFiltersOk() (*bool, bool)`

GetForceSqlFiltersOk returns a tuple with the ForceSqlFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSqlFilters

`func (o *QueryOptions) SetForceSqlFilters(v bool)`

SetForceSqlFilters sets ForceSqlFilters field to given value.

### HasForceSqlFilters

`func (o *QueryOptions) HasForceSqlFilters() bool`

HasForceSqlFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


