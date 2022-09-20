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

// GlusterRebalanceStatus0 struct for GlusterRebalanceStatus0
type GlusterRebalanceStatus0 struct {
	Name *string `json:"name,omitempty"`
}

// NewGlusterRebalanceStatus0 instantiates a new GlusterRebalanceStatus0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlusterRebalanceStatus0() *GlusterRebalanceStatus0 {
	this := GlusterRebalanceStatus0{}
	return &this
}

// NewGlusterRebalanceStatus0WithDefaults instantiates a new GlusterRebalanceStatus0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlusterRebalanceStatus0WithDefaults() *GlusterRebalanceStatus0 {
	this := GlusterRebalanceStatus0{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GlusterRebalanceStatus0) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlusterRebalanceStatus0) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GlusterRebalanceStatus0) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GlusterRebalanceStatus0) SetName(v string) {
	o.Name = &v
}

func (o GlusterRebalanceStatus0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableGlusterRebalanceStatus0 struct {
	value *GlusterRebalanceStatus0
	isSet bool
}

func (v NullableGlusterRebalanceStatus0) Get() *GlusterRebalanceStatus0 {
	return v.value
}

func (v *NullableGlusterRebalanceStatus0) Set(val *GlusterRebalanceStatus0) {
	v.value = val
	v.isSet = true
}

func (v NullableGlusterRebalanceStatus0) IsSet() bool {
	return v.isSet
}

func (v *NullableGlusterRebalanceStatus0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlusterRebalanceStatus0(val *GlusterRebalanceStatus0) *NullableGlusterRebalanceStatus0 {
	return &NullableGlusterRebalanceStatus0{value: val, isSet: true}
}

func (v NullableGlusterRebalanceStatus0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlusterRebalanceStatus0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
