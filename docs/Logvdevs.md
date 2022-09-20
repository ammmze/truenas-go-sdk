# Logvdevs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Disks** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewLogvdevs

`func NewLogvdevs() *Logvdevs`

NewLogvdevs instantiates a new Logvdevs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogvdevsWithDefaults

`func NewLogvdevsWithDefaults() *Logvdevs`

NewLogvdevsWithDefaults instantiates a new Logvdevs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Logvdevs) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Logvdevs) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Logvdevs) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Logvdevs) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisks

`func (o *Logvdevs) GetDisks() []string`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *Logvdevs) GetDisksOk() (*[]string, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *Logvdevs) SetDisks(v []string)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *Logvdevs) HasDisks() bool`

HasDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


