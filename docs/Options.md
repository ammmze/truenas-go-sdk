# Options

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recursive** | Pointer to **bool** |  | [optional] [default to false]
**Traverse** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewOptions

`func NewOptions() *Options`

NewOptions instantiates a new Options object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsWithDefaults

`func NewOptionsWithDefaults() *Options`

NewOptionsWithDefaults instantiates a new Options object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecursive

`func (o *Options) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *Options) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *Options) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *Options) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetTraverse

`func (o *Options) GetTraverse() bool`

GetTraverse returns the Traverse field if non-nil, zero value otherwise.

### GetTraverseOk

`func (o *Options) GetTraverseOk() (*bool, bool)`

GetTraverseOk returns a tuple with the Traverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraverse

`func (o *Options) SetTraverse(v bool)`

SetTraverse sets Traverse field to given value.

### HasTraverse

`func (o *Options) HasTraverse() bool`

HasTraverse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


