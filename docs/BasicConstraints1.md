# BasicConstraints1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ca** | Pointer to **bool** |  | [optional] [default to true]
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**PathLength** | Pointer to **NullableInt32** |  | [optional] 
**ExtensionCritical** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewBasicConstraints1

`func NewBasicConstraints1() *BasicConstraints1`

NewBasicConstraints1 instantiates a new BasicConstraints1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicConstraints1WithDefaults

`func NewBasicConstraints1WithDefaults() *BasicConstraints1`

NewBasicConstraints1WithDefaults instantiates a new BasicConstraints1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCa

`func (o *BasicConstraints1) GetCa() bool`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *BasicConstraints1) GetCaOk() (*bool, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *BasicConstraints1) SetCa(v bool)`

SetCa sets Ca field to given value.

### HasCa

`func (o *BasicConstraints1) HasCa() bool`

HasCa returns a boolean if a field has been set.

### GetEnabled

`func (o *BasicConstraints1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BasicConstraints1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BasicConstraints1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BasicConstraints1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPathLength

`func (o *BasicConstraints1) GetPathLength() int32`

GetPathLength returns the PathLength field if non-nil, zero value otherwise.

### GetPathLengthOk

`func (o *BasicConstraints1) GetPathLengthOk() (*int32, bool)`

GetPathLengthOk returns a tuple with the PathLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathLength

`func (o *BasicConstraints1) SetPathLength(v int32)`

SetPathLength sets PathLength field to given value.

### HasPathLength

`func (o *BasicConstraints1) HasPathLength() bool`

HasPathLength returns a boolean if a field has been set.

### SetPathLengthNil

`func (o *BasicConstraints1) SetPathLengthNil(b bool)`

 SetPathLengthNil sets the value for PathLength to be an explicit nil

### UnsetPathLength
`func (o *BasicConstraints1) UnsetPathLength()`

UnsetPathLength ensures that no value is present for PathLength, not even an explicit nil
### GetExtensionCritical

`func (o *BasicConstraints1) GetExtensionCritical() bool`

GetExtensionCritical returns the ExtensionCritical field if non-nil, zero value otherwise.

### GetExtensionCriticalOk

`func (o *BasicConstraints1) GetExtensionCriticalOk() (*bool, bool)`

GetExtensionCriticalOk returns a tuple with the ExtensionCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionCritical

`func (o *BasicConstraints1) SetExtensionCritical(v bool)`

SetExtensionCritical sets ExtensionCritical field to given value.

### HasExtensionCritical

`func (o *BasicConstraints1) HasExtensionCritical() bool`

HasExtensionCritical returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


