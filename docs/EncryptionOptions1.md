# EncryptionOptions1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenerateKey** | Pointer to **bool** |  | [optional] [default to false]
**Pbkdf2iters** | Pointer to **int32** |  | [optional] [default to 350000]
**Algorithm** | Pointer to **string** |  | [optional] [default to "AES-256-GCM"]
**Passphrase** | Pointer to **NullableString** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEncryptionOptions1

`func NewEncryptionOptions1() *EncryptionOptions1`

NewEncryptionOptions1 instantiates a new EncryptionOptions1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionOptions1WithDefaults

`func NewEncryptionOptions1WithDefaults() *EncryptionOptions1`

NewEncryptionOptions1WithDefaults instantiates a new EncryptionOptions1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerateKey

`func (o *EncryptionOptions1) GetGenerateKey() bool`

GetGenerateKey returns the GenerateKey field if non-nil, zero value otherwise.

### GetGenerateKeyOk

`func (o *EncryptionOptions1) GetGenerateKeyOk() (*bool, bool)`

GetGenerateKeyOk returns a tuple with the GenerateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateKey

`func (o *EncryptionOptions1) SetGenerateKey(v bool)`

SetGenerateKey sets GenerateKey field to given value.

### HasGenerateKey

`func (o *EncryptionOptions1) HasGenerateKey() bool`

HasGenerateKey returns a boolean if a field has been set.

### GetPbkdf2iters

`func (o *EncryptionOptions1) GetPbkdf2iters() int32`

GetPbkdf2iters returns the Pbkdf2iters field if non-nil, zero value otherwise.

### GetPbkdf2itersOk

`func (o *EncryptionOptions1) GetPbkdf2itersOk() (*int32, bool)`

GetPbkdf2itersOk returns a tuple with the Pbkdf2iters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbkdf2iters

`func (o *EncryptionOptions1) SetPbkdf2iters(v int32)`

SetPbkdf2iters sets Pbkdf2iters field to given value.

### HasPbkdf2iters

`func (o *EncryptionOptions1) HasPbkdf2iters() bool`

HasPbkdf2iters returns a boolean if a field has been set.

### GetAlgorithm

`func (o *EncryptionOptions1) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *EncryptionOptions1) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *EncryptionOptions1) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *EncryptionOptions1) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetPassphrase

`func (o *EncryptionOptions1) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *EncryptionOptions1) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *EncryptionOptions1) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *EncryptionOptions1) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### SetPassphraseNil

`func (o *EncryptionOptions1) SetPassphraseNil(b bool)`

 SetPassphraseNil sets the value for Passphrase to be an explicit nil

### UnsetPassphrase
`func (o *EncryptionOptions1) UnsetPassphrase()`

UnsetPassphrase ensures that no value is present for Passphrase, not even an explicit nil
### GetKey

`func (o *EncryptionOptions1) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EncryptionOptions1) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EncryptionOptions1) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EncryptionOptions1) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *EncryptionOptions1) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *EncryptionOptions1) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


