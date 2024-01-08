# NodeFullOs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Family of the OS | 
**Name** | **string** | Operating system name (distribution on Linux, etc.) | 
**Version** | **string** | OS version | 
**FullName** | **string** | Full operating system name | 
**ServicePack** | Pointer to **string** | If relevant, the service pack of the OS | [optional] 
**KernelVersion** | **string** | Version of the OS kernel | 

## Methods

### NewNodeFullOs

`func NewNodeFullOs(type_ string, name string, version string, fullName string, kernelVersion string, ) *NodeFullOs`

NewNodeFullOs instantiates a new NodeFullOs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFullOsWithDefaults

`func NewNodeFullOsWithDefaults() *NodeFullOs`

NewNodeFullOsWithDefaults instantiates a new NodeFullOs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NodeFullOs) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeFullOs) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeFullOs) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *NodeFullOs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeFullOs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeFullOs) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *NodeFullOs) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NodeFullOs) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NodeFullOs) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetFullName

`func (o *NodeFullOs) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *NodeFullOs) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *NodeFullOs) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetServicePack

`func (o *NodeFullOs) GetServicePack() string`

GetServicePack returns the ServicePack field if non-nil, zero value otherwise.

### GetServicePackOk

`func (o *NodeFullOs) GetServicePackOk() (*string, bool)`

GetServicePackOk returns a tuple with the ServicePack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePack

`func (o *NodeFullOs) SetServicePack(v string)`

SetServicePack sets ServicePack field to given value.

### HasServicePack

`func (o *NodeFullOs) HasServicePack() bool`

HasServicePack returns a boolean if a field has been set.

### GetKernelVersion

`func (o *NodeFullOs) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *NodeFullOs) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *NodeFullOs) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


