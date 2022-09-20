# Topology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Datavdevs**](Datavdevs.md) |  | [optional] [default to []]
**Special** | Pointer to [**[]Specialvdevs**](Specialvdevs.md) |  | [optional] [default to []]
**Dedup** | Pointer to [**[]Dedupvdevs**](Dedupvdevs.md) |  | [optional] [default to []]
**Cache** | Pointer to [**[]Cachevdevs**](Cachevdevs.md) |  | [optional] [default to []]
**Log** | Pointer to [**[]Logvdevs**](Logvdevs.md) |  | [optional] [default to []]
**Spares** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewTopology

`func NewTopology() *Topology`

NewTopology instantiates a new Topology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyWithDefaults

`func NewTopologyWithDefaults() *Topology`

NewTopologyWithDefaults instantiates a new Topology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Topology) GetData() []Datavdevs`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Topology) GetDataOk() (*[]Datavdevs, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Topology) SetData(v []Datavdevs)`

SetData sets Data field to given value.

### HasData

`func (o *Topology) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSpecial

`func (o *Topology) GetSpecial() []Specialvdevs`

GetSpecial returns the Special field if non-nil, zero value otherwise.

### GetSpecialOk

`func (o *Topology) GetSpecialOk() (*[]Specialvdevs, bool)`

GetSpecialOk returns a tuple with the Special field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecial

`func (o *Topology) SetSpecial(v []Specialvdevs)`

SetSpecial sets Special field to given value.

### HasSpecial

`func (o *Topology) HasSpecial() bool`

HasSpecial returns a boolean if a field has been set.

### GetDedup

`func (o *Topology) GetDedup() []Dedupvdevs`

GetDedup returns the Dedup field if non-nil, zero value otherwise.

### GetDedupOk

`func (o *Topology) GetDedupOk() (*[]Dedupvdevs, bool)`

GetDedupOk returns a tuple with the Dedup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedup

`func (o *Topology) SetDedup(v []Dedupvdevs)`

SetDedup sets Dedup field to given value.

### HasDedup

`func (o *Topology) HasDedup() bool`

HasDedup returns a boolean if a field has been set.

### GetCache

`func (o *Topology) GetCache() []Cachevdevs`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *Topology) GetCacheOk() (*[]Cachevdevs, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *Topology) SetCache(v []Cachevdevs)`

SetCache sets Cache field to given value.

### HasCache

`func (o *Topology) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetLog

`func (o *Topology) GetLog() []Logvdevs`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *Topology) GetLogOk() (*[]Logvdevs, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *Topology) SetLog(v []Logvdevs)`

SetLog sets Log field to given value.

### HasLog

`func (o *Topology) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetSpares

`func (o *Topology) GetSpares() []string`

GetSpares returns the Spares field if non-nil, zero value otherwise.

### GetSparesOk

`func (o *Topology) GetSparesOk() (*[]string, bool)`

GetSparesOk returns a tuple with the Spares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpares

`func (o *Topology) SetSpares(v []string)`

SetSpares sets Spares field to given value.

### HasSpares

`func (o *Topology) HasSpares() bool`

HasSpares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


