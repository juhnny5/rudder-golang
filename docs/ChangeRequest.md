# ChangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Acceptable** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Changes** | Pointer to [**ChangeRequestChanges**](ChangeRequestChanges.md) |  | [optional] 

## Methods

### NewChangeRequest

`func NewChangeRequest() *ChangeRequest`

NewChangeRequest instantiates a new ChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestWithDefaults

`func NewChangeRequestWithDefaults() *ChangeRequest`

NewChangeRequestWithDefaults instantiates a new ChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChangeRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangeRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangeRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ChangeRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ChangeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChangeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChangeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChangeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ChangeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChangeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChangeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChangeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *ChangeRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChangeRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChangeRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChangeRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAcceptable

`func (o *ChangeRequest) GetAcceptable() bool`

GetAcceptable returns the Acceptable field if non-nil, zero value otherwise.

### GetAcceptableOk

`func (o *ChangeRequest) GetAcceptableOk() (*bool, bool)`

GetAcceptableOk returns a tuple with the Acceptable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptable

`func (o *ChangeRequest) SetAcceptable(v bool)`

SetAcceptable sets Acceptable field to given value.

### HasAcceptable

`func (o *ChangeRequest) HasAcceptable() bool`

HasAcceptable returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ChangeRequest) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ChangeRequest) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ChangeRequest) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ChangeRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetChanges

`func (o *ChangeRequest) GetChanges() ChangeRequestChanges`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *ChangeRequest) GetChangesOk() (*ChangeRequestChanges, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *ChangeRequest) SetChanges(v ChangeRequestChanges)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *ChangeRequest) HasChanges() bool`

HasChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


