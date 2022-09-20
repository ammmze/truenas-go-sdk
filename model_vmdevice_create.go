/*
TrueNAS RESTful API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
)

// VmdeviceCreate struct for VmdeviceCreate
type VmdeviceCreate struct {
	Dtype      *string                `json:"dtype,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Order      NullableInt32          `json:"order,omitempty"`
}

// NewVmdeviceCreate instantiates a new VmdeviceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmdeviceCreate() *VmdeviceCreate {
	this := VmdeviceCreate{}
	return &this
}

// NewVmdeviceCreateWithDefaults instantiates a new VmdeviceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmdeviceCreateWithDefaults() *VmdeviceCreate {
	this := VmdeviceCreate{}
	return &this
}

// GetDtype returns the Dtype field value if set, zero value otherwise.
func (o *VmdeviceCreate) GetDtype() string {
	if o == nil || o.Dtype == nil {
		var ret string
		return ret
	}
	return *o.Dtype
}

// GetDtypeOk returns a tuple with the Dtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmdeviceCreate) GetDtypeOk() (*string, bool) {
	if o == nil || o.Dtype == nil {
		return nil, false
	}
	return o.Dtype, true
}

// HasDtype returns a boolean if a field has been set.
func (o *VmdeviceCreate) HasDtype() bool {
	if o != nil && o.Dtype != nil {
		return true
	}

	return false
}

// SetDtype gets a reference to the given string and assigns it to the Dtype field.
func (o *VmdeviceCreate) SetDtype(v string) {
	o.Dtype = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *VmdeviceCreate) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmdeviceCreate) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *VmdeviceCreate) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *VmdeviceCreate) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetOrder returns the Order field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmdeviceCreate) GetOrder() int32 {
	if o == nil || o.Order.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Order.Get()
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmdeviceCreate) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Order.Get(), o.Order.IsSet()
}

// HasOrder returns a boolean if a field has been set.
func (o *VmdeviceCreate) HasOrder() bool {
	if o != nil && o.Order.IsSet() {
		return true
	}

	return false
}

// SetOrder gets a reference to the given NullableInt32 and assigns it to the Order field.
func (o *VmdeviceCreate) SetOrder(v int32) {
	o.Order.Set(&v)
}

// SetOrderNil sets the value for Order to be an explicit nil
func (o *VmdeviceCreate) SetOrderNil() {
	o.Order.Set(nil)
}

// UnsetOrder ensures that no value is present for Order, not even an explicit nil
func (o *VmdeviceCreate) UnsetOrder() {
	o.Order.Unset()
}

func (o VmdeviceCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dtype != nil {
		toSerialize["dtype"] = o.Dtype
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Order.IsSet() {
		toSerialize["order"] = o.Order.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVmdeviceCreate struct {
	value *VmdeviceCreate
	isSet bool
}

func (v NullableVmdeviceCreate) Get() *VmdeviceCreate {
	return v.value
}

func (v *NullableVmdeviceCreate) Set(val *VmdeviceCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableVmdeviceCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableVmdeviceCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmdeviceCreate(val *VmdeviceCreate) *NullableVmdeviceCreate {
	return &NullableVmdeviceCreate{value: val, isSet: true}
}

func (v NullableVmdeviceCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmdeviceCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
