# RuleCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parent** | **string** | The parent category of the rules | 
**Id** | Pointer to **string** | Rule category id, only provide it when needed. | [optional] 
**Name** | **string** | Name of the rule category | 
**Description** | Pointer to **string** | Rules category description | [optional] 

## Methods

### NewRuleCategory

`func NewRuleCategory(parent string, name string, ) *RuleCategory`

NewRuleCategory instantiates a new RuleCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleCategoryWithDefaults

`func NewRuleCategoryWithDefaults() *RuleCategory`

NewRuleCategoryWithDefaults instantiates a new RuleCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParent

`func (o *RuleCategory) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *RuleCategory) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *RuleCategory) SetParent(v string)`

SetParent sets Parent field to given value.


### GetId

`func (o *RuleCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleCategory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuleCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RuleCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleCategory) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RuleCategory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleCategory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleCategory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuleCategory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


