# NodeFullMemoriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the memory slot or memory module | [optional] 
**Speed** | Pointer to **int32** | Memory speed (frequency) | [optional] 
**Type** | Pointer to **string** | Memory slot type | [optional] 
**Caption** | Pointer to **string** | Manufacturer provided information about the module | [optional] 
**Quantity** | Pointer to **int32** | Number of modules in that slot | [optional] 
**Capacity** | Pointer to **int32** | Size of modules | [optional] 
**SlotNumber** | Pointer to **int32** | Slot number | [optional] 
**Description** | Pointer to **string** | System provided description | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of the module | [optional] 

## Methods

### NewNodeFullMemoriesInner

`func NewNodeFullMemoriesInner() *NodeFullMemoriesInner`

NewNodeFullMemoriesInner instantiates a new NodeFullMemoriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFullMemoriesInnerWithDefaults

`func NewNodeFullMemoriesInnerWithDefaults() *NodeFullMemoriesInner`

NewNodeFullMemoriesInnerWithDefaults instantiates a new NodeFullMemoriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeFullMemoriesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeFullMemoriesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeFullMemoriesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeFullMemoriesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpeed

`func (o *NodeFullMemoriesInner) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *NodeFullMemoriesInner) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *NodeFullMemoriesInner) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *NodeFullMemoriesInner) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetType

`func (o *NodeFullMemoriesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeFullMemoriesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeFullMemoriesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NodeFullMemoriesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCaption

`func (o *NodeFullMemoriesInner) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *NodeFullMemoriesInner) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *NodeFullMemoriesInner) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *NodeFullMemoriesInner) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetQuantity

`func (o *NodeFullMemoriesInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *NodeFullMemoriesInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *NodeFullMemoriesInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *NodeFullMemoriesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCapacity

`func (o *NodeFullMemoriesInner) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *NodeFullMemoriesInner) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *NodeFullMemoriesInner) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *NodeFullMemoriesInner) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetSlotNumber

`func (o *NodeFullMemoriesInner) GetSlotNumber() int32`

GetSlotNumber returns the SlotNumber field if non-nil, zero value otherwise.

### GetSlotNumberOk

`func (o *NodeFullMemoriesInner) GetSlotNumberOk() (*int32, bool)`

GetSlotNumberOk returns a tuple with the SlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotNumber

`func (o *NodeFullMemoriesInner) SetSlotNumber(v int32)`

SetSlotNumber sets SlotNumber field to given value.

### HasSlotNumber

`func (o *NodeFullMemoriesInner) HasSlotNumber() bool`

HasSlotNumber returns a boolean if a field has been set.

### GetDescription

`func (o *NodeFullMemoriesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NodeFullMemoriesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NodeFullMemoriesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NodeFullMemoriesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSerialNumber

`func (o *NodeFullMemoriesInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NodeFullMemoriesInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NodeFullMemoriesInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NodeFullMemoriesInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


