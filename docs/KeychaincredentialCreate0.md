# KeychaincredentialCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Every Keychain Credential has a &#x60;name&#x60; which is used to distinguish it from others. | [optional] 
**Type** | Pointer to **string** | The following &#x60;type&#x60;s are supported:   * &#x60;SSH_KEY_PAIR&#x60;    Which &#x60;attributes&#x60; are:    * &#x60;private_key&#x60;    * &#x60;public_key&#x60; (which can be omitted and thus automatically derived from private key)    At least one attribute is required. | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | The following &#x60;type&#x60;s are supported:   * &#x60;SSH_KEY_PAIR&#x60;    Which &#x60;attributes&#x60; are:    * &#x60;private_key&#x60;    * &#x60;public_key&#x60; (which can be omitted and thus automatically derived from private key)    At least one attribute is required.  * &#x60;SSH_CREDENTIALS&#x60;    Which &#x60;attributes&#x60; are:    * &#x60;host&#x60;    * &#x60;port&#x60; (default 22)    * &#x60;username&#x60; (default root)    * &#x60;private_key&#x60; (Keychain Credential ID)    * &#x60;remote_host_key&#x60; (you can use &#x60;keychaincredential.remote_ssh_host_key_scan&#x60; do discover it)    * &#x60;cipher&#x60;: one of &#x60;STANDARD&#x60;, &#x60;FAST&#x60;, or &#x60;DISABLED&#x60; (last requires special support from both SSH server and      client)    * &#x60;connect_timeout&#x60; (default 10) | [optional] 

## Methods

### NewKeychaincredentialCreate0

`func NewKeychaincredentialCreate0() *KeychaincredentialCreate0`

NewKeychaincredentialCreate0 instantiates a new KeychaincredentialCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeychaincredentialCreate0WithDefaults

`func NewKeychaincredentialCreate0WithDefaults() *KeychaincredentialCreate0`

NewKeychaincredentialCreate0WithDefaults instantiates a new KeychaincredentialCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeychaincredentialCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeychaincredentialCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeychaincredentialCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeychaincredentialCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *KeychaincredentialCreate0) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeychaincredentialCreate0) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeychaincredentialCreate0) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KeychaincredentialCreate0) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *KeychaincredentialCreate0) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *KeychaincredentialCreate0) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *KeychaincredentialCreate0) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *KeychaincredentialCreate0) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


