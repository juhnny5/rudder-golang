# NodeFullVirtualMachinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the hosted virtual machine | [optional] 
**Type** | Pointer to **string** | Type of the hosted virtual machine | [optional] 
**Uuid** | Pointer to **string** | Unique identifier of the hosted virtual machine | [optional] 
**Vcpu** | Pointer to **string** | Number of virtual CPU allocated to the hosted virtual machine | [optional] 
**Owner** | Pointer to **string** | Owner of the hosted virtual machine | [optional] 
**Status** | Pointer to **string** | Status (up, starting, etc) of the hosted virtual machine | [optional] 
**Memory** | Pointer to **string** | Memory allocated to the hosted virtual machine | [optional] 
**Subsystem** | Pointer to **string** | Technology of the hosted virtual machine | [optional] 
**Description** | Pointer to **string** | System provided description of the hosted virtual machine | [optional] 

## Methods

### NewNodeFullVirtualMachinesInner

`func NewNodeFullVirtualMachinesInner() *NodeFullVirtualMachinesInner`

NewNodeFullVirtualMachinesInner instantiates a new NodeFullVirtualMachinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFullVirtualMachinesInnerWithDefaults

`func NewNodeFullVirtualMachinesInnerWithDefaults() *NodeFullVirtualMachinesInner`

NewNodeFullVirtualMachinesInnerWithDefaults instantiates a new NodeFullVirtualMachinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeFullVirtualMachinesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeFullVirtualMachinesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeFullVirtualMachinesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeFullVirtualMachinesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *NodeFullVirtualMachinesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeFullVirtualMachinesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeFullVirtualMachinesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NodeFullVirtualMachinesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *NodeFullVirtualMachinesInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NodeFullVirtualMachinesInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NodeFullVirtualMachinesInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NodeFullVirtualMachinesInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVcpu

`func (o *NodeFullVirtualMachinesInner) GetVcpu() string`

GetVcpu returns the Vcpu field if non-nil, zero value otherwise.

### GetVcpuOk

`func (o *NodeFullVirtualMachinesInner) GetVcpuOk() (*string, bool)`

GetVcpuOk returns a tuple with the Vcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpu

`func (o *NodeFullVirtualMachinesInner) SetVcpu(v string)`

SetVcpu sets Vcpu field to given value.

### HasVcpu

`func (o *NodeFullVirtualMachinesInner) HasVcpu() bool`

HasVcpu returns a boolean if a field has been set.

### GetOwner

`func (o *NodeFullVirtualMachinesInner) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NodeFullVirtualMachinesInner) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NodeFullVirtualMachinesInner) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NodeFullVirtualMachinesInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetStatus

`func (o *NodeFullVirtualMachinesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeFullVirtualMachinesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeFullVirtualMachinesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NodeFullVirtualMachinesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMemory

`func (o *NodeFullVirtualMachinesInner) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *NodeFullVirtualMachinesInner) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *NodeFullVirtualMachinesInner) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *NodeFullVirtualMachinesInner) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetSubsystem

`func (o *NodeFullVirtualMachinesInner) GetSubsystem() string`

GetSubsystem returns the Subsystem field if non-nil, zero value otherwise.

### GetSubsystemOk

`func (o *NodeFullVirtualMachinesInner) GetSubsystemOk() (*string, bool)`

GetSubsystemOk returns a tuple with the Subsystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsystem

`func (o *NodeFullVirtualMachinesInner) SetSubsystem(v string)`

SetSubsystem sets Subsystem field to given value.

### HasSubsystem

`func (o *NodeFullVirtualMachinesInner) HasSubsystem() bool`

HasSubsystem returns a boolean if a field has been set.

### GetDescription

`func (o *NodeFullVirtualMachinesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NodeFullVirtualMachinesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NodeFullVirtualMachinesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NodeFullVirtualMachinesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


