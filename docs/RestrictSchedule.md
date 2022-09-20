# RestrictSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Minute** | Pointer to **string** |  | [optional] [default to "00"]
**Hour** | Pointer to **string** |  | [optional] [default to "*"]
**Dom** | Pointer to **string** |  | [optional] [default to "*"]
**Month** | Pointer to **string** |  | [optional] [default to "*"]
**Dow** | Pointer to **string** |  | [optional] [default to "*"]
**Begin** | Pointer to **string** |  | [optional] [default to "00:00"]
**End** | Pointer to **string** |  | [optional] [default to "1439"]

## Methods

### NewRestrictSchedule

`func NewRestrictSchedule() *RestrictSchedule`

NewRestrictSchedule instantiates a new RestrictSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictScheduleWithDefaults

`func NewRestrictScheduleWithDefaults() *RestrictSchedule`

NewRestrictScheduleWithDefaults instantiates a new RestrictSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinute

`func (o *RestrictSchedule) GetMinute() string`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *RestrictSchedule) GetMinuteOk() (*string, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *RestrictSchedule) SetMinute(v string)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *RestrictSchedule) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetHour

`func (o *RestrictSchedule) GetHour() string`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *RestrictSchedule) GetHourOk() (*string, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *RestrictSchedule) SetHour(v string)`

SetHour sets Hour field to given value.

### HasHour

`func (o *RestrictSchedule) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetDom

`func (o *RestrictSchedule) GetDom() string`

GetDom returns the Dom field if non-nil, zero value otherwise.

### GetDomOk

`func (o *RestrictSchedule) GetDomOk() (*string, bool)`

GetDomOk returns a tuple with the Dom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDom

`func (o *RestrictSchedule) SetDom(v string)`

SetDom sets Dom field to given value.

### HasDom

`func (o *RestrictSchedule) HasDom() bool`

HasDom returns a boolean if a field has been set.

### GetMonth

`func (o *RestrictSchedule) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *RestrictSchedule) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *RestrictSchedule) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *RestrictSchedule) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetDow

`func (o *RestrictSchedule) GetDow() string`

GetDow returns the Dow field if non-nil, zero value otherwise.

### GetDowOk

`func (o *RestrictSchedule) GetDowOk() (*string, bool)`

GetDowOk returns a tuple with the Dow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDow

`func (o *RestrictSchedule) SetDow(v string)`

SetDow sets Dow field to given value.

### HasDow

`func (o *RestrictSchedule) HasDow() bool`

HasDow returns a boolean if a field has been set.

### GetBegin

`func (o *RestrictSchedule) GetBegin() string`

GetBegin returns the Begin field if non-nil, zero value otherwise.

### GetBeginOk

`func (o *RestrictSchedule) GetBeginOk() (*string, bool)`

GetBeginOk returns a tuple with the Begin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBegin

`func (o *RestrictSchedule) SetBegin(v string)`

SetBegin sets Begin field to given value.

### HasBegin

`func (o *RestrictSchedule) HasBegin() bool`

HasBegin returns a boolean if a field has been set.

### GetEnd

`func (o *RestrictSchedule) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *RestrictSchedule) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *RestrictSchedule) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *RestrictSchedule) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


