# GroupNewQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Select** | Pointer to **string** | What kind of data we want to include. Here we can get policy servers/relay by setting &#x60;nodeAndPolicyServer&#x60;. Only used if &#x60;where&#x60; is defined. | [optional] [default to "node"]
**Composition** | Pointer to **string** | Boolean operator to use between each  &#x60;where&#x60; criteria. | [optional] [default to "and"]
**Where** | Pointer to [**[]GroupQueryWhereInner**](GroupQueryWhereInner.md) | List of criteria | [optional] 

## Methods

### NewGroupNewQuery

`func NewGroupNewQuery() *GroupNewQuery`

NewGroupNewQuery instantiates a new GroupNewQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupNewQueryWithDefaults

`func NewGroupNewQueryWithDefaults() *GroupNewQuery`

NewGroupNewQueryWithDefaults instantiates a new GroupNewQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelect

`func (o *GroupNewQuery) GetSelect() string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *GroupNewQuery) GetSelectOk() (*string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *GroupNewQuery) SetSelect(v string)`

SetSelect sets Select field to given value.

### HasSelect

`func (o *GroupNewQuery) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### GetComposition

`func (o *GroupNewQuery) GetComposition() string`

GetComposition returns the Composition field if non-nil, zero value otherwise.

### GetCompositionOk

`func (o *GroupNewQuery) GetCompositionOk() (*string, bool)`

GetCompositionOk returns a tuple with the Composition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposition

`func (o *GroupNewQuery) SetComposition(v string)`

SetComposition sets Composition field to given value.

### HasComposition

`func (o *GroupNewQuery) HasComposition() bool`

HasComposition returns a boolean if a field has been set.

### GetWhere

`func (o *GroupNewQuery) GetWhere() []GroupQueryWhereInner`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *GroupNewQuery) GetWhereOk() (*[]GroupQueryWhereInner, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *GroupNewQuery) SetWhere(v []GroupQueryWhereInner)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *GroupNewQuery) HasWhere() bool`

HasWhere returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


