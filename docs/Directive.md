# Directive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Directive id | [optional] 
**DisplayName** | Pointer to **string** | Human readable name of the directive | [optional] 
**ShortDescription** | Pointer to **string** | One line directive description | [optional] 
**LongDescription** | Pointer to **string** | Description of the technique (rendered as markdown) | [optional] 
**TechniqueName** | Pointer to **string** | Directive id | [optional] 
**TechniqueVersion** | Pointer to **string** | Directive id | [optional] 
**Priority** | Pointer to **int32** | Directive priority. &#x60;0&#x60; has highest priority. | [optional] 
**Enabled** | Pointer to **bool** | Is the directive enabled | [optional] 
**System** | Pointer to **bool** | If true it is an internal Rudder directive | [optional] 
**PolicyMode** | Pointer to **string** | Policy mode of the directive | [optional] 
**Tags** | Pointer to [**[]DirectiveTagsInner**](DirectiveTagsInner.md) |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** | Directive parameters (depends on the source technique) | [optional] 

## Methods

### NewDirective

`func NewDirective() *Directive`

NewDirective instantiates a new Directive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectiveWithDefaults

`func NewDirectiveWithDefaults() *Directive`

NewDirectiveWithDefaults instantiates a new Directive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Directive) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Directive) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Directive) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Directive) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *Directive) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Directive) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Directive) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Directive) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetShortDescription

`func (o *Directive) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *Directive) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *Directive) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *Directive) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetLongDescription

`func (o *Directive) GetLongDescription() string`

GetLongDescription returns the LongDescription field if non-nil, zero value otherwise.

### GetLongDescriptionOk

`func (o *Directive) GetLongDescriptionOk() (*string, bool)`

GetLongDescriptionOk returns a tuple with the LongDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongDescription

`func (o *Directive) SetLongDescription(v string)`

SetLongDescription sets LongDescription field to given value.

### HasLongDescription

`func (o *Directive) HasLongDescription() bool`

HasLongDescription returns a boolean if a field has been set.

### GetTechniqueName

`func (o *Directive) GetTechniqueName() string`

GetTechniqueName returns the TechniqueName field if non-nil, zero value otherwise.

### GetTechniqueNameOk

`func (o *Directive) GetTechniqueNameOk() (*string, bool)`

GetTechniqueNameOk returns a tuple with the TechniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniqueName

`func (o *Directive) SetTechniqueName(v string)`

SetTechniqueName sets TechniqueName field to given value.

### HasTechniqueName

`func (o *Directive) HasTechniqueName() bool`

HasTechniqueName returns a boolean if a field has been set.

### GetTechniqueVersion

`func (o *Directive) GetTechniqueVersion() string`

GetTechniqueVersion returns the TechniqueVersion field if non-nil, zero value otherwise.

### GetTechniqueVersionOk

`func (o *Directive) GetTechniqueVersionOk() (*string, bool)`

GetTechniqueVersionOk returns a tuple with the TechniqueVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniqueVersion

`func (o *Directive) SetTechniqueVersion(v string)`

SetTechniqueVersion sets TechniqueVersion field to given value.

### HasTechniqueVersion

`func (o *Directive) HasTechniqueVersion() bool`

HasTechniqueVersion returns a boolean if a field has been set.

### GetPriority

`func (o *Directive) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Directive) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Directive) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Directive) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetEnabled

`func (o *Directive) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Directive) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Directive) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Directive) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSystem

`func (o *Directive) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Directive) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Directive) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *Directive) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetPolicyMode

`func (o *Directive) GetPolicyMode() string`

GetPolicyMode returns the PolicyMode field if non-nil, zero value otherwise.

### GetPolicyModeOk

`func (o *Directive) GetPolicyModeOk() (*string, bool)`

GetPolicyModeOk returns a tuple with the PolicyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMode

`func (o *Directive) SetPolicyMode(v string)`

SetPolicyMode sets PolicyMode field to given value.

### HasPolicyMode

`func (o *Directive) HasPolicyMode() bool`

HasPolicyMode returns a boolean if a field has been set.

### GetTags

`func (o *Directive) GetTags() []DirectiveTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Directive) GetTagsOk() (*[]DirectiveTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Directive) SetTags(v []DirectiveTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Directive) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetParameters

`func (o *Directive) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Directive) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Directive) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Directive) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


