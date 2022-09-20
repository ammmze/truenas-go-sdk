# ReplicationCountEligibleManualSnapshots0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasets** | Pointer to **[]string** |  | [optional] [default to []]
**NamingSchema** | Pointer to **[]string** | Count how many existing snapshots of &#x60;dataset&#x60; match &#x60;naming_schema&#x60;. | [optional] [default to []]
**NameRegex** | Pointer to **NullableString** |  | [optional] 
**Transport** | Pointer to **string** |  | [optional] 
**SshCredentials** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewReplicationCountEligibleManualSnapshots0

`func NewReplicationCountEligibleManualSnapshots0() *ReplicationCountEligibleManualSnapshots0`

NewReplicationCountEligibleManualSnapshots0 instantiates a new ReplicationCountEligibleManualSnapshots0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationCountEligibleManualSnapshots0WithDefaults

`func NewReplicationCountEligibleManualSnapshots0WithDefaults() *ReplicationCountEligibleManualSnapshots0`

NewReplicationCountEligibleManualSnapshots0WithDefaults instantiates a new ReplicationCountEligibleManualSnapshots0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasets

`func (o *ReplicationCountEligibleManualSnapshots0) GetDatasets() []string`

GetDatasets returns the Datasets field if non-nil, zero value otherwise.

### GetDatasetsOk

`func (o *ReplicationCountEligibleManualSnapshots0) GetDatasetsOk() (*[]string, bool)`

GetDatasetsOk returns a tuple with the Datasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasets

`func (o *ReplicationCountEligibleManualSnapshots0) SetDatasets(v []string)`

SetDatasets sets Datasets field to given value.

### HasDatasets

`func (o *ReplicationCountEligibleManualSnapshots0) HasDatasets() bool`

HasDatasets returns a boolean if a field has been set.

### GetNamingSchema

`func (o *ReplicationCountEligibleManualSnapshots0) GetNamingSchema() []string`

GetNamingSchema returns the NamingSchema field if non-nil, zero value otherwise.

### GetNamingSchemaOk

`func (o *ReplicationCountEligibleManualSnapshots0) GetNamingSchemaOk() (*[]string, bool)`

GetNamingSchemaOk returns a tuple with the NamingSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingSchema

`func (o *ReplicationCountEligibleManualSnapshots0) SetNamingSchema(v []string)`

SetNamingSchema sets NamingSchema field to given value.

### HasNamingSchema

`func (o *ReplicationCountEligibleManualSnapshots0) HasNamingSchema() bool`

HasNamingSchema returns a boolean if a field has been set.

### GetNameRegex

`func (o *ReplicationCountEligibleManualSnapshots0) GetNameRegex() string`

GetNameRegex returns the NameRegex field if non-nil, zero value otherwise.

### GetNameRegexOk

`func (o *ReplicationCountEligibleManualSnapshots0) GetNameRegexOk() (*string, bool)`

GetNameRegexOk returns a tuple with the NameRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameRegex

`func (o *ReplicationCountEligibleManualSnapshots0) SetNameRegex(v string)`

SetNameRegex sets NameRegex field to given value.

### HasNameRegex

`func (o *ReplicationCountEligibleManualSnapshots0) HasNameRegex() bool`

HasNameRegex returns a boolean if a field has been set.

### SetNameRegexNil

`func (o *ReplicationCountEligibleManualSnapshots0) SetNameRegexNil(b bool)`

 SetNameRegexNil sets the value for NameRegex to be an explicit nil

### UnsetNameRegex
`func (o *ReplicationCountEligibleManualSnapshots0) UnsetNameRegex()`

UnsetNameRegex ensures that no value is present for NameRegex, not even an explicit nil
### GetTransport

`func (o *ReplicationCountEligibleManualSnapshots0) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *ReplicationCountEligibleManualSnapshots0) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *ReplicationCountEligibleManualSnapshots0) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *ReplicationCountEligibleManualSnapshots0) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetSshCredentials

`func (o *ReplicationCountEligibleManualSnapshots0) GetSshCredentials() int32`

GetSshCredentials returns the SshCredentials field if non-nil, zero value otherwise.

### GetSshCredentialsOk

`func (o *ReplicationCountEligibleManualSnapshots0) GetSshCredentialsOk() (*int32, bool)`

GetSshCredentialsOk returns a tuple with the SshCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCredentials

`func (o *ReplicationCountEligibleManualSnapshots0) SetSshCredentials(v int32)`

SetSshCredentials sets SshCredentials field to given value.

### HasSshCredentials

`func (o *ReplicationCountEligibleManualSnapshots0) HasSshCredentials() bool`

HasSshCredentials returns a boolean if a field has been set.

### SetSshCredentialsNil

`func (o *ReplicationCountEligibleManualSnapshots0) SetSshCredentialsNil(b bool)`

 SetSshCredentialsNil sets the value for SshCredentials to be an explicit nil

### UnsetSshCredentials
`func (o *ReplicationCountEligibleManualSnapshots0) UnsetSshCredentials()`

UnsetSshCredentials ensures that no value is present for SshCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


