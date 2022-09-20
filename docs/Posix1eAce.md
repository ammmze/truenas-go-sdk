# Posix1eAce

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **bool** |  | [optional] [default to false]
**Tag** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [default to -1]
**Perms** | Pointer to [**Perms1**](Perms1.md) |  | [optional] 

## Methods

### NewPosix1eAce

`func NewPosix1eAce() *Posix1eAce`

NewPosix1eAce instantiates a new Posix1eAce object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPosix1eAceWithDefaults

`func NewPosix1eAceWithDefaults() *Posix1eAce`

NewPosix1eAceWithDefaults instantiates a new Posix1eAce object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *Posix1eAce) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Posix1eAce) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Posix1eAce) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Posix1eAce) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetTag

`func (o *Posix1eAce) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Posix1eAce) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Posix1eAce) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Posix1eAce) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetId

`func (o *Posix1eAce) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Posix1eAce) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Posix1eAce) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Posix1eAce) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPerms

`func (o *Posix1eAce) GetPerms() Perms1`

GetPerms returns the Perms field if non-nil, zero value otherwise.

### GetPermsOk

`func (o *Posix1eAce) GetPermsOk() (*Perms1, bool)`

GetPermsOk returns a tuple with the Perms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerms

`func (o *Posix1eAce) SetPerms(v Perms1)`

SetPerms sets Perms field to given value.

### HasPerms

`func (o *Posix1eAce) HasPerms() bool`

HasPerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


