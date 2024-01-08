# RuleNew

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | The id of the rule the clone will be based onto. If this parameter if provided, the new rule will be a clone of this source. Other value will override values from the source. | [optional] 
**Id** | Pointer to **string** | Rule id | [optional] 
**DisplayName** | Pointer to **string** | Rule name | [optional] 
**ShortDescription** | Pointer to **string** | One line rule description | [optional] 
**LongDescription** | Pointer to **string** | Rule documentation | [optional] 
**Category** | Pointer to **string** | The parent category id. If provided, the new rule will be in this parent category | [optional] 
**Directives** | Pointer to **[]string** | Directives linked to the rule | [optional] 
**Targets** | Pointer to [**[]RuleTargetsInner**](RuleTargetsInner.md) | Node and special groups targeted by that rule | [optional] 
**Enabled** | Pointer to **bool** | Is the rule enabled | [optional] 
**System** | Pointer to **bool** | If true it is an internal Rudder rule | [optional] 
**Tags** | Pointer to [**[]DirectiveTagsInner**](DirectiveTagsInner.md) |  | [optional] 

## Methods

### NewRuleNew

`func NewRuleNew() *RuleNew`

NewRuleNew instantiates a new RuleNew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleNewWithDefaults

`func NewRuleNewWithDefaults() *RuleNew`

NewRuleNewWithDefaults instantiates a new RuleNew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RuleNew) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RuleNew) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RuleNew) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RuleNew) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetId

`func (o *RuleNew) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleNew) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleNew) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuleNew) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *RuleNew) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RuleNew) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RuleNew) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RuleNew) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetShortDescription

`func (o *RuleNew) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *RuleNew) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *RuleNew) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *RuleNew) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetLongDescription

`func (o *RuleNew) GetLongDescription() string`

GetLongDescription returns the LongDescription field if non-nil, zero value otherwise.

### GetLongDescriptionOk

`func (o *RuleNew) GetLongDescriptionOk() (*string, bool)`

GetLongDescriptionOk returns a tuple with the LongDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongDescription

`func (o *RuleNew) SetLongDescription(v string)`

SetLongDescription sets LongDescription field to given value.

### HasLongDescription

`func (o *RuleNew) HasLongDescription() bool`

HasLongDescription returns a boolean if a field has been set.

### GetCategory

`func (o *RuleNew) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *RuleNew) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *RuleNew) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *RuleNew) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDirectives

`func (o *RuleNew) GetDirectives() []string`

GetDirectives returns the Directives field if non-nil, zero value otherwise.

### GetDirectivesOk

`func (o *RuleNew) GetDirectivesOk() (*[]string, bool)`

GetDirectivesOk returns a tuple with the Directives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectives

`func (o *RuleNew) SetDirectives(v []string)`

SetDirectives sets Directives field to given value.

### HasDirectives

`func (o *RuleNew) HasDirectives() bool`

HasDirectives returns a boolean if a field has been set.

### GetTargets

`func (o *RuleNew) GetTargets() []RuleTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *RuleNew) GetTargetsOk() (*[]RuleTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *RuleNew) SetTargets(v []RuleTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *RuleNew) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetEnabled

`func (o *RuleNew) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RuleNew) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RuleNew) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RuleNew) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSystem

`func (o *RuleNew) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *RuleNew) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *RuleNew) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *RuleNew) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTags

`func (o *RuleNew) GetTags() []DirectiveTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RuleNew) GetTagsOk() (*[]DirectiveTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RuleNew) SetTags(v []DirectiveTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RuleNew) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


