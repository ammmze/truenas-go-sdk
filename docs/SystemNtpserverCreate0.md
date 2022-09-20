# SystemNtpserverCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | &#x60;address&#x60; specifies the hostname/IP address of the NTP server. | [optional] 
**Burst** | Pointer to **bool** | &#x60;burst&#x60; when enabled makes sure that if server is reachable, sends a burst of eight packets instead of one. This is designed to improve timekeeping quality with the server command. | [optional] [default to false]
**Iburst** | Pointer to **bool** | &#x60;iburst&#x60; when enabled speeds up the initial synchronization, taking seconds rather than minutes. | [optional] [default to true]
**Prefer** | Pointer to **bool** | &#x60;prefer&#x60; marks the specified server as preferred. When all other things are equal, this host is chosen for synchronization acquisition with the server command. It is recommended that they be used for servers with time monitoring hardware. | [optional] [default to false]
**Minpoll** | Pointer to **int32** | &#x60;minpoll&#x60; is minimum polling time in seconds. It must be a power of 2 and less than &#x60;maxpoll&#x60;. &#x60;maxpoll&#x60; is maximum polling time in seconds. It must be a power of 2 and greater than &#x60;minpoll&#x60;. | [optional] [default to 6]
**Maxpoll** | Pointer to **int32** | &#x60;minpoll&#x60; is minimum polling time in seconds. It must be a power of 2 and less than &#x60;maxpoll&#x60;. &#x60;maxpoll&#x60; is maximum polling time in seconds. It must be a power of 2 and greater than &#x60;minpoll&#x60;. | [optional] [default to 10]
**Force** | Pointer to **bool** |  | [optional] 

## Methods

### NewSystemNtpserverCreate0

`func NewSystemNtpserverCreate0() *SystemNtpserverCreate0`

NewSystemNtpserverCreate0 instantiates a new SystemNtpserverCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemNtpserverCreate0WithDefaults

`func NewSystemNtpserverCreate0WithDefaults() *SystemNtpserverCreate0`

NewSystemNtpserverCreate0WithDefaults instantiates a new SystemNtpserverCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SystemNtpserverCreate0) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SystemNtpserverCreate0) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SystemNtpserverCreate0) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SystemNtpserverCreate0) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBurst

`func (o *SystemNtpserverCreate0) GetBurst() bool`

GetBurst returns the Burst field if non-nil, zero value otherwise.

### GetBurstOk

`func (o *SystemNtpserverCreate0) GetBurstOk() (*bool, bool)`

GetBurstOk returns a tuple with the Burst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurst

`func (o *SystemNtpserverCreate0) SetBurst(v bool)`

SetBurst sets Burst field to given value.

### HasBurst

`func (o *SystemNtpserverCreate0) HasBurst() bool`

HasBurst returns a boolean if a field has been set.

### GetIburst

`func (o *SystemNtpserverCreate0) GetIburst() bool`

GetIburst returns the Iburst field if non-nil, zero value otherwise.

### GetIburstOk

`func (o *SystemNtpserverCreate0) GetIburstOk() (*bool, bool)`

GetIburstOk returns a tuple with the Iburst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIburst

`func (o *SystemNtpserverCreate0) SetIburst(v bool)`

SetIburst sets Iburst field to given value.

### HasIburst

`func (o *SystemNtpserverCreate0) HasIburst() bool`

HasIburst returns a boolean if a field has been set.

### GetPrefer

`func (o *SystemNtpserverCreate0) GetPrefer() bool`

GetPrefer returns the Prefer field if non-nil, zero value otherwise.

### GetPreferOk

`func (o *SystemNtpserverCreate0) GetPreferOk() (*bool, bool)`

GetPreferOk returns a tuple with the Prefer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefer

`func (o *SystemNtpserverCreate0) SetPrefer(v bool)`

SetPrefer sets Prefer field to given value.

### HasPrefer

`func (o *SystemNtpserverCreate0) HasPrefer() bool`

HasPrefer returns a boolean if a field has been set.

### GetMinpoll

`func (o *SystemNtpserverCreate0) GetMinpoll() int32`

GetMinpoll returns the Minpoll field if non-nil, zero value otherwise.

### GetMinpollOk

`func (o *SystemNtpserverCreate0) GetMinpollOk() (*int32, bool)`

GetMinpollOk returns a tuple with the Minpoll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinpoll

`func (o *SystemNtpserverCreate0) SetMinpoll(v int32)`

SetMinpoll sets Minpoll field to given value.

### HasMinpoll

`func (o *SystemNtpserverCreate0) HasMinpoll() bool`

HasMinpoll returns a boolean if a field has been set.

### GetMaxpoll

`func (o *SystemNtpserverCreate0) GetMaxpoll() int32`

GetMaxpoll returns the Maxpoll field if non-nil, zero value otherwise.

### GetMaxpollOk

`func (o *SystemNtpserverCreate0) GetMaxpollOk() (*int32, bool)`

GetMaxpollOk returns a tuple with the Maxpoll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxpoll

`func (o *SystemNtpserverCreate0) SetMaxpoll(v int32)`

SetMaxpoll sets Maxpoll field to given value.

### HasMaxpoll

`func (o *SystemNtpserverCreate0) HasMaxpoll() bool`

HasMaxpoll returns a boolean if a field has been set.

### GetForce

`func (o *SystemNtpserverCreate0) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *SystemNtpserverCreate0) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *SystemNtpserverCreate0) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *SystemNtpserverCreate0) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


