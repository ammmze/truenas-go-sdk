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

// Options1 struct for Options1
type Options1 struct {
	Stripacl     *bool `json:"stripacl,omitempty"`
	Recursive    *bool `json:"recursive,omitempty"`
	Traverse     *bool `json:"traverse,omitempty"`
	Canonicalize *bool `json:"canonicalize,omitempty"`
}

// NewOptions1 instantiates a new Options1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptions1() *Options1 {
	this := Options1{}
	var stripacl bool = false
	this.Stripacl = &stripacl
	var recursive bool = false
	this.Recursive = &recursive
	var traverse bool = false
	this.Traverse = &traverse
	var canonicalize bool = true
	this.Canonicalize = &canonicalize
	return &this
}

// NewOptions1WithDefaults instantiates a new Options1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptions1WithDefaults() *Options1 {
	this := Options1{}
	var stripacl bool = false
	this.Stripacl = &stripacl
	var recursive bool = false
	this.Recursive = &recursive
	var traverse bool = false
	this.Traverse = &traverse
	var canonicalize bool = true
	this.Canonicalize = &canonicalize
	return &this
}

// GetStripacl returns the Stripacl field value if set, zero value otherwise.
func (o *Options1) GetStripacl() bool {
	if o == nil || o.Stripacl == nil {
		var ret bool
		return ret
	}
	return *o.Stripacl
}

// GetStripaclOk returns a tuple with the Stripacl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options1) GetStripaclOk() (*bool, bool) {
	if o == nil || o.Stripacl == nil {
		return nil, false
	}
	return o.Stripacl, true
}

// HasStripacl returns a boolean if a field has been set.
func (o *Options1) HasStripacl() bool {
	if o != nil && o.Stripacl != nil {
		return true
	}

	return false
}

// SetStripacl gets a reference to the given bool and assigns it to the Stripacl field.
func (o *Options1) SetStripacl(v bool) {
	o.Stripacl = &v
}

// GetRecursive returns the Recursive field value if set, zero value otherwise.
func (o *Options1) GetRecursive() bool {
	if o == nil || o.Recursive == nil {
		var ret bool
		return ret
	}
	return *o.Recursive
}

// GetRecursiveOk returns a tuple with the Recursive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options1) GetRecursiveOk() (*bool, bool) {
	if o == nil || o.Recursive == nil {
		return nil, false
	}
	return o.Recursive, true
}

// HasRecursive returns a boolean if a field has been set.
func (o *Options1) HasRecursive() bool {
	if o != nil && o.Recursive != nil {
		return true
	}

	return false
}

// SetRecursive gets a reference to the given bool and assigns it to the Recursive field.
func (o *Options1) SetRecursive(v bool) {
	o.Recursive = &v
}

// GetTraverse returns the Traverse field value if set, zero value otherwise.
func (o *Options1) GetTraverse() bool {
	if o == nil || o.Traverse == nil {
		var ret bool
		return ret
	}
	return *o.Traverse
}

// GetTraverseOk returns a tuple with the Traverse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options1) GetTraverseOk() (*bool, bool) {
	if o == nil || o.Traverse == nil {
		return nil, false
	}
	return o.Traverse, true
}

// HasTraverse returns a boolean if a field has been set.
func (o *Options1) HasTraverse() bool {
	if o != nil && o.Traverse != nil {
		return true
	}

	return false
}

// SetTraverse gets a reference to the given bool and assigns it to the Traverse field.
func (o *Options1) SetTraverse(v bool) {
	o.Traverse = &v
}

// GetCanonicalize returns the Canonicalize field value if set, zero value otherwise.
func (o *Options1) GetCanonicalize() bool {
	if o == nil || o.Canonicalize == nil {
		var ret bool
		return ret
	}
	return *o.Canonicalize
}

// GetCanonicalizeOk returns a tuple with the Canonicalize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options1) GetCanonicalizeOk() (*bool, bool) {
	if o == nil || o.Canonicalize == nil {
		return nil, false
	}
	return o.Canonicalize, true
}

// HasCanonicalize returns a boolean if a field has been set.
func (o *Options1) HasCanonicalize() bool {
	if o != nil && o.Canonicalize != nil {
		return true
	}

	return false
}

// SetCanonicalize gets a reference to the given bool and assigns it to the Canonicalize field.
func (o *Options1) SetCanonicalize(v bool) {
	o.Canonicalize = &v
}

func (o Options1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Stripacl != nil {
		toSerialize["stripacl"] = o.Stripacl
	}
	if o.Recursive != nil {
		toSerialize["recursive"] = o.Recursive
	}
	if o.Traverse != nil {
		toSerialize["traverse"] = o.Traverse
	}
	if o.Canonicalize != nil {
		toSerialize["canonicalize"] = o.Canonicalize
	}
	return json.Marshal(toSerialize)
}

type NullableOptions1 struct {
	value *Options1
	isSet bool
}

func (v NullableOptions1) Get() *Options1 {
	return v.value
}

func (v *NullableOptions1) Set(val *Options1) {
	v.value = val
	v.isSet = true
}

func (v NullableOptions1) IsSet() bool {
	return v.isSet
}

func (v *NullableOptions1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptions1(val *Options1) *NullableOptions1 {
	return &NullableOptions1{value: val, isSet: true}
}

func (v NullableOptions1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptions1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
