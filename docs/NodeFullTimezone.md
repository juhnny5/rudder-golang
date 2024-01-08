# NodeFullTimezone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Timezone name | 
**Offset** | Pointer to **string** | Timezone offset to UTC | [optional] 

## Methods

### NewNodeFullTimezone

`func NewNodeFullTimezone(name string, ) *NodeFullTimezone`

NewNodeFullTimezone instantiates a new NodeFullTimezone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFullTimezoneWithDefaults

`func NewNodeFullTimezoneWithDefaults() *NodeFullTimezone`

NewNodeFullTimezoneWithDefaults instantiates a new NodeFullTimezone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeFullTimezone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeFullTimezone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeFullTimezone) SetName(v string)`

SetName sets Name field to given value.


### GetOffset

`func (o *NodeFullTimezone) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *NodeFullTimezone) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *NodeFullTimezone) SetOffset(v string)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *NodeFullTimezone) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


