# Brick

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerName** | Pointer to **string** |  | [optional] 
**PeerPath** | Pointer to **string** |  | [optional] 

## Methods

### NewBrick

`func NewBrick() *Brick`

NewBrick instantiates a new Brick object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrickWithDefaults

`func NewBrickWithDefaults() *Brick`

NewBrickWithDefaults instantiates a new Brick object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerName

`func (o *Brick) GetPeerName() string`

GetPeerName returns the PeerName field if non-nil, zero value otherwise.

### GetPeerNameOk

`func (o *Brick) GetPeerNameOk() (*string, bool)`

GetPeerNameOk returns a tuple with the PeerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerName

`func (o *Brick) SetPeerName(v string)`

SetPeerName sets PeerName field to given value.

### HasPeerName

`func (o *Brick) HasPeerName() bool`

HasPeerName returns a boolean if a field has been set.

### GetPeerPath

`func (o *Brick) GetPeerPath() string`

GetPeerPath returns the PeerPath field if non-nil, zero value otherwise.

### GetPeerPathOk

`func (o *Brick) GetPeerPathOk() (*string, bool)`

GetPeerPathOk returns a tuple with the PeerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerPath

`func (o *Brick) SetPeerPath(v string)`

SetPeerPath sets PeerPath field to given value.

### HasPeerPath

`func (o *Brick) HasPeerPath() bool`

HasPeerPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


