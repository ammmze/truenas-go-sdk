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

// PoolImportDisk struct for PoolImportDisk
type PoolImportDisk struct {
	Device               *string                `json:"device,omitempty"`
	FsType               *string                `json:"fs_type,omitempty"`
	FsOptions            map[string]interface{} `json:"fs_options,omitempty"`
	DstPath              *string                `json:"dst_path,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoolImportDisk PoolImportDisk

// NewPoolImportDisk instantiates a new PoolImportDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolImportDisk() *PoolImportDisk {
	this := PoolImportDisk{}
	return &this
}

// NewPoolImportDiskWithDefaults instantiates a new PoolImportDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolImportDiskWithDefaults() *PoolImportDisk {
	this := PoolImportDisk{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *PoolImportDisk) GetDevice() string {
	if o == nil || o.Device == nil {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolImportDisk) GetDeviceOk() (*string, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *PoolImportDisk) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *PoolImportDisk) SetDevice(v string) {
	o.Device = &v
}

// GetFsType returns the FsType field value if set, zero value otherwise.
func (o *PoolImportDisk) GetFsType() string {
	if o == nil || o.FsType == nil {
		var ret string
		return ret
	}
	return *o.FsType
}

// GetFsTypeOk returns a tuple with the FsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolImportDisk) GetFsTypeOk() (*string, bool) {
	if o == nil || o.FsType == nil {
		return nil, false
	}
	return o.FsType, true
}

// HasFsType returns a boolean if a field has been set.
func (o *PoolImportDisk) HasFsType() bool {
	if o != nil && o.FsType != nil {
		return true
	}

	return false
}

// SetFsType gets a reference to the given string and assigns it to the FsType field.
func (o *PoolImportDisk) SetFsType(v string) {
	o.FsType = &v
}

// GetFsOptions returns the FsOptions field value if set, zero value otherwise.
func (o *PoolImportDisk) GetFsOptions() map[string]interface{} {
	if o == nil || o.FsOptions == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.FsOptions
}

// GetFsOptionsOk returns a tuple with the FsOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolImportDisk) GetFsOptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.FsOptions == nil {
		return nil, false
	}
	return o.FsOptions, true
}

// HasFsOptions returns a boolean if a field has been set.
func (o *PoolImportDisk) HasFsOptions() bool {
	if o != nil && o.FsOptions != nil {
		return true
	}

	return false
}

// SetFsOptions gets a reference to the given map[string]interface{} and assigns it to the FsOptions field.
func (o *PoolImportDisk) SetFsOptions(v map[string]interface{}) {
	o.FsOptions = v
}

// GetDstPath returns the DstPath field value if set, zero value otherwise.
func (o *PoolImportDisk) GetDstPath() string {
	if o == nil || o.DstPath == nil {
		var ret string
		return ret
	}
	return *o.DstPath
}

// GetDstPathOk returns a tuple with the DstPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolImportDisk) GetDstPathOk() (*string, bool) {
	if o == nil || o.DstPath == nil {
		return nil, false
	}
	return o.DstPath, true
}

// HasDstPath returns a boolean if a field has been set.
func (o *PoolImportDisk) HasDstPath() bool {
	if o != nil && o.DstPath != nil {
		return true
	}

	return false
}

// SetDstPath gets a reference to the given string and assigns it to the DstPath field.
func (o *PoolImportDisk) SetDstPath(v string) {
	o.DstPath = &v
}

func (o PoolImportDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.FsType != nil {
		toSerialize["fs_type"] = o.FsType
	}
	if o.FsOptions != nil {
		toSerialize["fs_options"] = o.FsOptions
	}
	if o.DstPath != nil {
		toSerialize["dst_path"] = o.DstPath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PoolImportDisk) UnmarshalJSON(bytes []byte) (err error) {
	varPoolImportDisk := _PoolImportDisk{}

	if err = json.Unmarshal(bytes, &varPoolImportDisk); err == nil {
		*o = PoolImportDisk(varPoolImportDisk)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "fs_type")
		delete(additionalProperties, "fs_options")
		delete(additionalProperties, "dst_path")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePoolImportDisk struct {
	value *PoolImportDisk
	isSet bool
}

func (v NullablePoolImportDisk) Get() *PoolImportDisk {
	return v.value
}

func (v *NullablePoolImportDisk) Set(val *PoolImportDisk) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolImportDisk) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolImportDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolImportDisk(val *PoolImportDisk) *NullablePoolImportDisk {
	return &NullablePoolImportDisk{value: val, isSet: true}
}

func (v NullablePoolImportDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolImportDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
