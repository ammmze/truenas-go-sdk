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

// InterfaceChoices0 struct for InterfaceChoices0
type InterfaceChoices0 struct {
	// `bridge_members` will include BRIDGE members.
	BridgeMembers *bool `json:"bridge_members,omitempty"`
	// `lag_ports` will include LINK_AGGREGATION ports.
	LagPorts *bool `json:"lag_ports,omitempty"`
	// `vlan_parent` will include VLAN parent interface.
	VlanParent *bool `json:"vlan_parent,omitempty"`
	// `exclude` is a list of interfaces prefix to remove.
	Exclude      []interface{} `json:"exclude,omitempty"`
	ExcludeTypes []string      `json:"exclude_types,omitempty"`
	Include      []interface{} `json:"include,omitempty"`
}

// NewInterfaceChoices0 instantiates a new InterfaceChoices0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceChoices0() *InterfaceChoices0 {
	this := InterfaceChoices0{}
	var bridgeMembers bool = false
	this.BridgeMembers = &bridgeMembers
	var lagPorts bool = false
	this.LagPorts = &lagPorts
	var vlanParent bool = true
	this.VlanParent = &vlanParent
	return &this
}

// NewInterfaceChoices0WithDefaults instantiates a new InterfaceChoices0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceChoices0WithDefaults() *InterfaceChoices0 {
	this := InterfaceChoices0{}
	var bridgeMembers bool = false
	this.BridgeMembers = &bridgeMembers
	var lagPorts bool = false
	this.LagPorts = &lagPorts
	var vlanParent bool = true
	this.VlanParent = &vlanParent
	return &this
}

// GetBridgeMembers returns the BridgeMembers field value if set, zero value otherwise.
func (o *InterfaceChoices0) GetBridgeMembers() bool {
	if o == nil || o.BridgeMembers == nil {
		var ret bool
		return ret
	}
	return *o.BridgeMembers
}

// GetBridgeMembersOk returns a tuple with the BridgeMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceChoices0) GetBridgeMembersOk() (*bool, bool) {
	if o == nil || o.BridgeMembers == nil {
		return nil, false
	}
	return o.BridgeMembers, true
}

// HasBridgeMembers returns a boolean if a field has been set.
func (o *InterfaceChoices0) HasBridgeMembers() bool {
	if o != nil && o.BridgeMembers != nil {
		return true
	}

	return false
}

// SetBridgeMembers gets a reference to the given bool and assigns it to the BridgeMembers field.
func (o *InterfaceChoices0) SetBridgeMembers(v bool) {
	o.BridgeMembers = &v
}

// GetLagPorts returns the LagPorts field value if set, zero value otherwise.
func (o *InterfaceChoices0) GetLagPorts() bool {
	if o == nil || o.LagPorts == nil {
		var ret bool
		return ret
	}
	return *o.LagPorts
}

// GetLagPortsOk returns a tuple with the LagPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceChoices0) GetLagPortsOk() (*bool, bool) {
	if o == nil || o.LagPorts == nil {
		return nil, false
	}
	return o.LagPorts, true
}

// HasLagPorts returns a boolean if a field has been set.
func (o *InterfaceChoices0) HasLagPorts() bool {
	if o != nil && o.LagPorts != nil {
		return true
	}

	return false
}

// SetLagPorts gets a reference to the given bool and assigns it to the LagPorts field.
func (o *InterfaceChoices0) SetLagPorts(v bool) {
	o.LagPorts = &v
}

// GetVlanParent returns the VlanParent field value if set, zero value otherwise.
func (o *InterfaceChoices0) GetVlanParent() bool {
	if o == nil || o.VlanParent == nil {
		var ret bool
		return ret
	}
	return *o.VlanParent
}

// GetVlanParentOk returns a tuple with the VlanParent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceChoices0) GetVlanParentOk() (*bool, bool) {
	if o == nil || o.VlanParent == nil {
		return nil, false
	}
	return o.VlanParent, true
}

// HasVlanParent returns a boolean if a field has been set.
func (o *InterfaceChoices0) HasVlanParent() bool {
	if o != nil && o.VlanParent != nil {
		return true
	}

	return false
}

// SetVlanParent gets a reference to the given bool and assigns it to the VlanParent field.
func (o *InterfaceChoices0) SetVlanParent(v bool) {
	o.VlanParent = &v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *InterfaceChoices0) GetExclude() []interface{} {
	if o == nil || o.Exclude == nil {
		var ret []interface{}
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceChoices0) GetExcludeOk() ([]interface{}, bool) {
	if o == nil || o.Exclude == nil {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *InterfaceChoices0) HasExclude() bool {
	if o != nil && o.Exclude != nil {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []interface{} and assigns it to the Exclude field.
func (o *InterfaceChoices0) SetExclude(v []interface{}) {
	o.Exclude = v
}

// GetExcludeTypes returns the ExcludeTypes field value if set, zero value otherwise.
func (o *InterfaceChoices0) GetExcludeTypes() []string {
	if o == nil || o.ExcludeTypes == nil {
		var ret []string
		return ret
	}
	return o.ExcludeTypes
}

// GetExcludeTypesOk returns a tuple with the ExcludeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceChoices0) GetExcludeTypesOk() ([]string, bool) {
	if o == nil || o.ExcludeTypes == nil {
		return nil, false
	}
	return o.ExcludeTypes, true
}

// HasExcludeTypes returns a boolean if a field has been set.
func (o *InterfaceChoices0) HasExcludeTypes() bool {
	if o != nil && o.ExcludeTypes != nil {
		return true
	}

	return false
}

// SetExcludeTypes gets a reference to the given []string and assigns it to the ExcludeTypes field.
func (o *InterfaceChoices0) SetExcludeTypes(v []string) {
	o.ExcludeTypes = v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *InterfaceChoices0) GetInclude() []interface{} {
	if o == nil || o.Include == nil {
		var ret []interface{}
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceChoices0) GetIncludeOk() ([]interface{}, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *InterfaceChoices0) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []interface{} and assigns it to the Include field.
func (o *InterfaceChoices0) SetInclude(v []interface{}) {
	o.Include = v
}

func (o InterfaceChoices0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BridgeMembers != nil {
		toSerialize["bridge_members"] = o.BridgeMembers
	}
	if o.LagPorts != nil {
		toSerialize["lag_ports"] = o.LagPorts
	}
	if o.VlanParent != nil {
		toSerialize["vlan_parent"] = o.VlanParent
	}
	if o.Exclude != nil {
		toSerialize["exclude"] = o.Exclude
	}
	if o.ExcludeTypes != nil {
		toSerialize["exclude_types"] = o.ExcludeTypes
	}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}
	return json.Marshal(toSerialize)
}

type NullableInterfaceChoices0 struct {
	value *InterfaceChoices0
	isSet bool
}

func (v NullableInterfaceChoices0) Get() *InterfaceChoices0 {
	return v.value
}

func (v *NullableInterfaceChoices0) Set(val *InterfaceChoices0) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceChoices0) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceChoices0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceChoices0(val *InterfaceChoices0) *NullableInterfaceChoices0 {
	return &NullableInterfaceChoices0{value: val, isSet: true}
}

func (v NullableInterfaceChoices0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceChoices0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
