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

// Aclentry struct for Aclentry
type Aclentry struct {
	AeWhoSid  *string    `json:"ae_who_sid,omitempty"`
	AeWhoName *AeWhoName `json:"ae_who_name,omitempty"`
	AePerm    *string    `json:"ae_perm,omitempty"`
	AeType    *string    `json:"ae_type,omitempty"`
}

// NewAclentry instantiates a new Aclentry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAclentry() *Aclentry {
	this := Aclentry{}
	return &this
}

// NewAclentryWithDefaults instantiates a new Aclentry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAclentryWithDefaults() *Aclentry {
	this := Aclentry{}
	return &this
}

// GetAeWhoSid returns the AeWhoSid field value if set, zero value otherwise.
func (o *Aclentry) GetAeWhoSid() string {
	if o == nil || o.AeWhoSid == nil {
		var ret string
		return ret
	}
	return *o.AeWhoSid
}

// GetAeWhoSidOk returns a tuple with the AeWhoSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aclentry) GetAeWhoSidOk() (*string, bool) {
	if o == nil || o.AeWhoSid == nil {
		return nil, false
	}
	return o.AeWhoSid, true
}

// HasAeWhoSid returns a boolean if a field has been set.
func (o *Aclentry) HasAeWhoSid() bool {
	if o != nil && o.AeWhoSid != nil {
		return true
	}

	return false
}

// SetAeWhoSid gets a reference to the given string and assigns it to the AeWhoSid field.
func (o *Aclentry) SetAeWhoSid(v string) {
	o.AeWhoSid = &v
}

// GetAeWhoName returns the AeWhoName field value if set, zero value otherwise.
func (o *Aclentry) GetAeWhoName() AeWhoName {
	if o == nil || o.AeWhoName == nil {
		var ret AeWhoName
		return ret
	}
	return *o.AeWhoName
}

// GetAeWhoNameOk returns a tuple with the AeWhoName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aclentry) GetAeWhoNameOk() (*AeWhoName, bool) {
	if o == nil || o.AeWhoName == nil {
		return nil, false
	}
	return o.AeWhoName, true
}

// HasAeWhoName returns a boolean if a field has been set.
func (o *Aclentry) HasAeWhoName() bool {
	if o != nil && o.AeWhoName != nil {
		return true
	}

	return false
}

// SetAeWhoName gets a reference to the given AeWhoName and assigns it to the AeWhoName field.
func (o *Aclentry) SetAeWhoName(v AeWhoName) {
	o.AeWhoName = &v
}

// GetAePerm returns the AePerm field value if set, zero value otherwise.
func (o *Aclentry) GetAePerm() string {
	if o == nil || o.AePerm == nil {
		var ret string
		return ret
	}
	return *o.AePerm
}

// GetAePermOk returns a tuple with the AePerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aclentry) GetAePermOk() (*string, bool) {
	if o == nil || o.AePerm == nil {
		return nil, false
	}
	return o.AePerm, true
}

// HasAePerm returns a boolean if a field has been set.
func (o *Aclentry) HasAePerm() bool {
	if o != nil && o.AePerm != nil {
		return true
	}

	return false
}

// SetAePerm gets a reference to the given string and assigns it to the AePerm field.
func (o *Aclentry) SetAePerm(v string) {
	o.AePerm = &v
}

// GetAeType returns the AeType field value if set, zero value otherwise.
func (o *Aclentry) GetAeType() string {
	if o == nil || o.AeType == nil {
		var ret string
		return ret
	}
	return *o.AeType
}

// GetAeTypeOk returns a tuple with the AeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aclentry) GetAeTypeOk() (*string, bool) {
	if o == nil || o.AeType == nil {
		return nil, false
	}
	return o.AeType, true
}

// HasAeType returns a boolean if a field has been set.
func (o *Aclentry) HasAeType() bool {
	if o != nil && o.AeType != nil {
		return true
	}

	return false
}

// SetAeType gets a reference to the given string and assigns it to the AeType field.
func (o *Aclentry) SetAeType(v string) {
	o.AeType = &v
}

func (o Aclentry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AeWhoSid != nil {
		toSerialize["ae_who_sid"] = o.AeWhoSid
	}
	if o.AeWhoName != nil {
		toSerialize["ae_who_name"] = o.AeWhoName
	}
	if o.AePerm != nil {
		toSerialize["ae_perm"] = o.AePerm
	}
	if o.AeType != nil {
		toSerialize["ae_type"] = o.AeType
	}
	return json.Marshal(toSerialize)
}

type NullableAclentry struct {
	value *Aclentry
	isSet bool
}

func (v NullableAclentry) Get() *Aclentry {
	return v.value
}

func (v *NullableAclentry) Set(val *Aclentry) {
	v.value = val
	v.isSet = true
}

func (v NullableAclentry) IsSet() bool {
	return v.isSet
}

func (v *NullableAclentry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAclentry(val *Aclentry) *NullableAclentry {
	return &NullableAclentry{value: val, isSet: true}
}

func (v NullableAclentry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAclentry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
