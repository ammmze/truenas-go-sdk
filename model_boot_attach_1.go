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

// BootAttach1 struct for BootAttach1
type BootAttach1 struct {
	Expand *bool `json:"expand,omitempty"`
}

// NewBootAttach1 instantiates a new BootAttach1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootAttach1() *BootAttach1 {
	this := BootAttach1{}
	var expand bool = false
	this.Expand = &expand
	return &this
}

// NewBootAttach1WithDefaults instantiates a new BootAttach1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootAttach1WithDefaults() *BootAttach1 {
	this := BootAttach1{}
	var expand bool = false
	this.Expand = &expand
	return &this
}

// GetExpand returns the Expand field value if set, zero value otherwise.
func (o *BootAttach1) GetExpand() bool {
	if o == nil || o.Expand == nil {
		var ret bool
		return ret
	}
	return *o.Expand
}

// GetExpandOk returns a tuple with the Expand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootAttach1) GetExpandOk() (*bool, bool) {
	if o == nil || o.Expand == nil {
		return nil, false
	}
	return o.Expand, true
}

// HasExpand returns a boolean if a field has been set.
func (o *BootAttach1) HasExpand() bool {
	if o != nil && o.Expand != nil {
		return true
	}

	return false
}

// SetExpand gets a reference to the given bool and assigns it to the Expand field.
func (o *BootAttach1) SetExpand(v bool) {
	o.Expand = &v
}

func (o BootAttach1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expand != nil {
		toSerialize["expand"] = o.Expand
	}
	return json.Marshal(toSerialize)
}

type NullableBootAttach1 struct {
	value *BootAttach1
	isSet bool
}

func (v NullableBootAttach1) Get() *BootAttach1 {
	return v.value
}

func (v *NullableBootAttach1) Set(val *BootAttach1) {
	v.value = val
	v.isSet = true
}

func (v NullableBootAttach1) IsSet() bool {
	return v.isSet
}

func (v *NullableBootAttach1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootAttach1(val *BootAttach1) *NullableBootAttach1 {
	return &NullableBootAttach1{value: val, isSet: true}
}

func (v NullableBootAttach1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootAttach1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
