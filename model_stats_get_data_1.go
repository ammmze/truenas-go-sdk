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

// StatsGetData1 struct for StatsGetData1
type StatsGetData1 struct {
	Step  *int32  `json:"step,omitempty"`
	Start *string `json:"start,omitempty"`
	End   *string `json:"end,omitempty"`
}

// NewStatsGetData1 instantiates a new StatsGetData1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsGetData1() *StatsGetData1 {
	this := StatsGetData1{}
	var step int32 = 10
	this.Step = &step
	var start string = "now-1h"
	this.Start = &start
	var end string = "now"
	this.End = &end
	return &this
}

// NewStatsGetData1WithDefaults instantiates a new StatsGetData1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsGetData1WithDefaults() *StatsGetData1 {
	this := StatsGetData1{}
	var step int32 = 10
	this.Step = &step
	var start string = "now-1h"
	this.Start = &start
	var end string = "now"
	this.End = &end
	return &this
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *StatsGetData1) GetStep() int32 {
	if o == nil || o.Step == nil {
		var ret int32
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsGetData1) GetStepOk() (*int32, bool) {
	if o == nil || o.Step == nil {
		return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *StatsGetData1) HasStep() bool {
	if o != nil && o.Step != nil {
		return true
	}

	return false
}

// SetStep gets a reference to the given int32 and assigns it to the Step field.
func (o *StatsGetData1) SetStep(v int32) {
	o.Step = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *StatsGetData1) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsGetData1) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *StatsGetData1) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *StatsGetData1) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *StatsGetData1) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsGetData1) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *StatsGetData1) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *StatsGetData1) SetEnd(v string) {
	o.End = &v
}

func (o StatsGetData1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Step != nil {
		toSerialize["step"] = o.Step
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	return json.Marshal(toSerialize)
}

type NullableStatsGetData1 struct {
	value *StatsGetData1
	isSet bool
}

func (v NullableStatsGetData1) Get() *StatsGetData1 {
	return v.value
}

func (v *NullableStatsGetData1) Set(val *StatsGetData1) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsGetData1) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsGetData1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsGetData1(val *StatsGetData1) *NullableStatsGetData1 {
	return &NullableStatsGetData1{value: val, isSet: true}
}

func (v NullableStatsGetData1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsGetData1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
