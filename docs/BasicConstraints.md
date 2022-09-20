# BasicConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ca** | Pointer to **bool** |  | [optional] [default to false]
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**PathLength** | Pointer to **NullableInt32** |  | [optional] 
**ExtensionCritical** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewBasicConstraints

`func NewBasicConstraints() *BasicConstraints`

NewBasicConstraints instantiates a new BasicConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicConstraintsWithDefaults

`func NewBasicConstraintsWithDefaults() *BasicConstraints`

NewBasicConstraintsWithDefaults instantiates a new BasicConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCa

`func (o *BasicConstraints) GetCa() bool`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *BasicConstraints) GetCaOk() (*bool, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *BasicConstraints) SetCa(v bool)`

SetCa sets Ca field to given value.

### HasCa

`func (o *BasicConstraints) HasCa() bool`

HasCa returns a boolean if a field has been set.

### GetEnabled

`func (o *BasicConstraints) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BasicConstraints) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BasicConstraints) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BasicConstraints) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPathLength

`func (o *BasicConstraints) GetPathLength() int32`

GetPathLength returns the PathLength field if non-nil, zero value otherwise.

### GetPathLengthOk

`func (o *BasicConstraints) GetPathLengthOk() (*int32, bool)`

GetPathLengthOk returns a tuple with the PathLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathLength

`func (o *BasicConstraints) SetPathLength(v int32)`

SetPathLength sets PathLength field to given value.

### HasPathLength

`func (o *BasicConstraints) HasPathLength() bool`

HasPathLength returns a boolean if a field has been set.

### SetPathLengthNil

`func (o *BasicConstraints) SetPathLengthNil(b bool)`

 SetPathLengthNil sets the value for PathLength to be an explicit nil

### UnsetPathLength
`func (o *BasicConstraints) UnsetPathLength()`

UnsetPathLength ensures that no value is present for PathLength, not even an explicit nil
### GetExtensionCritical

`func (o *BasicConstraints) GetExtensionCritical() bool`

GetExtensionCritical returns the ExtensionCritical field if non-nil, zero value otherwise.

### GetExtensionCriticalOk

`func (o *BasicConstraints) GetExtensionCriticalOk() (*bool, bool)`

GetExtensionCriticalOk returns a tuple with the ExtensionCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionCritical

`func (o *BasicConstraints) SetExtensionCritical(v bool)`

SetExtensionCritical sets ExtensionCritical field to given value.

### HasExtensionCritical

`func (o *BasicConstraints) HasExtensionCritical() bool`

HasExtensionCritical returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


