# ExtendedKeyUsage1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usages** | Pointer to **[]string** |  | [optional] [default to ["SERVER_AUTH"]]
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**ExtensionCritical** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewExtendedKeyUsage1

`func NewExtendedKeyUsage1() *ExtendedKeyUsage1`

NewExtendedKeyUsage1 instantiates a new ExtendedKeyUsage1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedKeyUsage1WithDefaults

`func NewExtendedKeyUsage1WithDefaults() *ExtendedKeyUsage1`

NewExtendedKeyUsage1WithDefaults instantiates a new ExtendedKeyUsage1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsages

`func (o *ExtendedKeyUsage1) GetUsages() []string`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *ExtendedKeyUsage1) GetUsagesOk() (*[]string, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *ExtendedKeyUsage1) SetUsages(v []string)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *ExtendedKeyUsage1) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetEnabled

`func (o *ExtendedKeyUsage1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExtendedKeyUsage1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExtendedKeyUsage1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ExtendedKeyUsage1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtensionCritical

`func (o *ExtendedKeyUsage1) GetExtensionCritical() bool`

GetExtensionCritical returns the ExtensionCritical field if non-nil, zero value otherwise.

### GetExtensionCriticalOk

`func (o *ExtendedKeyUsage1) GetExtensionCriticalOk() (*bool, bool)`

GetExtensionCriticalOk returns a tuple with the ExtensionCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionCritical

`func (o *ExtendedKeyUsage1) SetExtensionCritical(v bool)`

SetExtensionCritical sets ExtensionCritical field to given value.

### HasExtensionCritical

`func (o *ExtendedKeyUsage1) HasExtensionCritical() bool`

HasExtensionCritical returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


