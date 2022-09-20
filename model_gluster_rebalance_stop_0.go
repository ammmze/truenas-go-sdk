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

// GlusterRebalanceStop0 struct for GlusterRebalanceStop0
type GlusterRebalanceStop0 struct {
	Name *string `json:"name,omitempty"`
}

// NewGlusterRebalanceStop0 instantiates a new GlusterRebalanceStop0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlusterRebalanceStop0() *GlusterRebalanceStop0 {
	this := GlusterRebalanceStop0{}
	return &this
}

// NewGlusterRebalanceStop0WithDefaults instantiates a new GlusterRebalanceStop0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlusterRebalanceStop0WithDefaults() *GlusterRebalanceStop0 {
	this := GlusterRebalanceStop0{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GlusterRebalanceStop0) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlusterRebalanceStop0) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GlusterRebalanceStop0) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GlusterRebalanceStop0) SetName(v string) {
	o.Name = &v
}

func (o GlusterRebalanceStop0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableGlusterRebalanceStop0 struct {
	value *GlusterRebalanceStop0
	isSet bool
}

func (v NullableGlusterRebalanceStop0) Get() *GlusterRebalanceStop0 {
	return v.value
}

func (v *NullableGlusterRebalanceStop0) Set(val *GlusterRebalanceStop0) {
	v.value = val
	v.isSet = true
}

func (v NullableGlusterRebalanceStop0) IsSet() bool {
	return v.isSet
}

func (v *NullableGlusterRebalanceStop0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlusterRebalanceStop0(val *GlusterRebalanceStop0) *NullableGlusterRebalanceStop0 {
	return &NullableGlusterRebalanceStop0{value: val, isSet: true}
}

func (v NullableGlusterRebalanceStop0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlusterRebalanceStop0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
