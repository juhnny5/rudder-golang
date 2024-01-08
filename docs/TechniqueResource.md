# TechniqueResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | resource name. this is the relative path to the resource | [optional] 
**State** | Pointer to **string** | State of the resource file. it can be a value between new, modified, deleted, untouched | [optional] 

## Methods

### NewTechniqueResource

`func NewTechniqueResource() *TechniqueResource`

NewTechniqueResource instantiates a new TechniqueResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechniqueResourceWithDefaults

`func NewTechniqueResourceWithDefaults() *TechniqueResource`

NewTechniqueResourceWithDefaults instantiates a new TechniqueResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TechniqueResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TechniqueResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TechniqueResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TechniqueResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *TechniqueResource) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TechniqueResource) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TechniqueResource) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TechniqueResource) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


