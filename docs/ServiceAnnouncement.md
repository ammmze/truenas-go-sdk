# ServiceAnnouncement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Netbios** | Pointer to **bool** |  | [optional] 
**Mdns** | Pointer to **bool** |  | [optional] 
**Wsd** | Pointer to **bool** |  | [optional] 

## Methods

### NewServiceAnnouncement

`func NewServiceAnnouncement() *ServiceAnnouncement`

NewServiceAnnouncement instantiates a new ServiceAnnouncement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAnnouncementWithDefaults

`func NewServiceAnnouncementWithDefaults() *ServiceAnnouncement`

NewServiceAnnouncementWithDefaults instantiates a new ServiceAnnouncement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetbios

`func (o *ServiceAnnouncement) GetNetbios() bool`

GetNetbios returns the Netbios field if non-nil, zero value otherwise.

### GetNetbiosOk

`func (o *ServiceAnnouncement) GetNetbiosOk() (*bool, bool)`

GetNetbiosOk returns a tuple with the Netbios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbios

`func (o *ServiceAnnouncement) SetNetbios(v bool)`

SetNetbios sets Netbios field to given value.

### HasNetbios

`func (o *ServiceAnnouncement) HasNetbios() bool`

HasNetbios returns a boolean if a field has been set.

### GetMdns

`func (o *ServiceAnnouncement) GetMdns() bool`

GetMdns returns the Mdns field if non-nil, zero value otherwise.

### GetMdnsOk

`func (o *ServiceAnnouncement) GetMdnsOk() (*bool, bool)`

GetMdnsOk returns a tuple with the Mdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdns

`func (o *ServiceAnnouncement) SetMdns(v bool)`

SetMdns sets Mdns field to given value.

### HasMdns

`func (o *ServiceAnnouncement) HasMdns() bool`

HasMdns returns a boolean if a field has been set.

### GetWsd

`func (o *ServiceAnnouncement) GetWsd() bool`

GetWsd returns the Wsd field if non-nil, zero value otherwise.

### GetWsdOk

`func (o *ServiceAnnouncement) GetWsdOk() (*bool, bool)`

GetWsdOk returns a tuple with the Wsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsd

`func (o *ServiceAnnouncement) SetWsd(v bool)`

SetWsd sets Wsd field to given value.

### HasWsd

`func (o *ServiceAnnouncement) HasWsd() bool`

HasWsd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


