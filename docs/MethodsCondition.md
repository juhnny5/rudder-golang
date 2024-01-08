# MethodsCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **string** | Prefix used to generate condition | [optional] 
**Parameter** | Pointer to **string** | Id of the parameter used to generate condition | [optional] 

## Methods

### NewMethodsCondition

`func NewMethodsCondition() *MethodsCondition`

NewMethodsCondition instantiates a new MethodsCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMethodsConditionWithDefaults

`func NewMethodsConditionWithDefaults() *MethodsCondition`

NewMethodsConditionWithDefaults instantiates a new MethodsCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *MethodsCondition) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *MethodsCondition) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *MethodsCondition) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *MethodsCondition) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetParameter

`func (o *MethodsCondition) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *MethodsCondition) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *MethodsCondition) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *MethodsCondition) HasParameter() bool`

HasParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


