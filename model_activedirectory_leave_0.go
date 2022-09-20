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

// ActivedirectoryLeave0 struct for ActivedirectoryLeave0
type ActivedirectoryLeave0 struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

// NewActivedirectoryLeave0 instantiates a new ActivedirectoryLeave0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivedirectoryLeave0() *ActivedirectoryLeave0 {
	this := ActivedirectoryLeave0{}
	return &this
}

// NewActivedirectoryLeave0WithDefaults instantiates a new ActivedirectoryLeave0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivedirectoryLeave0WithDefaults() *ActivedirectoryLeave0 {
	this := ActivedirectoryLeave0{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ActivedirectoryLeave0) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivedirectoryLeave0) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ActivedirectoryLeave0) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ActivedirectoryLeave0) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ActivedirectoryLeave0) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivedirectoryLeave0) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ActivedirectoryLeave0) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ActivedirectoryLeave0) SetPassword(v string) {
	o.Password = &v
}

func (o ActivedirectoryLeave0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableActivedirectoryLeave0 struct {
	value *ActivedirectoryLeave0
	isSet bool
}

func (v NullableActivedirectoryLeave0) Get() *ActivedirectoryLeave0 {
	return v.value
}

func (v *NullableActivedirectoryLeave0) Set(val *ActivedirectoryLeave0) {
	v.value = val
	v.isSet = true
}

func (v NullableActivedirectoryLeave0) IsSet() bool {
	return v.isSet
}

func (v *NullableActivedirectoryLeave0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivedirectoryLeave0(val *ActivedirectoryLeave0) *NullableActivedirectoryLeave0 {
	return &NullableActivedirectoryLeave0{value: val, isSet: true}
}

func (v NullableActivedirectoryLeave0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivedirectoryLeave0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
