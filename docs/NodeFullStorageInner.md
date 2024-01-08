# NodeFullStorageInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Storage name | [optional] 
**Type** | Pointer to **string** | Storage type | [optional] 
**Size** | Pointer to **int32** | Storage size in MB | [optional] 
**Model** | Pointer to **string** | Storage model | [optional] 
**Firmware** | Pointer to **string** | Storage firmware information | [optional] 
**Quantity** | Pointer to **int32** | Quantity of similar storage | [optional] 
**Description** | Pointer to **string** | System provided information about the storage | [optional] 
**Manufacturer** | Pointer to **string** | Storage manufacturer | [optional] 
**SerialNumber** | Pointer to **string** | Storage serial number | [optional] 

## Methods

### NewNodeFullStorageInner

`func NewNodeFullStorageInner() *NodeFullStorageInner`

NewNodeFullStorageInner instantiates a new NodeFullStorageInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFullStorageInnerWithDefaults

`func NewNodeFullStorageInnerWithDefaults() *NodeFullStorageInner`

NewNodeFullStorageInnerWithDefaults instantiates a new NodeFullStorageInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeFullStorageInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeFullStorageInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeFullStorageInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeFullStorageInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *NodeFullStorageInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeFullStorageInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeFullStorageInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NodeFullStorageInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *NodeFullStorageInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *NodeFullStorageInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *NodeFullStorageInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *NodeFullStorageInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetModel

`func (o *NodeFullStorageInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NodeFullStorageInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NodeFullStorageInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *NodeFullStorageInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetFirmware

`func (o *NodeFullStorageInner) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *NodeFullStorageInner) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *NodeFullStorageInner) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *NodeFullStorageInner) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetQuantity

`func (o *NodeFullStorageInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *NodeFullStorageInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *NodeFullStorageInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *NodeFullStorageInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetDescription

`func (o *NodeFullStorageInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NodeFullStorageInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NodeFullStorageInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NodeFullStorageInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetManufacturer

`func (o *NodeFullStorageInner) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *NodeFullStorageInner) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *NodeFullStorageInner) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *NodeFullStorageInner) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetSerialNumber

`func (o *NodeFullStorageInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NodeFullStorageInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NodeFullStorageInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NodeFullStorageInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


