# NodeFullBios

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | BIOS name | [optional] 
**Version** | Pointer to **string** | BIOS version | [optional] 
**Editor** | Pointer to **string** | BIOS editor | [optional] 
**Quantity** | Pointer to **int32** | Number of BIOS on the machine | [optional] 
**ReleaseDate** | Pointer to **string** | Release date of the BIOS | [optional] 
**Description** | Pointer to **string** | System provided description of the BIOS (long name) | [optional] 

## Methods

### NewNodeFullBios

`func NewNodeFullBios() *NodeFullBios`

NewNodeFullBios instantiates a new NodeFullBios object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFullBiosWithDefaults

`func NewNodeFullBiosWithDefaults() *NodeFullBios`

NewNodeFullBiosWithDefaults instantiates a new NodeFullBios object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeFullBios) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeFullBios) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeFullBios) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeFullBios) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *NodeFullBios) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NodeFullBios) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NodeFullBios) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NodeFullBios) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEditor

`func (o *NodeFullBios) GetEditor() string`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *NodeFullBios) GetEditorOk() (*string, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *NodeFullBios) SetEditor(v string)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *NodeFullBios) HasEditor() bool`

HasEditor returns a boolean if a field has been set.

### GetQuantity

`func (o *NodeFullBios) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *NodeFullBios) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *NodeFullBios) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *NodeFullBios) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReleaseDate

`func (o *NodeFullBios) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *NodeFullBios) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *NodeFullBios) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *NodeFullBios) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetDescription

`func (o *NodeFullBios) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NodeFullBios) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NodeFullBios) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NodeFullBios) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


