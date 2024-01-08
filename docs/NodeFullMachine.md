# NodeFullMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Rudder unique identifier for the machine | [optional] 
**Type** | Pointer to **string** | Type of the machine | [optional] 
**Provider** | Pointer to **string** | In the case of VM, the VM technology | [optional] 
**Manufacturer** | Pointer to **string** | Information about machine manufacturer | [optional] 
**SerialNumber** | Pointer to **string** | If available, a unique identifier provided by the machine | [optional] 

## Methods

### NewNodeFullMachine

`func NewNodeFullMachine() *NodeFullMachine`

NewNodeFullMachine instantiates a new NodeFullMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFullMachineWithDefaults

`func NewNodeFullMachineWithDefaults() *NodeFullMachine`

NewNodeFullMachineWithDefaults instantiates a new NodeFullMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeFullMachine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeFullMachine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeFullMachine) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NodeFullMachine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NodeFullMachine) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeFullMachine) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeFullMachine) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NodeFullMachine) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProvider

`func (o *NodeFullMachine) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NodeFullMachine) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NodeFullMachine) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NodeFullMachine) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetManufacturer

`func (o *NodeFullMachine) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *NodeFullMachine) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *NodeFullMachine) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *NodeFullMachine) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetSerialNumber

`func (o *NodeFullMachine) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NodeFullMachine) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NodeFullMachine) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NodeFullMachine) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


