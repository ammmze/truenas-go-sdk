# GlusterVolumeRemovebrick0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | &#x60;name&#x60; String representing name of gluster volume | [optional] 
**Bricks** | Pointer to [**[]Brick**](Brick.md) | &#x60;bricks&#x60; List representing the brick paths     &#x60;peer_name&#x60; String representing IP or DNS name of the peer     &#x60;peer_path&#x60; String representing the full path of the brick | [optional] [default to []]
**Operation** | Pointer to **string** | &#x60;operation&#x60; String representing the operation to be performed     &#x60;START&#x60; Start the removal of the brick(s)     &#x60;STOP&#x60; Stop the removal of the brick(s)     &#x60;COMMIT&#x60; Commit the removal of the brick(s)     &#x60;STATUS&#x60; Display status of the removal of the brick(s)     &#x60;FORCE&#x60; Force the removal of the brick(s) | [optional] 
**Replica** | Pointer to **int32** |  | [optional] 

## Methods

### NewGlusterVolumeRemovebrick0

`func NewGlusterVolumeRemovebrick0() *GlusterVolumeRemovebrick0`

NewGlusterVolumeRemovebrick0 instantiates a new GlusterVolumeRemovebrick0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlusterVolumeRemovebrick0WithDefaults

`func NewGlusterVolumeRemovebrick0WithDefaults() *GlusterVolumeRemovebrick0`

NewGlusterVolumeRemovebrick0WithDefaults instantiates a new GlusterVolumeRemovebrick0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GlusterVolumeRemovebrick0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlusterVolumeRemovebrick0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlusterVolumeRemovebrick0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlusterVolumeRemovebrick0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBricks

`func (o *GlusterVolumeRemovebrick0) GetBricks() []Brick`

GetBricks returns the Bricks field if non-nil, zero value otherwise.

### GetBricksOk

`func (o *GlusterVolumeRemovebrick0) GetBricksOk() (*[]Brick, bool)`

GetBricksOk returns a tuple with the Bricks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBricks

`func (o *GlusterVolumeRemovebrick0) SetBricks(v []Brick)`

SetBricks sets Bricks field to given value.

### HasBricks

`func (o *GlusterVolumeRemovebrick0) HasBricks() bool`

HasBricks returns a boolean if a field has been set.

### GetOperation

`func (o *GlusterVolumeRemovebrick0) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *GlusterVolumeRemovebrick0) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *GlusterVolumeRemovebrick0) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *GlusterVolumeRemovebrick0) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetReplica

`func (o *GlusterVolumeRemovebrick0) GetReplica() int32`

GetReplica returns the Replica field if non-nil, zero value otherwise.

### GetReplicaOk

`func (o *GlusterVolumeRemovebrick0) GetReplicaOk() (*int32, bool)`

GetReplicaOk returns a tuple with the Replica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplica

`func (o *GlusterVolumeRemovebrick0) SetReplica(v int32)`

SetReplica sets Replica field to given value.

### HasReplica

`func (o *GlusterVolumeRemovebrick0) HasReplica() bool`

HasReplica returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


