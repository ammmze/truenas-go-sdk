/*
TrueNAS RESTful API

Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
)

// CompositeValue struct for CompositeValue
type CompositeValue struct {
	Value                *string `json:"value,omitempty"`
	Rawvalue             string  `json:"rawvalue"`
	Source               *string `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CompositeValue CompositeValue

// NewCompositeValue instantiates a new CompositeValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompositeValue(rawvalue string) *CompositeValue {
	this := CompositeValue{}
	this.Rawvalue = rawvalue
	return &this
}

// NewCompositeValueWithDefaults instantiates a new CompositeValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompositeValueWithDefaults() *CompositeValue {
	this := CompositeValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CompositeValue) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeValue) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CompositeValue) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CompositeValue) SetValue(v string) {
	o.Value = &v
}

// GetRawvalue returns the Rawvalue field value
func (o *CompositeValue) GetRawvalue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rawvalue
}

// GetRawvalueOk returns a tuple with the Rawvalue field value
// and a boolean to check if the value has been set.
func (o *CompositeValue) GetRawvalueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rawvalue, true
}

// SetRawvalue sets field value
func (o *CompositeValue) SetRawvalue(v string) {
	o.Rawvalue = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CompositeValue) GetSource() string {
	if o == nil || isNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeValue) GetSourceOk() (*string, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CompositeValue) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *CompositeValue) SetSource(v string) {
	o.Source = &v
}

func (o CompositeValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["rawvalue"] = o.Rawvalue
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CompositeValue) UnmarshalJSON(bytes []byte) (err error) {
	varCompositeValue := _CompositeValue{}

	if err = json.Unmarshal(bytes, &varCompositeValue); err == nil {
		*o = CompositeValue(varCompositeValue)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "rawvalue")
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCompositeValue struct {
	value *CompositeValue
	isSet bool
}

func (v NullableCompositeValue) Get() *CompositeValue {
	return v.value
}

func (v *NullableCompositeValue) Set(val *CompositeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCompositeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCompositeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompositeValue(val *CompositeValue) *NullableCompositeValue {
	return &NullableCompositeValue{value: val, isSet: true}
}

func (v NullableCompositeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompositeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
