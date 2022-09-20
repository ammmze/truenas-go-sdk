# SmbSharesecCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShareName** | Pointer to **string** | &#x60;share_name&#x60; - name of SMB share. | [optional] 
**ShareAcl** | Pointer to [**[]Aclentry**](Aclentry.md) | &#x60;share_acl&#x60; a list of ACL entries (dictionaries) with the following keys: | [optional] [default to [{"ae_who_sid":"S-1-1-0","ae_perm":"FULL","ae_type":"ALLOWED"}]]

## Methods

### NewSmbSharesecCreate0

`func NewSmbSharesecCreate0() *SmbSharesecCreate0`

NewSmbSharesecCreate0 instantiates a new SmbSharesecCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbSharesecCreate0WithDefaults

`func NewSmbSharesecCreate0WithDefaults() *SmbSharesecCreate0`

NewSmbSharesecCreate0WithDefaults instantiates a new SmbSharesecCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShareName

`func (o *SmbSharesecCreate0) GetShareName() string`

GetShareName returns the ShareName field if non-nil, zero value otherwise.

### GetShareNameOk

`func (o *SmbSharesecCreate0) GetShareNameOk() (*string, bool)`

GetShareNameOk returns a tuple with the ShareName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareName

`func (o *SmbSharesecCreate0) SetShareName(v string)`

SetShareName sets ShareName field to given value.

### HasShareName

`func (o *SmbSharesecCreate0) HasShareName() bool`

HasShareName returns a boolean if a field has been set.

### GetShareAcl

`func (o *SmbSharesecCreate0) GetShareAcl() []Aclentry`

GetShareAcl returns the ShareAcl field if non-nil, zero value otherwise.

### GetShareAclOk

`func (o *SmbSharesecCreate0) GetShareAclOk() (*[]Aclentry, bool)`

GetShareAclOk returns a tuple with the ShareAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareAcl

`func (o *SmbSharesecCreate0) SetShareAcl(v []Aclentry)`

SetShareAcl sets ShareAcl field to given value.

### HasShareAcl

`func (o *SmbSharesecCreate0) HasShareAcl() bool`

HasShareAcl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


