# Nfs4Ace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Perms** | Pointer to [**Perms**](Perms.md) |  | [optional] 
**Flags** | Pointer to [**Flags**](Flags.md) |  | [optional] 

## Methods

### NewNfs4Ace

`func NewNfs4Ace() *Nfs4Ace`

NewNfs4Ace instantiates a new Nfs4Ace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfs4AceWithDefaults

`func NewNfs4AceWithDefaults() *Nfs4Ace`

NewNfs4AceWithDefaults instantiates a new Nfs4Ace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *Nfs4Ace) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Nfs4Ace) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Nfs4Ace) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Nfs4Ace) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetId

`func (o *Nfs4Ace) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Nfs4Ace) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Nfs4Ace) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Nfs4Ace) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Nfs4Ace) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Nfs4Ace) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *Nfs4Ace) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Nfs4Ace) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Nfs4Ace) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Nfs4Ace) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPerms

`func (o *Nfs4Ace) GetPerms() Perms`

GetPerms returns the Perms field if non-nil, zero value otherwise.

### GetPermsOk

`func (o *Nfs4Ace) GetPermsOk() (*Perms, bool)`

GetPermsOk returns a tuple with the Perms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerms

`func (o *Nfs4Ace) SetPerms(v Perms)`

SetPerms sets Perms field to given value.

### HasPerms

`func (o *Nfs4Ace) HasPerms() bool`

HasPerms returns a boolean if a field has been set.

### GetFlags

`func (o *Nfs4Ace) GetFlags() Flags`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Nfs4Ace) GetFlagsOk() (*Flags, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *Nfs4Ace) SetFlags(v Flags)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *Nfs4Ace) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


