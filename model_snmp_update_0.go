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

// SnmpUpdate0 struct for SnmpUpdate0
type SnmpUpdate0 struct {
	Location *string `json:"location,omitempty"`
	Contact  *string `json:"contact,omitempty"`
	Traps    *bool   `json:"traps,omitempty"`
	// `v3` when set enables SNMP version 3.
	V3               *bool          `json:"v3,omitempty"`
	Community        *string        `json:"community,omitempty"`
	V3Username       *string        `json:"v3_username,omitempty"`
	V3Authtype       *string        `json:"v3_authtype,omitempty"`
	V3Password       *string        `json:"v3_password,omitempty"`
	V3Privproto      NullableString `json:"v3_privproto,omitempty"`
	V3Privpassphrase NullableString `json:"v3_privpassphrase,omitempty"`
	Loglevel         *int32         `json:"loglevel,omitempty"`
	Options          *string        `json:"options,omitempty"`
	Zilstat          *bool          `json:"zilstat,omitempty"`
}

// NewSnmpUpdate0 instantiates a new SnmpUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpUpdate0() *SnmpUpdate0 {
	this := SnmpUpdate0{}
	var community string = "public"
	this.Community = &community
	return &this
}

// NewSnmpUpdate0WithDefaults instantiates a new SnmpUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpUpdate0WithDefaults() *SnmpUpdate0 {
	this := SnmpUpdate0{}
	var community string = "public"
	this.Community = &community
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SnmpUpdate0) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUpdate0) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SnmpUpdate0) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *SnmpUpdate0) SetLocation(v string) {
	o.Location = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *SnmpUpdate0) GetContact() string {
	if o == nil || o.Contact == nil {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUpdate0) GetContactOk() (*string, bool) {
	if o == nil || o.Contact == nil {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *SnmpUpdate0) HasContact() bool {
	if o != nil && o.Contact != nil {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *SnmpUpdate0) SetContact(v string) {
	o.Contact = &v
}

// GetTraps returns the Traps field value if set, zero value otherwise.
func (o *SnmpUpdate0) GetTraps() bool {
	if o == nil || o.Traps == nil {
		var ret bool
		return ret
	}
	return *o.Traps
}

// GetTrapsOk returns a tuple with the Traps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUpdate0) GetTrapsOk() (*bool, bool) {
	if o == nil || o.Traps == nil {
		return nil, false
	}
	return o.Traps, true
}

// HasTraps returns a boolean if a field has been set.
func (o *SnmpUpdate0) HasTraps() bool {
	if o != nil && o.Traps != nil {
		return true
	}

	return false
}

// SetTraps gets a reference to the given bool and assigns it to the Traps field.
func (o *SnmpUpdate0) SetTraps(v bool) {
	o.Traps = &v
}

// GetV3 returns the V3 field value if set, zero value otherwise.
func (o *SnmpUpdate0) GetV3() bool {
	if o == nil || o.V3 == nil {
		var ret bool
		return ret
	}
	return *o.V3
}

// GetV3Ok returns a tuple with the V3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUpdate0) GetV3Ok() (*bool, bool) {
	if o == nil || o.V3 == nil {
		return nil, false
	}
	return o.V3, true
}

// HasV3 returns a boolean if a field has been set.
func (o *SnmpUpdate0) HasV3() bool {
	if o != nil && o.V3 != nil {
		return true
	}

	return false
}

// SetV3 gets a reference to the given bool and assigns it to the V3 field.
func (o *SnmpUpdate0) SetV3(v bool) {
	o.V3 = &v
}

// GetCommunity returns the Community field value if set, zero value otherwise.
func (o *SnmpUpdate0) GetCommunity() string {
	if o == nil || o.Community == nil {
		var ret string
		return ret
	}
	return *o.Community
}

// GetCommunityOk returns a tuple with the Community field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUpdate0) GetCommunityOk() (*string, bool) {
	if o == nil || o.Community == nil {
		return nil, false
	}
	return o.Community, true
}

// HasCommunity returns a boolean if a field has been set.
func (o *SnmpUpdate0) HasCommunity() bool {
	if o != nil && o.Community != nil {
		return true
	}

	return false
}

// SetCommunity gets a reference to the given string and assigns it to the Community field.
func (o *SnmpUpdate0) SetCommunity(v string) {
	o.Community = &v
}

// GetV3Username returns the V3Username field value if set, zero value otherwise.
func (o *SnmpUpdate0) GetV3Username() string {
	if o == nil || o.V3Username == nil {
		var ret string
		return ret
	}
	return *o.V3Username
}

// GetV3UsernameOk returns a tuple with the V3Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUpdate0) GetV3UsernameOk() (*string, bool) {
	if o == nil || o.V3Username == nil {
		return nil, false
	}
	return o.V3Username, true
}

// HasV3Username returns a boolean if a field has been set.
func (o *SnmpUpdate0) HasV3Username() bool {
	if o != nil && o.V3Username != nil {
		return true
	}

	return false
}

// SetV3Username gets a reference to the given string and assigns it to the V3Username field.
func (o *SnmpUpdate0) SetV3Username(v string) {
	o.V3Username = &v
}

// GetV3Authtype returns the V3Authtype field value if set, zero value otherwise.
func (o *SnmpUpdate0) GetV3Authtype() string {
	if o == nil || o.V3Authtype == nil {
		var ret string
		return ret
	}
	return *o.V3Authtype
}

// GetV3AuthtypeOk returns a tuple with the V3Authtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUpdate0) GetV3AuthtypeOk() (*string, bool) {
	if o == nil || o.V3Authtype == nil {
		return nil, false
	}
	return o.V3Authtype, true
}

// HasV3Authtype returns a boolean if a field has been set.
func (o *SnmpUpdate0) HasV3Authtype() bool {
	if o != nil && o.V3Authtype != nil {
		return true
	}

	return false
}

// SetV3Authtype gets a reference to the given string and assigns it to the V3Authtype field.
func (o *SnmpUpdate0) SetV3Authtype(v string) {
	o.V3Authtype = &v
}

// GetV3Password returns the V3Password field value if set, zero value otherwise.
func (o *SnmpUpdate0) GetV3Password() string {
	if o == nil || o.V3Password == nil {
		var ret string
		return ret
	}
	return *o.V3Password
}

// GetV3PasswordOk returns a tuple with the V3Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUpdate0) GetV3PasswordOk() (*string, bool) {
	if o == nil || o.V3Password == nil {
		return nil, false
	}
	return o.V3Password, true
}

// HasV3Password returns a boolean if a field has been set.
func (o *SnmpUpdate0) HasV3Password() bool {
	if o != nil && o.V3Password != nil {
		return true
	}

	return false
}

// SetV3Password gets a reference to the given string and assigns it to the V3Password field.
func (o *SnmpUpdate0) SetV3Password(v string) {
	o.V3Password = &v
}

// GetV3Privproto returns the V3Privproto field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnmpUpdate0) GetV3Privproto() string {
	if o == nil || o.V3Privproto.Get() == nil {
		var ret string
		return ret
	}
	return *o.V3Privproto.Get()
}

// GetV3PrivprotoOk returns a tuple with the V3Privproto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnmpUpdate0) GetV3PrivprotoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.V3Privproto.Get(), o.V3Privproto.IsSet()
}

// HasV3Privproto returns a boolean if a field has been set.
func (o *SnmpUpdate0) HasV3Privproto() bool {
	if o != nil && o.V3Privproto.IsSet() {
		return true
	}

	return false
}

// SetV3Privproto gets a reference to the given NullableString and assigns it to the V3Privproto field.
func (o *SnmpUpdate0) SetV3Privproto(v string) {
	o.V3Privproto.Set(&v)
}

// SetV3PrivprotoNil sets the value for V3Privproto to be an explicit nil
func (o *SnmpUpdate0) SetV3PrivprotoNil() {
	o.V3Privproto.Set(nil)
}

// UnsetV3Privproto ensures that no value is present for V3Privproto, not even an explicit nil
func (o *SnmpUpdate0) UnsetV3Privproto() {
	o.V3Privproto.Unset()
}

// GetV3Privpassphrase returns the V3Privpassphrase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnmpUpdate0) GetV3Privpassphrase() string {
	if o == nil || o.V3Privpassphrase.Get() == nil {
		var ret string
		return ret
	}
	return *o.V3Privpassphrase.Get()
}

// GetV3PrivpassphraseOk returns a tuple with the V3Privpassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnmpUpdate0) GetV3PrivpassphraseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.V3Privpassphrase.Get(), o.V3Privpassphrase.IsSet()
}

// HasV3Privpassphrase returns a boolean if a field has been set.
func (o *SnmpUpdate0) HasV3Privpassphrase() bool {
	if o != nil && o.V3Privpassphrase.IsSet() {
		return true
	}

	return false
}

// SetV3Privpassphrase gets a reference to the given NullableString and assigns it to the V3Privpassphrase field.
func (o *SnmpUpdate0) SetV3Privpassphrase(v string) {
	o.V3Privpassphrase.Set(&v)
}

// SetV3PrivpassphraseNil sets the value for V3Privpassphrase to be an explicit nil
func (o *SnmpUpdate0) SetV3PrivpassphraseNil() {
	o.V3Privpassphrase.Set(nil)
}

// UnsetV3Privpassphrase ensures that no value is present for V3Privpassphrase, not even an explicit nil
func (o *SnmpUpdate0) UnsetV3Privpassphrase() {
	o.V3Privpassphrase.Unset()
}

// GetLoglevel returns the Loglevel field value if set, zero value otherwise.
func (o *SnmpUpdate0) GetLoglevel() int32 {
	if o == nil || o.Loglevel == nil {
		var ret int32
		return ret
	}
	return *o.Loglevel
}

// GetLoglevelOk returns a tuple with the Loglevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUpdate0) GetLoglevelOk() (*int32, bool) {
	if o == nil || o.Loglevel == nil {
		return nil, false
	}
	return o.Loglevel, true
}

// HasLoglevel returns a boolean if a field has been set.
func (o *SnmpUpdate0) HasLoglevel() bool {
	if o != nil && o.Loglevel != nil {
		return true
	}

	return false
}

// SetLoglevel gets a reference to the given int32 and assigns it to the Loglevel field.
func (o *SnmpUpdate0) SetLoglevel(v int32) {
	o.Loglevel = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SnmpUpdate0) GetOptions() string {
	if o == nil || o.Options == nil {
		var ret string
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUpdate0) GetOptionsOk() (*string, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SnmpUpdate0) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given string and assigns it to the Options field.
func (o *SnmpUpdate0) SetOptions(v string) {
	o.Options = &v
}

// GetZilstat returns the Zilstat field value if set, zero value otherwise.
func (o *SnmpUpdate0) GetZilstat() bool {
	if o == nil || o.Zilstat == nil {
		var ret bool
		return ret
	}
	return *o.Zilstat
}

// GetZilstatOk returns a tuple with the Zilstat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUpdate0) GetZilstatOk() (*bool, bool) {
	if o == nil || o.Zilstat == nil {
		return nil, false
	}
	return o.Zilstat, true
}

// HasZilstat returns a boolean if a field has been set.
func (o *SnmpUpdate0) HasZilstat() bool {
	if o != nil && o.Zilstat != nil {
		return true
	}

	return false
}

// SetZilstat gets a reference to the given bool and assigns it to the Zilstat field.
func (o *SnmpUpdate0) SetZilstat(v bool) {
	o.Zilstat = &v
}

func (o SnmpUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Contact != nil {
		toSerialize["contact"] = o.Contact
	}
	if o.Traps != nil {
		toSerialize["traps"] = o.Traps
	}
	if o.V3 != nil {
		toSerialize["v3"] = o.V3
	}
	if o.Community != nil {
		toSerialize["community"] = o.Community
	}
	if o.V3Username != nil {
		toSerialize["v3_username"] = o.V3Username
	}
	if o.V3Authtype != nil {
		toSerialize["v3_authtype"] = o.V3Authtype
	}
	if o.V3Password != nil {
		toSerialize["v3_password"] = o.V3Password
	}
	if o.V3Privproto.IsSet() {
		toSerialize["v3_privproto"] = o.V3Privproto.Get()
	}
	if o.V3Privpassphrase.IsSet() {
		toSerialize["v3_privpassphrase"] = o.V3Privpassphrase.Get()
	}
	if o.Loglevel != nil {
		toSerialize["loglevel"] = o.Loglevel
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Zilstat != nil {
		toSerialize["zilstat"] = o.Zilstat
	}
	return json.Marshal(toSerialize)
}

type NullableSnmpUpdate0 struct {
	value *SnmpUpdate0
	isSet bool
}

func (v NullableSnmpUpdate0) Get() *SnmpUpdate0 {
	return v.value
}

func (v *NullableSnmpUpdate0) Set(val *SnmpUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpUpdate0(val *SnmpUpdate0) *NullableSnmpUpdate0 {
	return &NullableSnmpUpdate0{value: val, isSet: true}
}

func (v NullableSnmpUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
