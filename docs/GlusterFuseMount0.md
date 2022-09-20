# GlusterFuseMount0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | &#x60;name&#x60; String representing the name of the gluster volume | [optional] 
**All** | Pointer to **bool** | &#x60;all&#x60; Boolean if True locally FUSE mount all detected         gluster volumes | [optional] [default to false]
**Raise** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewGlusterFuseMount0

`func NewGlusterFuseMount0() *GlusterFuseMount0`

NewGlusterFuseMount0 instantiates a new GlusterFuseMount0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlusterFuseMount0WithDefaults

`func NewGlusterFuseMount0WithDefaults() *GlusterFuseMount0`

NewGlusterFuseMount0WithDefaults instantiates a new GlusterFuseMount0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GlusterFuseMount0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlusterFuseMount0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlusterFuseMount0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlusterFuseMount0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAll

`func (o *GlusterFuseMount0) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *GlusterFuseMount0) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *GlusterFuseMount0) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *GlusterFuseMount0) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetRaise

`func (o *GlusterFuseMount0) GetRaise() bool`

GetRaise returns the Raise field if non-nil, zero value otherwise.

### GetRaiseOk

`func (o *GlusterFuseMount0) GetRaiseOk() (*bool, bool)`

GetRaiseOk returns a tuple with the Raise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaise

`func (o *GlusterFuseMount0) SetRaise(v bool)`

SetRaise sets Raise field to given value.

### HasRaise

`func (o *GlusterFuseMount0) HasRaise() bool`

HasRaise returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


