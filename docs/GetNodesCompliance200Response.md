# GetNodesCompliance200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**GetNodesCompliance200ResponseData**](GetNodesCompliance200ResponseData.md) |  | 

## Methods

### NewGetNodesCompliance200Response

`func NewGetNodesCompliance200Response(result string, action string, data GetNodesCompliance200ResponseData, ) *GetNodesCompliance200Response`

NewGetNodesCompliance200Response instantiates a new GetNodesCompliance200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNodesCompliance200ResponseWithDefaults

`func NewGetNodesCompliance200ResponseWithDefaults() *GetNodesCompliance200Response`

NewGetNodesCompliance200ResponseWithDefaults instantiates a new GetNodesCompliance200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetNodesCompliance200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNodesCompliance200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNodesCompliance200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *GetNodesCompliance200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetNodesCompliance200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetNodesCompliance200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *GetNodesCompliance200Response) GetData() GetNodesCompliance200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetNodesCompliance200Response) GetDataOk() (*GetNodesCompliance200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetNodesCompliance200Response) SetData(v GetNodesCompliance200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


