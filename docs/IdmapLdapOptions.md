# IdmapLdapOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdapBaseDn** | Pointer to **string** |  | [optional] 
**LdapUserDn** | Pointer to **string** |  | [optional] 
**LdapUserDnPassword** | Pointer to **string** |  | [optional] 
**LdapUrl** | Pointer to **string** |  | [optional] 
**Readonly** | Pointer to **bool** |  | [optional] [default to false]
**Ssl** | Pointer to **string** |  | [optional] [default to "true"]
**ValidateCertificates** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewIdmapLdapOptions

`func NewIdmapLdapOptions() *IdmapLdapOptions`

NewIdmapLdapOptions instantiates a new IdmapLdapOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdmapLdapOptionsWithDefaults

`func NewIdmapLdapOptionsWithDefaults() *IdmapLdapOptions`

NewIdmapLdapOptionsWithDefaults instantiates a new IdmapLdapOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdapBaseDn

`func (o *IdmapLdapOptions) GetLdapBaseDn() string`

GetLdapBaseDn returns the LdapBaseDn field if non-nil, zero value otherwise.

### GetLdapBaseDnOk

`func (o *IdmapLdapOptions) GetLdapBaseDnOk() (*string, bool)`

GetLdapBaseDnOk returns a tuple with the LdapBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBaseDn

`func (o *IdmapLdapOptions) SetLdapBaseDn(v string)`

SetLdapBaseDn sets LdapBaseDn field to given value.

### HasLdapBaseDn

`func (o *IdmapLdapOptions) HasLdapBaseDn() bool`

HasLdapBaseDn returns a boolean if a field has been set.

### GetLdapUserDn

`func (o *IdmapLdapOptions) GetLdapUserDn() string`

GetLdapUserDn returns the LdapUserDn field if non-nil, zero value otherwise.

### GetLdapUserDnOk

`func (o *IdmapLdapOptions) GetLdapUserDnOk() (*string, bool)`

GetLdapUserDnOk returns a tuple with the LdapUserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserDn

`func (o *IdmapLdapOptions) SetLdapUserDn(v string)`

SetLdapUserDn sets LdapUserDn field to given value.

### HasLdapUserDn

`func (o *IdmapLdapOptions) HasLdapUserDn() bool`

HasLdapUserDn returns a boolean if a field has been set.

### GetLdapUserDnPassword

`func (o *IdmapLdapOptions) GetLdapUserDnPassword() string`

GetLdapUserDnPassword returns the LdapUserDnPassword field if non-nil, zero value otherwise.

### GetLdapUserDnPasswordOk

`func (o *IdmapLdapOptions) GetLdapUserDnPasswordOk() (*string, bool)`

GetLdapUserDnPasswordOk returns a tuple with the LdapUserDnPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserDnPassword

`func (o *IdmapLdapOptions) SetLdapUserDnPassword(v string)`

SetLdapUserDnPassword sets LdapUserDnPassword field to given value.

### HasLdapUserDnPassword

`func (o *IdmapLdapOptions) HasLdapUserDnPassword() bool`

HasLdapUserDnPassword returns a boolean if a field has been set.

### GetLdapUrl

`func (o *IdmapLdapOptions) GetLdapUrl() string`

GetLdapUrl returns the LdapUrl field if non-nil, zero value otherwise.

### GetLdapUrlOk

`func (o *IdmapLdapOptions) GetLdapUrlOk() (*string, bool)`

GetLdapUrlOk returns a tuple with the LdapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUrl

`func (o *IdmapLdapOptions) SetLdapUrl(v string)`

SetLdapUrl sets LdapUrl field to given value.

### HasLdapUrl

`func (o *IdmapLdapOptions) HasLdapUrl() bool`

HasLdapUrl returns a boolean if a field has been set.

### GetReadonly

`func (o *IdmapLdapOptions) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *IdmapLdapOptions) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *IdmapLdapOptions) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *IdmapLdapOptions) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetSsl

`func (o *IdmapLdapOptions) GetSsl() string`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *IdmapLdapOptions) GetSslOk() (*string, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *IdmapLdapOptions) SetSsl(v string)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *IdmapLdapOptions) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetValidateCertificates

`func (o *IdmapLdapOptions) GetValidateCertificates() bool`

GetValidateCertificates returns the ValidateCertificates field if non-nil, zero value otherwise.

### GetValidateCertificatesOk

`func (o *IdmapLdapOptions) GetValidateCertificatesOk() (*bool, bool)`

GetValidateCertificatesOk returns a tuple with the ValidateCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateCertificates

`func (o *IdmapLdapOptions) SetValidateCertificates(v bool)`

SetValidateCertificates sets ValidateCertificates field to given value.

### HasValidateCertificates

`func (o *IdmapLdapOptions) HasValidateCertificates() bool`

HasValidateCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


