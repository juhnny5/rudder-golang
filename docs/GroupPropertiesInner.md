# GroupPropertiesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Property name | 
**Value** | **interface{}** | Property value (can be a string or JSON object) | 

## Methods

### NewGroupPropertiesInner

`func NewGroupPropertiesInner(name string, value interface{}, ) *GroupPropertiesInner`

NewGroupPropertiesInner instantiates a new GroupPropertiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPropertiesInnerWithDefaults

`func NewGroupPropertiesInnerWithDefaults() *GroupPropertiesInner`

NewGroupPropertiesInnerWithDefaults instantiates a new GroupPropertiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupPropertiesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupPropertiesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupPropertiesInner) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *GroupPropertiesInner) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GroupPropertiesInner) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GroupPropertiesInner) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *GroupPropertiesInner) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *GroupPropertiesInner) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


