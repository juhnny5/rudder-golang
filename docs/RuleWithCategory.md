# RuleWithCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Rule id | [optional] 
**DisplayName** | Pointer to **string** | Rule name | [optional] 
**ShortDescription** | Pointer to **string** | One line rule description | [optional] 
**LongDescription** | Pointer to **string** | Rule documentation | [optional] 
**Category** | Pointer to **string** | The parent category id. | [optional] 
**Directives** | Pointer to **[]string** | Directives linked to the rule | [optional] 
**Targets** | Pointer to [**[]RuleTargetsInner**](RuleTargetsInner.md) | Node and special groups targeted by that rule | [optional] 
**Enabled** | Pointer to **bool** | Is the rule enabled | [optional] 
**System** | Pointer to **bool** | If true it is an internal Rudder rule | [optional] 
**Tags** | Pointer to [**[]DirectiveTagsInner**](DirectiveTagsInner.md) |  | [optional] 

## Methods

### NewRuleWithCategory

`func NewRuleWithCategory() *RuleWithCategory`

NewRuleWithCategory instantiates a new RuleWithCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithCategoryWithDefaults

`func NewRuleWithCategoryWithDefaults() *RuleWithCategory`

NewRuleWithCategoryWithDefaults instantiates a new RuleWithCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuleWithCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleWithCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleWithCategory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuleWithCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *RuleWithCategory) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RuleWithCategory) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RuleWithCategory) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RuleWithCategory) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetShortDescription

`func (o *RuleWithCategory) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *RuleWithCategory) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *RuleWithCategory) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *RuleWithCategory) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetLongDescription

`func (o *RuleWithCategory) GetLongDescription() string`

GetLongDescription returns the LongDescription field if non-nil, zero value otherwise.

### GetLongDescriptionOk

`func (o *RuleWithCategory) GetLongDescriptionOk() (*string, bool)`

GetLongDescriptionOk returns a tuple with the LongDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongDescription

`func (o *RuleWithCategory) SetLongDescription(v string)`

SetLongDescription sets LongDescription field to given value.

### HasLongDescription

`func (o *RuleWithCategory) HasLongDescription() bool`

HasLongDescription returns a boolean if a field has been set.

### GetCategory

`func (o *RuleWithCategory) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *RuleWithCategory) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *RuleWithCategory) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *RuleWithCategory) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDirectives

`func (o *RuleWithCategory) GetDirectives() []string`

GetDirectives returns the Directives field if non-nil, zero value otherwise.

### GetDirectivesOk

`func (o *RuleWithCategory) GetDirectivesOk() (*[]string, bool)`

GetDirectivesOk returns a tuple with the Directives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectives

`func (o *RuleWithCategory) SetDirectives(v []string)`

SetDirectives sets Directives field to given value.

### HasDirectives

`func (o *RuleWithCategory) HasDirectives() bool`

HasDirectives returns a boolean if a field has been set.

### GetTargets

`func (o *RuleWithCategory) GetTargets() []RuleTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *RuleWithCategory) GetTargetsOk() (*[]RuleTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *RuleWithCategory) SetTargets(v []RuleTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *RuleWithCategory) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetEnabled

`func (o *RuleWithCategory) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RuleWithCategory) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RuleWithCategory) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RuleWithCategory) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSystem

`func (o *RuleWithCategory) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *RuleWithCategory) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *RuleWithCategory) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *RuleWithCategory) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTags

`func (o *RuleWithCategory) GetTags() []DirectiveTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RuleWithCategory) GetTagsOk() (*[]DirectiveTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RuleWithCategory) SetTags(v []DirectiveTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RuleWithCategory) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


