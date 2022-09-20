# GlusterFuseUmount0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | &#x60;name&#x60; String representing the name of the gluster volume | [optional] 
**All** | Pointer to **bool** | &#x60;all&#x60; Boolean if True umount all locally detected FUSE         mounted gluster volumes | [optional] [default to false]
**Raise** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewGlusterFuseUmount0

`func NewGlusterFuseUmount0() *GlusterFuseUmount0`

NewGlusterFuseUmount0 instantiates a new GlusterFuseUmount0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlusterFuseUmount0WithDefaults

`func NewGlusterFuseUmount0WithDefaults() *GlusterFuseUmount0`

NewGlusterFuseUmount0WithDefaults instantiates a new GlusterFuseUmount0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GlusterFuseUmount0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlusterFuseUmount0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlusterFuseUmount0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlusterFuseUmount0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAll

`func (o *GlusterFuseUmount0) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *GlusterFuseUmount0) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *GlusterFuseUmount0) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *GlusterFuseUmount0) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetRaise

`func (o *GlusterFuseUmount0) GetRaise() bool`

GetRaise returns the Raise field if non-nil, zero value otherwise.

### GetRaiseOk

`func (o *GlusterFuseUmount0) GetRaiseOk() (*bool, bool)`

GetRaiseOk returns a tuple with the Raise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaise

`func (o *GlusterFuseUmount0) SetRaise(v bool)`

SetRaise sets Raise field to given value.

### HasRaise

`func (o *GlusterFuseUmount0) HasRaise() bool`

HasRaise returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


