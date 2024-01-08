# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Rule id | [optional] 
**DisplayName** | Pointer to **string** | Rule name | [optional] 
**ShortDescription** | Pointer to **string** | One line rule description | [optional] 
**LongDescription** | Pointer to **string** | Rule documentation | [optional] 
**Directives** | Pointer to **[]string** | Directives linked to the rule | [optional] 
**Targets** | Pointer to [**[]RuleTargetsInner**](RuleTargetsInner.md) | Node and special groups targeted by that rule | [optional] 
**Enabled** | Pointer to **bool** | Is the rule enabled | [optional] 
**System** | Pointer to **bool** | If true it is an internal Rudder rule | [optional] 
**Tags** | Pointer to [**[]DirectiveTagsInner**](DirectiveTagsInner.md) |  | [optional] 
**PolicyMode** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**RuleStatus**](RuleStatus.md) |  | [optional] 

## Methods

### NewRule

`func NewRule() *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Rule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Rule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *Rule) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Rule) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Rule) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Rule) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetShortDescription

`func (o *Rule) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *Rule) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *Rule) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *Rule) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetLongDescription

`func (o *Rule) GetLongDescription() string`

GetLongDescription returns the LongDescription field if non-nil, zero value otherwise.

### GetLongDescriptionOk

`func (o *Rule) GetLongDescriptionOk() (*string, bool)`

GetLongDescriptionOk returns a tuple with the LongDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongDescription

`func (o *Rule) SetLongDescription(v string)`

SetLongDescription sets LongDescription field to given value.

### HasLongDescription

`func (o *Rule) HasLongDescription() bool`

HasLongDescription returns a boolean if a field has been set.

### GetDirectives

`func (o *Rule) GetDirectives() []string`

GetDirectives returns the Directives field if non-nil, zero value otherwise.

### GetDirectivesOk

`func (o *Rule) GetDirectivesOk() (*[]string, bool)`

GetDirectivesOk returns a tuple with the Directives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectives

`func (o *Rule) SetDirectives(v []string)`

SetDirectives sets Directives field to given value.

### HasDirectives

`func (o *Rule) HasDirectives() bool`

HasDirectives returns a boolean if a field has been set.

### GetTargets

`func (o *Rule) GetTargets() []RuleTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *Rule) GetTargetsOk() (*[]RuleTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *Rule) SetTargets(v []RuleTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *Rule) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetEnabled

`func (o *Rule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Rule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Rule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Rule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSystem

`func (o *Rule) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Rule) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Rule) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *Rule) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTags

`func (o *Rule) GetTags() []DirectiveTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Rule) GetTagsOk() (*[]DirectiveTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Rule) SetTags(v []DirectiveTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Rule) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPolicyMode

`func (o *Rule) GetPolicyMode() string`

GetPolicyMode returns the PolicyMode field if non-nil, zero value otherwise.

### GetPolicyModeOk

`func (o *Rule) GetPolicyModeOk() (*string, bool)`

GetPolicyModeOk returns a tuple with the PolicyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMode

`func (o *Rule) SetPolicyMode(v string)`

SetPolicyMode sets PolicyMode field to given value.

### HasPolicyMode

`func (o *Rule) HasPolicyMode() bool`

HasPolicyMode returns a boolean if a field has been set.

### GetStatus

`func (o *Rule) GetStatus() RuleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Rule) GetStatusOk() (*RuleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Rule) SetStatus(v RuleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Rule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


