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

// FilesystemGetDefaultAcl1 the model 'FilesystemGetDefaultAcl1'
type FilesystemGetDefaultAcl1 string

// List of filesystem_get_default_acl_1
const (
	FILESYSTEMGETDEFAULTACL1_NONE FilesystemGetDefaultAcl1 = "NONE"
	FILESYSTEMGETDEFAULTACL1_SMB  FilesystemGetDefaultAcl1 = "SMB"
	FILESYSTEMGETDEFAULTACL1_NFS  FilesystemGetDefaultAcl1 = "NFS"
)

// All allowed values of FilesystemGetDefaultAcl1 enum
var AllowedFilesystemGetDefaultAcl1EnumValues = []FilesystemGetDefaultAcl1{
	"NONE",
	"SMB",
	"NFS",
}

func (v *FilesystemGetDefaultAcl1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FilesystemGetDefaultAcl1(value)
	for _, existing := range AllowedFilesystemGetDefaultAcl1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FilesystemGetDefaultAcl1", value)
}

// NewFilesystemGetDefaultAcl1FromValue returns a pointer to a valid FilesystemGetDefaultAcl1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFilesystemGetDefaultAcl1FromValue(v string) (*FilesystemGetDefaultAcl1, error) {
	ev := FilesystemGetDefaultAcl1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FilesystemGetDefaultAcl1: valid values are %v", v, AllowedFilesystemGetDefaultAcl1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FilesystemGetDefaultAcl1) IsValid() bool {
	for _, existing := range AllowedFilesystemGetDefaultAcl1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to filesystem_get_default_acl_1 value
func (v FilesystemGetDefaultAcl1) Ptr() *FilesystemGetDefaultAcl1 {
	return &v
}

type NullableFilesystemGetDefaultAcl1 struct {
	value *FilesystemGetDefaultAcl1
	isSet bool
}

func (v NullableFilesystemGetDefaultAcl1) Get() *FilesystemGetDefaultAcl1 {
	return v.value
}

func (v *NullableFilesystemGetDefaultAcl1) Set(val *FilesystemGetDefaultAcl1) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesystemGetDefaultAcl1) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesystemGetDefaultAcl1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesystemGetDefaultAcl1(val *FilesystemGetDefaultAcl1) *NullableFilesystemGetDefaultAcl1 {
	return &NullableFilesystemGetDefaultAcl1{value: val, isSet: true}
}

func (v NullableFilesystemGetDefaultAcl1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesystemGetDefaultAcl1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
