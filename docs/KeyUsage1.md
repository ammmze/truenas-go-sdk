# KeyUsage1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**DigitalSignature** | Pointer to **bool** |  | [optional] [default to false]
**ContentCommitment** | Pointer to **bool** |  | [optional] [default to false]
**KeyEncipherment** | Pointer to **bool** |  | [optional] [default to false]
**DataEncipherment** | Pointer to **bool** |  | [optional] [default to false]
**KeyAgreement** | Pointer to **bool** |  | [optional] [default to false]
**KeyCertSign** | Pointer to **bool** |  | [optional] [default to true]
**CrlSign** | Pointer to **bool** |  | [optional] [default to true]
**EncipherOnly** | Pointer to **bool** |  | [optional] [default to false]
**DecipherOnly** | Pointer to **bool** |  | [optional] [default to false]
**ExtensionCritical** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewKeyUsage1

`func NewKeyUsage1() *KeyUsage1`

NewKeyUsage1 instantiates a new KeyUsage1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyUsage1WithDefaults

`func NewKeyUsage1WithDefaults() *KeyUsage1`

NewKeyUsage1WithDefaults instantiates a new KeyUsage1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *KeyUsage1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyUsage1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyUsage1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyUsage1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDigitalSignature

`func (o *KeyUsage1) GetDigitalSignature() bool`

GetDigitalSignature returns the DigitalSignature field if non-nil, zero value otherwise.

### GetDigitalSignatureOk

`func (o *KeyUsage1) GetDigitalSignatureOk() (*bool, bool)`

GetDigitalSignatureOk returns a tuple with the DigitalSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalSignature

`func (o *KeyUsage1) SetDigitalSignature(v bool)`

SetDigitalSignature sets DigitalSignature field to given value.

### HasDigitalSignature

`func (o *KeyUsage1) HasDigitalSignature() bool`

HasDigitalSignature returns a boolean if a field has been set.

### GetContentCommitment

`func (o *KeyUsage1) GetContentCommitment() bool`

GetContentCommitment returns the ContentCommitment field if non-nil, zero value otherwise.

### GetContentCommitmentOk

`func (o *KeyUsage1) GetContentCommitmentOk() (*bool, bool)`

GetContentCommitmentOk returns a tuple with the ContentCommitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCommitment

`func (o *KeyUsage1) SetContentCommitment(v bool)`

SetContentCommitment sets ContentCommitment field to given value.

### HasContentCommitment

`func (o *KeyUsage1) HasContentCommitment() bool`

HasContentCommitment returns a boolean if a field has been set.

### GetKeyEncipherment

`func (o *KeyUsage1) GetKeyEncipherment() bool`

GetKeyEncipherment returns the KeyEncipherment field if non-nil, zero value otherwise.

### GetKeyEnciphermentOk

`func (o *KeyUsage1) GetKeyEnciphermentOk() (*bool, bool)`

GetKeyEnciphermentOk returns a tuple with the KeyEncipherment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyEncipherment

`func (o *KeyUsage1) SetKeyEncipherment(v bool)`

SetKeyEncipherment sets KeyEncipherment field to given value.

### HasKeyEncipherment

`func (o *KeyUsage1) HasKeyEncipherment() bool`

HasKeyEncipherment returns a boolean if a field has been set.

### GetDataEncipherment

`func (o *KeyUsage1) GetDataEncipherment() bool`

GetDataEncipherment returns the DataEncipherment field if non-nil, zero value otherwise.

### GetDataEnciphermentOk

`func (o *KeyUsage1) GetDataEnciphermentOk() (*bool, bool)`

GetDataEnciphermentOk returns a tuple with the DataEncipherment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataEncipherment

`func (o *KeyUsage1) SetDataEncipherment(v bool)`

SetDataEncipherment sets DataEncipherment field to given value.

### HasDataEncipherment

`func (o *KeyUsage1) HasDataEncipherment() bool`

HasDataEncipherment returns a boolean if a field has been set.

### GetKeyAgreement

`func (o *KeyUsage1) GetKeyAgreement() bool`

GetKeyAgreement returns the KeyAgreement field if non-nil, zero value otherwise.

### GetKeyAgreementOk

`func (o *KeyUsage1) GetKeyAgreementOk() (*bool, bool)`

GetKeyAgreementOk returns a tuple with the KeyAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAgreement

`func (o *KeyUsage1) SetKeyAgreement(v bool)`

SetKeyAgreement sets KeyAgreement field to given value.

### HasKeyAgreement

`func (o *KeyUsage1) HasKeyAgreement() bool`

HasKeyAgreement returns a boolean if a field has been set.

### GetKeyCertSign

`func (o *KeyUsage1) GetKeyCertSign() bool`

GetKeyCertSign returns the KeyCertSign field if non-nil, zero value otherwise.

### GetKeyCertSignOk

`func (o *KeyUsage1) GetKeyCertSignOk() (*bool, bool)`

GetKeyCertSignOk returns a tuple with the KeyCertSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyCertSign

`func (o *KeyUsage1) SetKeyCertSign(v bool)`

SetKeyCertSign sets KeyCertSign field to given value.

### HasKeyCertSign

`func (o *KeyUsage1) HasKeyCertSign() bool`

HasKeyCertSign returns a boolean if a field has been set.

### GetCrlSign

`func (o *KeyUsage1) GetCrlSign() bool`

GetCrlSign returns the CrlSign field if non-nil, zero value otherwise.

### GetCrlSignOk

`func (o *KeyUsage1) GetCrlSignOk() (*bool, bool)`

GetCrlSignOk returns a tuple with the CrlSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlSign

`func (o *KeyUsage1) SetCrlSign(v bool)`

SetCrlSign sets CrlSign field to given value.

### HasCrlSign

`func (o *KeyUsage1) HasCrlSign() bool`

HasCrlSign returns a boolean if a field has been set.

### GetEncipherOnly

`func (o *KeyUsage1) GetEncipherOnly() bool`

GetEncipherOnly returns the EncipherOnly field if non-nil, zero value otherwise.

### GetEncipherOnlyOk

`func (o *KeyUsage1) GetEncipherOnlyOk() (*bool, bool)`

GetEncipherOnlyOk returns a tuple with the EncipherOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncipherOnly

`func (o *KeyUsage1) SetEncipherOnly(v bool)`

SetEncipherOnly sets EncipherOnly field to given value.

### HasEncipherOnly

`func (o *KeyUsage1) HasEncipherOnly() bool`

HasEncipherOnly returns a boolean if a field has been set.

### GetDecipherOnly

`func (o *KeyUsage1) GetDecipherOnly() bool`

GetDecipherOnly returns the DecipherOnly field if non-nil, zero value otherwise.

### GetDecipherOnlyOk

`func (o *KeyUsage1) GetDecipherOnlyOk() (*bool, bool)`

GetDecipherOnlyOk returns a tuple with the DecipherOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecipherOnly

`func (o *KeyUsage1) SetDecipherOnly(v bool)`

SetDecipherOnly sets DecipherOnly field to given value.

### HasDecipherOnly

`func (o *KeyUsage1) HasDecipherOnly() bool`

HasDecipherOnly returns a boolean if a field has been set.

### GetExtensionCritical

`func (o *KeyUsage1) GetExtensionCritical() bool`

GetExtensionCritical returns the ExtensionCritical field if non-nil, zero value otherwise.

### GetExtensionCriticalOk

`func (o *KeyUsage1) GetExtensionCriticalOk() (*bool, bool)`

GetExtensionCriticalOk returns a tuple with the ExtensionCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionCritical

`func (o *KeyUsage1) SetExtensionCritical(v bool)`

SetExtensionCritical sets ExtensionCritical field to given value.

### HasExtensionCritical

`func (o *KeyUsage1) HasExtensionCritical() bool`

HasExtensionCritical returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


