# FilesystemSetImmutable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SetFlag** | Pointer to **bool** |  | [optional] 
**Path** | Pointer to **string** | Set/Unset immutable flag at &#x60;path&#x60;. | [optional] 

## Methods

### NewFilesystemSetImmutable

`func NewFilesystemSetImmutable() *FilesystemSetImmutable`

NewFilesystemSetImmutable instantiates a new FilesystemSetImmutable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemSetImmutableWithDefaults

`func NewFilesystemSetImmutableWithDefaults() *FilesystemSetImmutable`

NewFilesystemSetImmutableWithDefaults instantiates a new FilesystemSetImmutable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSetFlag

`func (o *FilesystemSetImmutable) GetSetFlag() bool`

GetSetFlag returns the SetFlag field if non-nil, zero value otherwise.

### GetSetFlagOk

`func (o *FilesystemSetImmutable) GetSetFlagOk() (*bool, bool)`

GetSetFlagOk returns a tuple with the SetFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetFlag

`func (o *FilesystemSetImmutable) SetSetFlag(v bool)`

SetSetFlag sets SetFlag field to given value.

### HasSetFlag

`func (o *FilesystemSetImmutable) HasSetFlag() bool`

HasSetFlag returns a boolean if a field has been set.

### GetPath

`func (o *FilesystemSetImmutable) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FilesystemSetImmutable) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FilesystemSetImmutable) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FilesystemSetImmutable) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


