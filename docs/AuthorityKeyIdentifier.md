# AuthorityKeyIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorityCertIssuer** | Pointer to **bool** |  | [optional] [default to false]
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**ExtensionCritical** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAuthorityKeyIdentifier

`func NewAuthorityKeyIdentifier() *AuthorityKeyIdentifier`

NewAuthorityKeyIdentifier instantiates a new AuthorityKeyIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorityKeyIdentifierWithDefaults

`func NewAuthorityKeyIdentifierWithDefaults() *AuthorityKeyIdentifier`

NewAuthorityKeyIdentifierWithDefaults instantiates a new AuthorityKeyIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorityCertIssuer

`func (o *AuthorityKeyIdentifier) GetAuthorityCertIssuer() bool`

GetAuthorityCertIssuer returns the AuthorityCertIssuer field if non-nil, zero value otherwise.

### GetAuthorityCertIssuerOk

`func (o *AuthorityKeyIdentifier) GetAuthorityCertIssuerOk() (*bool, bool)`

GetAuthorityCertIssuerOk returns a tuple with the AuthorityCertIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityCertIssuer

`func (o *AuthorityKeyIdentifier) SetAuthorityCertIssuer(v bool)`

SetAuthorityCertIssuer sets AuthorityCertIssuer field to given value.

### HasAuthorityCertIssuer

`func (o *AuthorityKeyIdentifier) HasAuthorityCertIssuer() bool`

HasAuthorityCertIssuer returns a boolean if a field has been set.

### GetEnabled

`func (o *AuthorityKeyIdentifier) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthorityKeyIdentifier) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthorityKeyIdentifier) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AuthorityKeyIdentifier) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtensionCritical

`func (o *AuthorityKeyIdentifier) GetExtensionCritical() bool`

GetExtensionCritical returns the ExtensionCritical field if non-nil, zero value otherwise.

### GetExtensionCriticalOk

`func (o *AuthorityKeyIdentifier) GetExtensionCriticalOk() (*bool, bool)`

GetExtensionCriticalOk returns a tuple with the ExtensionCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionCritical

`func (o *AuthorityKeyIdentifier) SetExtensionCritical(v bool)`

SetExtensionCritical sets ExtensionCritical field to given value.

### HasExtensionCritical

`func (o *AuthorityKeyIdentifier) HasExtensionCritical() bool`

HasExtensionCritical returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


