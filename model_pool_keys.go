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

// PoolKeys struct for PoolKeys
type PoolKeys struct {
	Name       *string `json:"name,omitempty"`
	Passphrase *string `json:"passphrase,omitempty"`
}

// NewPoolKeys instantiates a new PoolKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolKeys() *PoolKeys {
	this := PoolKeys{}
	return &this
}

// NewPoolKeysWithDefaults instantiates a new PoolKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolKeysWithDefaults() *PoolKeys {
	this := PoolKeys{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PoolKeys) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolKeys) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PoolKeys) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PoolKeys) SetName(v string) {
	o.Name = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *PoolKeys) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolKeys) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *PoolKeys) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *PoolKeys) SetPassphrase(v string) {
	o.Passphrase = &v
}

func (o PoolKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Passphrase != nil {
		toSerialize["passphrase"] = o.Passphrase
	}
	return json.Marshal(toSerialize)
}

type NullablePoolKeys struct {
	value *PoolKeys
	isSet bool
}

func (v NullablePoolKeys) Get() *PoolKeys {
	return v.value
}

func (v *NullablePoolKeys) Set(val *PoolKeys) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolKeys) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolKeys(val *PoolKeys) *NullablePoolKeys {
	return &NullablePoolKeys{value: val, isSet: true}
}

func (v NullablePoolKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
