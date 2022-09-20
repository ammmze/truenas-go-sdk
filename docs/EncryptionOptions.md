# EncryptionOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenerateKey** | Pointer to **bool** |  | [optional] [default to false]
**Pbkdf2iters** | Pointer to **int32** |  | [optional] [default to 350000]
**Algorithm** | Pointer to **string** |  | [optional] [default to "AES-256-GCM"]
**Passphrase** | Pointer to **NullableString** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEncryptionOptions

`func NewEncryptionOptions() *EncryptionOptions`

NewEncryptionOptions instantiates a new EncryptionOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionOptionsWithDefaults

`func NewEncryptionOptionsWithDefaults() *EncryptionOptions`

NewEncryptionOptionsWithDefaults instantiates a new EncryptionOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerateKey

`func (o *EncryptionOptions) GetGenerateKey() bool`

GetGenerateKey returns the GenerateKey field if non-nil, zero value otherwise.

### GetGenerateKeyOk

`func (o *EncryptionOptions) GetGenerateKeyOk() (*bool, bool)`

GetGenerateKeyOk returns a tuple with the GenerateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateKey

`func (o *EncryptionOptions) SetGenerateKey(v bool)`

SetGenerateKey sets GenerateKey field to given value.

### HasGenerateKey

`func (o *EncryptionOptions) HasGenerateKey() bool`

HasGenerateKey returns a boolean if a field has been set.

### GetPbkdf2iters

`func (o *EncryptionOptions) GetPbkdf2iters() int32`

GetPbkdf2iters returns the Pbkdf2iters field if non-nil, zero value otherwise.

### GetPbkdf2itersOk

`func (o *EncryptionOptions) GetPbkdf2itersOk() (*int32, bool)`

GetPbkdf2itersOk returns a tuple with the Pbkdf2iters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbkdf2iters

`func (o *EncryptionOptions) SetPbkdf2iters(v int32)`

SetPbkdf2iters sets Pbkdf2iters field to given value.

### HasPbkdf2iters

`func (o *EncryptionOptions) HasPbkdf2iters() bool`

HasPbkdf2iters returns a boolean if a field has been set.

### GetAlgorithm

`func (o *EncryptionOptions) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *EncryptionOptions) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *EncryptionOptions) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *EncryptionOptions) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetPassphrase

`func (o *EncryptionOptions) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *EncryptionOptions) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *EncryptionOptions) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *EncryptionOptions) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### SetPassphraseNil

`func (o *EncryptionOptions) SetPassphraseNil(b bool)`

 SetPassphraseNil sets the value for Passphrase to be an explicit nil

### UnsetPassphrase
`func (o *EncryptionOptions) UnsetPassphrase()`

UnsetPassphrase ensures that no value is present for Passphrase, not even an explicit nil
### GetKey

`func (o *EncryptionOptions) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EncryptionOptions) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EncryptionOptions) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EncryptionOptions) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *EncryptionOptions) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *EncryptionOptions) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


