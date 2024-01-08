# TechniqueCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Category&#39;s name | 
**Path** | **string** | Category&#39;s path | 
**Id** | **string** | Category ID | 
**Subcategories** | [**[]TechniqueCategory**](TechniqueCategory.md) | List of sub categories | 

## Methods

### NewTechniqueCategory

`func NewTechniqueCategory(name string, path string, id string, subcategories []TechniqueCategory, ) *TechniqueCategory`

NewTechniqueCategory instantiates a new TechniqueCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechniqueCategoryWithDefaults

`func NewTechniqueCategoryWithDefaults() *TechniqueCategory`

NewTechniqueCategoryWithDefaults instantiates a new TechniqueCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TechniqueCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TechniqueCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TechniqueCategory) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *TechniqueCategory) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TechniqueCategory) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TechniqueCategory) SetPath(v string)`

SetPath sets Path field to given value.


### GetId

`func (o *TechniqueCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TechniqueCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TechniqueCategory) SetId(v string)`

SetId sets Id field to given value.


### GetSubcategories

`func (o *TechniqueCategory) GetSubcategories() []TechniqueCategory`

GetSubcategories returns the Subcategories field if non-nil, zero value otherwise.

### GetSubcategoriesOk

`func (o *TechniqueCategory) GetSubcategoriesOk() (*[]TechniqueCategory, bool)`

GetSubcategoriesOk returns a tuple with the Subcategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategories

`func (o *TechniqueCategory) SetSubcategories(v []TechniqueCategory)`

SetSubcategories sets Subcategories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


