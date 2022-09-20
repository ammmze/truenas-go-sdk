# IdmapAdOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaMode** | Pointer to **string** |  | [optional] [default to "SFU"]
**UnixPrimaryGroup** | Pointer to **bool** |  | [optional] [default to false]
**UnixNssInfo** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewIdmapAdOptions

`func NewIdmapAdOptions() *IdmapAdOptions`

NewIdmapAdOptions instantiates a new IdmapAdOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdmapAdOptionsWithDefaults

`func NewIdmapAdOptionsWithDefaults() *IdmapAdOptions`

NewIdmapAdOptionsWithDefaults instantiates a new IdmapAdOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaMode

`func (o *IdmapAdOptions) GetSchemaMode() string`

GetSchemaMode returns the SchemaMode field if non-nil, zero value otherwise.

### GetSchemaModeOk

`func (o *IdmapAdOptions) GetSchemaModeOk() (*string, bool)`

GetSchemaModeOk returns a tuple with the SchemaMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaMode

`func (o *IdmapAdOptions) SetSchemaMode(v string)`

SetSchemaMode sets SchemaMode field to given value.

### HasSchemaMode

`func (o *IdmapAdOptions) HasSchemaMode() bool`

HasSchemaMode returns a boolean if a field has been set.

### GetUnixPrimaryGroup

`func (o *IdmapAdOptions) GetUnixPrimaryGroup() bool`

GetUnixPrimaryGroup returns the UnixPrimaryGroup field if non-nil, zero value otherwise.

### GetUnixPrimaryGroupOk

`func (o *IdmapAdOptions) GetUnixPrimaryGroupOk() (*bool, bool)`

GetUnixPrimaryGroupOk returns a tuple with the UnixPrimaryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixPrimaryGroup

`func (o *IdmapAdOptions) SetUnixPrimaryGroup(v bool)`

SetUnixPrimaryGroup sets UnixPrimaryGroup field to given value.

### HasUnixPrimaryGroup

`func (o *IdmapAdOptions) HasUnixPrimaryGroup() bool`

HasUnixPrimaryGroup returns a boolean if a field has been set.

### GetUnixNssInfo

`func (o *IdmapAdOptions) GetUnixNssInfo() bool`

GetUnixNssInfo returns the UnixNssInfo field if non-nil, zero value otherwise.

### GetUnixNssInfoOk

`func (o *IdmapAdOptions) GetUnixNssInfoOk() (*bool, bool)`

GetUnixNssInfoOk returns a tuple with the UnixNssInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixNssInfo

`func (o *IdmapAdOptions) SetUnixNssInfo(v bool)`

SetUnixNssInfo sets UnixNssInfo field to given value.

### HasUnixNssInfo

`func (o *IdmapAdOptions) HasUnixNssInfo() bool`

HasUnixNssInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


