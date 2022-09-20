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

// ScaleWorkload struct for ScaleWorkload
type ScaleWorkload struct {
	ReplicaCount *int32  `json:"replica_count,omitempty"`
	Type         *string `json:"type,omitempty"`
	Name         *string `json:"name,omitempty"`
}

// NewScaleWorkload instantiates a new ScaleWorkload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScaleWorkload() *ScaleWorkload {
	this := ScaleWorkload{}
	return &this
}

// NewScaleWorkloadWithDefaults instantiates a new ScaleWorkload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScaleWorkloadWithDefaults() *ScaleWorkload {
	this := ScaleWorkload{}
	return &this
}

// GetReplicaCount returns the ReplicaCount field value if set, zero value otherwise.
func (o *ScaleWorkload) GetReplicaCount() int32 {
	if o == nil || o.ReplicaCount == nil {
		var ret int32
		return ret
	}
	return *o.ReplicaCount
}

// GetReplicaCountOk returns a tuple with the ReplicaCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleWorkload) GetReplicaCountOk() (*int32, bool) {
	if o == nil || o.ReplicaCount == nil {
		return nil, false
	}
	return o.ReplicaCount, true
}

// HasReplicaCount returns a boolean if a field has been set.
func (o *ScaleWorkload) HasReplicaCount() bool {
	if o != nil && o.ReplicaCount != nil {
		return true
	}

	return false
}

// SetReplicaCount gets a reference to the given int32 and assigns it to the ReplicaCount field.
func (o *ScaleWorkload) SetReplicaCount(v int32) {
	o.ReplicaCount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ScaleWorkload) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleWorkload) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ScaleWorkload) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ScaleWorkload) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScaleWorkload) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleWorkload) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScaleWorkload) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScaleWorkload) SetName(v string) {
	o.Name = &v
}

func (o ScaleWorkload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReplicaCount != nil {
		toSerialize["replica_count"] = o.ReplicaCount
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableScaleWorkload struct {
	value *ScaleWorkload
	isSet bool
}

func (v NullableScaleWorkload) Get() *ScaleWorkload {
	return v.value
}

func (v *NullableScaleWorkload) Set(val *ScaleWorkload) {
	v.value = val
	v.isSet = true
}

func (v NullableScaleWorkload) IsSet() bool {
	return v.isSet
}

func (v *NullableScaleWorkload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScaleWorkload(val *ScaleWorkload) *NullableScaleWorkload {
	return &NullableScaleWorkload{value: val, isSet: true}
}

func (v NullableScaleWorkload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScaleWorkload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
