# GlusterVolumeCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | &#x60;name&#x60; String representing name to be given to the volume | [optional] 
**Bricks** | Pointer to [**[]Brick**](Brick.md) | &#x60;bricks&#x60; List representing the brick paths     &#x60;peer_name&#x60; String representing IP or DNS name of the peer     &#x60;peer_path&#x60; String representing the full path of the brick | [optional] [default to []]
**Replica** | Pointer to **int32** | &#x60;replica&#x60; Integer representing number of replica bricks | [optional] 
**Arbiter** | Pointer to **int32** | &#x60;arbiter&#x60; Integer representing number of arbiter bricks | [optional] 
**Disperse** | Pointer to **int32** | &#x60;disperse&#x60; Integer representing number of disperse bricks | [optional] 
**DisperseData** | Pointer to **int32** | &#x60;disperse_data&#x60; Integer representing number of disperse data bricks | [optional] 
**Redundancy** | Pointer to **int32** | &#x60;redundancy&#x60; Integer representing number of redundancy bricks | [optional] 
**Force** | Pointer to **bool** |  | [optional] 

## Methods

### NewGlusterVolumeCreate0

`func NewGlusterVolumeCreate0() *GlusterVolumeCreate0`

NewGlusterVolumeCreate0 instantiates a new GlusterVolumeCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlusterVolumeCreate0WithDefaults

`func NewGlusterVolumeCreate0WithDefaults() *GlusterVolumeCreate0`

NewGlusterVolumeCreate0WithDefaults instantiates a new GlusterVolumeCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GlusterVolumeCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlusterVolumeCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlusterVolumeCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlusterVolumeCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBricks

`func (o *GlusterVolumeCreate0) GetBricks() []Brick`

GetBricks returns the Bricks field if non-nil, zero value otherwise.

### GetBricksOk

`func (o *GlusterVolumeCreate0) GetBricksOk() (*[]Brick, bool)`

GetBricksOk returns a tuple with the Bricks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBricks

`func (o *GlusterVolumeCreate0) SetBricks(v []Brick)`

SetBricks sets Bricks field to given value.

### HasBricks

`func (o *GlusterVolumeCreate0) HasBricks() bool`

HasBricks returns a boolean if a field has been set.

### GetReplica

`func (o *GlusterVolumeCreate0) GetReplica() int32`

GetReplica returns the Replica field if non-nil, zero value otherwise.

### GetReplicaOk

`func (o *GlusterVolumeCreate0) GetReplicaOk() (*int32, bool)`

GetReplicaOk returns a tuple with the Replica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplica

`func (o *GlusterVolumeCreate0) SetReplica(v int32)`

SetReplica sets Replica field to given value.

### HasReplica

`func (o *GlusterVolumeCreate0) HasReplica() bool`

HasReplica returns a boolean if a field has been set.

### GetArbiter

`func (o *GlusterVolumeCreate0) GetArbiter() int32`

GetArbiter returns the Arbiter field if non-nil, zero value otherwise.

### GetArbiterOk

`func (o *GlusterVolumeCreate0) GetArbiterOk() (*int32, bool)`

GetArbiterOk returns a tuple with the Arbiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArbiter

`func (o *GlusterVolumeCreate0) SetArbiter(v int32)`

SetArbiter sets Arbiter field to given value.

### HasArbiter

`func (o *GlusterVolumeCreate0) HasArbiter() bool`

HasArbiter returns a boolean if a field has been set.

### GetDisperse

`func (o *GlusterVolumeCreate0) GetDisperse() int32`

GetDisperse returns the Disperse field if non-nil, zero value otherwise.

### GetDisperseOk

`func (o *GlusterVolumeCreate0) GetDisperseOk() (*int32, bool)`

GetDisperseOk returns a tuple with the Disperse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisperse

`func (o *GlusterVolumeCreate0) SetDisperse(v int32)`

SetDisperse sets Disperse field to given value.

### HasDisperse

`func (o *GlusterVolumeCreate0) HasDisperse() bool`

HasDisperse returns a boolean if a field has been set.

### GetDisperseData

`func (o *GlusterVolumeCreate0) GetDisperseData() int32`

GetDisperseData returns the DisperseData field if non-nil, zero value otherwise.

### GetDisperseDataOk

`func (o *GlusterVolumeCreate0) GetDisperseDataOk() (*int32, bool)`

GetDisperseDataOk returns a tuple with the DisperseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisperseData

`func (o *GlusterVolumeCreate0) SetDisperseData(v int32)`

SetDisperseData sets DisperseData field to given value.

### HasDisperseData

`func (o *GlusterVolumeCreate0) HasDisperseData() bool`

HasDisperseData returns a boolean if a field has been set.

### GetRedundancy

`func (o *GlusterVolumeCreate0) GetRedundancy() int32`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *GlusterVolumeCreate0) GetRedundancyOk() (*int32, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *GlusterVolumeCreate0) SetRedundancy(v int32)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *GlusterVolumeCreate0) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetForce

`func (o *GlusterVolumeCreate0) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *GlusterVolumeCreate0) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *GlusterVolumeCreate0) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *GlusterVolumeCreate0) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


