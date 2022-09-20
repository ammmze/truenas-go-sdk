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

// PoolDatasetDestroySnapshots1 struct for PoolDatasetDestroySnapshots1
type PoolDatasetDestroySnapshots1 struct {
	All       *bool            `json:"all,omitempty"`
	Recursive *bool            `json:"recursive,omitempty"`
	Snapshots []SnapshotsInner `json:"snapshots,omitempty"`
}

// NewPoolDatasetDestroySnapshots1 instantiates a new PoolDatasetDestroySnapshots1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolDatasetDestroySnapshots1() *PoolDatasetDestroySnapshots1 {
	this := PoolDatasetDestroySnapshots1{}
	var all bool = true
	this.All = &all
	var recursive bool = false
	this.Recursive = &recursive
	return &this
}

// NewPoolDatasetDestroySnapshots1WithDefaults instantiates a new PoolDatasetDestroySnapshots1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolDatasetDestroySnapshots1WithDefaults() *PoolDatasetDestroySnapshots1 {
	this := PoolDatasetDestroySnapshots1{}
	var all bool = true
	this.All = &all
	var recursive bool = false
	this.Recursive = &recursive
	return &this
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *PoolDatasetDestroySnapshots1) GetAll() bool {
	if o == nil || o.All == nil {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetDestroySnapshots1) GetAllOk() (*bool, bool) {
	if o == nil || o.All == nil {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *PoolDatasetDestroySnapshots1) HasAll() bool {
	if o != nil && o.All != nil {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *PoolDatasetDestroySnapshots1) SetAll(v bool) {
	o.All = &v
}

// GetRecursive returns the Recursive field value if set, zero value otherwise.
func (o *PoolDatasetDestroySnapshots1) GetRecursive() bool {
	if o == nil || o.Recursive == nil {
		var ret bool
		return ret
	}
	return *o.Recursive
}

// GetRecursiveOk returns a tuple with the Recursive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetDestroySnapshots1) GetRecursiveOk() (*bool, bool) {
	if o == nil || o.Recursive == nil {
		return nil, false
	}
	return o.Recursive, true
}

// HasRecursive returns a boolean if a field has been set.
func (o *PoolDatasetDestroySnapshots1) HasRecursive() bool {
	if o != nil && o.Recursive != nil {
		return true
	}

	return false
}

// SetRecursive gets a reference to the given bool and assigns it to the Recursive field.
func (o *PoolDatasetDestroySnapshots1) SetRecursive(v bool) {
	o.Recursive = &v
}

// GetSnapshots returns the Snapshots field value if set, zero value otherwise.
func (o *PoolDatasetDestroySnapshots1) GetSnapshots() []SnapshotsInner {
	if o == nil || o.Snapshots == nil {
		var ret []SnapshotsInner
		return ret
	}
	return o.Snapshots
}

// GetSnapshotsOk returns a tuple with the Snapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetDestroySnapshots1) GetSnapshotsOk() ([]SnapshotsInner, bool) {
	if o == nil || o.Snapshots == nil {
		return nil, false
	}
	return o.Snapshots, true
}

// HasSnapshots returns a boolean if a field has been set.
func (o *PoolDatasetDestroySnapshots1) HasSnapshots() bool {
	if o != nil && o.Snapshots != nil {
		return true
	}

	return false
}

// SetSnapshots gets a reference to the given []SnapshotsInner and assigns it to the Snapshots field.
func (o *PoolDatasetDestroySnapshots1) SetSnapshots(v []SnapshotsInner) {
	o.Snapshots = v
}

func (o PoolDatasetDestroySnapshots1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.All != nil {
		toSerialize["all"] = o.All
	}
	if o.Recursive != nil {
		toSerialize["recursive"] = o.Recursive
	}
	if o.Snapshots != nil {
		toSerialize["snapshots"] = o.Snapshots
	}
	return json.Marshal(toSerialize)
}

type NullablePoolDatasetDestroySnapshots1 struct {
	value *PoolDatasetDestroySnapshots1
	isSet bool
}

func (v NullablePoolDatasetDestroySnapshots1) Get() *PoolDatasetDestroySnapshots1 {
	return v.value
}

func (v *NullablePoolDatasetDestroySnapshots1) Set(val *PoolDatasetDestroySnapshots1) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolDatasetDestroySnapshots1) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolDatasetDestroySnapshots1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolDatasetDestroySnapshots1(val *PoolDatasetDestroySnapshots1) *NullablePoolDatasetDestroySnapshots1 {
	return &NullablePoolDatasetDestroySnapshots1{value: val, isSet: true}
}

func (v NullablePoolDatasetDestroySnapshots1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolDatasetDestroySnapshots1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
