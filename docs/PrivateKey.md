# PrivateKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenerateKey** | Pointer to **bool** |  | [optional] [default to true]
**ExistingKeyId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewPrivateKey

`func NewPrivateKey() *PrivateKey`

NewPrivateKey instantiates a new PrivateKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateKeyWithDefaults

`func NewPrivateKeyWithDefaults() *PrivateKey`

NewPrivateKeyWithDefaults instantiates a new PrivateKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerateKey

`func (o *PrivateKey) GetGenerateKey() bool`

GetGenerateKey returns the GenerateKey field if non-nil, zero value otherwise.

### GetGenerateKeyOk

`func (o *PrivateKey) GetGenerateKeyOk() (*bool, bool)`

GetGenerateKeyOk returns a tuple with the GenerateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateKey

`func (o *PrivateKey) SetGenerateKey(v bool)`

SetGenerateKey sets GenerateKey field to given value.

### HasGenerateKey

`func (o *PrivateKey) HasGenerateKey() bool`

HasGenerateKey returns a boolean if a field has been set.

### GetExistingKeyId

`func (o *PrivateKey) GetExistingKeyId() int32`

GetExistingKeyId returns the ExistingKeyId field if non-nil, zero value otherwise.

### GetExistingKeyIdOk

`func (o *PrivateKey) GetExistingKeyIdOk() (*int32, bool)`

GetExistingKeyIdOk returns a tuple with the ExistingKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingKeyId

`func (o *PrivateKey) SetExistingKeyId(v int32)`

SetExistingKeyId sets ExistingKeyId field to given value.

### HasExistingKeyId

`func (o *PrivateKey) HasExistingKeyId() bool`

HasExistingKeyId returns a boolean if a field has been set.

### GetName

`func (o *PrivateKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivateKey) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


