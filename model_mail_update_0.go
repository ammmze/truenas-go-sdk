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

// MailUpdate0 struct for MailUpdate0
type MailUpdate0 struct {
	// `fromemail` is used as a sending address which the mail server will use for sending emails.
	Fromemail *string `json:"fromemail,omitempty"`
	Fromname  *string `json:"fromname,omitempty"`
	// `outgoingserver` is the hostname or IP address of SMTP server used for sending an email.
	Outgoingserver *string `json:"outgoingserver,omitempty"`
	Port           *int32  `json:"port,omitempty"`
	// `security` is type of encryption desired.
	Security *string        `json:"security,omitempty"`
	Smtp     *bool          `json:"smtp,omitempty"`
	User     NullableString `json:"user,omitempty"`
	Pass     NullableString `json:"pass,omitempty"`
	Oauth    *Oauth         `json:"oauth,omitempty"`
}

// NewMailUpdate0 instantiates a new MailUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMailUpdate0() *MailUpdate0 {
	this := MailUpdate0{}
	return &this
}

// NewMailUpdate0WithDefaults instantiates a new MailUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMailUpdate0WithDefaults() *MailUpdate0 {
	this := MailUpdate0{}
	return &this
}

// GetFromemail returns the Fromemail field value if set, zero value otherwise.
func (o *MailUpdate0) GetFromemail() string {
	if o == nil || o.Fromemail == nil {
		var ret string
		return ret
	}
	return *o.Fromemail
}

// GetFromemailOk returns a tuple with the Fromemail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailUpdate0) GetFromemailOk() (*string, bool) {
	if o == nil || o.Fromemail == nil {
		return nil, false
	}
	return o.Fromemail, true
}

// HasFromemail returns a boolean if a field has been set.
func (o *MailUpdate0) HasFromemail() bool {
	if o != nil && o.Fromemail != nil {
		return true
	}

	return false
}

// SetFromemail gets a reference to the given string and assigns it to the Fromemail field.
func (o *MailUpdate0) SetFromemail(v string) {
	o.Fromemail = &v
}

// GetFromname returns the Fromname field value if set, zero value otherwise.
func (o *MailUpdate0) GetFromname() string {
	if o == nil || o.Fromname == nil {
		var ret string
		return ret
	}
	return *o.Fromname
}

// GetFromnameOk returns a tuple with the Fromname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailUpdate0) GetFromnameOk() (*string, bool) {
	if o == nil || o.Fromname == nil {
		return nil, false
	}
	return o.Fromname, true
}

// HasFromname returns a boolean if a field has been set.
func (o *MailUpdate0) HasFromname() bool {
	if o != nil && o.Fromname != nil {
		return true
	}

	return false
}

// SetFromname gets a reference to the given string and assigns it to the Fromname field.
func (o *MailUpdate0) SetFromname(v string) {
	o.Fromname = &v
}

// GetOutgoingserver returns the Outgoingserver field value if set, zero value otherwise.
func (o *MailUpdate0) GetOutgoingserver() string {
	if o == nil || o.Outgoingserver == nil {
		var ret string
		return ret
	}
	return *o.Outgoingserver
}

// GetOutgoingserverOk returns a tuple with the Outgoingserver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailUpdate0) GetOutgoingserverOk() (*string, bool) {
	if o == nil || o.Outgoingserver == nil {
		return nil, false
	}
	return o.Outgoingserver, true
}

// HasOutgoingserver returns a boolean if a field has been set.
func (o *MailUpdate0) HasOutgoingserver() bool {
	if o != nil && o.Outgoingserver != nil {
		return true
	}

	return false
}

// SetOutgoingserver gets a reference to the given string and assigns it to the Outgoingserver field.
func (o *MailUpdate0) SetOutgoingserver(v string) {
	o.Outgoingserver = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *MailUpdate0) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailUpdate0) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *MailUpdate0) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *MailUpdate0) SetPort(v int32) {
	o.Port = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *MailUpdate0) GetSecurity() string {
	if o == nil || o.Security == nil {
		var ret string
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailUpdate0) GetSecurityOk() (*string, bool) {
	if o == nil || o.Security == nil {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *MailUpdate0) HasSecurity() bool {
	if o != nil && o.Security != nil {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given string and assigns it to the Security field.
func (o *MailUpdate0) SetSecurity(v string) {
	o.Security = &v
}

// GetSmtp returns the Smtp field value if set, zero value otherwise.
func (o *MailUpdate0) GetSmtp() bool {
	if o == nil || o.Smtp == nil {
		var ret bool
		return ret
	}
	return *o.Smtp
}

// GetSmtpOk returns a tuple with the Smtp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailUpdate0) GetSmtpOk() (*bool, bool) {
	if o == nil || o.Smtp == nil {
		return nil, false
	}
	return o.Smtp, true
}

// HasSmtp returns a boolean if a field has been set.
func (o *MailUpdate0) HasSmtp() bool {
	if o != nil && o.Smtp != nil {
		return true
	}

	return false
}

// SetSmtp gets a reference to the given bool and assigns it to the Smtp field.
func (o *MailUpdate0) SetSmtp(v bool) {
	o.Smtp = &v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MailUpdate0) GetUser() string {
	if o == nil || o.User.Get() == nil {
		var ret string
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MailUpdate0) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *MailUpdate0) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableString and assigns it to the User field.
func (o *MailUpdate0) SetUser(v string) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *MailUpdate0) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *MailUpdate0) UnsetUser() {
	o.User.Unset()
}

// GetPass returns the Pass field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MailUpdate0) GetPass() string {
	if o == nil || o.Pass.Get() == nil {
		var ret string
		return ret
	}
	return *o.Pass.Get()
}

// GetPassOk returns a tuple with the Pass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MailUpdate0) GetPassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pass.Get(), o.Pass.IsSet()
}

// HasPass returns a boolean if a field has been set.
func (o *MailUpdate0) HasPass() bool {
	if o != nil && o.Pass.IsSet() {
		return true
	}

	return false
}

// SetPass gets a reference to the given NullableString and assigns it to the Pass field.
func (o *MailUpdate0) SetPass(v string) {
	o.Pass.Set(&v)
}

// SetPassNil sets the value for Pass to be an explicit nil
func (o *MailUpdate0) SetPassNil() {
	o.Pass.Set(nil)
}

// UnsetPass ensures that no value is present for Pass, not even an explicit nil
func (o *MailUpdate0) UnsetPass() {
	o.Pass.Unset()
}

// GetOauth returns the Oauth field value if set, zero value otherwise.
func (o *MailUpdate0) GetOauth() Oauth {
	if o == nil || o.Oauth == nil {
		var ret Oauth
		return ret
	}
	return *o.Oauth
}

// GetOauthOk returns a tuple with the Oauth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailUpdate0) GetOauthOk() (*Oauth, bool) {
	if o == nil || o.Oauth == nil {
		return nil, false
	}
	return o.Oauth, true
}

// HasOauth returns a boolean if a field has been set.
func (o *MailUpdate0) HasOauth() bool {
	if o != nil && o.Oauth != nil {
		return true
	}

	return false
}

// SetOauth gets a reference to the given Oauth and assigns it to the Oauth field.
func (o *MailUpdate0) SetOauth(v Oauth) {
	o.Oauth = &v
}

func (o MailUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fromemail != nil {
		toSerialize["fromemail"] = o.Fromemail
	}
	if o.Fromname != nil {
		toSerialize["fromname"] = o.Fromname
	}
	if o.Outgoingserver != nil {
		toSerialize["outgoingserver"] = o.Outgoingserver
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Security != nil {
		toSerialize["security"] = o.Security
	}
	if o.Smtp != nil {
		toSerialize["smtp"] = o.Smtp
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.Pass.IsSet() {
		toSerialize["pass"] = o.Pass.Get()
	}
	if o.Oauth != nil {
		toSerialize["oauth"] = o.Oauth
	}
	return json.Marshal(toSerialize)
}

type NullableMailUpdate0 struct {
	value *MailUpdate0
	isSet bool
}

func (v NullableMailUpdate0) Get() *MailUpdate0 {
	return v.value
}

func (v *NullableMailUpdate0) Set(val *MailUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableMailUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableMailUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMailUpdate0(val *MailUpdate0) *NullableMailUpdate0 {
	return &NullableMailUpdate0{value: val, isSet: true}
}

func (v NullableMailUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMailUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
