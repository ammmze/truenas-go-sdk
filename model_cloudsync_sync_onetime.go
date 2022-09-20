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

// CloudsyncSyncOnetime struct for CloudsyncSyncOnetime
type CloudsyncSyncOnetime struct {
	CloudSyncCreate      *CloudsyncSyncOnetime0 `json:"cloud_sync_create,omitempty"`
	CloudSyncSyncOptions *CloudsyncSyncOnetime1 `json:"cloud_sync_sync_options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudsyncSyncOnetime CloudsyncSyncOnetime

// NewCloudsyncSyncOnetime instantiates a new CloudsyncSyncOnetime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudsyncSyncOnetime() *CloudsyncSyncOnetime {
	this := CloudsyncSyncOnetime{}
	return &this
}

// NewCloudsyncSyncOnetimeWithDefaults instantiates a new CloudsyncSyncOnetime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudsyncSyncOnetimeWithDefaults() *CloudsyncSyncOnetime {
	this := CloudsyncSyncOnetime{}
	return &this
}

// GetCloudSyncCreate returns the CloudSyncCreate field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime) GetCloudSyncCreate() CloudsyncSyncOnetime0 {
	if o == nil || o.CloudSyncCreate == nil {
		var ret CloudsyncSyncOnetime0
		return ret
	}
	return *o.CloudSyncCreate
}

// GetCloudSyncCreateOk returns a tuple with the CloudSyncCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime) GetCloudSyncCreateOk() (*CloudsyncSyncOnetime0, bool) {
	if o == nil || o.CloudSyncCreate == nil {
		return nil, false
	}
	return o.CloudSyncCreate, true
}

// HasCloudSyncCreate returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime) HasCloudSyncCreate() bool {
	if o != nil && o.CloudSyncCreate != nil {
		return true
	}

	return false
}

// SetCloudSyncCreate gets a reference to the given CloudsyncSyncOnetime0 and assigns it to the CloudSyncCreate field.
func (o *CloudsyncSyncOnetime) SetCloudSyncCreate(v CloudsyncSyncOnetime0) {
	o.CloudSyncCreate = &v
}

// GetCloudSyncSyncOptions returns the CloudSyncSyncOptions field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime) GetCloudSyncSyncOptions() CloudsyncSyncOnetime1 {
	if o == nil || o.CloudSyncSyncOptions == nil {
		var ret CloudsyncSyncOnetime1
		return ret
	}
	return *o.CloudSyncSyncOptions
}

// GetCloudSyncSyncOptionsOk returns a tuple with the CloudSyncSyncOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime) GetCloudSyncSyncOptionsOk() (*CloudsyncSyncOnetime1, bool) {
	if o == nil || o.CloudSyncSyncOptions == nil {
		return nil, false
	}
	return o.CloudSyncSyncOptions, true
}

// HasCloudSyncSyncOptions returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime) HasCloudSyncSyncOptions() bool {
	if o != nil && o.CloudSyncSyncOptions != nil {
		return true
	}

	return false
}

// SetCloudSyncSyncOptions gets a reference to the given CloudsyncSyncOnetime1 and assigns it to the CloudSyncSyncOptions field.
func (o *CloudsyncSyncOnetime) SetCloudSyncSyncOptions(v CloudsyncSyncOnetime1) {
	o.CloudSyncSyncOptions = &v
}

func (o CloudsyncSyncOnetime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudSyncCreate != nil {
		toSerialize["cloud_sync_create"] = o.CloudSyncCreate
	}
	if o.CloudSyncSyncOptions != nil {
		toSerialize["cloud_sync_sync_options"] = o.CloudSyncSyncOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudsyncSyncOnetime) UnmarshalJSON(bytes []byte) (err error) {
	varCloudsyncSyncOnetime := _CloudsyncSyncOnetime{}

	if err = json.Unmarshal(bytes, &varCloudsyncSyncOnetime); err == nil {
		*o = CloudsyncSyncOnetime(varCloudsyncSyncOnetime)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cloud_sync_create")
		delete(additionalProperties, "cloud_sync_sync_options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudsyncSyncOnetime struct {
	value *CloudsyncSyncOnetime
	isSet bool
}

func (v NullableCloudsyncSyncOnetime) Get() *CloudsyncSyncOnetime {
	return v.value
}

func (v *NullableCloudsyncSyncOnetime) Set(val *CloudsyncSyncOnetime) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudsyncSyncOnetime) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudsyncSyncOnetime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudsyncSyncOnetime(val *CloudsyncSyncOnetime) *NullableCloudsyncSyncOnetime {
	return &NullableCloudsyncSyncOnetime{value: val, isSet: true}
}

func (v NullableCloudsyncSyncOnetime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudsyncSyncOnetime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
