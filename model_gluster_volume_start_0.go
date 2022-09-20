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

// GlusterVolumeStart0 struct for GlusterVolumeStart0
type GlusterVolumeStart0 struct {
	// `name` String representing name of gluster volume
	Name  *string `json:"name,omitempty"`
	Force *bool   `json:"force,omitempty"`
}

// NewGlusterVolumeStart0 instantiates a new GlusterVolumeStart0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlusterVolumeStart0() *GlusterVolumeStart0 {
	this := GlusterVolumeStart0{}
	var force bool = true
	this.Force = &force
	return &this
}

// NewGlusterVolumeStart0WithDefaults instantiates a new GlusterVolumeStart0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlusterVolumeStart0WithDefaults() *GlusterVolumeStart0 {
	this := GlusterVolumeStart0{}
	var force bool = true
	this.Force = &force
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GlusterVolumeStart0) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlusterVolumeStart0) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GlusterVolumeStart0) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GlusterVolumeStart0) SetName(v string) {
	o.Name = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *GlusterVolumeStart0) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlusterVolumeStart0) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *GlusterVolumeStart0) HasForce() bool {
	if o != nil && o.Force != nil {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *GlusterVolumeStart0) SetForce(v bool) {
	o.Force = &v
}

func (o GlusterVolumeStart0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Force != nil {
		toSerialize["force"] = o.Force
	}
	return json.Marshal(toSerialize)
}

type NullableGlusterVolumeStart0 struct {
	value *GlusterVolumeStart0
	isSet bool
}

func (v NullableGlusterVolumeStart0) Get() *GlusterVolumeStart0 {
	return v.value
}

func (v *NullableGlusterVolumeStart0) Set(val *GlusterVolumeStart0) {
	v.value = val
	v.isSet = true
}

func (v NullableGlusterVolumeStart0) IsSet() bool {
	return v.isSet
}

func (v *NullableGlusterVolumeStart0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlusterVolumeStart0(val *GlusterVolumeStart0) *NullableGlusterVolumeStart0 {
	return &NullableGlusterVolumeStart0{value: val, isSet: true}
}

func (v NullableGlusterVolumeStart0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlusterVolumeStart0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
