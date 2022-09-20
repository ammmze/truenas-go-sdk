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

// SnapshotsInner - struct for SnapshotsInner
type SnapshotsInner struct {
	SnapshotSpec *SnapshotSpec
	String       *string
}

// SnapshotSpecAsSnapshotsInner is a convenience function that returns SnapshotSpec wrapped in SnapshotsInner
func SnapshotSpecAsSnapshotsInner(v *SnapshotSpec) SnapshotsInner {
	return SnapshotsInner{
		SnapshotSpec: v,
	}
}

// stringAsSnapshotsInner is a convenience function that returns string wrapped in SnapshotsInner
func StringAsSnapshotsInner(v *string) SnapshotsInner {
	return SnapshotsInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SnapshotsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SnapshotSpec
	err = newStrictDecoder(data).Decode(&dst.SnapshotSpec)
	if err == nil {
		jsonSnapshotSpec, _ := json.Marshal(dst.SnapshotSpec)
		if string(jsonSnapshotSpec) == "{}" { // empty struct
			dst.SnapshotSpec = nil
		} else {
			match++
		}
	} else {
		dst.SnapshotSpec = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SnapshotSpec = nil
		dst.String = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(SnapshotsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SnapshotsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SnapshotsInner) MarshalJSON() ([]byte, error) {
	if src.SnapshotSpec != nil {
		return json.Marshal(&src.SnapshotSpec)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SnapshotsInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SnapshotSpec != nil {
		return obj.SnapshotSpec
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableSnapshotsInner struct {
	value *SnapshotsInner
	isSet bool
}

func (v NullableSnapshotsInner) Get() *SnapshotsInner {
	return v.value
}

func (v *NullableSnapshotsInner) Set(val *SnapshotsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotsInner(val *SnapshotsInner) *NullableSnapshotsInner {
	return &NullableSnapshotsInner{value: val, isSet: true}
}

func (v NullableSnapshotsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
