# NodeFullSlotsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Slot name or identifier | [optional] 
**Status** | Pointer to **string** | Slot status (used, powered, etc) | [optional] 
**Quantity** | Pointer to **int32** | Quantity of similar slots | [optional] 
**Description** | Pointer to **string** | System provided description of the slots | [optional] 

## Methods

### NewNodeFullSlotsInner

`func NewNodeFullSlotsInner() *NodeFullSlotsInner`

NewNodeFullSlotsInner instantiates a new NodeFullSlotsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFullSlotsInnerWithDefaults

`func NewNodeFullSlotsInnerWithDefaults() *NodeFullSlotsInner`

NewNodeFullSlotsInnerWithDefaults instantiates a new NodeFullSlotsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeFullSlotsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeFullSlotsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeFullSlotsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeFullSlotsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *NodeFullSlotsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeFullSlotsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeFullSlotsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NodeFullSlotsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetQuantity

`func (o *NodeFullSlotsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *NodeFullSlotsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *NodeFullSlotsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *NodeFullSlotsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetDescription

`func (o *NodeFullSlotsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NodeFullSlotsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NodeFullSlotsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NodeFullSlotsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


