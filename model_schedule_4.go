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

// Schedule4 * `auto` allows replication to run automatically on schedule or after bound periodic snapshot task * `schedule` is a schedule to run replication task. Only `auto` replication tasks without bound periodic    snapshot tasks can have a schedule * Enabling `only_matching_schedule` will only replicate snapshots that match `schedule` or    `restrict_schedule`
type Schedule4 struct {
	Minute *string `json:"minute,omitempty"`
	Hour   *string `json:"hour,omitempty"`
	Dom    *string `json:"dom,omitempty"`
	Month  *string `json:"month,omitempty"`
	Dow    *string `json:"dow,omitempty"`
	Begin  *string `json:"begin,omitempty"`
	End    *string `json:"end,omitempty"`
}

// NewSchedule4 instantiates a new Schedule4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule4() *Schedule4 {
	this := Schedule4{}
	var minute string = "00"
	this.Minute = &minute
	var hour string = "*"
	this.Hour = &hour
	var dom string = "*"
	this.Dom = &dom
	var month string = "*"
	this.Month = &month
	var dow string = "*"
	this.Dow = &dow
	var begin string = "00:00"
	this.Begin = &begin
	var end string = "1439"
	this.End = &end
	return &this
}

// NewSchedule4WithDefaults instantiates a new Schedule4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedule4WithDefaults() *Schedule4 {
	this := Schedule4{}
	var minute string = "00"
	this.Minute = &minute
	var hour string = "*"
	this.Hour = &hour
	var dom string = "*"
	this.Dom = &dom
	var month string = "*"
	this.Month = &month
	var dow string = "*"
	this.Dow = &dow
	var begin string = "00:00"
	this.Begin = &begin
	var end string = "1439"
	this.End = &end
	return &this
}

// GetMinute returns the Minute field value if set, zero value otherwise.
func (o *Schedule4) GetMinute() string {
	if o == nil || o.Minute == nil {
		var ret string
		return ret
	}
	return *o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule4) GetMinuteOk() (*string, bool) {
	if o == nil || o.Minute == nil {
		return nil, false
	}
	return o.Minute, true
}

// HasMinute returns a boolean if a field has been set.
func (o *Schedule4) HasMinute() bool {
	if o != nil && o.Minute != nil {
		return true
	}

	return false
}

// SetMinute gets a reference to the given string and assigns it to the Minute field.
func (o *Schedule4) SetMinute(v string) {
	o.Minute = &v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *Schedule4) GetHour() string {
	if o == nil || o.Hour == nil {
		var ret string
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule4) GetHourOk() (*string, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *Schedule4) HasHour() bool {
	if o != nil && o.Hour != nil {
		return true
	}

	return false
}

// SetHour gets a reference to the given string and assigns it to the Hour field.
func (o *Schedule4) SetHour(v string) {
	o.Hour = &v
}

// GetDom returns the Dom field value if set, zero value otherwise.
func (o *Schedule4) GetDom() string {
	if o == nil || o.Dom == nil {
		var ret string
		return ret
	}
	return *o.Dom
}

// GetDomOk returns a tuple with the Dom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule4) GetDomOk() (*string, bool) {
	if o == nil || o.Dom == nil {
		return nil, false
	}
	return o.Dom, true
}

// HasDom returns a boolean if a field has been set.
func (o *Schedule4) HasDom() bool {
	if o != nil && o.Dom != nil {
		return true
	}

	return false
}

// SetDom gets a reference to the given string and assigns it to the Dom field.
func (o *Schedule4) SetDom(v string) {
	o.Dom = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *Schedule4) GetMonth() string {
	if o == nil || o.Month == nil {
		var ret string
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule4) GetMonthOk() (*string, bool) {
	if o == nil || o.Month == nil {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *Schedule4) HasMonth() bool {
	if o != nil && o.Month != nil {
		return true
	}

	return false
}

// SetMonth gets a reference to the given string and assigns it to the Month field.
func (o *Schedule4) SetMonth(v string) {
	o.Month = &v
}

// GetDow returns the Dow field value if set, zero value otherwise.
func (o *Schedule4) GetDow() string {
	if o == nil || o.Dow == nil {
		var ret string
		return ret
	}
	return *o.Dow
}

// GetDowOk returns a tuple with the Dow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule4) GetDowOk() (*string, bool) {
	if o == nil || o.Dow == nil {
		return nil, false
	}
	return o.Dow, true
}

// HasDow returns a boolean if a field has been set.
func (o *Schedule4) HasDow() bool {
	if o != nil && o.Dow != nil {
		return true
	}

	return false
}

// SetDow gets a reference to the given string and assigns it to the Dow field.
func (o *Schedule4) SetDow(v string) {
	o.Dow = &v
}

// GetBegin returns the Begin field value if set, zero value otherwise.
func (o *Schedule4) GetBegin() string {
	if o == nil || o.Begin == nil {
		var ret string
		return ret
	}
	return *o.Begin
}

// GetBeginOk returns a tuple with the Begin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule4) GetBeginOk() (*string, bool) {
	if o == nil || o.Begin == nil {
		return nil, false
	}
	return o.Begin, true
}

// HasBegin returns a boolean if a field has been set.
func (o *Schedule4) HasBegin() bool {
	if o != nil && o.Begin != nil {
		return true
	}

	return false
}

// SetBegin gets a reference to the given string and assigns it to the Begin field.
func (o *Schedule4) SetBegin(v string) {
	o.Begin = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *Schedule4) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule4) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Schedule4) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *Schedule4) SetEnd(v string) {
	o.End = &v
}

func (o Schedule4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Minute != nil {
		toSerialize["minute"] = o.Minute
	}
	if o.Hour != nil {
		toSerialize["hour"] = o.Hour
	}
	if o.Dom != nil {
		toSerialize["dom"] = o.Dom
	}
	if o.Month != nil {
		toSerialize["month"] = o.Month
	}
	if o.Dow != nil {
		toSerialize["dow"] = o.Dow
	}
	if o.Begin != nil {
		toSerialize["begin"] = o.Begin
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	return json.Marshal(toSerialize)
}

type NullableSchedule4 struct {
	value *Schedule4
	isSet bool
}

func (v NullableSchedule4) Get() *Schedule4 {
	return v.value
}

func (v *NullableSchedule4) Set(val *Schedule4) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule4) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule4(val *Schedule4) *NullableSchedule4 {
	return &NullableSchedule4{value: val, isSet: true}
}

func (v NullableSchedule4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
