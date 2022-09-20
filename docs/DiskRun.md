# DiskRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] [default to "BACKGROUND"]
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewDiskRun

`func NewDiskRun() *DiskRun`

NewDiskRun instantiates a new DiskRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskRunWithDefaults

`func NewDiskRunWithDefaults() *DiskRun`

NewDiskRunWithDefaults instantiates a new DiskRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *DiskRun) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *DiskRun) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *DiskRun) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *DiskRun) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetMode

`func (o *DiskRun) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DiskRun) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DiskRun) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DiskRun) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetType

`func (o *DiskRun) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiskRun) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiskRun) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiskRun) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


