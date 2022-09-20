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

// GlusterPeerCreate0 struct for GlusterPeerCreate0
type GlusterPeerCreate0 struct {
	Hostname *string `json:"hostname,omitempty"`
}

// NewGlusterPeerCreate0 instantiates a new GlusterPeerCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlusterPeerCreate0() *GlusterPeerCreate0 {
	this := GlusterPeerCreate0{}
	return &this
}

// NewGlusterPeerCreate0WithDefaults instantiates a new GlusterPeerCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlusterPeerCreate0WithDefaults() *GlusterPeerCreate0 {
	this := GlusterPeerCreate0{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *GlusterPeerCreate0) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlusterPeerCreate0) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *GlusterPeerCreate0) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *GlusterPeerCreate0) SetHostname(v string) {
	o.Hostname = &v
}

func (o GlusterPeerCreate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	return json.Marshal(toSerialize)
}

type NullableGlusterPeerCreate0 struct {
	value *GlusterPeerCreate0
	isSet bool
}

func (v NullableGlusterPeerCreate0) Get() *GlusterPeerCreate0 {
	return v.value
}

func (v *NullableGlusterPeerCreate0) Set(val *GlusterPeerCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableGlusterPeerCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableGlusterPeerCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlusterPeerCreate0(val *GlusterPeerCreate0) *NullableGlusterPeerCreate0 {
	return &NullableGlusterPeerCreate0{value: val, isSet: true}
}

func (v NullableGlusterPeerCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlusterPeerCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
