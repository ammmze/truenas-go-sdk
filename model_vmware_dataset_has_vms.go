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

// VmwareDatasetHasVms struct for VmwareDatasetHasVms
type VmwareDatasetHasVms struct {
	Dataset              *string `json:"dataset,omitempty"`
	Recursive            *bool   `json:"recursive,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VmwareDatasetHasVms VmwareDatasetHasVms

// NewVmwareDatasetHasVms instantiates a new VmwareDatasetHasVms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareDatasetHasVms() *VmwareDatasetHasVms {
	this := VmwareDatasetHasVms{}
	return &this
}

// NewVmwareDatasetHasVmsWithDefaults instantiates a new VmwareDatasetHasVms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareDatasetHasVmsWithDefaults() *VmwareDatasetHasVms {
	this := VmwareDatasetHasVms{}
	return &this
}

// GetDataset returns the Dataset field value if set, zero value otherwise.
func (o *VmwareDatasetHasVms) GetDataset() string {
	if o == nil || o.Dataset == nil {
		var ret string
		return ret
	}
	return *o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareDatasetHasVms) GetDatasetOk() (*string, bool) {
	if o == nil || o.Dataset == nil {
		return nil, false
	}
	return o.Dataset, true
}

// HasDataset returns a boolean if a field has been set.
func (o *VmwareDatasetHasVms) HasDataset() bool {
	if o != nil && o.Dataset != nil {
		return true
	}

	return false
}

// SetDataset gets a reference to the given string and assigns it to the Dataset field.
func (o *VmwareDatasetHasVms) SetDataset(v string) {
	o.Dataset = &v
}

// GetRecursive returns the Recursive field value if set, zero value otherwise.
func (o *VmwareDatasetHasVms) GetRecursive() bool {
	if o == nil || o.Recursive == nil {
		var ret bool
		return ret
	}
	return *o.Recursive
}

// GetRecursiveOk returns a tuple with the Recursive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareDatasetHasVms) GetRecursiveOk() (*bool, bool) {
	if o == nil || o.Recursive == nil {
		return nil, false
	}
	return o.Recursive, true
}

// HasRecursive returns a boolean if a field has been set.
func (o *VmwareDatasetHasVms) HasRecursive() bool {
	if o != nil && o.Recursive != nil {
		return true
	}

	return false
}

// SetRecursive gets a reference to the given bool and assigns it to the Recursive field.
func (o *VmwareDatasetHasVms) SetRecursive(v bool) {
	o.Recursive = &v
}

func (o VmwareDatasetHasVms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dataset != nil {
		toSerialize["dataset"] = o.Dataset
	}
	if o.Recursive != nil {
		toSerialize["recursive"] = o.Recursive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VmwareDatasetHasVms) UnmarshalJSON(bytes []byte) (err error) {
	varVmwareDatasetHasVms := _VmwareDatasetHasVms{}

	if err = json.Unmarshal(bytes, &varVmwareDatasetHasVms); err == nil {
		*o = VmwareDatasetHasVms(varVmwareDatasetHasVms)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dataset")
		delete(additionalProperties, "recursive")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVmwareDatasetHasVms struct {
	value *VmwareDatasetHasVms
	isSet bool
}

func (v NullableVmwareDatasetHasVms) Get() *VmwareDatasetHasVms {
	return v.value
}

func (v *NullableVmwareDatasetHasVms) Set(val *VmwareDatasetHasVms) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareDatasetHasVms) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareDatasetHasVms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareDatasetHasVms(val *VmwareDatasetHasVms) *NullableVmwareDatasetHasVms {
	return &NullableVmwareDatasetHasVms{value: val, isSet: true}
}

func (v NullableVmwareDatasetHasVms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareDatasetHasVms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
