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

// CloudsyncOnedriveListDrives0 struct for CloudsyncOnedriveListDrives0
type CloudsyncOnedriveListDrives0 struct {
	ClientId     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`
	Token        *string `json:"token,omitempty"`
}

// NewCloudsyncOnedriveListDrives0 instantiates a new CloudsyncOnedriveListDrives0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudsyncOnedriveListDrives0() *CloudsyncOnedriveListDrives0 {
	this := CloudsyncOnedriveListDrives0{}
	var clientId string = ""
	this.ClientId = &clientId
	var clientSecret string = ""
	this.ClientSecret = &clientSecret
	return &this
}

// NewCloudsyncOnedriveListDrives0WithDefaults instantiates a new CloudsyncOnedriveListDrives0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudsyncOnedriveListDrives0WithDefaults() *CloudsyncOnedriveListDrives0 {
	this := CloudsyncOnedriveListDrives0{}
	var clientId string = ""
	this.ClientId = &clientId
	var clientSecret string = ""
	this.ClientSecret = &clientSecret
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CloudsyncOnedriveListDrives0) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncOnedriveListDrives0) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CloudsyncOnedriveListDrives0) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CloudsyncOnedriveListDrives0) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *CloudsyncOnedriveListDrives0) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncOnedriveListDrives0) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *CloudsyncOnedriveListDrives0) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *CloudsyncOnedriveListDrives0) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CloudsyncOnedriveListDrives0) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncOnedriveListDrives0) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CloudsyncOnedriveListDrives0) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CloudsyncOnedriveListDrives0) SetToken(v string) {
	o.Token = &v
}

func (o CloudsyncOnedriveListDrives0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableCloudsyncOnedriveListDrives0 struct {
	value *CloudsyncOnedriveListDrives0
	isSet bool
}

func (v NullableCloudsyncOnedriveListDrives0) Get() *CloudsyncOnedriveListDrives0 {
	return v.value
}

func (v *NullableCloudsyncOnedriveListDrives0) Set(val *CloudsyncOnedriveListDrives0) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudsyncOnedriveListDrives0) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudsyncOnedriveListDrives0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudsyncOnedriveListDrives0(val *CloudsyncOnedriveListDrives0) *NullableCloudsyncOnedriveListDrives0 {
	return &NullableCloudsyncOnedriveListDrives0{value: val, isSet: true}
}

func (v NullableCloudsyncOnedriveListDrives0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudsyncOnedriveListDrives0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
