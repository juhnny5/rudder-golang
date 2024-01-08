# MethodParameterConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowEmptyString** | **bool** | Can this parameter be empty? | 
**AllowWhitespaceString** | **bool** | Can this parameter allow trailing/ending spaces, or even a full whitespace string ? | 
**MaxLength** | **int32** | Maximum size of a parameter | 
**MinLength** | **int32** | Minimal size of a parameter | 
**Regex** | **string** | A regex to validate this parameter | 
**NotRegex** | **string** | A regexp to invalidate this parameter | 
**Select** | **[]string** | List of items authorized for this parameter | 

## Methods

### NewMethodParameterConstraints

`func NewMethodParameterConstraints(allowEmptyString bool, allowWhitespaceString bool, maxLength int32, minLength int32, regex string, notRegex string, select_ []string, ) *MethodParameterConstraints`

NewMethodParameterConstraints instantiates a new MethodParameterConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMethodParameterConstraintsWithDefaults

`func NewMethodParameterConstraintsWithDefaults() *MethodParameterConstraints`

NewMethodParameterConstraintsWithDefaults instantiates a new MethodParameterConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowEmptyString

`func (o *MethodParameterConstraints) GetAllowEmptyString() bool`

GetAllowEmptyString returns the AllowEmptyString field if non-nil, zero value otherwise.

### GetAllowEmptyStringOk

`func (o *MethodParameterConstraints) GetAllowEmptyStringOk() (*bool, bool)`

GetAllowEmptyStringOk returns a tuple with the AllowEmptyString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmptyString

`func (o *MethodParameterConstraints) SetAllowEmptyString(v bool)`

SetAllowEmptyString sets AllowEmptyString field to given value.


### GetAllowWhitespaceString

`func (o *MethodParameterConstraints) GetAllowWhitespaceString() bool`

GetAllowWhitespaceString returns the AllowWhitespaceString field if non-nil, zero value otherwise.

### GetAllowWhitespaceStringOk

`func (o *MethodParameterConstraints) GetAllowWhitespaceStringOk() (*bool, bool)`

GetAllowWhitespaceStringOk returns a tuple with the AllowWhitespaceString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWhitespaceString

`func (o *MethodParameterConstraints) SetAllowWhitespaceString(v bool)`

SetAllowWhitespaceString sets AllowWhitespaceString field to given value.


### GetMaxLength

`func (o *MethodParameterConstraints) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *MethodParameterConstraints) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *MethodParameterConstraints) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.


### GetMinLength

`func (o *MethodParameterConstraints) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *MethodParameterConstraints) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *MethodParameterConstraints) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.


### GetRegex

`func (o *MethodParameterConstraints) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *MethodParameterConstraints) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *MethodParameterConstraints) SetRegex(v string)`

SetRegex sets Regex field to given value.


### GetNotRegex

`func (o *MethodParameterConstraints) GetNotRegex() string`

GetNotRegex returns the NotRegex field if non-nil, zero value otherwise.

### GetNotRegexOk

`func (o *MethodParameterConstraints) GetNotRegexOk() (*string, bool)`

GetNotRegexOk returns a tuple with the NotRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotRegex

`func (o *MethodParameterConstraints) SetNotRegex(v string)`

SetNotRegex sets NotRegex field to given value.


### GetSelect

`func (o *MethodParameterConstraints) GetSelect() []string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *MethodParameterConstraints) GetSelectOk() (*[]string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *MethodParameterConstraints) SetSelect(v []string)`

SetSelect sets Select field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


