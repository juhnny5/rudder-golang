# AcceptChangeRequest200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**ListChangeRequests200ResponseData**](ListChangeRequests200ResponseData.md) |  | 

## Methods

### NewAcceptChangeRequest200Response

`func NewAcceptChangeRequest200Response(result string, action string, data ListChangeRequests200ResponseData, ) *AcceptChangeRequest200Response`

NewAcceptChangeRequest200Response instantiates a new AcceptChangeRequest200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptChangeRequest200ResponseWithDefaults

`func NewAcceptChangeRequest200ResponseWithDefaults() *AcceptChangeRequest200Response`

NewAcceptChangeRequest200ResponseWithDefaults instantiates a new AcceptChangeRequest200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *AcceptChangeRequest200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AcceptChangeRequest200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AcceptChangeRequest200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *AcceptChangeRequest200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AcceptChangeRequest200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AcceptChangeRequest200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *AcceptChangeRequest200Response) GetData() ListChangeRequests200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AcceptChangeRequest200Response) GetDataOk() (*ListChangeRequests200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AcceptChangeRequest200Response) SetData(v ListChangeRequests200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


