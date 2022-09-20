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

// AeWhoName struct for AeWhoName
type AeWhoName struct {
	Domain *string `json:"domain,omitempty"`
	Name   *string `json:"name,omitempty"`
}

// NewAeWhoName instantiates a new AeWhoName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeWhoName() *AeWhoName {
	this := AeWhoName{}
	var domain string = ""
	this.Domain = &domain
	var name string = ""
	this.Name = &name
	return &this
}

// NewAeWhoNameWithDefaults instantiates a new AeWhoName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeWhoNameWithDefaults() *AeWhoName {
	this := AeWhoName{}
	var domain string = ""
	this.Domain = &domain
	var name string = ""
	this.Name = &name
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *AeWhoName) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeWhoName) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *AeWhoName) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *AeWhoName) SetDomain(v string) {
	o.Domain = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AeWhoName) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeWhoName) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AeWhoName) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AeWhoName) SetName(v string) {
	o.Name = &v
}

func (o AeWhoName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableAeWhoName struct {
	value *AeWhoName
	isSet bool
}

func (v NullableAeWhoName) Get() *AeWhoName {
	return v.value
}

func (v *NullableAeWhoName) Set(val *AeWhoName) {
	v.value = val
	v.isSet = true
}

func (v NullableAeWhoName) IsSet() bool {
	return v.isSet
}

func (v *NullableAeWhoName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeWhoName(val *AeWhoName) *NullableAeWhoName {
	return &NullableAeWhoName{value: val, isSet: true}
}

func (v NullableAeWhoName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeWhoName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
