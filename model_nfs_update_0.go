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

// NfsUpdate0 struct for NfsUpdate0
type NfsUpdate0 struct {
	// `servers` represents number of servers to create.
	Servers *int32 `json:"servers,omitempty"`
	Udp     *bool  `json:"udp,omitempty"`
	// When `allow_nonroot` is set, it allows non-root mount requests to be served.
	AllowNonroot *bool `json:"allow_nonroot,omitempty"`
	// `v4` when set means that we switch from NFSv3 to NFSv4.
	V4 *bool `json:"v4,omitempty"`
	// `v4_v3owner` when set means that system will use NFSv3 ownership model for NFSv4.
	V4V3owner *bool `json:"v4_v3owner,omitempty"`
	// `v4_krb` will force NFS shares to fail if the Kerberos ticket is unavailable.
	V4Krb *bool `json:"v4_krb,omitempty"`
	// `v4_domain` overrides the default DNS domain name for NFSv4.
	V4Domain *string `json:"v4_domain,omitempty"`
	// `bindip` is a list of IP's on which NFS will listen for requests. When it is unset/empty, NFS listens on all available addresses.
	Bindip []string `json:"bindip,omitempty"`
	// `mountd_port` specifies the port mountd(8) binds to.
	MountdPort NullableInt32 `json:"mountd_port,omitempty"`
	// `rpcstatd_port` specifies the port rpc.statd(8) binds to.
	RpcstatdPort NullableInt32 `json:"rpcstatd_port,omitempty"`
	// `rpclockd_port` specifies the port rpclockd_port(8) binds to.
	RpclockdPort    NullableInt32 `json:"rpclockd_port,omitempty"`
	MountdLog       *bool         `json:"mountd_log,omitempty"`
	StatdLockdLog   *bool         `json:"statd_lockd_log,omitempty"`
	UserdManageGids *bool         `json:"userd_manage_gids,omitempty"`
}

// NewNfsUpdate0 instantiates a new NfsUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfsUpdate0() *NfsUpdate0 {
	this := NfsUpdate0{}
	return &this
}

// NewNfsUpdate0WithDefaults instantiates a new NfsUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfsUpdate0WithDefaults() *NfsUpdate0 {
	this := NfsUpdate0{}
	return &this
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *NfsUpdate0) GetServers() int32 {
	if o == nil || o.Servers == nil {
		var ret int32
		return ret
	}
	return *o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsUpdate0) GetServersOk() (*int32, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *NfsUpdate0) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given int32 and assigns it to the Servers field.
func (o *NfsUpdate0) SetServers(v int32) {
	o.Servers = &v
}

// GetUdp returns the Udp field value if set, zero value otherwise.
func (o *NfsUpdate0) GetUdp() bool {
	if o == nil || o.Udp == nil {
		var ret bool
		return ret
	}
	return *o.Udp
}

// GetUdpOk returns a tuple with the Udp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsUpdate0) GetUdpOk() (*bool, bool) {
	if o == nil || o.Udp == nil {
		return nil, false
	}
	return o.Udp, true
}

// HasUdp returns a boolean if a field has been set.
func (o *NfsUpdate0) HasUdp() bool {
	if o != nil && o.Udp != nil {
		return true
	}

	return false
}

// SetUdp gets a reference to the given bool and assigns it to the Udp field.
func (o *NfsUpdate0) SetUdp(v bool) {
	o.Udp = &v
}

// GetAllowNonroot returns the AllowNonroot field value if set, zero value otherwise.
func (o *NfsUpdate0) GetAllowNonroot() bool {
	if o == nil || o.AllowNonroot == nil {
		var ret bool
		return ret
	}
	return *o.AllowNonroot
}

// GetAllowNonrootOk returns a tuple with the AllowNonroot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsUpdate0) GetAllowNonrootOk() (*bool, bool) {
	if o == nil || o.AllowNonroot == nil {
		return nil, false
	}
	return o.AllowNonroot, true
}

// HasAllowNonroot returns a boolean if a field has been set.
func (o *NfsUpdate0) HasAllowNonroot() bool {
	if o != nil && o.AllowNonroot != nil {
		return true
	}

	return false
}

// SetAllowNonroot gets a reference to the given bool and assigns it to the AllowNonroot field.
func (o *NfsUpdate0) SetAllowNonroot(v bool) {
	o.AllowNonroot = &v
}

// GetV4 returns the V4 field value if set, zero value otherwise.
func (o *NfsUpdate0) GetV4() bool {
	if o == nil || o.V4 == nil {
		var ret bool
		return ret
	}
	return *o.V4
}

// GetV4Ok returns a tuple with the V4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsUpdate0) GetV4Ok() (*bool, bool) {
	if o == nil || o.V4 == nil {
		return nil, false
	}
	return o.V4, true
}

// HasV4 returns a boolean if a field has been set.
func (o *NfsUpdate0) HasV4() bool {
	if o != nil && o.V4 != nil {
		return true
	}

	return false
}

// SetV4 gets a reference to the given bool and assigns it to the V4 field.
func (o *NfsUpdate0) SetV4(v bool) {
	o.V4 = &v
}

// GetV4V3owner returns the V4V3owner field value if set, zero value otherwise.
func (o *NfsUpdate0) GetV4V3owner() bool {
	if o == nil || o.V4V3owner == nil {
		var ret bool
		return ret
	}
	return *o.V4V3owner
}

// GetV4V3ownerOk returns a tuple with the V4V3owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsUpdate0) GetV4V3ownerOk() (*bool, bool) {
	if o == nil || o.V4V3owner == nil {
		return nil, false
	}
	return o.V4V3owner, true
}

// HasV4V3owner returns a boolean if a field has been set.
func (o *NfsUpdate0) HasV4V3owner() bool {
	if o != nil && o.V4V3owner != nil {
		return true
	}

	return false
}

// SetV4V3owner gets a reference to the given bool and assigns it to the V4V3owner field.
func (o *NfsUpdate0) SetV4V3owner(v bool) {
	o.V4V3owner = &v
}

// GetV4Krb returns the V4Krb field value if set, zero value otherwise.
func (o *NfsUpdate0) GetV4Krb() bool {
	if o == nil || o.V4Krb == nil {
		var ret bool
		return ret
	}
	return *o.V4Krb
}

// GetV4KrbOk returns a tuple with the V4Krb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsUpdate0) GetV4KrbOk() (*bool, bool) {
	if o == nil || o.V4Krb == nil {
		return nil, false
	}
	return o.V4Krb, true
}

// HasV4Krb returns a boolean if a field has been set.
func (o *NfsUpdate0) HasV4Krb() bool {
	if o != nil && o.V4Krb != nil {
		return true
	}

	return false
}

// SetV4Krb gets a reference to the given bool and assigns it to the V4Krb field.
func (o *NfsUpdate0) SetV4Krb(v bool) {
	o.V4Krb = &v
}

// GetV4Domain returns the V4Domain field value if set, zero value otherwise.
func (o *NfsUpdate0) GetV4Domain() string {
	if o == nil || o.V4Domain == nil {
		var ret string
		return ret
	}
	return *o.V4Domain
}

// GetV4DomainOk returns a tuple with the V4Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsUpdate0) GetV4DomainOk() (*string, bool) {
	if o == nil || o.V4Domain == nil {
		return nil, false
	}
	return o.V4Domain, true
}

// HasV4Domain returns a boolean if a field has been set.
func (o *NfsUpdate0) HasV4Domain() bool {
	if o != nil && o.V4Domain != nil {
		return true
	}

	return false
}

// SetV4Domain gets a reference to the given string and assigns it to the V4Domain field.
func (o *NfsUpdate0) SetV4Domain(v string) {
	o.V4Domain = &v
}

// GetBindip returns the Bindip field value if set, zero value otherwise.
func (o *NfsUpdate0) GetBindip() []string {
	if o == nil || o.Bindip == nil {
		var ret []string
		return ret
	}
	return o.Bindip
}

// GetBindipOk returns a tuple with the Bindip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsUpdate0) GetBindipOk() ([]string, bool) {
	if o == nil || o.Bindip == nil {
		return nil, false
	}
	return o.Bindip, true
}

// HasBindip returns a boolean if a field has been set.
func (o *NfsUpdate0) HasBindip() bool {
	if o != nil && o.Bindip != nil {
		return true
	}

	return false
}

// SetBindip gets a reference to the given []string and assigns it to the Bindip field.
func (o *NfsUpdate0) SetBindip(v []string) {
	o.Bindip = v
}

// GetMountdPort returns the MountdPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NfsUpdate0) GetMountdPort() int32 {
	if o == nil || o.MountdPort.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MountdPort.Get()
}

// GetMountdPortOk returns a tuple with the MountdPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NfsUpdate0) GetMountdPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountdPort.Get(), o.MountdPort.IsSet()
}

// HasMountdPort returns a boolean if a field has been set.
func (o *NfsUpdate0) HasMountdPort() bool {
	if o != nil && o.MountdPort.IsSet() {
		return true
	}

	return false
}

// SetMountdPort gets a reference to the given NullableInt32 and assigns it to the MountdPort field.
func (o *NfsUpdate0) SetMountdPort(v int32) {
	o.MountdPort.Set(&v)
}

// SetMountdPortNil sets the value for MountdPort to be an explicit nil
func (o *NfsUpdate0) SetMountdPortNil() {
	o.MountdPort.Set(nil)
}

// UnsetMountdPort ensures that no value is present for MountdPort, not even an explicit nil
func (o *NfsUpdate0) UnsetMountdPort() {
	o.MountdPort.Unset()
}

// GetRpcstatdPort returns the RpcstatdPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NfsUpdate0) GetRpcstatdPort() int32 {
	if o == nil || o.RpcstatdPort.Get() == nil {
		var ret int32
		return ret
	}
	return *o.RpcstatdPort.Get()
}

// GetRpcstatdPortOk returns a tuple with the RpcstatdPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NfsUpdate0) GetRpcstatdPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RpcstatdPort.Get(), o.RpcstatdPort.IsSet()
}

// HasRpcstatdPort returns a boolean if a field has been set.
func (o *NfsUpdate0) HasRpcstatdPort() bool {
	if o != nil && o.RpcstatdPort.IsSet() {
		return true
	}

	return false
}

// SetRpcstatdPort gets a reference to the given NullableInt32 and assigns it to the RpcstatdPort field.
func (o *NfsUpdate0) SetRpcstatdPort(v int32) {
	o.RpcstatdPort.Set(&v)
}

// SetRpcstatdPortNil sets the value for RpcstatdPort to be an explicit nil
func (o *NfsUpdate0) SetRpcstatdPortNil() {
	o.RpcstatdPort.Set(nil)
}

// UnsetRpcstatdPort ensures that no value is present for RpcstatdPort, not even an explicit nil
func (o *NfsUpdate0) UnsetRpcstatdPort() {
	o.RpcstatdPort.Unset()
}

// GetRpclockdPort returns the RpclockdPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NfsUpdate0) GetRpclockdPort() int32 {
	if o == nil || o.RpclockdPort.Get() == nil {
		var ret int32
		return ret
	}
	return *o.RpclockdPort.Get()
}

// GetRpclockdPortOk returns a tuple with the RpclockdPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NfsUpdate0) GetRpclockdPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RpclockdPort.Get(), o.RpclockdPort.IsSet()
}

// HasRpclockdPort returns a boolean if a field has been set.
func (o *NfsUpdate0) HasRpclockdPort() bool {
	if o != nil && o.RpclockdPort.IsSet() {
		return true
	}

	return false
}

// SetRpclockdPort gets a reference to the given NullableInt32 and assigns it to the RpclockdPort field.
func (o *NfsUpdate0) SetRpclockdPort(v int32) {
	o.RpclockdPort.Set(&v)
}

// SetRpclockdPortNil sets the value for RpclockdPort to be an explicit nil
func (o *NfsUpdate0) SetRpclockdPortNil() {
	o.RpclockdPort.Set(nil)
}

// UnsetRpclockdPort ensures that no value is present for RpclockdPort, not even an explicit nil
func (o *NfsUpdate0) UnsetRpclockdPort() {
	o.RpclockdPort.Unset()
}

// GetMountdLog returns the MountdLog field value if set, zero value otherwise.
func (o *NfsUpdate0) GetMountdLog() bool {
	if o == nil || o.MountdLog == nil {
		var ret bool
		return ret
	}
	return *o.MountdLog
}

// GetMountdLogOk returns a tuple with the MountdLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsUpdate0) GetMountdLogOk() (*bool, bool) {
	if o == nil || o.MountdLog == nil {
		return nil, false
	}
	return o.MountdLog, true
}

// HasMountdLog returns a boolean if a field has been set.
func (o *NfsUpdate0) HasMountdLog() bool {
	if o != nil && o.MountdLog != nil {
		return true
	}

	return false
}

// SetMountdLog gets a reference to the given bool and assigns it to the MountdLog field.
func (o *NfsUpdate0) SetMountdLog(v bool) {
	o.MountdLog = &v
}

// GetStatdLockdLog returns the StatdLockdLog field value if set, zero value otherwise.
func (o *NfsUpdate0) GetStatdLockdLog() bool {
	if o == nil || o.StatdLockdLog == nil {
		var ret bool
		return ret
	}
	return *o.StatdLockdLog
}

// GetStatdLockdLogOk returns a tuple with the StatdLockdLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsUpdate0) GetStatdLockdLogOk() (*bool, bool) {
	if o == nil || o.StatdLockdLog == nil {
		return nil, false
	}
	return o.StatdLockdLog, true
}

// HasStatdLockdLog returns a boolean if a field has been set.
func (o *NfsUpdate0) HasStatdLockdLog() bool {
	if o != nil && o.StatdLockdLog != nil {
		return true
	}

	return false
}

// SetStatdLockdLog gets a reference to the given bool and assigns it to the StatdLockdLog field.
func (o *NfsUpdate0) SetStatdLockdLog(v bool) {
	o.StatdLockdLog = &v
}

// GetUserdManageGids returns the UserdManageGids field value if set, zero value otherwise.
func (o *NfsUpdate0) GetUserdManageGids() bool {
	if o == nil || o.UserdManageGids == nil {
		var ret bool
		return ret
	}
	return *o.UserdManageGids
}

// GetUserdManageGidsOk returns a tuple with the UserdManageGids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsUpdate0) GetUserdManageGidsOk() (*bool, bool) {
	if o == nil || o.UserdManageGids == nil {
		return nil, false
	}
	return o.UserdManageGids, true
}

// HasUserdManageGids returns a boolean if a field has been set.
func (o *NfsUpdate0) HasUserdManageGids() bool {
	if o != nil && o.UserdManageGids != nil {
		return true
	}

	return false
}

// SetUserdManageGids gets a reference to the given bool and assigns it to the UserdManageGids field.
func (o *NfsUpdate0) SetUserdManageGids(v bool) {
	o.UserdManageGids = &v
}

func (o NfsUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	if o.Udp != nil {
		toSerialize["udp"] = o.Udp
	}
	if o.AllowNonroot != nil {
		toSerialize["allow_nonroot"] = o.AllowNonroot
	}
	if o.V4 != nil {
		toSerialize["v4"] = o.V4
	}
	if o.V4V3owner != nil {
		toSerialize["v4_v3owner"] = o.V4V3owner
	}
	if o.V4Krb != nil {
		toSerialize["v4_krb"] = o.V4Krb
	}
	if o.V4Domain != nil {
		toSerialize["v4_domain"] = o.V4Domain
	}
	if o.Bindip != nil {
		toSerialize["bindip"] = o.Bindip
	}
	if o.MountdPort.IsSet() {
		toSerialize["mountd_port"] = o.MountdPort.Get()
	}
	if o.RpcstatdPort.IsSet() {
		toSerialize["rpcstatd_port"] = o.RpcstatdPort.Get()
	}
	if o.RpclockdPort.IsSet() {
		toSerialize["rpclockd_port"] = o.RpclockdPort.Get()
	}
	if o.MountdLog != nil {
		toSerialize["mountd_log"] = o.MountdLog
	}
	if o.StatdLockdLog != nil {
		toSerialize["statd_lockd_log"] = o.StatdLockdLog
	}
	if o.UserdManageGids != nil {
		toSerialize["userd_manage_gids"] = o.UserdManageGids
	}
	return json.Marshal(toSerialize)
}

type NullableNfsUpdate0 struct {
	value *NfsUpdate0
	isSet bool
}

func (v NullableNfsUpdate0) Get() *NfsUpdate0 {
	return v.value
}

func (v *NullableNfsUpdate0) Set(val *NfsUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableNfsUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableNfsUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfsUpdate0(val *NfsUpdate0) *NullableNfsUpdate0 {
	return &NullableNfsUpdate0{value: val, isSet: true}
}

func (v NullableNfsUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfsUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
