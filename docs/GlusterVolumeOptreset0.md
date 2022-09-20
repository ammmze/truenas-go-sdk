# GlusterVolumeOptreset0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | &#x60;name&#x60; String representing name of gluster volume | [optional] 
**Opt** | Pointer to **string** | Reset volumes options.     If &#x60;opt&#x60; is not provided, then all options     will be reset. &#x60;opt&#x60; String representing name of the option to reset | [optional] 
**Force** | Pointer to **bool** |  | [optional] 

## Methods

### NewGlusterVolumeOptreset0

`func NewGlusterVolumeOptreset0() *GlusterVolumeOptreset0`

NewGlusterVolumeOptreset0 instantiates a new GlusterVolumeOptreset0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlusterVolumeOptreset0WithDefaults

`func NewGlusterVolumeOptreset0WithDefaults() *GlusterVolumeOptreset0`

NewGlusterVolumeOptreset0WithDefaults instantiates a new GlusterVolumeOptreset0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GlusterVolumeOptreset0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlusterVolumeOptreset0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlusterVolumeOptreset0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlusterVolumeOptreset0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpt

`func (o *GlusterVolumeOptreset0) GetOpt() string`

GetOpt returns the Opt field if non-nil, zero value otherwise.

### GetOptOk

`func (o *GlusterVolumeOptreset0) GetOptOk() (*string, bool)`

GetOptOk returns a tuple with the Opt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpt

`func (o *GlusterVolumeOptreset0) SetOpt(v string)`

SetOpt sets Opt field to given value.

### HasOpt

`func (o *GlusterVolumeOptreset0) HasOpt() bool`

HasOpt returns a boolean if a field has been set.

### GetForce

`func (o *GlusterVolumeOptreset0) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *GlusterVolumeOptreset0) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *GlusterVolumeOptreset0) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *GlusterVolumeOptreset0) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


