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

// WebuiImageCreate0 struct for WebuiImageCreate0
type WebuiImageCreate0 struct {
	Identifier *string `json:"identifier,omitempty"`
}

// NewWebuiImageCreate0 instantiates a new WebuiImageCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebuiImageCreate0() *WebuiImageCreate0 {
	this := WebuiImageCreate0{}
	return &this
}

// NewWebuiImageCreate0WithDefaults instantiates a new WebuiImageCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebuiImageCreate0WithDefaults() *WebuiImageCreate0 {
	this := WebuiImageCreate0{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *WebuiImageCreate0) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebuiImageCreate0) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *WebuiImageCreate0) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *WebuiImageCreate0) SetIdentifier(v string) {
	o.Identifier = &v
}

func (o WebuiImageCreate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	return json.Marshal(toSerialize)
}

type NullableWebuiImageCreate0 struct {
	value *WebuiImageCreate0
	isSet bool
}

func (v NullableWebuiImageCreate0) Get() *WebuiImageCreate0 {
	return v.value
}

func (v *NullableWebuiImageCreate0) Set(val *WebuiImageCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableWebuiImageCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableWebuiImageCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebuiImageCreate0(val *WebuiImageCreate0) *NullableWebuiImageCreate0 {
	return &NullableWebuiImageCreate0{value: val, isSet: true}
}

func (v NullableWebuiImageCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebuiImageCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
