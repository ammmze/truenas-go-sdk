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

// FilesystemAcltemplateCreate0 struct for FilesystemAcltemplateCreate0
type FilesystemAcltemplateCreate0 struct {
	Name    *string                          `json:"name,omitempty"`
	Acltype *string                          `json:"acltype,omitempty"`
	Acl     *FilesystemAcltemplateCreate0Acl `json:"acl,omitempty"`
}

// NewFilesystemAcltemplateCreate0 instantiates a new FilesystemAcltemplateCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesystemAcltemplateCreate0() *FilesystemAcltemplateCreate0 {
	this := FilesystemAcltemplateCreate0{}
	return &this
}

// NewFilesystemAcltemplateCreate0WithDefaults instantiates a new FilesystemAcltemplateCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesystemAcltemplateCreate0WithDefaults() *FilesystemAcltemplateCreate0 {
	this := FilesystemAcltemplateCreate0{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FilesystemAcltemplateCreate0) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemAcltemplateCreate0) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FilesystemAcltemplateCreate0) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FilesystemAcltemplateCreate0) SetName(v string) {
	o.Name = &v
}

// GetAcltype returns the Acltype field value if set, zero value otherwise.
func (o *FilesystemAcltemplateCreate0) GetAcltype() string {
	if o == nil || o.Acltype == nil {
		var ret string
		return ret
	}
	return *o.Acltype
}

// GetAcltypeOk returns a tuple with the Acltype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemAcltemplateCreate0) GetAcltypeOk() (*string, bool) {
	if o == nil || o.Acltype == nil {
		return nil, false
	}
	return o.Acltype, true
}

// HasAcltype returns a boolean if a field has been set.
func (o *FilesystemAcltemplateCreate0) HasAcltype() bool {
	if o != nil && o.Acltype != nil {
		return true
	}

	return false
}

// SetAcltype gets a reference to the given string and assigns it to the Acltype field.
func (o *FilesystemAcltemplateCreate0) SetAcltype(v string) {
	o.Acltype = &v
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *FilesystemAcltemplateCreate0) GetAcl() FilesystemAcltemplateCreate0Acl {
	if o == nil || o.Acl == nil {
		var ret FilesystemAcltemplateCreate0Acl
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemAcltemplateCreate0) GetAclOk() (*FilesystemAcltemplateCreate0Acl, bool) {
	if o == nil || o.Acl == nil {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *FilesystemAcltemplateCreate0) HasAcl() bool {
	if o != nil && o.Acl != nil {
		return true
	}

	return false
}

// SetAcl gets a reference to the given FilesystemAcltemplateCreate0Acl and assigns it to the Acl field.
func (o *FilesystemAcltemplateCreate0) SetAcl(v FilesystemAcltemplateCreate0Acl) {
	o.Acl = &v
}

func (o FilesystemAcltemplateCreate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Acltype != nil {
		toSerialize["acltype"] = o.Acltype
	}
	if o.Acl != nil {
		toSerialize["acl"] = o.Acl
	}
	return json.Marshal(toSerialize)
}

type NullableFilesystemAcltemplateCreate0 struct {
	value *FilesystemAcltemplateCreate0
	isSet bool
}

func (v NullableFilesystemAcltemplateCreate0) Get() *FilesystemAcltemplateCreate0 {
	return v.value
}

func (v *NullableFilesystemAcltemplateCreate0) Set(val *FilesystemAcltemplateCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesystemAcltemplateCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesystemAcltemplateCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesystemAcltemplateCreate0(val *FilesystemAcltemplateCreate0) *NullableFilesystemAcltemplateCreate0 {
	return &NullableFilesystemAcltemplateCreate0{value: val, isSet: true}
}

func (v NullableFilesystemAcltemplateCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesystemAcltemplateCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
