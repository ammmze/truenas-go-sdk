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

// IscsiGlobalUpdate0 struct for IscsiGlobalUpdate0
type IscsiGlobalUpdate0 struct {
	Basename           *string       `json:"basename,omitempty"`
	IsnsServers        []string      `json:"isns_servers,omitempty"`
	PoolAvailThreshold NullableInt32 `json:"pool_avail_threshold,omitempty"`
	Alua               *bool         `json:"alua,omitempty"`
}

// NewIscsiGlobalUpdate0 instantiates a new IscsiGlobalUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIscsiGlobalUpdate0() *IscsiGlobalUpdate0 {
	this := IscsiGlobalUpdate0{}
	return &this
}

// NewIscsiGlobalUpdate0WithDefaults instantiates a new IscsiGlobalUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIscsiGlobalUpdate0WithDefaults() *IscsiGlobalUpdate0 {
	this := IscsiGlobalUpdate0{}
	return &this
}

// GetBasename returns the Basename field value if set, zero value otherwise.
func (o *IscsiGlobalUpdate0) GetBasename() string {
	if o == nil || o.Basename == nil {
		var ret string
		return ret
	}
	return *o.Basename
}

// GetBasenameOk returns a tuple with the Basename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiGlobalUpdate0) GetBasenameOk() (*string, bool) {
	if o == nil || o.Basename == nil {
		return nil, false
	}
	return o.Basename, true
}

// HasBasename returns a boolean if a field has been set.
func (o *IscsiGlobalUpdate0) HasBasename() bool {
	if o != nil && o.Basename != nil {
		return true
	}

	return false
}

// SetBasename gets a reference to the given string and assigns it to the Basename field.
func (o *IscsiGlobalUpdate0) SetBasename(v string) {
	o.Basename = &v
}

// GetIsnsServers returns the IsnsServers field value if set, zero value otherwise.
func (o *IscsiGlobalUpdate0) GetIsnsServers() []string {
	if o == nil || o.IsnsServers == nil {
		var ret []string
		return ret
	}
	return o.IsnsServers
}

// GetIsnsServersOk returns a tuple with the IsnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiGlobalUpdate0) GetIsnsServersOk() ([]string, bool) {
	if o == nil || o.IsnsServers == nil {
		return nil, false
	}
	return o.IsnsServers, true
}

// HasIsnsServers returns a boolean if a field has been set.
func (o *IscsiGlobalUpdate0) HasIsnsServers() bool {
	if o != nil && o.IsnsServers != nil {
		return true
	}

	return false
}

// SetIsnsServers gets a reference to the given []string and assigns it to the IsnsServers field.
func (o *IscsiGlobalUpdate0) SetIsnsServers(v []string) {
	o.IsnsServers = v
}

// GetPoolAvailThreshold returns the PoolAvailThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IscsiGlobalUpdate0) GetPoolAvailThreshold() int32 {
	if o == nil || o.PoolAvailThreshold.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PoolAvailThreshold.Get()
}

// GetPoolAvailThresholdOk returns a tuple with the PoolAvailThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IscsiGlobalUpdate0) GetPoolAvailThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoolAvailThreshold.Get(), o.PoolAvailThreshold.IsSet()
}

// HasPoolAvailThreshold returns a boolean if a field has been set.
func (o *IscsiGlobalUpdate0) HasPoolAvailThreshold() bool {
	if o != nil && o.PoolAvailThreshold.IsSet() {
		return true
	}

	return false
}

// SetPoolAvailThreshold gets a reference to the given NullableInt32 and assigns it to the PoolAvailThreshold field.
func (o *IscsiGlobalUpdate0) SetPoolAvailThreshold(v int32) {
	o.PoolAvailThreshold.Set(&v)
}

// SetPoolAvailThresholdNil sets the value for PoolAvailThreshold to be an explicit nil
func (o *IscsiGlobalUpdate0) SetPoolAvailThresholdNil() {
	o.PoolAvailThreshold.Set(nil)
}

// UnsetPoolAvailThreshold ensures that no value is present for PoolAvailThreshold, not even an explicit nil
func (o *IscsiGlobalUpdate0) UnsetPoolAvailThreshold() {
	o.PoolAvailThreshold.Unset()
}

// GetAlua returns the Alua field value if set, zero value otherwise.
func (o *IscsiGlobalUpdate0) GetAlua() bool {
	if o == nil || o.Alua == nil {
		var ret bool
		return ret
	}
	return *o.Alua
}

// GetAluaOk returns a tuple with the Alua field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiGlobalUpdate0) GetAluaOk() (*bool, bool) {
	if o == nil || o.Alua == nil {
		return nil, false
	}
	return o.Alua, true
}

// HasAlua returns a boolean if a field has been set.
func (o *IscsiGlobalUpdate0) HasAlua() bool {
	if o != nil && o.Alua != nil {
		return true
	}

	return false
}

// SetAlua gets a reference to the given bool and assigns it to the Alua field.
func (o *IscsiGlobalUpdate0) SetAlua(v bool) {
	o.Alua = &v
}

func (o IscsiGlobalUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Basename != nil {
		toSerialize["basename"] = o.Basename
	}
	if o.IsnsServers != nil {
		toSerialize["isns_servers"] = o.IsnsServers
	}
	if o.PoolAvailThreshold.IsSet() {
		toSerialize["pool_avail_threshold"] = o.PoolAvailThreshold.Get()
	}
	if o.Alua != nil {
		toSerialize["alua"] = o.Alua
	}
	return json.Marshal(toSerialize)
}

type NullableIscsiGlobalUpdate0 struct {
	value *IscsiGlobalUpdate0
	isSet bool
}

func (v NullableIscsiGlobalUpdate0) Get() *IscsiGlobalUpdate0 {
	return v.value
}

func (v *NullableIscsiGlobalUpdate0) Set(val *IscsiGlobalUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableIscsiGlobalUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableIscsiGlobalUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIscsiGlobalUpdate0(val *IscsiGlobalUpdate0) *NullableIscsiGlobalUpdate0 {
	return &NullableIscsiGlobalUpdate0{value: val, isSet: true}
}

func (v NullableIscsiGlobalUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIscsiGlobalUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
