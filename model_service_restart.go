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

// ServiceRestart struct for ServiceRestart
type ServiceRestart struct {
	Service              *string          `json:"service,omitempty"`
	ServiceControl       *ServiceRestart1 `json:"service-control,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceRestart ServiceRestart

// NewServiceRestart instantiates a new ServiceRestart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceRestart() *ServiceRestart {
	this := ServiceRestart{}
	return &this
}

// NewServiceRestartWithDefaults instantiates a new ServiceRestart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceRestartWithDefaults() *ServiceRestart {
	this := ServiceRestart{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ServiceRestart) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRestart) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ServiceRestart) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ServiceRestart) SetService(v string) {
	o.Service = &v
}

// GetServiceControl returns the ServiceControl field value if set, zero value otherwise.
func (o *ServiceRestart) GetServiceControl() ServiceRestart1 {
	if o == nil || o.ServiceControl == nil {
		var ret ServiceRestart1
		return ret
	}
	return *o.ServiceControl
}

// GetServiceControlOk returns a tuple with the ServiceControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRestart) GetServiceControlOk() (*ServiceRestart1, bool) {
	if o == nil || o.ServiceControl == nil {
		return nil, false
	}
	return o.ServiceControl, true
}

// HasServiceControl returns a boolean if a field has been set.
func (o *ServiceRestart) HasServiceControl() bool {
	if o != nil && o.ServiceControl != nil {
		return true
	}

	return false
}

// SetServiceControl gets a reference to the given ServiceRestart1 and assigns it to the ServiceControl field.
func (o *ServiceRestart) SetServiceControl(v ServiceRestart1) {
	o.ServiceControl = &v
}

func (o ServiceRestart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.ServiceControl != nil {
		toSerialize["service-control"] = o.ServiceControl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceRestart) UnmarshalJSON(bytes []byte) (err error) {
	varServiceRestart := _ServiceRestart{}

	if err = json.Unmarshal(bytes, &varServiceRestart); err == nil {
		*o = ServiceRestart(varServiceRestart)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service")
		delete(additionalProperties, "service-control")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceRestart struct {
	value *ServiceRestart
	isSet bool
}

func (v NullableServiceRestart) Get() *ServiceRestart {
	return v.value
}

func (v *NullableServiceRestart) Set(val *ServiceRestart) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceRestart) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceRestart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceRestart(val *ServiceRestart) *NullableServiceRestart {
	return &NullableServiceRestart{value: val, isSet: true}
}

func (v NullableServiceRestart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceRestart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
