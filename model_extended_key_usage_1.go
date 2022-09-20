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

// ExtendedKeyUsage1 struct for ExtendedKeyUsage1
type ExtendedKeyUsage1 struct {
	Usages            []string `json:"usages,omitempty"`
	Enabled           *bool    `json:"enabled,omitempty"`
	ExtensionCritical *bool    `json:"extension_critical,omitempty"`
}

// NewExtendedKeyUsage1 instantiates a new ExtendedKeyUsage1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendedKeyUsage1() *ExtendedKeyUsage1 {
	this := ExtendedKeyUsage1{}
	var enabled bool = true
	this.Enabled = &enabled
	var extensionCritical bool = false
	this.ExtensionCritical = &extensionCritical
	return &this
}

// NewExtendedKeyUsage1WithDefaults instantiates a new ExtendedKeyUsage1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendedKeyUsage1WithDefaults() *ExtendedKeyUsage1 {
	this := ExtendedKeyUsage1{}
	var enabled bool = true
	this.Enabled = &enabled
	var extensionCritical bool = false
	this.ExtensionCritical = &extensionCritical
	return &this
}

// GetUsages returns the Usages field value if set, zero value otherwise.
func (o *ExtendedKeyUsage1) GetUsages() []string {
	if o == nil || o.Usages == nil {
		var ret []string
		return ret
	}
	return o.Usages
}

// GetUsagesOk returns a tuple with the Usages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedKeyUsage1) GetUsagesOk() ([]string, bool) {
	if o == nil || o.Usages == nil {
		return nil, false
	}
	return o.Usages, true
}

// HasUsages returns a boolean if a field has been set.
func (o *ExtendedKeyUsage1) HasUsages() bool {
	if o != nil && o.Usages != nil {
		return true
	}

	return false
}

// SetUsages gets a reference to the given []string and assigns it to the Usages field.
func (o *ExtendedKeyUsage1) SetUsages(v []string) {
	o.Usages = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ExtendedKeyUsage1) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedKeyUsage1) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ExtendedKeyUsage1) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ExtendedKeyUsage1) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExtensionCritical returns the ExtensionCritical field value if set, zero value otherwise.
func (o *ExtendedKeyUsage1) GetExtensionCritical() bool {
	if o == nil || o.ExtensionCritical == nil {
		var ret bool
		return ret
	}
	return *o.ExtensionCritical
}

// GetExtensionCriticalOk returns a tuple with the ExtensionCritical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedKeyUsage1) GetExtensionCriticalOk() (*bool, bool) {
	if o == nil || o.ExtensionCritical == nil {
		return nil, false
	}
	return o.ExtensionCritical, true
}

// HasExtensionCritical returns a boolean if a field has been set.
func (o *ExtendedKeyUsage1) HasExtensionCritical() bool {
	if o != nil && o.ExtensionCritical != nil {
		return true
	}

	return false
}

// SetExtensionCritical gets a reference to the given bool and assigns it to the ExtensionCritical field.
func (o *ExtendedKeyUsage1) SetExtensionCritical(v bool) {
	o.ExtensionCritical = &v
}

func (o ExtendedKeyUsage1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Usages != nil {
		toSerialize["usages"] = o.Usages
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ExtensionCritical != nil {
		toSerialize["extension_critical"] = o.ExtensionCritical
	}
	return json.Marshal(toSerialize)
}

type NullableExtendedKeyUsage1 struct {
	value *ExtendedKeyUsage1
	isSet bool
}

func (v NullableExtendedKeyUsage1) Get() *ExtendedKeyUsage1 {
	return v.value
}

func (v *NullableExtendedKeyUsage1) Set(val *ExtendedKeyUsage1) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendedKeyUsage1) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendedKeyUsage1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendedKeyUsage1(val *ExtendedKeyUsage1) *NullableExtendedKeyUsage1 {
	return &NullableExtendedKeyUsage1{value: val, isSet: true}
}

func (v NullableExtendedKeyUsage1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendedKeyUsage1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
