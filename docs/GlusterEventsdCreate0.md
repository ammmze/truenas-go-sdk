# GlusterEventsdCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Add &#x60;url&#x60; webhook that will be called with a JSON formatted POST request that will include the event that was triggered along with the relevant data. &#x60;url&#x60; is a http address (i.e. http://192.168.1.50/endpoint) | [optional] 
**BearerToken** | Pointer to **string** | &#x60;bearer_token&#x60; is a bearer token | [optional] 
**Secret** | Pointer to **string** | &#x60;secret&#x60; secret to encode the JWT message | [optional] 

## Methods

### NewGlusterEventsdCreate0

`func NewGlusterEventsdCreate0() *GlusterEventsdCreate0`

NewGlusterEventsdCreate0 instantiates a new GlusterEventsdCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlusterEventsdCreate0WithDefaults

`func NewGlusterEventsdCreate0WithDefaults() *GlusterEventsdCreate0`

NewGlusterEventsdCreate0WithDefaults instantiates a new GlusterEventsdCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *GlusterEventsdCreate0) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GlusterEventsdCreate0) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GlusterEventsdCreate0) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GlusterEventsdCreate0) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetBearerToken

`func (o *GlusterEventsdCreate0) GetBearerToken() string`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *GlusterEventsdCreate0) GetBearerTokenOk() (*string, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *GlusterEventsdCreate0) SetBearerToken(v string)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *GlusterEventsdCreate0) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### GetSecret

`func (o *GlusterEventsdCreate0) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *GlusterEventsdCreate0) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *GlusterEventsdCreate0) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *GlusterEventsdCreate0) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


