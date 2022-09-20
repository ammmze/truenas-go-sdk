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

// SystemNtpserverCreate0 struct for SystemNtpserverCreate0
type SystemNtpserverCreate0 struct {
	// `address` specifies the hostname/IP address of the NTP server.
	Address *string `json:"address,omitempty"`
	// `burst` when enabled makes sure that if server is reachable, sends a burst of eight packets instead of one. This is designed to improve timekeeping quality with the server command.
	Burst *bool `json:"burst,omitempty"`
	// `iburst` when enabled speeds up the initial synchronization, taking seconds rather than minutes.
	Iburst *bool `json:"iburst,omitempty"`
	// `prefer` marks the specified server as preferred. When all other things are equal, this host is chosen for synchronization acquisition with the server command. It is recommended that they be used for servers with time monitoring hardware.
	Prefer *bool `json:"prefer,omitempty"`
	// `minpoll` is minimum polling time in seconds. It must be a power of 2 and less than `maxpoll`. `maxpoll` is maximum polling time in seconds. It must be a power of 2 and greater than `minpoll`.
	Minpoll *int32 `json:"minpoll,omitempty"`
	// `minpoll` is minimum polling time in seconds. It must be a power of 2 and less than `maxpoll`. `maxpoll` is maximum polling time in seconds. It must be a power of 2 and greater than `minpoll`.
	Maxpoll *int32 `json:"maxpoll,omitempty"`
	Force   *bool  `json:"force,omitempty"`
}

// NewSystemNtpserverCreate0 instantiates a new SystemNtpserverCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemNtpserverCreate0() *SystemNtpserverCreate0 {
	this := SystemNtpserverCreate0{}
	var burst bool = false
	this.Burst = &burst
	var iburst bool = true
	this.Iburst = &iburst
	var prefer bool = false
	this.Prefer = &prefer
	var minpoll int32 = 6
	this.Minpoll = &minpoll
	var maxpoll int32 = 10
	this.Maxpoll = &maxpoll
	return &this
}

// NewSystemNtpserverCreate0WithDefaults instantiates a new SystemNtpserverCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemNtpserverCreate0WithDefaults() *SystemNtpserverCreate0 {
	this := SystemNtpserverCreate0{}
	var burst bool = false
	this.Burst = &burst
	var iburst bool = true
	this.Iburst = &iburst
	var prefer bool = false
	this.Prefer = &prefer
	var minpoll int32 = 6
	this.Minpoll = &minpoll
	var maxpoll int32 = 10
	this.Maxpoll = &maxpoll
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *SystemNtpserverCreate0) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemNtpserverCreate0) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *SystemNtpserverCreate0) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *SystemNtpserverCreate0) SetAddress(v string) {
	o.Address = &v
}

// GetBurst returns the Burst field value if set, zero value otherwise.
func (o *SystemNtpserverCreate0) GetBurst() bool {
	if o == nil || o.Burst == nil {
		var ret bool
		return ret
	}
	return *o.Burst
}

// GetBurstOk returns a tuple with the Burst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemNtpserverCreate0) GetBurstOk() (*bool, bool) {
	if o == nil || o.Burst == nil {
		return nil, false
	}
	return o.Burst, true
}

// HasBurst returns a boolean if a field has been set.
func (o *SystemNtpserverCreate0) HasBurst() bool {
	if o != nil && o.Burst != nil {
		return true
	}

	return false
}

// SetBurst gets a reference to the given bool and assigns it to the Burst field.
func (o *SystemNtpserverCreate0) SetBurst(v bool) {
	o.Burst = &v
}

// GetIburst returns the Iburst field value if set, zero value otherwise.
func (o *SystemNtpserverCreate0) GetIburst() bool {
	if o == nil || o.Iburst == nil {
		var ret bool
		return ret
	}
	return *o.Iburst
}

// GetIburstOk returns a tuple with the Iburst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemNtpserverCreate0) GetIburstOk() (*bool, bool) {
	if o == nil || o.Iburst == nil {
		return nil, false
	}
	return o.Iburst, true
}

// HasIburst returns a boolean if a field has been set.
func (o *SystemNtpserverCreate0) HasIburst() bool {
	if o != nil && o.Iburst != nil {
		return true
	}

	return false
}

// SetIburst gets a reference to the given bool and assigns it to the Iburst field.
func (o *SystemNtpserverCreate0) SetIburst(v bool) {
	o.Iburst = &v
}

// GetPrefer returns the Prefer field value if set, zero value otherwise.
func (o *SystemNtpserverCreate0) GetPrefer() bool {
	if o == nil || o.Prefer == nil {
		var ret bool
		return ret
	}
	return *o.Prefer
}

// GetPreferOk returns a tuple with the Prefer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemNtpserverCreate0) GetPreferOk() (*bool, bool) {
	if o == nil || o.Prefer == nil {
		return nil, false
	}
	return o.Prefer, true
}

// HasPrefer returns a boolean if a field has been set.
func (o *SystemNtpserverCreate0) HasPrefer() bool {
	if o != nil && o.Prefer != nil {
		return true
	}

	return false
}

// SetPrefer gets a reference to the given bool and assigns it to the Prefer field.
func (o *SystemNtpserverCreate0) SetPrefer(v bool) {
	o.Prefer = &v
}

// GetMinpoll returns the Minpoll field value if set, zero value otherwise.
func (o *SystemNtpserverCreate0) GetMinpoll() int32 {
	if o == nil || o.Minpoll == nil {
		var ret int32
		return ret
	}
	return *o.Minpoll
}

// GetMinpollOk returns a tuple with the Minpoll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemNtpserverCreate0) GetMinpollOk() (*int32, bool) {
	if o == nil || o.Minpoll == nil {
		return nil, false
	}
	return o.Minpoll, true
}

// HasMinpoll returns a boolean if a field has been set.
func (o *SystemNtpserverCreate0) HasMinpoll() bool {
	if o != nil && o.Minpoll != nil {
		return true
	}

	return false
}

// SetMinpoll gets a reference to the given int32 and assigns it to the Minpoll field.
func (o *SystemNtpserverCreate0) SetMinpoll(v int32) {
	o.Minpoll = &v
}

// GetMaxpoll returns the Maxpoll field value if set, zero value otherwise.
func (o *SystemNtpserverCreate0) GetMaxpoll() int32 {
	if o == nil || o.Maxpoll == nil {
		var ret int32
		return ret
	}
	return *o.Maxpoll
}

// GetMaxpollOk returns a tuple with the Maxpoll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemNtpserverCreate0) GetMaxpollOk() (*int32, bool) {
	if o == nil || o.Maxpoll == nil {
		return nil, false
	}
	return o.Maxpoll, true
}

// HasMaxpoll returns a boolean if a field has been set.
func (o *SystemNtpserverCreate0) HasMaxpoll() bool {
	if o != nil && o.Maxpoll != nil {
		return true
	}

	return false
}

// SetMaxpoll gets a reference to the given int32 and assigns it to the Maxpoll field.
func (o *SystemNtpserverCreate0) SetMaxpoll(v int32) {
	o.Maxpoll = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *SystemNtpserverCreate0) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemNtpserverCreate0) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *SystemNtpserverCreate0) HasForce() bool {
	if o != nil && o.Force != nil {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *SystemNtpserverCreate0) SetForce(v bool) {
	o.Force = &v
}

func (o SystemNtpserverCreate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Burst != nil {
		toSerialize["burst"] = o.Burst
	}
	if o.Iburst != nil {
		toSerialize["iburst"] = o.Iburst
	}
	if o.Prefer != nil {
		toSerialize["prefer"] = o.Prefer
	}
	if o.Minpoll != nil {
		toSerialize["minpoll"] = o.Minpoll
	}
	if o.Maxpoll != nil {
		toSerialize["maxpoll"] = o.Maxpoll
	}
	if o.Force != nil {
		toSerialize["force"] = o.Force
	}
	return json.Marshal(toSerialize)
}

type NullableSystemNtpserverCreate0 struct {
	value *SystemNtpserverCreate0
	isSet bool
}

func (v NullableSystemNtpserverCreate0) Get() *SystemNtpserverCreate0 {
	return v.value
}

func (v *NullableSystemNtpserverCreate0) Set(val *SystemNtpserverCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemNtpserverCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemNtpserverCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemNtpserverCreate0(val *SystemNtpserverCreate0) *NullableSystemNtpserverCreate0 {
	return &NullableSystemNtpserverCreate0{value: val, isSet: true}
}

func (v NullableSystemNtpserverCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemNtpserverCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
