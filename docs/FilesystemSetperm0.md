# FilesystemSetperm0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Set unix permissions on given &#x60;path&#x60;. &#x60;stripacl&#x60; setperm will fail if an extended ACL is present on &#x60;path&#x60;, unless &#x60;stripacl&#x60; is set to True. | [optional] 
**Mode** | Pointer to **NullableString** | If &#x60;mode&#x60; is specified then the mode will be applied to the path and files and subdirectories depending on which &#x60;options&#x60; are selected. Mode should be formatted as string representation of octal permissions bits. | [optional] 
**Uid** | Pointer to **NullableInt32** | &#x60;uid&#x60; the desired UID of the file user. If set to None (the default), then user is not changed. | [optional] 
**Gid** | Pointer to **NullableInt32** | &#x60;gid&#x60; the desired GID of the file group. If set to None (the default), then group is not changed. | [optional] 
**Options** | Pointer to [**Options2**](Options2.md) |  | [optional] 

## Methods

### NewFilesystemSetperm0

`func NewFilesystemSetperm0() *FilesystemSetperm0`

NewFilesystemSetperm0 instantiates a new FilesystemSetperm0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemSetperm0WithDefaults

`func NewFilesystemSetperm0WithDefaults() *FilesystemSetperm0`

NewFilesystemSetperm0WithDefaults instantiates a new FilesystemSetperm0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *FilesystemSetperm0) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FilesystemSetperm0) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FilesystemSetperm0) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FilesystemSetperm0) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMode

`func (o *FilesystemSetperm0) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FilesystemSetperm0) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FilesystemSetperm0) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FilesystemSetperm0) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *FilesystemSetperm0) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *FilesystemSetperm0) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetUid

`func (o *FilesystemSetperm0) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *FilesystemSetperm0) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *FilesystemSetperm0) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *FilesystemSetperm0) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *FilesystemSetperm0) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *FilesystemSetperm0) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetGid

`func (o *FilesystemSetperm0) GetGid() int32`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *FilesystemSetperm0) GetGidOk() (*int32, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *FilesystemSetperm0) SetGid(v int32)`

SetGid sets Gid field to given value.

### HasGid

`func (o *FilesystemSetperm0) HasGid() bool`

HasGid returns a boolean if a field has been set.

### SetGidNil

`func (o *FilesystemSetperm0) SetGidNil(b bool)`

 SetGidNil sets the value for Gid to be an explicit nil

### UnsetGid
`func (o *FilesystemSetperm0) UnsetGid()`

UnsetGid ensures that no value is present for Gid, not even an explicit nil
### GetOptions

`func (o *FilesystemSetperm0) GetOptions() Options2`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FilesystemSetperm0) GetOptionsOk() (*Options2, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FilesystemSetperm0) SetOptions(v Options2)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *FilesystemSetperm0) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


