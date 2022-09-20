# Options3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseKerberos** | Pointer to **bool** |  | [optional] [default to false]
**OutputFormat** | Pointer to **string** |  | [optional] [default to "SMB"]

## Methods

### NewOptions3

`func NewOptions3() *Options3`

NewOptions3 instantiates a new Options3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptions3WithDefaults

`func NewOptions3WithDefaults() *Options3`

NewOptions3WithDefaults instantiates a new Options3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseKerberos

`func (o *Options3) GetUseKerberos() bool`

GetUseKerberos returns the UseKerberos field if non-nil, zero value otherwise.

### GetUseKerberosOk

`func (o *Options3) GetUseKerberosOk() (*bool, bool)`

GetUseKerberosOk returns a tuple with the UseKerberos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseKerberos

`func (o *Options3) SetUseKerberos(v bool)`

SetUseKerberos sets UseKerberos field to given value.

### HasUseKerberos

`func (o *Options3) HasUseKerberos() bool`

HasUseKerberos returns a boolean if a field has been set.

### GetOutputFormat

`func (o *Options3) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *Options3) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *Options3) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *Options3) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


