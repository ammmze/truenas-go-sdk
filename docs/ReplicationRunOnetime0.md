# ReplicationRunOnetime0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to **string** |  | [optional] 
**SshCredentials** | Pointer to **NullableInt32** |  | [optional] 
**NetcatActiveSide** | Pointer to **NullableString** |  | [optional] 
**NetcatActiveSideListenAddress** | Pointer to **NullableString** |  | [optional] 
**NetcatActiveSidePortMin** | Pointer to **NullableInt32** |  | [optional] 
**NetcatActiveSidePortMax** | Pointer to **NullableInt32** |  | [optional] 
**NetcatPassiveSideConnectAddress** | Pointer to **NullableString** |  | [optional] 
**SourceDatasets** | Pointer to **[]string** |  | [optional] [default to []]
**TargetDataset** | Pointer to **string** |  | [optional] 
**Recursive** | Pointer to **bool** |  | [optional] 
**Exclude** | Pointer to **[]string** |  | [optional] [default to []]
**Properties** | Pointer to **bool** |  | [optional] [default to true]
**PropertiesExclude** | Pointer to **[]string** |  | [optional] [default to []]
**PropertiesOverride** | Pointer to **map[string]interface{}** |  | [optional] 
**Replicate** | Pointer to **bool** |  | [optional] [default to false]
**Encryption** | Pointer to **bool** |  | [optional] [default to false]
**EncryptionKey** | Pointer to **NullableString** |  | [optional] 
**EncryptionKeyFormat** | Pointer to **NullableString** |  | [optional] 
**EncryptionKeyLocation** | Pointer to **NullableString** |  | [optional] 
**PeriodicSnapshotTasks** | Pointer to **[]int32** |  | [optional] [default to []]
**NamingSchema** | Pointer to **[]string** |  | [optional] [default to []]
**AlsoIncludeNamingSchema** | Pointer to **[]string** |  | [optional] [default to []]
**NameRegex** | Pointer to **NullableString** |  | [optional] 
**RestrictSchedule** | Pointer to [**RestrictSchedule1**](RestrictSchedule1.md) |  | [optional] 
**AllowFromScratch** | Pointer to **bool** |  | [optional] [default to false]
**Readonly** | Pointer to **string** |  | [optional] [default to "SET"]
**HoldPendingSnapshots** | Pointer to **bool** |  | [optional] [default to false]
**RetentionPolicy** | Pointer to **string** |  | [optional] 
**LifetimeValue** | Pointer to **NullableInt32** |  | [optional] 
**LifetimeUnit** | Pointer to **NullableString** |  | [optional] 
**Lifetimes** | Pointer to [**[]Lifetime**](Lifetime.md) |  | [optional] [default to []]
**Compression** | Pointer to **NullableString** |  | [optional] 
**SpeedLimit** | Pointer to **NullableInt32** |  | [optional] 
**LargeBlock** | Pointer to **bool** |  | [optional] [default to true]
**Embed** | Pointer to **bool** |  | [optional] [default to false]
**Compressed** | Pointer to **bool** |  | [optional] [default to true]
**Retries** | Pointer to **int32** |  | [optional] [default to 5]
**LoggingLevel** | Pointer to **NullableString** |  | [optional] 
**ExcludeMountpointProperty** | Pointer to **bool** |  | [optional] [default to true]
**OnlyFromScratch** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewReplicationRunOnetime0

`func NewReplicationRunOnetime0() *ReplicationRunOnetime0`

NewReplicationRunOnetime0 instantiates a new ReplicationRunOnetime0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationRunOnetime0WithDefaults

`func NewReplicationRunOnetime0WithDefaults() *ReplicationRunOnetime0`

NewReplicationRunOnetime0WithDefaults instantiates a new ReplicationRunOnetime0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *ReplicationRunOnetime0) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ReplicationRunOnetime0) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ReplicationRunOnetime0) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ReplicationRunOnetime0) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetTransport

`func (o *ReplicationRunOnetime0) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *ReplicationRunOnetime0) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *ReplicationRunOnetime0) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *ReplicationRunOnetime0) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetSshCredentials

`func (o *ReplicationRunOnetime0) GetSshCredentials() int32`

GetSshCredentials returns the SshCredentials field if non-nil, zero value otherwise.

### GetSshCredentialsOk

`func (o *ReplicationRunOnetime0) GetSshCredentialsOk() (*int32, bool)`

GetSshCredentialsOk returns a tuple with the SshCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCredentials

`func (o *ReplicationRunOnetime0) SetSshCredentials(v int32)`

SetSshCredentials sets SshCredentials field to given value.

### HasSshCredentials

`func (o *ReplicationRunOnetime0) HasSshCredentials() bool`

HasSshCredentials returns a boolean if a field has been set.

### SetSshCredentialsNil

`func (o *ReplicationRunOnetime0) SetSshCredentialsNil(b bool)`

 SetSshCredentialsNil sets the value for SshCredentials to be an explicit nil

### UnsetSshCredentials
`func (o *ReplicationRunOnetime0) UnsetSshCredentials()`

UnsetSshCredentials ensures that no value is present for SshCredentials, not even an explicit nil
### GetNetcatActiveSide

`func (o *ReplicationRunOnetime0) GetNetcatActiveSide() string`

GetNetcatActiveSide returns the NetcatActiveSide field if non-nil, zero value otherwise.

### GetNetcatActiveSideOk

`func (o *ReplicationRunOnetime0) GetNetcatActiveSideOk() (*string, bool)`

GetNetcatActiveSideOk returns a tuple with the NetcatActiveSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetcatActiveSide

`func (o *ReplicationRunOnetime0) SetNetcatActiveSide(v string)`

SetNetcatActiveSide sets NetcatActiveSide field to given value.

### HasNetcatActiveSide

`func (o *ReplicationRunOnetime0) HasNetcatActiveSide() bool`

HasNetcatActiveSide returns a boolean if a field has been set.

### SetNetcatActiveSideNil

`func (o *ReplicationRunOnetime0) SetNetcatActiveSideNil(b bool)`

 SetNetcatActiveSideNil sets the value for NetcatActiveSide to be an explicit nil

### UnsetNetcatActiveSide
`func (o *ReplicationRunOnetime0) UnsetNetcatActiveSide()`

UnsetNetcatActiveSide ensures that no value is present for NetcatActiveSide, not even an explicit nil
### GetNetcatActiveSideListenAddress

`func (o *ReplicationRunOnetime0) GetNetcatActiveSideListenAddress() string`

GetNetcatActiveSideListenAddress returns the NetcatActiveSideListenAddress field if non-nil, zero value otherwise.

### GetNetcatActiveSideListenAddressOk

`func (o *ReplicationRunOnetime0) GetNetcatActiveSideListenAddressOk() (*string, bool)`

GetNetcatActiveSideListenAddressOk returns a tuple with the NetcatActiveSideListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetcatActiveSideListenAddress

`func (o *ReplicationRunOnetime0) SetNetcatActiveSideListenAddress(v string)`

SetNetcatActiveSideListenAddress sets NetcatActiveSideListenAddress field to given value.

### HasNetcatActiveSideListenAddress

`func (o *ReplicationRunOnetime0) HasNetcatActiveSideListenAddress() bool`

HasNetcatActiveSideListenAddress returns a boolean if a field has been set.

### SetNetcatActiveSideListenAddressNil

`func (o *ReplicationRunOnetime0) SetNetcatActiveSideListenAddressNil(b bool)`

 SetNetcatActiveSideListenAddressNil sets the value for NetcatActiveSideListenAddress to be an explicit nil

### UnsetNetcatActiveSideListenAddress
`func (o *ReplicationRunOnetime0) UnsetNetcatActiveSideListenAddress()`

UnsetNetcatActiveSideListenAddress ensures that no value is present for NetcatActiveSideListenAddress, not even an explicit nil
### GetNetcatActiveSidePortMin

`func (o *ReplicationRunOnetime0) GetNetcatActiveSidePortMin() int32`

GetNetcatActiveSidePortMin returns the NetcatActiveSidePortMin field if non-nil, zero value otherwise.

### GetNetcatActiveSidePortMinOk

`func (o *ReplicationRunOnetime0) GetNetcatActiveSidePortMinOk() (*int32, bool)`

GetNetcatActiveSidePortMinOk returns a tuple with the NetcatActiveSidePortMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetcatActiveSidePortMin

`func (o *ReplicationRunOnetime0) SetNetcatActiveSidePortMin(v int32)`

SetNetcatActiveSidePortMin sets NetcatActiveSidePortMin field to given value.

### HasNetcatActiveSidePortMin

`func (o *ReplicationRunOnetime0) HasNetcatActiveSidePortMin() bool`

HasNetcatActiveSidePortMin returns a boolean if a field has been set.

### SetNetcatActiveSidePortMinNil

`func (o *ReplicationRunOnetime0) SetNetcatActiveSidePortMinNil(b bool)`

 SetNetcatActiveSidePortMinNil sets the value for NetcatActiveSidePortMin to be an explicit nil

### UnsetNetcatActiveSidePortMin
`func (o *ReplicationRunOnetime0) UnsetNetcatActiveSidePortMin()`

UnsetNetcatActiveSidePortMin ensures that no value is present for NetcatActiveSidePortMin, not even an explicit nil
### GetNetcatActiveSidePortMax

`func (o *ReplicationRunOnetime0) GetNetcatActiveSidePortMax() int32`

GetNetcatActiveSidePortMax returns the NetcatActiveSidePortMax field if non-nil, zero value otherwise.

### GetNetcatActiveSidePortMaxOk

`func (o *ReplicationRunOnetime0) GetNetcatActiveSidePortMaxOk() (*int32, bool)`

GetNetcatActiveSidePortMaxOk returns a tuple with the NetcatActiveSidePortMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetcatActiveSidePortMax

`func (o *ReplicationRunOnetime0) SetNetcatActiveSidePortMax(v int32)`

SetNetcatActiveSidePortMax sets NetcatActiveSidePortMax field to given value.

### HasNetcatActiveSidePortMax

`func (o *ReplicationRunOnetime0) HasNetcatActiveSidePortMax() bool`

HasNetcatActiveSidePortMax returns a boolean if a field has been set.

### SetNetcatActiveSidePortMaxNil

`func (o *ReplicationRunOnetime0) SetNetcatActiveSidePortMaxNil(b bool)`

 SetNetcatActiveSidePortMaxNil sets the value for NetcatActiveSidePortMax to be an explicit nil

### UnsetNetcatActiveSidePortMax
`func (o *ReplicationRunOnetime0) UnsetNetcatActiveSidePortMax()`

UnsetNetcatActiveSidePortMax ensures that no value is present for NetcatActiveSidePortMax, not even an explicit nil
### GetNetcatPassiveSideConnectAddress

`func (o *ReplicationRunOnetime0) GetNetcatPassiveSideConnectAddress() string`

GetNetcatPassiveSideConnectAddress returns the NetcatPassiveSideConnectAddress field if non-nil, zero value otherwise.

### GetNetcatPassiveSideConnectAddressOk

`func (o *ReplicationRunOnetime0) GetNetcatPassiveSideConnectAddressOk() (*string, bool)`

GetNetcatPassiveSideConnectAddressOk returns a tuple with the NetcatPassiveSideConnectAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetcatPassiveSideConnectAddress

`func (o *ReplicationRunOnetime0) SetNetcatPassiveSideConnectAddress(v string)`

SetNetcatPassiveSideConnectAddress sets NetcatPassiveSideConnectAddress field to given value.

### HasNetcatPassiveSideConnectAddress

`func (o *ReplicationRunOnetime0) HasNetcatPassiveSideConnectAddress() bool`

HasNetcatPassiveSideConnectAddress returns a boolean if a field has been set.

### SetNetcatPassiveSideConnectAddressNil

`func (o *ReplicationRunOnetime0) SetNetcatPassiveSideConnectAddressNil(b bool)`

 SetNetcatPassiveSideConnectAddressNil sets the value for NetcatPassiveSideConnectAddress to be an explicit nil

### UnsetNetcatPassiveSideConnectAddress
`func (o *ReplicationRunOnetime0) UnsetNetcatPassiveSideConnectAddress()`

UnsetNetcatPassiveSideConnectAddress ensures that no value is present for NetcatPassiveSideConnectAddress, not even an explicit nil
### GetSourceDatasets

`func (o *ReplicationRunOnetime0) GetSourceDatasets() []string`

GetSourceDatasets returns the SourceDatasets field if non-nil, zero value otherwise.

### GetSourceDatasetsOk

`func (o *ReplicationRunOnetime0) GetSourceDatasetsOk() (*[]string, bool)`

GetSourceDatasetsOk returns a tuple with the SourceDatasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDatasets

`func (o *ReplicationRunOnetime0) SetSourceDatasets(v []string)`

SetSourceDatasets sets SourceDatasets field to given value.

### HasSourceDatasets

`func (o *ReplicationRunOnetime0) HasSourceDatasets() bool`

HasSourceDatasets returns a boolean if a field has been set.

### GetTargetDataset

`func (o *ReplicationRunOnetime0) GetTargetDataset() string`

GetTargetDataset returns the TargetDataset field if non-nil, zero value otherwise.

### GetTargetDatasetOk

`func (o *ReplicationRunOnetime0) GetTargetDatasetOk() (*string, bool)`

GetTargetDatasetOk returns a tuple with the TargetDataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDataset

`func (o *ReplicationRunOnetime0) SetTargetDataset(v string)`

SetTargetDataset sets TargetDataset field to given value.

### HasTargetDataset

`func (o *ReplicationRunOnetime0) HasTargetDataset() bool`

HasTargetDataset returns a boolean if a field has been set.

### GetRecursive

`func (o *ReplicationRunOnetime0) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *ReplicationRunOnetime0) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *ReplicationRunOnetime0) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *ReplicationRunOnetime0) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetExclude

`func (o *ReplicationRunOnetime0) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *ReplicationRunOnetime0) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *ReplicationRunOnetime0) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *ReplicationRunOnetime0) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetProperties

`func (o *ReplicationRunOnetime0) GetProperties() bool`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ReplicationRunOnetime0) GetPropertiesOk() (*bool, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ReplicationRunOnetime0) SetProperties(v bool)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ReplicationRunOnetime0) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPropertiesExclude

`func (o *ReplicationRunOnetime0) GetPropertiesExclude() []string`

GetPropertiesExclude returns the PropertiesExclude field if non-nil, zero value otherwise.

### GetPropertiesExcludeOk

`func (o *ReplicationRunOnetime0) GetPropertiesExcludeOk() (*[]string, bool)`

GetPropertiesExcludeOk returns a tuple with the PropertiesExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesExclude

`func (o *ReplicationRunOnetime0) SetPropertiesExclude(v []string)`

SetPropertiesExclude sets PropertiesExclude field to given value.

### HasPropertiesExclude

`func (o *ReplicationRunOnetime0) HasPropertiesExclude() bool`

HasPropertiesExclude returns a boolean if a field has been set.

### GetPropertiesOverride

`func (o *ReplicationRunOnetime0) GetPropertiesOverride() map[string]interface{}`

GetPropertiesOverride returns the PropertiesOverride field if non-nil, zero value otherwise.

### GetPropertiesOverrideOk

`func (o *ReplicationRunOnetime0) GetPropertiesOverrideOk() (*map[string]interface{}, bool)`

GetPropertiesOverrideOk returns a tuple with the PropertiesOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesOverride

`func (o *ReplicationRunOnetime0) SetPropertiesOverride(v map[string]interface{})`

SetPropertiesOverride sets PropertiesOverride field to given value.

### HasPropertiesOverride

`func (o *ReplicationRunOnetime0) HasPropertiesOverride() bool`

HasPropertiesOverride returns a boolean if a field has been set.

### GetReplicate

`func (o *ReplicationRunOnetime0) GetReplicate() bool`

GetReplicate returns the Replicate field if non-nil, zero value otherwise.

### GetReplicateOk

`func (o *ReplicationRunOnetime0) GetReplicateOk() (*bool, bool)`

GetReplicateOk returns a tuple with the Replicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicate

`func (o *ReplicationRunOnetime0) SetReplicate(v bool)`

SetReplicate sets Replicate field to given value.

### HasReplicate

`func (o *ReplicationRunOnetime0) HasReplicate() bool`

HasReplicate returns a boolean if a field has been set.

### GetEncryption

`func (o *ReplicationRunOnetime0) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *ReplicationRunOnetime0) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *ReplicationRunOnetime0) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *ReplicationRunOnetime0) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *ReplicationRunOnetime0) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *ReplicationRunOnetime0) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *ReplicationRunOnetime0) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *ReplicationRunOnetime0) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### SetEncryptionKeyNil

`func (o *ReplicationRunOnetime0) SetEncryptionKeyNil(b bool)`

 SetEncryptionKeyNil sets the value for EncryptionKey to be an explicit nil

### UnsetEncryptionKey
`func (o *ReplicationRunOnetime0) UnsetEncryptionKey()`

UnsetEncryptionKey ensures that no value is present for EncryptionKey, not even an explicit nil
### GetEncryptionKeyFormat

`func (o *ReplicationRunOnetime0) GetEncryptionKeyFormat() string`

GetEncryptionKeyFormat returns the EncryptionKeyFormat field if non-nil, zero value otherwise.

### GetEncryptionKeyFormatOk

`func (o *ReplicationRunOnetime0) GetEncryptionKeyFormatOk() (*string, bool)`

GetEncryptionKeyFormatOk returns a tuple with the EncryptionKeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyFormat

`func (o *ReplicationRunOnetime0) SetEncryptionKeyFormat(v string)`

SetEncryptionKeyFormat sets EncryptionKeyFormat field to given value.

### HasEncryptionKeyFormat

`func (o *ReplicationRunOnetime0) HasEncryptionKeyFormat() bool`

HasEncryptionKeyFormat returns a boolean if a field has been set.

### SetEncryptionKeyFormatNil

`func (o *ReplicationRunOnetime0) SetEncryptionKeyFormatNil(b bool)`

 SetEncryptionKeyFormatNil sets the value for EncryptionKeyFormat to be an explicit nil

### UnsetEncryptionKeyFormat
`func (o *ReplicationRunOnetime0) UnsetEncryptionKeyFormat()`

UnsetEncryptionKeyFormat ensures that no value is present for EncryptionKeyFormat, not even an explicit nil
### GetEncryptionKeyLocation

`func (o *ReplicationRunOnetime0) GetEncryptionKeyLocation() string`

GetEncryptionKeyLocation returns the EncryptionKeyLocation field if non-nil, zero value otherwise.

### GetEncryptionKeyLocationOk

`func (o *ReplicationRunOnetime0) GetEncryptionKeyLocationOk() (*string, bool)`

GetEncryptionKeyLocationOk returns a tuple with the EncryptionKeyLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyLocation

`func (o *ReplicationRunOnetime0) SetEncryptionKeyLocation(v string)`

SetEncryptionKeyLocation sets EncryptionKeyLocation field to given value.

### HasEncryptionKeyLocation

`func (o *ReplicationRunOnetime0) HasEncryptionKeyLocation() bool`

HasEncryptionKeyLocation returns a boolean if a field has been set.

### SetEncryptionKeyLocationNil

`func (o *ReplicationRunOnetime0) SetEncryptionKeyLocationNil(b bool)`

 SetEncryptionKeyLocationNil sets the value for EncryptionKeyLocation to be an explicit nil

### UnsetEncryptionKeyLocation
`func (o *ReplicationRunOnetime0) UnsetEncryptionKeyLocation()`

UnsetEncryptionKeyLocation ensures that no value is present for EncryptionKeyLocation, not even an explicit nil
### GetPeriodicSnapshotTasks

`func (o *ReplicationRunOnetime0) GetPeriodicSnapshotTasks() []int32`

GetPeriodicSnapshotTasks returns the PeriodicSnapshotTasks field if non-nil, zero value otherwise.

### GetPeriodicSnapshotTasksOk

`func (o *ReplicationRunOnetime0) GetPeriodicSnapshotTasksOk() (*[]int32, bool)`

GetPeriodicSnapshotTasksOk returns a tuple with the PeriodicSnapshotTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicSnapshotTasks

`func (o *ReplicationRunOnetime0) SetPeriodicSnapshotTasks(v []int32)`

SetPeriodicSnapshotTasks sets PeriodicSnapshotTasks field to given value.

### HasPeriodicSnapshotTasks

`func (o *ReplicationRunOnetime0) HasPeriodicSnapshotTasks() bool`

HasPeriodicSnapshotTasks returns a boolean if a field has been set.

### GetNamingSchema

`func (o *ReplicationRunOnetime0) GetNamingSchema() []string`

GetNamingSchema returns the NamingSchema field if non-nil, zero value otherwise.

### GetNamingSchemaOk

`func (o *ReplicationRunOnetime0) GetNamingSchemaOk() (*[]string, bool)`

GetNamingSchemaOk returns a tuple with the NamingSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingSchema

`func (o *ReplicationRunOnetime0) SetNamingSchema(v []string)`

SetNamingSchema sets NamingSchema field to given value.

### HasNamingSchema

`func (o *ReplicationRunOnetime0) HasNamingSchema() bool`

HasNamingSchema returns a boolean if a field has been set.

### GetAlsoIncludeNamingSchema

`func (o *ReplicationRunOnetime0) GetAlsoIncludeNamingSchema() []string`

GetAlsoIncludeNamingSchema returns the AlsoIncludeNamingSchema field if non-nil, zero value otherwise.

### GetAlsoIncludeNamingSchemaOk

`func (o *ReplicationRunOnetime0) GetAlsoIncludeNamingSchemaOk() (*[]string, bool)`

GetAlsoIncludeNamingSchemaOk returns a tuple with the AlsoIncludeNamingSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlsoIncludeNamingSchema

`func (o *ReplicationRunOnetime0) SetAlsoIncludeNamingSchema(v []string)`

SetAlsoIncludeNamingSchema sets AlsoIncludeNamingSchema field to given value.

### HasAlsoIncludeNamingSchema

`func (o *ReplicationRunOnetime0) HasAlsoIncludeNamingSchema() bool`

HasAlsoIncludeNamingSchema returns a boolean if a field has been set.

### GetNameRegex

`func (o *ReplicationRunOnetime0) GetNameRegex() string`

GetNameRegex returns the NameRegex field if non-nil, zero value otherwise.

### GetNameRegexOk

`func (o *ReplicationRunOnetime0) GetNameRegexOk() (*string, bool)`

GetNameRegexOk returns a tuple with the NameRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameRegex

`func (o *ReplicationRunOnetime0) SetNameRegex(v string)`

SetNameRegex sets NameRegex field to given value.

### HasNameRegex

`func (o *ReplicationRunOnetime0) HasNameRegex() bool`

HasNameRegex returns a boolean if a field has been set.

### SetNameRegexNil

`func (o *ReplicationRunOnetime0) SetNameRegexNil(b bool)`

 SetNameRegexNil sets the value for NameRegex to be an explicit nil

### UnsetNameRegex
`func (o *ReplicationRunOnetime0) UnsetNameRegex()`

UnsetNameRegex ensures that no value is present for NameRegex, not even an explicit nil
### GetRestrictSchedule

`func (o *ReplicationRunOnetime0) GetRestrictSchedule() RestrictSchedule1`

GetRestrictSchedule returns the RestrictSchedule field if non-nil, zero value otherwise.

### GetRestrictScheduleOk

`func (o *ReplicationRunOnetime0) GetRestrictScheduleOk() (*RestrictSchedule1, bool)`

GetRestrictScheduleOk returns a tuple with the RestrictSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictSchedule

`func (o *ReplicationRunOnetime0) SetRestrictSchedule(v RestrictSchedule1)`

SetRestrictSchedule sets RestrictSchedule field to given value.

### HasRestrictSchedule

`func (o *ReplicationRunOnetime0) HasRestrictSchedule() bool`

HasRestrictSchedule returns a boolean if a field has been set.

### GetAllowFromScratch

`func (o *ReplicationRunOnetime0) GetAllowFromScratch() bool`

GetAllowFromScratch returns the AllowFromScratch field if non-nil, zero value otherwise.

### GetAllowFromScratchOk

`func (o *ReplicationRunOnetime0) GetAllowFromScratchOk() (*bool, bool)`

GetAllowFromScratchOk returns a tuple with the AllowFromScratch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFromScratch

`func (o *ReplicationRunOnetime0) SetAllowFromScratch(v bool)`

SetAllowFromScratch sets AllowFromScratch field to given value.

### HasAllowFromScratch

`func (o *ReplicationRunOnetime0) HasAllowFromScratch() bool`

HasAllowFromScratch returns a boolean if a field has been set.

### GetReadonly

`func (o *ReplicationRunOnetime0) GetReadonly() string`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *ReplicationRunOnetime0) GetReadonlyOk() (*string, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *ReplicationRunOnetime0) SetReadonly(v string)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *ReplicationRunOnetime0) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetHoldPendingSnapshots

`func (o *ReplicationRunOnetime0) GetHoldPendingSnapshots() bool`

GetHoldPendingSnapshots returns the HoldPendingSnapshots field if non-nil, zero value otherwise.

### GetHoldPendingSnapshotsOk

`func (o *ReplicationRunOnetime0) GetHoldPendingSnapshotsOk() (*bool, bool)`

GetHoldPendingSnapshotsOk returns a tuple with the HoldPendingSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldPendingSnapshots

`func (o *ReplicationRunOnetime0) SetHoldPendingSnapshots(v bool)`

SetHoldPendingSnapshots sets HoldPendingSnapshots field to given value.

### HasHoldPendingSnapshots

`func (o *ReplicationRunOnetime0) HasHoldPendingSnapshots() bool`

HasHoldPendingSnapshots returns a boolean if a field has been set.

### GetRetentionPolicy

`func (o *ReplicationRunOnetime0) GetRetentionPolicy() string`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *ReplicationRunOnetime0) GetRetentionPolicyOk() (*string, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *ReplicationRunOnetime0) SetRetentionPolicy(v string)`

SetRetentionPolicy sets RetentionPolicy field to given value.

### HasRetentionPolicy

`func (o *ReplicationRunOnetime0) HasRetentionPolicy() bool`

HasRetentionPolicy returns a boolean if a field has been set.

### GetLifetimeValue

`func (o *ReplicationRunOnetime0) GetLifetimeValue() int32`

GetLifetimeValue returns the LifetimeValue field if non-nil, zero value otherwise.

### GetLifetimeValueOk

`func (o *ReplicationRunOnetime0) GetLifetimeValueOk() (*int32, bool)`

GetLifetimeValueOk returns a tuple with the LifetimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeValue

`func (o *ReplicationRunOnetime0) SetLifetimeValue(v int32)`

SetLifetimeValue sets LifetimeValue field to given value.

### HasLifetimeValue

`func (o *ReplicationRunOnetime0) HasLifetimeValue() bool`

HasLifetimeValue returns a boolean if a field has been set.

### SetLifetimeValueNil

`func (o *ReplicationRunOnetime0) SetLifetimeValueNil(b bool)`

 SetLifetimeValueNil sets the value for LifetimeValue to be an explicit nil

### UnsetLifetimeValue
`func (o *ReplicationRunOnetime0) UnsetLifetimeValue()`

UnsetLifetimeValue ensures that no value is present for LifetimeValue, not even an explicit nil
### GetLifetimeUnit

`func (o *ReplicationRunOnetime0) GetLifetimeUnit() string`

GetLifetimeUnit returns the LifetimeUnit field if non-nil, zero value otherwise.

### GetLifetimeUnitOk

`func (o *ReplicationRunOnetime0) GetLifetimeUnitOk() (*string, bool)`

GetLifetimeUnitOk returns a tuple with the LifetimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeUnit

`func (o *ReplicationRunOnetime0) SetLifetimeUnit(v string)`

SetLifetimeUnit sets LifetimeUnit field to given value.

### HasLifetimeUnit

`func (o *ReplicationRunOnetime0) HasLifetimeUnit() bool`

HasLifetimeUnit returns a boolean if a field has been set.

### SetLifetimeUnitNil

`func (o *ReplicationRunOnetime0) SetLifetimeUnitNil(b bool)`

 SetLifetimeUnitNil sets the value for LifetimeUnit to be an explicit nil

### UnsetLifetimeUnit
`func (o *ReplicationRunOnetime0) UnsetLifetimeUnit()`

UnsetLifetimeUnit ensures that no value is present for LifetimeUnit, not even an explicit nil
### GetLifetimes

`func (o *ReplicationRunOnetime0) GetLifetimes() []Lifetime`

GetLifetimes returns the Lifetimes field if non-nil, zero value otherwise.

### GetLifetimesOk

`func (o *ReplicationRunOnetime0) GetLifetimesOk() (*[]Lifetime, bool)`

GetLifetimesOk returns a tuple with the Lifetimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimes

`func (o *ReplicationRunOnetime0) SetLifetimes(v []Lifetime)`

SetLifetimes sets Lifetimes field to given value.

### HasLifetimes

`func (o *ReplicationRunOnetime0) HasLifetimes() bool`

HasLifetimes returns a boolean if a field has been set.

### GetCompression

`func (o *ReplicationRunOnetime0) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *ReplicationRunOnetime0) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *ReplicationRunOnetime0) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *ReplicationRunOnetime0) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### SetCompressionNil

`func (o *ReplicationRunOnetime0) SetCompressionNil(b bool)`

 SetCompressionNil sets the value for Compression to be an explicit nil

### UnsetCompression
`func (o *ReplicationRunOnetime0) UnsetCompression()`

UnsetCompression ensures that no value is present for Compression, not even an explicit nil
### GetSpeedLimit

`func (o *ReplicationRunOnetime0) GetSpeedLimit() int32`

GetSpeedLimit returns the SpeedLimit field if non-nil, zero value otherwise.

### GetSpeedLimitOk

`func (o *ReplicationRunOnetime0) GetSpeedLimitOk() (*int32, bool)`

GetSpeedLimitOk returns a tuple with the SpeedLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedLimit

`func (o *ReplicationRunOnetime0) SetSpeedLimit(v int32)`

SetSpeedLimit sets SpeedLimit field to given value.

### HasSpeedLimit

`func (o *ReplicationRunOnetime0) HasSpeedLimit() bool`

HasSpeedLimit returns a boolean if a field has been set.

### SetSpeedLimitNil

`func (o *ReplicationRunOnetime0) SetSpeedLimitNil(b bool)`

 SetSpeedLimitNil sets the value for SpeedLimit to be an explicit nil

### UnsetSpeedLimit
`func (o *ReplicationRunOnetime0) UnsetSpeedLimit()`

UnsetSpeedLimit ensures that no value is present for SpeedLimit, not even an explicit nil
### GetLargeBlock

`func (o *ReplicationRunOnetime0) GetLargeBlock() bool`

GetLargeBlock returns the LargeBlock field if non-nil, zero value otherwise.

### GetLargeBlockOk

`func (o *ReplicationRunOnetime0) GetLargeBlockOk() (*bool, bool)`

GetLargeBlockOk returns a tuple with the LargeBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeBlock

`func (o *ReplicationRunOnetime0) SetLargeBlock(v bool)`

SetLargeBlock sets LargeBlock field to given value.

### HasLargeBlock

`func (o *ReplicationRunOnetime0) HasLargeBlock() bool`

HasLargeBlock returns a boolean if a field has been set.

### GetEmbed

`func (o *ReplicationRunOnetime0) GetEmbed() bool`

GetEmbed returns the Embed field if non-nil, zero value otherwise.

### GetEmbedOk

`func (o *ReplicationRunOnetime0) GetEmbedOk() (*bool, bool)`

GetEmbedOk returns a tuple with the Embed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbed

`func (o *ReplicationRunOnetime0) SetEmbed(v bool)`

SetEmbed sets Embed field to given value.

### HasEmbed

`func (o *ReplicationRunOnetime0) HasEmbed() bool`

HasEmbed returns a boolean if a field has been set.

### GetCompressed

`func (o *ReplicationRunOnetime0) GetCompressed() bool`

GetCompressed returns the Compressed field if non-nil, zero value otherwise.

### GetCompressedOk

`func (o *ReplicationRunOnetime0) GetCompressedOk() (*bool, bool)`

GetCompressedOk returns a tuple with the Compressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressed

`func (o *ReplicationRunOnetime0) SetCompressed(v bool)`

SetCompressed sets Compressed field to given value.

### HasCompressed

`func (o *ReplicationRunOnetime0) HasCompressed() bool`

HasCompressed returns a boolean if a field has been set.

### GetRetries

`func (o *ReplicationRunOnetime0) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *ReplicationRunOnetime0) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *ReplicationRunOnetime0) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *ReplicationRunOnetime0) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetLoggingLevel

`func (o *ReplicationRunOnetime0) GetLoggingLevel() string`

GetLoggingLevel returns the LoggingLevel field if non-nil, zero value otherwise.

### GetLoggingLevelOk

`func (o *ReplicationRunOnetime0) GetLoggingLevelOk() (*string, bool)`

GetLoggingLevelOk returns a tuple with the LoggingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingLevel

`func (o *ReplicationRunOnetime0) SetLoggingLevel(v string)`

SetLoggingLevel sets LoggingLevel field to given value.

### HasLoggingLevel

`func (o *ReplicationRunOnetime0) HasLoggingLevel() bool`

HasLoggingLevel returns a boolean if a field has been set.

### SetLoggingLevelNil

`func (o *ReplicationRunOnetime0) SetLoggingLevelNil(b bool)`

 SetLoggingLevelNil sets the value for LoggingLevel to be an explicit nil

### UnsetLoggingLevel
`func (o *ReplicationRunOnetime0) UnsetLoggingLevel()`

UnsetLoggingLevel ensures that no value is present for LoggingLevel, not even an explicit nil
### GetExcludeMountpointProperty

`func (o *ReplicationRunOnetime0) GetExcludeMountpointProperty() bool`

GetExcludeMountpointProperty returns the ExcludeMountpointProperty field if non-nil, zero value otherwise.

### GetExcludeMountpointPropertyOk

`func (o *ReplicationRunOnetime0) GetExcludeMountpointPropertyOk() (*bool, bool)`

GetExcludeMountpointPropertyOk returns a tuple with the ExcludeMountpointProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeMountpointProperty

`func (o *ReplicationRunOnetime0) SetExcludeMountpointProperty(v bool)`

SetExcludeMountpointProperty sets ExcludeMountpointProperty field to given value.

### HasExcludeMountpointProperty

`func (o *ReplicationRunOnetime0) HasExcludeMountpointProperty() bool`

HasExcludeMountpointProperty returns a boolean if a field has been set.

### GetOnlyFromScratch

`func (o *ReplicationRunOnetime0) GetOnlyFromScratch() bool`

GetOnlyFromScratch returns the OnlyFromScratch field if non-nil, zero value otherwise.

### GetOnlyFromScratchOk

`func (o *ReplicationRunOnetime0) GetOnlyFromScratchOk() (*bool, bool)`

GetOnlyFromScratchOk returns a tuple with the OnlyFromScratch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyFromScratch

`func (o *ReplicationRunOnetime0) SetOnlyFromScratch(v bool)`

SetOnlyFromScratch sets OnlyFromScratch field to given value.

### HasOnlyFromScratch

`func (o *ReplicationRunOnetime0) HasOnlyFromScratch() bool`

HasOnlyFromScratch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


