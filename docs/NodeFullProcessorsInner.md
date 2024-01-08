# NodeFullProcessorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | CPU name | [optional] 
**Arch** | Pointer to **string** | CPU architecture | [optional] 
**Model** | Pointer to **int32** | CPU model | [optional] 
**FamilyName** | Pointer to **string** | CPU family | [optional] 
**Core** | Pointer to **int32** | Number of core for that CPU | [optional] 
**Speed** | Pointer to **int32** | Speed (frequency) of the CPU | [optional] 
**Thread** | Pointer to **int32** | Number of thread by core for the CPU | [optional] 
**Stepping** | Pointer to **int32** | Stepping or power management information | [optional] 
**Manufacturer** | Pointer to **string** | CPU manufacturer | [optional] 
**Quantity** | Pointer to **int32** | Number of CPU with these features | [optional] 
**Cpuid** | Pointer to **string** | Identifier of the CPU | [optional] 
**ExternalClock** | Pointer to **string** | External clock used by the CPU | [optional] 
**Description** | Pointer to **string** | System provided description of the CPU | [optional] 

## Methods

### NewNodeFullProcessorsInner

`func NewNodeFullProcessorsInner() *NodeFullProcessorsInner`

NewNodeFullProcessorsInner instantiates a new NodeFullProcessorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFullProcessorsInnerWithDefaults

`func NewNodeFullProcessorsInnerWithDefaults() *NodeFullProcessorsInner`

NewNodeFullProcessorsInnerWithDefaults instantiates a new NodeFullProcessorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeFullProcessorsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeFullProcessorsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeFullProcessorsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeFullProcessorsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArch

`func (o *NodeFullProcessorsInner) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *NodeFullProcessorsInner) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *NodeFullProcessorsInner) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *NodeFullProcessorsInner) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetModel

`func (o *NodeFullProcessorsInner) GetModel() int32`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NodeFullProcessorsInner) GetModelOk() (*int32, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NodeFullProcessorsInner) SetModel(v int32)`

SetModel sets Model field to given value.

### HasModel

`func (o *NodeFullProcessorsInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetFamilyName

`func (o *NodeFullProcessorsInner) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *NodeFullProcessorsInner) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *NodeFullProcessorsInner) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *NodeFullProcessorsInner) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetCore

`func (o *NodeFullProcessorsInner) GetCore() int32`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *NodeFullProcessorsInner) GetCoreOk() (*int32, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *NodeFullProcessorsInner) SetCore(v int32)`

SetCore sets Core field to given value.

### HasCore

`func (o *NodeFullProcessorsInner) HasCore() bool`

HasCore returns a boolean if a field has been set.

### GetSpeed

`func (o *NodeFullProcessorsInner) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *NodeFullProcessorsInner) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *NodeFullProcessorsInner) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *NodeFullProcessorsInner) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetThread

`func (o *NodeFullProcessorsInner) GetThread() int32`

GetThread returns the Thread field if non-nil, zero value otherwise.

### GetThreadOk

`func (o *NodeFullProcessorsInner) GetThreadOk() (*int32, bool)`

GetThreadOk returns a tuple with the Thread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThread

`func (o *NodeFullProcessorsInner) SetThread(v int32)`

SetThread sets Thread field to given value.

### HasThread

`func (o *NodeFullProcessorsInner) HasThread() bool`

HasThread returns a boolean if a field has been set.

### GetStepping

`func (o *NodeFullProcessorsInner) GetStepping() int32`

GetStepping returns the Stepping field if non-nil, zero value otherwise.

### GetSteppingOk

`func (o *NodeFullProcessorsInner) GetSteppingOk() (*int32, bool)`

GetSteppingOk returns a tuple with the Stepping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepping

`func (o *NodeFullProcessorsInner) SetStepping(v int32)`

SetStepping sets Stepping field to given value.

### HasStepping

`func (o *NodeFullProcessorsInner) HasStepping() bool`

HasStepping returns a boolean if a field has been set.

### GetManufacturer

`func (o *NodeFullProcessorsInner) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *NodeFullProcessorsInner) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *NodeFullProcessorsInner) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *NodeFullProcessorsInner) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetQuantity

`func (o *NodeFullProcessorsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *NodeFullProcessorsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *NodeFullProcessorsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *NodeFullProcessorsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCpuid

`func (o *NodeFullProcessorsInner) GetCpuid() string`

GetCpuid returns the Cpuid field if non-nil, zero value otherwise.

### GetCpuidOk

`func (o *NodeFullProcessorsInner) GetCpuidOk() (*string, bool)`

GetCpuidOk returns a tuple with the Cpuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuid

`func (o *NodeFullProcessorsInner) SetCpuid(v string)`

SetCpuid sets Cpuid field to given value.

### HasCpuid

`func (o *NodeFullProcessorsInner) HasCpuid() bool`

HasCpuid returns a boolean if a field has been set.

### GetExternalClock

`func (o *NodeFullProcessorsInner) GetExternalClock() string`

GetExternalClock returns the ExternalClock field if non-nil, zero value otherwise.

### GetExternalClockOk

`func (o *NodeFullProcessorsInner) GetExternalClockOk() (*string, bool)`

GetExternalClockOk returns a tuple with the ExternalClock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalClock

`func (o *NodeFullProcessorsInner) SetExternalClock(v string)`

SetExternalClock sets ExternalClock field to given value.

### HasExternalClock

`func (o *NodeFullProcessorsInner) HasExternalClock() bool`

HasExternalClock returns a boolean if a field has been set.

### GetDescription

`func (o *NodeFullProcessorsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NodeFullProcessorsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NodeFullProcessorsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NodeFullProcessorsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


