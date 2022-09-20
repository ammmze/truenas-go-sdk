# GlusterVolumeAddbrick0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | &#x60;name&#x60; String representing name of gluster volume | [optional] 
**Bricks** | Pointer to [**[]Brick**](Brick.md) | &#x60;bricks&#x60; List representing the brick paths     &#x60;peer_name&#x60; String representing IP or DNS name of the peer     &#x60;peer_path&#x60; String representing the full path of the brick | [optional] [default to []]
**Replica** | Pointer to **int32** | &#x60;replica&#x60; Integer replicating replica count | [optional] 
**Arbiter** | Pointer to **int32** | &#x60;arbiter&#x60; Integer replicating arbiter count | [optional] 
**Force** | Pointer to **bool** |  | [optional] 

## Methods

### NewGlusterVolumeAddbrick0

`func NewGlusterVolumeAddbrick0() *GlusterVolumeAddbrick0`

NewGlusterVolumeAddbrick0 instantiates a new GlusterVolumeAddbrick0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlusterVolumeAddbrick0WithDefaults

`func NewGlusterVolumeAddbrick0WithDefaults() *GlusterVolumeAddbrick0`

NewGlusterVolumeAddbrick0WithDefaults instantiates a new GlusterVolumeAddbrick0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GlusterVolumeAddbrick0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlusterVolumeAddbrick0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlusterVolumeAddbrick0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlusterVolumeAddbrick0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBricks

`func (o *GlusterVolumeAddbrick0) GetBricks() []Brick`

GetBricks returns the Bricks field if non-nil, zero value otherwise.

### GetBricksOk

`func (o *GlusterVolumeAddbrick0) GetBricksOk() (*[]Brick, bool)`

GetBricksOk returns a tuple with the Bricks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBricks

`func (o *GlusterVolumeAddbrick0) SetBricks(v []Brick)`

SetBricks sets Bricks field to given value.

### HasBricks

`func (o *GlusterVolumeAddbrick0) HasBricks() bool`

HasBricks returns a boolean if a field has been set.

### GetReplica

`func (o *GlusterVolumeAddbrick0) GetReplica() int32`

GetReplica returns the Replica field if non-nil, zero value otherwise.

### GetReplicaOk

`func (o *GlusterVolumeAddbrick0) GetReplicaOk() (*int32, bool)`

GetReplicaOk returns a tuple with the Replica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplica

`func (o *GlusterVolumeAddbrick0) SetReplica(v int32)`

SetReplica sets Replica field to given value.

### HasReplica

`func (o *GlusterVolumeAddbrick0) HasReplica() bool`

HasReplica returns a boolean if a field has been set.

### GetArbiter

`func (o *GlusterVolumeAddbrick0) GetArbiter() int32`

GetArbiter returns the Arbiter field if non-nil, zero value otherwise.

### GetArbiterOk

`func (o *GlusterVolumeAddbrick0) GetArbiterOk() (*int32, bool)`

GetArbiterOk returns a tuple with the Arbiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArbiter

`func (o *GlusterVolumeAddbrick0) SetArbiter(v int32)`

SetArbiter sets Arbiter field to given value.

### HasArbiter

`func (o *GlusterVolumeAddbrick0) HasArbiter() bool`

HasArbiter returns a boolean if a field has been set.

### GetForce

`func (o *GlusterVolumeAddbrick0) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *GlusterVolumeAddbrick0) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *GlusterVolumeAddbrick0) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *GlusterVolumeAddbrick0) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


