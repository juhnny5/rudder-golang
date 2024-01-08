# CategoriesTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Category&#39;s name | 
**Path** | **string** | Category&#39;s path | 
**Id** | **string** | Category ID | 
**Subcategories** | [**[]TechniqueCategory**](TechniqueCategory.md) | List of sub categories | 

## Methods

### NewCategoriesTree

`func NewCategoriesTree(name string, path string, id string, subcategories []TechniqueCategory, ) *CategoriesTree`

NewCategoriesTree instantiates a new CategoriesTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoriesTreeWithDefaults

`func NewCategoriesTreeWithDefaults() *CategoriesTree`

NewCategoriesTreeWithDefaults instantiates a new CategoriesTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CategoriesTree) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoriesTree) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoriesTree) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *CategoriesTree) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CategoriesTree) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CategoriesTree) SetPath(v string)`

SetPath sets Path field to given value.


### GetId

`func (o *CategoriesTree) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CategoriesTree) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CategoriesTree) SetId(v string)`

SetId sets Id field to given value.


### GetSubcategories

`func (o *CategoriesTree) GetSubcategories() []TechniqueCategory`

GetSubcategories returns the Subcategories field if non-nil, zero value otherwise.

### GetSubcategoriesOk

`func (o *CategoriesTree) GetSubcategoriesOk() (*[]TechniqueCategory, bool)`

GetSubcategoriesOk returns a tuple with the Subcategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategories

`func (o *CategoriesTree) SetSubcategories(v []TechniqueCategory)`

SetSubcategories sets Subcategories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


