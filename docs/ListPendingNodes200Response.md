# ListPendingNodes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**ListAcceptedNodes200ResponseData**](ListAcceptedNodes200ResponseData.md) |  | 

## Methods

### NewListPendingNodes200Response

`func NewListPendingNodes200Response(result string, action string, data ListAcceptedNodes200ResponseData, ) *ListPendingNodes200Response`

NewListPendingNodes200Response instantiates a new ListPendingNodes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPendingNodes200ResponseWithDefaults

`func NewListPendingNodes200ResponseWithDefaults() *ListPendingNodes200Response`

NewListPendingNodes200ResponseWithDefaults instantiates a new ListPendingNodes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ListPendingNodes200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListPendingNodes200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListPendingNodes200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *ListPendingNodes200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ListPendingNodes200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ListPendingNodes200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *ListPendingNodes200Response) GetData() ListAcceptedNodes200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPendingNodes200Response) GetDataOk() (*ListAcceptedNodes200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPendingNodes200Response) SetData(v ListAcceptedNodes200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


