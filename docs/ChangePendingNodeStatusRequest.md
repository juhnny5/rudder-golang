# ChangePendingNodeStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | New status of the pending node | [optional] 

## Methods

### NewChangePendingNodeStatusRequest

`func NewChangePendingNodeStatusRequest() *ChangePendingNodeStatusRequest`

NewChangePendingNodeStatusRequest instantiates a new ChangePendingNodeStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangePendingNodeStatusRequestWithDefaults

`func NewChangePendingNodeStatusRequestWithDefaults() *ChangePendingNodeStatusRequest`

NewChangePendingNodeStatusRequestWithDefaults instantiates a new ChangePendingNodeStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChangePendingNodeStatusRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChangePendingNodeStatusRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChangePendingNodeStatusRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChangePendingNodeStatusRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


