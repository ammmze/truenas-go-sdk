# CertExtensions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicConstraints** | Pointer to [**BasicConstraints**](BasicConstraints.md) |  | [optional] 
**AuthorityKeyIdentifier** | Pointer to [**AuthorityKeyIdentifier**](AuthorityKeyIdentifier.md) |  | [optional] 
**ExtendedKeyUsage** | Pointer to [**ExtendedKeyUsage**](ExtendedKeyUsage.md) |  | [optional] 
**KeyUsage** | Pointer to [**KeyUsage**](KeyUsage.md) |  | [optional] 

## Methods

### NewCertExtensions

`func NewCertExtensions() *CertExtensions`

NewCertExtensions instantiates a new CertExtensions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertExtensionsWithDefaults

`func NewCertExtensionsWithDefaults() *CertExtensions`

NewCertExtensionsWithDefaults instantiates a new CertExtensions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicConstraints

`func (o *CertExtensions) GetBasicConstraints() BasicConstraints`

GetBasicConstraints returns the BasicConstraints field if non-nil, zero value otherwise.

### GetBasicConstraintsOk

`func (o *CertExtensions) GetBasicConstraintsOk() (*BasicConstraints, bool)`

GetBasicConstraintsOk returns a tuple with the BasicConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicConstraints

`func (o *CertExtensions) SetBasicConstraints(v BasicConstraints)`

SetBasicConstraints sets BasicConstraints field to given value.

### HasBasicConstraints

`func (o *CertExtensions) HasBasicConstraints() bool`

HasBasicConstraints returns a boolean if a field has been set.

### GetAuthorityKeyIdentifier

`func (o *CertExtensions) GetAuthorityKeyIdentifier() AuthorityKeyIdentifier`

GetAuthorityKeyIdentifier returns the AuthorityKeyIdentifier field if non-nil, zero value otherwise.

### GetAuthorityKeyIdentifierOk

`func (o *CertExtensions) GetAuthorityKeyIdentifierOk() (*AuthorityKeyIdentifier, bool)`

GetAuthorityKeyIdentifierOk returns a tuple with the AuthorityKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityKeyIdentifier

`func (o *CertExtensions) SetAuthorityKeyIdentifier(v AuthorityKeyIdentifier)`

SetAuthorityKeyIdentifier sets AuthorityKeyIdentifier field to given value.

### HasAuthorityKeyIdentifier

`func (o *CertExtensions) HasAuthorityKeyIdentifier() bool`

HasAuthorityKeyIdentifier returns a boolean if a field has been set.

### GetExtendedKeyUsage

`func (o *CertExtensions) GetExtendedKeyUsage() ExtendedKeyUsage`

GetExtendedKeyUsage returns the ExtendedKeyUsage field if non-nil, zero value otherwise.

### GetExtendedKeyUsageOk

`func (o *CertExtensions) GetExtendedKeyUsageOk() (*ExtendedKeyUsage, bool)`

GetExtendedKeyUsageOk returns a tuple with the ExtendedKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsage

`func (o *CertExtensions) SetExtendedKeyUsage(v ExtendedKeyUsage)`

SetExtendedKeyUsage sets ExtendedKeyUsage field to given value.

### HasExtendedKeyUsage

`func (o *CertExtensions) HasExtendedKeyUsage() bool`

HasExtendedKeyUsage returns a boolean if a field has been set.

### GetKeyUsage

`func (o *CertExtensions) GetKeyUsage() KeyUsage`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *CertExtensions) GetKeyUsageOk() (*KeyUsage, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *CertExtensions) SetKeyUsage(v KeyUsage)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *CertExtensions) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


