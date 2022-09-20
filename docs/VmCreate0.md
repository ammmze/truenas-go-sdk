# VmCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuMode** | Pointer to **string** |  | [optional] [default to "CUSTOM"]
**CpuModel** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Vcpus** | Pointer to **int32** | Maximum of 16 guest virtual CPUs are allowed. By default, every virtual CPU is configured as a separate package. Multiple cores can be configured per CPU by specifying &#x60;cores&#x60; attributes. &#x60;vcpus&#x60; specifies total number of CPU sockets. &#x60;cores&#x60; specifies number of cores per socket. &#x60;threads&#x60; specifies number of threads per core. | [optional] [default to 1]
**Cores** | Pointer to **int32** | Maximum of 16 guest virtual CPUs are allowed. By default, every virtual CPU is configured as a separate package. Multiple cores can be configured per CPU by specifying &#x60;cores&#x60; attributes. &#x60;vcpus&#x60; specifies total number of CPU sockets. &#x60;cores&#x60; specifies number of cores per socket. &#x60;threads&#x60; specifies number of threads per core. | [optional] [default to 1]
**Threads** | Pointer to **int32** | Maximum of 16 guest virtual CPUs are allowed. By default, every virtual CPU is configured as a separate package. Multiple cores can be configured per CPU by specifying &#x60;cores&#x60; attributes. &#x60;vcpus&#x60; specifies total number of CPU sockets. &#x60;cores&#x60; specifies number of cores per socket. &#x60;threads&#x60; specifies number of threads per core. | [optional] [default to 1]
**Memory** | Pointer to **int32** |  | [optional] 
**Bootloader** | Pointer to **string** |  | [optional] [default to "UEFI"]
**Devices** | Pointer to [**[]VmdeviceCreate**](VmdeviceCreate.md) | &#x60;devices&#x60; is a list of virtualized hardware to add to the newly created Virtual Machine. Failure to attach a device destroys the VM and any resources allocated by the VM devices. | [optional] [default to []]
**Autostart** | Pointer to **bool** |  | [optional] [default to true]
**HideFromMsr** | Pointer to **bool** | &#x60;hide_from_msr&#x60; is a boolean which when set will hide the KVM hypervisor from standard MSR based discovery and is useful to enable when doing GPU passthrough. | [optional] [default to false]
**EnsureDisplayDevice** | Pointer to **bool** | &#x60;ensure_display_device&#x60; when set ( the default ) will ensure that the guest always has access to a video device. For headless installations like ubuntu server this is required for the guest to operate properly. However for cases where consumer would like to use GPU passthrough and does not want a display device added should set this to &#x60;false&#x60;. | [optional] [default to true]
**Time** | Pointer to **string** |  | [optional] [default to "LOCAL"]
**ShutdownTimeout** | Pointer to **int32** | &#x60;shutdown_timeout&#x60; indicates the time in seconds the system waits for the VM to cleanly shutdown. During system shutdown, if the VM hasn&#39;t exited after a hardware shutdown signal has been sent by the system within &#x60;shutdown_timeout&#x60; seconds, system initiates poweroff for the VM to stop it. | [optional] [default to 90]
**ArchType** | Pointer to **NullableString** | &#x60;arch_type&#x60; refers to architecture type and can be specified for the guest. By default the value is &#x60;null&#x60; and system in this case will choose a reasonable default based on host. &#x60;machine_type&#x60; refers to machine type of the guest based on the architecture type selected with &#x60;arch_type&#x60;. By default the value is &#x60;null&#x60; and system in this case will choose a reasonable default based on &#x60;arch_type&#x60; configuration. | [optional] 
**MachineType** | Pointer to **NullableString** | &#x60;machine_type&#x60; refers to machine type of the guest based on the architecture type selected with &#x60;arch_type&#x60;. By default the value is &#x60;null&#x60; and system in this case will choose a reasonable default based on &#x60;arch_type&#x60; configuration. | [optional] 
**Uuid** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVmCreate0

`func NewVmCreate0() *VmCreate0`

NewVmCreate0 instantiates a new VmCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmCreate0WithDefaults

`func NewVmCreate0WithDefaults() *VmCreate0`

NewVmCreate0WithDefaults instantiates a new VmCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuMode

`func (o *VmCreate0) GetCpuMode() string`

GetCpuMode returns the CpuMode field if non-nil, zero value otherwise.

### GetCpuModeOk

`func (o *VmCreate0) GetCpuModeOk() (*string, bool)`

GetCpuModeOk returns a tuple with the CpuMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMode

`func (o *VmCreate0) SetCpuMode(v string)`

SetCpuMode sets CpuMode field to given value.

### HasCpuMode

`func (o *VmCreate0) HasCpuMode() bool`

HasCpuMode returns a boolean if a field has been set.

### GetCpuModel

`func (o *VmCreate0) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *VmCreate0) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *VmCreate0) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *VmCreate0) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### SetCpuModelNil

`func (o *VmCreate0) SetCpuModelNil(b bool)`

 SetCpuModelNil sets the value for CpuModel to be an explicit nil

### UnsetCpuModel
`func (o *VmCreate0) UnsetCpuModel()`

UnsetCpuModel ensures that no value is present for CpuModel, not even an explicit nil
### GetName

`func (o *VmCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VmCreate0) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VmCreate0) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VmCreate0) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VmCreate0) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVcpus

`func (o *VmCreate0) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *VmCreate0) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *VmCreate0) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *VmCreate0) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetCores

`func (o *VmCreate0) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *VmCreate0) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *VmCreate0) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *VmCreate0) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetThreads

`func (o *VmCreate0) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *VmCreate0) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *VmCreate0) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *VmCreate0) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetMemory

`func (o *VmCreate0) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VmCreate0) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VmCreate0) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *VmCreate0) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetBootloader

`func (o *VmCreate0) GetBootloader() string`

GetBootloader returns the Bootloader field if non-nil, zero value otherwise.

### GetBootloaderOk

`func (o *VmCreate0) GetBootloaderOk() (*string, bool)`

GetBootloaderOk returns a tuple with the Bootloader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootloader

`func (o *VmCreate0) SetBootloader(v string)`

SetBootloader sets Bootloader field to given value.

### HasBootloader

`func (o *VmCreate0) HasBootloader() bool`

HasBootloader returns a boolean if a field has been set.

### GetDevices

`func (o *VmCreate0) GetDevices() []VmdeviceCreate`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *VmCreate0) GetDevicesOk() (*[]VmdeviceCreate, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *VmCreate0) SetDevices(v []VmdeviceCreate)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *VmCreate0) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetAutostart

`func (o *VmCreate0) GetAutostart() bool`

GetAutostart returns the Autostart field if non-nil, zero value otherwise.

### GetAutostartOk

`func (o *VmCreate0) GetAutostartOk() (*bool, bool)`

GetAutostartOk returns a tuple with the Autostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostart

`func (o *VmCreate0) SetAutostart(v bool)`

SetAutostart sets Autostart field to given value.

### HasAutostart

`func (o *VmCreate0) HasAutostart() bool`

HasAutostart returns a boolean if a field has been set.

### GetHideFromMsr

`func (o *VmCreate0) GetHideFromMsr() bool`

GetHideFromMsr returns the HideFromMsr field if non-nil, zero value otherwise.

### GetHideFromMsrOk

`func (o *VmCreate0) GetHideFromMsrOk() (*bool, bool)`

GetHideFromMsrOk returns a tuple with the HideFromMsr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromMsr

`func (o *VmCreate0) SetHideFromMsr(v bool)`

SetHideFromMsr sets HideFromMsr field to given value.

### HasHideFromMsr

`func (o *VmCreate0) HasHideFromMsr() bool`

HasHideFromMsr returns a boolean if a field has been set.

### GetEnsureDisplayDevice

`func (o *VmCreate0) GetEnsureDisplayDevice() bool`

GetEnsureDisplayDevice returns the EnsureDisplayDevice field if non-nil, zero value otherwise.

### GetEnsureDisplayDeviceOk

`func (o *VmCreate0) GetEnsureDisplayDeviceOk() (*bool, bool)`

GetEnsureDisplayDeviceOk returns a tuple with the EnsureDisplayDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsureDisplayDevice

`func (o *VmCreate0) SetEnsureDisplayDevice(v bool)`

SetEnsureDisplayDevice sets EnsureDisplayDevice field to given value.

### HasEnsureDisplayDevice

`func (o *VmCreate0) HasEnsureDisplayDevice() bool`

HasEnsureDisplayDevice returns a boolean if a field has been set.

### GetTime

`func (o *VmCreate0) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *VmCreate0) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *VmCreate0) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *VmCreate0) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetShutdownTimeout

`func (o *VmCreate0) GetShutdownTimeout() int32`

GetShutdownTimeout returns the ShutdownTimeout field if non-nil, zero value otherwise.

### GetShutdownTimeoutOk

`func (o *VmCreate0) GetShutdownTimeoutOk() (*int32, bool)`

GetShutdownTimeoutOk returns a tuple with the ShutdownTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownTimeout

`func (o *VmCreate0) SetShutdownTimeout(v int32)`

SetShutdownTimeout sets ShutdownTimeout field to given value.

### HasShutdownTimeout

`func (o *VmCreate0) HasShutdownTimeout() bool`

HasShutdownTimeout returns a boolean if a field has been set.

### GetArchType

`func (o *VmCreate0) GetArchType() string`

GetArchType returns the ArchType field if non-nil, zero value otherwise.

### GetArchTypeOk

`func (o *VmCreate0) GetArchTypeOk() (*string, bool)`

GetArchTypeOk returns a tuple with the ArchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchType

`func (o *VmCreate0) SetArchType(v string)`

SetArchType sets ArchType field to given value.

### HasArchType

`func (o *VmCreate0) HasArchType() bool`

HasArchType returns a boolean if a field has been set.

### SetArchTypeNil

`func (o *VmCreate0) SetArchTypeNil(b bool)`

 SetArchTypeNil sets the value for ArchType to be an explicit nil

### UnsetArchType
`func (o *VmCreate0) UnsetArchType()`

UnsetArchType ensures that no value is present for ArchType, not even an explicit nil
### GetMachineType

`func (o *VmCreate0) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *VmCreate0) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *VmCreate0) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *VmCreate0) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### SetMachineTypeNil

`func (o *VmCreate0) SetMachineTypeNil(b bool)`

 SetMachineTypeNil sets the value for MachineType to be an explicit nil

### UnsetMachineType
`func (o *VmCreate0) UnsetMachineType()`

UnsetMachineType ensures that no value is present for MachineType, not even an explicit nil
### GetUuid

`func (o *VmCreate0) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VmCreate0) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VmCreate0) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VmCreate0) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *VmCreate0) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *VmCreate0) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


