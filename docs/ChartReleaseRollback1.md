# ChartReleaseRollback1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceRollback** | Pointer to **bool** | &#x60;force_rollback&#x60; is a boolean which when set will force rollback operation to move forward even if no snapshots are found. This is only useful when &#x60;rollback_snapshot&#x60; is set. | [optional] [default to false]
**RecreateResources** | Pointer to **bool** | &#x60;recreate_resources&#x60; is a boolean which will delete and then create the kubernetes resources on rollback of chart release. This should be used with caution as if chart release is consuming immutable objects like a PVC, the rollback operation can&#39;t be performed and will fail as helm tries to do a 3 way patch for rollback. | [optional] [default to false]
**RollbackSnapshot** | Pointer to **bool** | &#x60;rollback_snapshot&#x60; is a boolean value which when set will rollback snapshots of any PVC&#39;s or ix volumes being consumed by the chart release. &#x60;force_rollback&#x60; is a boolean which when set will force rollback operation to move forward even if no snapshots are found. This is only useful when &#x60;rollback_snapshot&#x60; is set. | [optional] [default to true]
**ItemVersion** | Pointer to **string** | &#x60;item_version&#x60; is version which we want to rollback a chart release to. | [optional] 

## Methods

### NewChartReleaseRollback1

`func NewChartReleaseRollback1() *ChartReleaseRollback1`

NewChartReleaseRollback1 instantiates a new ChartReleaseRollback1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartReleaseRollback1WithDefaults

`func NewChartReleaseRollback1WithDefaults() *ChartReleaseRollback1`

NewChartReleaseRollback1WithDefaults instantiates a new ChartReleaseRollback1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceRollback

`func (o *ChartReleaseRollback1) GetForceRollback() bool`

GetForceRollback returns the ForceRollback field if non-nil, zero value otherwise.

### GetForceRollbackOk

`func (o *ChartReleaseRollback1) GetForceRollbackOk() (*bool, bool)`

GetForceRollbackOk returns a tuple with the ForceRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRollback

`func (o *ChartReleaseRollback1) SetForceRollback(v bool)`

SetForceRollback sets ForceRollback field to given value.

### HasForceRollback

`func (o *ChartReleaseRollback1) HasForceRollback() bool`

HasForceRollback returns a boolean if a field has been set.

### GetRecreateResources

`func (o *ChartReleaseRollback1) GetRecreateResources() bool`

GetRecreateResources returns the RecreateResources field if non-nil, zero value otherwise.

### GetRecreateResourcesOk

`func (o *ChartReleaseRollback1) GetRecreateResourcesOk() (*bool, bool)`

GetRecreateResourcesOk returns a tuple with the RecreateResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecreateResources

`func (o *ChartReleaseRollback1) SetRecreateResources(v bool)`

SetRecreateResources sets RecreateResources field to given value.

### HasRecreateResources

`func (o *ChartReleaseRollback1) HasRecreateResources() bool`

HasRecreateResources returns a boolean if a field has been set.

### GetRollbackSnapshot

`func (o *ChartReleaseRollback1) GetRollbackSnapshot() bool`

GetRollbackSnapshot returns the RollbackSnapshot field if non-nil, zero value otherwise.

### GetRollbackSnapshotOk

`func (o *ChartReleaseRollback1) GetRollbackSnapshotOk() (*bool, bool)`

GetRollbackSnapshotOk returns a tuple with the RollbackSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackSnapshot

`func (o *ChartReleaseRollback1) SetRollbackSnapshot(v bool)`

SetRollbackSnapshot sets RollbackSnapshot field to given value.

### HasRollbackSnapshot

`func (o *ChartReleaseRollback1) HasRollbackSnapshot() bool`

HasRollbackSnapshot returns a boolean if a field has been set.

### GetItemVersion

`func (o *ChartReleaseRollback1) GetItemVersion() string`

GetItemVersion returns the ItemVersion field if non-nil, zero value otherwise.

### GetItemVersionOk

`func (o *ChartReleaseRollback1) GetItemVersionOk() (*string, bool)`

GetItemVersionOk returns a tuple with the ItemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersion

`func (o *ChartReleaseRollback1) SetItemVersion(v string)`

SetItemVersion sets ItemVersion field to given value.

### HasItemVersion

`func (o *ChartReleaseRollback1) HasItemVersion() bool`

HasItemVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


