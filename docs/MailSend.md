# MailSend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailMessage** | Pointer to [**MailSend0**](MailSend0.md) |  | [optional] 
**MailEntry** | Pointer to [**MailSend1**](MailSend1.md) |  | [optional] 

## Methods

### NewMailSend

`func NewMailSend() *MailSend`

NewMailSend instantiates a new MailSend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailSendWithDefaults

`func NewMailSendWithDefaults() *MailSend`

NewMailSendWithDefaults instantiates a new MailSend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailMessage

`func (o *MailSend) GetMailMessage() MailSend0`

GetMailMessage returns the MailMessage field if non-nil, zero value otherwise.

### GetMailMessageOk

`func (o *MailSend) GetMailMessageOk() (*MailSend0, bool)`

GetMailMessageOk returns a tuple with the MailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailMessage

`func (o *MailSend) SetMailMessage(v MailSend0)`

SetMailMessage sets MailMessage field to given value.

### HasMailMessage

`func (o *MailSend) HasMailMessage() bool`

HasMailMessage returns a boolean if a field has been set.

### GetMailEntry

`func (o *MailSend) GetMailEntry() MailSend1`

GetMailEntry returns the MailEntry field if non-nil, zero value otherwise.

### GetMailEntryOk

`func (o *MailSend) GetMailEntryOk() (*MailSend1, bool)`

GetMailEntryOk returns a tuple with the MailEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailEntry

`func (o *MailSend) SetMailEntry(v MailSend1)`

SetMailEntry sets MailEntry field to given value.

### HasMailEntry

`func (o *MailSend) HasMailEntry() bool`

HasMailEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


