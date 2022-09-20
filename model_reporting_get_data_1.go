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

// ReportingGetData1 struct for ReportingGetData1
type ReportingGetData1 struct {
	// For the time period of the graph either `unit` and `page` OR `start` and `end` should be used, not both.
	Unit *string `json:"unit,omitempty"`
	// For the time period of the graph either `unit` and `page` OR `start` and `end` should be used, not both.
	Page *int32 `json:"page,omitempty"`
	// For the time period of the graph either `unit` and `page` OR `start` and `end` should be used, not both.
	Start *string `json:"start,omitempty"`
	// For the time period of the graph either `unit` and `page` OR `start` and `end` should be used, not both.
	End *string `json:"end,omitempty"`
	// `aggregate` will return aggregate available data for each graph (e.g. min, max, mean).
	Aggregate *bool `json:"aggregate,omitempty"`
}

// NewReportingGetData1 instantiates a new ReportingGetData1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingGetData1() *ReportingGetData1 {
	this := ReportingGetData1{}
	var page int32 = 0
	this.Page = &page
	var aggregate bool = true
	this.Aggregate = &aggregate
	return &this
}

// NewReportingGetData1WithDefaults instantiates a new ReportingGetData1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingGetData1WithDefaults() *ReportingGetData1 {
	this := ReportingGetData1{}
	var page int32 = 0
	this.Page = &page
	var aggregate bool = true
	this.Aggregate = &aggregate
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ReportingGetData1) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingGetData1) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ReportingGetData1) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *ReportingGetData1) SetUnit(v string) {
	o.Unit = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ReportingGetData1) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingGetData1) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ReportingGetData1) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ReportingGetData1) SetPage(v int32) {
	o.Page = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ReportingGetData1) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingGetData1) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ReportingGetData1) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *ReportingGetData1) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ReportingGetData1) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingGetData1) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ReportingGetData1) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *ReportingGetData1) SetEnd(v string) {
	o.End = &v
}

// GetAggregate returns the Aggregate field value if set, zero value otherwise.
func (o *ReportingGetData1) GetAggregate() bool {
	if o == nil || o.Aggregate == nil {
		var ret bool
		return ret
	}
	return *o.Aggregate
}

// GetAggregateOk returns a tuple with the Aggregate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingGetData1) GetAggregateOk() (*bool, bool) {
	if o == nil || o.Aggregate == nil {
		return nil, false
	}
	return o.Aggregate, true
}

// HasAggregate returns a boolean if a field has been set.
func (o *ReportingGetData1) HasAggregate() bool {
	if o != nil && o.Aggregate != nil {
		return true
	}

	return false
}

// SetAggregate gets a reference to the given bool and assigns it to the Aggregate field.
func (o *ReportingGetData1) SetAggregate(v bool) {
	o.Aggregate = &v
}

func (o ReportingGetData1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Aggregate != nil {
		toSerialize["aggregate"] = o.Aggregate
	}
	return json.Marshal(toSerialize)
}

type NullableReportingGetData1 struct {
	value *ReportingGetData1
	isSet bool
}

func (v NullableReportingGetData1) Get() *ReportingGetData1 {
	return v.value
}

func (v *NullableReportingGetData1) Set(val *ReportingGetData1) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingGetData1) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingGetData1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingGetData1(val *ReportingGetData1) *NullableReportingGetData1 {
	return &NullableReportingGetData1{value: val, isSet: true}
}

func (v NullableReportingGetData1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingGetData1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
