# NodeFullProcessesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pid** | Pointer to **int32** | PID of the process | [optional] 
**Tty** | Pointer to **string** | TTY to which the process is | [optional] 
**Name** | Pointer to **string** | Process name | [optional] 
**User** | Pointer to **string** | User account who started the process | [optional] 
**Started** | Pointer to **string** | Started date and time of the process | [optional] 
**Memory** | Pointer to **float32** | Memory allocated to the process (at inventory time) | [optional] 
**VirtualMemory** | Pointer to **int32** | Virtual memory allocated to the process (at inventory time) | [optional] 
**CpuUsage** | Pointer to **int32** | CPU used by the process (at inventory time) | [optional] 
**Description** | Pointer to **string** | System provided description about the process | [optional] 

## Methods

### NewNodeFullProcessesInner

`func NewNodeFullProcessesInner() *NodeFullProcessesInner`

NewNodeFullProcessesInner instantiates a new NodeFullProcessesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFullProcessesInnerWithDefaults

`func NewNodeFullProcessesInnerWithDefaults() *NodeFullProcessesInner`

NewNodeFullProcessesInnerWithDefaults instantiates a new NodeFullProcessesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPid

`func (o *NodeFullProcessesInner) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *NodeFullProcessesInner) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *NodeFullProcessesInner) SetPid(v int32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *NodeFullProcessesInner) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetTty

`func (o *NodeFullProcessesInner) GetTty() string`

GetTty returns the Tty field if non-nil, zero value otherwise.

### GetTtyOk

`func (o *NodeFullProcessesInner) GetTtyOk() (*string, bool)`

GetTtyOk returns a tuple with the Tty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTty

`func (o *NodeFullProcessesInner) SetTty(v string)`

SetTty sets Tty field to given value.

### HasTty

`func (o *NodeFullProcessesInner) HasTty() bool`

HasTty returns a boolean if a field has been set.

### GetName

`func (o *NodeFullProcessesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeFullProcessesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeFullProcessesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeFullProcessesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUser

`func (o *NodeFullProcessesInner) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *NodeFullProcessesInner) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *NodeFullProcessesInner) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *NodeFullProcessesInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetStarted

`func (o *NodeFullProcessesInner) GetStarted() string`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *NodeFullProcessesInner) GetStartedOk() (*string, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *NodeFullProcessesInner) SetStarted(v string)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *NodeFullProcessesInner) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetMemory

`func (o *NodeFullProcessesInner) GetMemory() float32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *NodeFullProcessesInner) GetMemoryOk() (*float32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *NodeFullProcessesInner) SetMemory(v float32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *NodeFullProcessesInner) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetVirtualMemory

`func (o *NodeFullProcessesInner) GetVirtualMemory() int32`

GetVirtualMemory returns the VirtualMemory field if non-nil, zero value otherwise.

### GetVirtualMemoryOk

`func (o *NodeFullProcessesInner) GetVirtualMemoryOk() (*int32, bool)`

GetVirtualMemoryOk returns a tuple with the VirtualMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMemory

`func (o *NodeFullProcessesInner) SetVirtualMemory(v int32)`

SetVirtualMemory sets VirtualMemory field to given value.

### HasVirtualMemory

`func (o *NodeFullProcessesInner) HasVirtualMemory() bool`

HasVirtualMemory returns a boolean if a field has been set.

### GetCpuUsage

`func (o *NodeFullProcessesInner) GetCpuUsage() int32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *NodeFullProcessesInner) GetCpuUsageOk() (*int32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *NodeFullProcessesInner) SetCpuUsage(v int32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *NodeFullProcessesInner) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetDescription

`func (o *NodeFullProcessesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NodeFullProcessesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NodeFullProcessesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NodeFullProcessesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


