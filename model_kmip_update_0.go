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

// KmipUpdate0 struct for KmipUpdate0
type KmipUpdate0 struct {
	// `enabled` if true, cannot be set to disabled if there are existing keys pending to be synced. However users can still perform this action by enabling `force_clear`.
	Enabled *bool `json:"enabled,omitempty"`
	// `manage_zfs_keys`/`manage_sed_disks` when enabled will sync keys from local database to remote KMIP server. When disabled, if there are any keys left to be retrieved from the KMIP server, it will sync them back to local database.
	ManageSedDisks *bool `json:"manage_sed_disks,omitempty"`
	// `manage_zfs_keys`/`manage_sed_disks` when enabled will sync keys from local database to remote KMIP server. When disabled, if there are any keys left to be retrieved from the KMIP server, it will sync them back to local database.
	ManageZfsKeys *bool `json:"manage_zfs_keys,omitempty"`
	// System currently authenticates connection with remote KMIP Server with a TLS handshake. `certificate` and
	Certificate NullableInt32 `json:"certificate,omitempty"`
	// `certificate_authority` determine the certs which will be used to initiate the TLS handshake with `server`.
	CertificateAuthority NullableInt32 `json:"certificate_authority,omitempty"`
	Port                 *int32        `json:"port,omitempty"`
	// `certificate_authority` determine the certs which will be used to initiate the TLS handshake with `server`. `validate` is enabled by default. When enabled, system will test connection to `server` making sure it's reachable.
	Server NullableString `json:"server,omitempty"`
	// `ssl_version` can be specified to match the ssl configuration being used by KMIP server.
	SslVersion *string `json:"ssl_version,omitempty"`
	// `enabled` if true, cannot be set to disabled if there are existing keys pending to be synced. However users can still perform this action by enabling `force_clear`. `change_server` is a boolean field which allows users to migrate data between two KMIP servers. System will first migrate keys from old KMIP server to local database and then migrate the keys from local database to new KMIP server. If it is unable to retrieve all the keys from old server, this will fail. Users can bypass this by enabling `force_clear`.
	ForceClear *bool `json:"force_clear,omitempty"`
	// `change_server` is a boolean field which allows users to migrate data between two KMIP servers. System will first migrate keys from old KMIP server to local database and then migrate the keys from local database to new KMIP server. If it is unable to retrieve all the keys from old server, this will fail. Users can bypass this by enabling `force_clear`.
	ChangeServer *bool `json:"change_server,omitempty"`
	// `validate` is enabled by default. When enabled, system will test connection to `server` making sure it's reachable.
	Validate *bool `json:"validate,omitempty"`
}

// NewKmipUpdate0 instantiates a new KmipUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKmipUpdate0() *KmipUpdate0 {
	this := KmipUpdate0{}
	return &this
}

// NewKmipUpdate0WithDefaults instantiates a new KmipUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKmipUpdate0WithDefaults() *KmipUpdate0 {
	this := KmipUpdate0{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *KmipUpdate0) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipUpdate0) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *KmipUpdate0) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *KmipUpdate0) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetManageSedDisks returns the ManageSedDisks field value if set, zero value otherwise.
func (o *KmipUpdate0) GetManageSedDisks() bool {
	if o == nil || o.ManageSedDisks == nil {
		var ret bool
		return ret
	}
	return *o.ManageSedDisks
}

// GetManageSedDisksOk returns a tuple with the ManageSedDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipUpdate0) GetManageSedDisksOk() (*bool, bool) {
	if o == nil || o.ManageSedDisks == nil {
		return nil, false
	}
	return o.ManageSedDisks, true
}

// HasManageSedDisks returns a boolean if a field has been set.
func (o *KmipUpdate0) HasManageSedDisks() bool {
	if o != nil && o.ManageSedDisks != nil {
		return true
	}

	return false
}

// SetManageSedDisks gets a reference to the given bool and assigns it to the ManageSedDisks field.
func (o *KmipUpdate0) SetManageSedDisks(v bool) {
	o.ManageSedDisks = &v
}

// GetManageZfsKeys returns the ManageZfsKeys field value if set, zero value otherwise.
func (o *KmipUpdate0) GetManageZfsKeys() bool {
	if o == nil || o.ManageZfsKeys == nil {
		var ret bool
		return ret
	}
	return *o.ManageZfsKeys
}

// GetManageZfsKeysOk returns a tuple with the ManageZfsKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipUpdate0) GetManageZfsKeysOk() (*bool, bool) {
	if o == nil || o.ManageZfsKeys == nil {
		return nil, false
	}
	return o.ManageZfsKeys, true
}

// HasManageZfsKeys returns a boolean if a field has been set.
func (o *KmipUpdate0) HasManageZfsKeys() bool {
	if o != nil && o.ManageZfsKeys != nil {
		return true
	}

	return false
}

// SetManageZfsKeys gets a reference to the given bool and assigns it to the ManageZfsKeys field.
func (o *KmipUpdate0) SetManageZfsKeys(v bool) {
	o.ManageZfsKeys = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KmipUpdate0) GetCertificate() int32 {
	if o == nil || o.Certificate.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KmipUpdate0) GetCertificateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *KmipUpdate0) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableInt32 and assigns it to the Certificate field.
func (o *KmipUpdate0) SetCertificate(v int32) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *KmipUpdate0) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *KmipUpdate0) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetCertificateAuthority returns the CertificateAuthority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KmipUpdate0) GetCertificateAuthority() int32 {
	if o == nil || o.CertificateAuthority.Get() == nil {
		var ret int32
		return ret
	}
	return *o.CertificateAuthority.Get()
}

// GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KmipUpdate0) GetCertificateAuthorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateAuthority.Get(), o.CertificateAuthority.IsSet()
}

// HasCertificateAuthority returns a boolean if a field has been set.
func (o *KmipUpdate0) HasCertificateAuthority() bool {
	if o != nil && o.CertificateAuthority.IsSet() {
		return true
	}

	return false
}

// SetCertificateAuthority gets a reference to the given NullableInt32 and assigns it to the CertificateAuthority field.
func (o *KmipUpdate0) SetCertificateAuthority(v int32) {
	o.CertificateAuthority.Set(&v)
}

// SetCertificateAuthorityNil sets the value for CertificateAuthority to be an explicit nil
func (o *KmipUpdate0) SetCertificateAuthorityNil() {
	o.CertificateAuthority.Set(nil)
}

// UnsetCertificateAuthority ensures that no value is present for CertificateAuthority, not even an explicit nil
func (o *KmipUpdate0) UnsetCertificateAuthority() {
	o.CertificateAuthority.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *KmipUpdate0) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipUpdate0) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *KmipUpdate0) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *KmipUpdate0) SetPort(v int32) {
	o.Port = &v
}

// GetServer returns the Server field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KmipUpdate0) GetServer() string {
	if o == nil || o.Server.Get() == nil {
		var ret string
		return ret
	}
	return *o.Server.Get()
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KmipUpdate0) GetServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Server.Get(), o.Server.IsSet()
}

// HasServer returns a boolean if a field has been set.
func (o *KmipUpdate0) HasServer() bool {
	if o != nil && o.Server.IsSet() {
		return true
	}

	return false
}

// SetServer gets a reference to the given NullableString and assigns it to the Server field.
func (o *KmipUpdate0) SetServer(v string) {
	o.Server.Set(&v)
}

// SetServerNil sets the value for Server to be an explicit nil
func (o *KmipUpdate0) SetServerNil() {
	o.Server.Set(nil)
}

// UnsetServer ensures that no value is present for Server, not even an explicit nil
func (o *KmipUpdate0) UnsetServer() {
	o.Server.Unset()
}

// GetSslVersion returns the SslVersion field value if set, zero value otherwise.
func (o *KmipUpdate0) GetSslVersion() string {
	if o == nil || o.SslVersion == nil {
		var ret string
		return ret
	}
	return *o.SslVersion
}

// GetSslVersionOk returns a tuple with the SslVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipUpdate0) GetSslVersionOk() (*string, bool) {
	if o == nil || o.SslVersion == nil {
		return nil, false
	}
	return o.SslVersion, true
}

// HasSslVersion returns a boolean if a field has been set.
func (o *KmipUpdate0) HasSslVersion() bool {
	if o != nil && o.SslVersion != nil {
		return true
	}

	return false
}

// SetSslVersion gets a reference to the given string and assigns it to the SslVersion field.
func (o *KmipUpdate0) SetSslVersion(v string) {
	o.SslVersion = &v
}

// GetForceClear returns the ForceClear field value if set, zero value otherwise.
func (o *KmipUpdate0) GetForceClear() bool {
	if o == nil || o.ForceClear == nil {
		var ret bool
		return ret
	}
	return *o.ForceClear
}

// GetForceClearOk returns a tuple with the ForceClear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipUpdate0) GetForceClearOk() (*bool, bool) {
	if o == nil || o.ForceClear == nil {
		return nil, false
	}
	return o.ForceClear, true
}

// HasForceClear returns a boolean if a field has been set.
func (o *KmipUpdate0) HasForceClear() bool {
	if o != nil && o.ForceClear != nil {
		return true
	}

	return false
}

// SetForceClear gets a reference to the given bool and assigns it to the ForceClear field.
func (o *KmipUpdate0) SetForceClear(v bool) {
	o.ForceClear = &v
}

// GetChangeServer returns the ChangeServer field value if set, zero value otherwise.
func (o *KmipUpdate0) GetChangeServer() bool {
	if o == nil || o.ChangeServer == nil {
		var ret bool
		return ret
	}
	return *o.ChangeServer
}

// GetChangeServerOk returns a tuple with the ChangeServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipUpdate0) GetChangeServerOk() (*bool, bool) {
	if o == nil || o.ChangeServer == nil {
		return nil, false
	}
	return o.ChangeServer, true
}

// HasChangeServer returns a boolean if a field has been set.
func (o *KmipUpdate0) HasChangeServer() bool {
	if o != nil && o.ChangeServer != nil {
		return true
	}

	return false
}

// SetChangeServer gets a reference to the given bool and assigns it to the ChangeServer field.
func (o *KmipUpdate0) SetChangeServer(v bool) {
	o.ChangeServer = &v
}

// GetValidate returns the Validate field value if set, zero value otherwise.
func (o *KmipUpdate0) GetValidate() bool {
	if o == nil || o.Validate == nil {
		var ret bool
		return ret
	}
	return *o.Validate
}

// GetValidateOk returns a tuple with the Validate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmipUpdate0) GetValidateOk() (*bool, bool) {
	if o == nil || o.Validate == nil {
		return nil, false
	}
	return o.Validate, true
}

// HasValidate returns a boolean if a field has been set.
func (o *KmipUpdate0) HasValidate() bool {
	if o != nil && o.Validate != nil {
		return true
	}

	return false
}

// SetValidate gets a reference to the given bool and assigns it to the Validate field.
func (o *KmipUpdate0) SetValidate(v bool) {
	o.Validate = &v
}

func (o KmipUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ManageSedDisks != nil {
		toSerialize["manage_sed_disks"] = o.ManageSedDisks
	}
	if o.ManageZfsKeys != nil {
		toSerialize["manage_zfs_keys"] = o.ManageZfsKeys
	}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.CertificateAuthority.IsSet() {
		toSerialize["certificate_authority"] = o.CertificateAuthority.Get()
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Server.IsSet() {
		toSerialize["server"] = o.Server.Get()
	}
	if o.SslVersion != nil {
		toSerialize["ssl_version"] = o.SslVersion
	}
	if o.ForceClear != nil {
		toSerialize["force_clear"] = o.ForceClear
	}
	if o.ChangeServer != nil {
		toSerialize["change_server"] = o.ChangeServer
	}
	if o.Validate != nil {
		toSerialize["validate"] = o.Validate
	}
	return json.Marshal(toSerialize)
}

type NullableKmipUpdate0 struct {
	value *KmipUpdate0
	isSet bool
}

func (v NullableKmipUpdate0) Get() *KmipUpdate0 {
	return v.value
}

func (v *NullableKmipUpdate0) Set(val *KmipUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableKmipUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableKmipUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKmipUpdate0(val *KmipUpdate0) *NullableKmipUpdate0 {
	return &NullableKmipUpdate0{value: val, isSet: true}
}

func (v NullableKmipUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKmipUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
