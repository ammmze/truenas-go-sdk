# KeychaincredentialRemoteSshSemiautomaticSetup0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Perform semi-automatic SSH connection setup with other FreeNAS machine. It creates a &#x60;SSH_CREDENTIALS&#x60; credential with specified &#x60;name&#x60; that can be used to connect to FreeNAS machine with specified &#x60;url&#x60; and temporary auth &#x60;token&#x60;. Other FreeNAS machine adds &#x60;private_key&#x60; to allowed &#x60;username&#x60;&#39;s private keys. Other | [optional] 
**Url** | Pointer to **string** | Perform semi-automatic SSH connection setup with other FreeNAS machine. It creates a &#x60;SSH_CREDENTIALS&#x60; credential with specified &#x60;name&#x60; that can be used to connect to FreeNAS machine with specified &#x60;url&#x60; and temporary auth &#x60;token&#x60;. Other FreeNAS machine adds &#x60;private_key&#x60; to allowed &#x60;username&#x60;&#39;s private keys. Other | [optional] 
**Token** | Pointer to **string** | Perform semi-automatic SSH connection setup with other FreeNAS machine. It creates a &#x60;SSH_CREDENTIALS&#x60; credential with specified &#x60;name&#x60; that can be used to connect to FreeNAS machine with specified &#x60;url&#x60; and temporary auth &#x60;token&#x60;. Other FreeNAS machine adds &#x60;private_key&#x60; to allowed &#x60;username&#x60;&#39;s private keys. Other | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** | Perform semi-automatic SSH connection setup with other FreeNAS machine. It creates a &#x60;SSH_CREDENTIALS&#x60; credential with specified &#x60;name&#x60; that can be used to connect to FreeNAS machine with specified &#x60;url&#x60; and temporary auth &#x60;token&#x60;. Other FreeNAS machine adds &#x60;private_key&#x60; to allowed &#x60;username&#x60;&#39;s private keys. Other | [optional] [default to "root"]
**PrivateKey** | Pointer to **int32** | Perform semi-automatic SSH connection setup with other FreeNAS machine. It creates a &#x60;SSH_CREDENTIALS&#x60; credential with specified &#x60;name&#x60; that can be used to connect to FreeNAS machine with specified &#x60;url&#x60; and temporary auth &#x60;token&#x60;. Other FreeNAS machine adds &#x60;private_key&#x60; to allowed &#x60;username&#x60;&#39;s private keys. Other | [optional] 
**Cipher** | Pointer to **string** | &#x60;SSH_CREDENTIALS&#x60; attributes such as &#x60;cipher&#x60; and &#x60;connect_timeout&#x60; can be specified as well. | [optional] [default to "STANDARD"]
**ConnectTimeout** | Pointer to **int32** | &#x60;SSH_CREDENTIALS&#x60; attributes such as &#x60;cipher&#x60; and &#x60;connect_timeout&#x60; can be specified as well. | [optional] [default to 10]

## Methods

### NewKeychaincredentialRemoteSshSemiautomaticSetup0

`func NewKeychaincredentialRemoteSshSemiautomaticSetup0() *KeychaincredentialRemoteSshSemiautomaticSetup0`

NewKeychaincredentialRemoteSshSemiautomaticSetup0 instantiates a new KeychaincredentialRemoteSshSemiautomaticSetup0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeychaincredentialRemoteSshSemiautomaticSetup0WithDefaults

`func NewKeychaincredentialRemoteSshSemiautomaticSetup0WithDefaults() *KeychaincredentialRemoteSshSemiautomaticSetup0`

NewKeychaincredentialRemoteSshSemiautomaticSetup0WithDefaults instantiates a new KeychaincredentialRemoteSshSemiautomaticSetup0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetToken

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetPassword

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPrivateKey

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) GetPrivateKey() int32`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) GetPrivateKeyOk() (*int32, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) SetPrivateKey(v int32)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetCipher

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) GetCipher() string`

GetCipher returns the Cipher field if non-nil, zero value otherwise.

### GetCipherOk

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) GetCipherOk() (*string, bool)`

GetCipherOk returns a tuple with the Cipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipher

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) SetCipher(v string)`

SetCipher sets Cipher field to given value.

### HasCipher

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) HasCipher() bool`

HasCipher returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) GetConnectTimeout() int32`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) GetConnectTimeoutOk() (*int32, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) SetConnectTimeout(v int32)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *KeychaincredentialRemoteSshSemiautomaticSetup0) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


