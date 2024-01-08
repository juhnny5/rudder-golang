# Parameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the parameter | 
**Value** | Pointer to **interface{}** | Value of the parameter | [optional] 
**Description** | Pointer to **string** | Description of the parameter | [optional] 
**Overridable** | Pointer to **bool** | Is the global parameter overridable by node | [optional] 

## Methods

### NewParameter

`func NewParameter(id string, ) *Parameter`

NewParameter instantiates a new Parameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterWithDefaults

`func NewParameterWithDefaults() *Parameter`

NewParameterWithDefaults instantiates a new Parameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Parameter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Parameter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Parameter) SetId(v string)`

SetId sets Id field to given value.


### GetValue

`func (o *Parameter) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Parameter) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Parameter) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *Parameter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Parameter) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Parameter) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetDescription

`func (o *Parameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Parameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Parameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Parameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOverridable

`func (o *Parameter) GetOverridable() bool`

GetOverridable returns the Overridable field if non-nil, zero value otherwise.

### GetOverridableOk

`func (o *Parameter) GetOverridableOk() (*bool, bool)`

GetOverridableOk returns a tuple with the Overridable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridable

`func (o *Parameter) SetOverridable(v bool)`

SetOverridable sets Overridable field to given value.

### HasOverridable

`func (o *Parameter) HasOverridable() bool`

HasOverridable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


