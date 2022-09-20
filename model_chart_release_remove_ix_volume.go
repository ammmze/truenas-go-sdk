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

// ChartReleaseRemoveIxVolume struct for ChartReleaseRemoveIxVolume
type ChartReleaseRemoveIxVolume struct {
	ReleaseName          *string `json:"release_name,omitempty"`
	VolumeName           *string `json:"volume_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChartReleaseRemoveIxVolume ChartReleaseRemoveIxVolume

// NewChartReleaseRemoveIxVolume instantiates a new ChartReleaseRemoveIxVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartReleaseRemoveIxVolume() *ChartReleaseRemoveIxVolume {
	this := ChartReleaseRemoveIxVolume{}
	return &this
}

// NewChartReleaseRemoveIxVolumeWithDefaults instantiates a new ChartReleaseRemoveIxVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartReleaseRemoveIxVolumeWithDefaults() *ChartReleaseRemoveIxVolume {
	this := ChartReleaseRemoveIxVolume{}
	return &this
}

// GetReleaseName returns the ReleaseName field value if set, zero value otherwise.
func (o *ChartReleaseRemoveIxVolume) GetReleaseName() string {
	if o == nil || o.ReleaseName == nil {
		var ret string
		return ret
	}
	return *o.ReleaseName
}

// GetReleaseNameOk returns a tuple with the ReleaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartReleaseRemoveIxVolume) GetReleaseNameOk() (*string, bool) {
	if o == nil || o.ReleaseName == nil {
		return nil, false
	}
	return o.ReleaseName, true
}

// HasReleaseName returns a boolean if a field has been set.
func (o *ChartReleaseRemoveIxVolume) HasReleaseName() bool {
	if o != nil && o.ReleaseName != nil {
		return true
	}

	return false
}

// SetReleaseName gets a reference to the given string and assigns it to the ReleaseName field.
func (o *ChartReleaseRemoveIxVolume) SetReleaseName(v string) {
	o.ReleaseName = &v
}

// GetVolumeName returns the VolumeName field value if set, zero value otherwise.
func (o *ChartReleaseRemoveIxVolume) GetVolumeName() string {
	if o == nil || o.VolumeName == nil {
		var ret string
		return ret
	}
	return *o.VolumeName
}

// GetVolumeNameOk returns a tuple with the VolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartReleaseRemoveIxVolume) GetVolumeNameOk() (*string, bool) {
	if o == nil || o.VolumeName == nil {
		return nil, false
	}
	return o.VolumeName, true
}

// HasVolumeName returns a boolean if a field has been set.
func (o *ChartReleaseRemoveIxVolume) HasVolumeName() bool {
	if o != nil && o.VolumeName != nil {
		return true
	}

	return false
}

// SetVolumeName gets a reference to the given string and assigns it to the VolumeName field.
func (o *ChartReleaseRemoveIxVolume) SetVolumeName(v string) {
	o.VolumeName = &v
}

func (o ChartReleaseRemoveIxVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReleaseName != nil {
		toSerialize["release_name"] = o.ReleaseName
	}
	if o.VolumeName != nil {
		toSerialize["volume_name"] = o.VolumeName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChartReleaseRemoveIxVolume) UnmarshalJSON(bytes []byte) (err error) {
	varChartReleaseRemoveIxVolume := _ChartReleaseRemoveIxVolume{}

	if err = json.Unmarshal(bytes, &varChartReleaseRemoveIxVolume); err == nil {
		*o = ChartReleaseRemoveIxVolume(varChartReleaseRemoveIxVolume)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "release_name")
		delete(additionalProperties, "volume_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChartReleaseRemoveIxVolume struct {
	value *ChartReleaseRemoveIxVolume
	isSet bool
}

func (v NullableChartReleaseRemoveIxVolume) Get() *ChartReleaseRemoveIxVolume {
	return v.value
}

func (v *NullableChartReleaseRemoveIxVolume) Set(val *ChartReleaseRemoveIxVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableChartReleaseRemoveIxVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableChartReleaseRemoveIxVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartReleaseRemoveIxVolume(val *ChartReleaseRemoveIxVolume) *NullableChartReleaseRemoveIxVolume {
	return &NullableChartReleaseRemoveIxVolume{value: val, isSet: true}
}

func (v NullableChartReleaseRemoveIxVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartReleaseRemoveIxVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
