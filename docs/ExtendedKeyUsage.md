# ExtendedKeyUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usages** | Pointer to **[]string** |  | [optional] [default to []]
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**ExtensionCritical** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewExtendedKeyUsage

`func NewExtendedKeyUsage() *ExtendedKeyUsage`

NewExtendedKeyUsage instantiates a new ExtendedKeyUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedKeyUsageWithDefaults

`func NewExtendedKeyUsageWithDefaults() *ExtendedKeyUsage`

NewExtendedKeyUsageWithDefaults instantiates a new ExtendedKeyUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsages

`func (o *ExtendedKeyUsage) GetUsages() []string`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *ExtendedKeyUsage) GetUsagesOk() (*[]string, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *ExtendedKeyUsage) SetUsages(v []string)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *ExtendedKeyUsage) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetEnabled

`func (o *ExtendedKeyUsage) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExtendedKeyUsage) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExtendedKeyUsage) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ExtendedKeyUsage) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtensionCritical

`func (o *ExtendedKeyUsage) GetExtensionCritical() bool`

GetExtensionCritical returns the ExtensionCritical field if non-nil, zero value otherwise.

### GetExtensionCriticalOk

`func (o *ExtendedKeyUsage) GetExtensionCriticalOk() (*bool, bool)`

GetExtensionCriticalOk returns a tuple with the ExtensionCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionCritical

`func (o *ExtendedKeyUsage) SetExtensionCritical(v bool)`

SetExtensionCritical sets ExtensionCritical field to given value.

### HasExtensionCritical

`func (o *ExtendedKeyUsage) HasExtensionCritical() bool`

HasExtensionCritical returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


