/*
TrueNAS RESTful API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
	"fmt"
)

// Managedby struct for Managedby
type Managedby struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Managedby) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(Managedby)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Managedby) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableManagedby struct {
	value *Managedby
	isSet bool
}

func (v NullableManagedby) Get() *Managedby {
	return v.value
}

func (v *NullableManagedby) Set(val *Managedby) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedby) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedby) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedby(val *Managedby) *NullableManagedby {
	return &NullableManagedby{value: val, isSet: true}
}

func (v NullableManagedby) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedby) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
