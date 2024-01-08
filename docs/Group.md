# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Group id | [optional] 
**DisplayName** | Pointer to **string** | Name of the group | [optional] 
**Description** | Pointer to **string** | Group description | [optional] 
**Query** | Pointer to [**GroupQuery**](GroupQuery.md) |  | [optional] 
**NodeIds** | Pointer to **[]string** | List of nodes in the group | [optional] 
**Dynamic** | Pointer to **bool** | Should the group be dynamically refreshed (if not, it is a static group) | [optional] [default to true]
**Enabled** | Pointer to **bool** | Enable or disable the group | [optional] [default to true]
**GroupClass** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to [**[]GroupPropertiesInner**](GroupPropertiesInner.md) | Group properties | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *Group) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Group) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Group) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Group) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Group) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Group) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Group) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuery

`func (o *Group) GetQuery() GroupQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Group) GetQueryOk() (*GroupQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Group) SetQuery(v GroupQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Group) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetNodeIds

`func (o *Group) GetNodeIds() []string`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *Group) GetNodeIdsOk() (*[]string, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *Group) SetNodeIds(v []string)`

SetNodeIds sets NodeIds field to given value.

### HasNodeIds

`func (o *Group) HasNodeIds() bool`

HasNodeIds returns a boolean if a field has been set.

### GetDynamic

`func (o *Group) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *Group) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *Group) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *Group) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetEnabled

`func (o *Group) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Group) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Group) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Group) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGroupClass

`func (o *Group) GetGroupClass() []string`

GetGroupClass returns the GroupClass field if non-nil, zero value otherwise.

### GetGroupClassOk

`func (o *Group) GetGroupClassOk() (*[]string, bool)`

GetGroupClassOk returns a tuple with the GroupClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClass

`func (o *Group) SetGroupClass(v []string)`

SetGroupClass sets GroupClass field to given value.

### HasGroupClass

`func (o *Group) HasGroupClass() bool`

HasGroupClass returns a boolean if a field has been set.

### GetProperties

`func (o *Group) GetProperties() []GroupPropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Group) GetPropertiesOk() (*[]GroupPropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Group) SetProperties(v []GroupPropertiesInner)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Group) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


