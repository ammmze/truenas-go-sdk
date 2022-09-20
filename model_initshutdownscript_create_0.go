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

// InitshutdownscriptCreate0 struct for InitshutdownscriptCreate0
type InitshutdownscriptCreate0 struct {
	// `type` indicates if a command or script should be executed at `when`.
	Type       *string        `json:"type,omitempty"`
	Command    NullableString `json:"command,omitempty"`
	ScriptText NullableString `json:"script_text,omitempty"`
	Script     NullableString `json:"script,omitempty"`
	// `type` indicates if a command or script should be executed at `when`. There are three choices for `when`:
	When    *string `json:"when,omitempty"`
	Enabled *bool   `json:"enabled,omitempty"`
	Timeout *int32  `json:"timeout,omitempty"`
	Comment *string `json:"comment,omitempty"`
}

// NewInitshutdownscriptCreate0 instantiates a new InitshutdownscriptCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitshutdownscriptCreate0() *InitshutdownscriptCreate0 {
	this := InitshutdownscriptCreate0{}
	var enabled bool = true
	this.Enabled = &enabled
	var timeout int32 = 10
	this.Timeout = &timeout
	var comment string = ""
	this.Comment = &comment
	return &this
}

// NewInitshutdownscriptCreate0WithDefaults instantiates a new InitshutdownscriptCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitshutdownscriptCreate0WithDefaults() *InitshutdownscriptCreate0 {
	this := InitshutdownscriptCreate0{}
	var enabled bool = true
	this.Enabled = &enabled
	var timeout int32 = 10
	this.Timeout = &timeout
	var comment string = ""
	this.Comment = &comment
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InitshutdownscriptCreate0) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitshutdownscriptCreate0) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InitshutdownscriptCreate0) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InitshutdownscriptCreate0) SetType(v string) {
	o.Type = &v
}

// GetCommand returns the Command field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InitshutdownscriptCreate0) GetCommand() string {
	if o == nil || o.Command.Get() == nil {
		var ret string
		return ret
	}
	return *o.Command.Get()
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InitshutdownscriptCreate0) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Command.Get(), o.Command.IsSet()
}

// HasCommand returns a boolean if a field has been set.
func (o *InitshutdownscriptCreate0) HasCommand() bool {
	if o != nil && o.Command.IsSet() {
		return true
	}

	return false
}

// SetCommand gets a reference to the given NullableString and assigns it to the Command field.
func (o *InitshutdownscriptCreate0) SetCommand(v string) {
	o.Command.Set(&v)
}

// SetCommandNil sets the value for Command to be an explicit nil
func (o *InitshutdownscriptCreate0) SetCommandNil() {
	o.Command.Set(nil)
}

// UnsetCommand ensures that no value is present for Command, not even an explicit nil
func (o *InitshutdownscriptCreate0) UnsetCommand() {
	o.Command.Unset()
}

// GetScriptText returns the ScriptText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InitshutdownscriptCreate0) GetScriptText() string {
	if o == nil || o.ScriptText.Get() == nil {
		var ret string
		return ret
	}
	return *o.ScriptText.Get()
}

// GetScriptTextOk returns a tuple with the ScriptText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InitshutdownscriptCreate0) GetScriptTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScriptText.Get(), o.ScriptText.IsSet()
}

// HasScriptText returns a boolean if a field has been set.
func (o *InitshutdownscriptCreate0) HasScriptText() bool {
	if o != nil && o.ScriptText.IsSet() {
		return true
	}

	return false
}

// SetScriptText gets a reference to the given NullableString and assigns it to the ScriptText field.
func (o *InitshutdownscriptCreate0) SetScriptText(v string) {
	o.ScriptText.Set(&v)
}

// SetScriptTextNil sets the value for ScriptText to be an explicit nil
func (o *InitshutdownscriptCreate0) SetScriptTextNil() {
	o.ScriptText.Set(nil)
}

// UnsetScriptText ensures that no value is present for ScriptText, not even an explicit nil
func (o *InitshutdownscriptCreate0) UnsetScriptText() {
	o.ScriptText.Unset()
}

// GetScript returns the Script field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InitshutdownscriptCreate0) GetScript() string {
	if o == nil || o.Script.Get() == nil {
		var ret string
		return ret
	}
	return *o.Script.Get()
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InitshutdownscriptCreate0) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Script.Get(), o.Script.IsSet()
}

// HasScript returns a boolean if a field has been set.
func (o *InitshutdownscriptCreate0) HasScript() bool {
	if o != nil && o.Script.IsSet() {
		return true
	}

	return false
}

// SetScript gets a reference to the given NullableString and assigns it to the Script field.
func (o *InitshutdownscriptCreate0) SetScript(v string) {
	o.Script.Set(&v)
}

// SetScriptNil sets the value for Script to be an explicit nil
func (o *InitshutdownscriptCreate0) SetScriptNil() {
	o.Script.Set(nil)
}

// UnsetScript ensures that no value is present for Script, not even an explicit nil
func (o *InitshutdownscriptCreate0) UnsetScript() {
	o.Script.Unset()
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *InitshutdownscriptCreate0) GetWhen() string {
	if o == nil || o.When == nil {
		var ret string
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitshutdownscriptCreate0) GetWhenOk() (*string, bool) {
	if o == nil || o.When == nil {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *InitshutdownscriptCreate0) HasWhen() bool {
	if o != nil && o.When != nil {
		return true
	}

	return false
}

// SetWhen gets a reference to the given string and assigns it to the When field.
func (o *InitshutdownscriptCreate0) SetWhen(v string) {
	o.When = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InitshutdownscriptCreate0) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitshutdownscriptCreate0) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InitshutdownscriptCreate0) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InitshutdownscriptCreate0) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *InitshutdownscriptCreate0) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitshutdownscriptCreate0) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *InitshutdownscriptCreate0) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *InitshutdownscriptCreate0) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *InitshutdownscriptCreate0) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitshutdownscriptCreate0) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *InitshutdownscriptCreate0) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *InitshutdownscriptCreate0) SetComment(v string) {
	o.Comment = &v
}

func (o InitshutdownscriptCreate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Command.IsSet() {
		toSerialize["command"] = o.Command.Get()
	}
	if o.ScriptText.IsSet() {
		toSerialize["script_text"] = o.ScriptText.Get()
	}
	if o.Script.IsSet() {
		toSerialize["script"] = o.Script.Get()
	}
	if o.When != nil {
		toSerialize["when"] = o.When
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableInitshutdownscriptCreate0 struct {
	value *InitshutdownscriptCreate0
	isSet bool
}

func (v NullableInitshutdownscriptCreate0) Get() *InitshutdownscriptCreate0 {
	return v.value
}

func (v *NullableInitshutdownscriptCreate0) Set(val *InitshutdownscriptCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableInitshutdownscriptCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableInitshutdownscriptCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitshutdownscriptCreate0(val *InitshutdownscriptCreate0) *NullableInitshutdownscriptCreate0 {
	return &NullableInitshutdownscriptCreate0{value: val, isSet: true}
}

func (v NullableInitshutdownscriptCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitshutdownscriptCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
