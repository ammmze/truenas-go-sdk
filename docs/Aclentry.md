# Aclentry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AeWhoSid** | Pointer to **string** |  | [optional] 
**AeWhoName** | Pointer to [**AeWhoName**](AeWhoName.md) |  | [optional] 
**AePerm** | Pointer to **string** |  | [optional] 
**AeType** | Pointer to **string** |  | [optional] 

## Methods

### NewAclentry

`func NewAclentry() *Aclentry`

NewAclentry instantiates a new Aclentry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAclentryWithDefaults

`func NewAclentryWithDefaults() *Aclentry`

NewAclentryWithDefaults instantiates a new Aclentry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAeWhoSid

`func (o *Aclentry) GetAeWhoSid() string`

GetAeWhoSid returns the AeWhoSid field if non-nil, zero value otherwise.

### GetAeWhoSidOk

`func (o *Aclentry) GetAeWhoSidOk() (*string, bool)`

GetAeWhoSidOk returns a tuple with the AeWhoSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeWhoSid

`func (o *Aclentry) SetAeWhoSid(v string)`

SetAeWhoSid sets AeWhoSid field to given value.

### HasAeWhoSid

`func (o *Aclentry) HasAeWhoSid() bool`

HasAeWhoSid returns a boolean if a field has been set.

### GetAeWhoName

`func (o *Aclentry) GetAeWhoName() AeWhoName`

GetAeWhoName returns the AeWhoName field if non-nil, zero value otherwise.

### GetAeWhoNameOk

`func (o *Aclentry) GetAeWhoNameOk() (*AeWhoName, bool)`

GetAeWhoNameOk returns a tuple with the AeWhoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeWhoName

`func (o *Aclentry) SetAeWhoName(v AeWhoName)`

SetAeWhoName sets AeWhoName field to given value.

### HasAeWhoName

`func (o *Aclentry) HasAeWhoName() bool`

HasAeWhoName returns a boolean if a field has been set.

### GetAePerm

`func (o *Aclentry) GetAePerm() string`

GetAePerm returns the AePerm field if non-nil, zero value otherwise.

### GetAePermOk

`func (o *Aclentry) GetAePermOk() (*string, bool)`

GetAePermOk returns a tuple with the AePerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAePerm

`func (o *Aclentry) SetAePerm(v string)`

SetAePerm sets AePerm field to given value.

### HasAePerm

`func (o *Aclentry) HasAePerm() bool`

HasAePerm returns a boolean if a field has been set.

### GetAeType

`func (o *Aclentry) GetAeType() string`

GetAeType returns the AeType field if non-nil, zero value otherwise.

### GetAeTypeOk

`func (o *Aclentry) GetAeTypeOk() (*string, bool)`

GetAeTypeOk returns a tuple with the AeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeType

`func (o *Aclentry) SetAeType(v string)`

SetAeType sets AeType field to given value.

### HasAeType

`func (o *Aclentry) HasAeType() bool`

HasAeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


