# DirectiveNew

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | The id of the directive the clone will be based onto. If this parameter if provided,  the new directive will be a clone of this source. Other value will override values from the source. | [optional] 
**Id** | Pointer to **string** | Directive id | [optional] 
**DisplayName** | Pointer to **string** | Human readable name of the directive | [optional] 
**ShortDescription** | Pointer to **string** | One line directive description | [optional] 
**LongDescription** | Pointer to **string** | Description of the technique (rendered as markdown) | [optional] 
**TechniqueName** | Pointer to **string** | Directive id | [optional] 
**TechniqueVersion** | Pointer to **string** | Directive id | [optional] 
**Priority** | Pointer to **int32** | Directive priority. &#x60;0&#x60; has highest priority. | [optional] 
**Enabled** | Pointer to **bool** | Is the directive enabled | [optional] 
**System** | Pointer to **bool** | If true it is an internal Rudder directive | [optional] 
**Tags** | Pointer to [**[]DirectiveTagsInner**](DirectiveTagsInner.md) |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** | Directive parameters (depends on the source technique) | [optional] 

## Methods

### NewDirectiveNew

`func NewDirectiveNew() *DirectiveNew`

NewDirectiveNew instantiates a new DirectiveNew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectiveNewWithDefaults

`func NewDirectiveNewWithDefaults() *DirectiveNew`

NewDirectiveNewWithDefaults instantiates a new DirectiveNew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *DirectiveNew) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DirectiveNew) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DirectiveNew) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DirectiveNew) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetId

`func (o *DirectiveNew) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DirectiveNew) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DirectiveNew) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DirectiveNew) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *DirectiveNew) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DirectiveNew) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DirectiveNew) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DirectiveNew) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetShortDescription

`func (o *DirectiveNew) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *DirectiveNew) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *DirectiveNew) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *DirectiveNew) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetLongDescription

`func (o *DirectiveNew) GetLongDescription() string`

GetLongDescription returns the LongDescription field if non-nil, zero value otherwise.

### GetLongDescriptionOk

`func (o *DirectiveNew) GetLongDescriptionOk() (*string, bool)`

GetLongDescriptionOk returns a tuple with the LongDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongDescription

`func (o *DirectiveNew) SetLongDescription(v string)`

SetLongDescription sets LongDescription field to given value.

### HasLongDescription

`func (o *DirectiveNew) HasLongDescription() bool`

HasLongDescription returns a boolean if a field has been set.

### GetTechniqueName

`func (o *DirectiveNew) GetTechniqueName() string`

GetTechniqueName returns the TechniqueName field if non-nil, zero value otherwise.

### GetTechniqueNameOk

`func (o *DirectiveNew) GetTechniqueNameOk() (*string, bool)`

GetTechniqueNameOk returns a tuple with the TechniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniqueName

`func (o *DirectiveNew) SetTechniqueName(v string)`

SetTechniqueName sets TechniqueName field to given value.

### HasTechniqueName

`func (o *DirectiveNew) HasTechniqueName() bool`

HasTechniqueName returns a boolean if a field has been set.

### GetTechniqueVersion

`func (o *DirectiveNew) GetTechniqueVersion() string`

GetTechniqueVersion returns the TechniqueVersion field if non-nil, zero value otherwise.

### GetTechniqueVersionOk

`func (o *DirectiveNew) GetTechniqueVersionOk() (*string, bool)`

GetTechniqueVersionOk returns a tuple with the TechniqueVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniqueVersion

`func (o *DirectiveNew) SetTechniqueVersion(v string)`

SetTechniqueVersion sets TechniqueVersion field to given value.

### HasTechniqueVersion

`func (o *DirectiveNew) HasTechniqueVersion() bool`

HasTechniqueVersion returns a boolean if a field has been set.

### GetPriority

`func (o *DirectiveNew) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DirectiveNew) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DirectiveNew) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DirectiveNew) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetEnabled

`func (o *DirectiveNew) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DirectiveNew) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DirectiveNew) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DirectiveNew) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSystem

`func (o *DirectiveNew) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *DirectiveNew) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *DirectiveNew) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *DirectiveNew) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTags

`func (o *DirectiveNew) GetTags() []DirectiveTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DirectiveNew) GetTagsOk() (*[]DirectiveTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DirectiveNew) SetTags(v []DirectiveTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DirectiveNew) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetParameters

`func (o *DirectiveNew) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *DirectiveNew) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *DirectiveNew) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *DirectiveNew) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


