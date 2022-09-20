# Lifetime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | Pointer to [**Schedule5**](Schedule5.md) |  | [optional] 
**LifetimeValue** | Pointer to **int32** |  | [optional] 
**LifetimeUnit** | Pointer to **string** |  | [optional] 

## Methods

### NewLifetime

`func NewLifetime() *Lifetime`

NewLifetime instantiates a new Lifetime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifetimeWithDefaults

`func NewLifetimeWithDefaults() *Lifetime`

NewLifetimeWithDefaults instantiates a new Lifetime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *Lifetime) GetSchedule() Schedule5`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Lifetime) GetScheduleOk() (*Schedule5, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Lifetime) SetSchedule(v Schedule5)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Lifetime) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetLifetimeValue

`func (o *Lifetime) GetLifetimeValue() int32`

GetLifetimeValue returns the LifetimeValue field if non-nil, zero value otherwise.

### GetLifetimeValueOk

`func (o *Lifetime) GetLifetimeValueOk() (*int32, bool)`

GetLifetimeValueOk returns a tuple with the LifetimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeValue

`func (o *Lifetime) SetLifetimeValue(v int32)`

SetLifetimeValue sets LifetimeValue field to given value.

### HasLifetimeValue

`func (o *Lifetime) HasLifetimeValue() bool`

HasLifetimeValue returns a boolean if a field has been set.

### GetLifetimeUnit

`func (o *Lifetime) GetLifetimeUnit() string`

GetLifetimeUnit returns the LifetimeUnit field if non-nil, zero value otherwise.

### GetLifetimeUnitOk

`func (o *Lifetime) GetLifetimeUnitOk() (*string, bool)`

GetLifetimeUnitOk returns a tuple with the LifetimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeUnit

`func (o *Lifetime) SetLifetimeUnit(v string)`

SetLifetimeUnit sets LifetimeUnit field to given value.

### HasLifetimeUnit

`func (o *Lifetime) HasLifetimeUnit() bool`

HasLifetimeUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


