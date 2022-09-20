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

// FilesystemPut struct for FilesystemPut
type FilesystemPut struct {
	Path                 *string         `json:"path,omitempty"`
	Options              *FilesystemPut1 `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FilesystemPut FilesystemPut

// NewFilesystemPut instantiates a new FilesystemPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesystemPut() *FilesystemPut {
	this := FilesystemPut{}
	return &this
}

// NewFilesystemPutWithDefaults instantiates a new FilesystemPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesystemPutWithDefaults() *FilesystemPut {
	this := FilesystemPut{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *FilesystemPut) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemPut) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *FilesystemPut) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *FilesystemPut) SetPath(v string) {
	o.Path = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *FilesystemPut) GetOptions() FilesystemPut1 {
	if o == nil || o.Options == nil {
		var ret FilesystemPut1
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemPut) GetOptionsOk() (*FilesystemPut1, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *FilesystemPut) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given FilesystemPut1 and assigns it to the Options field.
func (o *FilesystemPut) SetOptions(v FilesystemPut1) {
	o.Options = &v
}

func (o FilesystemPut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FilesystemPut) UnmarshalJSON(bytes []byte) (err error) {
	varFilesystemPut := _FilesystemPut{}

	if err = json.Unmarshal(bytes, &varFilesystemPut); err == nil {
		*o = FilesystemPut(varFilesystemPut)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "path")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesystemPut struct {
	value *FilesystemPut
	isSet bool
}

func (v NullableFilesystemPut) Get() *FilesystemPut {
	return v.value
}

func (v *NullableFilesystemPut) Set(val *FilesystemPut) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesystemPut) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesystemPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesystemPut(val *FilesystemPut) *NullableFilesystemPut {
	return &NullableFilesystemPut{value: val, isSet: true}
}

func (v NullableFilesystemPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesystemPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
