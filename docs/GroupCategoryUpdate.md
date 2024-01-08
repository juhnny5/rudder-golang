# GroupCategoryUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parent** | **string** | The parent category of the groups | 
**Name** | **string** | Name of the group category | 
**Description** | Pointer to **string** | Group category description | [optional] 

## Methods

### NewGroupCategoryUpdate

`func NewGroupCategoryUpdate(parent string, name string, ) *GroupCategoryUpdate`

NewGroupCategoryUpdate instantiates a new GroupCategoryUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupCategoryUpdateWithDefaults

`func NewGroupCategoryUpdateWithDefaults() *GroupCategoryUpdate`

NewGroupCategoryUpdateWithDefaults instantiates a new GroupCategoryUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParent

`func (o *GroupCategoryUpdate) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GroupCategoryUpdate) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GroupCategoryUpdate) SetParent(v string)`

SetParent sets Parent field to given value.


### GetName

`func (o *GroupCategoryUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupCategoryUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupCategoryUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GroupCategoryUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupCategoryUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupCategoryUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupCategoryUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


