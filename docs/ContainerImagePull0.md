# ContainerImagePull0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DockerAuthentication** | Pointer to [**DockerAuthentication**](DockerAuthentication.md) |  | [optional] 
**FromImage** | Pointer to **string** | &#x60;from_image&#x60; is the name of the image to pull. Format for the name is \&quot;registry/repo/image\&quot; where registry may be omitted and it will default to docker registry in this case. | [optional] 
**Tag** | Pointer to **NullableString** | &#x60;tag&#x60; specifies tag of the image and defaults to &#x60;null&#x60;. In case of &#x60;null&#x60; it will retrieve all the tags of the image. | [optional] 

## Methods

### NewContainerImagePull0

`func NewContainerImagePull0() *ContainerImagePull0`

NewContainerImagePull0 instantiates a new ContainerImagePull0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerImagePull0WithDefaults

`func NewContainerImagePull0WithDefaults() *ContainerImagePull0`

NewContainerImagePull0WithDefaults instantiates a new ContainerImagePull0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDockerAuthentication

`func (o *ContainerImagePull0) GetDockerAuthentication() DockerAuthentication`

GetDockerAuthentication returns the DockerAuthentication field if non-nil, zero value otherwise.

### GetDockerAuthenticationOk

`func (o *ContainerImagePull0) GetDockerAuthenticationOk() (*DockerAuthentication, bool)`

GetDockerAuthenticationOk returns a tuple with the DockerAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerAuthentication

`func (o *ContainerImagePull0) SetDockerAuthentication(v DockerAuthentication)`

SetDockerAuthentication sets DockerAuthentication field to given value.

### HasDockerAuthentication

`func (o *ContainerImagePull0) HasDockerAuthentication() bool`

HasDockerAuthentication returns a boolean if a field has been set.

### GetFromImage

`func (o *ContainerImagePull0) GetFromImage() string`

GetFromImage returns the FromImage field if non-nil, zero value otherwise.

### GetFromImageOk

`func (o *ContainerImagePull0) GetFromImageOk() (*string, bool)`

GetFromImageOk returns a tuple with the FromImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromImage

`func (o *ContainerImagePull0) SetFromImage(v string)`

SetFromImage sets FromImage field to given value.

### HasFromImage

`func (o *ContainerImagePull0) HasFromImage() bool`

HasFromImage returns a boolean if a field has been set.

### GetTag

`func (o *ContainerImagePull0) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ContainerImagePull0) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ContainerImagePull0) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ContainerImagePull0) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *ContainerImagePull0) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *ContainerImagePull0) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


