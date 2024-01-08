# EditorTechnique

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Technique id | [optional] 
**Name** | Pointer to **string** | Technique name | [optional] 
**Version** | Pointer to **string** | version of this technique | [optional] 
**Category** | Pointer to **string** | category of this technique | [optional] 
**Description** | Pointer to **string** | description of this technique | [optional] 
**Source** | Pointer to **string** | Source of the technique, always editor here | [optional] 
**Parameters** | Pointer to [**[]TechniqueParameter**](TechniqueParameter.md) | Parameters for this technique | [optional] 
**Resources** | Pointer to [**[]TechniqueResource**](TechniqueResource.md) | Resources for this technique | [optional] 
**Calls** | Pointer to [**[]TechniqueBlockCallsInner**](TechniqueBlockCallsInner.md) | Method and blocks contained by this technique | [optional] 

## Methods

### NewEditorTechnique

`func NewEditorTechnique() *EditorTechnique`

NewEditorTechnique instantiates a new EditorTechnique object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditorTechniqueWithDefaults

`func NewEditorTechniqueWithDefaults() *EditorTechnique`

NewEditorTechniqueWithDefaults instantiates a new EditorTechnique object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditorTechnique) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditorTechnique) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditorTechnique) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EditorTechnique) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EditorTechnique) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditorTechnique) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditorTechnique) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditorTechnique) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *EditorTechnique) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EditorTechnique) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EditorTechnique) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EditorTechnique) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCategory

`func (o *EditorTechnique) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EditorTechnique) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EditorTechnique) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EditorTechnique) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *EditorTechnique) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditorTechnique) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditorTechnique) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditorTechnique) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSource

`func (o *EditorTechnique) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EditorTechnique) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EditorTechnique) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *EditorTechnique) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetParameters

`func (o *EditorTechnique) GetParameters() []TechniqueParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *EditorTechnique) GetParametersOk() (*[]TechniqueParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *EditorTechnique) SetParameters(v []TechniqueParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *EditorTechnique) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetResources

`func (o *EditorTechnique) GetResources() []TechniqueResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *EditorTechnique) GetResourcesOk() (*[]TechniqueResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *EditorTechnique) SetResources(v []TechniqueResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *EditorTechnique) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetCalls

`func (o *EditorTechnique) GetCalls() []TechniqueBlockCallsInner`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *EditorTechnique) GetCallsOk() (*[]TechniqueBlockCallsInner, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *EditorTechnique) SetCalls(v []TechniqueBlockCallsInner)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *EditorTechnique) HasCalls() bool`

HasCalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


