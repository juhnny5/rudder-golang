# NodeInheritedPropertiesPropertiesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Property name | 
**Value** | **interface{}** | Resolved (ie, with inheritance and overriding done) property value (can be a string or JSON object) | 
**Provider** | Pointer to **string** | Property provider (if the property is not a simple node property) | [optional] 
**Hierarchy** | Pointer to [**[]NodeInheritedPropertiesPropertiesInnerHierarchyInner**](NodeInheritedPropertiesPropertiesInnerHierarchyInner.md) | A description of the inheritance hierarchy for that property, with most direct parent at head and oldest one at tail | [optional] 
**Origval** | Pointer to **interface{}** | The original value (ie, before overriding and inheritance resolution) for that node | [optional] 

## Methods

### NewNodeInheritedPropertiesPropertiesInner

`func NewNodeInheritedPropertiesPropertiesInner(name string, value interface{}, ) *NodeInheritedPropertiesPropertiesInner`

NewNodeInheritedPropertiesPropertiesInner instantiates a new NodeInheritedPropertiesPropertiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeInheritedPropertiesPropertiesInnerWithDefaults

`func NewNodeInheritedPropertiesPropertiesInnerWithDefaults() *NodeInheritedPropertiesPropertiesInner`

NewNodeInheritedPropertiesPropertiesInnerWithDefaults instantiates a new NodeInheritedPropertiesPropertiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeInheritedPropertiesPropertiesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeInheritedPropertiesPropertiesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeInheritedPropertiesPropertiesInner) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *NodeInheritedPropertiesPropertiesInner) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NodeInheritedPropertiesPropertiesInner) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NodeInheritedPropertiesPropertiesInner) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *NodeInheritedPropertiesPropertiesInner) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *NodeInheritedPropertiesPropertiesInner) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetProvider

`func (o *NodeInheritedPropertiesPropertiesInner) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NodeInheritedPropertiesPropertiesInner) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NodeInheritedPropertiesPropertiesInner) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NodeInheritedPropertiesPropertiesInner) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetHierarchy

`func (o *NodeInheritedPropertiesPropertiesInner) GetHierarchy() []NodeInheritedPropertiesPropertiesInnerHierarchyInner`

GetHierarchy returns the Hierarchy field if non-nil, zero value otherwise.

### GetHierarchyOk

`func (o *NodeInheritedPropertiesPropertiesInner) GetHierarchyOk() (*[]NodeInheritedPropertiesPropertiesInnerHierarchyInner, bool)`

GetHierarchyOk returns a tuple with the Hierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchy

`func (o *NodeInheritedPropertiesPropertiesInner) SetHierarchy(v []NodeInheritedPropertiesPropertiesInnerHierarchyInner)`

SetHierarchy sets Hierarchy field to given value.

### HasHierarchy

`func (o *NodeInheritedPropertiesPropertiesInner) HasHierarchy() bool`

HasHierarchy returns a boolean if a field has been set.

### GetOrigval

`func (o *NodeInheritedPropertiesPropertiesInner) GetOrigval() interface{}`

GetOrigval returns the Origval field if non-nil, zero value otherwise.

### GetOrigvalOk

`func (o *NodeInheritedPropertiesPropertiesInner) GetOrigvalOk() (*interface{}, bool)`

GetOrigvalOk returns a tuple with the Origval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigval

`func (o *NodeInheritedPropertiesPropertiesInner) SetOrigval(v interface{})`

SetOrigval sets Origval field to given value.

### HasOrigval

`func (o *NodeInheritedPropertiesPropertiesInner) HasOrigval() bool`

HasOrigval returns a boolean if a field has been set.

### SetOrigvalNil

`func (o *NodeInheritedPropertiesPropertiesInner) SetOrigvalNil(b bool)`

 SetOrigvalNil sets the value for Origval to be an explicit nil

### UnsetOrigval
`func (o *NodeInheritedPropertiesPropertiesInner) UnsetOrigval()`

UnsetOrigval ensures that no value is present for Origval, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


