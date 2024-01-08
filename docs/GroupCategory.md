# GroupCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parent** | **string** | The parent category of the groups | 
**Id** | Pointer to **string** | Group category id, only provide it when needed. | [optional] 
**Name** | **string** | Name of the group category | 
**Description** | Pointer to **string** | Group category description | [optional] 

## Methods

### NewGroupCategory

`func NewGroupCategory(parent string, name string, ) *GroupCategory`

NewGroupCategory instantiates a new GroupCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupCategoryWithDefaults

`func NewGroupCategoryWithDefaults() *GroupCategory`

NewGroupCategoryWithDefaults instantiates a new GroupCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParent

`func (o *GroupCategory) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GroupCategory) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GroupCategory) SetParent(v string)`

SetParent sets Parent field to given value.


### GetId

`func (o *GroupCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupCategory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GroupCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupCategory) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GroupCategory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupCategory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupCategory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupCategory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


