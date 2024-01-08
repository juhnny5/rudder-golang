# TechniqueParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | parameter id | [optional] 
**Name** | Pointer to **string** | Parameter name | [optional] 
**Description** | Pointer to **string** | description of this parameter | [optional] 
**MayBeEmpty** | Pointer to **bool** | May the value given when creating a directive be empty | [optional] 

## Methods

### NewTechniqueParameter

`func NewTechniqueParameter() *TechniqueParameter`

NewTechniqueParameter instantiates a new TechniqueParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechniqueParameterWithDefaults

`func NewTechniqueParameterWithDefaults() *TechniqueParameter`

NewTechniqueParameterWithDefaults instantiates a new TechniqueParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TechniqueParameter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TechniqueParameter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TechniqueParameter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TechniqueParameter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TechniqueParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TechniqueParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TechniqueParameter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TechniqueParameter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TechniqueParameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TechniqueParameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TechniqueParameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TechniqueParameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMayBeEmpty

`func (o *TechniqueParameter) GetMayBeEmpty() bool`

GetMayBeEmpty returns the MayBeEmpty field if non-nil, zero value otherwise.

### GetMayBeEmptyOk

`func (o *TechniqueParameter) GetMayBeEmptyOk() (*bool, bool)`

GetMayBeEmptyOk returns a tuple with the MayBeEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMayBeEmpty

`func (o *TechniqueParameter) SetMayBeEmpty(v bool)`

SetMayBeEmpty sets MayBeEmpty field to given value.

### HasMayBeEmpty

`func (o *TechniqueParameter) HasMayBeEmpty() bool`

HasMayBeEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


