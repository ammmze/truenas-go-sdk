# Perms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**READ_DATA** | Pointer to **bool** |  | [optional] 
**WRITE_DATA** | Pointer to **bool** |  | [optional] 
**APPEND_DATA** | Pointer to **bool** |  | [optional] 
**READ_NAMED_ATTRS** | Pointer to **bool** |  | [optional] 
**WRITE_NAMED_ATTRS** | Pointer to **bool** |  | [optional] 
**EXECUTE** | Pointer to **bool** |  | [optional] 
**DELETE_CHILD** | Pointer to **bool** |  | [optional] 
**READ_ATTRIBUTES** | Pointer to **bool** |  | [optional] 
**WRITE_ATTRIBUTES** | Pointer to **bool** |  | [optional] 
**DELETE** | Pointer to **bool** |  | [optional] 
**READ_ACL** | Pointer to **bool** |  | [optional] 
**WRITE_ACL** | Pointer to **bool** |  | [optional] 
**WRITE_OWNER** | Pointer to **bool** |  | [optional] 
**SYNCHRONIZE** | Pointer to **bool** |  | [optional] 
**BASIC** | Pointer to **string** |  | [optional] 

## Methods

### NewPerms

`func NewPerms() *Perms`

NewPerms instantiates a new Perms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermsWithDefaults

`func NewPermsWithDefaults() *Perms`

NewPermsWithDefaults instantiates a new Perms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetREAD_DATA

`func (o *Perms) GetREAD_DATA() bool`

GetREAD_DATA returns the READ_DATA field if non-nil, zero value otherwise.

### GetREAD_DATAOk

`func (o *Perms) GetREAD_DATAOk() (*bool, bool)`

GetREAD_DATAOk returns a tuple with the READ_DATA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREAD_DATA

`func (o *Perms) SetREAD_DATA(v bool)`

SetREAD_DATA sets READ_DATA field to given value.

### HasREAD_DATA

`func (o *Perms) HasREAD_DATA() bool`

HasREAD_DATA returns a boolean if a field has been set.

### GetWRITE_DATA

`func (o *Perms) GetWRITE_DATA() bool`

GetWRITE_DATA returns the WRITE_DATA field if non-nil, zero value otherwise.

### GetWRITE_DATAOk

`func (o *Perms) GetWRITE_DATAOk() (*bool, bool)`

GetWRITE_DATAOk returns a tuple with the WRITE_DATA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWRITE_DATA

`func (o *Perms) SetWRITE_DATA(v bool)`

SetWRITE_DATA sets WRITE_DATA field to given value.

### HasWRITE_DATA

`func (o *Perms) HasWRITE_DATA() bool`

HasWRITE_DATA returns a boolean if a field has been set.

### GetAPPEND_DATA

`func (o *Perms) GetAPPEND_DATA() bool`

GetAPPEND_DATA returns the APPEND_DATA field if non-nil, zero value otherwise.

### GetAPPEND_DATAOk

`func (o *Perms) GetAPPEND_DATAOk() (*bool, bool)`

GetAPPEND_DATAOk returns a tuple with the APPEND_DATA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPPEND_DATA

`func (o *Perms) SetAPPEND_DATA(v bool)`

SetAPPEND_DATA sets APPEND_DATA field to given value.

### HasAPPEND_DATA

`func (o *Perms) HasAPPEND_DATA() bool`

HasAPPEND_DATA returns a boolean if a field has been set.

### GetREAD_NAMED_ATTRS

`func (o *Perms) GetREAD_NAMED_ATTRS() bool`

GetREAD_NAMED_ATTRS returns the READ_NAMED_ATTRS field if non-nil, zero value otherwise.

### GetREAD_NAMED_ATTRSOk

`func (o *Perms) GetREAD_NAMED_ATTRSOk() (*bool, bool)`

GetREAD_NAMED_ATTRSOk returns a tuple with the READ_NAMED_ATTRS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREAD_NAMED_ATTRS

`func (o *Perms) SetREAD_NAMED_ATTRS(v bool)`

SetREAD_NAMED_ATTRS sets READ_NAMED_ATTRS field to given value.

### HasREAD_NAMED_ATTRS

`func (o *Perms) HasREAD_NAMED_ATTRS() bool`

HasREAD_NAMED_ATTRS returns a boolean if a field has been set.

### GetWRITE_NAMED_ATTRS

`func (o *Perms) GetWRITE_NAMED_ATTRS() bool`

GetWRITE_NAMED_ATTRS returns the WRITE_NAMED_ATTRS field if non-nil, zero value otherwise.

### GetWRITE_NAMED_ATTRSOk

`func (o *Perms) GetWRITE_NAMED_ATTRSOk() (*bool, bool)`

GetWRITE_NAMED_ATTRSOk returns a tuple with the WRITE_NAMED_ATTRS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWRITE_NAMED_ATTRS

`func (o *Perms) SetWRITE_NAMED_ATTRS(v bool)`

SetWRITE_NAMED_ATTRS sets WRITE_NAMED_ATTRS field to given value.

### HasWRITE_NAMED_ATTRS

`func (o *Perms) HasWRITE_NAMED_ATTRS() bool`

HasWRITE_NAMED_ATTRS returns a boolean if a field has been set.

### GetEXECUTE

`func (o *Perms) GetEXECUTE() bool`

GetEXECUTE returns the EXECUTE field if non-nil, zero value otherwise.

### GetEXECUTEOk

`func (o *Perms) GetEXECUTEOk() (*bool, bool)`

GetEXECUTEOk returns a tuple with the EXECUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEXECUTE

`func (o *Perms) SetEXECUTE(v bool)`

SetEXECUTE sets EXECUTE field to given value.

### HasEXECUTE

`func (o *Perms) HasEXECUTE() bool`

HasEXECUTE returns a boolean if a field has been set.

### GetDELETE_CHILD

`func (o *Perms) GetDELETE_CHILD() bool`

GetDELETE_CHILD returns the DELETE_CHILD field if non-nil, zero value otherwise.

### GetDELETE_CHILDOk

`func (o *Perms) GetDELETE_CHILDOk() (*bool, bool)`

GetDELETE_CHILDOk returns a tuple with the DELETE_CHILD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDELETE_CHILD

`func (o *Perms) SetDELETE_CHILD(v bool)`

SetDELETE_CHILD sets DELETE_CHILD field to given value.

### HasDELETE_CHILD

`func (o *Perms) HasDELETE_CHILD() bool`

HasDELETE_CHILD returns a boolean if a field has been set.

### GetREAD_ATTRIBUTES

`func (o *Perms) GetREAD_ATTRIBUTES() bool`

GetREAD_ATTRIBUTES returns the READ_ATTRIBUTES field if non-nil, zero value otherwise.

### GetREAD_ATTRIBUTESOk

`func (o *Perms) GetREAD_ATTRIBUTESOk() (*bool, bool)`

GetREAD_ATTRIBUTESOk returns a tuple with the READ_ATTRIBUTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREAD_ATTRIBUTES

`func (o *Perms) SetREAD_ATTRIBUTES(v bool)`

SetREAD_ATTRIBUTES sets READ_ATTRIBUTES field to given value.

### HasREAD_ATTRIBUTES

`func (o *Perms) HasREAD_ATTRIBUTES() bool`

HasREAD_ATTRIBUTES returns a boolean if a field has been set.

### GetWRITE_ATTRIBUTES

`func (o *Perms) GetWRITE_ATTRIBUTES() bool`

GetWRITE_ATTRIBUTES returns the WRITE_ATTRIBUTES field if non-nil, zero value otherwise.

### GetWRITE_ATTRIBUTESOk

`func (o *Perms) GetWRITE_ATTRIBUTESOk() (*bool, bool)`

GetWRITE_ATTRIBUTESOk returns a tuple with the WRITE_ATTRIBUTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWRITE_ATTRIBUTES

`func (o *Perms) SetWRITE_ATTRIBUTES(v bool)`

SetWRITE_ATTRIBUTES sets WRITE_ATTRIBUTES field to given value.

### HasWRITE_ATTRIBUTES

`func (o *Perms) HasWRITE_ATTRIBUTES() bool`

HasWRITE_ATTRIBUTES returns a boolean if a field has been set.

### GetDELETE

`func (o *Perms) GetDELETE() bool`

GetDELETE returns the DELETE field if non-nil, zero value otherwise.

### GetDELETEOk

`func (o *Perms) GetDELETEOk() (*bool, bool)`

GetDELETEOk returns a tuple with the DELETE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDELETE

`func (o *Perms) SetDELETE(v bool)`

SetDELETE sets DELETE field to given value.

### HasDELETE

`func (o *Perms) HasDELETE() bool`

HasDELETE returns a boolean if a field has been set.

### GetREAD_ACL

`func (o *Perms) GetREAD_ACL() bool`

GetREAD_ACL returns the READ_ACL field if non-nil, zero value otherwise.

### GetREAD_ACLOk

`func (o *Perms) GetREAD_ACLOk() (*bool, bool)`

GetREAD_ACLOk returns a tuple with the READ_ACL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREAD_ACL

`func (o *Perms) SetREAD_ACL(v bool)`

SetREAD_ACL sets READ_ACL field to given value.

### HasREAD_ACL

`func (o *Perms) HasREAD_ACL() bool`

HasREAD_ACL returns a boolean if a field has been set.

### GetWRITE_ACL

`func (o *Perms) GetWRITE_ACL() bool`

GetWRITE_ACL returns the WRITE_ACL field if non-nil, zero value otherwise.

### GetWRITE_ACLOk

`func (o *Perms) GetWRITE_ACLOk() (*bool, bool)`

GetWRITE_ACLOk returns a tuple with the WRITE_ACL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWRITE_ACL

`func (o *Perms) SetWRITE_ACL(v bool)`

SetWRITE_ACL sets WRITE_ACL field to given value.

### HasWRITE_ACL

`func (o *Perms) HasWRITE_ACL() bool`

HasWRITE_ACL returns a boolean if a field has been set.

### GetWRITE_OWNER

`func (o *Perms) GetWRITE_OWNER() bool`

GetWRITE_OWNER returns the WRITE_OWNER field if non-nil, zero value otherwise.

### GetWRITE_OWNEROk

`func (o *Perms) GetWRITE_OWNEROk() (*bool, bool)`

GetWRITE_OWNEROk returns a tuple with the WRITE_OWNER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWRITE_OWNER

`func (o *Perms) SetWRITE_OWNER(v bool)`

SetWRITE_OWNER sets WRITE_OWNER field to given value.

### HasWRITE_OWNER

`func (o *Perms) HasWRITE_OWNER() bool`

HasWRITE_OWNER returns a boolean if a field has been set.

### GetSYNCHRONIZE

`func (o *Perms) GetSYNCHRONIZE() bool`

GetSYNCHRONIZE returns the SYNCHRONIZE field if non-nil, zero value otherwise.

### GetSYNCHRONIZEOk

`func (o *Perms) GetSYNCHRONIZEOk() (*bool, bool)`

GetSYNCHRONIZEOk returns a tuple with the SYNCHRONIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSYNCHRONIZE

`func (o *Perms) SetSYNCHRONIZE(v bool)`

SetSYNCHRONIZE sets SYNCHRONIZE field to given value.

### HasSYNCHRONIZE

`func (o *Perms) HasSYNCHRONIZE() bool`

HasSYNCHRONIZE returns a boolean if a field has been set.

### GetBASIC

`func (o *Perms) GetBASIC() string`

GetBASIC returns the BASIC field if non-nil, zero value otherwise.

### GetBASICOk

`func (o *Perms) GetBASICOk() (*string, bool)`

GetBASICOk returns a tuple with the BASIC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASIC

`func (o *Perms) SetBASIC(v string)`

SetBASIC sets BASIC field to given value.

### HasBASIC

`func (o *Perms) HasBASIC() bool`

HasBASIC returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


