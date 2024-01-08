# Methods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Method id | 
**Name** | **string** | Method name | 
**Version** | **string** | Version of this technique | 
**Category** | **string** | Category of this technique | 
**Desc** | **string** | Description of this method | 
**Documentation** | **string** | Full documentation of this method | 
**Parameters** | [**[]MethodParameter**](MethodParameter.md) | Parameters for this technique | 
**Agents** | **[]string** | List of agents compatible with this method | 
**Condition** | [**MethodsCondition**](MethodsCondition.md) |  | 
**Deprecated** | [**MethodsDeprecated**](MethodsDeprecated.md) |  | 

## Methods

### NewMethods

`func NewMethods(id string, name string, version string, category string, desc string, documentation string, parameters []MethodParameter, agents []string, condition MethodsCondition, deprecated MethodsDeprecated, ) *Methods`

NewMethods instantiates a new Methods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMethodsWithDefaults

`func NewMethodsWithDefaults() *Methods`

NewMethodsWithDefaults instantiates a new Methods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Methods) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Methods) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Methods) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Methods) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Methods) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Methods) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *Methods) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Methods) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Methods) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetCategory

`func (o *Methods) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Methods) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Methods) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetDesc

`func (o *Methods) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Methods) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Methods) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetDocumentation

`func (o *Methods) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *Methods) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *Methods) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.


### GetParameters

`func (o *Methods) GetParameters() []MethodParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Methods) GetParametersOk() (*[]MethodParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Methods) SetParameters(v []MethodParameter)`

SetParameters sets Parameters field to given value.


### GetAgents

`func (o *Methods) GetAgents() []string`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *Methods) GetAgentsOk() (*[]string, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *Methods) SetAgents(v []string)`

SetAgents sets Agents field to given value.


### GetCondition

`func (o *Methods) GetCondition() MethodsCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *Methods) GetConditionOk() (*MethodsCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *Methods) SetCondition(v MethodsCondition)`

SetCondition sets Condition field to given value.


### GetDeprecated

`func (o *Methods) GetDeprecated() MethodsDeprecated`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *Methods) GetDeprecatedOk() (*MethodsDeprecated, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *Methods) SetDeprecated(v MethodsDeprecated)`

SetDeprecated sets Deprecated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


