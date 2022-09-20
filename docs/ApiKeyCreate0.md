# ApiKeyCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Allowlist** | Pointer to [**[]AllowlistItem**](AllowlistItem.md) |  | [optional] [default to []]

## Methods

### NewApiKeyCreate0

`func NewApiKeyCreate0() *ApiKeyCreate0`

NewApiKeyCreate0 instantiates a new ApiKeyCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyCreate0WithDefaults

`func NewApiKeyCreate0WithDefaults() *ApiKeyCreate0`

NewApiKeyCreate0WithDefaults instantiates a new ApiKeyCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiKeyCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeyCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiKeyCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAllowlist

`func (o *ApiKeyCreate0) GetAllowlist() []AllowlistItem`

GetAllowlist returns the Allowlist field if non-nil, zero value otherwise.

### GetAllowlistOk

`func (o *ApiKeyCreate0) GetAllowlistOk() (*[]AllowlistItem, bool)`

GetAllowlistOk returns a tuple with the Allowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlist

`func (o *ApiKeyCreate0) SetAllowlist(v []AllowlistItem)`

SetAllowlist sets Allowlist field to given value.

### HasAllowlist

`func (o *ApiKeyCreate0) HasAllowlist() bool`

HasAllowlist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


