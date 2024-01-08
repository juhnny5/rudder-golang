# GroupNew

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | The id of the group the clone will be based onto. If this parameter if provided,  the new group will be a clone of this source. Other value will override values from the source. | [optional] 
**Category** | **string** | Id of the new group&#39;s category | 
**Id** | Pointer to **string** | Group id, only provide it when needed. | [optional] 
**DisplayName** | **string** | Name of the group | 
**Description** | Pointer to **string** | Group description | [optional] 
**Query** | Pointer to [**GroupNewQuery**](GroupNewQuery.md) |  | [optional] 
**Dynamic** | Pointer to **bool** | Should the group be dynamically refreshed (if not, it is a static group) | [optional] [default to true]
**Enabled** | Pointer to **bool** | Enable or disable the group | [optional] [default to true]
**Properties** | Pointer to [**[]GroupPropertiesInner**](GroupPropertiesInner.md) | Group properties | [optional] 

## Methods

### NewGroupNew

`func NewGroupNew(category string, displayName string, ) *GroupNew`

NewGroupNew instantiates a new GroupNew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupNewWithDefaults

`func NewGroupNewWithDefaults() *GroupNew`

NewGroupNewWithDefaults instantiates a new GroupNew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *GroupNew) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GroupNew) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GroupNew) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GroupNew) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetCategory

`func (o *GroupNew) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GroupNew) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GroupNew) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetId

`func (o *GroupNew) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupNew) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupNew) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupNew) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *GroupNew) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GroupNew) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GroupNew) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *GroupNew) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupNew) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupNew) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupNew) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuery

`func (o *GroupNew) GetQuery() GroupNewQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GroupNew) GetQueryOk() (*GroupNewQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GroupNew) SetQuery(v GroupNewQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *GroupNew) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetDynamic

`func (o *GroupNew) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *GroupNew) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *GroupNew) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *GroupNew) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetEnabled

`func (o *GroupNew) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GroupNew) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GroupNew) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GroupNew) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProperties

`func (o *GroupNew) GetProperties() []GroupPropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GroupNew) GetPropertiesOk() (*[]GroupPropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GroupNew) SetProperties(v []GroupPropertiesInner)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *GroupNew) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


