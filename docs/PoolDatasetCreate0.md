# PoolDatasetCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | &#x60;encryption&#x60; when enabled will create an ZFS encrypted root dataset for &#x60;name&#x60; pool. There are 2 cases where ZFS encryption is not allowed for a dataset: 1) Pool in question is GELI encrypted. 2) If the parent dataset is encrypted with a passphrase and &#x60;name&#x60; is being created     with a key for encrypting the dataset. &#x60;encryption_options&#x60; specifies configuration for encryption of dataset for &#x60;name&#x60; pool. &#x60;encryption_options.passphrase&#x60; must be specified if encryption for dataset is desired with a passphrase as a key. Otherwise a hex encoded key can be specified by providing &#x60;encryption_options.key&#x60;. &#x60;encryption_options.generate_key&#x60; when enabled automatically generates the key to be used for dataset encryption. | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "FILESYSTEM"]
**Volsize** | Pointer to **int32** | &#x60;volsize&#x60; is required for type&#x3D;VOLUME and is supposed to be a multiple of the block size. | [optional] 
**Volblocksize** | Pointer to **string** | &#x60;sparse&#x60; and &#x60;volblocksize&#x60; are only used for type&#x3D;VOLUME. | [optional] 
**Sparse** | Pointer to **bool** | &#x60;sparse&#x60; and &#x60;volblocksize&#x60; are only used for type&#x3D;VOLUME. | [optional] 
**ForceSize** | Pointer to **bool** |  | [optional] 
**Comments** | Pointer to [**Comments**](Comments.md) |  | [optional] 
**Sync** | Pointer to **string** |  | [optional] 
**Compression** | Pointer to **string** |  | [optional] 
**Atime** | Pointer to **string** |  | [optional] 
**Exec** | Pointer to **string** |  | [optional] 
**Managedby** | Pointer to [**Managedby**](Managedby.md) |  | [optional] 
**Quota** | Pointer to **NullableInt32** |  | [optional] 
**QuotaWarning** | Pointer to [**QuotaWarning**](QuotaWarning.md) |  | [optional] 
**QuotaCritical** | Pointer to [**QuotaCritical**](QuotaCritical.md) |  | [optional] 
**Refquota** | Pointer to **NullableInt32** |  | [optional] 
**RefquotaWarning** | Pointer to [**RefquotaWarning**](RefquotaWarning.md) |  | [optional] 
**RefquotaCritical** | Pointer to [**RefquotaCritical**](RefquotaCritical.md) |  | [optional] 
**Reservation** | Pointer to **int32** |  | [optional] 
**Refreservation** | Pointer to **int32** |  | [optional] 
**SpecialSmallBlockSize** | Pointer to [**SpecialSmallBlockSize**](SpecialSmallBlockSize.md) |  | [optional] 
**Copies** | Pointer to [**Copies**](Copies.md) |  | [optional] 
**Snapdir** | Pointer to **string** |  | [optional] 
**Deduplication** | Pointer to **string** |  | [optional] 
**Checksum** | Pointer to **string** |  | [optional] 
**Readonly** | Pointer to **string** |  | [optional] 
**Recordsize** | Pointer to [**Recordsize**](Recordsize.md) |  | [optional] 
**Casesensitivity** | Pointer to **string** |  | [optional] 
**Aclmode** | Pointer to **string** |  | [optional] 
**Acltype** | Pointer to **string** |  | [optional] 
**ShareType** | Pointer to **string** |  | [optional] [default to "GENERIC"]
**Xattr** | Pointer to **string** |  | [optional] [default to "SA"]
**EncryptionOptions** | Pointer to [**EncryptionOptions1**](EncryptionOptions1.md) |  | [optional] 
**Encryption** | Pointer to **bool** | &#x60;encryption&#x60; when enabled will create an ZFS encrypted root dataset for &#x60;name&#x60; pool. There are 2 cases where ZFS encryption is not allowed for a dataset: 1) Pool in question is GELI encrypted. 2) If the parent dataset is encrypted with a passphrase and &#x60;name&#x60; is being created     with a key for encrypting the dataset. | [optional] [default to false]
**InheritEncryption** | Pointer to **bool** |  | [optional] [default to true]
**UserProperties** | Pointer to [**[]UserProperty**](UserProperty.md) |  | [optional] [default to []]
**CreateAncestors** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPoolDatasetCreate0

`func NewPoolDatasetCreate0() *PoolDatasetCreate0`

NewPoolDatasetCreate0 instantiates a new PoolDatasetCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolDatasetCreate0WithDefaults

`func NewPoolDatasetCreate0WithDefaults() *PoolDatasetCreate0`

NewPoolDatasetCreate0WithDefaults instantiates a new PoolDatasetCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PoolDatasetCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolDatasetCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolDatasetCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoolDatasetCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PoolDatasetCreate0) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PoolDatasetCreate0) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PoolDatasetCreate0) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PoolDatasetCreate0) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVolsize

`func (o *PoolDatasetCreate0) GetVolsize() int32`

GetVolsize returns the Volsize field if non-nil, zero value otherwise.

### GetVolsizeOk

`func (o *PoolDatasetCreate0) GetVolsizeOk() (*int32, bool)`

GetVolsizeOk returns a tuple with the Volsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolsize

`func (o *PoolDatasetCreate0) SetVolsize(v int32)`

SetVolsize sets Volsize field to given value.

### HasVolsize

`func (o *PoolDatasetCreate0) HasVolsize() bool`

HasVolsize returns a boolean if a field has been set.

### GetVolblocksize

`func (o *PoolDatasetCreate0) GetVolblocksize() string`

GetVolblocksize returns the Volblocksize field if non-nil, zero value otherwise.

### GetVolblocksizeOk

`func (o *PoolDatasetCreate0) GetVolblocksizeOk() (*string, bool)`

GetVolblocksizeOk returns a tuple with the Volblocksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolblocksize

`func (o *PoolDatasetCreate0) SetVolblocksize(v string)`

SetVolblocksize sets Volblocksize field to given value.

### HasVolblocksize

`func (o *PoolDatasetCreate0) HasVolblocksize() bool`

HasVolblocksize returns a boolean if a field has been set.

### GetSparse

`func (o *PoolDatasetCreate0) GetSparse() bool`

GetSparse returns the Sparse field if non-nil, zero value otherwise.

### GetSparseOk

`func (o *PoolDatasetCreate0) GetSparseOk() (*bool, bool)`

GetSparseOk returns a tuple with the Sparse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparse

`func (o *PoolDatasetCreate0) SetSparse(v bool)`

SetSparse sets Sparse field to given value.

### HasSparse

`func (o *PoolDatasetCreate0) HasSparse() bool`

HasSparse returns a boolean if a field has been set.

### GetForceSize

`func (o *PoolDatasetCreate0) GetForceSize() bool`

GetForceSize returns the ForceSize field if non-nil, zero value otherwise.

### GetForceSizeOk

`func (o *PoolDatasetCreate0) GetForceSizeOk() (*bool, bool)`

GetForceSizeOk returns a tuple with the ForceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSize

`func (o *PoolDatasetCreate0) SetForceSize(v bool)`

SetForceSize sets ForceSize field to given value.

### HasForceSize

`func (o *PoolDatasetCreate0) HasForceSize() bool`

HasForceSize returns a boolean if a field has been set.

### GetComments

`func (o *PoolDatasetCreate0) GetComments() Comments`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PoolDatasetCreate0) GetCommentsOk() (*Comments, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PoolDatasetCreate0) SetComments(v Comments)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PoolDatasetCreate0) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetSync

`func (o *PoolDatasetCreate0) GetSync() string`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *PoolDatasetCreate0) GetSyncOk() (*string, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *PoolDatasetCreate0) SetSync(v string)`

SetSync sets Sync field to given value.

### HasSync

`func (o *PoolDatasetCreate0) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetCompression

`func (o *PoolDatasetCreate0) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *PoolDatasetCreate0) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *PoolDatasetCreate0) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *PoolDatasetCreate0) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetAtime

`func (o *PoolDatasetCreate0) GetAtime() string`

GetAtime returns the Atime field if non-nil, zero value otherwise.

### GetAtimeOk

`func (o *PoolDatasetCreate0) GetAtimeOk() (*string, bool)`

GetAtimeOk returns a tuple with the Atime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtime

`func (o *PoolDatasetCreate0) SetAtime(v string)`

SetAtime sets Atime field to given value.

### HasAtime

`func (o *PoolDatasetCreate0) HasAtime() bool`

HasAtime returns a boolean if a field has been set.

### GetExec

`func (o *PoolDatasetCreate0) GetExec() string`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *PoolDatasetCreate0) GetExecOk() (*string, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *PoolDatasetCreate0) SetExec(v string)`

SetExec sets Exec field to given value.

### HasExec

`func (o *PoolDatasetCreate0) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetManagedby

`func (o *PoolDatasetCreate0) GetManagedby() Managedby`

GetManagedby returns the Managedby field if non-nil, zero value otherwise.

### GetManagedbyOk

`func (o *PoolDatasetCreate0) GetManagedbyOk() (*Managedby, bool)`

GetManagedbyOk returns a tuple with the Managedby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedby

`func (o *PoolDatasetCreate0) SetManagedby(v Managedby)`

SetManagedby sets Managedby field to given value.

### HasManagedby

`func (o *PoolDatasetCreate0) HasManagedby() bool`

HasManagedby returns a boolean if a field has been set.

### GetQuota

`func (o *PoolDatasetCreate0) GetQuota() int32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *PoolDatasetCreate0) GetQuotaOk() (*int32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *PoolDatasetCreate0) SetQuota(v int32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *PoolDatasetCreate0) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuotaNil

`func (o *PoolDatasetCreate0) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *PoolDatasetCreate0) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil
### GetQuotaWarning

`func (o *PoolDatasetCreate0) GetQuotaWarning() QuotaWarning`

GetQuotaWarning returns the QuotaWarning field if non-nil, zero value otherwise.

### GetQuotaWarningOk

`func (o *PoolDatasetCreate0) GetQuotaWarningOk() (*QuotaWarning, bool)`

GetQuotaWarningOk returns a tuple with the QuotaWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaWarning

`func (o *PoolDatasetCreate0) SetQuotaWarning(v QuotaWarning)`

SetQuotaWarning sets QuotaWarning field to given value.

### HasQuotaWarning

`func (o *PoolDatasetCreate0) HasQuotaWarning() bool`

HasQuotaWarning returns a boolean if a field has been set.

### GetQuotaCritical

`func (o *PoolDatasetCreate0) GetQuotaCritical() QuotaCritical`

GetQuotaCritical returns the QuotaCritical field if non-nil, zero value otherwise.

### GetQuotaCriticalOk

`func (o *PoolDatasetCreate0) GetQuotaCriticalOk() (*QuotaCritical, bool)`

GetQuotaCriticalOk returns a tuple with the QuotaCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaCritical

`func (o *PoolDatasetCreate0) SetQuotaCritical(v QuotaCritical)`

SetQuotaCritical sets QuotaCritical field to given value.

### HasQuotaCritical

`func (o *PoolDatasetCreate0) HasQuotaCritical() bool`

HasQuotaCritical returns a boolean if a field has been set.

### GetRefquota

`func (o *PoolDatasetCreate0) GetRefquota() int32`

GetRefquota returns the Refquota field if non-nil, zero value otherwise.

### GetRefquotaOk

`func (o *PoolDatasetCreate0) GetRefquotaOk() (*int32, bool)`

GetRefquotaOk returns a tuple with the Refquota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefquota

`func (o *PoolDatasetCreate0) SetRefquota(v int32)`

SetRefquota sets Refquota field to given value.

### HasRefquota

`func (o *PoolDatasetCreate0) HasRefquota() bool`

HasRefquota returns a boolean if a field has been set.

### SetRefquotaNil

`func (o *PoolDatasetCreate0) SetRefquotaNil(b bool)`

 SetRefquotaNil sets the value for Refquota to be an explicit nil

### UnsetRefquota
`func (o *PoolDatasetCreate0) UnsetRefquota()`

UnsetRefquota ensures that no value is present for Refquota, not even an explicit nil
### GetRefquotaWarning

`func (o *PoolDatasetCreate0) GetRefquotaWarning() RefquotaWarning`

GetRefquotaWarning returns the RefquotaWarning field if non-nil, zero value otherwise.

### GetRefquotaWarningOk

`func (o *PoolDatasetCreate0) GetRefquotaWarningOk() (*RefquotaWarning, bool)`

GetRefquotaWarningOk returns a tuple with the RefquotaWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefquotaWarning

`func (o *PoolDatasetCreate0) SetRefquotaWarning(v RefquotaWarning)`

SetRefquotaWarning sets RefquotaWarning field to given value.

### HasRefquotaWarning

`func (o *PoolDatasetCreate0) HasRefquotaWarning() bool`

HasRefquotaWarning returns a boolean if a field has been set.

### GetRefquotaCritical

`func (o *PoolDatasetCreate0) GetRefquotaCritical() RefquotaCritical`

GetRefquotaCritical returns the RefquotaCritical field if non-nil, zero value otherwise.

### GetRefquotaCriticalOk

`func (o *PoolDatasetCreate0) GetRefquotaCriticalOk() (*RefquotaCritical, bool)`

GetRefquotaCriticalOk returns a tuple with the RefquotaCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefquotaCritical

`func (o *PoolDatasetCreate0) SetRefquotaCritical(v RefquotaCritical)`

SetRefquotaCritical sets RefquotaCritical field to given value.

### HasRefquotaCritical

`func (o *PoolDatasetCreate0) HasRefquotaCritical() bool`

HasRefquotaCritical returns a boolean if a field has been set.

### GetReservation

`func (o *PoolDatasetCreate0) GetReservation() int32`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *PoolDatasetCreate0) GetReservationOk() (*int32, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *PoolDatasetCreate0) SetReservation(v int32)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *PoolDatasetCreate0) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetRefreservation

`func (o *PoolDatasetCreate0) GetRefreservation() int32`

GetRefreservation returns the Refreservation field if non-nil, zero value otherwise.

### GetRefreservationOk

`func (o *PoolDatasetCreate0) GetRefreservationOk() (*int32, bool)`

GetRefreservationOk returns a tuple with the Refreservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreservation

`func (o *PoolDatasetCreate0) SetRefreservation(v int32)`

SetRefreservation sets Refreservation field to given value.

### HasRefreservation

`func (o *PoolDatasetCreate0) HasRefreservation() bool`

HasRefreservation returns a boolean if a field has been set.

### GetSpecialSmallBlockSize

`func (o *PoolDatasetCreate0) GetSpecialSmallBlockSize() SpecialSmallBlockSize`

GetSpecialSmallBlockSize returns the SpecialSmallBlockSize field if non-nil, zero value otherwise.

### GetSpecialSmallBlockSizeOk

`func (o *PoolDatasetCreate0) GetSpecialSmallBlockSizeOk() (*SpecialSmallBlockSize, bool)`

GetSpecialSmallBlockSizeOk returns a tuple with the SpecialSmallBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialSmallBlockSize

`func (o *PoolDatasetCreate0) SetSpecialSmallBlockSize(v SpecialSmallBlockSize)`

SetSpecialSmallBlockSize sets SpecialSmallBlockSize field to given value.

### HasSpecialSmallBlockSize

`func (o *PoolDatasetCreate0) HasSpecialSmallBlockSize() bool`

HasSpecialSmallBlockSize returns a boolean if a field has been set.

### GetCopies

`func (o *PoolDatasetCreate0) GetCopies() Copies`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *PoolDatasetCreate0) GetCopiesOk() (*Copies, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *PoolDatasetCreate0) SetCopies(v Copies)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *PoolDatasetCreate0) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetSnapdir

`func (o *PoolDatasetCreate0) GetSnapdir() string`

GetSnapdir returns the Snapdir field if non-nil, zero value otherwise.

### GetSnapdirOk

`func (o *PoolDatasetCreate0) GetSnapdirOk() (*string, bool)`

GetSnapdirOk returns a tuple with the Snapdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapdir

`func (o *PoolDatasetCreate0) SetSnapdir(v string)`

SetSnapdir sets Snapdir field to given value.

### HasSnapdir

`func (o *PoolDatasetCreate0) HasSnapdir() bool`

HasSnapdir returns a boolean if a field has been set.

### GetDeduplication

`func (o *PoolDatasetCreate0) GetDeduplication() string`

GetDeduplication returns the Deduplication field if non-nil, zero value otherwise.

### GetDeduplicationOk

`func (o *PoolDatasetCreate0) GetDeduplicationOk() (*string, bool)`

GetDeduplicationOk returns a tuple with the Deduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplication

`func (o *PoolDatasetCreate0) SetDeduplication(v string)`

SetDeduplication sets Deduplication field to given value.

### HasDeduplication

`func (o *PoolDatasetCreate0) HasDeduplication() bool`

HasDeduplication returns a boolean if a field has been set.

### GetChecksum

`func (o *PoolDatasetCreate0) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *PoolDatasetCreate0) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *PoolDatasetCreate0) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *PoolDatasetCreate0) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetReadonly

`func (o *PoolDatasetCreate0) GetReadonly() string`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *PoolDatasetCreate0) GetReadonlyOk() (*string, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *PoolDatasetCreate0) SetReadonly(v string)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *PoolDatasetCreate0) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetRecordsize

`func (o *PoolDatasetCreate0) GetRecordsize() Recordsize`

GetRecordsize returns the Recordsize field if non-nil, zero value otherwise.

### GetRecordsizeOk

`func (o *PoolDatasetCreate0) GetRecordsizeOk() (*Recordsize, bool)`

GetRecordsizeOk returns a tuple with the Recordsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsize

`func (o *PoolDatasetCreate0) SetRecordsize(v Recordsize)`

SetRecordsize sets Recordsize field to given value.

### HasRecordsize

`func (o *PoolDatasetCreate0) HasRecordsize() bool`

HasRecordsize returns a boolean if a field has been set.

### GetCasesensitivity

`func (o *PoolDatasetCreate0) GetCasesensitivity() string`

GetCasesensitivity returns the Casesensitivity field if non-nil, zero value otherwise.

### GetCasesensitivityOk

`func (o *PoolDatasetCreate0) GetCasesensitivityOk() (*string, bool)`

GetCasesensitivityOk returns a tuple with the Casesensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasesensitivity

`func (o *PoolDatasetCreate0) SetCasesensitivity(v string)`

SetCasesensitivity sets Casesensitivity field to given value.

### HasCasesensitivity

`func (o *PoolDatasetCreate0) HasCasesensitivity() bool`

HasCasesensitivity returns a boolean if a field has been set.

### GetAclmode

`func (o *PoolDatasetCreate0) GetAclmode() string`

GetAclmode returns the Aclmode field if non-nil, zero value otherwise.

### GetAclmodeOk

`func (o *PoolDatasetCreate0) GetAclmodeOk() (*string, bool)`

GetAclmodeOk returns a tuple with the Aclmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclmode

`func (o *PoolDatasetCreate0) SetAclmode(v string)`

SetAclmode sets Aclmode field to given value.

### HasAclmode

`func (o *PoolDatasetCreate0) HasAclmode() bool`

HasAclmode returns a boolean if a field has been set.

### GetAcltype

`func (o *PoolDatasetCreate0) GetAcltype() string`

GetAcltype returns the Acltype field if non-nil, zero value otherwise.

### GetAcltypeOk

`func (o *PoolDatasetCreate0) GetAcltypeOk() (*string, bool)`

GetAcltypeOk returns a tuple with the Acltype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcltype

`func (o *PoolDatasetCreate0) SetAcltype(v string)`

SetAcltype sets Acltype field to given value.

### HasAcltype

`func (o *PoolDatasetCreate0) HasAcltype() bool`

HasAcltype returns a boolean if a field has been set.

### GetShareType

`func (o *PoolDatasetCreate0) GetShareType() string`

GetShareType returns the ShareType field if non-nil, zero value otherwise.

### GetShareTypeOk

`func (o *PoolDatasetCreate0) GetShareTypeOk() (*string, bool)`

GetShareTypeOk returns a tuple with the ShareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareType

`func (o *PoolDatasetCreate0) SetShareType(v string)`

SetShareType sets ShareType field to given value.

### HasShareType

`func (o *PoolDatasetCreate0) HasShareType() bool`

HasShareType returns a boolean if a field has been set.

### GetXattr

`func (o *PoolDatasetCreate0) GetXattr() string`

GetXattr returns the Xattr field if non-nil, zero value otherwise.

### GetXattrOk

`func (o *PoolDatasetCreate0) GetXattrOk() (*string, bool)`

GetXattrOk returns a tuple with the Xattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXattr

`func (o *PoolDatasetCreate0) SetXattr(v string)`

SetXattr sets Xattr field to given value.

### HasXattr

`func (o *PoolDatasetCreate0) HasXattr() bool`

HasXattr returns a boolean if a field has been set.

### GetEncryptionOptions

`func (o *PoolDatasetCreate0) GetEncryptionOptions() EncryptionOptions1`

GetEncryptionOptions returns the EncryptionOptions field if non-nil, zero value otherwise.

### GetEncryptionOptionsOk

`func (o *PoolDatasetCreate0) GetEncryptionOptionsOk() (*EncryptionOptions1, bool)`

GetEncryptionOptionsOk returns a tuple with the EncryptionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionOptions

`func (o *PoolDatasetCreate0) SetEncryptionOptions(v EncryptionOptions1)`

SetEncryptionOptions sets EncryptionOptions field to given value.

### HasEncryptionOptions

`func (o *PoolDatasetCreate0) HasEncryptionOptions() bool`

HasEncryptionOptions returns a boolean if a field has been set.

### GetEncryption

`func (o *PoolDatasetCreate0) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *PoolDatasetCreate0) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *PoolDatasetCreate0) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *PoolDatasetCreate0) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetInheritEncryption

`func (o *PoolDatasetCreate0) GetInheritEncryption() bool`

GetInheritEncryption returns the InheritEncryption field if non-nil, zero value otherwise.

### GetInheritEncryptionOk

`func (o *PoolDatasetCreate0) GetInheritEncryptionOk() (*bool, bool)`

GetInheritEncryptionOk returns a tuple with the InheritEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritEncryption

`func (o *PoolDatasetCreate0) SetInheritEncryption(v bool)`

SetInheritEncryption sets InheritEncryption field to given value.

### HasInheritEncryption

`func (o *PoolDatasetCreate0) HasInheritEncryption() bool`

HasInheritEncryption returns a boolean if a field has been set.

### GetUserProperties

`func (o *PoolDatasetCreate0) GetUserProperties() []UserProperty`

GetUserProperties returns the UserProperties field if non-nil, zero value otherwise.

### GetUserPropertiesOk

`func (o *PoolDatasetCreate0) GetUserPropertiesOk() (*[]UserProperty, bool)`

GetUserPropertiesOk returns a tuple with the UserProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProperties

`func (o *PoolDatasetCreate0) SetUserProperties(v []UserProperty)`

SetUserProperties sets UserProperties field to given value.

### HasUserProperties

`func (o *PoolDatasetCreate0) HasUserProperties() bool`

HasUserProperties returns a boolean if a field has been set.

### GetCreateAncestors

`func (o *PoolDatasetCreate0) GetCreateAncestors() bool`

GetCreateAncestors returns the CreateAncestors field if non-nil, zero value otherwise.

### GetCreateAncestorsOk

`func (o *PoolDatasetCreate0) GetCreateAncestorsOk() (*bool, bool)`

GetCreateAncestorsOk returns a tuple with the CreateAncestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAncestors

`func (o *PoolDatasetCreate0) SetCreateAncestors(v bool)`

SetCreateAncestors sets CreateAncestors field to given value.

### HasCreateAncestors

`func (o *PoolDatasetCreate0) HasCreateAncestors() bool`

HasCreateAncestors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


