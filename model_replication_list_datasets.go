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

// ReplicationListDatasets struct for ReplicationListDatasets
type ReplicationListDatasets struct {
	Transport            *ReplicationListDatasets0 `json:"transport,omitempty"`
	SshCredentials       NullableInt32             `json:"ssh_credentials,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReplicationListDatasets ReplicationListDatasets

// NewReplicationListDatasets instantiates a new ReplicationListDatasets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicationListDatasets() *ReplicationListDatasets {
	this := ReplicationListDatasets{}
	return &this
}

// NewReplicationListDatasetsWithDefaults instantiates a new ReplicationListDatasets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicationListDatasetsWithDefaults() *ReplicationListDatasets {
	this := ReplicationListDatasets{}
	return &this
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *ReplicationListDatasets) GetTransport() ReplicationListDatasets0 {
	if o == nil || o.Transport == nil {
		var ret ReplicationListDatasets0
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationListDatasets) GetTransportOk() (*ReplicationListDatasets0, bool) {
	if o == nil || o.Transport == nil {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *ReplicationListDatasets) HasTransport() bool {
	if o != nil && o.Transport != nil {
		return true
	}

	return false
}

// SetTransport gets a reference to the given ReplicationListDatasets0 and assigns it to the Transport field.
func (o *ReplicationListDatasets) SetTransport(v ReplicationListDatasets0) {
	o.Transport = &v
}

// GetSshCredentials returns the SshCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplicationListDatasets) GetSshCredentials() int32 {
	if o == nil || o.SshCredentials.Get() == nil {
		var ret int32
		return ret
	}
	return *o.SshCredentials.Get()
}

// GetSshCredentialsOk returns a tuple with the SshCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicationListDatasets) GetSshCredentialsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SshCredentials.Get(), o.SshCredentials.IsSet()
}

// HasSshCredentials returns a boolean if a field has been set.
func (o *ReplicationListDatasets) HasSshCredentials() bool {
	if o != nil && o.SshCredentials.IsSet() {
		return true
	}

	return false
}

// SetSshCredentials gets a reference to the given NullableInt32 and assigns it to the SshCredentials field.
func (o *ReplicationListDatasets) SetSshCredentials(v int32) {
	o.SshCredentials.Set(&v)
}

// SetSshCredentialsNil sets the value for SshCredentials to be an explicit nil
func (o *ReplicationListDatasets) SetSshCredentialsNil() {
	o.SshCredentials.Set(nil)
}

// UnsetSshCredentials ensures that no value is present for SshCredentials, not even an explicit nil
func (o *ReplicationListDatasets) UnsetSshCredentials() {
	o.SshCredentials.Unset()
}

func (o ReplicationListDatasets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Transport != nil {
		toSerialize["transport"] = o.Transport
	}
	if o.SshCredentials.IsSet() {
		toSerialize["ssh_credentials"] = o.SshCredentials.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ReplicationListDatasets) UnmarshalJSON(bytes []byte) (err error) {
	varReplicationListDatasets := _ReplicationListDatasets{}

	if err = json.Unmarshal(bytes, &varReplicationListDatasets); err == nil {
		*o = ReplicationListDatasets(varReplicationListDatasets)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "transport")
		delete(additionalProperties, "ssh_credentials")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReplicationListDatasets struct {
	value *ReplicationListDatasets
	isSet bool
}

func (v NullableReplicationListDatasets) Get() *ReplicationListDatasets {
	return v.value
}

func (v *NullableReplicationListDatasets) Set(val *ReplicationListDatasets) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicationListDatasets) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicationListDatasets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicationListDatasets(val *ReplicationListDatasets) *NullableReplicationListDatasets {
	return &NullableReplicationListDatasets{value: val, isSet: true}
}

func (v NullableReplicationListDatasets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicationListDatasets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
