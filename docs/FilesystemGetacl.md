# FilesystemGetacl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Simplified** | Pointer to **bool** | &#x60;simplified&#x60; - effect of this depends on ACL type on underlying filesystem. In the case of NFSv4 ACLs simplified permissions and flags are returned for ACL entries where applicable. NFSv4 errata below. In the case of POSIX1E ACls, this setting has no impact on returned ACL. &#x60;simplified&#x60; returns a shortened form of the ACL permset and flags where applicable. If permissions have been simplified, then the &#x60;perms&#x60; object will contain only a single &#x60;BASIC&#x60; key with a string describing the underlying permissions set. | [optional] [default to true]
**ResolveIds** | Pointer to **bool** | &#x60;resolve_ids&#x60; - adds additional &#x60;who&#x60; key to each ACL entry, that converts the numeric id to a user name or group name. In the case of owner@ and group@ (NFSv4) or USER_OBJ and GROUP_OBJ (POSIX1E), st_uid or st_gid will be converted from stat() return for file. In the case of MASK (POSIX1E), OTHER (POSIX1E), everyone@ (NFSv4), key &#x60;who&#x60; will be included, but set to null. In case of failure to resolve the id to a name, &#x60;who&#x60; will be set to null. This option should only be used if resolving ids to names is required. | [optional] [default to false]

## Methods

### NewFilesystemGetacl

`func NewFilesystemGetacl() *FilesystemGetacl`

NewFilesystemGetacl instantiates a new FilesystemGetacl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemGetaclWithDefaults

`func NewFilesystemGetaclWithDefaults() *FilesystemGetacl`

NewFilesystemGetaclWithDefaults instantiates a new FilesystemGetacl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *FilesystemGetacl) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FilesystemGetacl) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FilesystemGetacl) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FilesystemGetacl) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSimplified

`func (o *FilesystemGetacl) GetSimplified() bool`

GetSimplified returns the Simplified field if non-nil, zero value otherwise.

### GetSimplifiedOk

`func (o *FilesystemGetacl) GetSimplifiedOk() (*bool, bool)`

GetSimplifiedOk returns a tuple with the Simplified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplified

`func (o *FilesystemGetacl) SetSimplified(v bool)`

SetSimplified sets Simplified field to given value.

### HasSimplified

`func (o *FilesystemGetacl) HasSimplified() bool`

HasSimplified returns a boolean if a field has been set.

### GetResolveIds

`func (o *FilesystemGetacl) GetResolveIds() bool`

GetResolveIds returns the ResolveIds field if non-nil, zero value otherwise.

### GetResolveIdsOk

`func (o *FilesystemGetacl) GetResolveIdsOk() (*bool, bool)`

GetResolveIdsOk returns a tuple with the ResolveIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveIds

`func (o *FilesystemGetacl) SetResolveIds(v bool)`

SetResolveIds sets ResolveIds field to given value.

### HasResolveIds

`func (o *FilesystemGetacl) HasResolveIds() bool`

HasResolveIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


