# ChangePendingNodeStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**ChangePendingNodeStatus200ResponseData**](ChangePendingNodeStatus200ResponseData.md) |  | 

## Methods

### NewChangePendingNodeStatus200Response

`func NewChangePendingNodeStatus200Response(result string, action string, data ChangePendingNodeStatus200ResponseData, ) *ChangePendingNodeStatus200Response`

NewChangePendingNodeStatus200Response instantiates a new ChangePendingNodeStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangePendingNodeStatus200ResponseWithDefaults

`func NewChangePendingNodeStatus200ResponseWithDefaults() *ChangePendingNodeStatus200Response`

NewChangePendingNodeStatus200ResponseWithDefaults instantiates a new ChangePendingNodeStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ChangePendingNodeStatus200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ChangePendingNodeStatus200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ChangePendingNodeStatus200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *ChangePendingNodeStatus200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ChangePendingNodeStatus200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ChangePendingNodeStatus200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *ChangePendingNodeStatus200Response) GetData() ChangePendingNodeStatus200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ChangePendingNodeStatus200Response) GetDataOk() (*ChangePendingNodeStatus200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ChangePendingNodeStatus200Response) SetData(v ChangePendingNodeStatus200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


