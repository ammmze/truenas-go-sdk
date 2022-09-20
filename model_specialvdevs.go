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

// Specialvdevs struct for Specialvdevs
type Specialvdevs struct {
	Type  *string  `json:"type,omitempty"`
	Disks []string `json:"disks,omitempty"`
}

// NewSpecialvdevs instantiates a new Specialvdevs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecialvdevs() *Specialvdevs {
	this := Specialvdevs{}
	return &this
}

// NewSpecialvdevsWithDefaults instantiates a new Specialvdevs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecialvdevsWithDefaults() *Specialvdevs {
	this := Specialvdevs{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Specialvdevs) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Specialvdevs) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Specialvdevs) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Specialvdevs) SetType(v string) {
	o.Type = &v
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *Specialvdevs) GetDisks() []string {
	if o == nil || o.Disks == nil {
		var ret []string
		return ret
	}
	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Specialvdevs) GetDisksOk() ([]string, bool) {
	if o == nil || o.Disks == nil {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *Specialvdevs) HasDisks() bool {
	if o != nil && o.Disks != nil {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []string and assigns it to the Disks field.
func (o *Specialvdevs) SetDisks(v []string) {
	o.Disks = v
}

func (o Specialvdevs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Disks != nil {
		toSerialize["disks"] = o.Disks
	}
	return json.Marshal(toSerialize)
}

type NullableSpecialvdevs struct {
	value *Specialvdevs
	isSet bool
}

func (v NullableSpecialvdevs) Get() *Specialvdevs {
	return v.value
}

func (v *NullableSpecialvdevs) Set(val *Specialvdevs) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecialvdevs) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecialvdevs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecialvdevs(val *Specialvdevs) *NullableSpecialvdevs {
	return &NullableSpecialvdevs{value: val, isSet: true}
}

func (v NullableSpecialvdevs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecialvdevs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
