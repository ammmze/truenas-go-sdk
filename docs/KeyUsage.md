# KeyUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**DigitalSignature** | Pointer to **bool** |  | [optional] [default to false]
**ContentCommitment** | Pointer to **bool** |  | [optional] [default to false]
**KeyEncipherment** | Pointer to **bool** |  | [optional] [default to false]
**DataEncipherment** | Pointer to **bool** |  | [optional] [default to false]
**KeyAgreement** | Pointer to **bool** |  | [optional] [default to false]
**KeyCertSign** | Pointer to **bool** |  | [optional] [default to false]
**CrlSign** | Pointer to **bool** |  | [optional] [default to false]
**EncipherOnly** | Pointer to **bool** |  | [optional] [default to false]
**DecipherOnly** | Pointer to **bool** |  | [optional] [default to false]
**ExtensionCritical** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewKeyUsage

`func NewKeyUsage() *KeyUsage`

NewKeyUsage instantiates a new KeyUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyUsageWithDefaults

`func NewKeyUsageWithDefaults() *KeyUsage`

NewKeyUsageWithDefaults instantiates a new KeyUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *KeyUsage) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyUsage) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyUsage) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyUsage) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDigitalSignature

`func (o *KeyUsage) GetDigitalSignature() bool`

GetDigitalSignature returns the DigitalSignature field if non-nil, zero value otherwise.

### GetDigitalSignatureOk

`func (o *KeyUsage) GetDigitalSignatureOk() (*bool, bool)`

GetDigitalSignatureOk returns a tuple with the DigitalSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalSignature

`func (o *KeyUsage) SetDigitalSignature(v bool)`

SetDigitalSignature sets DigitalSignature field to given value.

### HasDigitalSignature

`func (o *KeyUsage) HasDigitalSignature() bool`

HasDigitalSignature returns a boolean if a field has been set.

### GetContentCommitment

`func (o *KeyUsage) GetContentCommitment() bool`

GetContentCommitment returns the ContentCommitment field if non-nil, zero value otherwise.

### GetContentCommitmentOk

`func (o *KeyUsage) GetContentCommitmentOk() (*bool, bool)`

GetContentCommitmentOk returns a tuple with the ContentCommitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCommitment

`func (o *KeyUsage) SetContentCommitment(v bool)`

SetContentCommitment sets ContentCommitment field to given value.

### HasContentCommitment

`func (o *KeyUsage) HasContentCommitment() bool`

HasContentCommitment returns a boolean if a field has been set.

### GetKeyEncipherment

`func (o *KeyUsage) GetKeyEncipherment() bool`

GetKeyEncipherment returns the KeyEncipherment field if non-nil, zero value otherwise.

### GetKeyEnciphermentOk

`func (o *KeyUsage) GetKeyEnciphermentOk() (*bool, bool)`

GetKeyEnciphermentOk returns a tuple with the KeyEncipherment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyEncipherment

`func (o *KeyUsage) SetKeyEncipherment(v bool)`

SetKeyEncipherment sets KeyEncipherment field to given value.

### HasKeyEncipherment

`func (o *KeyUsage) HasKeyEncipherment() bool`

HasKeyEncipherment returns a boolean if a field has been set.

### GetDataEncipherment

`func (o *KeyUsage) GetDataEncipherment() bool`

GetDataEncipherment returns the DataEncipherment field if non-nil, zero value otherwise.

### GetDataEnciphermentOk

`func (o *KeyUsage) GetDataEnciphermentOk() (*bool, bool)`

GetDataEnciphermentOk returns a tuple with the DataEncipherment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataEncipherment

`func (o *KeyUsage) SetDataEncipherment(v bool)`

SetDataEncipherment sets DataEncipherment field to given value.

### HasDataEncipherment

`func (o *KeyUsage) HasDataEncipherment() bool`

HasDataEncipherment returns a boolean if a field has been set.

### GetKeyAgreement

`func (o *KeyUsage) GetKeyAgreement() bool`

GetKeyAgreement returns the KeyAgreement field if non-nil, zero value otherwise.

### GetKeyAgreementOk

`func (o *KeyUsage) GetKeyAgreementOk() (*bool, bool)`

GetKeyAgreementOk returns a tuple with the KeyAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAgreement

`func (o *KeyUsage) SetKeyAgreement(v bool)`

SetKeyAgreement sets KeyAgreement field to given value.

### HasKeyAgreement

`func (o *KeyUsage) HasKeyAgreement() bool`

HasKeyAgreement returns a boolean if a field has been set.

### GetKeyCertSign

`func (o *KeyUsage) GetKeyCertSign() bool`

GetKeyCertSign returns the KeyCertSign field if non-nil, zero value otherwise.

### GetKeyCertSignOk

`func (o *KeyUsage) GetKeyCertSignOk() (*bool, bool)`

GetKeyCertSignOk returns a tuple with the KeyCertSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyCertSign

`func (o *KeyUsage) SetKeyCertSign(v bool)`

SetKeyCertSign sets KeyCertSign field to given value.

### HasKeyCertSign

`func (o *KeyUsage) HasKeyCertSign() bool`

HasKeyCertSign returns a boolean if a field has been set.

### GetCrlSign

`func (o *KeyUsage) GetCrlSign() bool`

GetCrlSign returns the CrlSign field if non-nil, zero value otherwise.

### GetCrlSignOk

`func (o *KeyUsage) GetCrlSignOk() (*bool, bool)`

GetCrlSignOk returns a tuple with the CrlSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlSign

`func (o *KeyUsage) SetCrlSign(v bool)`

SetCrlSign sets CrlSign field to given value.

### HasCrlSign

`func (o *KeyUsage) HasCrlSign() bool`

HasCrlSign returns a boolean if a field has been set.

### GetEncipherOnly

`func (o *KeyUsage) GetEncipherOnly() bool`

GetEncipherOnly returns the EncipherOnly field if non-nil, zero value otherwise.

### GetEncipherOnlyOk

`func (o *KeyUsage) GetEncipherOnlyOk() (*bool, bool)`

GetEncipherOnlyOk returns a tuple with the EncipherOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncipherOnly

`func (o *KeyUsage) SetEncipherOnly(v bool)`

SetEncipherOnly sets EncipherOnly field to given value.

### HasEncipherOnly

`func (o *KeyUsage) HasEncipherOnly() bool`

HasEncipherOnly returns a boolean if a field has been set.

### GetDecipherOnly

`func (o *KeyUsage) GetDecipherOnly() bool`

GetDecipherOnly returns the DecipherOnly field if non-nil, zero value otherwise.

### GetDecipherOnlyOk

`func (o *KeyUsage) GetDecipherOnlyOk() (*bool, bool)`

GetDecipherOnlyOk returns a tuple with the DecipherOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecipherOnly

`func (o *KeyUsage) SetDecipherOnly(v bool)`

SetDecipherOnly sets DecipherOnly field to given value.

### HasDecipherOnly

`func (o *KeyUsage) HasDecipherOnly() bool`

HasDecipherOnly returns a boolean if a field has been set.

### GetExtensionCritical

`func (o *KeyUsage) GetExtensionCritical() bool`

GetExtensionCritical returns the ExtensionCritical field if non-nil, zero value otherwise.

### GetExtensionCriticalOk

`func (o *KeyUsage) GetExtensionCriticalOk() (*bool, bool)`

GetExtensionCriticalOk returns a tuple with the ExtensionCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionCritical

`func (o *KeyUsage) SetExtensionCritical(v bool)`

SetExtensionCritical sets ExtensionCritical field to given value.

### HasExtensionCritical

`func (o *KeyUsage) HasExtensionCritical() bool`

HasExtensionCritical returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


