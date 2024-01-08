# MethodParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Parameter name | 
**Description** | **string** | Description of this parameter | 
**Constraints** | [**MethodParameterConstraints**](MethodParameterConstraints.md) |  | 
**Type** | Pointer to **string** | Type of the parameter | [optional] [default to "String"]

## Methods

### NewMethodParameter

`func NewMethodParameter(name string, description string, constraints MethodParameterConstraints, ) *MethodParameter`

NewMethodParameter instantiates a new MethodParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMethodParameterWithDefaults

`func NewMethodParameterWithDefaults() *MethodParameter`

NewMethodParameterWithDefaults instantiates a new MethodParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MethodParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MethodParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MethodParameter) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MethodParameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MethodParameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MethodParameter) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetConstraints

`func (o *MethodParameter) GetConstraints() MethodParameterConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *MethodParameter) GetConstraintsOk() (*MethodParameterConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *MethodParameter) SetConstraints(v MethodParameterConstraints)`

SetConstraints sets Constraints field to given value.


### GetType

`func (o *MethodParameter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MethodParameter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MethodParameter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MethodParameter) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


