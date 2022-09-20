# KubernetesRestoreBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupName** | Pointer to **string** | Restore &#x60;backup_name&#x60; chart releases backup. | [optional] 
**Options** | Pointer to [**KubernetesRestoreBackup1**](KubernetesRestoreBackup1.md) |  | [optional] 

## Methods

### NewKubernetesRestoreBackup

`func NewKubernetesRestoreBackup() *KubernetesRestoreBackup`

NewKubernetesRestoreBackup instantiates a new KubernetesRestoreBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesRestoreBackupWithDefaults

`func NewKubernetesRestoreBackupWithDefaults() *KubernetesRestoreBackup`

NewKubernetesRestoreBackupWithDefaults instantiates a new KubernetesRestoreBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupName

`func (o *KubernetesRestoreBackup) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *KubernetesRestoreBackup) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *KubernetesRestoreBackup) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *KubernetesRestoreBackup) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### GetOptions

`func (o *KubernetesRestoreBackup) GetOptions() KubernetesRestoreBackup1`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *KubernetesRestoreBackup) GetOptionsOk() (*KubernetesRestoreBackup1, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *KubernetesRestoreBackup) SetOptions(v KubernetesRestoreBackup1)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *KubernetesRestoreBackup) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


