# IdmapNssOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkedService** | Pointer to **string** |  | [optional] [default to "LOCAL_ACCOUNT"]

## Methods

### NewIdmapNssOptions

`func NewIdmapNssOptions() *IdmapNssOptions`

NewIdmapNssOptions instantiates a new IdmapNssOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdmapNssOptionsWithDefaults

`func NewIdmapNssOptionsWithDefaults() *IdmapNssOptions`

NewIdmapNssOptionsWithDefaults instantiates a new IdmapNssOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkedService

`func (o *IdmapNssOptions) GetLinkedService() string`

GetLinkedService returns the LinkedService field if non-nil, zero value otherwise.

### GetLinkedServiceOk

`func (o *IdmapNssOptions) GetLinkedServiceOk() (*string, bool)`

GetLinkedServiceOk returns a tuple with the LinkedService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedService

`func (o *IdmapNssOptions) SetLinkedService(v string)`

SetLinkedService sets LinkedService field to given value.

### HasLinkedService

`func (o *IdmapNssOptions) HasLinkedService() bool`

HasLinkedService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


