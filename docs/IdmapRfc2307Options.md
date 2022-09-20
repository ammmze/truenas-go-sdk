# IdmapRfc2307Options

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdapServer** | Pointer to **string** |  | [optional] 
**LdapRealm** | Pointer to **bool** |  | [optional] [default to false]
**BindPathUser** | Pointer to **string** |  | [optional] 
**BindPathGroup** | Pointer to **string** |  | [optional] 
**UserCn** | Pointer to **bool** |  | [optional] [default to false]
**CnRealm** | Pointer to **string** |  | [optional] 
**LdapDomain** | Pointer to **string** |  | [optional] 
**LdapUrl** | Pointer to **string** |  | [optional] 
**LdapUserDn** | Pointer to **string** |  | [optional] 
**LdapUserDnPassword** | Pointer to **string** |  | [optional] 
**Ssl** | Pointer to **string** |  | [optional] [default to "true"]
**ValidateCertificates** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewIdmapRfc2307Options

`func NewIdmapRfc2307Options() *IdmapRfc2307Options`

NewIdmapRfc2307Options instantiates a new IdmapRfc2307Options object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdmapRfc2307OptionsWithDefaults

`func NewIdmapRfc2307OptionsWithDefaults() *IdmapRfc2307Options`

NewIdmapRfc2307OptionsWithDefaults instantiates a new IdmapRfc2307Options object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdapServer

`func (o *IdmapRfc2307Options) GetLdapServer() string`

GetLdapServer returns the LdapServer field if non-nil, zero value otherwise.

### GetLdapServerOk

`func (o *IdmapRfc2307Options) GetLdapServerOk() (*string, bool)`

GetLdapServerOk returns a tuple with the LdapServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServer

`func (o *IdmapRfc2307Options) SetLdapServer(v string)`

SetLdapServer sets LdapServer field to given value.

### HasLdapServer

`func (o *IdmapRfc2307Options) HasLdapServer() bool`

HasLdapServer returns a boolean if a field has been set.

### GetLdapRealm

`func (o *IdmapRfc2307Options) GetLdapRealm() bool`

GetLdapRealm returns the LdapRealm field if non-nil, zero value otherwise.

### GetLdapRealmOk

`func (o *IdmapRfc2307Options) GetLdapRealmOk() (*bool, bool)`

GetLdapRealmOk returns a tuple with the LdapRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapRealm

`func (o *IdmapRfc2307Options) SetLdapRealm(v bool)`

SetLdapRealm sets LdapRealm field to given value.

### HasLdapRealm

`func (o *IdmapRfc2307Options) HasLdapRealm() bool`

HasLdapRealm returns a boolean if a field has been set.

### GetBindPathUser

`func (o *IdmapRfc2307Options) GetBindPathUser() string`

GetBindPathUser returns the BindPathUser field if non-nil, zero value otherwise.

### GetBindPathUserOk

`func (o *IdmapRfc2307Options) GetBindPathUserOk() (*string, bool)`

GetBindPathUserOk returns a tuple with the BindPathUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPathUser

`func (o *IdmapRfc2307Options) SetBindPathUser(v string)`

SetBindPathUser sets BindPathUser field to given value.

### HasBindPathUser

`func (o *IdmapRfc2307Options) HasBindPathUser() bool`

HasBindPathUser returns a boolean if a field has been set.

### GetBindPathGroup

`func (o *IdmapRfc2307Options) GetBindPathGroup() string`

GetBindPathGroup returns the BindPathGroup field if non-nil, zero value otherwise.

### GetBindPathGroupOk

`func (o *IdmapRfc2307Options) GetBindPathGroupOk() (*string, bool)`

GetBindPathGroupOk returns a tuple with the BindPathGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPathGroup

`func (o *IdmapRfc2307Options) SetBindPathGroup(v string)`

SetBindPathGroup sets BindPathGroup field to given value.

### HasBindPathGroup

`func (o *IdmapRfc2307Options) HasBindPathGroup() bool`

HasBindPathGroup returns a boolean if a field has been set.

### GetUserCn

`func (o *IdmapRfc2307Options) GetUserCn() bool`

GetUserCn returns the UserCn field if non-nil, zero value otherwise.

### GetUserCnOk

`func (o *IdmapRfc2307Options) GetUserCnOk() (*bool, bool)`

GetUserCnOk returns a tuple with the UserCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCn

`func (o *IdmapRfc2307Options) SetUserCn(v bool)`

SetUserCn sets UserCn field to given value.

### HasUserCn

`func (o *IdmapRfc2307Options) HasUserCn() bool`

HasUserCn returns a boolean if a field has been set.

### GetCnRealm

`func (o *IdmapRfc2307Options) GetCnRealm() string`

GetCnRealm returns the CnRealm field if non-nil, zero value otherwise.

### GetCnRealmOk

`func (o *IdmapRfc2307Options) GetCnRealmOk() (*string, bool)`

GetCnRealmOk returns a tuple with the CnRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnRealm

`func (o *IdmapRfc2307Options) SetCnRealm(v string)`

SetCnRealm sets CnRealm field to given value.

### HasCnRealm

`func (o *IdmapRfc2307Options) HasCnRealm() bool`

HasCnRealm returns a boolean if a field has been set.

### GetLdapDomain

`func (o *IdmapRfc2307Options) GetLdapDomain() string`

GetLdapDomain returns the LdapDomain field if non-nil, zero value otherwise.

### GetLdapDomainOk

`func (o *IdmapRfc2307Options) GetLdapDomainOk() (*string, bool)`

GetLdapDomainOk returns a tuple with the LdapDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapDomain

`func (o *IdmapRfc2307Options) SetLdapDomain(v string)`

SetLdapDomain sets LdapDomain field to given value.

### HasLdapDomain

`func (o *IdmapRfc2307Options) HasLdapDomain() bool`

HasLdapDomain returns a boolean if a field has been set.

### GetLdapUrl

`func (o *IdmapRfc2307Options) GetLdapUrl() string`

GetLdapUrl returns the LdapUrl field if non-nil, zero value otherwise.

### GetLdapUrlOk

`func (o *IdmapRfc2307Options) GetLdapUrlOk() (*string, bool)`

GetLdapUrlOk returns a tuple with the LdapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUrl

`func (o *IdmapRfc2307Options) SetLdapUrl(v string)`

SetLdapUrl sets LdapUrl field to given value.

### HasLdapUrl

`func (o *IdmapRfc2307Options) HasLdapUrl() bool`

HasLdapUrl returns a boolean if a field has been set.

### GetLdapUserDn

`func (o *IdmapRfc2307Options) GetLdapUserDn() string`

GetLdapUserDn returns the LdapUserDn field if non-nil, zero value otherwise.

### GetLdapUserDnOk

`func (o *IdmapRfc2307Options) GetLdapUserDnOk() (*string, bool)`

GetLdapUserDnOk returns a tuple with the LdapUserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserDn

`func (o *IdmapRfc2307Options) SetLdapUserDn(v string)`

SetLdapUserDn sets LdapUserDn field to given value.

### HasLdapUserDn

`func (o *IdmapRfc2307Options) HasLdapUserDn() bool`

HasLdapUserDn returns a boolean if a field has been set.

### GetLdapUserDnPassword

`func (o *IdmapRfc2307Options) GetLdapUserDnPassword() string`

GetLdapUserDnPassword returns the LdapUserDnPassword field if non-nil, zero value otherwise.

### GetLdapUserDnPasswordOk

`func (o *IdmapRfc2307Options) GetLdapUserDnPasswordOk() (*string, bool)`

GetLdapUserDnPasswordOk returns a tuple with the LdapUserDnPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserDnPassword

`func (o *IdmapRfc2307Options) SetLdapUserDnPassword(v string)`

SetLdapUserDnPassword sets LdapUserDnPassword field to given value.

### HasLdapUserDnPassword

`func (o *IdmapRfc2307Options) HasLdapUserDnPassword() bool`

HasLdapUserDnPassword returns a boolean if a field has been set.

### GetSsl

`func (o *IdmapRfc2307Options) GetSsl() string`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *IdmapRfc2307Options) GetSslOk() (*string, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *IdmapRfc2307Options) SetSsl(v string)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *IdmapRfc2307Options) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetValidateCertificates

`func (o *IdmapRfc2307Options) GetValidateCertificates() bool`

GetValidateCertificates returns the ValidateCertificates field if non-nil, zero value otherwise.

### GetValidateCertificatesOk

`func (o *IdmapRfc2307Options) GetValidateCertificatesOk() (*bool, bool)`

GetValidateCertificatesOk returns a tuple with the ValidateCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateCertificates

`func (o *IdmapRfc2307Options) SetValidateCertificates(v bool)`

SetValidateCertificates sets ValidateCertificates field to given value.

### HasValidateCertificates

`func (o *IdmapRfc2307Options) HasValidateCertificates() bool`

HasValidateCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


