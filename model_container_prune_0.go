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

// ContainerPrune0 struct for ContainerPrune0
type ContainerPrune0 struct {
	RemoveUnusedImages      *bool `json:"remove_unused_images,omitempty"`
	RemoveStoppedContainers *bool `json:"remove_stopped_containers,omitempty"`
}

// NewContainerPrune0 instantiates a new ContainerPrune0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerPrune0() *ContainerPrune0 {
	this := ContainerPrune0{}
	var removeUnusedImages bool = false
	this.RemoveUnusedImages = &removeUnusedImages
	var removeStoppedContainers bool = false
	this.RemoveStoppedContainers = &removeStoppedContainers
	return &this
}

// NewContainerPrune0WithDefaults instantiates a new ContainerPrune0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerPrune0WithDefaults() *ContainerPrune0 {
	this := ContainerPrune0{}
	var removeUnusedImages bool = false
	this.RemoveUnusedImages = &removeUnusedImages
	var removeStoppedContainers bool = false
	this.RemoveStoppedContainers = &removeStoppedContainers
	return &this
}

// GetRemoveUnusedImages returns the RemoveUnusedImages field value if set, zero value otherwise.
func (o *ContainerPrune0) GetRemoveUnusedImages() bool {
	if o == nil || o.RemoveUnusedImages == nil {
		var ret bool
		return ret
	}
	return *o.RemoveUnusedImages
}

// GetRemoveUnusedImagesOk returns a tuple with the RemoveUnusedImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPrune0) GetRemoveUnusedImagesOk() (*bool, bool) {
	if o == nil || o.RemoveUnusedImages == nil {
		return nil, false
	}
	return o.RemoveUnusedImages, true
}

// HasRemoveUnusedImages returns a boolean if a field has been set.
func (o *ContainerPrune0) HasRemoveUnusedImages() bool {
	if o != nil && o.RemoveUnusedImages != nil {
		return true
	}

	return false
}

// SetRemoveUnusedImages gets a reference to the given bool and assigns it to the RemoveUnusedImages field.
func (o *ContainerPrune0) SetRemoveUnusedImages(v bool) {
	o.RemoveUnusedImages = &v
}

// GetRemoveStoppedContainers returns the RemoveStoppedContainers field value if set, zero value otherwise.
func (o *ContainerPrune0) GetRemoveStoppedContainers() bool {
	if o == nil || o.RemoveStoppedContainers == nil {
		var ret bool
		return ret
	}
	return *o.RemoveStoppedContainers
}

// GetRemoveStoppedContainersOk returns a tuple with the RemoveStoppedContainers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPrune0) GetRemoveStoppedContainersOk() (*bool, bool) {
	if o == nil || o.RemoveStoppedContainers == nil {
		return nil, false
	}
	return o.RemoveStoppedContainers, true
}

// HasRemoveStoppedContainers returns a boolean if a field has been set.
func (o *ContainerPrune0) HasRemoveStoppedContainers() bool {
	if o != nil && o.RemoveStoppedContainers != nil {
		return true
	}

	return false
}

// SetRemoveStoppedContainers gets a reference to the given bool and assigns it to the RemoveStoppedContainers field.
func (o *ContainerPrune0) SetRemoveStoppedContainers(v bool) {
	o.RemoveStoppedContainers = &v
}

func (o ContainerPrune0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RemoveUnusedImages != nil {
		toSerialize["remove_unused_images"] = o.RemoveUnusedImages
	}
	if o.RemoveStoppedContainers != nil {
		toSerialize["remove_stopped_containers"] = o.RemoveStoppedContainers
	}
	return json.Marshal(toSerialize)
}

type NullableContainerPrune0 struct {
	value *ContainerPrune0
	isSet bool
}

func (v NullableContainerPrune0) Get() *ContainerPrune0 {
	return v.value
}

func (v *NullableContainerPrune0) Set(val *ContainerPrune0) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerPrune0) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerPrune0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerPrune0(val *ContainerPrune0) *NullableContainerPrune0 {
	return &NullableContainerPrune0{value: val, isSet: true}
}

func (v NullableContainerPrune0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerPrune0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
