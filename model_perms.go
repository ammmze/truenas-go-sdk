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

// Perms struct for Perms
type Perms struct {
	READ_DATA         *bool   `json:"READ_DATA,omitempty"`
	WRITE_DATA        *bool   `json:"WRITE_DATA,omitempty"`
	APPEND_DATA       *bool   `json:"APPEND_DATA,omitempty"`
	READ_NAMED_ATTRS  *bool   `json:"READ_NAMED_ATTRS,omitempty"`
	WRITE_NAMED_ATTRS *bool   `json:"WRITE_NAMED_ATTRS,omitempty"`
	EXECUTE           *bool   `json:"EXECUTE,omitempty"`
	DELETE_CHILD      *bool   `json:"DELETE_CHILD,omitempty"`
	READ_ATTRIBUTES   *bool   `json:"READ_ATTRIBUTES,omitempty"`
	WRITE_ATTRIBUTES  *bool   `json:"WRITE_ATTRIBUTES,omitempty"`
	DELETE            *bool   `json:"DELETE,omitempty"`
	READ_ACL          *bool   `json:"READ_ACL,omitempty"`
	WRITE_ACL         *bool   `json:"WRITE_ACL,omitempty"`
	WRITE_OWNER       *bool   `json:"WRITE_OWNER,omitempty"`
	SYNCHRONIZE       *bool   `json:"SYNCHRONIZE,omitempty"`
	BASIC             *string `json:"BASIC,omitempty"`
}

// NewPerms instantiates a new Perms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerms() *Perms {
	this := Perms{}
	return &this
}

// NewPermsWithDefaults instantiates a new Perms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermsWithDefaults() *Perms {
	this := Perms{}
	return &this
}

// GetREAD_DATA returns the READ_DATA field value if set, zero value otherwise.
func (o *Perms) GetREAD_DATA() bool {
	if o == nil || o.READ_DATA == nil {
		var ret bool
		return ret
	}
	return *o.READ_DATA
}

// GetREAD_DATAOk returns a tuple with the READ_DATA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perms) GetREAD_DATAOk() (*bool, bool) {
	if o == nil || o.READ_DATA == nil {
		return nil, false
	}
	return o.READ_DATA, true
}

// HasREAD_DATA returns a boolean if a field has been set.
func (o *Perms) HasREAD_DATA() bool {
	if o != nil && o.READ_DATA != nil {
		return true
	}

	return false
}

// SetREAD_DATA gets a reference to the given bool and assigns it to the READ_DATA field.
func (o *Perms) SetREAD_DATA(v bool) {
	o.READ_DATA = &v
}

// GetWRITE_DATA returns the WRITE_DATA field value if set, zero value otherwise.
func (o *Perms) GetWRITE_DATA() bool {
	if o == nil || o.WRITE_DATA == nil {
		var ret bool
		return ret
	}
	return *o.WRITE_DATA
}

// GetWRITE_DATAOk returns a tuple with the WRITE_DATA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perms) GetWRITE_DATAOk() (*bool, bool) {
	if o == nil || o.WRITE_DATA == nil {
		return nil, false
	}
	return o.WRITE_DATA, true
}

// HasWRITE_DATA returns a boolean if a field has been set.
func (o *Perms) HasWRITE_DATA() bool {
	if o != nil && o.WRITE_DATA != nil {
		return true
	}

	return false
}

// SetWRITE_DATA gets a reference to the given bool and assigns it to the WRITE_DATA field.
func (o *Perms) SetWRITE_DATA(v bool) {
	o.WRITE_DATA = &v
}

// GetAPPEND_DATA returns the APPEND_DATA field value if set, zero value otherwise.
func (o *Perms) GetAPPEND_DATA() bool {
	if o == nil || o.APPEND_DATA == nil {
		var ret bool
		return ret
	}
	return *o.APPEND_DATA
}

// GetAPPEND_DATAOk returns a tuple with the APPEND_DATA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perms) GetAPPEND_DATAOk() (*bool, bool) {
	if o == nil || o.APPEND_DATA == nil {
		return nil, false
	}
	return o.APPEND_DATA, true
}

// HasAPPEND_DATA returns a boolean if a field has been set.
func (o *Perms) HasAPPEND_DATA() bool {
	if o != nil && o.APPEND_DATA != nil {
		return true
	}

	return false
}

// SetAPPEND_DATA gets a reference to the given bool and assigns it to the APPEND_DATA field.
func (o *Perms) SetAPPEND_DATA(v bool) {
	o.APPEND_DATA = &v
}

// GetREAD_NAMED_ATTRS returns the READ_NAMED_ATTRS field value if set, zero value otherwise.
func (o *Perms) GetREAD_NAMED_ATTRS() bool {
	if o == nil || o.READ_NAMED_ATTRS == nil {
		var ret bool
		return ret
	}
	return *o.READ_NAMED_ATTRS
}

// GetREAD_NAMED_ATTRSOk returns a tuple with the READ_NAMED_ATTRS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perms) GetREAD_NAMED_ATTRSOk() (*bool, bool) {
	if o == nil || o.READ_NAMED_ATTRS == nil {
		return nil, false
	}
	return o.READ_NAMED_ATTRS, true
}

// HasREAD_NAMED_ATTRS returns a boolean if a field has been set.
func (o *Perms) HasREAD_NAMED_ATTRS() bool {
	if o != nil && o.READ_NAMED_ATTRS != nil {
		return true
	}

	return false
}

// SetREAD_NAMED_ATTRS gets a reference to the given bool and assigns it to the READ_NAMED_ATTRS field.
func (o *Perms) SetREAD_NAMED_ATTRS(v bool) {
	o.READ_NAMED_ATTRS = &v
}

// GetWRITE_NAMED_ATTRS returns the WRITE_NAMED_ATTRS field value if set, zero value otherwise.
func (o *Perms) GetWRITE_NAMED_ATTRS() bool {
	if o == nil || o.WRITE_NAMED_ATTRS == nil {
		var ret bool
		return ret
	}
	return *o.WRITE_NAMED_ATTRS
}

// GetWRITE_NAMED_ATTRSOk returns a tuple with the WRITE_NAMED_ATTRS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perms) GetWRITE_NAMED_ATTRSOk() (*bool, bool) {
	if o == nil || o.WRITE_NAMED_ATTRS == nil {
		return nil, false
	}
	return o.WRITE_NAMED_ATTRS, true
}

// HasWRITE_NAMED_ATTRS returns a boolean if a field has been set.
func (o *Perms) HasWRITE_NAMED_ATTRS() bool {
	if o != nil && o.WRITE_NAMED_ATTRS != nil {
		return true
	}

	return false
}

// SetWRITE_NAMED_ATTRS gets a reference to the given bool and assigns it to the WRITE_NAMED_ATTRS field.
func (o *Perms) SetWRITE_NAMED_ATTRS(v bool) {
	o.WRITE_NAMED_ATTRS = &v
}

// GetEXECUTE returns the EXECUTE field value if set, zero value otherwise.
func (o *Perms) GetEXECUTE() bool {
	if o == nil || o.EXECUTE == nil {
		var ret bool
		return ret
	}
	return *o.EXECUTE
}

// GetEXECUTEOk returns a tuple with the EXECUTE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perms) GetEXECUTEOk() (*bool, bool) {
	if o == nil || o.EXECUTE == nil {
		return nil, false
	}
	return o.EXECUTE, true
}

// HasEXECUTE returns a boolean if a field has been set.
func (o *Perms) HasEXECUTE() bool {
	if o != nil && o.EXECUTE != nil {
		return true
	}

	return false
}

// SetEXECUTE gets a reference to the given bool and assigns it to the EXECUTE field.
func (o *Perms) SetEXECUTE(v bool) {
	o.EXECUTE = &v
}

// GetDELETE_CHILD returns the DELETE_CHILD field value if set, zero value otherwise.
func (o *Perms) GetDELETE_CHILD() bool {
	if o == nil || o.DELETE_CHILD == nil {
		var ret bool
		return ret
	}
	return *o.DELETE_CHILD
}

// GetDELETE_CHILDOk returns a tuple with the DELETE_CHILD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perms) GetDELETE_CHILDOk() (*bool, bool) {
	if o == nil || o.DELETE_CHILD == nil {
		return nil, false
	}
	return o.DELETE_CHILD, true
}

// HasDELETE_CHILD returns a boolean if a field has been set.
func (o *Perms) HasDELETE_CHILD() bool {
	if o != nil && o.DELETE_CHILD != nil {
		return true
	}

	return false
}

// SetDELETE_CHILD gets a reference to the given bool and assigns it to the DELETE_CHILD field.
func (o *Perms) SetDELETE_CHILD(v bool) {
	o.DELETE_CHILD = &v
}

// GetREAD_ATTRIBUTES returns the READ_ATTRIBUTES field value if set, zero value otherwise.
func (o *Perms) GetREAD_ATTRIBUTES() bool {
	if o == nil || o.READ_ATTRIBUTES == nil {
		var ret bool
		return ret
	}
	return *o.READ_ATTRIBUTES
}

// GetREAD_ATTRIBUTESOk returns a tuple with the READ_ATTRIBUTES field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perms) GetREAD_ATTRIBUTESOk() (*bool, bool) {
	if o == nil || o.READ_ATTRIBUTES == nil {
		return nil, false
	}
	return o.READ_ATTRIBUTES, true
}

// HasREAD_ATTRIBUTES returns a boolean if a field has been set.
func (o *Perms) HasREAD_ATTRIBUTES() bool {
	if o != nil && o.READ_ATTRIBUTES != nil {
		return true
	}

	return false
}

// SetREAD_ATTRIBUTES gets a reference to the given bool and assigns it to the READ_ATTRIBUTES field.
func (o *Perms) SetREAD_ATTRIBUTES(v bool) {
	o.READ_ATTRIBUTES = &v
}

// GetWRITE_ATTRIBUTES returns the WRITE_ATTRIBUTES field value if set, zero value otherwise.
func (o *Perms) GetWRITE_ATTRIBUTES() bool {
	if o == nil || o.WRITE_ATTRIBUTES == nil {
		var ret bool
		return ret
	}
	return *o.WRITE_ATTRIBUTES
}

// GetWRITE_ATTRIBUTESOk returns a tuple with the WRITE_ATTRIBUTES field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perms) GetWRITE_ATTRIBUTESOk() (*bool, bool) {
	if o == nil || o.WRITE_ATTRIBUTES == nil {
		return nil, false
	}
	return o.WRITE_ATTRIBUTES, true
}

// HasWRITE_ATTRIBUTES returns a boolean if a field has been set.
func (o *Perms) HasWRITE_ATTRIBUTES() bool {
	if o != nil && o.WRITE_ATTRIBUTES != nil {
		return true
	}

	return false
}

// SetWRITE_ATTRIBUTES gets a reference to the given bool and assigns it to the WRITE_ATTRIBUTES field.
func (o *Perms) SetWRITE_ATTRIBUTES(v bool) {
	o.WRITE_ATTRIBUTES = &v
}

// GetDELETE returns the DELETE field value if set, zero value otherwise.
func (o *Perms) GetDELETE() bool {
	if o == nil || o.DELETE == nil {
		var ret bool
		return ret
	}
	return *o.DELETE
}

// GetDELETEOk returns a tuple with the DELETE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perms) GetDELETEOk() (*bool, bool) {
	if o == nil || o.DELETE == nil {
		return nil, false
	}
	return o.DELETE, true
}

// HasDELETE returns a boolean if a field has been set.
func (o *Perms) HasDELETE() bool {
	if o != nil && o.DELETE != nil {
		return true
	}

	return false
}

// SetDELETE gets a reference to the given bool and assigns it to the DELETE field.
func (o *Perms) SetDELETE(v bool) {
	o.DELETE = &v
}

// GetREAD_ACL returns the READ_ACL field value if set, zero value otherwise.
func (o *Perms) GetREAD_ACL() bool {
	if o == nil || o.READ_ACL == nil {
		var ret bool
		return ret
	}
	return *o.READ_ACL
}

// GetREAD_ACLOk returns a tuple with the READ_ACL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perms) GetREAD_ACLOk() (*bool, bool) {
	if o == nil || o.READ_ACL == nil {
		return nil, false
	}
	return o.READ_ACL, true
}

// HasREAD_ACL returns a boolean if a field has been set.
func (o *Perms) HasREAD_ACL() bool {
	if o != nil && o.READ_ACL != nil {
		return true
	}

	return false
}

// SetREAD_ACL gets a reference to the given bool and assigns it to the READ_ACL field.
func (o *Perms) SetREAD_ACL(v bool) {
	o.READ_ACL = &v
}

// GetWRITE_ACL returns the WRITE_ACL field value if set, zero value otherwise.
func (o *Perms) GetWRITE_ACL() bool {
	if o == nil || o.WRITE_ACL == nil {
		var ret bool
		return ret
	}
	return *o.WRITE_ACL
}

// GetWRITE_ACLOk returns a tuple with the WRITE_ACL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perms) GetWRITE_ACLOk() (*bool, bool) {
	if o == nil || o.WRITE_ACL == nil {
		return nil, false
	}
	return o.WRITE_ACL, true
}

// HasWRITE_ACL returns a boolean if a field has been set.
func (o *Perms) HasWRITE_ACL() bool {
	if o != nil && o.WRITE_ACL != nil {
		return true
	}

	return false
}

// SetWRITE_ACL gets a reference to the given bool and assigns it to the WRITE_ACL field.
func (o *Perms) SetWRITE_ACL(v bool) {
	o.WRITE_ACL = &v
}

// GetWRITE_OWNER returns the WRITE_OWNER field value if set, zero value otherwise.
func (o *Perms) GetWRITE_OWNER() bool {
	if o == nil || o.WRITE_OWNER == nil {
		var ret bool
		return ret
	}
	return *o.WRITE_OWNER
}

// GetWRITE_OWNEROk returns a tuple with the WRITE_OWNER field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perms) GetWRITE_OWNEROk() (*bool, bool) {
	if o == nil || o.WRITE_OWNER == nil {
		return nil, false
	}
	return o.WRITE_OWNER, true
}

// HasWRITE_OWNER returns a boolean if a field has been set.
func (o *Perms) HasWRITE_OWNER() bool {
	if o != nil && o.WRITE_OWNER != nil {
		return true
	}

	return false
}

// SetWRITE_OWNER gets a reference to the given bool and assigns it to the WRITE_OWNER field.
func (o *Perms) SetWRITE_OWNER(v bool) {
	o.WRITE_OWNER = &v
}

// GetSYNCHRONIZE returns the SYNCHRONIZE field value if set, zero value otherwise.
func (o *Perms) GetSYNCHRONIZE() bool {
	if o == nil || o.SYNCHRONIZE == nil {
		var ret bool
		return ret
	}
	return *o.SYNCHRONIZE
}

// GetSYNCHRONIZEOk returns a tuple with the SYNCHRONIZE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perms) GetSYNCHRONIZEOk() (*bool, bool) {
	if o == nil || o.SYNCHRONIZE == nil {
		return nil, false
	}
	return o.SYNCHRONIZE, true
}

// HasSYNCHRONIZE returns a boolean if a field has been set.
func (o *Perms) HasSYNCHRONIZE() bool {
	if o != nil && o.SYNCHRONIZE != nil {
		return true
	}

	return false
}

// SetSYNCHRONIZE gets a reference to the given bool and assigns it to the SYNCHRONIZE field.
func (o *Perms) SetSYNCHRONIZE(v bool) {
	o.SYNCHRONIZE = &v
}

// GetBASIC returns the BASIC field value if set, zero value otherwise.
func (o *Perms) GetBASIC() string {
	if o == nil || o.BASIC == nil {
		var ret string
		return ret
	}
	return *o.BASIC
}

// GetBASICOk returns a tuple with the BASIC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perms) GetBASICOk() (*string, bool) {
	if o == nil || o.BASIC == nil {
		return nil, false
	}
	return o.BASIC, true
}

// HasBASIC returns a boolean if a field has been set.
func (o *Perms) HasBASIC() bool {
	if o != nil && o.BASIC != nil {
		return true
	}

	return false
}

// SetBASIC gets a reference to the given string and assigns it to the BASIC field.
func (o *Perms) SetBASIC(v string) {
	o.BASIC = &v
}

func (o Perms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.READ_DATA != nil {
		toSerialize["READ_DATA"] = o.READ_DATA
	}
	if o.WRITE_DATA != nil {
		toSerialize["WRITE_DATA"] = o.WRITE_DATA
	}
	if o.APPEND_DATA != nil {
		toSerialize["APPEND_DATA"] = o.APPEND_DATA
	}
	if o.READ_NAMED_ATTRS != nil {
		toSerialize["READ_NAMED_ATTRS"] = o.READ_NAMED_ATTRS
	}
	if o.WRITE_NAMED_ATTRS != nil {
		toSerialize["WRITE_NAMED_ATTRS"] = o.WRITE_NAMED_ATTRS
	}
	if o.EXECUTE != nil {
		toSerialize["EXECUTE"] = o.EXECUTE
	}
	if o.DELETE_CHILD != nil {
		toSerialize["DELETE_CHILD"] = o.DELETE_CHILD
	}
	if o.READ_ATTRIBUTES != nil {
		toSerialize["READ_ATTRIBUTES"] = o.READ_ATTRIBUTES
	}
	if o.WRITE_ATTRIBUTES != nil {
		toSerialize["WRITE_ATTRIBUTES"] = o.WRITE_ATTRIBUTES
	}
	if o.DELETE != nil {
		toSerialize["DELETE"] = o.DELETE
	}
	if o.READ_ACL != nil {
		toSerialize["READ_ACL"] = o.READ_ACL
	}
	if o.WRITE_ACL != nil {
		toSerialize["WRITE_ACL"] = o.WRITE_ACL
	}
	if o.WRITE_OWNER != nil {
		toSerialize["WRITE_OWNER"] = o.WRITE_OWNER
	}
	if o.SYNCHRONIZE != nil {
		toSerialize["SYNCHRONIZE"] = o.SYNCHRONIZE
	}
	if o.BASIC != nil {
		toSerialize["BASIC"] = o.BASIC
	}
	return json.Marshal(toSerialize)
}

type NullablePerms struct {
	value *Perms
	isSet bool
}

func (v NullablePerms) Get() *Perms {
	return v.value
}

func (v *NullablePerms) Set(val *Perms) {
	v.value = val
	v.isSet = true
}

func (v NullablePerms) IsSet() bool {
	return v.isSet
}

func (v *NullablePerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerms(val *Perms) *NullablePerms {
	return &NullablePerms{value: val, isSet: true}
}

func (v NullablePerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
