# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Portal** | Pointer to **int32** |  | [optional] 
**Initiator** | Pointer to **NullableInt32** |  | [optional] 
**Authmethod** | Pointer to **string** |  | [optional] [default to "NONE"]
**Auth** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortal

`func (o *Group) GetPortal() int32`

GetPortal returns the Portal field if non-nil, zero value otherwise.

### GetPortalOk

`func (o *Group) GetPortalOk() (*int32, bool)`

GetPortalOk returns a tuple with the Portal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortal

`func (o *Group) SetPortal(v int32)`

SetPortal sets Portal field to given value.

### HasPortal

`func (o *Group) HasPortal() bool`

HasPortal returns a boolean if a field has been set.

### GetInitiator

`func (o *Group) GetInitiator() int32`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *Group) GetInitiatorOk() (*int32, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *Group) SetInitiator(v int32)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *Group) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### SetInitiatorNil

`func (o *Group) SetInitiatorNil(b bool)`

 SetInitiatorNil sets the value for Initiator to be an explicit nil

### UnsetInitiator
`func (o *Group) UnsetInitiator()`

UnsetInitiator ensures that no value is present for Initiator, not even an explicit nil
### GetAuthmethod

`func (o *Group) GetAuthmethod() string`

GetAuthmethod returns the Authmethod field if non-nil, zero value otherwise.

### GetAuthmethodOk

`func (o *Group) GetAuthmethodOk() (*string, bool)`

GetAuthmethodOk returns a tuple with the Authmethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthmethod

`func (o *Group) SetAuthmethod(v string)`

SetAuthmethod sets Authmethod field to given value.

### HasAuthmethod

`func (o *Group) HasAuthmethod() bool`

HasAuthmethod returns a boolean if a field has been set.

### GetAuth

`func (o *Group) GetAuth() int32`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *Group) GetAuthOk() (*int32, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *Group) SetAuth(v int32)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *Group) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### SetAuthNil

`func (o *Group) SetAuthNil(b bool)`

 SetAuthNil sets the value for Auth to be an explicit nil

### UnsetAuth
`func (o *Group) UnsetAuth()`

UnsetAuth ensures that no value is present for Auth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


