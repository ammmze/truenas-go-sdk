# PoolDatasetUnlock1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Force** | Pointer to **bool** |  | [optional] [default to false]
**KeyFile** | Pointer to **bool** |  | [optional] [default to false]
**Recursive** | Pointer to **bool** |  | [optional] [default to false]
**ToggleAttachments** | Pointer to **bool** | &#x60;toggle_attachments&#x60; controls whether attachments  should be put in action after unlocking dataset(s). Toggling attachments can theoretically lead to service interruption when daemons configurations are reloaded (this should not happen,  and if this happens it should be considered a bug). As TrueNAS does not have a state for resources that should be unlocked but are still locked, disabling this option will put the system into an inconsistent state so it should really never be disabled. | [optional] [default to true]
**Datasets** | Pointer to [**[]Dataset**](Dataset.md) |  | [optional] [default to []]

## Methods

### NewPoolDatasetUnlock1

`func NewPoolDatasetUnlock1() *PoolDatasetUnlock1`

NewPoolDatasetUnlock1 instantiates a new PoolDatasetUnlock1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolDatasetUnlock1WithDefaults

`func NewPoolDatasetUnlock1WithDefaults() *PoolDatasetUnlock1`

NewPoolDatasetUnlock1WithDefaults instantiates a new PoolDatasetUnlock1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForce

`func (o *PoolDatasetUnlock1) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *PoolDatasetUnlock1) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *PoolDatasetUnlock1) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *PoolDatasetUnlock1) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetKeyFile

`func (o *PoolDatasetUnlock1) GetKeyFile() bool`

GetKeyFile returns the KeyFile field if non-nil, zero value otherwise.

### GetKeyFileOk

`func (o *PoolDatasetUnlock1) GetKeyFileOk() (*bool, bool)`

GetKeyFileOk returns a tuple with the KeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFile

`func (o *PoolDatasetUnlock1) SetKeyFile(v bool)`

SetKeyFile sets KeyFile field to given value.

### HasKeyFile

`func (o *PoolDatasetUnlock1) HasKeyFile() bool`

HasKeyFile returns a boolean if a field has been set.

### GetRecursive

`func (o *PoolDatasetUnlock1) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *PoolDatasetUnlock1) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *PoolDatasetUnlock1) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *PoolDatasetUnlock1) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetToggleAttachments

`func (o *PoolDatasetUnlock1) GetToggleAttachments() bool`

GetToggleAttachments returns the ToggleAttachments field if non-nil, zero value otherwise.

### GetToggleAttachmentsOk

`func (o *PoolDatasetUnlock1) GetToggleAttachmentsOk() (*bool, bool)`

GetToggleAttachmentsOk returns a tuple with the ToggleAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToggleAttachments

`func (o *PoolDatasetUnlock1) SetToggleAttachments(v bool)`

SetToggleAttachments sets ToggleAttachments field to given value.

### HasToggleAttachments

`func (o *PoolDatasetUnlock1) HasToggleAttachments() bool`

HasToggleAttachments returns a boolean if a field has been set.

### GetDatasets

`func (o *PoolDatasetUnlock1) GetDatasets() []Dataset`

GetDatasets returns the Datasets field if non-nil, zero value otherwise.

### GetDatasetsOk

`func (o *PoolDatasetUnlock1) GetDatasetsOk() (*[]Dataset, bool)`

GetDatasetsOk returns a tuple with the Datasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasets

`func (o *PoolDatasetUnlock1) SetDatasets(v []Dataset)`

SetDatasets sets Datasets field to given value.

### HasDatasets

`func (o *PoolDatasetUnlock1) HasDatasets() bool`

HasDatasets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


