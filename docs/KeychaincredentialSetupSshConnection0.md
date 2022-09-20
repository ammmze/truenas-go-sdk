# KeychaincredentialSetupSshConnection0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | Pointer to [**PrivateKey**](PrivateKey.md) |  | [optional] 
**ConnectionName** | Pointer to **string** |  | [optional] 
**SetupType** | Pointer to **string** | 1) Generating SSH Key Pair if required 2) Setting up SSH Credentials based on &#x60;setup_type&#x60; | [optional] [default to "MANUAL"]
**SemiAutomaticSetup** | Pointer to [**KeychainRemoteSshSemiautomaticSetup**](KeychainRemoteSshSemiautomaticSetup.md) |  | [optional] 
**ManualSetup** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewKeychaincredentialSetupSshConnection0

`func NewKeychaincredentialSetupSshConnection0() *KeychaincredentialSetupSshConnection0`

NewKeychaincredentialSetupSshConnection0 instantiates a new KeychaincredentialSetupSshConnection0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeychaincredentialSetupSshConnection0WithDefaults

`func NewKeychaincredentialSetupSshConnection0WithDefaults() *KeychaincredentialSetupSshConnection0`

NewKeychaincredentialSetupSshConnection0WithDefaults instantiates a new KeychaincredentialSetupSshConnection0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *KeychaincredentialSetupSshConnection0) GetPrivateKey() PrivateKey`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *KeychaincredentialSetupSshConnection0) GetPrivateKeyOk() (*PrivateKey, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *KeychaincredentialSetupSshConnection0) SetPrivateKey(v PrivateKey)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *KeychaincredentialSetupSshConnection0) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetConnectionName

`func (o *KeychaincredentialSetupSshConnection0) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *KeychaincredentialSetupSshConnection0) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *KeychaincredentialSetupSshConnection0) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *KeychaincredentialSetupSshConnection0) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetSetupType

`func (o *KeychaincredentialSetupSshConnection0) GetSetupType() string`

GetSetupType returns the SetupType field if non-nil, zero value otherwise.

### GetSetupTypeOk

`func (o *KeychaincredentialSetupSshConnection0) GetSetupTypeOk() (*string, bool)`

GetSetupTypeOk returns a tuple with the SetupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupType

`func (o *KeychaincredentialSetupSshConnection0) SetSetupType(v string)`

SetSetupType sets SetupType field to given value.

### HasSetupType

`func (o *KeychaincredentialSetupSshConnection0) HasSetupType() bool`

HasSetupType returns a boolean if a field has been set.

### GetSemiAutomaticSetup

`func (o *KeychaincredentialSetupSshConnection0) GetSemiAutomaticSetup() KeychainRemoteSshSemiautomaticSetup`

GetSemiAutomaticSetup returns the SemiAutomaticSetup field if non-nil, zero value otherwise.

### GetSemiAutomaticSetupOk

`func (o *KeychaincredentialSetupSshConnection0) GetSemiAutomaticSetupOk() (*KeychainRemoteSshSemiautomaticSetup, bool)`

GetSemiAutomaticSetupOk returns a tuple with the SemiAutomaticSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemiAutomaticSetup

`func (o *KeychaincredentialSetupSshConnection0) SetSemiAutomaticSetup(v KeychainRemoteSshSemiautomaticSetup)`

SetSemiAutomaticSetup sets SemiAutomaticSetup field to given value.

### HasSemiAutomaticSetup

`func (o *KeychaincredentialSetupSshConnection0) HasSemiAutomaticSetup() bool`

HasSemiAutomaticSetup returns a boolean if a field has been set.

### GetManualSetup

`func (o *KeychaincredentialSetupSshConnection0) GetManualSetup() map[string]interface{}`

GetManualSetup returns the ManualSetup field if non-nil, zero value otherwise.

### GetManualSetupOk

`func (o *KeychaincredentialSetupSshConnection0) GetManualSetupOk() (*map[string]interface{}, bool)`

GetManualSetupOk returns a tuple with the ManualSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualSetup

`func (o *KeychaincredentialSetupSshConnection0) SetManualSetup(v map[string]interface{})`

SetManualSetup sets ManualSetup field to given value.

### HasManualSetup

`func (o *KeychaincredentialSetupSshConnection0) HasManualSetup() bool`

HasManualSetup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


