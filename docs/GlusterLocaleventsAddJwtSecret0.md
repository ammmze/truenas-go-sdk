# GlusterLocaleventsAddJwtSecret0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** | Add a &#x60;secret&#x60; key used to encode/decode JWT messages for sending/receiving gluster events. &#x60;secret&#x60; String representing the key to be used             to encode/decode JWT messages | [optional] 
**Force** | Pointer to **bool** | &#x60;force&#x60; Boolean if set to True, will forcefully             wipe any existing jwt key for this             peer. Note, if forcefully adding a             new key, the other peers in the TSP             will also need to be sent this key. | [optional] [default to false]

## Methods

### NewGlusterLocaleventsAddJwtSecret0

`func NewGlusterLocaleventsAddJwtSecret0() *GlusterLocaleventsAddJwtSecret0`

NewGlusterLocaleventsAddJwtSecret0 instantiates a new GlusterLocaleventsAddJwtSecret0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlusterLocaleventsAddJwtSecret0WithDefaults

`func NewGlusterLocaleventsAddJwtSecret0WithDefaults() *GlusterLocaleventsAddJwtSecret0`

NewGlusterLocaleventsAddJwtSecret0WithDefaults instantiates a new GlusterLocaleventsAddJwtSecret0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *GlusterLocaleventsAddJwtSecret0) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *GlusterLocaleventsAddJwtSecret0) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *GlusterLocaleventsAddJwtSecret0) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *GlusterLocaleventsAddJwtSecret0) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetForce

`func (o *GlusterLocaleventsAddJwtSecret0) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *GlusterLocaleventsAddJwtSecret0) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *GlusterLocaleventsAddJwtSecret0) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *GlusterLocaleventsAddJwtSecret0) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


