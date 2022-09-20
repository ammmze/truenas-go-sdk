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

// ChartReleasePullContainerImages1 struct for ChartReleasePullContainerImages1
type ChartReleasePullContainerImages1 struct {
	Redeploy *bool `json:"redeploy,omitempty"`
}

// NewChartReleasePullContainerImages1 instantiates a new ChartReleasePullContainerImages1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartReleasePullContainerImages1() *ChartReleasePullContainerImages1 {
	this := ChartReleasePullContainerImages1{}
	var redeploy bool = true
	this.Redeploy = &redeploy
	return &this
}

// NewChartReleasePullContainerImages1WithDefaults instantiates a new ChartReleasePullContainerImages1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartReleasePullContainerImages1WithDefaults() *ChartReleasePullContainerImages1 {
	this := ChartReleasePullContainerImages1{}
	var redeploy bool = true
	this.Redeploy = &redeploy
	return &this
}

// GetRedeploy returns the Redeploy field value if set, zero value otherwise.
func (o *ChartReleasePullContainerImages1) GetRedeploy() bool {
	if o == nil || o.Redeploy == nil {
		var ret bool
		return ret
	}
	return *o.Redeploy
}

// GetRedeployOk returns a tuple with the Redeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartReleasePullContainerImages1) GetRedeployOk() (*bool, bool) {
	if o == nil || o.Redeploy == nil {
		return nil, false
	}
	return o.Redeploy, true
}

// HasRedeploy returns a boolean if a field has been set.
func (o *ChartReleasePullContainerImages1) HasRedeploy() bool {
	if o != nil && o.Redeploy != nil {
		return true
	}

	return false
}

// SetRedeploy gets a reference to the given bool and assigns it to the Redeploy field.
func (o *ChartReleasePullContainerImages1) SetRedeploy(v bool) {
	o.Redeploy = &v
}

func (o ChartReleasePullContainerImages1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Redeploy != nil {
		toSerialize["redeploy"] = o.Redeploy
	}
	return json.Marshal(toSerialize)
}

type NullableChartReleasePullContainerImages1 struct {
	value *ChartReleasePullContainerImages1
	isSet bool
}

func (v NullableChartReleasePullContainerImages1) Get() *ChartReleasePullContainerImages1 {
	return v.value
}

func (v *NullableChartReleasePullContainerImages1) Set(val *ChartReleasePullContainerImages1) {
	v.value = val
	v.isSet = true
}

func (v NullableChartReleasePullContainerImages1) IsSet() bool {
	return v.isSet
}

func (v *NullableChartReleasePullContainerImages1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartReleasePullContainerImages1(val *ChartReleasePullContainerImages1) *NullableChartReleasePullContainerImages1 {
	return &NullableChartReleasePullContainerImages1{value: val, isSet: true}
}

func (v NullableChartReleasePullContainerImages1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartReleasePullContainerImages1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
