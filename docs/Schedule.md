# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Minute** | Pointer to **string** |  | [optional] [default to "00"]
**Hour** | Pointer to **string** |  | [optional] [default to "*"]
**Dom** | Pointer to **string** |  | [optional] [default to "*"]
**Month** | Pointer to **string** |  | [optional] [default to "*"]
**Dow** | Pointer to **string** |  | [optional] [default to "*"]

## Methods

### NewSchedule

`func NewSchedule() *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinute

`func (o *Schedule) GetMinute() string`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *Schedule) GetMinuteOk() (*string, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *Schedule) SetMinute(v string)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *Schedule) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetHour

`func (o *Schedule) GetHour() string`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *Schedule) GetHourOk() (*string, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *Schedule) SetHour(v string)`

SetHour sets Hour field to given value.

### HasHour

`func (o *Schedule) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetDom

`func (o *Schedule) GetDom() string`

GetDom returns the Dom field if non-nil, zero value otherwise.

### GetDomOk

`func (o *Schedule) GetDomOk() (*string, bool)`

GetDomOk returns a tuple with the Dom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDom

`func (o *Schedule) SetDom(v string)`

SetDom sets Dom field to given value.

### HasDom

`func (o *Schedule) HasDom() bool`

HasDom returns a boolean if a field has been set.

### GetMonth

`func (o *Schedule) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *Schedule) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *Schedule) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *Schedule) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetDow

`func (o *Schedule) GetDow() string`

GetDow returns the Dow field if non-nil, zero value otherwise.

### GetDowOk

`func (o *Schedule) GetDowOk() (*string, bool)`

GetDowOk returns a tuple with the Dow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDow

`func (o *Schedule) SetDow(v string)`

SetDow sets Dow field to given value.

### HasDow

`func (o *Schedule) HasDow() bool`

HasDow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


