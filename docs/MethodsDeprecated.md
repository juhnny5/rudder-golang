# MethodsDeprecated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to **string** | Information notice about the deprecation, especially how to replace it | [optional] 
**ReplacedBy** | Pointer to **string** | Id of the method replacing this method | [optional] 

## Methods

### NewMethodsDeprecated

`func NewMethodsDeprecated() *MethodsDeprecated`

NewMethodsDeprecated instantiates a new MethodsDeprecated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMethodsDeprecatedWithDefaults

`func NewMethodsDeprecatedWithDefaults() *MethodsDeprecated`

NewMethodsDeprecatedWithDefaults instantiates a new MethodsDeprecated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *MethodsDeprecated) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MethodsDeprecated) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MethodsDeprecated) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MethodsDeprecated) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetReplacedBy

`func (o *MethodsDeprecated) GetReplacedBy() string`

GetReplacedBy returns the ReplacedBy field if non-nil, zero value otherwise.

### GetReplacedByOk

`func (o *MethodsDeprecated) GetReplacedByOk() (*string, bool)`

GetReplacedByOk returns a tuple with the ReplacedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedBy

`func (o *MethodsDeprecated) SetReplacedBy(v string)`

SetReplacedBy sets ReplacedBy field to given value.

### HasReplacedBy

`func (o *MethodsDeprecated) HasReplacedBy() bool`

HasReplacedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


