/*
TrueNAS RESTful API

Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package angelfish

import (
	"encoding/json"
)

// Service struct for Service
type Service struct {
	Id                   int32   `json:"id"`
	Service              string  `json:"service"`
	Enable               *bool   `json:"enable,omitempty"`
	State                *string `json:"state,omitempty"`
	Pids                 []int32 `json:"pids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Service Service

// NewService instantiates a new Service object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewService(id int32, service string) *Service {
	this := Service{}
	this.Id = id
	this.Service = service
	return &this
}

// NewServiceWithDefaults instantiates a new Service object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithDefaults() *Service {
	this := Service{}
	return &this
}

// GetId returns the Id field value
func (o *Service) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Service) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Service) SetId(v int32) {
	o.Id = v
}

// GetService returns the Service field value
func (o *Service) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *Service) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *Service) SetService(v string) {
	o.Service = v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *Service) GetEnable() bool {
	if o == nil || isNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetEnableOk() (*bool, bool) {
	if o == nil || isNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *Service) HasEnable() bool {
	if o != nil && !isNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *Service) SetEnable(v bool) {
	o.Enable = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Service) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Service) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Service) SetState(v string) {
	o.State = &v
}

// GetPids returns the Pids field value if set, zero value otherwise.
func (o *Service) GetPids() []int32 {
	if o == nil || isNil(o.Pids) {
		var ret []int32
		return ret
	}
	return o.Pids
}

// GetPidsOk returns a tuple with the Pids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetPidsOk() ([]int32, bool) {
	if o == nil || isNil(o.Pids) {
		return nil, false
	}
	return o.Pids, true
}

// HasPids returns a boolean if a field has been set.
func (o *Service) HasPids() bool {
	if o != nil && !isNil(o.Pids) {
		return true
	}

	return false
}

// SetPids gets a reference to the given []int32 and assigns it to the Pids field.
func (o *Service) SetPids(v []int32) {
	o.Pids = v
}

func (o Service) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["service"] = o.Service
	}
	if !isNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Pids) {
		toSerialize["pids"] = o.Pids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Service) UnmarshalJSON(bytes []byte) (err error) {
	varService := _Service{}

	if err = json.Unmarshal(bytes, &varService); err == nil {
		*o = Service(varService)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "service")
		delete(additionalProperties, "enable")
		delete(additionalProperties, "state")
		delete(additionalProperties, "pids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableService struct {
	value *Service
	isSet bool
}

func (v NullableService) Get() *Service {
	return v.value
}

func (v *NullableService) Set(val *Service) {
	v.value = val
	v.isSet = true
}

func (v NullableService) IsSet() bool {
	return v.isSet
}

func (v *NullableService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableService(val *Service) *NullableService {
	return &NullableService{value: val, isSet: true}
}

func (v NullableService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
