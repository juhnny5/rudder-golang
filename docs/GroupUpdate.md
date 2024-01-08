# GroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Id of the new group&#39;s category | [optional] 
**DisplayName** | Pointer to **string** | Name of the group | [optional] 
**Description** | Pointer to **string** | Group description | [optional] 
**Query** | Pointer to [**GroupNewQuery**](GroupNewQuery.md) |  | [optional] 
**Dynamic** | Pointer to **bool** | Should the group be dynamically refreshed (if not, it is a static group) | [optional] [default to true]
**Enabled** | Pointer to **bool** | Enable or disable the group | [optional] [default to true]

## Methods

### NewGroupUpdate

`func NewGroupUpdate() *GroupUpdate`

NewGroupUpdate instantiates a new GroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUpdateWithDefaults

`func NewGroupUpdateWithDefaults() *GroupUpdate`

NewGroupUpdateWithDefaults instantiates a new GroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *GroupUpdate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GroupUpdate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GroupUpdate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GroupUpdate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDisplayName

`func (o *GroupUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GroupUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GroupUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GroupUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *GroupUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuery

`func (o *GroupUpdate) GetQuery() GroupNewQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GroupUpdate) GetQueryOk() (*GroupNewQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GroupUpdate) SetQuery(v GroupNewQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *GroupUpdate) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetDynamic

`func (o *GroupUpdate) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *GroupUpdate) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *GroupUpdate) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *GroupUpdate) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetEnabled

`func (o *GroupUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GroupUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GroupUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GroupUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


