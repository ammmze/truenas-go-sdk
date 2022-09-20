# VmdeviceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dtype** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Order** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewVmdeviceCreate

`func NewVmdeviceCreate() *VmdeviceCreate`

NewVmdeviceCreate instantiates a new VmdeviceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmdeviceCreateWithDefaults

`func NewVmdeviceCreateWithDefaults() *VmdeviceCreate`

NewVmdeviceCreateWithDefaults instantiates a new VmdeviceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDtype

`func (o *VmdeviceCreate) GetDtype() string`

GetDtype returns the Dtype field if non-nil, zero value otherwise.

### GetDtypeOk

`func (o *VmdeviceCreate) GetDtypeOk() (*string, bool)`

GetDtypeOk returns a tuple with the Dtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtype

`func (o *VmdeviceCreate) SetDtype(v string)`

SetDtype sets Dtype field to given value.

### HasDtype

`func (o *VmdeviceCreate) HasDtype() bool`

HasDtype returns a boolean if a field has been set.

### GetAttributes

`func (o *VmdeviceCreate) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *VmdeviceCreate) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *VmdeviceCreate) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *VmdeviceCreate) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetOrder

`func (o *VmdeviceCreate) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VmdeviceCreate) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VmdeviceCreate) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VmdeviceCreate) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *VmdeviceCreate) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *VmdeviceCreate) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


